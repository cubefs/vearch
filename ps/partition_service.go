// Copyright 2019 The Vearch Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package ps

import (
	"context"
    "strconv"
	"sync"

	"github.com/tiglabs/raft"
	"github.com/tiglabs/raft/proto"
	"github.com/vearch/vearch/config"
	"github.com/vearch/vearch/proto/entity"
	"github.com/vearch/vearch/proto/vearchpb"
	"github.com/vearch/vearch/ps/engine"
	"github.com/vearch/vearch/ps/psutil"
    "github.com/vearch/vearch/util/fileutil"
	"github.com/vearch/vearch/ps/storage/raftstore"
	"github.com/vearch/vearch/util/log"
)

type Base interface {
	Start() error

	// Destroy close partition store if it running currently.
	Close() error

	// Destroy close partition store if it running currently and remove all data file from filesystem.
	Destroy() error

	// GetMeta returns meta information about this store.
	GetPartition() *entity.Partition

	//GetEngine return engine
	GetEngine() engine.Engine

	//space change API
	GetSpace() entity.Space

	// SetSpace
	SetSpace(space *entity.Space)
}

type Raft interface {
	GetLeader() (entity.NodeID, uint64)

	IsLeader() bool

	TryToLeader() error

	Status() *raft.Status

	GetVersion() uint64

	GetUnreachable(id uint64) []uint64

	ChangeMember(changeType proto.ConfChangeType, server *entity.Server) error
}

type PartitionStore interface {
	Base

	Raft

	UpdateSpace(ctx context.Context, space *entity.Space) error

	GetDocument(ctx context.Context, readLeader bool, doc *vearchpb.Document) (err error)

	Write(ctx context.Context, request *vearchpb.DocCmd, query *vearchpb.SearchRequest, response *vearchpb.SearchResponse) (err error)

	Flush(ctx context.Context) error

	Search(ctx context.Context, query *vearchpb.SearchRequest, response *vearchpb.SearchResponse) error
}

func (s *Server) GetPartition(id entity.PartitionID) (partition PartitionStore) {
	if p, _ := s.partitions.Load(id); p != nil {
		partition = p.(PartitionStore)
	}
	return
}

func (s *Server) RangePartition(fun func(entity.PartitionID, PartitionStore)) {

	s.partitions.Range(func(key, value interface{}) bool {
		fun(key.(entity.PartitionID), value.(PartitionStore))
		return true
	})
}

//load partition for in disk
func (s *Server) LoadPartition(ctx context.Context, pid entity.PartitionID) (PartitionStore, error) {

    meta_dir := config.Conf().GetDataDirBySlot(config.PS, pid);
	space, err := psutil.LoadPartitionMeta(meta_dir, pid)

	if err != nil {
		return nil, err
	}
// delete 
	store, err := raftstore.CreateStore(ctx, pid, s.nodeID, space, s.raftServer, s, s.client)
	if err != nil {
		return nil, err
	}

	if err := store.Start(); err != nil {
		return nil, err
	}

	s.partitions.Store(pid, store)

	return store, nil
}

func (s *Server) LoadFailPartition(ctx context.Context, pid entity.PartitionID, ip string) (PartitionStore, error) {

    source_dir := config.Conf().GetSourceDataDir()

    meta_dir := source_dir + "/" + ip

	space, err := psutil.LoadPartitionMeta(meta_dir, pid)

	if err != nil {
		return nil, err
	}

	store, err := raftstore.CreateFailStore(ctx, pid, s.nodeID, space, s.raftServer, s, s.client, ip)
	if err != nil {
		return nil, err
	}

	if err := store.Start(); err != nil {
		return nil, err
	}

	s.partitions.Store(pid, store)

	return store, nil
}

func (s *Server) CreatePartition(ctx context.Context, space *entity.Space, pid entity.PartitionID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	store, err := raftstore.CreateStore(ctx, pid, s.nodeID, space, s.raftServer, s, s.client)
	if err != nil {
		return err
	}
	store.RsStatusC = s.replicasStatusC
	if _, ok := s.partitions.LoadOrStore(pid, store); ok {
		log.Warn("partition is already exist partition id:[%d]", pid)
		if err := store.Close(); err != nil {
			log.Error("partitions close err : %s", err.Error())
		}
	} else {
        	for _, nodeId := range store.Partition.Replicas {
        		if server, err := s.client.Master().QueryServer(ctx, nodeId); err != nil {
        			log.Error("get server info err %s", err.Error())
        			return err
        			s.raftResolver.AddNode(nodeId, server.Replica())
        		}
        	}
		if err = store.Start(); err != nil {
			return err
		}
	}

	s.partitions.Store(pid, store)
	store.Partition.SetStatus(entity.PA_READWRITE)
        for _, nodeId := range store.Partition.Replicas {
		s.registerMaster(nodeId, pid)
        }
	return nil
}

func (s *Server) DeleteReplica(id entity.PartitionID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p, ok := s.partitions.Load(id); ok {
		s.partitions.Delete(id)
		if partition, is := p.(PartitionStore); is {
			s.raftResolver.DeleteNode(s.nodeID)
			if err := partition.Destroy(); err != nil {
				log.Error("delete partition[%v] fail cause: %v", id, err)
			}
		}

	}

	psutil.ClearPartition(config.Conf().GetDataDirBySlot(config.PS, id), id)
	log.Info("delete partition[%d] success", id)
}

func (s *Server) DeletePartition(id entity.PartitionID) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if p, ok := s.partitions.Load(id); ok {
		if partition, is := p.(PartitionStore); is {
			for _, r := range partition.GetPartition().Replicas {
				s.raftResolver.DeleteNode(r)
			}
			if err := partition.Destroy(); err != nil {
				log.Error("delete partition[%v] fail cause: %v", id, err)
				return
			}

			if partition.GetPartition().GetStatus() == entity.PA_INVALID {
				s.partitions.Delete(id)
			}
		}
	}

	psutil.ClearPartition(config.Conf().GetDataDirBySlot(config.PS, id), id)
	log.Info("delete partition:[%d] success", id)

	//delete partition cache
	if _, ok := s.partitions.Load(id); ok {
		s.partitions.Delete(id)
	}
}

func (s *Server) PartitionNum() int {
	var count int
	s.partitions.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (s *Server) ClosePartitions() {
	log.Info("ClosePartitions() invoked...")
	s.partitions.Range(func(key, value interface{}) bool {
		partition := value.(PartitionStore)
		err := partition.Close()
		if err != nil {
			log.Error("ClosePartitions() partition close has err: [%v]", err)
		}
		return true
	})
}

func (s *Server) recoverPartitions(pids []entity.PartitionID) {
	wg := new(sync.WaitGroup)
	wg.Add(len(pids))

	ctx := context.Background()

	// parallel execution recovery
	for i := 0; i < len(pids); i++ {
		idx := i
		go func(pid entity.PartitionID) {
			defer wg.Done()

			log.Debug("starting recover partition[%d]...", pid)
			_, err := s.LoadPartition(ctx, pid)
			if err != nil {
				log.Error("init partition err :[%s]", err.Error())
			}
			log.Debug("partition[%d] recovered complete", pid)
		}(pids[idx])
	}

	wg.Wait()
}

func (s *Server) recoverFailPartitions(ip string) ([]entity.PartitionID) {
// func (s *Server) recoverFailPartitions(ip string) {
	wg := new(sync.WaitGroup)

	ctx := context.Background()

    source_dir := config.Conf().GetSourceDataDir()

    fail_server_dir := source_dir + "/" + ip + "/meta"

    pids, _ := fileutil.GetAllDirsNames(fail_server_dir)

    PartitionIds :=  make([]entity.PartitionID, 0, len(pids))

	wg.Add(len(pids))

	log.Debug("load fail over pids size=%d fail_server_dir=%s source_dir=%s", len(pids), fail_server_dir, source_dir)

	// parallel execution recovery
	for i := 0; i < len(pids); i++ {
		idx := i
        partition, _ := strconv.ParseUint(pids[idx], 10, 64)
        // PartitionIds[idx] = uint32(partition)
        PartitionIds = append(PartitionIds, uint32(partition))
		go func(pid entity.PartitionID) {
			defer wg.Done()

			log.Debug("starting recover partition[%d]...", pid)
			_, err := s.LoadFailPartition(ctx, pid, ip)
			if err != nil {
				log.Error("WFQ init fail partition err :[%s]", err.Error())
			}
			log.Debug("partition[%d] recovered complete", pid)
		}(uint32(partition))
	}
	wg.Wait()
    return PartitionIds
}


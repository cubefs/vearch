// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc.proto

package pspb

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RpcRequestHead struct {
	TimeOutMs            int64    `protobuf:"varint,1,opt,name=time_out_ms,json=timeOutMs,proto3" json:"time_out_ms,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcRequestHead) Reset()      { *m = RpcRequestHead{} }
func (*RpcRequestHead) ProtoMessage() {}
func (*RpcRequestHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}
func (m *RpcRequestHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcRequestHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcRequestHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcRequestHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcRequestHead.Merge(m, src)
}
func (m *RpcRequestHead) XXX_Size() int {
	return m.Size()
}
func (m *RpcRequestHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcRequestHead.DiscardUnknown(m)
}

var xxx_messageInfo_RpcRequestHead proto.InternalMessageInfo

type RpcResponseHead struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Code                 int64    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcResponseHead) Reset()      { *m = RpcResponseHead{} }
func (*RpcResponseHead) ProtoMessage() {}
func (*RpcResponseHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}
func (m *RpcResponseHead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcResponseHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcResponseHead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcResponseHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcResponseHead.Merge(m, src)
}
func (m *RpcResponseHead) XXX_Size() int {
	return m.Size()
}
func (m *RpcResponseHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcResponseHead.DiscardUnknown(m)
}

var xxx_messageInfo_RpcResponseHead proto.InternalMessageInfo

type RpcSearchRequest struct {
	Head                 *RpcRequestHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	DbName               string          `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	SpaceName            string          `protobuf:"bytes,3,opt,name=space_name,json=spaceName,proto3" json:"space_name,omitempty"`
	Query                string          `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RpcSearchRequest) Reset()      { *m = RpcSearchRequest{} }
func (*RpcSearchRequest) ProtoMessage() {}
func (*RpcSearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}
func (m *RpcSearchRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcSearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcSearchRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcSearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcSearchRequest.Merge(m, src)
}
func (m *RpcSearchRequest) XXX_Size() int {
	return m.Size()
}
func (m *RpcSearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcSearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcSearchRequest proto.InternalMessageInfo

type RpcSearchResponse struct {
	Head                 *RpcResponseHead `protobuf:"bytes,1,opt,name=head,proto3" json:"head,omitempty"`
	Body                 string           `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RpcSearchResponse) Reset()      { *m = RpcSearchResponse{} }
func (*RpcSearchResponse) ProtoMessage() {}
func (*RpcSearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}
func (m *RpcSearchResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RpcSearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RpcSearchResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RpcSearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcSearchResponse.Merge(m, src)
}
func (m *RpcSearchResponse) XXX_Size() int {
	return m.Size()
}
func (m *RpcSearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcSearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcSearchResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RpcRequestHead)(nil), "RpcRequestHead")
	proto.RegisterType((*RpcResponseHead)(nil), "RpcResponseHead")
	proto.RegisterType((*RpcSearchRequest)(nil), "RpcSearchRequest")
	proto.RegisterType((*RpcSearchResponse)(nil), "RpcSearchResponse")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xb1, 0x8e, 0xd3, 0x40,
	0x10, 0x86, 0x77, 0x89, 0xcf, 0x9c, 0xe7, 0x24, 0xee, 0x6e, 0x75, 0x12, 0x56, 0x10, 0xab, 0x93,
	0xa1, 0xb8, 0x06, 0x47, 0x84, 0x32, 0x34, 0x50, 0xa5, 0x09, 0x48, 0x4b, 0x47, 0x13, 0xd9, 0xeb,
	0xc5, 0x71, 0xe1, 0xec, 0x66, 0xd7, 0x16, 0xa2, 0xa3, 0xe1, 0x1d, 0x78, 0x04, 0x1e, 0x81, 0x92,
	0x32, 0x25, 0x25, 0x65, 0x6c, 0x5e, 0x80, 0x92, 0x12, 0x65, 0x36, 0x91, 0x92, 0xd0, 0xcd, 0x3f,
	0xbf, 0x3d, 0xff, 0xf7, 0x4b, 0x0b, 0x50, 0x5a, 0x23, 0x53, 0x63, 0x75, 0xa3, 0x87, 0xcf, 0xca,
	0xaa, 0x59, 0xb4, 0x79, 0x2a, 0x75, 0x3d, 0x2a, 0x75, 0xa9, 0x47, 0xb8, 0xce, 0xdb, 0x0f, 0xa8,
	0x50, 0xe0, 0xe4, 0x3f, 0x4f, 0x2a, 0x78, 0x20, 0x8c, 0x14, 0x6a, 0xd5, 0x2a, 0xd7, 0x4c, 0x55,
	0x56, 0x30, 0x0e, 0x17, 0x4d, 0x55, 0xab, 0xb9, 0x6e, 0x9b, 0x79, 0xed, 0x62, 0x7a, 0x4b, 0xef,
	0x06, 0x22, 0xda, 0xae, 0xde, 0xb6, 0xcd, 0xcc, 0xb1, 0x47, 0x10, 0xb5, 0x4e, 0xd9, 0xf9, 0x32,
	0xab, 0x55, 0x7c, 0xef, 0x96, 0xde, 0x45, 0xe2, 0x7c, 0xbb, 0x78, 0x93, 0xd5, 0x8a, 0x0d, 0xe1,
	0xdc, 0x64, 0xce, 0x7d, 0xd4, 0xb6, 0x88, 0x07, 0xde, 0xdb, 0xeb, 0x64, 0x02, 0x97, 0x18, 0xe5,
	0x8c, 0x5e, 0x3a, 0x85, 0x59, 0x37, 0x70, 0xa6, 0xac, 0xd5, 0x16, 0x53, 0x22, 0xe1, 0x05, 0x63,
	0x10, 0x48, 0x5d, 0xf8, 0xe3, 0x03, 0x81, 0x73, 0xf2, 0x85, 0xc2, 0x95, 0x30, 0xf2, 0x9d, 0xca,
	0xac, 0x5c, 0xec, 0x70, 0xd9, 0x13, 0x08, 0x16, 0x2a, 0x2b, 0xf0, 0xef, 0x8b, 0xf1, 0x65, 0x7a,
	0xdc, 0x44, 0xa0, 0xc9, 0x1e, 0xc2, 0xfd, 0x22, 0x3f, 0xa4, 0x0d, 0x8b, 0x1c, 0x59, 0x1f, 0x03,
	0x38, 0x93, 0x49, 0xe5, 0x3d, 0x4f, 0x1b, 0xe1, 0x06, 0xed, 0x1b, 0x38, 0x5b, 0xb5, 0xca, 0x7e,
	0x8a, 0x03, 0xcf, 0x86, 0x22, 0x99, 0xc1, 0xf5, 0x01, 0x86, 0xaf, 0xc2, 0x9e, 0x1e, 0x71, 0x5c,
	0xa5, 0x27, 0x35, 0x77, 0x20, 0x0c, 0x82, 0x5c, 0x17, 0xfb, 0x7b, 0x38, 0x8f, 0x27, 0x10, 0x0a,
	0x23, 0x5f, 0x99, 0x8a, 0x3d, 0x87, 0xd0, 0x5f, 0x65, 0xd7, 0xe9, 0x69, 0xd1, 0x21, 0x4b, 0xff,
	0x0b, 0x4d, 0xc8, 0xeb, 0x97, 0xeb, 0x8e, 0x93, 0x5f, 0x1d, 0x27, 0x9b, 0x8e, 0x93, 0x3f, 0x1d,
	0x27, 0x7f, 0x3b, 0x4e, 0x3f, 0xf7, 0x9c, 0x7e, 0xeb, 0x39, 0xfd, 0xde, 0x73, 0xf2, 0xa3, 0xe7,
	0x64, 0xdd, 0x73, 0xfa, 0xb3, 0xe7, 0x74, 0xd3, 0x73, 0xfa, 0xf5, 0x37, 0x27, 0x53, 0xfa, 0x3e,
	0x30, 0xce, 0xe4, 0x79, 0x88, 0x0f, 0xe0, 0xc5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x7a,
	0xb3, 0x7e, 0x3d, 0x02, 0x00, 0x00,
}

func (this *RpcRequestHead) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RpcRequestHead)
	if !ok {
		that2, ok := that.(RpcRequestHead)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TimeOutMs != that1.TimeOutMs {
		return false
	}
	if this.UserName != that1.UserName {
		return false
	}
	if this.Password != that1.Password {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpcResponseHead) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RpcResponseHead)
	if !ok {
		that2, ok := that.(RpcResponseHead)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpcSearchRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RpcSearchRequest)
	if !ok {
		that2, ok := that.(RpcSearchRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Head.Equal(that1.Head) {
		return false
	}
	if this.DbName != that1.DbName {
		return false
	}
	if this.SpaceName != that1.SpaceName {
		return false
	}
	if this.Query != that1.Query {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RpcSearchResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RpcSearchResponse)
	if !ok {
		that2, ok := that.(RpcSearchResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Head.Equal(that1.Head) {
		return false
	}
	if this.Body != that1.Body {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RpcApiClient is the client API for RpcApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RpcApiClient interface {
	// Sends a greeting
	Search(ctx context.Context, in *RpcSearchRequest, opts ...grpc.CallOption) (*RpcSearchResponse, error)
}

type rpcApiClient struct {
	cc *grpc.ClientConn
}

func NewRpcApiClient(cc *grpc.ClientConn) RpcApiClient {
	return &rpcApiClient{cc}
}

func (c *rpcApiClient) Search(ctx context.Context, in *RpcSearchRequest, opts ...grpc.CallOption) (*RpcSearchResponse, error) {
	out := new(RpcSearchResponse)
	err := c.cc.Invoke(ctx, "/RpcApi/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcApiServer is the server API for RpcApi service.
type RpcApiServer interface {
	// Sends a greeting
	Search(context.Context, *RpcSearchRequest) (*RpcSearchResponse, error)
}

// UnimplementedRpcApiServer can be embedded to have forward compatible implementations.
type UnimplementedRpcApiServer struct {
}

func (*UnimplementedRpcApiServer) Search(ctx context.Context, req *RpcSearchRequest) (*RpcSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterRpcApiServer(s *grpc.Server, srv RpcApiServer) {
	s.RegisterService(&_RpcApi_serviceDesc, srv)
}

func _RpcApi_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RpcSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcApiServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RpcApi/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcApiServer).Search(ctx, req.(*RpcSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RpcApi",
	HandlerType: (*RpcApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _RpcApi_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

func (m *RpcRequestHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcRequestHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcRequestHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UserName) > 0 {
		i -= len(m.UserName)
		copy(dAtA[i:], m.UserName)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.UserName)))
		i--
		dAtA[i] = 0x12
	}
	if m.TimeOutMs != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.TimeOutMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RpcResponseHead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcResponseHead) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcResponseHead) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintGrpc(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Error) > 0 {
		i -= len(m.Error)
		copy(dAtA[i:], m.Error)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Error)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcSearchRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcSearchRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcSearchRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Query) > 0 {
		i -= len(m.Query)
		copy(dAtA[i:], m.Query)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Query)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SpaceName) > 0 {
		i -= len(m.SpaceName)
		copy(dAtA[i:], m.SpaceName)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.SpaceName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DbName) > 0 {
		i -= len(m.DbName)
		copy(dAtA[i:], m.DbName)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.DbName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RpcSearchResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RpcSearchResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RpcSearchResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintGrpc(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if m.Head != nil {
		{
			size, err := m.Head.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGrpc(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedRpcRequestHead(r randyGrpc, easy bool) *RpcRequestHead {
	this := &RpcRequestHead{}
	this.TimeOutMs = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.TimeOutMs *= -1
	}
	this.UserName = string(randStringGrpc(r))
	this.Password = string(randStringGrpc(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedGrpc(r, 4)
	}
	return this
}

func NewPopulatedRpcResponseHead(r randyGrpc, easy bool) *RpcResponseHead {
	this := &RpcResponseHead{}
	this.Error = string(randStringGrpc(r))
	this.Code = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Code *= -1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedGrpc(r, 3)
	}
	return this
}

func NewPopulatedRpcSearchRequest(r randyGrpc, easy bool) *RpcSearchRequest {
	this := &RpcSearchRequest{}
	if r.Intn(5) != 0 {
		this.Head = NewPopulatedRpcRequestHead(r, easy)
	}
	this.DbName = string(randStringGrpc(r))
	this.SpaceName = string(randStringGrpc(r))
	this.Query = string(randStringGrpc(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedGrpc(r, 5)
	}
	return this
}

func NewPopulatedRpcSearchResponse(r randyGrpc, easy bool) *RpcSearchResponse {
	this := &RpcSearchResponse{}
	if r.Intn(5) != 0 {
		this.Head = NewPopulatedRpcResponseHead(r, easy)
	}
	this.Body = string(randStringGrpc(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedGrpc(r, 5)
	}
	return this
}

type randyGrpc interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneGrpc(r randyGrpc) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringGrpc(r randyGrpc) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneGrpc(r)
	}
	return string(tmps)
}
func randUnrecognizedGrpc(r randyGrpc, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldGrpc(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldGrpc(dAtA []byte, r randyGrpc, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateGrpc(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateGrpc(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *RpcRequestHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeOutMs != 0 {
		n += 1 + sovGrpc(uint64(m.TimeOutMs))
	}
	l = len(m.UserName)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RpcResponseHead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.Code != 0 {
		n += 1 + sovGrpc(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RpcSearchRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.DbName)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.SpaceName)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RpcSearchResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Head != nil {
		l = m.Head.Size()
		n += 1 + l + sovGrpc(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovGrpc(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGrpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpc(x uint64) (n int) {
	return sovGrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RpcRequestHead) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RpcRequestHead{`,
		`TimeOutMs:` + fmt.Sprintf("%v", this.TimeOutMs) + `,`,
		`UserName:` + fmt.Sprintf("%v", this.UserName) + `,`,
		`Password:` + fmt.Sprintf("%v", this.Password) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RpcResponseHead) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RpcResponseHead{`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RpcSearchRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RpcSearchRequest{`,
		`Head:` + strings.Replace(this.Head.String(), "RpcRequestHead", "RpcRequestHead", 1) + `,`,
		`DbName:` + fmt.Sprintf("%v", this.DbName) + `,`,
		`SpaceName:` + fmt.Sprintf("%v", this.SpaceName) + `,`,
		`Query:` + fmt.Sprintf("%v", this.Query) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RpcSearchResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RpcSearchResponse{`,
		`Head:` + strings.Replace(this.Head.String(), "RpcResponseHead", "RpcResponseHead", 1) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGrpc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RpcRequestHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RpcRequestHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcRequestHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeOutMs", wireType)
			}
			m.TimeOutMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeOutMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcResponseHead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RpcResponseHead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcResponseHead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcSearchRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RpcSearchRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcSearchRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &RpcRequestHead{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DbName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DbName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RpcSearchResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RpcSearchResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RpcSearchResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Head", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Head == nil {
				m.Head = &RpcResponseHead{}
			}
			if err := m.Head.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpc
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpc
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGrpc
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGrpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrpc = fmt.Errorf("proto: unexpected end of group")
)

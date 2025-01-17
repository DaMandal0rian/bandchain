// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bandchain/v1/node/query.proto

package node

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryChainIDRequest is request type for the Service/ChainID RPC method.
type QueryChainIDRequest struct {
}

func (m *QueryChainIDRequest) Reset()         { *m = QueryChainIDRequest{} }
func (m *QueryChainIDRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChainIDRequest) ProtoMessage()    {}
func (*QueryChainIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25c18d173830f62, []int{0}
}
func (m *QueryChainIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainIDRequest.Merge(m, src)
}
func (m *QueryChainIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainIDRequest proto.InternalMessageInfo

// QueryChainIDResponse is response type for the Service/ChainID RPC method.
type QueryChainIDResponse struct {
	ChainID string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *QueryChainIDResponse) Reset()         { *m = QueryChainIDResponse{} }
func (m *QueryChainIDResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChainIDResponse) ProtoMessage()    {}
func (*QueryChainIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25c18d173830f62, []int{1}
}
func (m *QueryChainIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainIDResponse.Merge(m, src)
}
func (m *QueryChainIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainIDResponse proto.InternalMessageInfo

func (m *QueryChainIDResponse) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

// QueryEVMValidatorsRequest is request type for the Service/EVMValidators RPC method.
type QueryEVMValidatorsRequest struct {
}

func (m *QueryEVMValidatorsRequest) Reset()         { *m = QueryEVMValidatorsRequest{} }
func (m *QueryEVMValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEVMValidatorsRequest) ProtoMessage()    {}
func (*QueryEVMValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25c18d173830f62, []int{2}
}
func (m *QueryEVMValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEVMValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEVMValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEVMValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEVMValidatorsRequest.Merge(m, src)
}
func (m *QueryEVMValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEVMValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEVMValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEVMValidatorsRequest proto.InternalMessageInfo

// QueryEVMValidatorsResponse is response type for the Service/EVMValidators RPC method.
type QueryEVMValidatorsResponse struct {
	// BlockHeight is the latest block height
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Validators is list of validator's addresss and voting power
	Validators []ValidatorMinimal `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators"`
}

func (m *QueryEVMValidatorsResponse) Reset()         { *m = QueryEVMValidatorsResponse{} }
func (m *QueryEVMValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEVMValidatorsResponse) ProtoMessage()    {}
func (*QueryEVMValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25c18d173830f62, []int{3}
}
func (m *QueryEVMValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEVMValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEVMValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEVMValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEVMValidatorsResponse.Merge(m, src)
}
func (m *QueryEVMValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEVMValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEVMValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEVMValidatorsResponse proto.InternalMessageInfo

func (m *QueryEVMValidatorsResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *QueryEVMValidatorsResponse) GetValidators() []ValidatorMinimal {
	if m != nil {
		return m.Validators
	}
	return nil
}

// ValidatorMinimal is the data structure for storing validator's address and voting power
type ValidatorMinimal struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	VotingPower int64  `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *ValidatorMinimal) Reset()         { *m = ValidatorMinimal{} }
func (m *ValidatorMinimal) String() string { return proto.CompactTextString(m) }
func (*ValidatorMinimal) ProtoMessage()    {}
func (*ValidatorMinimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a25c18d173830f62, []int{4}
}
func (m *ValidatorMinimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMinimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMinimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMinimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMinimal.Merge(m, src)
}
func (m *ValidatorMinimal) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMinimal) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMinimal.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMinimal proto.InternalMessageInfo

func (m *ValidatorMinimal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorMinimal) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryChainIDRequest)(nil), "bandchain.v1.node.QueryChainIDRequest")
	proto.RegisterType((*QueryChainIDResponse)(nil), "bandchain.v1.node.QueryChainIDResponse")
	proto.RegisterType((*QueryEVMValidatorsRequest)(nil), "bandchain.v1.node.QueryEVMValidatorsRequest")
	proto.RegisterType((*QueryEVMValidatorsResponse)(nil), "bandchain.v1.node.QueryEVMValidatorsResponse")
	proto.RegisterType((*ValidatorMinimal)(nil), "bandchain.v1.node.ValidatorMinimal")
}

func init() { proto.RegisterFile("bandchain/v1/node/query.proto", fileDescriptor_a25c18d173830f62) }

var fileDescriptor_a25c18d173830f62 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0x3a, 0x44, 0xc1, 0x05, 0x09, 0xcc, 0x40, 0x21, 0x0c, 0x6f, 0x04, 0x34, 0x76, 0x60,
	0xb1, 0x5a, 0xee, 0x1c, 0x0a, 0x48, 0x54, 0x68, 0x02, 0x8a, 0xb4, 0x03, 0x97, 0xca, 0x49, 0x2c,
	0xc7, 0x22, 0xf5, 0xcb, 0x12, 0x37, 0x08, 0x8e, 0x1c, 0x39, 0x21, 0xb8, 0xf1, 0x8b, 0x76, 0x9c,
	0xc4, 0x85, 0x53, 0x85, 0x52, 0x7e, 0x08, 0x8a, 0xd3, 0x14, 0x02, 0x41, 0xec, 0xe6, 0x7c, 0xef,
	0x7b, 0xdf, 0xfb, 0xf2, 0xf9, 0x19, 0xdd, 0xf4, 0x99, 0x0a, 0x83, 0x88, 0x49, 0x45, 0xf3, 0x01,
	0x55, 0x10, 0x72, 0x7a, 0x34, 0xe7, 0xe9, 0x5b, 0x2f, 0x49, 0x41, 0x03, 0xbe, 0xbc, 0x2e, 0x7b,
	0xf9, 0xc0, 0x2b, 0xcb, 0xce, 0x96, 0x00, 0x10, 0x31, 0xa7, 0x2c, 0x91, 0x94, 0x29, 0x05, 0x9a,
	0x69, 0x09, 0x2a, 0xab, 0x1a, 0x9c, 0x4d, 0x01, 0x02, 0xcc, 0x91, 0x96, 0xa7, 0x0a, 0x75, 0xaf,
	0xa2, 0x2b, 0x2f, 0x4a, 0xd5, 0x87, 0xa5, 0xd2, 0xf8, 0xd1, 0x84, 0x1f, 0xcd, 0x79, 0xa6, 0xdd,
	0x07, 0x68, 0xb3, 0x09, 0x67, 0x09, 0xa8, 0x8c, 0xe3, 0x5d, 0x74, 0xce, 0xcc, 0x9c, 0xca, 0xd0,
	0xb6, 0x76, 0xac, 0xbd, 0xf3, 0xa3, 0x7e, 0xb1, 0xd8, 0xee, 0xd5, 0xb4, 0x9e, 0x29, 0x8e, 0x43,
	0xf7, 0x06, 0xba, 0x6e, 0xfa, 0x1f, 0x1f, 0x1e, 0x1c, 0xb2, 0x58, 0x86, 0x4c, 0x43, 0x9a, 0xd5,
	0xe2, 0x1f, 0x2c, 0xe4, 0xb4, 0x55, 0x57, 0x33, 0x6e, 0xa1, 0x0b, 0x7e, 0x0c, 0xc1, 0xeb, 0x69,
	0xc4, 0xa5, 0x88, 0xb4, 0x99, 0xb3, 0x31, 0xe9, 0x1b, 0xec, 0x89, 0x81, 0xf0, 0x18, 0xa1, 0x7c,
	0xdd, 0x68, 0x77, 0x77, 0x36, 0xf6, 0xfa, 0xc3, 0xdb, 0xde, 0x5f, 0x89, 0x78, 0x6b, 0xf5, 0x03,
	0xa9, 0xe4, 0x8c, 0xc5, 0xa3, 0x33, 0xc7, 0x8b, 0xed, 0xce, 0xe4, 0xb7, 0x66, 0xf7, 0x19, 0xba,
	0xf4, 0x27, 0x0b, 0xdb, 0xa8, 0xc7, 0xc2, 0x30, 0xe5, 0x59, 0x56, 0xfd, 0xe4, 0xa4, 0xfe, 0x2c,
	0xbd, 0xe5, 0xa0, 0xa5, 0x12, 0xd3, 0x04, 0xde, 0xf0, 0xd4, 0xee, 0x56, 0xde, 0x2a, 0xec, 0x79,
	0x09, 0x0d, 0xbf, 0x74, 0x51, 0xef, 0x25, 0x4f, 0x73, 0x19, 0x70, 0xfc, 0x0e, 0xd5, 0xd1, 0xe0,
	0xdd, 0x16, 0x7b, 0x2d, 0xc9, 0x3b, 0x77, 0xff, 0xcb, 0xab, 0x62, 0x72, 0xc9, 0xfb, 0xaf, 0x3f,
	0x3e, 0x77, 0x6d, 0x7c, 0x8d, 0x36, 0x16, 0xa5, 0xbe, 0x1e, 0xfc, 0xc9, 0x42, 0x17, 0x1b, 0x01,
	0xe3, 0x7b, 0xff, 0x92, 0x6e, 0xbb, 0x25, 0x67, 0xff, 0x94, 0xec, 0x95, 0x9d, 0x3b, 0xc6, 0x0e,
	0xc1, 0x5b, 0x4d, 0x3b, 0x3c, 0x9f, 0xed, 0xff, 0x4a, 0x7b, 0xf4, 0xf4, 0xb8, 0x20, 0xd6, 0x49,
	0x41, 0xac, 0xef, 0x05, 0xb1, 0x3e, 0x2e, 0x49, 0xe7, 0x64, 0x49, 0x3a, 0xdf, 0x96, 0xa4, 0xf3,
	0x6a, 0x20, 0xa4, 0x8e, 0xe6, 0xbe, 0x17, 0xc0, 0xcc, 0x28, 0x98, 0xf5, 0x0c, 0x20, 0xa6, 0x2b,
	0xa9, 0x21, 0x0d, 0x62, 0xc9, 0x95, 0xa6, 0x22, 0x4d, 0x02, 0xf3, 0x1c, 0xfc, 0xb3, 0x86, 0x73,
	0xff, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x13, 0x94, 0xa2, 0x2a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// ChainID queries the chain ID of this node
	ChainID(ctx context.Context, in *QueryChainIDRequest, opts ...grpc.CallOption) (*QueryChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(ctx context.Context, in *QueryEVMValidatorsRequest, opts ...grpc.CallOption) (*QueryEVMValidatorsResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ChainID(ctx context.Context, in *QueryChainIDRequest, opts ...grpc.CallOption) (*QueryChainIDResponse, error) {
	out := new(QueryChainIDResponse)
	err := c.cc.Invoke(ctx, "/bandchain.v1.node.Service/ChainID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) EVMValidators(ctx context.Context, in *QueryEVMValidatorsRequest, opts ...grpc.CallOption) (*QueryEVMValidatorsResponse, error) {
	out := new(QueryEVMValidatorsResponse)
	err := c.cc.Invoke(ctx, "/bandchain.v1.node.Service/EVMValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// ChainID queries the chain ID of this node
	ChainID(context.Context, *QueryChainIDRequest) (*QueryChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(context.Context, *QueryEVMValidatorsRequest) (*QueryEVMValidatorsResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) ChainID(ctx context.Context, req *QueryChainIDRequest) (*QueryChainIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainID not implemented")
}
func (*UnimplementedServiceServer) EVMValidators(ctx context.Context, req *QueryEVMValidatorsRequest) (*QueryEVMValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EVMValidators not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_ChainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ChainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bandchain.v1.node.Service/ChainID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ChainID(ctx, req.(*QueryChainIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_EVMValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEVMValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).EVMValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bandchain.v1.node.Service/EVMValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).EVMValidators(ctx, req.(*QueryEVMValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bandchain.v1.node.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainID",
			Handler:    _Service_ChainID_Handler,
		},
		{
			MethodName: "EVMValidators",
			Handler:    _Service_EVMValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bandchain/v1/node/query.proto",
}

func (m *QueryChainIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryChainIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEVMValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEVMValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEVMValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEVMValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEVMValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEVMValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMinimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMinimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMinimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryChainIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryChainIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEVMValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEVMValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *ValidatorMinimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovQuery(uint64(m.VotingPower))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryChainIDRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChainIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryChainIDResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryChainIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEVMValidatorsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEVMValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEVMValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEVMValidatorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEVMValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEVMValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, ValidatorMinimal{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ValidatorMinimal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ValidatorMinimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMinimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

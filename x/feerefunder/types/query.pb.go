// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/feerefunder/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20b5686ec46d4e6, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20b5686ec46d4e6, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type FeeInfoRequest struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	PortId    string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *FeeInfoRequest) Reset()         { *m = FeeInfoRequest{} }
func (m *FeeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*FeeInfoRequest) ProtoMessage()    {}
func (*FeeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20b5686ec46d4e6, []int{2}
}
func (m *FeeInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeInfoRequest.Merge(m, src)
}
func (m *FeeInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *FeeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FeeInfoRequest proto.InternalMessageInfo

func (m *FeeInfoRequest) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *FeeInfoRequest) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *FeeInfoRequest) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type FeeInfoResponse struct {
	FeeInfo *FeeInfo `protobuf:"bytes,1,opt,name=fee_info,json=feeInfo,proto3" json:"fee_info,omitempty"`
}

func (m *FeeInfoResponse) Reset()         { *m = FeeInfoResponse{} }
func (m *FeeInfoResponse) String() string { return proto.CompactTextString(m) }
func (*FeeInfoResponse) ProtoMessage()    {}
func (*FeeInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c20b5686ec46d4e6, []int{3}
}
func (m *FeeInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeInfoResponse.Merge(m, src)
}
func (m *FeeInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *FeeInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FeeInfoResponse proto.InternalMessageInfo

func (m *FeeInfoResponse) GetFeeInfo() *FeeInfo {
	if m != nil {
		return m.FeeInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "neutron.feerefunder.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "neutron.feerefunder.QueryParamsResponse")
	proto.RegisterType((*FeeInfoRequest)(nil), "neutron.feerefunder.FeeInfoRequest")
	proto.RegisterType((*FeeInfoResponse)(nil), "neutron.feerefunder.FeeInfoResponse")
}

func init() { proto.RegisterFile("neutron/feerefunder/query.proto", fileDescriptor_c20b5686ec46d4e6) }

var fileDescriptor_c20b5686ec46d4e6 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6b, 0xd4, 0x40,
	0x14, 0xc6, 0x77, 0xd6, 0xba, 0xdb, 0x8e, 0xa0, 0x30, 0x2d, 0x58, 0xd6, 0x9a, 0xae, 0x51, 0xd9,
	0x55, 0x68, 0x86, 0xae, 0x88, 0x78, 0xed, 0x41, 0x58, 0x2f, 0xd6, 0x1c, 0xbd, 0x94, 0x49, 0xf2,
	0x32, 0x0d, 0x74, 0xe7, 0xa5, 0x33, 0x93, 0x62, 0xaf, 0x7a, 0xf5, 0x20, 0x78, 0xf5, 0x0f, 0xea,
	0xb1, 0xe0, 0xc5, 0x93, 0xc8, 0xae, 0x7f, 0x88, 0x64, 0x32, 0x2d, 0x16, 0xd3, 0xee, 0x6d, 0xe6,
	0xbd, 0xef, 0xfb, 0xe6, 0xf7, 0x5e, 0x42, 0xb7, 0x15, 0x54, 0x56, 0xa3, 0xe2, 0x39, 0x80, 0x86,
	0xbc, 0x52, 0x19, 0x68, 0x7e, 0x5c, 0x81, 0x3e, 0x8d, 0x4a, 0x8d, 0x16, 0xd9, 0xba, 0x17, 0x44,
	0xff, 0x08, 0x06, 0xcf, 0x53, 0x34, 0x33, 0x34, 0x3c, 0x11, 0x06, 0x1a, 0x35, 0x3f, 0xd9, 0x4d,
	0xc0, 0x8a, 0x5d, 0x5e, 0x0a, 0x59, 0x28, 0x61, 0x0b, 0x54, 0x4d, 0xc0, 0x60, 0x43, 0xa2, 0x44,
	0x77, 0xe4, 0xf5, 0xc9, 0x57, 0xb7, 0x24, 0xa2, 0x3c, 0x02, 0x2e, 0xca, 0x82, 0x0b, 0xa5, 0xd0,
	0x3a, 0x8b, 0xf1, 0xdd, 0x47, 0x6d, 0x54, 0x12, 0x14, 0x98, 0xe2, 0x42, 0x32, 0x6c, 0x93, 0x94,
	0x42, 0x8b, 0x99, 0x57, 0x84, 0x1b, 0x94, 0xbd, 0xaf, 0xd1, 0xf6, 0x5d, 0x31, 0x86, 0xe3, 0x0a,
	0x8c, 0x0d, 0xf7, 0xe9, 0xfa, 0x95, 0xaa, 0x29, 0x51, 0x19, 0x60, 0xaf, 0x69, 0xaf, 0x31, 0x6f,
	0x92, 0x21, 0x19, 0xdf, 0x99, 0x3c, 0x88, 0x5a, 0xe6, 0x8e, 0x1a, 0xd3, 0xde, 0xca, 0xd9, 0xaf,
	0xed, 0x4e, 0xec, 0x0d, 0x61, 0x46, 0xef, 0xbe, 0x01, 0x98, 0xaa, 0x1c, 0xfd, 0x1b, 0xec, 0x21,
	0xa5, 0xe9, 0xa1, 0x50, 0x0a, 0x8e, 0x0e, 0x8a, 0xcc, 0x05, 0xae, 0xc5, 0x6b, 0xbe, 0x32, 0xcd,
	0xd8, 0x7d, 0xda, 0x2f, 0x51, 0xdb, 0xba, 0xd7, 0x75, 0xbd, 0x5e, 0x7d, 0x9d, 0x66, 0x6c, 0x40,
	0x57, 0x4d, 0x1d, 0xa1, 0x52, 0xd8, 0xbc, 0x35, 0x24, 0xe3, 0x95, 0xf8, 0xf2, 0x1e, 0xbe, 0xa5,
	0xf7, 0x2e, 0x5f, 0xf1, 0xcc, 0xaf, 0xe8, 0x6a, 0x0e, 0x70, 0x50, 0xa8, 0x1c, 0x3d, 0xf5, 0x56,
	0x2b, 0xf5, 0x85, 0xaf, 0x9f, 0x37, 0x87, 0xc9, 0xf7, 0x2e, 0xbd, 0xed, 0x96, 0xc0, 0xbe, 0x10,
	0xda, 0x6b, 0x86, 0x62, 0xa3, 0x56, 0xef, 0xff, 0x1b, 0x1c, 0x8c, 0x97, 0x0b, 0x1b, 0xc0, 0x90,
	0x7f, 0xfa, 0xf1, 0xe7, 0x5b, 0xf7, 0x19, 0x1b, 0x71, 0xef, 0xd8, 0x41, 0x2d, 0xf9, 0xf5, 0x1f,
	0x8e, 0x7d, 0x26, 0xb4, 0xef, 0x69, 0xd9, 0xe3, 0x1b, 0x67, 0xf1, 0x2c, 0x4f, 0x6e, 0x16, 0x79,
	0x8e, 0x1d, 0xc7, 0x31, 0x62, 0x4f, 0x97, 0x72, 0xd4, 0xbb, 0xdc, 0x7b, 0x77, 0x36, 0x0f, 0xc8,
	0xf9, 0x3c, 0x20, 0xbf, 0xe7, 0x01, 0xf9, 0xba, 0x08, 0x3a, 0xe7, 0x8b, 0xa0, 0xf3, 0x73, 0x11,
	0x74, 0x3e, 0xbc, 0x94, 0x85, 0x3d, 0xac, 0x92, 0x28, 0xc5, 0x59, 0x6b, 0xd4, 0xc9, 0x84, 0x7f,
	0xbc, 0x92, 0x67, 0x4f, 0x4b, 0x30, 0x49, 0xcf, 0xfd, 0x90, 0x2f, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x22, 0xe1, 0xfd, 0x6d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	FeeInfo(ctx context.Context, in *FeeInfoRequest, opts ...grpc.CallOption) (*FeeInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/neutron.feerefunder.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeeInfo(ctx context.Context, in *FeeInfoRequest, opts ...grpc.CallOption) (*FeeInfoResponse, error) {
	out := new(FeeInfoResponse)
	err := c.cc.Invoke(ctx, "/neutron.feerefunder.Query/FeeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	FeeInfo(context.Context, *FeeInfoRequest) (*FeeInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) FeeInfo(ctx context.Context, req *FeeInfoRequest) (*FeeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.feerefunder.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neutron.feerefunder.Query/FeeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeInfo(ctx, req.(*FeeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "neutron.feerefunder.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "FeeInfo",
			Handler:    _Query_FeeInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "neutron/feerefunder/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FeeInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FeeInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeInfo != nil {
		{
			size, err := m.FeeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *FeeInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovQuery(uint64(m.Sequence))
	}
	return n
}

func (m *FeeInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeInfo != nil {
		l = m.FeeInfo.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *FeeInfoRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FeeInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
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
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
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
func (m *FeeInfoResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: FeeInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeInfo", wireType)
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
			if m.FeeInfo == nil {
				m.FeeInfo = &FeeInfo{}
			}
			if err := m.FeeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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

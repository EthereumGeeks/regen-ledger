// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1alpha1/tx.proto

package data

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type MsgAnchorDataRequest struct {
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Cid    []byte `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *MsgAnchorDataRequest) Reset()         { *m = MsgAnchorDataRequest{} }
func (m *MsgAnchorDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAnchorDataRequest) ProtoMessage()    {}
func (*MsgAnchorDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{0}
}
func (m *MsgAnchorDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchorDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchorDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchorDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchorDataRequest.Merge(m, src)
}
func (m *MsgAnchorDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchorDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchorDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchorDataRequest proto.InternalMessageInfo

func (m *MsgAnchorDataRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgAnchorDataRequest) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

type MsgAnchorDataResponse struct {
	Timestamp *types.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MsgAnchorDataResponse) Reset()         { *m = MsgAnchorDataResponse{} }
func (m *MsgAnchorDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAnchorDataResponse) ProtoMessage()    {}
func (*MsgAnchorDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{1}
}
func (m *MsgAnchorDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAnchorDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAnchorDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAnchorDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAnchorDataResponse.Merge(m, src)
}
func (m *MsgAnchorDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAnchorDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAnchorDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAnchorDataResponse proto.InternalMessageInfo

func (m *MsgAnchorDataResponse) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type MsgSignDataRequest struct {
	Signers []string `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	Cid     []byte   `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *MsgSignDataRequest) Reset()         { *m = MsgSignDataRequest{} }
func (m *MsgSignDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSignDataRequest) ProtoMessage()    {}
func (*MsgSignDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{2}
}
func (m *MsgSignDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignDataRequest.Merge(m, src)
}
func (m *MsgSignDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignDataRequest proto.InternalMessageInfo

type MsgSignDataResponse struct {
}

func (m *MsgSignDataResponse) Reset()         { *m = MsgSignDataResponse{} }
func (m *MsgSignDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSignDataResponse) ProtoMessage()    {}
func (*MsgSignDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{3}
}
func (m *MsgSignDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSignDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSignDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSignDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSignDataResponse.Merge(m, src)
}
func (m *MsgSignDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSignDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSignDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSignDataResponse proto.InternalMessageInfo

type MsgStoreDataRequest struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Cid     []byte `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
	Content []byte `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *MsgStoreDataRequest) Reset()         { *m = MsgStoreDataRequest{} }
func (m *MsgStoreDataRequest) String() string { return proto.CompactTextString(m) }
func (*MsgStoreDataRequest) ProtoMessage()    {}
func (*MsgStoreDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{4}
}
func (m *MsgStoreDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreDataRequest.Merge(m, src)
}
func (m *MsgStoreDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreDataRequest proto.InternalMessageInfo

func (m *MsgStoreDataRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStoreDataRequest) GetCid() []byte {
	if m != nil {
		return m.Cid
	}
	return nil
}

func (m *MsgStoreDataRequest) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type MsgStoreDataResponse struct {
}

func (m *MsgStoreDataResponse) Reset()         { *m = MsgStoreDataResponse{} }
func (m *MsgStoreDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStoreDataResponse) ProtoMessage()    {}
func (*MsgStoreDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5051179b7de7ca1d, []int{5}
}
func (m *MsgStoreDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStoreDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStoreDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStoreDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStoreDataResponse.Merge(m, src)
}
func (m *MsgStoreDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStoreDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStoreDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStoreDataResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAnchorDataRequest)(nil), "regen.data.v1alpha1.MsgAnchorDataRequest")
	proto.RegisterType((*MsgAnchorDataResponse)(nil), "regen.data.v1alpha1.MsgAnchorDataResponse")
	proto.RegisterType((*MsgSignDataRequest)(nil), "regen.data.v1alpha1.MsgSignDataRequest")
	proto.RegisterType((*MsgSignDataResponse)(nil), "regen.data.v1alpha1.MsgSignDataResponse")
	proto.RegisterType((*MsgStoreDataRequest)(nil), "regen.data.v1alpha1.MsgStoreDataRequest")
	proto.RegisterType((*MsgStoreDataResponse)(nil), "regen.data.v1alpha1.MsgStoreDataResponse")
}

func init() { proto.RegisterFile("regen/data/v1alpha1/tx.proto", fileDescriptor_5051179b7de7ca1d) }

var fileDescriptor_5051179b7de7ca1d = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x18, 0x4c, 0x1a, 0x54, 0x88, 0xe1, 0x80, 0xdc, 0x1f, 0x45, 0x11, 0x4a, 0xab, 0x5c, 0x48, 0x11,
	0xd8, 0x6a, 0xb9, 0x20, 0x4e, 0x80, 0x50, 0x6f, 0x3d, 0x10, 0xb8, 0x80, 0x84, 0x44, 0x9a, 0x18,
	0x37, 0xa2, 0xb5, 0x83, 0xed, 0x42, 0x1f, 0x81, 0x23, 0x8f, 0xb0, 0xf7, 0x7d, 0x91, 0x3d, 0xf6,
	0xb8, 0xc7, 0x55, 0xfb, 0x22, 0xab, 0x38, 0x49, 0xdb, 0xed, 0xb6, 0xda, 0x6a, 0x6f, 0x1e, 0x7f,
	0xe3, 0xf9, 0xe6, 0x9b, 0x2f, 0x01, 0xcf, 0x04, 0xa1, 0x84, 0xe1, 0x24, 0x52, 0x11, 0xfe, 0xd3,
	0x8f, 0xa6, 0xd9, 0x24, 0xea, 0x63, 0xb5, 0x40, 0x99, 0xe0, 0x8a, 0xc3, 0x86, 0xae, 0xa2, 0xbc,
	0x8a, 0xaa, 0xaa, 0xdb, 0xa4, 0x9c, 0x72, 0x5d, 0xc7, 0xf9, 0xa9, 0xa0, 0xba, 0x1d, 0xca, 0x39,
	0x9d, 0x12, 0xac, 0xd1, 0x78, 0xfe, 0x13, 0xab, 0x74, 0x46, 0xa4, 0x8a, 0x66, 0x59, 0x41, 0xf0,
	0xdf, 0x81, 0xe6, 0x48, 0xd2, 0xf7, 0x2c, 0x9e, 0x70, 0xf1, 0x31, 0x52, 0x51, 0x48, 0x7e, 0xcf,
	0x89, 0x54, 0xb0, 0x0d, 0xea, 0x92, 0xb0, 0x84, 0x08, 0xc7, 0xec, 0x9a, 0x81, 0x1d, 0x96, 0x08,
	0x3e, 0x05, 0x56, 0x9c, 0x26, 0x4e, 0xad, 0x6b, 0x06, 0x4f, 0xc2, 0xfc, 0xe8, 0x7f, 0x02, 0xad,
	0x3d, 0x05, 0x99, 0x71, 0x26, 0x09, 0x7c, 0x03, 0xec, 0x4d, 0x37, 0xad, 0xf2, 0x78, 0xe0, 0xa2,
	0xc2, 0x0f, 0xaa, 0xfc, 0xa0, 0x2f, 0x15, 0x23, 0xdc, 0x92, 0xfd, 0x21, 0x80, 0x23, 0x49, 0x3f,
	0xa7, 0x94, 0xed, 0x5a, 0x72, 0xc0, 0x43, 0x99, 0x52, 0x46, 0x84, 0x74, 0xcc, 0xae, 0x15, 0xd8,
	0x61, 0x05, 0x6f, 0x9b, 0x7a, 0xfb, 0xe0, 0xdf, 0x59, 0xc7, 0xf0, 0x5b, 0xa0, 0x71, 0x43, 0xa7,
	0x30, 0xe6, 0x7f, 0x2d, 0xae, 0x15, 0x17, 0xe4, 0x5e, 0x23, 0xe7, 0x4e, 0x62, 0xce, 0x14, 0x61,
	0xca, 0xb1, 0xf4, 0x6d, 0x05, 0xfd, 0xb6, 0x8e, 0x73, 0x47, 0xba, 0x68, 0x39, 0x38, 0xaf, 0x01,
	0x6b, 0x24, 0x29, 0x8c, 0x01, 0xd8, 0x26, 0x05, 0x7b, 0xe8, 0xc0, 0x26, 0xd1, 0xa1, 0x7d, 0xb8,
	0x2f, 0x4e, 0xa1, 0x96, 0xc1, 0x7f, 0x07, 0x8f, 0xaa, 0x99, 0xe1, 0xf3, 0x63, 0xef, 0xf6, 0xd2,
	0x75, 0x83, 0xbb, 0x89, 0xa5, 0xfc, 0x0f, 0x60, 0x6f, 0x06, 0x84, 0xc7, 0x9f, 0xed, 0xc5, 0xeb,
	0xf6, 0x4e, 0x60, 0x16, 0x1d, 0x3e, 0x0c, 0x2f, 0x56, 0x9e, 0xb9, 0x5c, 0x79, 0xe6, 0xd5, 0xca,
	0x33, 0xff, 0xaf, 0x3d, 0x63, 0xb9, 0xf6, 0x8c, 0xcb, 0xb5, 0x67, 0x7c, 0x7b, 0x49, 0x53, 0x35,
	0x99, 0x8f, 0x51, 0xcc, 0x67, 0x58, 0xcb, 0xbd, 0x62, 0x44, 0xfd, 0xe5, 0xe2, 0x57, 0x89, 0xa6,
	0x24, 0xa1, 0x44, 0xe0, 0x85, 0xfe, 0x75, 0xc6, 0x75, 0xfd, 0x99, 0xbd, 0xbe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0xf3, 0xd6, 0xef, 0x4f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AnchorData(ctx context.Context, in *MsgAnchorDataRequest, opts ...grpc.CallOption) (*MsgAnchorDataResponse, error)
	SignData(ctx context.Context, in *MsgSignDataRequest, opts ...grpc.CallOption) (*MsgSignDataResponse, error)
	StoreData(ctx context.Context, in *MsgStoreDataRequest, opts ...grpc.CallOption) (*MsgStoreDataResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AnchorData(ctx context.Context, in *MsgAnchorDataRequest, opts ...grpc.CallOption) (*MsgAnchorDataResponse, error) {
	out := new(MsgAnchorDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/AnchorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SignData(ctx context.Context, in *MsgSignDataRequest, opts ...grpc.CallOption) (*MsgSignDataResponse, error) {
	out := new(MsgSignDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/SignData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StoreData(ctx context.Context, in *MsgStoreDataRequest, opts ...grpc.CallOption) (*MsgStoreDataResponse, error) {
	out := new(MsgStoreDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/StoreData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AnchorData(context.Context, *MsgAnchorDataRequest) (*MsgAnchorDataResponse, error)
	SignData(context.Context, *MsgSignDataRequest) (*MsgSignDataResponse, error)
	StoreData(context.Context, *MsgStoreDataRequest) (*MsgStoreDataResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AnchorData(ctx context.Context, req *MsgAnchorDataRequest) (*MsgAnchorDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnchorData not implemented")
}
func (*UnimplementedMsgServer) SignData(ctx context.Context, req *MsgSignDataRequest) (*MsgSignDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignData not implemented")
}
func (*UnimplementedMsgServer) StoreData(ctx context.Context, req *MsgStoreDataRequest) (*MsgStoreDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreData not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AnchorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAnchorDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AnchorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/AnchorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AnchorData(ctx, req.(*MsgAnchorDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SignData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/SignData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignData(ctx, req.(*MsgSignDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StoreData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/StoreData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreData(ctx, req.(*MsgStoreDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnchorData",
			Handler:    _Msg_AnchorData_Handler,
		},
		{
			MethodName: "SignData",
			Handler:    _Msg_SignData_Handler,
		},
		{
			MethodName: "StoreData",
			Handler:    _Msg_StoreData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/data/v1alpha1/tx.proto",
}

func (m *MsgAnchorDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchorDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchorDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAnchorDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAnchorDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAnchorDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSignDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgSignDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSignDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSignDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgStoreDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStoreDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStoreDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStoreDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAnchorDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAnchorDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSignDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSignDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgStoreDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgStoreDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAnchorDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnchorDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchorDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = append(m.Cid[:0], dAtA[iNdEx:postIndex]...)
			if m.Cid == nil {
				m.Cid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgAnchorDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAnchorDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAnchorDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSignDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSignDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = append(m.Cid[:0], dAtA[iNdEx:postIndex]...)
			if m.Cid == nil {
				m.Cid = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSignDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSignDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSignDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = append(m.Cid[:0], dAtA[iNdEx:postIndex]...)
			if m.Cid == nil {
				m.Cid = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgStoreDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgStoreDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStoreDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/ // Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api.proto

package v1alpha1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type Void struct {
}

func (m *Void) Reset()      { *m = Void{} }
func (*Void) ProtoMessage() {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Void.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return m.Size()
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type RegisterRequest struct {
	// Name of the adaptor in the form `adaptor-vendor.com/adaptor-vendor`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the API the adaptor was built against
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Name of the unix socket the adaptor is listening on, it's in the form `*.socket`
	Endpoint string `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (m *RegisterRequest) Reset()      { *m = RegisterRequest{} }
func (*RegisterRequest) ProtoMessage() {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegisterRequest) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

type ConnectionRequest struct {
	// Parameters for the connection, it's in form JSON bytes
	Parameters []byte `protobuf:"bytes,1,opt,name=parameters,proto3" json:"parameters,omitempty"`
	// Desired device, it's in form JSON bytes
	Device []byte `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *ConnectionRequest) Reset()      { *m = ConnectionRequest{} }
func (*ConnectionRequest) ProtoMessage() {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}
func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(m, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetParameters() []byte {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ConnectionRequest) GetDevice() []byte {
	if m != nil {
		return m.Device
	}
	return nil
}

type ConnectionResponse struct {
	// Observed device, it's in form JSON bytes
	Device []byte `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *ConnectionResponse) Reset()      { *m = ConnectionResponse{} }
func (*ConnectionResponse) ProtoMessage() {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}
func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetDevice() []byte {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*Void)(nil), "v1alpha1.Void")
	proto.RegisterType((*RegisterRequest)(nil), "v1alpha1.RegisterRequest")
	proto.RegisterType((*ConnectionRequest)(nil), "v1alpha1.ConnectionRequest")
	proto.RegisterType((*ConnectionResponse)(nil), "v1alpha1.ConnectionResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x3d, 0x6f, 0x1a, 0x41,
	0x14, 0xbc, 0x4d, 0x10, 0x1f, 0x4f, 0x88, 0x28, 0x1b, 0x29, 0xba, 0x5c, 0xa2, 0x55, 0x74, 0x4a,
	0x41, 0x91, 0x1c, 0x81, 0x14, 0xa9, 0xb1, 0x3b, 0x23, 0xb9, 0x38, 0x4b, 0x6e, 0x5c, 0x2d, 0xdc,
	0xf3, 0xb1, 0x12, 0xec, 0xae, 0xf7, 0x96, 0xab, 0xfd, 0x07, 0x2c, 0xf9, 0x67, 0x51, 0x52, 0x52,
	0x9a, 0xe3, 0x8f, 0x58, 0x2c, 0x9c, 0x39, 0x81, 0xbb, 0x9d, 0x37, 0xf3, 0xe6, 0x69, 0x66, 0xa1,
	0xc5, 0xb5, 0x88, 0xb4, 0x51, 0x56, 0xd1, 0x66, 0xde, 0xe7, 0x33, 0x3d, 0xe5, 0xfd, 0xe0, 0x4f,
	0x2a, 0xec, 0x74, 0x31, 0x8e, 0x26, 0x6a, 0xde, 0x4b, 0x55, 0xaa, 0x7a, 0x4e, 0x30, 0x5e, 0xdc,
	0x3b, 0xe4, 0x80, 0x7b, 0xed, 0x17, 0xc3, 0x3a, 0xd4, 0x6e, 0x95, 0x48, 0xc2, 0x3b, 0xf8, 0x14,
	0x63, 0x2a, 0x32, 0x8b, 0x26, 0xc6, 0x87, 0x05, 0x66, 0x96, 0x52, 0xa8, 0x49, 0x3e, 0x47, 0x9f,
	0xfc, 0x24, 0xdd, 0x56, 0xec, 0xde, 0xd4, 0x87, 0x46, 0x8e, 0x26, 0x13, 0x4a, 0xfa, 0x1f, 0xdc,
	0xb8, 0x84, 0x34, 0x80, 0x26, 0xca, 0x44, 0x2b, 0x21, 0xad, 0xff, 0xd1, 0x51, 0x6f, 0x38, 0x1c,
	0xc1, 0xe7, 0x4b, 0x25, 0x25, 0x4e, 0xac, 0x50, 0xb2, 0xb4, 0x67, 0x00, 0x9a, 0x1b, 0x3e, 0x47,
	0x8b, 0x26, 0x73, 0x47, 0xda, 0x71, 0x65, 0x42, 0xbf, 0x42, 0x3d, 0xc1, 0x5c, 0x4c, 0xd0, 0x5d,
	0x6a, 0xc7, 0x07, 0x14, 0xfe, 0x06, 0x5a, 0x35, 0xcb, 0xb4, 0x92, 0x19, 0x56, 0xd4, 0xa4, 0xaa,
	0x1e, 0x5c, 0xc3, 0x97, 0x7d, 0x2e, 0xc3, 0x77, 0xfa, 0x1b, 0x34, 0xbb, 0x31, 0xfd, 0x0f, 0xcd,
	0x32, 0x2e, 0xfd, 0x16, 0x95, 0xe5, 0x45, 0x27, 0x15, 0x04, 0x9d, 0x23, 0xe5, 0x5a, 0xf2, 0x06,
	0x4f, 0x04, 0x3a, 0xc3, 0x84, 0x6b, 0xab, 0x4c, 0xe9, 0xd5, 0x83, 0xd6, 0x08, 0x51, 0x0f, 0x67,
	0x22, 0x47, 0x7a, 0xb2, 0x71, 0xee, 0xd0, 0x25, 0xf4, 0x0a, 0x1a, 0x87, 0x04, 0xf4, 0xfb, 0x91,
	0x3e, 0x6b, 0x28, 0xf8, 0xf1, 0x3e, 0xb9, 0x4f, 0xbc, 0x73, 0xfa, 0x4b, 0x2e, 0x7e, 0x2d, 0x37,
	0x8c, 0xac, 0x37, 0xcc, 0x7b, 0x2c, 0x18, 0x59, 0x16, 0x8c, 0xac, 0x0a, 0x46, 0x5e, 0x0a, 0x46,
	0x9e, 0xb7, 0xcc, 0x5b, 0x6d, 0x99, 0xb7, 0xde, 0x32, 0x6f, 0x5c, 0x77, 0x9f, 0xfd, 0xef, 0x35,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0x33, 0x1b, 0x04, 0x32, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationServiceClient is the client API for RegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Void, error)
}

type registrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationServiceClient(cc *grpc.ClientConn) RegistrationServiceClient {
	return &registrationServiceClient{cc}
}

func (c *registrationServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/v1alpha1.RegistrationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServiceServer is the server API for RegistrationService service.
type RegistrationServiceServer interface {
	Register(context.Context, *RegisterRequest) (*Void, error)
}

// UnimplementedRegistrationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRegistrationServiceServer struct {
}

func (*UnimplementedRegistrationServiceServer) Register(ctx context.Context, req *RegisterRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegistrationServiceServer(s *grpc.Server, srv RegistrationServiceServer) {
	s.RegisterService(&_RegistrationService_serviceDesc, srv)
}

func _RegistrationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.RegistrationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.RegistrationService",
	HandlerType: (*RegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RegistrationService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// AdaptorServiceClient is the client API for AdaptorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdaptorServiceClient interface {
	KeepAlive(ctx context.Context, opts ...grpc.CallOption) (AdaptorService_KeepAliveClient, error)
	Connect(ctx context.Context, opts ...grpc.CallOption) (AdaptorService_ConnectClient, error)
}

type adaptorServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdaptorServiceClient(cc *grpc.ClientConn) AdaptorServiceClient {
	return &adaptorServiceClient{cc}
}

func (c *adaptorServiceClient) KeepAlive(ctx context.Context, opts ...grpc.CallOption) (AdaptorService_KeepAliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdaptorService_serviceDesc.Streams[0], "/v1alpha1.AdaptorService/KeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &adaptorServiceKeepAliveClient{stream}
	return x, nil
}

type AdaptorService_KeepAliveClient interface {
	Send(*Void) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type adaptorServiceKeepAliveClient struct {
	grpc.ClientStream
}

func (x *adaptorServiceKeepAliveClient) Send(m *Void) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adaptorServiceKeepAliveClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adaptorServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (AdaptorService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdaptorService_serviceDesc.Streams[1], "/v1alpha1.AdaptorService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &adaptorServiceConnectClient{stream}
	return x, nil
}

type AdaptorService_ConnectClient interface {
	Send(*ConnectionRequest) error
	Recv() (*ConnectionResponse, error)
	grpc.ClientStream
}

type adaptorServiceConnectClient struct {
	grpc.ClientStream
}

func (x *adaptorServiceConnectClient) Send(m *ConnectionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adaptorServiceConnectClient) Recv() (*ConnectionResponse, error) {
	m := new(ConnectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdaptorServiceServer is the server API for AdaptorService service.
type AdaptorServiceServer interface {
	KeepAlive(AdaptorService_KeepAliveServer) error
	Connect(AdaptorService_ConnectServer) error
}

// UnimplementedAdaptorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdaptorServiceServer struct {
}

func (*UnimplementedAdaptorServiceServer) KeepAlive(srv AdaptorService_KeepAliveServer) error {
	return status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedAdaptorServiceServer) Connect(srv AdaptorService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterAdaptorServiceServer(s *grpc.Server, srv AdaptorServiceServer) {
	s.RegisterService(&_AdaptorService_serviceDesc, srv)
}

func _AdaptorService_KeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdaptorServiceServer).KeepAlive(&adaptorServiceKeepAliveServer{stream})
}

type AdaptorService_KeepAliveServer interface {
	SendAndClose(*Void) error
	Recv() (*Void, error)
	grpc.ServerStream
}

type adaptorServiceKeepAliveServer struct {
	grpc.ServerStream
}

func (x *adaptorServiceKeepAliveServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adaptorServiceKeepAliveServer) Recv() (*Void, error) {
	m := new(Void)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AdaptorService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdaptorServiceServer).Connect(&adaptorServiceConnectServer{stream})
}

type AdaptorService_ConnectServer interface {
	Send(*ConnectionResponse) error
	Recv() (*ConnectionRequest, error)
	grpc.ServerStream
}

type adaptorServiceConnectServer struct {
	grpc.ServerStream
}

func (x *adaptorServiceConnectServer) Send(m *ConnectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adaptorServiceConnectServer) Recv() (*ConnectionRequest, error) {
	m := new(ConnectionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AdaptorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.AdaptorService",
	HandlerType: (*AdaptorServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "KeepAlive",
			Handler:       _AdaptorService_KeepAlive_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Connect",
			Handler:       _AdaptorService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}

func (m *Void) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Void) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Void) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RegisterRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConnectionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		i -= len(m.Device)
		copy(dAtA[i:], m.Device)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Device)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Parameters) > 0 {
		i -= len(m.Parameters)
		copy(dAtA[i:], m.Parameters)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Parameters)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConnectionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Device) > 0 {
		i -= len(m.Device)
		copy(dAtA[i:], m.Device)
		i = encodeVarintApi(dAtA, i, uint64(len(m.Device)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	offset -= sovApi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Void) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RegisterRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *ConnectionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Parameters)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *ConnectionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func sovApi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Void) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Void{`,
		`}`,
	}, "")
	return s
}
func (this *RegisterRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Endpoint:` + fmt.Sprintf("%v", this.Endpoint) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionRequest{`,
		`Parameters:` + fmt.Sprintf("%v", this.Parameters) + `,`,
		`Device:` + fmt.Sprintf("%v", this.Device) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConnectionResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConnectionResponse{`,
		`Device:` + fmt.Sprintf("%v", this.Device) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApi(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Void) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: Void: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Void: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *RegisterRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: RegisterRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *ConnectionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ConnectionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameters", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameters = append(m.Parameters[:0], dAtA[iNdEx:postIndex]...)
			if m.Parameters == nil {
				m.Parameters = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = append(m.Device[:0], dAtA[iNdEx:postIndex]...)
			if m.Device == nil {
				m.Device = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func (m *ConnectionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
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
			return fmt.Errorf("proto: ConnectionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
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
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = append(m.Device[:0], dAtA[iNdEx:postIndex]...)
			if m.Device == nil {
				m.Device = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
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
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
					return 0, ErrIntOverflowApi
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
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApi = fmt.Errorf("proto: unexpected end of group")
)
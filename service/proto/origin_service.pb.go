// Code generated by protoc-gen-go. DO NOT EDIT.
// source: origin_service.proto

package proto

import (
	"context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
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

type StreamRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a78a6e360964aa4, []int{0}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamRequest) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type StreamResponse struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a78a6e360964aa4, []int{1}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

type EnterRoomRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnterId              string   `protobuf:"bytes,2,opt,name=enter_id,json=enterId,proto3" json:"enter_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterRoomRequest) Reset()         { *m = EnterRoomRequest{} }
func (m *EnterRoomRequest) String() string { return proto.CompactTextString(m) }
func (*EnterRoomRequest) ProtoMessage()    {}
func (*EnterRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a78a6e360964aa4, []int{2}
}

func (m *EnterRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterRoomRequest.Unmarshal(m, b)
}
func (m *EnterRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterRoomRequest.Marshal(b, m, deterministic)
}
func (m *EnterRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterRoomRequest.Merge(m, src)
}
func (m *EnterRoomRequest) XXX_Size() int {
	return xxx_messageInfo_EnterRoomRequest.Size(m)
}
func (m *EnterRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnterRoomRequest proto.InternalMessageInfo

func (m *EnterRoomRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EnterRoomRequest) GetEnterId() string {
	if m != nil {
		return m.EnterId
	}
	return ""
}

type EnterRoomResponse struct {
	Info                 string   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnterRoomResponse) Reset()         { *m = EnterRoomResponse{} }
func (m *EnterRoomResponse) String() string { return proto.CompactTextString(m) }
func (*EnterRoomResponse) ProtoMessage()    {}
func (*EnterRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a78a6e360964aa4, []int{3}
}

func (m *EnterRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnterRoomResponse.Unmarshal(m, b)
}
func (m *EnterRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnterRoomResponse.Marshal(b, m, deterministic)
}
func (m *EnterRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnterRoomResponse.Merge(m, src)
}
func (m *EnterRoomResponse) XXX_Size() int {
	return xxx_messageInfo_EnterRoomResponse.Size(m)
}
func (m *EnterRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnterRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnterRoomResponse proto.InternalMessageInfo

func (m *EnterRoomResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "proto.streamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.streamResponse")
	proto.RegisterType((*EnterRoomRequest)(nil), "proto.enterRoomRequest")
	proto.RegisterType((*EnterRoomResponse)(nil), "proto.enterRoomResponse")
}

func init() { proto.RegisterFile("origin_service.proto", fileDescriptor_4a78a6e360964aa4) }

var fileDescriptor_4a78a6e360964aa4 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x2f, 0xca, 0x4c,
	0xcf, 0xcc, 0x8b, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x05, 0x53, 0x4a, 0xc6, 0x5c, 0xbc, 0xc5, 0x25, 0x45, 0xa9, 0x89, 0xb9, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x99, 0x79, 0x69, 0xf9, 0x12, 0x4c, 0x60,
	0x11, 0x30, 0x5b, 0x49, 0x85, 0x8b, 0x0f, 0xa6, 0xa9, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xae,
	0x8a, 0x11, 0x49, 0x95, 0x2d, 0x97, 0x40, 0x6a, 0x5e, 0x49, 0x6a, 0x51, 0x50, 0x7e, 0x3e, 0x4e,
	0xd3, 0x25, 0xb9, 0x38, 0xc0, 0x6a, 0xe2, 0x33, 0x53, 0xa0, 0x36, 0xb0, 0x83, 0xf9, 0x9e, 0x29,
	0x4a, 0xea, 0x5c, 0x82, 0x48, 0xda, 0x71, 0xdb, 0x63, 0x34, 0x85, 0x91, 0x8b, 0xd3, 0xb7, 0x32,
	0x18, 0xe2, 0x3b, 0x21, 0x47, 0x2e, 0x9e, 0xa0, 0xd4, 0xf4, 0xcc, 0x62, 0xa8, 0x4e, 0x21, 0x11,
	0x88, 0x7f, 0xf5, 0x50, 0x7c, 0x29, 0x25, 0x8a, 0x26, 0x0a, 0x31, 0x5e, 0x89, 0x41, 0x83, 0xd1,
	0x80, 0x51, 0xc8, 0x85, 0x8b, 0xd3, 0x15, 0x66, 0xb3, 0x90, 0x38, 0x54, 0x25, 0xba, 0x57, 0xa4,
	0x24, 0x30, 0x25, 0x90, 0x4d, 0x49, 0x62, 0x03, 0x4b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xff, 0x33, 0x70, 0x7f, 0x01, 0x00, 0x00,
}

type StreamServiceClient interface {
	RegisterRoom(ctx context.Context, opts ...grpc.CallOption) (StreamClient, error)
	EnterRoom(ctx context.Context, opts ...grpc.CallOption) (StreamClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) RegisterRoom(ctx context.Context, opts ...grpc.CallOption) (StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.MyService/RegisterRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClient{stream}
	return x, nil
}

func (c *streamServiceClient) EnterRoom(ctx context.Context, opts ...grpc.CallOption) (StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.MyService/EnterRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamClient{stream}
	return x, nil
}

type StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	Send2(*EnterRoomRequest) error
	Recv2() (*EnterRoomResponse, error)
	grpc.ClientStream
}

type streamClient struct {
	grpc.ClientStream
}

func (x *streamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}
func (x *streamClient) Send2(m *EnterRoomRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *streamClient) Recv2() (*EnterRoomResponse, error) {
	m := new(EnterRoomResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MyService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RegisterRoom",
			Handler:       _MyService_RegisterRoom_Handler,
			ClientStreams: true,
			ServerStreams: true,
		},
		{
			StreamName:    "EnterRoom",
			Handler:       _MyService_EnterRoom_Handler,
			ClientStreams: true,
			ServerStreams: true,
		},
	},
	Metadata: "origin_service.proto",
}

func _MyService_RegisterRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).RegisterRoom(&streamServiceRecordServer{stream})
}

func _MyService_EnterRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).EnterRoom(&streamServiceRecordServer{stream})
}

type StreamServiceServer interface {
	RegisterRoom(StreamService_RecordServer) error
	EnterRoom(StreamService_RecordServer) error
}

type StreamService_RecordServer interface {
	Recv() (*StreamRequest, error)
	Send(*StreamResponse) error
	Recv2() (*EnterRoomRequest, error)
	Send2(*EnterRoomResponse) error
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Send2(m *EnterRoomResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}
func (x *streamServiceRecordServer) Recv2() (*EnterRoomRequest, error) {
	m := new(EnterRoomRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

package proto

import (
	"context"

	"google.golang.org/grpc"
)

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

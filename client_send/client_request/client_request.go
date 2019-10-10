package client_request

import (
	pb "stream_learn/room"

	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type GreeterClient struct {
	// Sends a greeting
	SayHello(ctx context.Context, in *pb.RoomReq, opts ...grpc.CallOption) (*pb.RoomRsp, error)
}

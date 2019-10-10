package main

import (
	"context"
	"log"
	"net"

	pb "stream_learn/room"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type streamServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *streamServer) SayHello(ctx context.Context, in *pb.RoomReq) (*pb.RoomRsp, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.RoomRsp{Name: "Server name "}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStreamServer(s, &streamServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"io"
	"log"
	"net"
	// "context"

	"google.golang.org/grpc"

	pb "stream_learn/service/proto"
)

type StreamServer struct{}

const (
	PORT = ":50051"
)

func (s *StreamServer) Service1(stream pb.StreamService_RecordServer) error {
    for {
        r, err := stream.Recv()
        if err == io.EOF {
            return stream.SendAndClose(&pb.StreamResponse{Name:"close stream"})
        }
        if err != nil {
            return err
        }

        log.Printf("stream.Recv pt.name: %s", r.Name)
    }

    return nil
}
func main() {
	
	server := grpc.NewServer()
    pb.RegisterStreamServiceServer(server, &StreamServer{})

    lis, err := net.Listen("tcp", PORT)
    if err != nil {
        log.Fatalf("net.Listen err: %v", err)
    }

    server.Serve(lis)
}
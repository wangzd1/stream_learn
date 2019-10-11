package main

import (
	"context"
	"log"
	// "os"
	// "time"

	"google.golang.org/grpc"
	pb "stream_learn/service/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world22"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStreamServiceClient(conn)

	stream, err := c.Service1(context.Background())
    if err != nil {
		log.Fatalf("err1",err)
    }

    for n := 0; n < 6; n++ {
        err := stream.Send(&pb.StreamRequest{Name:"qa"})
        if err != nil {
            log.Fatalf("err2",err)
        }
    }

    resp, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("err3",err)
    }

    log.Printf("resp: pj.name: %s", resp.Name)
}
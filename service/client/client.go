package main

import (
	"context"
	"log"
	"strconv"

	"fmt"
	// "time"

	pb "stream_learn/service/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world22"
)

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go clients(i, ch)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	defer close(ch)
}

func clients(index int, ch chan int) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStreamServiceClient(conn)

	stream, err := c.Service1(context.Background())
	if err != nil {
		log.Fatalf("err1", err)
	}

	for n := 0; n < 3; n++ {
		err := stream.Send(&pb.StreamRequest{Name: "request" + strconv.Itoa(index)})
		if err != nil {
			log.Fatalf("err2", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err3", err)
	}
	ch <- 1
	log.Printf("resp: pj.name: %s", resp.Name)
}

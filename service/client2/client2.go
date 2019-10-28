package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	// "time"

	pb "stream_learn/service/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world22"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStreamServiceClient(conn)
	stream, err := c.Service1(context.Background())
	if err != nil {
		fmt.Println("err1")
	}
	var n = 0
	for {
		err = stream.Send(&pb.StreamRequest{
			Id:         "a0101",
			Info:       "request" + strconv.Itoa(n),
			ClientType: 2,
		})
		if err != nil {
			fmt.Println("err2")
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("err3")
		}

		log.Printf("resp: pj.info: %s", resp.Info)
	}

	stream.CloseSend()

}

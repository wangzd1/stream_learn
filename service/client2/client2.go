package main

import (
	"context"
	"fmt"
	"io"
	"log"

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
	stream, err := c.EnterRoom(context.Background())
	if err != nil {
		fmt.Println("err1")
	}
	err = stream.Send2(&pb.EnterRoomRequest{
		Id:      "a0101",
		EnterId: "0101",
	})
	if err != nil {
		fmt.Println("err2")
	}
	for {
		resp, err := stream.Recv2()
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

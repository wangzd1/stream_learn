package main

import (
	"bufio"
	"context"
	"log"
	"os"
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
	for i := 0; i < 1; i++ {
		go clients(i, ch)
	}
	for i := 0; i < 1; i++ {
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
	i := 0
	for {
		input := bufio.NewReader(os.Stdin)
		inputstring, _, _ := input.ReadLine()
		err := stream.Send(&pb.StreamRequest{Name: "request" + strconv.Itoa(i) + string(inputstring)})
		// time.Sleep(time.Second * 3)
		if err != nil {
			log.Fatalf("err2", err)
		}
		i++
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err3", err)
	}
	ch <- 1
	log.Printf("resp: pj.name: %s", resp.Name)
}

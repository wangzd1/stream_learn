package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
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
	stream, err := c.RegisterRoom(context.Background())
	if err != nil {
		fmt.Println("err1")
	}
	var n = 0
	for {
		n++
		fmt.Println("send1a", n)
		//读键盘
		reader := bufio.NewReader(os.Stdin)
		//以换行符结束
		str, _ := reader.ReadString('\n')
		err = stream.Send(&pb.StreamRequest{
			Id:   "0101",
			Info: "request" + strconv.Itoa(n) + "say:" + str,
		})
		fmt.Println("send1d", n)
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

// func main() {
// 	ch := make(chan int, 10)
// 	for i := 0; i < 1; i++ {
// 		go clients(i, ch)
// 	}
// 	for i := 0; i < 1; i++ {
// 		fmt.Println(<-ch)
// 	}
// 	defer close(ch)
// }

// func clients(index int, ch chan int) {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewStreamServiceClient(conn)

// 	stream, err := c.RegisterRoom(context.Background())
// 	if err != nil {
// 		log.Fatalf("err1", err)
// 	}
// 	i := 0
// 	for {
// 		input := bufio.NewReader(os.Stdin)
// 		inputstring, _, _ := input.ReadLine()
// 		err := stream.Send(&pb.StreamRequest{
// 			Id:         "0101",
// 			Info:       "request" + strconv.Itoa(i) + string(inputstring),
// 			ClientType: 1,
// 		})
// 		// time.Sleep(time.Second * 3)
// 		if err != nil {
// 			log.Fatalf("err2", err)
// 		}
// 		i++
// 	}

// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Fatalf("err3", err)
// 	}
// 	ch <- 1
// 	log.Printf("resp: pj.Info: %s", resp.Info)
// }

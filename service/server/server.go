package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	// "context"

	"google.golang.org/grpc"

	pb "stream_learn/service/proto"

	"github.com/pkg/profile"
)

type StreamServer struct{}

const (
	PORT = ":50051"
)

type Rm struct {
	Id        string
	Origin    chan string
	Audiences []Audience
}

type Audience struct {
	SecondId string
	Origin   chan string
}

var a Rm

func (s *StreamServer) RegisterRoom(stream pb.StreamService_RecordServer) error {
	n := 0
	fmt.Println(a)
	for {
		time.Sleep(time.Second)
		fmt.Println("a.Recvqqq")
		r, err := stream.Recv()
		fmt.Println("a.Recvwww")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		a.Id = r.Id
		if len(a.Audiences) > 0 {
			for _, v := range a.Audiences {
				v.Origin <- r.Info
			}
			err = stream.Send(&pb.StreamResponse{
				Info: "some receive,server responce " + r.Info,
			})
			if err != nil {
				return err
			}
		} else {
			err = stream.Send(&pb.StreamResponse{
				Info: "no one receive,server responce " + r.Info,
			})
			if err != nil {
				return err
			}
		}
		log.Printf("stream.send1 pt.info: %s", r.Info)
	}
	fmt.Println("regist exit")
	return nil
}
func (s *StreamServer) EnterRoom(stream pb.StreamService_RecordServer) error {
	n := 0
	fmt.Println(a)
	curchan := make(chan string, 10)
	time.Sleep(time.Second)
	r, err := stream.Recv2()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if a.Id == r.EnterId {
		a.Audiences = append(a.Audiences, Audience{
			SecondId: r.Id,
			Origin:   curchan,
		})
	}
	fmt.Println("curchan", &curchan)
	for {
		n++
		err = stream.Send2(&pb.EnterRoomResponse{
			Info: "server responce " + <-curchan,
		})
		if err != nil {
			return err
		}
		log.Printf("stream.send2 pt.Id: %s", r.Id)
	}
	fmt.Println("enter exit")
	return nil
}

func main() {
	defer profile.Start().Stop()
	a.Origin = make(chan string, 10)
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamServer{})

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

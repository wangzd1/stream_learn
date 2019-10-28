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

func CheckJoin(secondId string) (int, error) {
	for k, v := range a.Audiences {
		if v.SecondId == secondId {
			return k, nil
		}
	}
	return -1, nil
}

func (audience *Audience) Join() {

}

func (audience *Audience) Leave() {

}

func (s *StreamServer) Service1(stream pb.StreamService_RecordServer) error {
	n := 0
	fmt.Println(a)
	for {

		time.Sleep(time.Second)

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		if r.ClientType == 1 {
			a.Id = r.Id
			a.Origin <- r.Info
			err = stream.Send(&pb.StreamResponse{
				Info: "server responce " + r.Info,
			})
			if err != nil {
				return err
			}
			log.Printf("stream.send1 pt.info: %s", r.Info)
		} else {
			if r.Id == a.Id {
				// index,err:=CheckJoin(r.Id)
				err = stream.Send(&pb.StreamResponse{
					Info: <-a.Origin,
				})
				if err != nil {
					return err
				}
				log.Printf("stream.send2 info: %s", r.Info)
			}
		}

	}

	return nil
}

func main() {
	a.Origin = make(chan string, 10)
	server := grpc.NewServer()
	pb.RegisterStreamServiceServer(server, &StreamServer{})

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server.Serve(lis)
}

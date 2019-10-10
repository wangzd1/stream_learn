package main

import (
	"fmt"
	"log"

	pb "stream_learn/room"

	"github.com/golang/protobuf/proto"
)

func main() {
	test := &pb.Room{
		Id:   17,
		Name: "hello",
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println("Marshaldata", data)
	newTest := &pb.Room{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("Unmarshaldata", newTest)
	// Now test and newTest contain the same data.
	if test.GetId() != newTest.GetId() {
		log.Fatalf("data mismatch %q != %q", test.GetId(), newTest.GetId())
	}
	// etc.
}

package main

import (
	"fmt"
	"stream_learn/tarsgo/TestApp"

	"github.com/TarsCloud/TarsGo/tars"
)

//tars.Communicator should only init once and be global
var comm *tars.Communicator

func main() {
	comm = tars.NewCommunicator()
	obj := "TestApp.HelloServer.HelloObj@tcp -h 127.0.0.1 -p 20011 -t 60000"
	app := new(TestApp.Hello)
	comm.StringToProxy(obj, app)
	var req string = "Hello Wold"
	var out string = "Hello Wold"
	ret, err := app.TestHello(req, &out)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(ret, out)
}

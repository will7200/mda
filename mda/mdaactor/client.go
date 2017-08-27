package main

import (
	"fmt"
	"time"

	console "github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/will7200/mda/mda/grpc/pb"
)

func main() {
	timeout := 5 * time.Second
	remote.Start("127.0.0.1:8081")
	pid, _ := remote.SpawnNamed("127.0.0.1:8000", "remote", "start", timeout)
	res, _ := pid.RequestFuture(&pb.StartRequest{Id: "123"}, timeout).Result()
	response := res.(*pb.StartReply)
	fmt.Printf("Response from remote %v", response.Message)

	console.ReadLine()

}

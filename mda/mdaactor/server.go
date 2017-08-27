package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/will7200/mda/mda/grpc/pb"
	"github.com/will7200/mjs/job"
)

type starterActor struct{}

func (*starterActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *pb.StartRequest:
		fmt.Println("Starting...")
		ctx.Respond(&pb.StartReply{
			Message: "Yeah",
			Error:   0,
		})
	case *pb.GetRequest:
		fmt.Println("Getting...")
	case *pb.DisableRequest:
		fmt.Println("Disabling")
	default:
		a := reflect.TypeOf(ctx.Message())
		fmt.Printf("Type %s is not implemented", a)
	}
}

func newStarterActor() actor.Actor {
	return &starterActor{}
}

func init() {
	remote.Register("start", actor.FromProducer(newStarterActor))
}
func main() {
	port := 8000
	fmt.Println(strconv.Itoa(port))
	c := make(chan struct{})
	f := func() {
		c <- struct{}{}
	}
	job.NewAfterFunc(time.Second*15, f)
	remote.Start("0.0.0.0:" + strconv.Itoa(port))
	<-c
}

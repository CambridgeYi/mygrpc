package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"mygrpc/proto/helloworld"
	"os"
)

const (
	address = "localhost:50051"
	defaultName = "world"
)

func main(){
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("did not connect:%v",err)

	}
	defer conn.Close()
	c:= helloworld.NewGreeterClient(conn)

	name:= defaultName
	if len(os.Args)>1{
		name =os.Args[1]
	}
	r,err := c.SayHello(context.Background(),&helloworld.HelloRequest{Name: name})

	if err !=nil {
		log.Fatal("could not greet :%v",err)

	}
	log.Printf("Greeting :%s",r.Message)
}

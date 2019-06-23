package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"mygrpc/proto/helloworld"
	"net"
)

const (
	port = ":50051"
)
type server struct{}

func (s *server) SayHello(ctx context.Context,in *helloworld.HelloRequest)(*helloworld.HelloReply,error){
	return &helloworld.HelloReply{Message: "hello" +in.Name},nil
}

func main(){
	lis,err := net.Listen("tcp",port)
	if err !=nil{
		log.Fatalf("failed to listen: %v",err)

	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s,&server{})
	s.Serve(lis)
}
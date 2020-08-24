package main

import (
	"context"
	"log"
	"net"

	pb "gorpc1/hello"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "你好" + in.Name}, nil
}

//服务端
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	s := grpc.NewServer()                        //grpc生成一个新的服务
	pb.RegisterHellosServiceServer(s, &server{}) //注册
	s.Serve(lis)
}

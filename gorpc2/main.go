package main

import (
	"context"
	"log"
	"os"

	pb "gorpc1/hello"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

//客户端
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) //建立grpc的连接
	if err != nil {
		log.Fatal("did not connect: ", err)
	}
	defer conn.Close()
	c := pb.NewHellosServiceClient(conn) //把连接传进去

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name}) //调服务端的函数
	if err != nil {
		log.Fatal("counld not greet: ", err)
	}

	log.Printf("Greeting: %s", r.Reply)
}

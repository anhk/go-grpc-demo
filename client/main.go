package main

import (
	"context"
	"fmt"
	"go-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := proto.NewGreeterClient(conn)

	stream, err := cli.SayHello(context.Background())
	if err != nil {
		panic(err)
	}

out:
	for {
		select {
		case <-stream.Context().Done():
			goto out
		default:
			msg, err := stream.Recv()
			if err != nil {
				break out
			}
			fmt.Println("recv:", msg)
		}
	}
	//r, err := cli.SayHello(ctx, &proto.HelloRequest{
	//	Type: proto.MessageType_Bar,
	//	Name: "hello world",
	//})

}

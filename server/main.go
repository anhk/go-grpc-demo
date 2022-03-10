package main

import (
	"fmt"
	"go-grpc-demo/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Servcie struct {
	proto.UnimplementedGreeterServer
}

var clients = make([]proto.Greeter_SayHelloServer, 0, 1)

func (s Servcie) SayHello(stream proto.Greeter_SayHelloServer) error {
	clients = append(clients, stream)
	for {
		select {
		case <-stream.Context().Done():
			return nil
		default:
			msg, err := stream.Recv()
			if err != nil || msg == nil {
				fmt.Println("err: ", err)
				continue
			}
			fmt.Println("receive:", msg)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8091")
	if err != nil {
		panic(err)
	}

	svc := grpc.NewServer()
	proto.RegisterGreeterServer(svc, &Servcie{})

	go unsafeSendMessage()

	if err := svc.Serve(listen); err != nil {
		panic(err)
	}
}

func unsafeSendMessage() {
out:
	for {
		for _, c := range clients {
			fmt.Println("Send:")
			if err := c.Send(&proto.HelloMessage{
				Type: 0,
				Name: "hello World-0089",
			}); err != nil {
				fmt.Println("err: ", err)
				break out
			}
		}
		time.Sleep(time.Second)
	}
}

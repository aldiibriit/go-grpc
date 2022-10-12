package main

import (
	"context"
	"fmt"
	pb "go-grpc-2/proto/echo"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func callerUnaryEcho(c pb.EchoServiceClient, message string) {
	fmt.Println("call unary echo")

	md := metadata.Pairs("timestamps", time.Now().Format(time.StampNano))

	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		os.Exit(1)
	}

	fmt.Print("response")
	fmt.Printf("-> %s", r.Message)
}

func main() {
	conn, err := grpc.Dial(":7177", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("error : %v", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	callerUnaryEcho(client, "message -1")
	callerUnaryEcho(client, "aldi")
}

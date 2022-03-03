package main

import (
	"context"
	"fmt"
	pb "github.com/Zaric666/learning/grpc/proto/echo"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)

	res, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{
		Message: "Hi, I'm client A",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(res.Message)
}

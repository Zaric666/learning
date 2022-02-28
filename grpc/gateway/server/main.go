package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Zaric666/learning/grpc/gateway/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("port", 50052, "the port to serve on")

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello " + in.Name}, nil
}

func main() {
	// Create a gRPC server object
	s := grpc.NewServer()

	// Attach the Greeter service to the server
	pb.RegisterHelloServiceServer(s, &server{})

	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", *port))

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

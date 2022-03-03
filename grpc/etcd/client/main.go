package main

import (
	"context"
	"fmt"
	discovery "github.com/Zaric666/learning/grpc/etcd/dicovery"
	pb "github.com/Zaric666/learning/grpc/proto/echo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
)

const App = "base-hello"

func main() {
	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	}

	addrs := []string{"127.0.0.1:12379"}
	r := discovery.NewResolver(addrs, zap.NewNop())
	resolver.Register(r)

	conn, err := grpc.Dial("etcd:///"+App, options...)
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

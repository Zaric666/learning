package main

import (
	"context"
	"fmt"
	"github.com/Zaric666/learning/grpc/etcd/dicovery"
	pb "github.com/Zaric666/learning/grpc/proto/echo"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

const (
	app         = "base-hello"
	grpcAddress = "127.0.0.1:8083"
)

type ecServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *ecServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}

func main() {
	grpcAdders := []string{"127.0.0.1:8083", "127.0.0.1:8084", "127.0.0.1:8085", "127.0.0.1:8086"}
	adders := []string{"127.0.0.1:12379"}
	etcdRegister := discovery.NewRegister(adders, zap.NewNop())
	servers := make([]*grpc.Server, 10)
	for _, grpcAdder := range grpcAdders {
		node := discovery.Server{
			Name: app,
			Addr: grpcAdder,
		}
		server, err := StartServer(grpcAdder)
		servers = append(servers, server)
		if err != nil {
			panic(fmt.Sprintf("start server failed : %v", err))
		}
		if _, err := etcdRegister.Register(node, 10); err != nil {
			panic(fmt.Sprintf("server register failed: %v", err))
		}

		fmt.Println("service started listen on", grpcAdder)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			for _, s := range servers {
				s.Stop()
			}
			etcdRegister.Stop()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func StartServer(addr string) (*grpc.Server, error) {
	s := grpc.NewServer()

	pb.RegisterEchoServer(s, &ecServer{addr: addr})
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return s, nil
}

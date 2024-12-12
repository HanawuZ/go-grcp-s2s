package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/HanawuZ/go-grcp-s2s/shared/grpc/pingpong"
	"google.golang.org/grpc"
)

type PingPongServerImpl struct {
	pb.UnimplementedPingPongServer
}

func (s *PingPongServerImpl) StartPing(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {

	fmt.Println("Ping Received")

	resp := pb.Pong{
		Id:      ping.Id,
		Message: "Received " + ping.Message,
	}

	return &resp, nil
}

func StartServer() {
	server := &PingPongServerImpl{}
	listener, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPingPongServer(grpcServer, server)

	log.Printf("server listening at %v", listener.Addr())

	// Start grpcServer
	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
	fmt.Println("pingpong started")

}

func main() {
	StartServer()
}

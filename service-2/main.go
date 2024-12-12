package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/HanawuZ/go-grcp-s2s/shared/grpc/pingpong"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	connection, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := pb.NewPingPongClient(connection)

	response, err := client.StartPing(ctx, &pb.Ping{Id: 1, Message: "ping"})
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
}

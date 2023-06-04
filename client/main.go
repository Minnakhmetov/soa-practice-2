package main

import (
	"context"
	"fmt"

	pb "github.com/Minnakhmetov/soa-practice-2/mafia"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 65434

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewReverseClient(conn)
	ctx := context.Background()
	resp, err := client.Do(ctx, &pb.Request{Message: "moscow", X: 6, Y: 17})
	if err != nil {
		print(err)
	}
	println(resp.GetMessage())
	println(resp.GetZ())
}

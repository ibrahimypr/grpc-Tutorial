package main

import (
	"context"
	"log"
	pb "github.com/ibrahimypr/grpc-tutorial/clientStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.Dial("localhost:4848", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect :%v", err)
	}
	defer conn.Close()

	c := pb.NewRatingServiceClient(conn)

	stream, err := c.RateProduct(context.Background())
	if err != nil {
		log.Fatalf("Error while calling RateProduct: %v", err)
	}

	ratings := []*pb.ProductRating{
		{ProductId: "p1", Rating: 5},
		{ProductId: "p2", Rating: 4},
		{ProductId: "p3", Rating: 3},
		{ProductId: "p4", Rating: 5},
		{ProductId: "p5", Rating: 4},
	}
	for _, rating := range ratings {
		stream.Send(rating)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response: %v", err)
	}

	log.Printf("Average Rating: %v", res.GetAverage())
}
package main

import (
	"io"
	"log"
	"net"

	pb "github.com/ibrahimypr/grpc-tutorial/clientStreaming/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.RatingServiceServer
}

func (s *server) RateProduct(stream pb.RatingService_RateProductServer) error {
	var sum int32
	var count int32

	for {
		rating, err := stream.Recv()
		if err == io.EOF {
			average := float32(sum) / float32(count)
			return stream.SendAndClose(&pb.AverageRating{Average: average})
		}
		if err != nil {
			return err
		}
		sum += rating.Rating
		count++
	}
}

func main() {
	lis, err := net.Listen("tcp", ":4848")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRatingServiceServer(s, &server{})
	log.Println("Server is running on port 4848")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
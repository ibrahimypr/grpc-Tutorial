package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "github.com/ibrahimypr/grpc-tutorial/unary/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.PersonServiceServer
}

func (s *server) GetPerson(ctx context.Context, req *pb.PersonRequest) (*pb.Person, error) {
	if req.Name == "Ali" {
		return &pb.Person{
			Name: "Ali",
			Age: 30,
		}, nil
	}
	return nil, fmt.Errorf("Person not found")
}

func (s *server) GetPersons(ctx context.Context, req *pb.Empty) (*pb.PersonList, error) {
	return &pb.PersonList{
		Persons: []*pb.Person{
			{Name: "Ali", Age: 30},
			{Name: "ibo", Age: 24},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4848")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{})
	log.Printf("Listening on %v\n", ":4848")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err);
	}
}
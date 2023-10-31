package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ibrahimypr/grpc-tutorial/unary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4848", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect :%v", err)
	}
	defer conn.Close()

	c := pb.NewPersonServiceClient(conn)

	personName := "Ali"
	response, err := c.GetPerson(context.Background(), &pb.PersonRequest{Name: personName})
	if err != nil {
		log.Fatalf("error calling GetPerson RPC:%v", err)
	}
	fmt.Printf("Received response for %s: %v\n", personName, response)

	responseList, err := c.GetPersons(context.Background(), &pb.Empty{})

	if err != nil {
		log.Fatalf("Error calling GetPersons RPC: %v", err)
	}
	fmt.Printf("Received person list: %v\n", responseList)
}
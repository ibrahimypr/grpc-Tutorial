package main

import (
	"log"
	"net"
	"io"
	pb "github.com/ibrahimypr/grpc-tutorial/biDirectionalStreaming/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.ChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message from %s: %s", msg.GetUser(), msg.GetContent())
		err = stream.Send(&pb.Message{
			User:    "Server",
			Content: "Hello " + msg.GetUser(),
		})
		if err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":4848")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Println("Server is running on port 4848")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

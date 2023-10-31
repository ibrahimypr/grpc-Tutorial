package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	pb "github.com/ibrahimypr/grpc-tutorial/biDirectionalStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main() {
	conn, err := grpc.Dial("localhost:4848", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect :%v", err)
	}
	defer conn.Close()

	c := pb.NewChatServiceClient(conn)

	stream, err := c.Chat(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}
			log.Printf("%s: %s", msg.GetUser(), msg.GetContent())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		err := stream.Send(&pb.Message{
			User:    "Client",
			Content: text,
		})
		if err != nil {
			log.Fatalf("Error while sending message: %v", err)
		}
	}
}
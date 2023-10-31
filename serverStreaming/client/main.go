package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/ibrahimypr/grpc-tutorial/serverStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4848", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}
	defer conn.Close()

	client := pb.NewWeatherServiceClient(conn)

	request := &pb.CityRequest{CityName: "Istanbul"}
	stream, err := client.GetWeeklyForecast(context.Background(), request)
	if err != nil {
		log.Fatalf("Error on get weekly forecast : %v", err)
	}

	for {
		forecast, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error when reading stream: %v", err)
		}
		fmt.Printf("Weather on %s: %s, Temperature: %f celcius\n", forecast.Data, forecast.Description, forecast.Temperature)
	}
}
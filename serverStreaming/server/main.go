package main

import (
	"log"
	"net"
	"time"
	pb "github.com/ibrahimypr/grpc-tutorial/serverStreaming/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.WeatherServiceServer
}

func (s *server) GetWeeklyForecast(req *pb.CityRequest, stream pb.WeatherService_GetWeeklyForecastServer) error{
	for i := 0; i < 5; i++ {
		forecast := &pb.WeatherForecast{
			Data: time.Now().AddDate(0, 0, i).Format("2008-04-04"),
			Description: "Sunny",
			Temperature: 20 + float32(i),
		}
		if err := stream.Send(forecast); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":4848")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWeatherServiceServer(s, &server{})

	log.Println("Server is running on port 4848")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
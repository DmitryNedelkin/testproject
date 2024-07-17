package grpc_server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"testproject/database"
	pb "testproject/metrics"
	"time"

	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMetricsServer
}

func (s *server) Do(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetIndex())
	i, err := strconv.Atoi(in.GetIndex())
	if err != nil {
		panic(err)
	}

	response := database.GetMessageFromDatabase(i)
	if response != nil {
		jsonString, err := json.Marshal(&response)
		if err != nil {
			log.Fatal("Error marshaling to JSON: ", err)
		}

		return &pb.Response{Message: "Response: " + string(jsonString)}, nil
	} else {
		return &pb.Response{Message: "Error"}, nil
	}
}

func (s *server) DoStreamResponse(in *pb.Request, stream pb.Metrics_DoStreamResponseServer) error {
	log.Println("0")
	allMessages := database.GetAllMessages()
	log.Println("allMessages")
	for message := range allMessages {
		jsonString, err := json.Marshal(&message)
		if err != nil {
			log.Fatal("Error marshaling to JSON: ", err)
		}
		log.Println("1")
		if err := stream.Send(&pb.Response{Message: string(jsonString)}); err != nil {
			return err
		}
		log.Println("2")
		time.Sleep(1 * time.Second)
		log.Println("3")
	}

	return nil
}

func StartGrpsServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Starting grpc server on :50051")

	s := grpc.NewServer()
	pb.RegisterMetricsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

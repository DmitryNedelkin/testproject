package grpc_server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"testproject/database"
	pb "testproject/metrics"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMetricsServer
}

func (s *server) ListMetrics(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.GetIndex())
	response := database.GetMessageFromDatabase(in.GetIndex())
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

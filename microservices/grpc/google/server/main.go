package main

import (
	"context"
	"log"
	"net"

	"github.com/williamnoble/goExercises/microservices/grpc/google/helloworld"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Recieved %v", in.Name)
	return &helloworld.HelloReply{Message: "Hello" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	s := grpc.NewServer()

	helloworld.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}

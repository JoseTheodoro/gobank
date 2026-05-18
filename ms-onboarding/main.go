package main

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/onboarding"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedOnboardingServer
}

func (s *Server) StartOnboarding(ctx context.Context, in *pb.OnboardingRequest) (*pb.OnboardingResponse, error) {

	msg := fmt.Sprintf("%s, %s", in.GetM(), "Estou respondendo do servidor")

	return &pb.OnboardingResponse{
		A: msg,
	}, nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on load .env file")
	}

	server, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterOnboardingServer(s, &Server{})
	fmt.Printf("gRPC server started at localhost%s", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(server); err != nil {
		log.Fatalf("error on grpc server: %v\n", err)
	}

}

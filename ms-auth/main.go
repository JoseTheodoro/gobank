package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "gobank/contracts/pb/auth"
)

type Server struct {
	pb.UnimplementedAuthServer
}

func (s *Server) CreateCredentials(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	fmt.Printf("request received: %s", in.String())
	return &pb.AuthResponse{
		B: fmt.Sprintf("%s, hello from ms-auth", in.GetName()),
	}, nil
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error on load .env file")
	}

	conn, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error on connect tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &Server{})
	fmt.Printf("ms-auth: gRPC server started at localhost%v\n", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(conn); err != nil {
		log.Fatalf("error on grpc server: %v", err)
	}

}

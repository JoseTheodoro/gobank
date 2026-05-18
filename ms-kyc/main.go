package main

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/kyc"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedKYCServer
}

func (s *Server) StartKYC(ctx context.Context, in *pb.StartKYCRequest) (*pb.StartKYCResponse, error) {
	return &pb.StartKYCResponse{
		B: fmt.Sprintf("%s, hello from ms-kyc", in.GetName()),
	}, nil
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error on load .env file")
	}

	conn, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error on connect to tcp: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterKYCServer(s, &Server{})
	fmt.Printf("ms-kyc: gRPC server started at localhost%s\n", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(conn); err != nil {
		log.Fatalf("error on grpc server: %v", err)
	}

}

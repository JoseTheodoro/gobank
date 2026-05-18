package main

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/account"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAccountServer
}

func (s *Server) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	return &pb.AccountResponse{
		B: fmt.Sprintf("%s, hello from ms-account", in.GetA()),
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
	pb.RegisterAccountServer(s, &Server{})
	fmt.Printf("ms-account: gRPC server started at localhost%s\n", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(conn); err != nil {
		log.Fatalf("error on grpc server %v", err)
	}

}

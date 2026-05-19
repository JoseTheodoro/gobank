package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "gobank/contracts/pb/customer"
)

type Server struct {
	pb.UnimplementedCustomerServer
}

func (s *Server) CreateCustomer(ctx context.Context, in *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	fmt.Printf("request received: %s", in.String())
	msg := fmt.Sprintf("%s, hello from ms-customer", in.GetName())

	return &pb.CreateCustomerResponse{
		B: msg,
	}, nil
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error on load .env file")
	}

	conn, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error on tcp connect: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, &Server{})
	fmt.Printf("ms-customer: gRPC server started at localhost%s\n", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(conn); err != nil {
		log.Fatalf("error on grpc server: %v", err)
	}

}

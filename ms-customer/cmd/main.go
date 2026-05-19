package main

import (
	"context"
	"fmt"
	"gobank/ms-customer/internal/infrastructure/database/postgres"
	"gobank/ms-customer/internal/services"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "gobank/contracts/pb/customer"
	"gobank/ms-customer/internal/infrastructure/handler"
	"gobank/pkg/database"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error on load .env file")
	}

	db, err := database.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatalf("error on conect postgres > %v", err)
	}
	defer db.Close()

	conn, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error on tcp connect: %v", err)
	}

	customerRepo := postgres.NewCustomerRepositoryPostgres(db)
	customerSrv := services.NewCustomerService(customerRepo)

	h := handler.NewHandle(customerSrv)

	s := grpc.NewServer()
	pb.RegisterCustomerServer(s, h)
	fmt.Printf("ms-customer: gRPC server started at localhost%s\n", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(conn); err != nil {
		log.Fatalf("error on grpc server: %v", err)
	}

}

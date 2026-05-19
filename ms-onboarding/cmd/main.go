package main

import (
	"context"
	"fmt"
	"gobank/ms-onboarding/internal/infraestructure/database/postgres"
	"gobank/ms-onboarding/internal/infraestructure/handler"
	"gobank/ms-onboarding/internal/services"
	"gobank/pkg/database"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbAccount "gobank/contracts/pb/account"
	pbAuth "gobank/contracts/pb/auth"
	pbCustomer "gobank/contracts/pb/customer"
	pbKYC "gobank/contracts/pb/kyc"
	pb "gobank/contracts/pb/onboarding"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error on load .env file")
	}

	server, err := net.Listen("tcp", os.Getenv("ADDR_LISTEN"))
	if err != nil {
		log.Fatalf("error to listen: %v\n", err)
	}

	// new client ms-customer
	cCustomer, err := grpc.NewClient(os.Getenv("CUSTOMER_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on ms-customer client: %v", err)
	}
	// new client ms-auth
	cAuth, err := grpc.NewClient(os.Getenv("AUTH_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on ms-auth client: %v", err)
	}
	// new client ms-account
	cAccount, err := grpc.NewClient(os.Getenv("ACCOUNT_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on ms-account client: %v", err)
	}
	// new client ms-kyc
	cKYC, err := grpc.NewClient(os.Getenv("KYC_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on ms-kyc client: %v", err)
	}

	conn, err := database.Connect(context.Background(), os.Getenv("DATABASE_DSN"))
	if err != nil {
		log.Fatalf("error on connect postgres: %v", err)
	}
	defer conn.Close()

	//clients
	customerClient := pbCustomer.NewCustomerClient(cCustomer)
	authClient := pbAuth.NewAuthClient(cAuth)
	accountClient := pbAccount.NewAccountClient(cAccount)
	KYCClient := pbKYC.NewKYCClient(cKYC)

	// services
	customerService := services.NewCustomerService(customerClient)
	authService := services.NewAuthService(authClient)
	accountService := services.NewAccountService(accountClient)
	KYCService := services.NewKYCService(KYCClient)

	// postgres
	onboardingRepository := postgres.NewOnboardingPostgress(conn)

	s := grpc.NewServer()
	onboardingService := services.NewOnboardingService(customerService, authService, accountService, KYCService, onboardingRepository)
	h := handler.NewHandle(onboardingService)

	pb.RegisterOnboardingServer(s, h)
	fmt.Printf("ms-onboarding: gRPC server started at localhost%s", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(server); err != nil {
		log.Fatalf("error on grpc server: %v\n", err)
	}

}

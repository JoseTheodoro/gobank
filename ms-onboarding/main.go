package main

import (
	"context"
	"fmt"
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

type Server struct {
	pb.UnimplementedOnboardingServer
	customerClient pbCustomer.CustomerClient
	authClient     pbAuth.AuthClient
	accountClient  pbAccount.AccountClient
	KYCClient      pbKYC.KYCClient
}

func (s *Server) StartOnboarding(ctx context.Context, in *pb.OnboardingRequest) (*pb.OnboardingResponse, error) {

	// request to ms-customer
	reqCreateCustomer := pbCustomer.CreateCustomerRequest{
		Name: in.CustomerInfo.Name,
	}
	reqCreateCredentials := pbAuth.AuthRequest{
		Name: in.CustomerInfo.Name,
	}
	reqCreateAccount := pbAccount.AccountRequest{
		A: in.CustomerInfo.Name,
	}
	reqStartKYC := pbKYC.StartKYCRequest{
		Name: in.CustomerInfo.Name,
	}

	c, err := s.customerClient.CreateCustomer(ctx, &reqCreateCustomer)
	if err != nil {
		log.Printf("error on creating customer: %v", err)
		return nil, err
	}
	// request to ms-auth
	cre, err := s.authClient.CreateCredentials(ctx, &reqCreateCredentials)
	if err != nil {
		fmt.Printf("error on creating credentials: %v", err)
		return nil, err
	}
	// request to ms-account
	a, err := s.accountClient.CreateAccount(ctx, &reqCreateAccount)
	if err != nil {
		fmt.Printf("error on creating account: %v", err)
		return nil, err
	}
	// request to ms-kyc
	k, err := s.KYCClient.StartKYC(ctx, &reqStartKYC)
	if err != nil {
		fmt.Printf("error on starting kyc: %v", err)
		return nil, err
	}

	fmt.Printf("ms-customer=%s, ms-auth=%s, ms-account=%s, ms-kyc=%s", c.GetB(), cre.GetB(), a.GetB(), k.GetB())

	return &pb.OnboardingResponse{Message: "Onboarding started"}, nil
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

	customerClient := pbCustomer.NewCustomerClient(cCustomer)
	authClient := pbAuth.NewAuthClient(cAuth)
	accountClient := pbAccount.NewAccountClient(cAccount)
	KYCClient := pbKYC.NewKYCClient(cKYC)

	s := grpc.NewServer()
	pb.RegisterOnboardingServer(s, &Server{customerClient: customerClient, authClient: authClient, accountClient: accountClient, KYCClient: KYCClient})
	fmt.Printf("ms-onboarding: gRPC server started at localhost%s", os.Getenv("ADDR_LISTEN"))
	if err := s.Serve(server); err != nil {
		log.Fatalf("error on grpc server: %v\n", err)
	}

}

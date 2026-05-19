package handler

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/onboarding"
	"gobank/ms-onboarding/internal/domain"
	"gobank/ms-onboarding/internal/services"
)

type Handle struct {
	pb.UnimplementedOnboardingServer
	OnboardingService *services.OnboardingService
}

func NewHandle(o *services.OnboardingService) *Handle {
	return &Handle{
		OnboardingService: o,
	}
}

func (s *Handle) StartOnboarding(ctx context.Context, in *pb.OnboardingRequest) (*pb.OnboardingResponse, error) {

	fmt.Println("request received: ", in.String())

	input, err := s.validate(in)
	if err != nil {
		return &pb.OnboardingResponse{Message: "Bad Request"}, err
	}

	if err := s.OnboardingService.Run(ctx, input); err != nil {
		return &pb.OnboardingResponse{Message: "Failed to run onboarding"}, fmt.Errorf("error on running onboarding > %w", err)
	}

	return &pb.OnboardingResponse{Message: "Onboading Started Successful"}, nil
}

func (s *Handle) validate(r *pb.OnboardingRequest) (domain.StartOnboardingInput, error) {

	customer := domain.CustomerInput{
		Name:     r.CustomerInfo.GetName(),
		Document: r.CustomerInfo.GetDocument(),
		Type:     r.CustomerInfo.GetType(),
	}

	credentials := domain.CredentialsInput{
		Email:    r.AccountCredentials.GetEmail(),
		Password: r.AccountCredentials.GetPassowrd(),
	}

	device := domain.DeviceInput{
		IPAddr:    r.DeviceInfo.GetIPAddr(),
		UserAgent: r.DeviceInfo.GetUserAgent(),
		DeviceID:  r.DeviceInfo.GetDeviceID(),
	}

	return domain.StartOnboardingInput{
		Customer:    customer,
		Credentials: credentials,
		Device:      device,
	}, nil

}

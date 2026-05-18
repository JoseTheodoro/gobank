package main

import (
	"context"
	"fmt"
	pb "gobank/contracts/pb/onboarding"
	"net/http"
)

type Handle struct {
	onboardingClient pb.OnboardingClient
}

func NewHandle(ob pb.OnboardingClient) *Handle {
	return &Handle{
		onboardingClient: ob,
	}
}

func (h *Handle) handleOnboarding(w http.ResponseWriter, _ *http.Request) {

	startingOnboardingRequest := pb.OnboardingRequest{
		M: "starting onboarding from api-gateway",
	}
	res, err := h.onboardingClient.StartOnboarding(context.Background(), &startingOnboardingRequest)
	if err != nil {
		fmt.Printf("cannot send request for starting onboarding: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.GetA()))

}

package main

import (
	"encoding/json"
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

func (h *Handle) handleOnboarding(w http.ResponseWriter, r *http.Request) {

	var obRequest OnboardingRequest
	json.NewDecoder(r.Body).Decode(&obRequest)

	fmt.Printf("OnboardingRequest: %#v", obRequest)

}

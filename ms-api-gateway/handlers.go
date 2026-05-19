package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "gobank/contracts/pb/onboarding"
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

	var obRequest OnboardingRequestt
	json.NewDecoder(r.Body).Decode(&obRequest)

	req := pb.OnboardingRequest{
		CustomerInfo: &pb.CustomerInfo{
			Name:     obRequest.CustomerInfo.Name,
			Document: obRequest.CustomerInfo.Document,
			Type:     string(INDIVIDUAL),
		},
		AccountCredentials: &pb.AccountCredentials{
			Email:    obRequest.AccountCredentials.Email,
			Passowrd: obRequest.AccountCredentials.Password,
		},
		DeviceInfo: &pb.DeviceInfo{
			IPAddr:    obRequest.DeviceInfo.IPAddr,
			UserAgent: obRequest.DeviceInfo.UserAgent,
			DeviceID:  obRequest.DeviceInfo.DeviceID,
		},
	}

	res, err := h.onboardingClient.StartOnboarding(context.Background(), &req)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(res.Message))

}

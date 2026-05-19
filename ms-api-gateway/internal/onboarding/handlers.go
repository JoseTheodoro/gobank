package onboarding

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

func (h *Handle) Start(w http.ResponseWriter, r *http.Request) {

	var onboardingRequest Request
	json.NewDecoder(r.Body).Decode(&onboardingRequest)

	req := pb.OnboardingRequest{
		CustomerInfo: &pb.CustomerInfo{
			Name:     onboardingRequest.CustomerInfo.Name,
			Document: onboardingRequest.CustomerInfo.Document,
			Type:     string(INDIVIDUAL),
		},
		AccountCredentials: &pb.AccountCredentials{
			Email:    onboardingRequest.AccountCredentials.Email,
			Passowrd: onboardingRequest.AccountCredentials.Password,
		},
		DeviceInfo: &pb.DeviceInfo{
			IPAddr:    onboardingRequest.DeviceInfo.IPAddr,
			UserAgent: onboardingRequest.DeviceInfo.UserAgent,
			DeviceID:  onboardingRequest.DeviceInfo.DeviceID,
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

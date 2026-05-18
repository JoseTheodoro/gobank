package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gobank/contracts/pb/onboarding"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	ctx := context.Background()
	ctxSignal, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()

	// client onboarding
	obConn, err := grpc.NewClient(os.Getenv("ONBOARDING_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on connect onboarding server: %v", err)
	}
	defer obConn.Close()
	obClient := pb.NewOnboardingClient(obConn)

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    os.Getenv("ADDR_LISTEN"),
		Handler: mux,
	}

	h := NewHandle(obClient)

	mux.HandleFunc("POST /api/v1/onboarding", h.handleOnboarding)

	fmt.Printf("http sever is running at localhost%s\n", os.Getenv("ADDR_LISTEN"))

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Default().Printf("error received: %v\n", err)
		}
	}()

	<-ctxSignal.Done()

	ctxTimemout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxTimemout); err != nil {
		log.Default().Printf("error on graceful shutdown: %v", err)
	}

}

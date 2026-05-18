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
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	ctx := context.Background()
	ctxSignal, stop := signal.NotifyContext(ctx, syscall.SIGTERM)
	defer stop()

	mux := http.NewServeMux()

	server := http.Server{
		Addr:    os.Getenv("ADDR_LISTEN"),
		Handler: mux,
	}

	mux.HandleFunc("POST /api/v1/onboarding", handleOnboarding)

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

func handleOnboarding(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Onboarding"))

}

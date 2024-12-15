package main

import (
	"context"
	"fmt"
	"github.com/k33pup/Booking/payment_svc/internal/pkg/app"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	bs, err := app.NewPaymentService()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to create payment service: %v\n", err)
		return
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	doneChan := make(chan struct{})

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		if err := bs.Start(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting payment service: %v\n", err)
		}
	}()

	go func() {
		<-signalChan
		fmt.Println("Received termination signal. Shutting down...")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer shutdownCancel()

		if err := bs.Stop(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "Error during shutdown: %v\n", err)
		}

		close(doneChan)
		cancel()
	}()

	wg.Wait()
	<-doneChan

	fmt.Println("Service stopped gracefully.")
}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/k33pup/Booking/notification_svc/internal/pkg/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	nService, err := app.NewNotificationService()
	if err != nil {
		panic(err)
	}

	if err := nService.Start(ctx); err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("Received termination signal. Shutting down...")
		if err := nService.Stop(); err != nil {
			fmt.Println("Error during shutdown: ", err.Error())
		}
		cancel()
	}()
}

package main

import (
	"context"
	"fmt"
	"github.com/k33pup/Booking.git/internal/hotel/pkg/app"
	"github.com/k33pup/Booking.git/internal/hotel/pkg/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewConfig()

	hotelService := app.NewHotelService(cfg)

	if err := hotelService.Init(ctx); err != nil {
		panic(err)
	}

	if err := hotelService.Start(ctx); err != nil {
		panic(err)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("Received termination signal. Shutting down...")
		if err := hotelService.Stop(ctx); err != nil {
			fmt.Println("Error during shutdown: ", err.Error())
		}
		cancel()
	}()

	hotelService.Wait(ctx)
	fmt.Println("Hotel service stopped. Exiting...")
}

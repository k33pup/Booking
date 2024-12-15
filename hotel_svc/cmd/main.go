package main

import (
	"booking/hotel_svc/internal/pkg/app"
	"booking/hotel_svc/internal/pkg/config"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg := config.NewConfig()

	hotelService := app.NewHotelService(cfg)

	if err := hotelService.Init(ctx); err != nil {
		panic(err)
	}

	if err := hotelService.Start(ctx); err != nil {
		panic(err)
	}

	fmt.Println("hotel service started...")

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

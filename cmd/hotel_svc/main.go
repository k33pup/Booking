package main

import (
	"context"
	"github.com/k33pup/Booking.git/internal/hotel/pkg/app"
	logBase "log"
)

func main() {
	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := app.NewHotelService()

	if err := service.Init(ctxWithCancel); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := service.Start(ctxWithCancel); err != nil {
		logBase.Fatal("start app failed")
	}

	if err := service.Stop(ctxWithCancel); err != nil {
		logBase.Fatalf("stop app failed")
	}
}

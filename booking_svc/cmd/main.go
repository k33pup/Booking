package main

import (
	"context"
	"github.com/k33pup/Booking.git/pkg/app"
)

func main() {
	bs, err := app.NewBookingService()
	if err != nil {
		panic(err)
	}
	err = bs.Start(context.Background())
	if err != nil {
		panic(err)
	}
}

package app

import (
	"context"
	"github.com/k33pup/Booking.git/internal/booking/config"
	"github.com/k33pup/Booking.git/internal/booking/pkg/api/http"
	"github.com/k33pup/Booking.git/internal/booking/pkg/repository"
	"github.com/k33pup/Booking.git/internal/booking/usecases"
)

type BookingService struct {
	server  *http.Server
	useCase usecases.IBookedRoomUseCase
	repo    *repository.BookedRoomRepository
}

func NewBookingService() (*BookingService, error) {
	dsn, err := config.GetDsnString()
	if err != nil {
		return nil, err
	}
	repo, err := repository.NewBookedRoomRepository(dsn)
	if err != nil {
		return nil, err
	}
	useCase := usecases.NewBookedRoomUseCase(repo)
	server := http.NewServer(useCase)
	return &BookingService{server: server, useCase: useCase, repo: repo}, nil
}

func (b *BookingService) Start(ctx context.Context) error {
	err := b.server.Start(ctx)
	if err != nil {
		return err
	}
	return nil
}
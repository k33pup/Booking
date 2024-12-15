package app

import (
	"context"
	my_http "github.com/k33pup/Booking/internal/pkg/api/http"
	"github.com/k33pup/Booking/internal/pkg/config"
	"github.com/k33pup/Booking/internal/pkg/logger"
	"github.com/k33pup/Booking/internal/pkg/repository"
	"github.com/k33pup/Booking/internal/usecases"
	"log/slog"
)

type BookingService struct {
	server *my_http.Server
	repo   *repository.BookedRoomRepository
	Log    *slog.Logger
}

func NewBookingService() (*BookingService, error) {
	log, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	dsn, err := config.GetDsnString()
	if err != nil {
		return nil, err
	}
	repo, err := repository.NewBookedRoomRepository(dsn)
	if err != nil {
		return nil, err
	}
	useCase := usecases.NewBookedRoomUseCase(repo)
	server := my_http.NewServer(useCase)
	return &BookingService{server: server, repo: repo, Log: log}, nil
}

func (b *BookingService) Start(ctx context.Context) error {
	b.Log.Info("Booking service started")
	err := b.server.Start()
	if err != nil {
		return err
	}
	return nil
}

func (b *BookingService) Stop(ctx context.Context) error {
	b.Log.Info("Booking service stopped")
	err := b.server.Stop(ctx)
	if err != nil {
		return err
	}
	return nil
}

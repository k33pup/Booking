package app

import (
	"context"
	"github.com/k33pup/Booking.git/internal/pkg/api/http"
	"github.com/k33pup/Booking.git/internal/pkg/config"
	"github.com/k33pup/Booking.git/internal/pkg/repository"
	"github.com/k33pup/Booking.git/internal/usecases"
	"log/slog"
	"os"
)

type BookingService struct {
	server *http.Server
	repo   *repository.BookedRoomRepository
	log    *slog.Logger
}

func NewBookingService() (*BookingService, error) {
	logFilePath, err := config.LoadLogFilePath()
	if err != nil {
		return nil, err
	}
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	handler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Устанавливаем минимальный уровень логирования
	})
	log := slog.New(handler)
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
	return &BookingService{server: server, repo: repo, log: log}, nil
}

func (b *BookingService) Start(ctx context.Context) error {
	b.log.Info("Booking service started")
	err := b.server.Start()
	if err != nil {
		return err
	}
	return nil
}

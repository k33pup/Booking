package app

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/k33pup/Booking.git/internal/hotel/pkg/repository"
	"github.com/k33pup/Booking.git/internal/hotel/usecases"
)

type HotelService struct {
	useCase *usecases.HotelUseCase
}

func NewHotelService() *HotelService {
	return &HotelService{}
}

func (h *HotelService) Init(ctx context.Context) error {
	// initialization of grpc, postgres, kafka and other api

	pgxConn := &pgxpool.Conn{}
	repo := repository.NewHotelRepository(pgxConn)

	h.useCase = usecases.NewHotelUseCase(repo)

	return nil
}

func (h *HotelService) Start(ctx context.Context) error {
	// grpc server start

	return nil
}

func (h *HotelService) Stop(ctx context.Context) error {
	//<-ctx.Done()

	// stopping service, graceful shutdown
	return nil
}

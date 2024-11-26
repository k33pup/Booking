package usecases

import (
	"context"
	"github.com/k33pup/Booking.git/internal/hotel/domain"
)

type IHotelUseCase interface {
	GetHotels(ctx context.Context) ([]*domain.Hotel, error)
	GetHotel(ctx context.Context, id string) (*domain.Hotel, error)
	GetRooms(ctx context.Context, hotel domain.Hotel) ([]*domain.Room, error)
	GetRoom(ctx context.Context, id string) (*domain.Room, error)
}

type HotelUseCase struct {
	repo IHotelRepository
}

func NewHotelUseCase(repo IHotelRepository) *HotelUseCase {
	return &HotelUseCase{repo: repo}
}

func (h *HotelUseCase) GetHotels(ctx context.Context) ([]*domain.Hotel, error) {
	return h.repo.GetHotels(ctx)
}

func (h *HotelUseCase) GetHotel(ctx context.Context, id string) (*domain.Hotel, error) {
	return h.repo.GetHotel(ctx, id)
}

package usecases

import (
	"context"
	"github.com/k33pup/Booking.git/internal/booking/domain"
)

type IBookedRoomUseCase interface {
	BookRoom(ctx context.Context, room *domain.BookedRoom) error
	GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error)
	IsRoomBooked(ctx context.Context, roomID string) (bool, error)
}

type BookedRoomUseCase struct {
	repo IBookedRoomRepository
}

func NewBookedRoomUseCase(repo IBookedRoomRepository) *BookedRoomUseCase {
	return &BookedRoomUseCase{repo: repo}
}

func (br *BookedRoomUseCase) BookRoom(ctx context.Context, room *domain.BookedRoom) error {
	return br.repo.BookRoom(ctx, room)
}

func (br *BookedRoomUseCase) GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error) {
	return br.repo.GetBookedRoomsList(ctx, hotelId)
}

func (br *BookedRoomUseCase) IsRoomBooked(ctx context.Context, roomID string) (bool, error) {
	return br.repo.IsRoomBooked(ctx, roomID)
}

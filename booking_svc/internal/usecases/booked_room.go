package usecases

import (
	"context"
	"github.com/k33pup/Booking/internal/domain"
)

type IBookedRoomUseCase interface {
	ReserveRoom(ctx context.Context, room *domain.BookedRoom) error
	UnReserveRoom(ctx context.Context, roomId string) error
	ApproveRoom(ctx context.Context, roomId string) error
	GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error)
	IsRoomBooked(ctx context.Context, roomID string) (bool, error)
}

type BookedRoomUseCase struct {
	repo IBookedRoomRepository
}

func NewBookedRoomUseCase(repo IBookedRoomRepository) *BookedRoomUseCase {
	return &BookedRoomUseCase{repo: repo}
}

func (br *BookedRoomUseCase) ReserveRoom(ctx context.Context, room *domain.BookedRoom) error {
	return br.repo.ReserveRoom(ctx, room)
}

func (br *BookedRoomUseCase) ApproveRoom(ctx context.Context, roomId string) error {
	return br.repo.ApproveRoom(ctx, roomId)
}

func (br *BookedRoomUseCase) GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error) {
	return br.repo.GetBookedRoomsList(ctx, hotelId)
}

func (br *BookedRoomUseCase) IsRoomBooked(ctx context.Context, roomID string) (bool, error) {
	return br.repo.IsRoomBooked(ctx, roomID)
}

func (br *BookedRoomUseCase) UnReserveRoom(ctx context.Context, roomId string) error {
	return br.repo.UnReserveRoom(ctx, roomId)
}

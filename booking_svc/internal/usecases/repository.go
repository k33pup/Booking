package usecases

import (
	"context"
	"github.com/k33pup/Booking/internal/domain"
)

type IBookedRoomRepository interface {
	ReserveRoom(ctx context.Context, room *domain.BookedRoom) error
	ApproveRoom(ctx context.Context, roomId string) error
	GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error)
	IsRoomBooked(ctx context.Context, roomID string) (bool, error)
	UnReserveRoom(ctx context.Context, roomId string) error
}

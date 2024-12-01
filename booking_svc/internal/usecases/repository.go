package usecases

import (
	"context"
	"github.com/k33pup/Booking.git/internal/domain"
)

type IBookedRoomRepository interface {
	BookRoom(ctx context.Context, room *domain.BookedRoom) error
	GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error)
	IsRoomBooked(ctx context.Context, roomID string) (bool, error)
}

package usecases

import (
	"context"
	"github.com/k33pup/Booking.git/internal/hotel/domain"
)

type IHotelRepository interface {
	GetHotels(ctx context.Context) ([]*domain.Hotel, error)
	GetHotel(ctx context.Context, hotelId string) (*domain.Hotel, error)
	GetRooms(ctx context.Context, hotelId string) ([]*domain.Room, error)
	GetRoom(ctx context.Context, roomId string) (*domain.Room, error)
	AddHotel(ctx context.Context, hotel domain.Hotel) error
	AddRoom(ctx context.Context, room domain.Room) error
}

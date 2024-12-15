package usecases

import (
	"context"
	"github.com/k33pup/Booking/hotel_svc/internal/domain"
)

type IHotelRepository interface {
	GetHotelsList(ctx context.Context) ([]*domain.Hotel, error)
	GetHotelById(ctx context.Context, hotelId string) (*domain.Hotel, error)
	CreateHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error)
	DeleteHotel(ctx context.Context, hotelId string) error
	GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*domain.Room, error)
	GetRoomById(ctx context.Context, roomId string) (*domain.Room, error)
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	DeleteRoom(ctx context.Context, roomId string) error
}

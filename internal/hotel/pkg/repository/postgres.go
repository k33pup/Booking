package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/k33pup/Booking.git/internal/hotel/domain"
)

type HotelRepository struct {
	pgx *pgxpool.Conn
}

func NewHotelRepository(pgx *pgxpool.Conn) *HotelRepository {
	return &HotelRepository{pgx: pgx}
}

func (h *HotelRepository) GetHotels(ctx context.Context) ([]*domain.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) GetHotel(ctx context.Context, hotelId string) (*domain.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) GetRooms(ctx context.Context, hotelId string) ([]*domain.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) GetRoom(ctx context.Context, roomId string) (*domain.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) AddHotel(ctx context.Context, hotel domain.Hotel) error {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) AddRoom(ctx context.Context, room domain.Room) error {
	//TODO implement me
	panic("implement me")
}

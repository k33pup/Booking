package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/k33pup/Booking.git/internal/hotel/domain"
	"time"
)

type HotelRepository struct {
	pgx *pgxpool.Conn
}

func NewHotelRepository(pgx *pgxpool.Conn) *HotelRepository {
	return &HotelRepository{pgx: pgx}
}

func (h *HotelRepository) GetHotelsList(ctx context.Context) ([]*domain.Hotel, error) {
	//TODO implement me
	list := make([]*domain.Hotel, 0)
	list = append(list, &domain.Hotel{ID: "1", Name: "Сочи да"})
	time.Sleep(15 * time.Second)
	return list, nil
	panic("implement me")
}

func (h *HotelRepository) GetHotelById(ctx context.Context, hotelId string) (*domain.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) CreateHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) DeleteHotel(ctx context.Context, hotelId string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*domain.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) GetRoomById(ctx context.Context, roomId string) (*domain.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	//TODO implement me
	panic("implement me")
}

func (h *HotelRepository) DeleteRoom(ctx context.Context, roomId string) error {
	//TODO implement me
	panic("implement me")
}

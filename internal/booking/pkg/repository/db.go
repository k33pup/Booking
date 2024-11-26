package repository

import (
	"context"
	"github.com/k33pup/Booking.git/internal/booking/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookedRoomRepository struct {
	db *gorm.DB
}

func NewBookedRoomRepository(dsn string) (*BookedRoomRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &BookedRoomRepository{db: db}, nil
}

func (r *BookedRoomRepository) BookRoom(ctx context.Context, room *domain.BookedRoom) error {
	return nil
}

func (r *BookedRoomRepository) GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error) {
	return nil, nil
}

func (r *BookedRoomRepository) IsRoomBooked(ctx context.Context, roomID string) (bool, error) {
	return false, nil
}

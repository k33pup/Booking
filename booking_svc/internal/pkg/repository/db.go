package repository

import (
	"context"
	"fmt"
	"github.com/k33pup/Booking.git/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type BookedRoomRepository struct {
	db *gorm.DB
}

func NewBookedRoomRepository(dsn string) (*BookedRoomRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.BookedRoom{})
	if err != nil {
		return nil, err
	}

	return &BookedRoomRepository{db: db}, nil
}

func (r *BookedRoomRepository) ReserveRoom(ctx context.Context, room *domain.BookedRoom) error {
	if err := r.db.Table("booked_rooms").Create(room).Error; err != nil {
		return fmt.Errorf("adding room to database: %v", err)
	}
	return nil
}

func (r *BookedRoomRepository) ApproveRoom(ctx context.Context, roomId string) error {
	// Выполняем обновление поля IsPaid
	result := r.db.WithContext(ctx).
		Table("booked_rooms").
		Where("id = ?", roomId).
		Update("is_paid", true)

	// Проверяем на ошибки
	if result.Error != nil {
		return fmt.Errorf("updating room payment status: %v", result.Error)
	}

	// Проверяем, была ли обновлена хотя бы одна строка
	if result.RowsAffected == 0 {
		return fmt.Errorf("no room found with id: %s", roomId)
	}

	return nil
}

func (r *BookedRoomRepository) GetBookedRoomsList(ctx context.Context, hotelId string) ([]domain.BookedRoom, error) {
	var bookedRooms []domain.BookedRoom

	if err := r.db.Table("booked_rooms").Find(&bookedRooms).Error; err != nil {
		return nil, fmt.Errorf("retrieving rooms list from database: %v", err)
	}

	return bookedRooms, nil
}

func (r *BookedRoomRepository) IsRoomBooked(ctx context.Context, roomID string) (bool, error) {
	var count int64
	currentDate := time.Now()

	if err := r.db.Table("booked_rooms").
		Where("id = ? AND exit > ?", roomID, currentDate).
		Count(&count).Error; err != nil {
		return false, fmt.Errorf("checking room booking_svc status: %v", err)
	}

	return count > 0, nil
}

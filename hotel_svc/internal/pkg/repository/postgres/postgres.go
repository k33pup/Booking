package postgres

import (
	"booking/hotel_svc/internal/domain"
	"booking/hotel_svc/internal/usecases"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type HotelRepository struct {
	db *gorm.DB
}

func NewHotelRepository(dsn string) (*HotelRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&HotelModel{}, &RoomModel{})
	if err != nil {
		return nil, err
	}
	return &HotelRepository{db: db}, nil
}

func (h *HotelRepository) generateUUID() string {
	return uuid.NewString()
}

func (h *HotelRepository) GetHotelsList(ctx context.Context) ([]*domain.Hotel, error) {
	var hotels []*HotelModel
	if err := h.db.Table("hotels").Find(&hotels).Error; err != nil {
		return nil, fmt.Errorf("error while getting all hotels: %v", err)
	}
	return ToListDomainHotel(hotels), nil
}

func (h *HotelRepository) GetHotelById(ctx context.Context, hotelId string) (*domain.Hotel, error) {
	var hotel *HotelModel
	if err := h.db.Table("hotels").Where("id = ?", hotelId).First(&hotel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: hotel with id %q", usecases.ErrHotelNotFound, hotelId)
		}
		return nil, fmt.Errorf("error while getting hotel by id %q: %v", hotelId, err)
	}
	return ToDomainHotel(hotel), nil
}

func (h *HotelRepository) CreateHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	hotel.Id = h.generateUUID()
	if err := h.db.Table("hotels").Create(ToRepositoryHotel(hotel)).Error; err != nil {
		return nil, fmt.Errorf("error while creating hotel: %v", err)
	}
	return hotel, nil
}

func (h *HotelRepository) DeleteHotel(ctx context.Context, hotelId string) error {
	result := h.db.Table("hotels").Where("id = ?", hotelId).Delete(&HotelModel{})
	if result.Error != nil {
		return fmt.Errorf("error while deleting hotel with id %s: %v", hotelId, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: hotel with id %s not found", usecases.ErrHotelNotFound, hotelId)
	}
	return nil
}

func (h *HotelRepository) GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*domain.Room, error) {
	if _, err := h.GetHotelById(ctx, hotelId); err != nil {
		if errors.Is(err, usecases.ErrHotelNotFound) {
			return nil, fmt.Errorf("%w: hotel with id %q", usecases.ErrHotelNotFound, hotelId)
		}
		return nil, fmt.Errorf("error while finding hotel by id %q: %v", hotelId, err)
	}
	var roomsList []*RoomModel
	if err := h.db.Table("rooms").Where("hotel_id = ?", hotelId).Find(&roomsList).Error; err != nil {
		return nil, fmt.Errorf("error while getting rooms for hotel with id %s: %v", hotelId, err)
	}
	// may be error if list is empty
	return ToListDomainRoom(roomsList), nil
}

func (h *HotelRepository) GetRoomById(ctx context.Context, roomId string) (*domain.Room, error) {
	var room *RoomModel
	if err := h.db.Table("rooms").Where("id = ?", roomId).First(&room).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: room with id %q", usecases.ErrRoomNotFound, roomId)
		}
		return nil, fmt.Errorf("error while getting room by id %q: %v", roomId, err)
	}
	return ToDomainRoom(room), nil
}

func (h *HotelRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	room.Id = h.generateUUID()
	if err := h.db.Table("rooms").Create(ToRepositoryRoom(room)).Error; err != nil {
		return nil, fmt.Errorf("error while creating room: %v", err)
	}
	return room, nil
}

func (h *HotelRepository) DeleteRoom(ctx context.Context, roomId string) error {
	result := h.db.Table("rooms").Where("id = ?", roomId).Delete(&RoomModel{})
	if result.Error != nil {
		return fmt.Errorf("error while deleting room with id %s: %v", roomId, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: room with id %s not found", usecases.ErrRoomNotFound, roomId)
	}
	return nil
}

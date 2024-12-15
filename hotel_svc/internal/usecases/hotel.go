package usecases

import (
	"booking/hotel_svc/internal/domain"
	"context"
	"errors"
)

type IHotelUseCase interface {
	GetHotelsList(ctx context.Context) ([]*domain.Hotel, error)
	GetHotelById(ctx context.Context, hotelId string) (*domain.Hotel, error)
	CreateHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error)
	DeleteHotel(ctx context.Context, hotelId string) error
	GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*domain.Room, error)
	GetRoomById(ctx context.Context, roomId string) (*domain.Room, error)
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	DeleteRoom(ctx context.Context, roomId string) error
}

type HotelUseCase struct {
	repo IHotelRepository
}

func NewHotelUseCase(repo IHotelRepository) *HotelUseCase {
	return &HotelUseCase{repo: repo}
}

var ErrHotelNotFound = errors.New("hotel not found")
var ErrRoomNotFound = errors.New("room not found")
var ErrRoomsNotFound = errors.New("rooms not found")

func (h *HotelUseCase) GetHotelsList(ctx context.Context) ([]*domain.Hotel, error) {
	return h.repo.GetHotelsList(ctx)
}

func (h *HotelUseCase) GetHotelById(ctx context.Context, hotelId string) (*domain.Hotel, error) {
	return h.repo.GetHotelById(ctx, hotelId)
}

func (h *HotelUseCase) CreateHotel(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
	return h.repo.CreateHotel(ctx, hotel)
}

func (h *HotelUseCase) DeleteHotel(ctx context.Context, hotelId string) error {
	return h.repo.DeleteHotel(ctx, hotelId)
}

func (h *HotelUseCase) GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*domain.Room, error) {
	return h.repo.GetRoomsByHotelId(ctx, hotelId)
}

func (h *HotelUseCase) GetRoomById(ctx context.Context, roomId string) (*domain.Room, error) {
	return h.repo.GetRoomById(ctx, roomId)
}

func (h *HotelUseCase) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	return h.repo.CreateRoom(ctx, room)
}

func (h *HotelUseCase) DeleteRoom(ctx context.Context, roomId string) error {
	return h.repo.DeleteRoom(ctx, roomId)
}

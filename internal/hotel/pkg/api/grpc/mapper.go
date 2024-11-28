package grpc

import (
	"github.com/k33pup/Booking.git/internal/hotel/domain"
	apiv1pb "github.com/k33pup/Booking.git/internal/hotel/pkg/api/generated"
	"github.com/k33pup/Booking.git/pkg/models"
)

func ToModelsHotel(hotel *domain.Hotel) *models.Hotel {
	return &models.Hotel{
		ID:          hotel.ID,
		Name:        hotel.Name,
		Description: hotel.Description,
		City:        hotel.City,
		Address:     hotel.Address,
	}
}

func ToApiv1pbHotel(hotel *domain.Hotel) *apiv1pb.Hotel {
	return &apiv1pb.Hotel{
		HotelId:     hotel.ID,
		Name:        hotel.Name,
		Description: hotel.Description,
		City:        hotel.City,
		Address:     hotel.Address,
	}
}

func ToDomainHotel(in *apiv1pb.CreateHotelRequest) *domain.Hotel {
	return &domain.Hotel{
		Name:        in.Name,
		Description: in.Description,
		City:        in.City,
		Address:     in.Address,
	}
}

func ToApiv1pbRoom(room *domain.Room) *apiv1pb.Room {
	return &apiv1pb.Room{
		RoomId:      room.Id,
		HotelId:     room.HotelId,
		Name:        room.Name,
		Description: room.Description,
		Price:       room.Price,
	}
}

func ToDomainRoom(in *apiv1pb.CreateRoomRequest) *domain.Room {
	return &domain.Room{
		HotelId:     in.HotelId,
		Name:        in.Name,
		Description: in.Description,
		Price:       in.Price,
	}
}

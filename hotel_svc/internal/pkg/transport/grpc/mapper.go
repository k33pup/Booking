package grpc

import (
	"booking/hotel_svc/internal/domain"
	apiv1pb "booking/hotel_svc/internal/pkg/transport/grpc/generated"
	"booking/hotel_svc/pkg/models"
)

func ToModelsHotel(hotel *domain.Hotel) *models.Hotel {
	return &models.Hotel{
		Id:          hotel.Id,
		Name:        hotel.Name,
		Description: hotel.Description,
		City:        hotel.City,
		Address:     hotel.Address,
		Contacts:    hotel.Contacts,
	}
}

func ToApiv1pbHotel(hotel *domain.Hotel) *apiv1pb.Hotel {
	return &apiv1pb.Hotel{
		HotelId:     hotel.Id,
		Name:        hotel.Name,
		Description: hotel.Description,
		City:        hotel.City,
		Address:     hotel.Address,
		Contacts:    hotel.Contacts,
	}
}

func ToDomainHotel(in *apiv1pb.CreateHotelRequest) *domain.Hotel {
	return &domain.Hotel{
		Name:        in.Name,
		Description: in.Description,
		City:        in.City,
		Address:     in.Address,
		Contacts:    in.Contacts,
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

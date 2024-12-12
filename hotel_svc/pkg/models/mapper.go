package models

import (
	"booking/hotel_svc/internal/pkg/transport/grpc/generated"
)

func ToModelHotel(hotel *generated.Hotel) *Hotel {
	return &Hotel{
		Id:          hotel.HotelId,
		Name:        hotel.Name,
		Description: hotel.Description,
		City:        hotel.City,
		Address:     hotel.Address,
		Contacts:    hotel.Contacts,
	}
}

func ToListModelHotel(hotels []*generated.Hotel) []*Hotel {
	var hotelsList []*Hotel
	for _, hotel := range hotels {
		addHotel := ToModelHotel(hotel)
		hotelsList = append(hotelsList, addHotel)
	}
	return hotelsList
}

func ToProtocHotel(model *Hotel) *generated.Hotel {
	return &generated.Hotel{
		HotelId:     model.Id,
		Name:        model.Name,
		Description: model.Description,
		City:        model.City,
		Address:     model.Address,
		Contacts:    model.Contacts,
	}
}

func ToCreateHotelRequest(model *Hotel) *generated.CreateHotelRequest {
	return &generated.CreateHotelRequest{
		Name:        model.Name,
		Description: model.Description,
		City:        model.City,
		Address:     model.Address,
		Contacts:    model.Contacts,
	}
}

func ToModelRoom(room *generated.Room) *Room {
	return &Room{
		Id:          room.RoomId,
		HotelId:     room.HotelId,
		Name:        room.Name,
		Description: room.Description,
		Price:       room.Price,
	}
}

func ToCreateRoomRequest(model *Room) *generated.CreateRoomRequest {
	return &generated.CreateRoomRequest{
		HotelId:     model.HotelId,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
	}
}

func ToListModelRoom(rooms []*generated.Room) []*Room {
	var roomsList []*Room
	for _, room := range rooms {
		addRoom := ToModelRoom(room)
		roomsList = append(roomsList, addRoom)
	}
	return roomsList
}

func ToProtocRoom(room *Room) *generated.Room {
	return &generated.Room{
		RoomId:      room.Id,
		HotelId:     room.HotelId,
		Name:        room.Name,
		Description: room.Description,
		Price:       room.Price,
	}
}

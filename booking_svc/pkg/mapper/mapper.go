package mapper

import (
	openapi "github.com/k33pup/Booking/internal/pkg/api/http/generated_api/generated_client"
	"github.com/k33pup/Booking/pkg/models"
)

func ToOpenApiBookedRoom(bookedRoom models.BookedRoom) openapi.BookedRoom {
	return openapi.BookedRoom{
		ID:      bookedRoom.ID,
		HotelID: bookedRoom.HotelID,
		Entry:   bookedRoom.Entry,
		Exit:    bookedRoom.Exit,
		IsPaid:  bookedRoom.IsPaid,
	}
}

func ToModelsBookedRoom(bookedRoom openapi.BookedRoom) models.BookedRoom {
	return models.BookedRoom{
		ID:      bookedRoom.ID,
		HotelID: bookedRoom.HotelID,
		Entry:   bookedRoom.Entry,
		Exit:    bookedRoom.Exit,
		IsPaid:  bookedRoom.IsPaid,
		Email:   bookedRoom.Email,
	}
}

func ToModelsRoom(bookedRoom openapi.Room) models.Room {
	return models.Room{
		Id:          bookedRoom.ID,
		HotelId:     bookedRoom.HotelID,
		Name:        bookedRoom.Name,
		Description: *bookedRoom.Description,
		Price:       uint64(bookedRoom.Price),
	}
}

func ToModelsBookedRoomList(bookedRooms []openapi.BookedRoom) []models.BookedRoom {
	bookedRoomList := make([]models.BookedRoom, 0, len(bookedRooms))
	for idx, bookedRoom := range bookedRooms {
		bookedRoomList[idx] = ToModelsBookedRoom(bookedRoom)
	}
	return bookedRoomList
}

func ToModelsRoomList(rooms []openapi.Room) []models.Room {
	roomList := make([]models.Room, 0, len(rooms))
	for idx, room := range rooms {
		roomList[idx] = ToModelsRoom(room)
	}
	return roomList
}

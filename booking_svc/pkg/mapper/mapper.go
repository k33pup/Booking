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

func ToModelsBookedRoomList(bookedRoom []openapi.BookedRoom) []models.BookedRoom {
	var bookedRoomList []models.BookedRoom
	for _, bookedRoom := range bookedRoom {
		bookedRoomList = append(bookedRoomList, ToModelsBookedRoom(bookedRoom))
	}
	return bookedRoomList
}

func ToModelsRoomList(bookedRoom []openapi.Room) []models.Room {
	var roomList []models.Room
	for _, room := range bookedRoom {
		roomList = append(roomList, ToModelsRoom(room))
	}
	return roomList
}

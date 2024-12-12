package openapi

import (
	"github.com/k33pup/Booking.git/internal/domain"
)

func ToDomainBookedRoom(bookedRoom BookedRoom) *domain.BookedRoom {
	return &domain.BookedRoom{
		HotelID: bookedRoom.HotelID,
		ID:      bookedRoom.ID,
		Entry:   bookedRoom.Entry,
		Exit:    bookedRoom.Exit,
		Email:   bookedRoom.Email,
		IsPaid:  bookedRoom.IsPaid,
	}
}

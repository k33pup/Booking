package grpc

import (
	"github.com/k33pup/Booking.git/internal/hotel/domain"
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

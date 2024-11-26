package grpc

import (
	"context"
	"github.com/k33pup/Booking.git/internal/hotel/usecases"
	"github.com/k33pup/Booking.git/pkg/models"
)

type Server struct {
	HotelUseCase usecases.IHotelUseCase
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetHotel(ctx context.Context, in *api.HotelRequest) (*models.Hotel, error) {
	hotel, err := s.HotelUseCase.GetHotel(ctx, in.GetID())
	if err != nil {
		// TODO log error
		return nil, nil
	}
	return ToModelsHotel(hotel), nil
}

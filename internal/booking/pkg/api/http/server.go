package http

import "github.com/k33pup/Booking.git/internal/booking/usecases"

type Server struct {
	BookedRoomUseCase usecases.IBookedRoomUseCase
}

func NewServer(BookedRoomUseCase usecases.IBookedRoomUseCase) *Server {
	return &Server{BookedRoomUseCase}
}

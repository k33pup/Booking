package grpc

import (
	"booking/hotel_svc/internal/pkg/api/generated"
	"booking/hotel_svc/internal/usecases"
	"context"
	"errors"
)

type Server struct {
	HotelUseCase usecases.IHotelUseCase
	generated.UnimplementedHotelServiceServer
}

func NewServer(hotelUseCase usecases.IHotelUseCase) *Server {
	return &Server{
		HotelUseCase: hotelUseCase,
	}
}

func (s *Server) GetHotelsList(ctx context.Context, _ *generated.GetHotelsListRequest) (*generated.GetHotelsListResponse, error) {
	hotelsList := &generated.GetHotelsListResponse{
		Hotels: make([]*generated.Hotel, 0),
	}
	hotels, err := s.HotelUseCase.GetHotelsList(ctx)
	if err != nil {
		return nil, errors.New("error while getting all hotels")
	}
	for _, hotel := range hotels {
		addHotel := ToApiv1pbHotel(hotel)
		hotelsList.Hotels = append(hotelsList.Hotels, addHotel)
	}
	return hotelsList, nil
}

func (s *Server) GetHotelById(ctx context.Context, in *generated.GetHotelByIdRequest) (*generated.GetHotelByIdResponse, error) {
	hotel, err := s.HotelUseCase.GetHotelById(ctx, in.HotelId)
	if err != nil {
		return nil, errors.New("failed to find internal by id")
	}
	return &generated.GetHotelByIdResponse{Hotel: ToApiv1pbHotel(hotel)}, nil
}

func (s *Server) CreateHotel(ctx context.Context, in *generated.CreateHotelRequest) (*generated.CreateHotelResponse, error) {
	hotel, err := s.HotelUseCase.CreateHotel(ctx, ToDomainHotel(in))
	if err != nil {
		return nil, errors.New("failed to create a internal")
	}
	return &generated.CreateHotelResponse{Hotel: ToApiv1pbHotel(hotel)}, nil
}

func (s *Server) DeleteHotel(ctx context.Context, in *generated.DeleteHotelRequest) (*generated.DeleteHotelResponse, error) {
	err := s.HotelUseCase.DeleteHotel(ctx, in.HotelId)
	if err != nil {
		return &generated.DeleteHotelResponse{Success: false}, errors.New("failed to delete a internal")
	}
	return &generated.DeleteHotelResponse{Success: true}, nil
}

func (s *Server) GetRoomsByHotelId(ctx context.Context, in *generated.GetRoomsByHotelIdRequest) (*generated.GetRoomsByHotelIdResponse, error) {
	roomsList := &generated.GetRoomsByHotelIdResponse{Rooms: make([]*generated.Room, 0)}
	rooms, err := s.HotelUseCase.GetRoomsByHotelId(ctx, in.HotelId)
	if err != nil {
		return nil, errors.New("failed to get rooms from internal")
	}
	for _, room := range rooms {
		addRoom := ToApiv1pbRoom(room)
		roomsList.Rooms = append(roomsList.Rooms, addRoom)
	}
	return roomsList, nil
}

func (s *Server) GetRoomById(ctx context.Context, in *generated.GetRoomByIdRequest) (*generated.GetRoomByIdResponse, error) {
	room, err := s.HotelUseCase.GetRoomById(ctx, in.RoomId)
	if err != nil {
		return nil, errors.New("failed to find room by id")
	}
	return &generated.GetRoomByIdResponse{Room: ToApiv1pbRoom(room)}, nil
}

func (s *Server) CreateRoom(ctx context.Context, in *generated.CreateRoomRequest) (*generated.CreateRoomResponse, error) {
	room, err := s.HotelUseCase.CreateRoom(ctx, ToDomainRoom(in))
	if err != nil {
		return nil, errors.New("failed to create new room")
	}
	return &generated.CreateRoomResponse{Room: ToApiv1pbRoom(room)}, nil
}

func (s *Server) DeleteRoom(ctx context.Context, in *generated.DeleteRoomRequest) (*generated.DeleteRoomResponse, error) {
	err := s.HotelUseCase.DeleteRoom(ctx, in.RoomId)
	if err != nil {
		return &generated.DeleteRoomResponse{Success: false}, errors.New("failed to delete room")
	}
	return &generated.DeleteRoomResponse{Success: true}, nil
}

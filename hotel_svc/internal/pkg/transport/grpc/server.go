package grpc

import (
	apiv1pb "booking/hotel_svc/internal/pkg/transport/grpc/generated"
	"booking/hotel_svc/internal/pkg/transport/grpc/validation"
	"booking/hotel_svc/internal/usecases"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	HotelUseCase usecases.IHotelUseCase
	apiv1pb.UnimplementedHotelServiceServer
}

func NewServer(hotelUseCase usecases.IHotelUseCase) *Server {
	return &Server{
		HotelUseCase: hotelUseCase,
	}
}

func (s *Server) GetHotelsList(ctx context.Context, _ *apiv1pb.GetHotelsListRequest) (*apiv1pb.GetHotelsListResponse, error) {
	hotelsList := &apiv1pb.GetHotelsListResponse{
		Hotels: make([]*apiv1pb.Hotel, 0),
	}
	hotels, err := s.HotelUseCase.GetHotelsList(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	for _, hotel := range hotels {
		addHotel := ToApiv1pbHotel(hotel)
		hotelsList.Hotels = append(hotelsList.Hotels, addHotel)
	}
	return hotelsList, nil
}

func (s *Server) GetHotelById(ctx context.Context, in *apiv1pb.GetHotelByIdRequest) (*apiv1pb.GetHotelByIdResponse, error) {
	if err := validation.ValidateField(in.HotelId, "required,min=1"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	hotel, err := s.HotelUseCase.GetHotelById(ctx, in.HotelId)
	if err != nil {
		if errors.Is(err, usecases.ErrHotelNotFound) {
			return nil, status.Errorf(codes.NotFound, "hotel with id %s not found", in.HotelId)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.GetHotelByIdResponse{Hotel: ToApiv1pbHotel(hotel)}, nil
}

func (s *Server) CreateHotel(ctx context.Context, in *apiv1pb.CreateHotelRequest) (*apiv1pb.CreateHotelResponse, error) {
	if err := validation.ValidateCreateHotel(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	hotel, err := s.HotelUseCase.CreateHotel(ctx, ToDomainHotel(in))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.CreateHotelResponse{Hotel: ToApiv1pbHotel(hotel)}, nil
}

func (s *Server) DeleteHotel(ctx context.Context, in *apiv1pb.DeleteHotelRequest) (*apiv1pb.DeleteHotelResponse, error) {
	if err := validation.ValidateField(in.HotelId, "required,min=1"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	err := s.HotelUseCase.DeleteHotel(ctx, in.HotelId)
	if err != nil {
		if errors.Is(err, usecases.ErrHotelNotFound) {
			return &apiv1pb.DeleteHotelResponse{Success: false}, status.Errorf(codes.NotFound, "hotel with id %s not found", in.HotelId)
		}
		return &apiv1pb.DeleteHotelResponse{Success: false}, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.DeleteHotelResponse{Success: true}, nil
}

func (s *Server) GetRoomsByHotelId(ctx context.Context, in *apiv1pb.GetRoomsByHotelIdRequest) (*apiv1pb.GetRoomsByHotelIdResponse, error) {
	if err := validation.ValidateField(in.HotelId, "required,min=1"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	roomsList := &apiv1pb.GetRoomsByHotelIdResponse{Rooms: make([]*apiv1pb.Room, 0)}
	if _, err := s.HotelUseCase.GetHotelById(ctx, in.HotelId); err != nil {
		if errors.Is(err, usecases.ErrHotelNotFound) {
			return nil, status.Errorf(codes.NotFound, "hotel with id %s not found", in.HotelId)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	rooms, err := s.HotelUseCase.GetRoomsByHotelId(ctx, in.HotelId)
	if err != nil {
		if errors.Is(err, usecases.ErrRoomsNotFound) {
			return nil, status.Errorf(codes.NotFound, "rooms in hotel with id %s not found", in.HotelId)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	for _, room := range rooms {
		addRoom := ToApiv1pbRoom(room)
		roomsList.Rooms = append(roomsList.Rooms, addRoom)
	}
	return roomsList, nil
}

func (s *Server) GetRoomById(ctx context.Context, in *apiv1pb.GetRoomByIdRequest) (*apiv1pb.GetRoomByIdResponse, error) {
	if err := validation.ValidateField(in.RoomId, "required,min=1"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	room, err := s.HotelUseCase.GetRoomById(ctx, in.RoomId)
	if err != nil {
		if errors.Is(err, usecases.ErrRoomNotFound) {
			return nil, status.Errorf(codes.NotFound, "room with id %s not found", in.RoomId)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.GetRoomByIdResponse{Room: ToApiv1pbRoom(room)}, nil
}

func (s *Server) CreateRoom(ctx context.Context, in *apiv1pb.CreateRoomRequest) (*apiv1pb.CreateRoomResponse, error) {
	if err := validation.ValidateCreateRoom(in); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	if _, err := s.HotelUseCase.GetHotelById(ctx, in.HotelId); err != nil {
		if errors.Is(err, usecases.ErrHotelNotFound) {
			return nil, status.Errorf(codes.NotFound, "hotel with id %s not found", in.HotelId)
		}
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	room, err := s.HotelUseCase.CreateRoom(ctx, ToDomainRoom(in))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.CreateRoomResponse{Room: ToApiv1pbRoom(room)}, nil
}

func (s *Server) DeleteRoom(ctx context.Context, in *apiv1pb.DeleteRoomRequest) (*apiv1pb.DeleteRoomResponse, error) {
	if err := validation.ValidateField(in.RoomId, "required,min=1"); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}
	err := s.HotelUseCase.DeleteRoom(ctx, in.RoomId)
	if err != nil {
		if errors.Is(err, usecases.ErrRoomNotFound) {
			return &apiv1pb.DeleteRoomResponse{Success: false}, status.Errorf(codes.NotFound, "room with id %s not found", in.RoomId)
		}
		return &apiv1pb.DeleteRoomResponse{Success: false}, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	return &apiv1pb.DeleteRoomResponse{Success: true}, nil
}

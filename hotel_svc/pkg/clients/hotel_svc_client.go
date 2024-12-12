package clients

import (
	"booking/hotel_svc/internal/pkg/transport/grpc/generated"
	"booking/hotel_svc/pkg/models"
	"context"
)

type GrpcHotelClient struct {
	grpcClient generated.HotelServiceClient
}

func NewGrpcHotelClient(client generated.HotelServiceClient) *GrpcHotelClient {
	return &GrpcHotelClient{grpcClient: client}
}

func (h *GrpcHotelClient) GetHotelsList(ctx context.Context) ([]*models.Hotel, error) {
	resp, err := h.grpcClient.GetHotelsList(ctx, nil)
	if err != nil {
		return nil, err
	}
	return models.ToListModelHotel(resp.Hotels), nil
}

func (h *GrpcHotelClient) GetHotelById(ctx context.Context, hotelId string) (*models.Hotel, error) {
	resp, err := h.grpcClient.GetHotelById(ctx, &generated.GetHotelByIdRequest{HotelId: hotelId})
	if err != nil {
		return nil, err
	}
	return models.ToModelHotel(resp.Hotel), nil
}

func (h *GrpcHotelClient) CreateHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	resp, err := h.grpcClient.CreateHotel(ctx, models.ToCreateHotelRequest(hotel))
	if err != nil {
		return nil, err
	}
	return models.ToModelHotel(resp.Hotel), nil
}

func (h *GrpcHotelClient) DeleteHotel(ctx context.Context, hotelId string) (bool, error) {
	resp, err := h.grpcClient.DeleteHotel(ctx, &generated.DeleteHotelRequest{HotelId: hotelId})
	if err != nil {
		return resp.Success, err
	}
	return resp.Success, nil
}

func (h *GrpcHotelClient) GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*models.Room, error) {
	resp, err := h.grpcClient.GetRoomsByHotelId(ctx, &generated.GetRoomsByHotelIdRequest{HotelId: hotelId})
	if err != nil {
		return nil, err
	}
	return models.ToListModelRoom(resp.Rooms), nil
}

func (h *GrpcHotelClient) GetRoomById(ctx context.Context, roomId string) (*models.Room, error) {
	resp, err := h.grpcClient.GetRoomById(ctx, &generated.GetRoomByIdRequest{RoomId: roomId})
	if err != nil {
		return nil, err
	}
	return models.ToModelRoom(resp.Room), nil
}

func (h *GrpcHotelClient) CreateRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	resp, err := h.grpcClient.CreateRoom(ctx, models.ToCreateRoomRequest(room))
	if err != nil {
		return nil, err
	}
	return models.ToModelRoom(resp.Room), nil
}

func (h *GrpcHotelClient) DeleteRoom(ctx context.Context, roomId string) (bool, error) {
	resp, err := h.grpcClient.DeleteRoom(ctx, &generated.DeleteRoomRequest{RoomId: roomId})
	if err != nil {
		return resp.Success, err
	}
	return resp.Success, nil
}

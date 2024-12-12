package clients

import (
	"booking/hotel_svc/internal/pkg/transport/grpc/generated"
	"booking/hotel_svc/pkg/models"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

type GrpcHotelClient struct {
	grpcClient generated.HotelServiceClient
	conn       *grpc.ClientConn
}

func NewGrpcHotelClient(url string) (*GrpcHotelClient, error) {
	clientConn, err := grpc.NewClient(url)
	if err != nil {
		return nil, err
	}
	return &GrpcHotelClient{
		grpcClient: generated.NewHotelServiceClient(clientConn),
		conn:       clientConn,
	}, nil
}

func (g *GrpcHotelClient) Close() {
	if g.conn != nil && g.conn.GetState() != connectivity.Shutdown {
		g.conn.Close()
	}
}
func (g *GrpcHotelClient) GetHotelsList(ctx context.Context) ([]*models.Hotel, error) {
	resp, err := g.grpcClient.GetHotelsList(ctx, nil)
	if err != nil {
		return nil, err
	}
	return models.ToListModelHotel(resp.Hotels), nil
}

func (g *GrpcHotelClient) GetHotelById(ctx context.Context, hotelId string) (*models.Hotel, error) {
	resp, err := g.grpcClient.GetHotelById(ctx, &generated.GetHotelByIdRequest{HotelId: hotelId})
	if err != nil {
		return nil, err
	}
	return models.ToModelHotel(resp.Hotel), nil
}

func (g *GrpcHotelClient) CreateHotel(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	resp, err := g.grpcClient.CreateHotel(ctx, models.ToCreateHotelRequest(hotel))
	if err != nil {
		return nil, err
	}
	return models.ToModelHotel(resp.Hotel), nil
}

func (g *GrpcHotelClient) DeleteHotel(ctx context.Context, hotelId string) (bool, error) {
	resp, err := g.grpcClient.DeleteHotel(ctx, &generated.DeleteHotelRequest{HotelId: hotelId})
	if err != nil {
		return resp.Success, err
	}
	return resp.Success, nil
}

func (g *GrpcHotelClient) GetRoomsByHotelId(ctx context.Context, hotelId string) ([]*models.Room, error) {
	resp, err := g.grpcClient.GetRoomsByHotelId(ctx, &generated.GetRoomsByHotelIdRequest{HotelId: hotelId})
	if err != nil {
		return nil, err
	}
	return models.ToListModelRoom(resp.Rooms), nil
}

func (g *GrpcHotelClient) GetRoomById(ctx context.Context, roomId string) (*models.Room, error) {
	resp, err := g.grpcClient.GetRoomById(ctx, &generated.GetRoomByIdRequest{RoomId: roomId})
	if err != nil {
		return nil, err
	}
	return models.ToModelRoom(resp.Room), nil
}

func (g *GrpcHotelClient) CreateRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	resp, err := g.grpcClient.CreateRoom(ctx, models.ToCreateRoomRequest(room))
	if err != nil {
		return nil, err
	}
	return models.ToModelRoom(resp.Room), nil
}

func (g *GrpcHotelClient) DeleteRoom(ctx context.Context, roomId string) (bool, error) {
	resp, err := g.grpcClient.DeleteRoom(ctx, &generated.DeleteRoomRequest{RoomId: roomId})
	if err != nil {
		return resp.Success, err
	}
	return resp.Success, nil
}

package client

import (
	"context"
	"fmt"
	openapi "github.com/k33pup/Booking.git/internal/pkg/api/http/generated_api/generated_client"
	"net/http"
)

type Client struct {
	client *openapi.APIClient
}

func NewClient(url string) *Client {
	cfg := &openapi.Configuration{
		Host:       url,
		HTTPClient: http.DefaultClient,
	}
	client := openapi.NewAPIClient(cfg)
	return &Client{client: client}
}

func (c *Client) BookRoom(room openapi.BookedRoom) error {

	req := c.client.DefaultAPI.BookRoomPost(context.Background()).BookedRoom(room)

	_, httpResp, err := c.client.DefaultAPI.BookRoomPostExecute(req)
	if err != nil {
		return fmt.Errorf("failed to handle payment: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to handle payment, status code: %v", httpResp.StatusCode)
	}

	return nil
}

func (c *Client) GetBookedRoomsList(hotelID string) ([]openapi.BookedRoom, error) {
	req := c.client.DefaultAPI.GetBookedRooms(context.Background(), hotelID)

	resp, httpResp, err := c.client.DefaultAPI.GetBookedRoomsExecute(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get booked rooms: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to handle list of booked rooms, status code: %v", httpResp.StatusCode)
	}
	return resp, nil
}

func (c *Client) GetUnbookedRoomsList(hotelID string) ([]openapi.Room, error) {
	req := c.client.DefaultAPI.GetUnbookedRooms(context.Background(), hotelID)

	resp, httpResp, err := c.client.DefaultAPI.GetUnbookedRoomsExecute(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get unbooked rooms: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to handle list of unbooked rooms, status code: %v", httpResp.StatusCode)
	}
	return resp, nil
}

package client

import (
	"context"
	"fmt"
	"github.com/k33pup/Booking.git/internal/domain"
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

func (c *Client) HandlePayment(pay domain.Payment) error {
	// Преобразование domain.Payment в сгенерированный Payment
	payment := openapi.Payment{
		Price:  float32(pay.Price),
		RoomId: pay.RoomId,
	}

	req := c.client.DefaultAPI.HandlePayment(context.Background()).Payment(payment)

	resp, httpResp, err := c.client.DefaultAPI.HandlePaymentExecute(req)
	if err != nil {
		// Обработка ошибки
		return fmt.Errorf("failed to handle payment: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}

	fmt.Printf("Response: %v\n", resp.Message)
	return nil
}

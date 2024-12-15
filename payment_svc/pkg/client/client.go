package client

import (
	"context"
	"fmt"
	openapi "github.com/k33pup/Booking/payment_svc/internal/pkg/transport/http/generated_api/generated_client"
	mapper "github.com/k33pup/Booking/payment_svc/internal/pkg/transport/http/generated_api/generated_server/go"
	"github.com/k33pup/Booking/payment_svc/pkg/models"
	"net/http"
)

type Client struct {
	client *openapi.APIClient
}

func NewClient(url string) *Client {
	cfg := &openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{URL: url},
		},
		HTTPClient: http.DefaultClient,
	}
	client := openapi.NewAPIClient(cfg)
	return &Client{client: client}
}

func (c *Client) CreatePayment(payment models.PaymentM) error {
	apiPayment := mapper.FromDomainPayment(models.ToDomainPayment(payment))
	req := c.client.DefaultAPI.CreatePaymentPost(context.Background()).CreatePaymentPostRequest(*apiPayment)

	resp, httpResp, err := c.client.DefaultAPI.CreatePaymentPostExecute(req)
	if err != nil {
		return fmt.Errorf("failed to create payment: %v", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", httpResp.StatusCode)
	}
	fmt.Printf("Response: %s\n", resp)
	return nil
}

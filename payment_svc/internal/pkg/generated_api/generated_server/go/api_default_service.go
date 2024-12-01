// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Payment Webhook API
 *
 * API для обработки платежей по комнатам.
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService() *DefaultAPIService {
	return &DefaultAPIService{}
}

// HandlePayment - Обработка данных о платеже
func (s *DefaultAPIService) HandlePayment(ctx context.Context, payment Payment) (ImplResponse, error) {
	// TODO - update HandlePayment with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, HandlePayment200Response{}) or use other options such as http.Ok ...
	// return Response(200, HandlePayment200Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, HandlePayment400Response{}) or use other options such as http.Ok ...
	// return Response(400, HandlePayment400Response{}), nil

	// TODO: Uncomment the next line to return response Response(500, HandlePayment500Response{}) or use other options such as http.Ok ...
	// return Response(500, HandlePayment500Response{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("HandlePayment method not implemented")
}

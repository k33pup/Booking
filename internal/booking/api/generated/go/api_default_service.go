// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Hotel Booking API
 *
 * API для управления бронированием отелей
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"github.com/k33pup/Booking.git/internal/booking/usecases"
	"net/http"
	"errors"
	"time"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	useCase usecases.IBookedRoomRepository
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService(useCase usecases.IBookedRoomRepository) *DefaultAPIService {
	return &DefaultAPIService{useCase}
}

// GetUnbookedRooms - Получить список свободных комнат по ID отеля
func (s *DefaultAPIService) GetUnbookedRooms(ctx context.Context, hotelId string) (ImplResponse, error) {
	// TODO - update GetUnbookedRooms with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []Room{}) or use other options such as http.Ok ...
	// return Response(200, []Room{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	// return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUnbookedRooms method not implemented")
}

// GetBookedRooms - Получить список забронированных комнат по ID отеля
func (s *DefaultAPIService) GetBookedRooms(ctx context.Context, hotelId string) (ImplResponse, error) {
	// TODO - update GetBookedRooms with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []BookedRoom{}) or use other options such as http.Ok ...
	// return Response(200, []BookedRoom{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	// TODO: Uncomment the next line to return response Response(500, {}) or use other options such as http.Ok ...
	// return Response(500, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetBookedRooms method not implemented")
}

// BookRoomRoomIdPost - Book a room by ID
func (s *DefaultAPIService) BookRoomRoomIdPost(ctx context.Context, roomId string, hotelId string, entry time.Time, exit time.Time, email string) (ImplResponse, error) {
	// TODO - update BookRoomRoomIdPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, BookedRoom{}) or use other options such as http.Ok ...
	// return Response(201, BookedRoom{}), nil

	// TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	// return Response(400, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("BookRoomRoomIdPost method not implemented")
}
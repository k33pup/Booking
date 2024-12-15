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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service DefaultAPIServicer
	errorHandler ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) *DefaultAPIController {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
		"ApprovePaymentWebhook": Route{
			strings.ToUpper("Post"),
			"/webhook/approve-payment",
			c.ApprovePaymentWebhook,
		},
		"GetUnbookedRooms": Route{
			strings.ToUpper("Get"),
			"/unbooked-rooms/{hotel_id}",
			c.GetUnbookedRooms,
		},
		"GetBookedRooms": Route{
			strings.ToUpper("Get"),
			"/booked-rooms/{hotel_id}",
			c.GetBookedRooms,
		},
		"BookRoomPost": Route{
			strings.ToUpper("Post"),
			"/book-room/",
			c.BookRoomPost,
		},
	}
}

// ApprovePaymentWebhook - Webhook для подтверждения платежа
func (c *DefaultAPIController) ApprovePaymentWebhook(w http.ResponseWriter, r *http.Request) {
	approvePaymentWebhookRequestParam := ApprovePaymentWebhookRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&approvePaymentWebhookRequestParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertApprovePaymentWebhookRequestRequired(approvePaymentWebhookRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertApprovePaymentWebhookRequestConstraints(approvePaymentWebhookRequestParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ApprovePaymentWebhook(r.Context(), approvePaymentWebhookRequestParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetUnbookedRooms - Получить список свободных комнат по ID отеля
func (c *DefaultAPIController) GetUnbookedRooms(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotelIdParam := params["hotel_id"]
	if hotelIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"hotel_id"}, nil)
		return
	}
	result, err := c.service.GetUnbookedRooms(r.Context(), hotelIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetBookedRooms - Получить список забронированных комнат по ID отеля
func (c *DefaultAPIController) GetBookedRooms(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotelIdParam := params["hotel_id"]
	if hotelIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"hotel_id"}, nil)
		return
	}
	result, err := c.service.GetBookedRooms(r.Context(), hotelIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// BookRoomPost - Book a room by ID
func (c *DefaultAPIController) BookRoomPost(w http.ResponseWriter, r *http.Request) {
	bookedRoomParam := BookedRoom{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookedRoomParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertBookedRoomRequired(bookedRoomParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertBookedRoomConstraints(bookedRoomParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.BookRoomPost(r.Context(), bookedRoomParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

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
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// DefaultAPIController binds http requests to an generated_api service and writes the service results to the http response
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

// NewDefaultAPIController creates a default generated_api controller
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

// Routes returns all the generated_api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() Routes {
	return Routes{
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
		"BookRoomRoomIdPost": Route{
			strings.ToUpper("Post"),
			"/book-room/{room_id}",
			c.BookRoomRoomIdPost,
		},
	}
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

// BookRoomRoomIdPost - Book a room by ID
func (c *DefaultAPIController) BookRoomRoomIdPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	roomIdParam := params["room_id"]
	if roomIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"room_id"}, nil)
		return
	}
	var hotelIdParam string
	if query.Has("hotelId") {
		param := query.Get("hotelId")

		hotelIdParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "hotelId"}, nil)
		return
	}
	var entryParam time.Time
	if query.Has("Entry"){
		param, err := parseTime(query.Get("Entry"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "Entry", Err: err}, nil)
			return
		}

		entryParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{"Entry"}, nil)
		return
	}
	var exitParam time.Time
	if query.Has("Exit"){
		param, err := parseTime(query.Get("Exit"))
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "Exit", Err: err}, nil)
			return
		}

		exitParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{"Exit"}, nil)
		return
	}
	var emailParam string
	if query.Has("Email") {
		param := query.Get("Email")

		emailParam = param
	} else {
		c.errorHandler(w, r, &RequiredError{Field: "Email"}, nil)
		return
	}
	result, err := c.service.BookRoomRoomIdPost(r.Context(), roomIdParam, hotelIdParam, entryParam, exitParam, emailParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Hotel Booking API
 *
 * API for fetching and booking rooms by hotel ID.
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"net/http"
	"time"
)



// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	BookedRoomsHotelIdGet(http.ResponseWriter, *http.Request)
	UnbookedRoomsHotelIdGet(http.ResponseWriter, *http.Request)
	BookRoomRoomIdPost(http.ResponseWriter, *http.Request)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	BookedRoomsHotelIdGet(context.Context, string) (ImplResponse, error)
	UnbookedRoomsHotelIdGet(context.Context, string) (ImplResponse, error)
	BookRoomRoomIdPost(context.Context, string, time.Time, time.Time, string) (ImplResponse, error)
}

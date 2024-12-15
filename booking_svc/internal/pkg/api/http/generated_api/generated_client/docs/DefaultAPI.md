# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BookRoomPost**](DefaultAPI.md#BookRoomPost) | **Post** /book-room/ | Book a room by ID
[**GetBookedRooms**](DefaultAPI.md#GetBookedRooms) | **Get** /booked-rooms/{hotel_id} | Получить список забронированных комнат по ID отеля
[**GetUnbookedRooms**](DefaultAPI.md#GetUnbookedRooms) | **Get** /unbooked-rooms/{hotel_id} | Получить список свободных комнат по ID отеля



## BookRoomPost

> BookedRoom BookRoomPost(ctx).BookedRoom(bookedRoom).Execute()

Book a room by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	bookedRoom := *openapiclient.NewBookedRoom("ID_example", "HotelID_example", time.Now(), time.Now(), "Email_example", false) // BookedRoom | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.BookRoomPost(context.Background()).BookedRoom(bookedRoom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.BookRoomPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BookRoomPost`: BookedRoom
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.BookRoomPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBookRoomPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookedRoom** | [**BookedRoom**](BookedRoom.md) |  | 

### Return type

[**BookedRoom**](BookedRoom.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBookedRooms

> []BookedRoom GetBookedRooms(ctx, hotelId).Execute()

Получить список забронированных комнат по ID отеля

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hotelId := "hotelId_example" // string | ID отеля

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetBookedRooms(context.Background(), hotelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetBookedRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookedRooms`: []BookedRoom
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetBookedRooms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hotelId** | **string** | ID отеля | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookedRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]BookedRoom**](BookedRoom.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnbookedRooms

> []Room GetUnbookedRooms(ctx, hotelId).Execute()

Получить список свободных комнат по ID отеля

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	hotelId := "hotelId_example" // string | ID отеля

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetUnbookedRooms(context.Background(), hotelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetUnbookedRooms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnbookedRooms`: []Room
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetUnbookedRooms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hotelId** | **string** | ID отеля | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnbookedRoomsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Room**](Room.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandlePayment**](DefaultAPI.md#HandlePayment) | **Post** /webhook/payment | Обработка данных о платеже



## HandlePayment

> HandlePayment200Response HandlePayment(ctx).Payment(payment).Execute()

Обработка данных о платеже



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
	payment := *openapiclient.NewPayment(float32(150.75), "101") // Payment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HandlePayment(context.Background()).Payment(payment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HandlePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HandlePayment`: HandlePayment200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HandlePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandlePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payment** | [**Payment**](Payment.md) |  | 

### Return type

[**HandlePayment200Response**](HandlePayment200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


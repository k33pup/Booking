# CreatePaymentPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoomId** | Pointer to **string** | ID комнаты | [optional] 
**Amount** | Pointer to **float32** | Сумма платежа | [optional] 
**WebhookUrl** | Pointer to **string** | URL для получения уведомлений о статусе платежа | [optional] 

## Methods

### NewCreatePaymentPostRequest

`func NewCreatePaymentPostRequest() *CreatePaymentPostRequest`

NewCreatePaymentPostRequest instantiates a new CreatePaymentPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentPostRequestWithDefaults

`func NewCreatePaymentPostRequestWithDefaults() *CreatePaymentPostRequest`

NewCreatePaymentPostRequestWithDefaults instantiates a new CreatePaymentPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoomId

`func (o *CreatePaymentPostRequest) GetRoomId() string`

GetRoomId returns the RoomId field if non-nil, zero value otherwise.

### GetRoomIdOk

`func (o *CreatePaymentPostRequest) GetRoomIdOk() (*string, bool)`

GetRoomIdOk returns a tuple with the RoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomId

`func (o *CreatePaymentPostRequest) SetRoomId(v string)`

SetRoomId sets RoomId field to given value.

### HasRoomId

`func (o *CreatePaymentPostRequest) HasRoomId() bool`

HasRoomId returns a boolean if a field has been set.

### GetAmount

`func (o *CreatePaymentPostRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreatePaymentPostRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreatePaymentPostRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *CreatePaymentPostRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *CreatePaymentPostRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreatePaymentPostRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreatePaymentPostRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreatePaymentPostRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



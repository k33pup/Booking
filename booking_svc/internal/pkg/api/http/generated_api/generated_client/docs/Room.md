# Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | ID комнаты | 
**HotelID** | **string** | ID отеля | 
**Name** | **string** | Название комнаты | 
**Description** | Pointer to **string** | Описание комнаты | [optional] 
**Price** | **int32** | Цена за комнату | 

## Methods

### NewRoom

`func NewRoom(iD string, hotelID string, name string, price int32, ) *Room`

NewRoom instantiates a new Room object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomWithDefaults

`func NewRoomWithDefaults() *Room`

NewRoomWithDefaults instantiates a new Room object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *Room) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Room) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Room) SetID(v string)`

SetID sets ID field to given value.


### GetHotelID

`func (o *Room) GetHotelID() string`

GetHotelID returns the HotelID field if non-nil, zero value otherwise.

### GetHotelIDOk

`func (o *Room) GetHotelIDOk() (*string, bool)`

GetHotelIDOk returns a tuple with the HotelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotelID

`func (o *Room) SetHotelID(v string)`

SetHotelID sets HotelID field to given value.


### GetName

`func (o *Room) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Room) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Room) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Room) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Room) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Room) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Room) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrice

`func (o *Room) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Room) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Room) SetPrice(v int32)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



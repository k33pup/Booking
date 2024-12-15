# BookedRoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | ID забронированной комнаты | 
**HotelID** | **string** | ID отеля | 
**Entry** | **time.Time** | Дата заезда | 
**Exit** | **time.Time** | Дата выезда | 
**Email** | **string** | Email пользователя, забронировавшего комнату | 
**IsPaid** | **bool** | оплачена ли комната | 

## Methods

### NewBookedRoom

`func NewBookedRoom(iD string, hotelID string, entry time.Time, exit time.Time, email string, isPaid bool, ) *BookedRoom`

NewBookedRoom instantiates a new BookedRoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookedRoomWithDefaults

`func NewBookedRoomWithDefaults() *BookedRoom`

NewBookedRoomWithDefaults instantiates a new BookedRoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *BookedRoom) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BookedRoom) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BookedRoom) SetID(v string)`

SetID sets ID field to given value.


### GetHotelID

`func (o *BookedRoom) GetHotelID() string`

GetHotelID returns the HotelID field if non-nil, zero value otherwise.

### GetHotelIDOk

`func (o *BookedRoom) GetHotelIDOk() (*string, bool)`

GetHotelIDOk returns a tuple with the HotelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotelID

`func (o *BookedRoom) SetHotelID(v string)`

SetHotelID sets HotelID field to given value.


### GetEntry

`func (o *BookedRoom) GetEntry() time.Time`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *BookedRoom) GetEntryOk() (*time.Time, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *BookedRoom) SetEntry(v time.Time)`

SetEntry sets Entry field to given value.


### GetExit

`func (o *BookedRoom) GetExit() time.Time`

GetExit returns the Exit field if non-nil, zero value otherwise.

### GetExitOk

`func (o *BookedRoom) GetExitOk() (*time.Time, bool)`

GetExitOk returns a tuple with the Exit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExit

`func (o *BookedRoom) SetExit(v time.Time)`

SetExit sets Exit field to given value.


### GetEmail

`func (o *BookedRoom) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BookedRoom) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BookedRoom) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetIsPaid

`func (o *BookedRoom) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *BookedRoom) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *BookedRoom) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



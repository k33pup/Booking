/*
Hotel Booking API

API для управления бронированием отелей

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Room type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Room{}

// Room struct for Room
type Room struct {
	// ID комнаты
	ID string `json:"ID"`
	// ID отеля
	HotelID string `json:"HotelID"`
	// Название комнаты
	Name string `json:"Name"`
	// Описание комнаты
	Description *string `json:"Description,omitempty"`
	// Цена за комнату
	Price int32 `json:"Price"`
}

type _Room Room

// NewRoom instantiates a new Room object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoom(iD string, hotelID string, name string, price int32) *Room {
	this := Room{}
	this.ID = iD
	this.HotelID = hotelID
	this.Name = name
	this.Price = price
	return &this
}

// NewRoomWithDefaults instantiates a new Room object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomWithDefaults() *Room {
	this := Room{}
	return &this
}

// GetID returns the ID field value
func (o *Room) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *Room) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *Room) SetID(v string) {
	o.ID = v
}

// GetHotelID returns the HotelID field value
func (o *Room) GetHotelID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HotelID
}

// GetHotelIDOk returns a tuple with the HotelID field value
// and a boolean to check if the value has been set.
func (o *Room) GetHotelIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HotelID, true
}

// SetHotelID sets field value
func (o *Room) SetHotelID(v string) {
	o.HotelID = v
}

// GetName returns the Name field value
func (o *Room) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Room) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Room) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Room) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Room) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Room) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Room) SetDescription(v string) {
	o.Description = &v
}

// GetPrice returns the Price field value
func (o *Room) GetPrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *Room) GetPriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *Room) SetPrice(v int32) {
	o.Price = v
}

func (o Room) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Room) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ID"] = o.ID
	toSerialize["HotelID"] = o.HotelID
	toSerialize["Name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["Price"] = o.Price
	return toSerialize, nil
}

func (o *Room) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ID",
		"HotelID",
		"Name",
		"Price",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoom := _Room{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoom)

	if err != nil {
		return err
	}

	*o = Room(varRoom)

	return err
}

type NullableRoom struct {
	value *Room
	isSet bool
}

func (v NullableRoom) Get() *Room {
	return v.value
}

func (v *NullableRoom) Set(val *Room) {
	v.value = val
	v.isSet = true
}

func (v NullableRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoom(val *Room) *NullableRoom {
	return &NullableRoom{value: val, isSet: true}
}

func (v NullableRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



/*
Payment Webhook API

API для обработки платежей по комнатам.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HandlePayment400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlePayment400Response{}

// HandlePayment400Response struct for HandlePayment400Response
type HandlePayment400Response struct {
	Error *string `json:"error,omitempty"`
}

// NewHandlePayment400Response instantiates a new HandlePayment400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlePayment400Response() *HandlePayment400Response {
	this := HandlePayment400Response{}
	return &this
}

// NewHandlePayment400ResponseWithDefaults instantiates a new HandlePayment400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlePayment400ResponseWithDefaults() *HandlePayment400Response {
	this := HandlePayment400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *HandlePayment400Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlePayment400Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HandlePayment400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *HandlePayment400Response) SetError(v string) {
	o.Error = &v
}

func (o HandlePayment400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlePayment400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableHandlePayment400Response struct {
	value *HandlePayment400Response
	isSet bool
}

func (v NullableHandlePayment400Response) Get() *HandlePayment400Response {
	return v.value
}

func (v *NullableHandlePayment400Response) Set(val *HandlePayment400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlePayment400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlePayment400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlePayment400Response(val *HandlePayment400Response) *NullableHandlePayment400Response {
	return &NullableHandlePayment400Response{value: val, isSet: true}
}

func (v NullableHandlePayment400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlePayment400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



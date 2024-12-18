/*
Treasury Management API

A comprehensive API for managing addresses, offering endpoints for creation, update, deletion, and retrieval of address data.

API version: 1.0
Contact: support_cept@indiapost.gov.in
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the HandlerStationeryDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerStationeryDetail{}

// HandlerStationeryDetail struct for HandlerStationeryDetail
type HandlerStationeryDetail struct {
	Name *string `json:"name,omitempty"`
	Quantity *string `json:"quantity,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewHandlerStationeryDetail instantiates a new HandlerStationeryDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerStationeryDetail() *HandlerStationeryDetail {
	this := HandlerStationeryDetail{}
	return &this
}

// NewHandlerStationeryDetailWithDefaults instantiates a new HandlerStationeryDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerStationeryDetailWithDefaults() *HandlerStationeryDetail {
	this := HandlerStationeryDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HandlerStationeryDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStationeryDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HandlerStationeryDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HandlerStationeryDetail) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *HandlerStationeryDetail) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStationeryDetail) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *HandlerStationeryDetail) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *HandlerStationeryDetail) SetQuantity(v string) {
	o.Quantity = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HandlerStationeryDetail) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStationeryDetail) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HandlerStationeryDetail) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HandlerStationeryDetail) SetValue(v string) {
	o.Value = &v
}

func (o HandlerStationeryDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerStationeryDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableHandlerStationeryDetail struct {
	value *HandlerStationeryDetail
	isSet bool
}

func (v NullableHandlerStationeryDetail) Get() *HandlerStationeryDetail {
	return v.value
}

func (v *NullableHandlerStationeryDetail) Set(val *HandlerStationeryDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerStationeryDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerStationeryDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerStationeryDetail(val *HandlerStationeryDetail) *NullableHandlerStationeryDetail {
	return &NullableHandlerStationeryDetail{value: val, isSet: true}
}

func (v NullableHandlerStationeryDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerStationeryDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



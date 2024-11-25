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

// checks if the ApierrorsFieldError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApierrorsFieldError{}

// ApierrorsFieldError struct for ApierrorsFieldError
type ApierrorsFieldError struct {
	Field *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewApierrorsFieldError instantiates a new ApierrorsFieldError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApierrorsFieldError() *ApierrorsFieldError {
	this := ApierrorsFieldError{}
	return &this
}

// NewApierrorsFieldErrorWithDefaults instantiates a new ApierrorsFieldError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApierrorsFieldErrorWithDefaults() *ApierrorsFieldError {
	this := ApierrorsFieldError{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ApierrorsFieldError) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsFieldError) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ApierrorsFieldError) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ApierrorsFieldError) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApierrorsFieldError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsFieldError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApierrorsFieldError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApierrorsFieldError) SetMessage(v string) {
	o.Message = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApierrorsFieldError) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsFieldError) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApierrorsFieldError) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ApierrorsFieldError) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o ApierrorsFieldError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApierrorsFieldError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApierrorsFieldError struct {
	value *ApierrorsFieldError
	isSet bool
}

func (v NullableApierrorsFieldError) Get() *ApierrorsFieldError {
	return v.value
}

func (v *NullableApierrorsFieldError) Set(val *ApierrorsFieldError) {
	v.value = val
	v.isSet = true
}

func (v NullableApierrorsFieldError) IsSet() bool {
	return v.isSet
}

func (v *NullableApierrorsFieldError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApierrorsFieldError(val *ApierrorsFieldError) *NullableApierrorsFieldError {
	return &NullableApierrorsFieldError{value: val, isSet: true}
}

func (v NullableApierrorsFieldError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApierrorsFieldError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ApierrorsAppError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApierrorsAppError{}

// ApierrorsAppError struct for ApierrorsAppError
type ApierrorsAppError struct {
	Code *string `json:"code,omitempty"`
	FieldErrors []ApierrorsFieldError `json:"field_errors,omitempty"`
	Id *string `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewApierrorsAppError instantiates a new ApierrorsAppError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApierrorsAppError() *ApierrorsAppError {
	this := ApierrorsAppError{}
	return &this
}

// NewApierrorsAppErrorWithDefaults instantiates a new ApierrorsAppError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApierrorsAppErrorWithDefaults() *ApierrorsAppError {
	this := ApierrorsAppError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApierrorsAppError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsAppError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApierrorsAppError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ApierrorsAppError) SetCode(v string) {
	o.Code = &v
}

// GetFieldErrors returns the FieldErrors field value if set, zero value otherwise.
func (o *ApierrorsAppError) GetFieldErrors() []ApierrorsFieldError {
	if o == nil || IsNil(o.FieldErrors) {
		var ret []ApierrorsFieldError
		return ret
	}
	return o.FieldErrors
}

// GetFieldErrorsOk returns a tuple with the FieldErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsAppError) GetFieldErrorsOk() ([]ApierrorsFieldError, bool) {
	if o == nil || IsNil(o.FieldErrors) {
		return nil, false
	}
	return o.FieldErrors, true
}

// HasFieldErrors returns a boolean if a field has been set.
func (o *ApierrorsAppError) HasFieldErrors() bool {
	if o != nil && !IsNil(o.FieldErrors) {
		return true
	}

	return false
}

// SetFieldErrors gets a reference to the given []ApierrorsFieldError and assigns it to the FieldErrors field.
func (o *ApierrorsAppError) SetFieldErrors(v []ApierrorsFieldError) {
	o.FieldErrors = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApierrorsAppError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsAppError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApierrorsAppError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApierrorsAppError) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApierrorsAppError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApierrorsAppError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApierrorsAppError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApierrorsAppError) SetMessage(v string) {
	o.Message = &v
}

func (o ApierrorsAppError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApierrorsAppError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.FieldErrors) {
		toSerialize["field_errors"] = o.FieldErrors
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableApierrorsAppError struct {
	value *ApierrorsAppError
	isSet bool
}

func (v NullableApierrorsAppError) Get() *ApierrorsAppError {
	return v.value
}

func (v *NullableApierrorsAppError) Set(val *ApierrorsAppError) {
	v.value = val
	v.isSet = true
}

func (v NullableApierrorsAppError) IsSet() bool {
	return v.isSet
}

func (v *NullableApierrorsAppError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApierrorsAppError(val *ApierrorsAppError) *NullableApierrorsAppError {
	return &NullableApierrorsAppError{value: val, isSet: true}
}

func (v NullableApierrorsAppError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApierrorsAppError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



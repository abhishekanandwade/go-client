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

// checks if the ResponseCreateCashTransferLineLimitsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateCashTransferLineLimitsResponse{}

// ResponseCreateCashTransferLineLimitsResponse struct for ResponseCreateCashTransferLineLimitsResponse
type ResponseCreateCashTransferLineLimitsResponse struct {
	LimitId *string `json:"limit_id,omitempty"`
}

// NewResponseCreateCashTransferLineLimitsResponse instantiates a new ResponseCreateCashTransferLineLimitsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateCashTransferLineLimitsResponse() *ResponseCreateCashTransferLineLimitsResponse {
	this := ResponseCreateCashTransferLineLimitsResponse{}
	return &this
}

// NewResponseCreateCashTransferLineLimitsResponseWithDefaults instantiates a new ResponseCreateCashTransferLineLimitsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateCashTransferLineLimitsResponseWithDefaults() *ResponseCreateCashTransferLineLimitsResponse {
	this := ResponseCreateCashTransferLineLimitsResponse{}
	return &this
}

// GetLimitId returns the LimitId field value if set, zero value otherwise.
func (o *ResponseCreateCashTransferLineLimitsResponse) GetLimitId() string {
	if o == nil || IsNil(o.LimitId) {
		var ret string
		return ret
	}
	return *o.LimitId
}

// GetLimitIdOk returns a tuple with the LimitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateCashTransferLineLimitsResponse) GetLimitIdOk() (*string, bool) {
	if o == nil || IsNil(o.LimitId) {
		return nil, false
	}
	return o.LimitId, true
}

// HasLimitId returns a boolean if a field has been set.
func (o *ResponseCreateCashTransferLineLimitsResponse) HasLimitId() bool {
	if o != nil && !IsNil(o.LimitId) {
		return true
	}

	return false
}

// SetLimitId gets a reference to the given string and assigns it to the LimitId field.
func (o *ResponseCreateCashTransferLineLimitsResponse) SetLimitId(v string) {
	o.LimitId = &v
}

func (o ResponseCreateCashTransferLineLimitsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateCashTransferLineLimitsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LimitId) {
		toSerialize["limit_id"] = o.LimitId
	}
	return toSerialize, nil
}

type NullableResponseCreateCashTransferLineLimitsResponse struct {
	value *ResponseCreateCashTransferLineLimitsResponse
	isSet bool
}

func (v NullableResponseCreateCashTransferLineLimitsResponse) Get() *ResponseCreateCashTransferLineLimitsResponse {
	return v.value
}

func (v *NullableResponseCreateCashTransferLineLimitsResponse) Set(val *ResponseCreateCashTransferLineLimitsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateCashTransferLineLimitsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateCashTransferLineLimitsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateCashTransferLineLimitsResponse(val *ResponseCreateCashTransferLineLimitsResponse) *NullableResponseCreateCashTransferLineLimitsResponse {
	return &NullableResponseCreateCashTransferLineLimitsResponse{value: val, isSet: true}
}

func (v NullableResponseCreateCashTransferLineLimitsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateCashTransferLineLimitsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



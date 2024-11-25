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

// checks if the ResponseCreateStampObReceiptsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateStampObReceiptsResponse{}

// ResponseCreateStampObReceiptsResponse struct for ResponseCreateStampObReceiptsResponse
type ResponseCreateStampObReceiptsResponse struct {
	ObTransactionId *string `json:"ob_transaction_id,omitempty"`
}

// NewResponseCreateStampObReceiptsResponse instantiates a new ResponseCreateStampObReceiptsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateStampObReceiptsResponse() *ResponseCreateStampObReceiptsResponse {
	this := ResponseCreateStampObReceiptsResponse{}
	return &this
}

// NewResponseCreateStampObReceiptsResponseWithDefaults instantiates a new ResponseCreateStampObReceiptsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateStampObReceiptsResponseWithDefaults() *ResponseCreateStampObReceiptsResponse {
	this := ResponseCreateStampObReceiptsResponse{}
	return &this
}

// GetObTransactionId returns the ObTransactionId field value if set, zero value otherwise.
func (o *ResponseCreateStampObReceiptsResponse) GetObTransactionId() string {
	if o == nil || IsNil(o.ObTransactionId) {
		var ret string
		return ret
	}
	return *o.ObTransactionId
}

// GetObTransactionIdOk returns a tuple with the ObTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateStampObReceiptsResponse) GetObTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObTransactionId) {
		return nil, false
	}
	return o.ObTransactionId, true
}

// HasObTransactionId returns a boolean if a field has been set.
func (o *ResponseCreateStampObReceiptsResponse) HasObTransactionId() bool {
	if o != nil && !IsNil(o.ObTransactionId) {
		return true
	}

	return false
}

// SetObTransactionId gets a reference to the given string and assigns it to the ObTransactionId field.
func (o *ResponseCreateStampObReceiptsResponse) SetObTransactionId(v string) {
	o.ObTransactionId = &v
}

func (o ResponseCreateStampObReceiptsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateStampObReceiptsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ObTransactionId) {
		toSerialize["ob_transaction_id"] = o.ObTransactionId
	}
	return toSerialize, nil
}

type NullableResponseCreateStampObReceiptsResponse struct {
	value *ResponseCreateStampObReceiptsResponse
	isSet bool
}

func (v NullableResponseCreateStampObReceiptsResponse) Get() *ResponseCreateStampObReceiptsResponse {
	return v.value
}

func (v *NullableResponseCreateStampObReceiptsResponse) Set(val *ResponseCreateStampObReceiptsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateStampObReceiptsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateStampObReceiptsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateStampObReceiptsResponse(val *ResponseCreateStampObReceiptsResponse) *NullableResponseCreateStampObReceiptsResponse {
	return &NullableResponseCreateStampObReceiptsResponse{value: val, isSet: true}
}

func (v NullableResponseCreateStampObReceiptsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateStampObReceiptsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


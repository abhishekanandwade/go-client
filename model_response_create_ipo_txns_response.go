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

// checks if the ResponseCreateIpoTxnsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateIpoTxnsResponse{}

// ResponseCreateIpoTxnsResponse struct for ResponseCreateIpoTxnsResponse
type ResponseCreateIpoTxnsResponse struct {
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewResponseCreateIpoTxnsResponse instantiates a new ResponseCreateIpoTxnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateIpoTxnsResponse() *ResponseCreateIpoTxnsResponse {
	this := ResponseCreateIpoTxnsResponse{}
	return &this
}

// NewResponseCreateIpoTxnsResponseWithDefaults instantiates a new ResponseCreateIpoTxnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateIpoTxnsResponseWithDefaults() *ResponseCreateIpoTxnsResponse {
	this := ResponseCreateIpoTxnsResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseCreateIpoTxnsResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateIpoTxnsResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseCreateIpoTxnsResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseCreateIpoTxnsResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o ResponseCreateIpoTxnsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateIpoTxnsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableResponseCreateIpoTxnsResponse struct {
	value *ResponseCreateIpoTxnsResponse
	isSet bool
}

func (v NullableResponseCreateIpoTxnsResponse) Get() *ResponseCreateIpoTxnsResponse {
	return v.value
}

func (v *NullableResponseCreateIpoTxnsResponse) Set(val *ResponseCreateIpoTxnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateIpoTxnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateIpoTxnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateIpoTxnsResponse(val *ResponseCreateIpoTxnsResponse) *NullableResponseCreateIpoTxnsResponse {
	return &NullableResponseCreateIpoTxnsResponse{value: val, isSet: true}
}

func (v NullableResponseCreateIpoTxnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateIpoTxnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



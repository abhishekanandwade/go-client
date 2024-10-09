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

// checks if the ResponseCreateTreasuryTxnsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateTreasuryTxnsResponse{}

// ResponseCreateTreasuryTxnsResponse struct for ResponseCreateTreasuryTxnsResponse
type ResponseCreateTreasuryTxnsResponse struct {
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewResponseCreateTreasuryTxnsResponse instantiates a new ResponseCreateTreasuryTxnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateTreasuryTxnsResponse() *ResponseCreateTreasuryTxnsResponse {
	this := ResponseCreateTreasuryTxnsResponse{}
	return &this
}

// NewResponseCreateTreasuryTxnsResponseWithDefaults instantiates a new ResponseCreateTreasuryTxnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateTreasuryTxnsResponseWithDefaults() *ResponseCreateTreasuryTxnsResponse {
	this := ResponseCreateTreasuryTxnsResponse{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseCreateTreasuryTxnsResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateTreasuryTxnsResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseCreateTreasuryTxnsResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseCreateTreasuryTxnsResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o ResponseCreateTreasuryTxnsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateTreasuryTxnsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableResponseCreateTreasuryTxnsResponse struct {
	value *ResponseCreateTreasuryTxnsResponse
	isSet bool
}

func (v NullableResponseCreateTreasuryTxnsResponse) Get() *ResponseCreateTreasuryTxnsResponse {
	return v.value
}

func (v *NullableResponseCreateTreasuryTxnsResponse) Set(val *ResponseCreateTreasuryTxnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateTreasuryTxnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateTreasuryTxnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateTreasuryTxnsResponse(val *ResponseCreateTreasuryTxnsResponse) *NullableResponseCreateTreasuryTxnsResponse {
	return &NullableResponseCreateTreasuryTxnsResponse{value: val, isSet: true}
}

func (v NullableResponseCreateTreasuryTxnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateTreasuryTxnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


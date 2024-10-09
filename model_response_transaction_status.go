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

// checks if the ResponseTransactionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTransactionStatus{}

// ResponseTransactionStatus struct for ResponseTransactionStatus
type ResponseTransactionStatus struct {
	Status *string `json:"status,omitempty"`
	TransactionCount *int32 `json:"transaction_count,omitempty"`
	TransactionType *string `json:"transaction_type,omitempty"`
}

// NewResponseTransactionStatus instantiates a new ResponseTransactionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTransactionStatus() *ResponseTransactionStatus {
	this := ResponseTransactionStatus{}
	return &this
}

// NewResponseTransactionStatusWithDefaults instantiates a new ResponseTransactionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTransactionStatusWithDefaults() *ResponseTransactionStatus {
	this := ResponseTransactionStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseTransactionStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseTransactionStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseTransactionStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTransactionCount returns the TransactionCount field value if set, zero value otherwise.
func (o *ResponseTransactionStatus) GetTransactionCount() int32 {
	if o == nil || IsNil(o.TransactionCount) {
		var ret int32
		return ret
	}
	return *o.TransactionCount
}

// GetTransactionCountOk returns a tuple with the TransactionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionStatus) GetTransactionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransactionCount) {
		return nil, false
	}
	return o.TransactionCount, true
}

// HasTransactionCount returns a boolean if a field has been set.
func (o *ResponseTransactionStatus) HasTransactionCount() bool {
	if o != nil && !IsNil(o.TransactionCount) {
		return true
	}

	return false
}

// SetTransactionCount gets a reference to the given int32 and assigns it to the TransactionCount field.
func (o *ResponseTransactionStatus) SetTransactionCount(v int32) {
	o.TransactionCount = &v
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *ResponseTransactionStatus) GetTransactionType() string {
	if o == nil || IsNil(o.TransactionType) {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransactionStatus) GetTransactionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionType) {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *ResponseTransactionStatus) HasTransactionType() bool {
	if o != nil && !IsNil(o.TransactionType) {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *ResponseTransactionStatus) SetTransactionType(v string) {
	o.TransactionType = &v
}

func (o ResponseTransactionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTransactionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransactionCount) {
		toSerialize["transaction_count"] = o.TransactionCount
	}
	if !IsNil(o.TransactionType) {
		toSerialize["transaction_type"] = o.TransactionType
	}
	return toSerialize, nil
}

type NullableResponseTransactionStatus struct {
	value *ResponseTransactionStatus
	isSet bool
}

func (v NullableResponseTransactionStatus) Get() *ResponseTransactionStatus {
	return v.value
}

func (v *NullableResponseTransactionStatus) Set(val *ResponseTransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTransactionStatus(val *ResponseTransactionStatus) *NullableResponseTransactionStatus {
	return &NullableResponseTransactionStatus{value: val, isSet: true}
}

func (v NullableResponseTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


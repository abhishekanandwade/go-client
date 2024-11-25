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

// checks if the ResponseCreateIPOSales type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateIPOSales{}

// ResponseCreateIPOSales struct for ResponseCreateIPOSales
type ResponseCreateIPOSales struct {
	TransactionId *string `json:"transaction_id,omitempty"`
}

// NewResponseCreateIPOSales instantiates a new ResponseCreateIPOSales object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateIPOSales() *ResponseCreateIPOSales {
	this := ResponseCreateIPOSales{}
	return &this
}

// NewResponseCreateIPOSalesWithDefaults instantiates a new ResponseCreateIPOSales object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateIPOSalesWithDefaults() *ResponseCreateIPOSales {
	this := ResponseCreateIPOSales{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseCreateIPOSales) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateIPOSales) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseCreateIPOSales) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseCreateIPOSales) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o ResponseCreateIPOSales) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateIPOSales) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableResponseCreateIPOSales struct {
	value *ResponseCreateIPOSales
	isSet bool
}

func (v NullableResponseCreateIPOSales) Get() *ResponseCreateIPOSales {
	return v.value
}

func (v *NullableResponseCreateIPOSales) Set(val *ResponseCreateIPOSales) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateIPOSales) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateIPOSales) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateIPOSales(val *ResponseCreateIPOSales) *NullableResponseCreateIPOSales {
	return &NullableResponseCreateIPOSales{value: val, isSet: true}
}

func (v NullableResponseCreateIPOSales) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateIPOSales) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



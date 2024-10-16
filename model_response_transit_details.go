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

// checks if the ResponseTransitDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTransitDetails{}

// ResponseTransitDetails struct for ResponseTransitDetails
type ResponseTransitDetails struct {
	TransAmount *float32 `json:"trans_amount,omitempty"`
	TransCount *int32 `json:"trans_count,omitempty"`
	TransType *string `json:"trans_type,omitempty"`
}

// NewResponseTransitDetails instantiates a new ResponseTransitDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTransitDetails() *ResponseTransitDetails {
	this := ResponseTransitDetails{}
	return &this
}

// NewResponseTransitDetailsWithDefaults instantiates a new ResponseTransitDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTransitDetailsWithDefaults() *ResponseTransitDetails {
	this := ResponseTransitDetails{}
	return &this
}

// GetTransAmount returns the TransAmount field value if set, zero value otherwise.
func (o *ResponseTransitDetails) GetTransAmount() float32 {
	if o == nil || IsNil(o.TransAmount) {
		var ret float32
		return ret
	}
	return *o.TransAmount
}

// GetTransAmountOk returns a tuple with the TransAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransitDetails) GetTransAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TransAmount) {
		return nil, false
	}
	return o.TransAmount, true
}

// HasTransAmount returns a boolean if a field has been set.
func (o *ResponseTransitDetails) HasTransAmount() bool {
	if o != nil && !IsNil(o.TransAmount) {
		return true
	}

	return false
}

// SetTransAmount gets a reference to the given float32 and assigns it to the TransAmount field.
func (o *ResponseTransitDetails) SetTransAmount(v float32) {
	o.TransAmount = &v
}

// GetTransCount returns the TransCount field value if set, zero value otherwise.
func (o *ResponseTransitDetails) GetTransCount() int32 {
	if o == nil || IsNil(o.TransCount) {
		var ret int32
		return ret
	}
	return *o.TransCount
}

// GetTransCountOk returns a tuple with the TransCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransitDetails) GetTransCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TransCount) {
		return nil, false
	}
	return o.TransCount, true
}

// HasTransCount returns a boolean if a field has been set.
func (o *ResponseTransitDetails) HasTransCount() bool {
	if o != nil && !IsNil(o.TransCount) {
		return true
	}

	return false
}

// SetTransCount gets a reference to the given int32 and assigns it to the TransCount field.
func (o *ResponseTransitDetails) SetTransCount(v int32) {
	o.TransCount = &v
}

// GetTransType returns the TransType field value if set, zero value otherwise.
func (o *ResponseTransitDetails) GetTransType() string {
	if o == nil || IsNil(o.TransType) {
		var ret string
		return ret
	}
	return *o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTransitDetails) GetTransTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransType) {
		return nil, false
	}
	return o.TransType, true
}

// HasTransType returns a boolean if a field has been set.
func (o *ResponseTransitDetails) HasTransType() bool {
	if o != nil && !IsNil(o.TransType) {
		return true
	}

	return false
}

// SetTransType gets a reference to the given string and assigns it to the TransType field.
func (o *ResponseTransitDetails) SetTransType(v string) {
	o.TransType = &v
}

func (o ResponseTransitDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTransitDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransAmount) {
		toSerialize["trans_amount"] = o.TransAmount
	}
	if !IsNil(o.TransCount) {
		toSerialize["trans_count"] = o.TransCount
	}
	if !IsNil(o.TransType) {
		toSerialize["trans_type"] = o.TransType
	}
	return toSerialize, nil
}

type NullableResponseTransitDetails struct {
	value *ResponseTransitDetails
	isSet bool
}

func (v NullableResponseTransitDetails) Get() *ResponseTransitDetails {
	return v.value
}

func (v *NullableResponseTransitDetails) Set(val *ResponseTransitDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTransitDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTransitDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTransitDetails(val *ResponseTransitDetails) *NullableResponseTransitDetails {
	return &NullableResponseTransitDetails{value: val, isSet: true}
}

func (v NullableResponseTransitDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTransitDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



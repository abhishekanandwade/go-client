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

// checks if the HandlerTemporalBagOpenGetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerTemporalBagOpenGetRequest{}

// HandlerTemporalBagOpenGetRequest struct for HandlerTemporalBagOpenGetRequest
type HandlerTemporalBagOpenGetRequest struct {
	Bagnumber *string `json:"bagnumber,omitempty"`
	Officeid *int32 `json:"officeid,omitempty"`
}

// NewHandlerTemporalBagOpenGetRequest instantiates a new HandlerTemporalBagOpenGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerTemporalBagOpenGetRequest() *HandlerTemporalBagOpenGetRequest {
	this := HandlerTemporalBagOpenGetRequest{}
	return &this
}

// NewHandlerTemporalBagOpenGetRequestWithDefaults instantiates a new HandlerTemporalBagOpenGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerTemporalBagOpenGetRequestWithDefaults() *HandlerTemporalBagOpenGetRequest {
	this := HandlerTemporalBagOpenGetRequest{}
	return &this
}

// GetBagnumber returns the Bagnumber field value if set, zero value otherwise.
func (o *HandlerTemporalBagOpenGetRequest) GetBagnumber() string {
	if o == nil || IsNil(o.Bagnumber) {
		var ret string
		return ret
	}
	return *o.Bagnumber
}

// GetBagnumberOk returns a tuple with the Bagnumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalBagOpenGetRequest) GetBagnumberOk() (*string, bool) {
	if o == nil || IsNil(o.Bagnumber) {
		return nil, false
	}
	return o.Bagnumber, true
}

// HasBagnumber returns a boolean if a field has been set.
func (o *HandlerTemporalBagOpenGetRequest) HasBagnumber() bool {
	if o != nil && !IsNil(o.Bagnumber) {
		return true
	}

	return false
}

// SetBagnumber gets a reference to the given string and assigns it to the Bagnumber field.
func (o *HandlerTemporalBagOpenGetRequest) SetBagnumber(v string) {
	o.Bagnumber = &v
}

// GetOfficeid returns the Officeid field value if set, zero value otherwise.
func (o *HandlerTemporalBagOpenGetRequest) GetOfficeid() int32 {
	if o == nil || IsNil(o.Officeid) {
		var ret int32
		return ret
	}
	return *o.Officeid
}

// GetOfficeidOk returns a tuple with the Officeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalBagOpenGetRequest) GetOfficeidOk() (*int32, bool) {
	if o == nil || IsNil(o.Officeid) {
		return nil, false
	}
	return o.Officeid, true
}

// HasOfficeid returns a boolean if a field has been set.
func (o *HandlerTemporalBagOpenGetRequest) HasOfficeid() bool {
	if o != nil && !IsNil(o.Officeid) {
		return true
	}

	return false
}

// SetOfficeid gets a reference to the given int32 and assigns it to the Officeid field.
func (o *HandlerTemporalBagOpenGetRequest) SetOfficeid(v int32) {
	o.Officeid = &v
}

func (o HandlerTemporalBagOpenGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerTemporalBagOpenGetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bagnumber) {
		toSerialize["bagnumber"] = o.Bagnumber
	}
	if !IsNil(o.Officeid) {
		toSerialize["officeid"] = o.Officeid
	}
	return toSerialize, nil
}

type NullableHandlerTemporalBagOpenGetRequest struct {
	value *HandlerTemporalBagOpenGetRequest
	isSet bool
}

func (v NullableHandlerTemporalBagOpenGetRequest) Get() *HandlerTemporalBagOpenGetRequest {
	return v.value
}

func (v *NullableHandlerTemporalBagOpenGetRequest) Set(val *HandlerTemporalBagOpenGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerTemporalBagOpenGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerTemporalBagOpenGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerTemporalBagOpenGetRequest(val *HandlerTemporalBagOpenGetRequest) *NullableHandlerTemporalBagOpenGetRequest {
	return &NullableHandlerTemporalBagOpenGetRequest{value: val, isSet: true}
}

func (v NullableHandlerTemporalBagOpenGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerTemporalBagOpenGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the HandlerEmoDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerEmoDetail{}

// HandlerEmoDetail struct for HandlerEmoDetail
type HandlerEmoDetail struct {
	Addresse *string `json:"addresse,omitempty"`
	EmoNo *string `json:"emo_no,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewHandlerEmoDetail instantiates a new HandlerEmoDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerEmoDetail() *HandlerEmoDetail {
	this := HandlerEmoDetail{}
	return &this
}

// NewHandlerEmoDetailWithDefaults instantiates a new HandlerEmoDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerEmoDetailWithDefaults() *HandlerEmoDetail {
	this := HandlerEmoDetail{}
	return &this
}

// GetAddresse returns the Addresse field value if set, zero value otherwise.
func (o *HandlerEmoDetail) GetAddresse() string {
	if o == nil || IsNil(o.Addresse) {
		var ret string
		return ret
	}
	return *o.Addresse
}

// GetAddresseOk returns a tuple with the Addresse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerEmoDetail) GetAddresseOk() (*string, bool) {
	if o == nil || IsNil(o.Addresse) {
		return nil, false
	}
	return o.Addresse, true
}

// HasAddresse returns a boolean if a field has been set.
func (o *HandlerEmoDetail) HasAddresse() bool {
	if o != nil && !IsNil(o.Addresse) {
		return true
	}

	return false
}

// SetAddresse gets a reference to the given string and assigns it to the Addresse field.
func (o *HandlerEmoDetail) SetAddresse(v string) {
	o.Addresse = &v
}

// GetEmoNo returns the EmoNo field value if set, zero value otherwise.
func (o *HandlerEmoDetail) GetEmoNo() string {
	if o == nil || IsNil(o.EmoNo) {
		var ret string
		return ret
	}
	return *o.EmoNo
}

// GetEmoNoOk returns a tuple with the EmoNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerEmoDetail) GetEmoNoOk() (*string, bool) {
	if o == nil || IsNil(o.EmoNo) {
		return nil, false
	}
	return o.EmoNo, true
}

// HasEmoNo returns a boolean if a field has been set.
func (o *HandlerEmoDetail) HasEmoNo() bool {
	if o != nil && !IsNil(o.EmoNo) {
		return true
	}

	return false
}

// SetEmoNo gets a reference to the given string and assigns it to the EmoNo field.
func (o *HandlerEmoDetail) SetEmoNo(v string) {
	o.EmoNo = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *HandlerEmoDetail) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerEmoDetail) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *HandlerEmoDetail) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *HandlerEmoDetail) SetValue(v string) {
	o.Value = &v
}

func (o HandlerEmoDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerEmoDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresse) {
		toSerialize["addresse"] = o.Addresse
	}
	if !IsNil(o.EmoNo) {
		toSerialize["emo_no"] = o.EmoNo
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableHandlerEmoDetail struct {
	value *HandlerEmoDetail
	isSet bool
}

func (v NullableHandlerEmoDetail) Get() *HandlerEmoDetail {
	return v.value
}

func (v *NullableHandlerEmoDetail) Set(val *HandlerEmoDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerEmoDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerEmoDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerEmoDetail(val *HandlerEmoDetail) *NullableHandlerEmoDetail {
	return &NullableHandlerEmoDetail{value: val, isSet: true}
}

func (v NullableHandlerEmoDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerEmoDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



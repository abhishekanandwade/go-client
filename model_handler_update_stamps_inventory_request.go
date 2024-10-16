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
	"bytes"
	"fmt"
)

// checks if the HandlerUpdateStampsInventoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerUpdateStampsInventoryRequest{}

// HandlerUpdateStampsInventoryRequest struct for HandlerUpdateStampsInventoryRequest
type HandlerUpdateStampsInventoryRequest struct {
	SubRequestId int32 `json:"sub_request_id"`
	UserId int32 `json:"user_id"`
}

type _HandlerUpdateStampsInventoryRequest HandlerUpdateStampsInventoryRequest

// NewHandlerUpdateStampsInventoryRequest instantiates a new HandlerUpdateStampsInventoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerUpdateStampsInventoryRequest(subRequestId int32, userId int32) *HandlerUpdateStampsInventoryRequest {
	this := HandlerUpdateStampsInventoryRequest{}
	this.SubRequestId = subRequestId
	this.UserId = userId
	return &this
}

// NewHandlerUpdateStampsInventoryRequestWithDefaults instantiates a new HandlerUpdateStampsInventoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerUpdateStampsInventoryRequestWithDefaults() *HandlerUpdateStampsInventoryRequest {
	this := HandlerUpdateStampsInventoryRequest{}
	return &this
}

// GetSubRequestId returns the SubRequestId field value
func (o *HandlerUpdateStampsInventoryRequest) GetSubRequestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SubRequestId
}

// GetSubRequestIdOk returns a tuple with the SubRequestId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampsInventoryRequest) GetSubRequestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubRequestId, true
}

// SetSubRequestId sets field value
func (o *HandlerUpdateStampsInventoryRequest) SetSubRequestId(v int32) {
	o.SubRequestId = v
}

// GetUserId returns the UserId field value
func (o *HandlerUpdateStampsInventoryRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampsInventoryRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *HandlerUpdateStampsInventoryRequest) SetUserId(v int32) {
	o.UserId = v
}

func (o HandlerUpdateStampsInventoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerUpdateStampsInventoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sub_request_id"] = o.SubRequestId
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *HandlerUpdateStampsInventoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sub_request_id",
		"user_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHandlerUpdateStampsInventoryRequest := _HandlerUpdateStampsInventoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerUpdateStampsInventoryRequest)

	if err != nil {
		return err
	}

	*o = HandlerUpdateStampsInventoryRequest(varHandlerUpdateStampsInventoryRequest)

	return err
}

type NullableHandlerUpdateStampsInventoryRequest struct {
	value *HandlerUpdateStampsInventoryRequest
	isSet bool
}

func (v NullableHandlerUpdateStampsInventoryRequest) Get() *HandlerUpdateStampsInventoryRequest {
	return v.value
}

func (v *NullableHandlerUpdateStampsInventoryRequest) Set(val *HandlerUpdateStampsInventoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerUpdateStampsInventoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerUpdateStampsInventoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerUpdateStampsInventoryRequest(val *HandlerUpdateStampsInventoryRequest) *NullableHandlerUpdateStampsInventoryRequest {
	return &NullableHandlerUpdateStampsInventoryRequest{value: val, isSet: true}
}

func (v NullableHandlerUpdateStampsInventoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerUpdateStampsInventoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



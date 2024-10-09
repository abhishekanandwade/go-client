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

// checks if the HandlerUpdateStampAdvanceTxnRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerUpdateStampAdvanceTxnRequest{}

// HandlerUpdateStampAdvanceTxnRequest struct for HandlerUpdateStampAdvanceTxnRequest
type HandlerUpdateStampAdvanceTxnRequest struct {
	IsApproved *bool `json:"is_approved,omitempty"`
	TryApproverId int32 `json:"try_approver_id"`
}

type _HandlerUpdateStampAdvanceTxnRequest HandlerUpdateStampAdvanceTxnRequest

// NewHandlerUpdateStampAdvanceTxnRequest instantiates a new HandlerUpdateStampAdvanceTxnRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerUpdateStampAdvanceTxnRequest(tryApproverId int32) *HandlerUpdateStampAdvanceTxnRequest {
	this := HandlerUpdateStampAdvanceTxnRequest{}
	this.TryApproverId = tryApproverId
	return &this
}

// NewHandlerUpdateStampAdvanceTxnRequestWithDefaults instantiates a new HandlerUpdateStampAdvanceTxnRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerUpdateStampAdvanceTxnRequestWithDefaults() *HandlerUpdateStampAdvanceTxnRequest {
	this := HandlerUpdateStampAdvanceTxnRequest{}
	return &this
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *HandlerUpdateStampAdvanceTxnRequest) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampAdvanceTxnRequest) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *HandlerUpdateStampAdvanceTxnRequest) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *HandlerUpdateStampAdvanceTxnRequest) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetTryApproverId returns the TryApproverId field value
func (o *HandlerUpdateStampAdvanceTxnRequest) GetTryApproverId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TryApproverId
}

// GetTryApproverIdOk returns a tuple with the TryApproverId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampAdvanceTxnRequest) GetTryApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TryApproverId, true
}

// SetTryApproverId sets field value
func (o *HandlerUpdateStampAdvanceTxnRequest) SetTryApproverId(v int32) {
	o.TryApproverId = v
}

func (o HandlerUpdateStampAdvanceTxnRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerUpdateStampAdvanceTxnRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	toSerialize["try_approver_id"] = o.TryApproverId
	return toSerialize, nil
}

func (o *HandlerUpdateStampAdvanceTxnRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"try_approver_id",
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

	varHandlerUpdateStampAdvanceTxnRequest := _HandlerUpdateStampAdvanceTxnRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerUpdateStampAdvanceTxnRequest)

	if err != nil {
		return err
	}

	*o = HandlerUpdateStampAdvanceTxnRequest(varHandlerUpdateStampAdvanceTxnRequest)

	return err
}

type NullableHandlerUpdateStampAdvanceTxnRequest struct {
	value *HandlerUpdateStampAdvanceTxnRequest
	isSet bool
}

func (v NullableHandlerUpdateStampAdvanceTxnRequest) Get() *HandlerUpdateStampAdvanceTxnRequest {
	return v.value
}

func (v *NullableHandlerUpdateStampAdvanceTxnRequest) Set(val *HandlerUpdateStampAdvanceTxnRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerUpdateStampAdvanceTxnRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerUpdateStampAdvanceTxnRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerUpdateStampAdvanceTxnRequest(val *HandlerUpdateStampAdvanceTxnRequest) *NullableHandlerUpdateStampAdvanceTxnRequest {
	return &NullableHandlerUpdateStampAdvanceTxnRequest{value: val, isSet: true}
}

func (v NullableHandlerUpdateStampAdvanceTxnRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerUpdateStampAdvanceTxnRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



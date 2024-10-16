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

// checks if the HandlerUpdateStampBulkSalesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerUpdateStampBulkSalesRequest{}

// HandlerUpdateStampBulkSalesRequest struct for HandlerUpdateStampBulkSalesRequest
type HandlerUpdateStampBulkSalesRequest struct {
	ApproverId int32 `json:"approver_id"`
	IsApproved *bool `json:"is_approved,omitempty"`
}

type _HandlerUpdateStampBulkSalesRequest HandlerUpdateStampBulkSalesRequest

// NewHandlerUpdateStampBulkSalesRequest instantiates a new HandlerUpdateStampBulkSalesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerUpdateStampBulkSalesRequest(approverId int32) *HandlerUpdateStampBulkSalesRequest {
	this := HandlerUpdateStampBulkSalesRequest{}
	this.ApproverId = approverId
	return &this
}

// NewHandlerUpdateStampBulkSalesRequestWithDefaults instantiates a new HandlerUpdateStampBulkSalesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerUpdateStampBulkSalesRequestWithDefaults() *HandlerUpdateStampBulkSalesRequest {
	this := HandlerUpdateStampBulkSalesRequest{}
	return &this
}

// GetApproverId returns the ApproverId field value
func (o *HandlerUpdateStampBulkSalesRequest) GetApproverId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampBulkSalesRequest) GetApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *HandlerUpdateStampBulkSalesRequest) SetApproverId(v int32) {
	o.ApproverId = v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *HandlerUpdateStampBulkSalesRequest) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerUpdateStampBulkSalesRequest) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *HandlerUpdateStampBulkSalesRequest) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *HandlerUpdateStampBulkSalesRequest) SetIsApproved(v bool) {
	o.IsApproved = &v
}

func (o HandlerUpdateStampBulkSalesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerUpdateStampBulkSalesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approver_id"] = o.ApproverId
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	return toSerialize, nil
}

func (o *HandlerUpdateStampBulkSalesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approver_id",
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

	varHandlerUpdateStampBulkSalesRequest := _HandlerUpdateStampBulkSalesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerUpdateStampBulkSalesRequest)

	if err != nil {
		return err
	}

	*o = HandlerUpdateStampBulkSalesRequest(varHandlerUpdateStampBulkSalesRequest)

	return err
}

type NullableHandlerUpdateStampBulkSalesRequest struct {
	value *HandlerUpdateStampBulkSalesRequest
	isSet bool
}

func (v NullableHandlerUpdateStampBulkSalesRequest) Get() *HandlerUpdateStampBulkSalesRequest {
	return v.value
}

func (v *NullableHandlerUpdateStampBulkSalesRequest) Set(val *HandlerUpdateStampBulkSalesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerUpdateStampBulkSalesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerUpdateStampBulkSalesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerUpdateStampBulkSalesRequest(val *HandlerUpdateStampBulkSalesRequest) *NullableHandlerUpdateStampBulkSalesRequest {
	return &NullableHandlerUpdateStampBulkSalesRequest{value: val, isSet: true}
}

func (v NullableHandlerUpdateStampBulkSalesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerUpdateStampBulkSalesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



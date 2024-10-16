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

// checks if the HandlerApproveCustomerTransactionsApproveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerApproveCustomerTransactionsApproveRequest{}

// HandlerApproveCustomerTransactionsApproveRequest struct for HandlerApproveCustomerTransactionsApproveRequest
type HandlerApproveCustomerTransactionsApproveRequest struct {
	ApproverId string `json:"approver-id"`
	IsApproved *bool `json:"is-approved,omitempty"`
}

type _HandlerApproveCustomerTransactionsApproveRequest HandlerApproveCustomerTransactionsApproveRequest

// NewHandlerApproveCustomerTransactionsApproveRequest instantiates a new HandlerApproveCustomerTransactionsApproveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerApproveCustomerTransactionsApproveRequest(approverId string) *HandlerApproveCustomerTransactionsApproveRequest {
	this := HandlerApproveCustomerTransactionsApproveRequest{}
	this.ApproverId = approverId
	return &this
}

// NewHandlerApproveCustomerTransactionsApproveRequestWithDefaults instantiates a new HandlerApproveCustomerTransactionsApproveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerApproveCustomerTransactionsApproveRequestWithDefaults() *HandlerApproveCustomerTransactionsApproveRequest {
	this := HandlerApproveCustomerTransactionsApproveRequest{}
	return &this
}

// GetApproverId returns the ApproverId field value
func (o *HandlerApproveCustomerTransactionsApproveRequest) GetApproverId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *HandlerApproveCustomerTransactionsApproveRequest) GetApproverIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *HandlerApproveCustomerTransactionsApproveRequest) SetApproverId(v string) {
	o.ApproverId = v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *HandlerApproveCustomerTransactionsApproveRequest) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerApproveCustomerTransactionsApproveRequest) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *HandlerApproveCustomerTransactionsApproveRequest) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *HandlerApproveCustomerTransactionsApproveRequest) SetIsApproved(v bool) {
	o.IsApproved = &v
}

func (o HandlerApproveCustomerTransactionsApproveRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerApproveCustomerTransactionsApproveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approver-id"] = o.ApproverId
	if !IsNil(o.IsApproved) {
		toSerialize["is-approved"] = o.IsApproved
	}
	return toSerialize, nil
}

func (o *HandlerApproveCustomerTransactionsApproveRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approver-id",
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

	varHandlerApproveCustomerTransactionsApproveRequest := _HandlerApproveCustomerTransactionsApproveRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerApproveCustomerTransactionsApproveRequest)

	if err != nil {
		return err
	}

	*o = HandlerApproveCustomerTransactionsApproveRequest(varHandlerApproveCustomerTransactionsApproveRequest)

	return err
}

type NullableHandlerApproveCustomerTransactionsApproveRequest struct {
	value *HandlerApproveCustomerTransactionsApproveRequest
	isSet bool
}

func (v NullableHandlerApproveCustomerTransactionsApproveRequest) Get() *HandlerApproveCustomerTransactionsApproveRequest {
	return v.value
}

func (v *NullableHandlerApproveCustomerTransactionsApproveRequest) Set(val *HandlerApproveCustomerTransactionsApproveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerApproveCustomerTransactionsApproveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerApproveCustomerTransactionsApproveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerApproveCustomerTransactionsApproveRequest(val *HandlerApproveCustomerTransactionsApproveRequest) *NullableHandlerApproveCustomerTransactionsApproveRequest {
	return &NullableHandlerApproveCustomerTransactionsApproveRequest{value: val, isSet: true}
}

func (v NullableHandlerApproveCustomerTransactionsApproveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerApproveCustomerTransactionsApproveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



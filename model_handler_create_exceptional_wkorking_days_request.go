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

// checks if the HandlerCreateExceptionalWkorkingDaysRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateExceptionalWkorkingDaysRequest{}

// HandlerCreateExceptionalWkorkingDaysRequest struct for HandlerCreateExceptionalWkorkingDaysRequest
type HandlerCreateExceptionalWkorkingDaysRequest struct {
	ApproverId int32 `json:"approver_id"`
	ExceptionalWkgDayDate string `json:"exceptional_wkg_day_date"`
	Remarks string `json:"remarks"`
}

type _HandlerCreateExceptionalWkorkingDaysRequest HandlerCreateExceptionalWkorkingDaysRequest

// NewHandlerCreateExceptionalWkorkingDaysRequest instantiates a new HandlerCreateExceptionalWkorkingDaysRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateExceptionalWkorkingDaysRequest(approverId int32, exceptionalWkgDayDate string, remarks string) *HandlerCreateExceptionalWkorkingDaysRequest {
	this := HandlerCreateExceptionalWkorkingDaysRequest{}
	this.ApproverId = approverId
	this.ExceptionalWkgDayDate = exceptionalWkgDayDate
	this.Remarks = remarks
	return &this
}

// NewHandlerCreateExceptionalWkorkingDaysRequestWithDefaults instantiates a new HandlerCreateExceptionalWkorkingDaysRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateExceptionalWkorkingDaysRequestWithDefaults() *HandlerCreateExceptionalWkorkingDaysRequest {
	this := HandlerCreateExceptionalWkorkingDaysRequest{}
	return &this
}

// GetApproverId returns the ApproverId field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetApproverId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) SetApproverId(v int32) {
	o.ApproverId = v
}

// GetExceptionalWkgDayDate returns the ExceptionalWkgDayDate field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetExceptionalWkgDayDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExceptionalWkgDayDate
}

// GetExceptionalWkgDayDateOk returns a tuple with the ExceptionalWkgDayDate field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetExceptionalWkgDayDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExceptionalWkgDayDate, true
}

// SetExceptionalWkgDayDate sets field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) SetExceptionalWkgDayDate(v string) {
	o.ExceptionalWkgDayDate = v
}

// GetRemarks returns the Remarks field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateExceptionalWkorkingDaysRequest) GetRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remarks, true
}

// SetRemarks sets field value
func (o *HandlerCreateExceptionalWkorkingDaysRequest) SetRemarks(v string) {
	o.Remarks = v
}

func (o HandlerCreateExceptionalWkorkingDaysRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateExceptionalWkorkingDaysRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approver_id"] = o.ApproverId
	toSerialize["exceptional_wkg_day_date"] = o.ExceptionalWkgDayDate
	toSerialize["remarks"] = o.Remarks
	return toSerialize, nil
}

func (o *HandlerCreateExceptionalWkorkingDaysRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approver_id",
		"exceptional_wkg_day_date",
		"remarks",
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

	varHandlerCreateExceptionalWkorkingDaysRequest := _HandlerCreateExceptionalWkorkingDaysRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateExceptionalWkorkingDaysRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateExceptionalWkorkingDaysRequest(varHandlerCreateExceptionalWkorkingDaysRequest)

	return err
}

type NullableHandlerCreateExceptionalWkorkingDaysRequest struct {
	value *HandlerCreateExceptionalWkorkingDaysRequest
	isSet bool
}

func (v NullableHandlerCreateExceptionalWkorkingDaysRequest) Get() *HandlerCreateExceptionalWkorkingDaysRequest {
	return v.value
}

func (v *NullableHandlerCreateExceptionalWkorkingDaysRequest) Set(val *HandlerCreateExceptionalWkorkingDaysRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateExceptionalWkorkingDaysRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateExceptionalWkorkingDaysRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateExceptionalWkorkingDaysRequest(val *HandlerCreateExceptionalWkorkingDaysRequest) *NullableHandlerCreateExceptionalWkorkingDaysRequest {
	return &NullableHandlerCreateExceptionalWkorkingDaysRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateExceptionalWkorkingDaysRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateExceptionalWkorkingDaysRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



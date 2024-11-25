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

// checks if the HandlerUpdateClosingBalancesTransferRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerUpdateClosingBalancesTransferRequest{}

// HandlerUpdateClosingBalancesTransferRequest struct for HandlerUpdateClosingBalancesTransferRequest
type HandlerUpdateClosingBalancesTransferRequest struct {
	ApproverId int32 `json:"approver_id"`
	CashBal *float32 `json:"cash_bal,omitempty"`
	FromUserId int32 `json:"from_user_id"`
	Remarks string `json:"remarks"`
	ToUserId int32 `json:"to_user_id"`
	Type HandlerStampStatus `json:"type"`
}

type _HandlerUpdateClosingBalancesTransferRequest HandlerUpdateClosingBalancesTransferRequest

// NewHandlerUpdateClosingBalancesTransferRequest instantiates a new HandlerUpdateClosingBalancesTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerUpdateClosingBalancesTransferRequest(approverId int32, fromUserId int32, remarks string, toUserId int32, type_ HandlerStampStatus) *HandlerUpdateClosingBalancesTransferRequest {
	this := HandlerUpdateClosingBalancesTransferRequest{}
	this.ApproverId = approverId
	this.FromUserId = fromUserId
	this.Remarks = remarks
	this.ToUserId = toUserId
	this.Type = type_
	return &this
}

// NewHandlerUpdateClosingBalancesTransferRequestWithDefaults instantiates a new HandlerUpdateClosingBalancesTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerUpdateClosingBalancesTransferRequestWithDefaults() *HandlerUpdateClosingBalancesTransferRequest {
	this := HandlerUpdateClosingBalancesTransferRequest{}
	return &this
}

// GetApproverId returns the ApproverId field value
func (o *HandlerUpdateClosingBalancesTransferRequest) GetApproverId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *HandlerUpdateClosingBalancesTransferRequest) SetApproverId(v int32) {
	o.ApproverId = v
}

// GetCashBal returns the CashBal field value if set, zero value otherwise.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetCashBal() float32 {
	if o == nil || IsNil(o.CashBal) {
		var ret float32
		return ret
	}
	return *o.CashBal
}

// GetCashBalOk returns a tuple with the CashBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetCashBalOk() (*float32, bool) {
	if o == nil || IsNil(o.CashBal) {
		return nil, false
	}
	return o.CashBal, true
}

// HasCashBal returns a boolean if a field has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) HasCashBal() bool {
	if o != nil && !IsNil(o.CashBal) {
		return true
	}

	return false
}

// SetCashBal gets a reference to the given float32 and assigns it to the CashBal field.
func (o *HandlerUpdateClosingBalancesTransferRequest) SetCashBal(v float32) {
	o.CashBal = &v
}

// GetFromUserId returns the FromUserId field value
func (o *HandlerUpdateClosingBalancesTransferRequest) GetFromUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetFromUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromUserId, true
}

// SetFromUserId sets field value
func (o *HandlerUpdateClosingBalancesTransferRequest) SetFromUserId(v int32) {
	o.FromUserId = v
}

// GetRemarks returns the Remarks field value
func (o *HandlerUpdateClosingBalancesTransferRequest) GetRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remarks, true
}

// SetRemarks sets field value
func (o *HandlerUpdateClosingBalancesTransferRequest) SetRemarks(v string) {
	o.Remarks = v
}

// GetToUserId returns the ToUserId field value
func (o *HandlerUpdateClosingBalancesTransferRequest) GetToUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToUserId
}

// GetToUserIdOk returns a tuple with the ToUserId field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetToUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToUserId, true
}

// SetToUserId sets field value
func (o *HandlerUpdateClosingBalancesTransferRequest) SetToUserId(v int32) {
	o.ToUserId = v
}

// GetType returns the Type field value
func (o *HandlerUpdateClosingBalancesTransferRequest) GetType() HandlerStampStatus {
	if o == nil {
		var ret HandlerStampStatus
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateClosingBalancesTransferRequest) GetTypeOk() (*HandlerStampStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HandlerUpdateClosingBalancesTransferRequest) SetType(v HandlerStampStatus) {
	o.Type = v
}

func (o HandlerUpdateClosingBalancesTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerUpdateClosingBalancesTransferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["approver_id"] = o.ApproverId
	if !IsNil(o.CashBal) {
		toSerialize["cash_bal"] = o.CashBal
	}
	toSerialize["from_user_id"] = o.FromUserId
	toSerialize["remarks"] = o.Remarks
	toSerialize["to_user_id"] = o.ToUserId
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *HandlerUpdateClosingBalancesTransferRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"approver_id",
		"from_user_id",
		"remarks",
		"to_user_id",
		"type",
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

	varHandlerUpdateClosingBalancesTransferRequest := _HandlerUpdateClosingBalancesTransferRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerUpdateClosingBalancesTransferRequest)

	if err != nil {
		return err
	}

	*o = HandlerUpdateClosingBalancesTransferRequest(varHandlerUpdateClosingBalancesTransferRequest)

	return err
}

type NullableHandlerUpdateClosingBalancesTransferRequest struct {
	value *HandlerUpdateClosingBalancesTransferRequest
	isSet bool
}

func (v NullableHandlerUpdateClosingBalancesTransferRequest) Get() *HandlerUpdateClosingBalancesTransferRequest {
	return v.value
}

func (v *NullableHandlerUpdateClosingBalancesTransferRequest) Set(val *HandlerUpdateClosingBalancesTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerUpdateClosingBalancesTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerUpdateClosingBalancesTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerUpdateClosingBalancesTransferRequest(val *HandlerUpdateClosingBalancesTransferRequest) *NullableHandlerUpdateClosingBalancesTransferRequest {
	return &NullableHandlerUpdateClosingBalancesTransferRequest{value: val, isSet: true}
}

func (v NullableHandlerUpdateClosingBalancesTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerUpdateClosingBalancesTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


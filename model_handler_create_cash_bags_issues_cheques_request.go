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

// checks if the HandlerCreateCashBagsIssuesChequesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateCashBagsIssuesChequesRequest{}

// HandlerCreateCashBagsIssuesChequesRequest struct for HandlerCreateCashBagsIssuesChequesRequest
type HandlerCreateCashBagsIssuesChequesRequest struct {
	FromOfficeId int32 `json:"from-office-id"`
	Remarks string `json:"remarks"`
	ToOfficeId int32 `json:"to-office-id"`
	TransAmount float32 `json:"trans-amount"`
	TransDetails string `json:"trans-details"`
	TransId string `json:"trans-id"`
}

type _HandlerCreateCashBagsIssuesChequesRequest HandlerCreateCashBagsIssuesChequesRequest

// NewHandlerCreateCashBagsIssuesChequesRequest instantiates a new HandlerCreateCashBagsIssuesChequesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateCashBagsIssuesChequesRequest(fromOfficeId int32, remarks string, toOfficeId int32, transAmount float32, transDetails string, transId string) *HandlerCreateCashBagsIssuesChequesRequest {
	this := HandlerCreateCashBagsIssuesChequesRequest{}
	this.FromOfficeId = fromOfficeId
	this.Remarks = remarks
	this.ToOfficeId = toOfficeId
	this.TransAmount = transAmount
	this.TransDetails = transDetails
	this.TransId = transId
	return &this
}

// NewHandlerCreateCashBagsIssuesChequesRequestWithDefaults instantiates a new HandlerCreateCashBagsIssuesChequesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateCashBagsIssuesChequesRequestWithDefaults() *HandlerCreateCashBagsIssuesChequesRequest {
	this := HandlerCreateCashBagsIssuesChequesRequest{}
	return &this
}

// GetFromOfficeId returns the FromOfficeId field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetFromOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FromOfficeId
}

// GetFromOfficeIdOk returns a tuple with the FromOfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetFromOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromOfficeId, true
}

// SetFromOfficeId sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetFromOfficeId(v int32) {
	o.FromOfficeId = v
}

// GetRemarks returns the Remarks field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remarks, true
}

// SetRemarks sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetRemarks(v string) {
	o.Remarks = v
}

// GetToOfficeId returns the ToOfficeId field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetToOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToOfficeId
}

// GetToOfficeIdOk returns a tuple with the ToOfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetToOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToOfficeId, true
}

// SetToOfficeId sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetToOfficeId(v int32) {
	o.ToOfficeId = v
}

// GetTransAmount returns the TransAmount field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransAmount
}

// GetTransAmountOk returns a tuple with the TransAmount field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransAmount, true
}

// SetTransAmount sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetTransAmount(v float32) {
	o.TransAmount = v
}

// GetTransDetails returns the TransDetails field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransDetails
}

// GetTransDetailsOk returns a tuple with the TransDetails field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransDetails, true
}

// SetTransDetails sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetTransDetails(v string) {
	o.TransDetails = v
}

// GetTransId returns the TransId field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateCashBagsIssuesChequesRequest) GetTransIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransId, true
}

// SetTransId sets field value
func (o *HandlerCreateCashBagsIssuesChequesRequest) SetTransId(v string) {
	o.TransId = v
}

func (o HandlerCreateCashBagsIssuesChequesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateCashBagsIssuesChequesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from-office-id"] = o.FromOfficeId
	toSerialize["remarks"] = o.Remarks
	toSerialize["to-office-id"] = o.ToOfficeId
	toSerialize["trans-amount"] = o.TransAmount
	toSerialize["trans-details"] = o.TransDetails
	toSerialize["trans-id"] = o.TransId
	return toSerialize, nil
}

func (o *HandlerCreateCashBagsIssuesChequesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from-office-id",
		"remarks",
		"to-office-id",
		"trans-amount",
		"trans-details",
		"trans-id",
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

	varHandlerCreateCashBagsIssuesChequesRequest := _HandlerCreateCashBagsIssuesChequesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateCashBagsIssuesChequesRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateCashBagsIssuesChequesRequest(varHandlerCreateCashBagsIssuesChequesRequest)

	return err
}

type NullableHandlerCreateCashBagsIssuesChequesRequest struct {
	value *HandlerCreateCashBagsIssuesChequesRequest
	isSet bool
}

func (v NullableHandlerCreateCashBagsIssuesChequesRequest) Get() *HandlerCreateCashBagsIssuesChequesRequest {
	return v.value
}

func (v *NullableHandlerCreateCashBagsIssuesChequesRequest) Set(val *HandlerCreateCashBagsIssuesChequesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateCashBagsIssuesChequesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateCashBagsIssuesChequesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateCashBagsIssuesChequesRequest(val *HandlerCreateCashBagsIssuesChequesRequest) *NullableHandlerCreateCashBagsIssuesChequesRequest {
	return &NullableHandlerCreateCashBagsIssuesChequesRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateCashBagsIssuesChequesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateCashBagsIssuesChequesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


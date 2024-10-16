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

// checks if the HandlerUpdateOfficeTransactionsBankCreditLimitsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerUpdateOfficeTransactionsBankCreditLimitsRequest{}

// HandlerUpdateOfficeTransactionsBankCreditLimitsRequest struct for HandlerUpdateOfficeTransactionsBankCreditLimitsRequest
type HandlerUpdateOfficeTransactionsBankCreditLimitsRequest struct {
	BankCreditLimit float32 `json:"bank_credit_limit"`
}

type _HandlerUpdateOfficeTransactionsBankCreditLimitsRequest HandlerUpdateOfficeTransactionsBankCreditLimitsRequest

// NewHandlerUpdateOfficeTransactionsBankCreditLimitsRequest instantiates a new HandlerUpdateOfficeTransactionsBankCreditLimitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerUpdateOfficeTransactionsBankCreditLimitsRequest(bankCreditLimit float32) *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest {
	this := HandlerUpdateOfficeTransactionsBankCreditLimitsRequest{}
	this.BankCreditLimit = bankCreditLimit
	return &this
}

// NewHandlerUpdateOfficeTransactionsBankCreditLimitsRequestWithDefaults instantiates a new HandlerUpdateOfficeTransactionsBankCreditLimitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerUpdateOfficeTransactionsBankCreditLimitsRequestWithDefaults() *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest {
	this := HandlerUpdateOfficeTransactionsBankCreditLimitsRequest{}
	return &this
}

// GetBankCreditLimit returns the BankCreditLimit field value
func (o *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) GetBankCreditLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BankCreditLimit
}

// GetBankCreditLimitOk returns a tuple with the BankCreditLimit field value
// and a boolean to check if the value has been set.
func (o *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) GetBankCreditLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankCreditLimit, true
}

// SetBankCreditLimit sets field value
func (o *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) SetBankCreditLimit(v float32) {
	o.BankCreditLimit = v
}

func (o HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bank_credit_limit"] = o.BankCreditLimit
	return toSerialize, nil
}

func (o *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bank_credit_limit",
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

	varHandlerUpdateOfficeTransactionsBankCreditLimitsRequest := _HandlerUpdateOfficeTransactionsBankCreditLimitsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerUpdateOfficeTransactionsBankCreditLimitsRequest)

	if err != nil {
		return err
	}

	*o = HandlerUpdateOfficeTransactionsBankCreditLimitsRequest(varHandlerUpdateOfficeTransactionsBankCreditLimitsRequest)

	return err
}

type NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest struct {
	value *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest
	isSet bool
}

func (v NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) Get() *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest {
	return v.value
}

func (v *NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) Set(val *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest(val *HandlerUpdateOfficeTransactionsBankCreditLimitsRequest) *NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest {
	return &NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest{value: val, isSet: true}
}

func (v NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerUpdateOfficeTransactionsBankCreditLimitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



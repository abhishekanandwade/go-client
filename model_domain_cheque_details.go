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

// checks if the DomainChequeDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainChequeDetails{}

// DomainChequeDetails struct for DomainChequeDetails
type DomainChequeDetails struct {
	ChequeAmount *float32 `json:"cheque_amount,omitempty"`
	ChequeDate *string `json:"cheque_date,omitempty"`
	ChequeNo *string `json:"cheque_no,omitempty"`
}

// NewDomainChequeDetails instantiates a new DomainChequeDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainChequeDetails() *DomainChequeDetails {
	this := DomainChequeDetails{}
	return &this
}

// NewDomainChequeDetailsWithDefaults instantiates a new DomainChequeDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainChequeDetailsWithDefaults() *DomainChequeDetails {
	this := DomainChequeDetails{}
	return &this
}

// GetChequeAmount returns the ChequeAmount field value if set, zero value otherwise.
func (o *DomainChequeDetails) GetChequeAmount() float32 {
	if o == nil || IsNil(o.ChequeAmount) {
		var ret float32
		return ret
	}
	return *o.ChequeAmount
}

// GetChequeAmountOk returns a tuple with the ChequeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainChequeDetails) GetChequeAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ChequeAmount) {
		return nil, false
	}
	return o.ChequeAmount, true
}

// HasChequeAmount returns a boolean if a field has been set.
func (o *DomainChequeDetails) HasChequeAmount() bool {
	if o != nil && !IsNil(o.ChequeAmount) {
		return true
	}

	return false
}

// SetChequeAmount gets a reference to the given float32 and assigns it to the ChequeAmount field.
func (o *DomainChequeDetails) SetChequeAmount(v float32) {
	o.ChequeAmount = &v
}

// GetChequeDate returns the ChequeDate field value if set, zero value otherwise.
func (o *DomainChequeDetails) GetChequeDate() string {
	if o == nil || IsNil(o.ChequeDate) {
		var ret string
		return ret
	}
	return *o.ChequeDate
}

// GetChequeDateOk returns a tuple with the ChequeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainChequeDetails) GetChequeDateOk() (*string, bool) {
	if o == nil || IsNil(o.ChequeDate) {
		return nil, false
	}
	return o.ChequeDate, true
}

// HasChequeDate returns a boolean if a field has been set.
func (o *DomainChequeDetails) HasChequeDate() bool {
	if o != nil && !IsNil(o.ChequeDate) {
		return true
	}

	return false
}

// SetChequeDate gets a reference to the given string and assigns it to the ChequeDate field.
func (o *DomainChequeDetails) SetChequeDate(v string) {
	o.ChequeDate = &v
}

// GetChequeNo returns the ChequeNo field value if set, zero value otherwise.
func (o *DomainChequeDetails) GetChequeNo() string {
	if o == nil || IsNil(o.ChequeNo) {
		var ret string
		return ret
	}
	return *o.ChequeNo
}

// GetChequeNoOk returns a tuple with the ChequeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainChequeDetails) GetChequeNoOk() (*string, bool) {
	if o == nil || IsNil(o.ChequeNo) {
		return nil, false
	}
	return o.ChequeNo, true
}

// HasChequeNo returns a boolean if a field has been set.
func (o *DomainChequeDetails) HasChequeNo() bool {
	if o != nil && !IsNil(o.ChequeNo) {
		return true
	}

	return false
}

// SetChequeNo gets a reference to the given string and assigns it to the ChequeNo field.
func (o *DomainChequeDetails) SetChequeNo(v string) {
	o.ChequeNo = &v
}

func (o DomainChequeDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainChequeDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChequeAmount) {
		toSerialize["cheque_amount"] = o.ChequeAmount
	}
	if !IsNil(o.ChequeDate) {
		toSerialize["cheque_date"] = o.ChequeDate
	}
	if !IsNil(o.ChequeNo) {
		toSerialize["cheque_no"] = o.ChequeNo
	}
	return toSerialize, nil
}

type NullableDomainChequeDetails struct {
	value *DomainChequeDetails
	isSet bool
}

func (v NullableDomainChequeDetails) Get() *DomainChequeDetails {
	return v.value
}

func (v *NullableDomainChequeDetails) Set(val *DomainChequeDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainChequeDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainChequeDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainChequeDetails(val *DomainChequeDetails) *NullableDomainChequeDetails {
	return &NullableDomainChequeDetails{value: val, isSet: true}
}

func (v NullableDomainChequeDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainChequeDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



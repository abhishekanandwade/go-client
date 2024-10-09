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

// checks if the ResponseCashBagDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCashBagDetail{}

// ResponseCashBagDetail struct for ResponseCashBagDetail
type ResponseCashBagDetail struct {
	BagWeight *string `json:"bag_weight,omitempty"`
	CashBagId *string `json:"cash_bag_id,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	TransDetails *string `json:"trans_details,omitempty"`
	TransType *string `json:"trans_type,omitempty"`
}

// NewResponseCashBagDetail instantiates a new ResponseCashBagDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCashBagDetail() *ResponseCashBagDetail {
	this := ResponseCashBagDetail{}
	return &this
}

// NewResponseCashBagDetailWithDefaults instantiates a new ResponseCashBagDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCashBagDetailWithDefaults() *ResponseCashBagDetail {
	this := ResponseCashBagDetail{}
	return &this
}

// GetBagWeight returns the BagWeight field value if set, zero value otherwise.
func (o *ResponseCashBagDetail) GetBagWeight() string {
	if o == nil || IsNil(o.BagWeight) {
		var ret string
		return ret
	}
	return *o.BagWeight
}

// GetBagWeightOk returns a tuple with the BagWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashBagDetail) GetBagWeightOk() (*string, bool) {
	if o == nil || IsNil(o.BagWeight) {
		return nil, false
	}
	return o.BagWeight, true
}

// HasBagWeight returns a boolean if a field has been set.
func (o *ResponseCashBagDetail) HasBagWeight() bool {
	if o != nil && !IsNil(o.BagWeight) {
		return true
	}

	return false
}

// SetBagWeight gets a reference to the given string and assigns it to the BagWeight field.
func (o *ResponseCashBagDetail) SetBagWeight(v string) {
	o.BagWeight = &v
}

// GetCashBagId returns the CashBagId field value if set, zero value otherwise.
func (o *ResponseCashBagDetail) GetCashBagId() string {
	if o == nil || IsNil(o.CashBagId) {
		var ret string
		return ret
	}
	return *o.CashBagId
}

// GetCashBagIdOk returns a tuple with the CashBagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashBagDetail) GetCashBagIdOk() (*string, bool) {
	if o == nil || IsNil(o.CashBagId) {
		return nil, false
	}
	return o.CashBagId, true
}

// HasCashBagId returns a boolean if a field has been set.
func (o *ResponseCashBagDetail) HasCashBagId() bool {
	if o != nil && !IsNil(o.CashBagId) {
		return true
	}

	return false
}

// SetCashBagId gets a reference to the given string and assigns it to the CashBagId field.
func (o *ResponseCashBagDetail) SetCashBagId(v string) {
	o.CashBagId = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseCashBagDetail) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashBagDetail) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseCashBagDetail) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseCashBagDetail) SetRemarks(v string) {
	o.Remarks = &v
}

// GetTransDetails returns the TransDetails field value if set, zero value otherwise.
func (o *ResponseCashBagDetail) GetTransDetails() string {
	if o == nil || IsNil(o.TransDetails) {
		var ret string
		return ret
	}
	return *o.TransDetails
}

// GetTransDetailsOk returns a tuple with the TransDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashBagDetail) GetTransDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.TransDetails) {
		return nil, false
	}
	return o.TransDetails, true
}

// HasTransDetails returns a boolean if a field has been set.
func (o *ResponseCashBagDetail) HasTransDetails() bool {
	if o != nil && !IsNil(o.TransDetails) {
		return true
	}

	return false
}

// SetTransDetails gets a reference to the given string and assigns it to the TransDetails field.
func (o *ResponseCashBagDetail) SetTransDetails(v string) {
	o.TransDetails = &v
}

// GetTransType returns the TransType field value if set, zero value otherwise.
func (o *ResponseCashBagDetail) GetTransType() string {
	if o == nil || IsNil(o.TransType) {
		var ret string
		return ret
	}
	return *o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashBagDetail) GetTransTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransType) {
		return nil, false
	}
	return o.TransType, true
}

// HasTransType returns a boolean if a field has been set.
func (o *ResponseCashBagDetail) HasTransType() bool {
	if o != nil && !IsNil(o.TransType) {
		return true
	}

	return false
}

// SetTransType gets a reference to the given string and assigns it to the TransType field.
func (o *ResponseCashBagDetail) SetTransType(v string) {
	o.TransType = &v
}

func (o ResponseCashBagDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCashBagDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BagWeight) {
		toSerialize["bag_weight"] = o.BagWeight
	}
	if !IsNil(o.CashBagId) {
		toSerialize["cash_bag_id"] = o.CashBagId
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.TransDetails) {
		toSerialize["trans_details"] = o.TransDetails
	}
	if !IsNil(o.TransType) {
		toSerialize["trans_type"] = o.TransType
	}
	return toSerialize, nil
}

type NullableResponseCashBagDetail struct {
	value *ResponseCashBagDetail
	isSet bool
}

func (v NullableResponseCashBagDetail) Get() *ResponseCashBagDetail {
	return v.value
}

func (v *NullableResponseCashBagDetail) Set(val *ResponseCashBagDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCashBagDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCashBagDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCashBagDetail(val *ResponseCashBagDetail) *NullableResponseCashBagDetail {
	return &NullableResponseCashBagDetail{value: val, isSet: true}
}

func (v NullableResponseCashBagDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCashBagDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ResponseSubOrdinateOfficeSummarydetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSubOrdinateOfficeSummarydetails{}

// ResponseSubOrdinateOfficeSummarydetails struct for ResponseSubOrdinateOfficeSummarydetails
type ResponseSubOrdinateOfficeSummarydetails struct {
	ClosingBalance *float32 `json:"closing_balance,omitempty"`
	OfficeDate *string `json:"office_date,omitempty"`
	OpeningBalance *float32 `json:"opening_balance,omitempty"`
	SubordinateOfficeId *string `json:"subordinate_office_id,omitempty"`
}

// NewResponseSubOrdinateOfficeSummarydetails instantiates a new ResponseSubOrdinateOfficeSummarydetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSubOrdinateOfficeSummarydetails() *ResponseSubOrdinateOfficeSummarydetails {
	this := ResponseSubOrdinateOfficeSummarydetails{}
	return &this
}

// NewResponseSubOrdinateOfficeSummarydetailsWithDefaults instantiates a new ResponseSubOrdinateOfficeSummarydetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSubOrdinateOfficeSummarydetailsWithDefaults() *ResponseSubOrdinateOfficeSummarydetails {
	this := ResponseSubOrdinateOfficeSummarydetails{}
	return &this
}

// GetClosingBalance returns the ClosingBalance field value if set, zero value otherwise.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetClosingBalance() float32 {
	if o == nil || IsNil(o.ClosingBalance) {
		var ret float32
		return ret
	}
	return *o.ClosingBalance
}

// GetClosingBalanceOk returns a tuple with the ClosingBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetClosingBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.ClosingBalance) {
		return nil, false
	}
	return o.ClosingBalance, true
}

// HasClosingBalance returns a boolean if a field has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) HasClosingBalance() bool {
	if o != nil && !IsNil(o.ClosingBalance) {
		return true
	}

	return false
}

// SetClosingBalance gets a reference to the given float32 and assigns it to the ClosingBalance field.
func (o *ResponseSubOrdinateOfficeSummarydetails) SetClosingBalance(v float32) {
	o.ClosingBalance = &v
}

// GetOfficeDate returns the OfficeDate field value if set, zero value otherwise.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetOfficeDate() string {
	if o == nil || IsNil(o.OfficeDate) {
		var ret string
		return ret
	}
	return *o.OfficeDate
}

// GetOfficeDateOk returns a tuple with the OfficeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetOfficeDateOk() (*string, bool) {
	if o == nil || IsNil(o.OfficeDate) {
		return nil, false
	}
	return o.OfficeDate, true
}

// HasOfficeDate returns a boolean if a field has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) HasOfficeDate() bool {
	if o != nil && !IsNil(o.OfficeDate) {
		return true
	}

	return false
}

// SetOfficeDate gets a reference to the given string and assigns it to the OfficeDate field.
func (o *ResponseSubOrdinateOfficeSummarydetails) SetOfficeDate(v string) {
	o.OfficeDate = &v
}

// GetOpeningBalance returns the OpeningBalance field value if set, zero value otherwise.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetOpeningBalance() float32 {
	if o == nil || IsNil(o.OpeningBalance) {
		var ret float32
		return ret
	}
	return *o.OpeningBalance
}

// GetOpeningBalanceOk returns a tuple with the OpeningBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetOpeningBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.OpeningBalance) {
		return nil, false
	}
	return o.OpeningBalance, true
}

// HasOpeningBalance returns a boolean if a field has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) HasOpeningBalance() bool {
	if o != nil && !IsNil(o.OpeningBalance) {
		return true
	}

	return false
}

// SetOpeningBalance gets a reference to the given float32 and assigns it to the OpeningBalance field.
func (o *ResponseSubOrdinateOfficeSummarydetails) SetOpeningBalance(v float32) {
	o.OpeningBalance = &v
}

// GetSubordinateOfficeId returns the SubordinateOfficeId field value if set, zero value otherwise.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetSubordinateOfficeId() string {
	if o == nil || IsNil(o.SubordinateOfficeId) {
		var ret string
		return ret
	}
	return *o.SubordinateOfficeId
}

// GetSubordinateOfficeIdOk returns a tuple with the SubordinateOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) GetSubordinateOfficeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubordinateOfficeId) {
		return nil, false
	}
	return o.SubordinateOfficeId, true
}

// HasSubordinateOfficeId returns a boolean if a field has been set.
func (o *ResponseSubOrdinateOfficeSummarydetails) HasSubordinateOfficeId() bool {
	if o != nil && !IsNil(o.SubordinateOfficeId) {
		return true
	}

	return false
}

// SetSubordinateOfficeId gets a reference to the given string and assigns it to the SubordinateOfficeId field.
func (o *ResponseSubOrdinateOfficeSummarydetails) SetSubordinateOfficeId(v string) {
	o.SubordinateOfficeId = &v
}

func (o ResponseSubOrdinateOfficeSummarydetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSubOrdinateOfficeSummarydetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClosingBalance) {
		toSerialize["closing_balance"] = o.ClosingBalance
	}
	if !IsNil(o.OfficeDate) {
		toSerialize["office_date"] = o.OfficeDate
	}
	if !IsNil(o.OpeningBalance) {
		toSerialize["opening_balance"] = o.OpeningBalance
	}
	if !IsNil(o.SubordinateOfficeId) {
		toSerialize["subordinate_office_id"] = o.SubordinateOfficeId
	}
	return toSerialize, nil
}

type NullableResponseSubOrdinateOfficeSummarydetails struct {
	value *ResponseSubOrdinateOfficeSummarydetails
	isSet bool
}

func (v NullableResponseSubOrdinateOfficeSummarydetails) Get() *ResponseSubOrdinateOfficeSummarydetails {
	return v.value
}

func (v *NullableResponseSubOrdinateOfficeSummarydetails) Set(val *ResponseSubOrdinateOfficeSummarydetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSubOrdinateOfficeSummarydetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSubOrdinateOfficeSummarydetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSubOrdinateOfficeSummarydetails(val *ResponseSubOrdinateOfficeSummarydetails) *NullableResponseSubOrdinateOfficeSummarydetails {
	return &NullableResponseSubOrdinateOfficeSummarydetails{value: val, isSet: true}
}

func (v NullableResponseSubOrdinateOfficeSummarydetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSubOrdinateOfficeSummarydetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



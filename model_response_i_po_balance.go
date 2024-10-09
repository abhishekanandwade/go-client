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

// checks if the ResponseIPOBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIPOBalance{}

// ResponseIPOBalance struct for ResponseIPOBalance
type ResponseIPOBalance struct {
	ClosingBal *int32 `json:"closing_bal,omitempty"`
	DenominationId *string `json:"denomination_id,omitempty"`
	IpoDetails *string `json:"ipo_details,omitempty"`
	Issues *int32 `json:"issues,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	OpeningBal *int32 `json:"opening_bal,omitempty"`
	Receipts *int32 `json:"receipts,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
}

// NewResponseIPOBalance instantiates a new ResponseIPOBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIPOBalance() *ResponseIPOBalance {
	this := ResponseIPOBalance{}
	return &this
}

// NewResponseIPOBalanceWithDefaults instantiates a new ResponseIPOBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIPOBalanceWithDefaults() *ResponseIPOBalance {
	this := ResponseIPOBalance{}
	return &this
}

// GetClosingBal returns the ClosingBal field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetClosingBal() int32 {
	if o == nil || IsNil(o.ClosingBal) {
		var ret int32
		return ret
	}
	return *o.ClosingBal
}

// GetClosingBalOk returns a tuple with the ClosingBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetClosingBalOk() (*int32, bool) {
	if o == nil || IsNil(o.ClosingBal) {
		return nil, false
	}
	return o.ClosingBal, true
}

// HasClosingBal returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasClosingBal() bool {
	if o != nil && !IsNil(o.ClosingBal) {
		return true
	}

	return false
}

// SetClosingBal gets a reference to the given int32 and assigns it to the ClosingBal field.
func (o *ResponseIPOBalance) SetClosingBal(v int32) {
	o.ClosingBal = &v
}

// GetDenominationId returns the DenominationId field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetDenominationId() string {
	if o == nil || IsNil(o.DenominationId) {
		var ret string
		return ret
	}
	return *o.DenominationId
}

// GetDenominationIdOk returns a tuple with the DenominationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetDenominationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DenominationId) {
		return nil, false
	}
	return o.DenominationId, true
}

// HasDenominationId returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasDenominationId() bool {
	if o != nil && !IsNil(o.DenominationId) {
		return true
	}

	return false
}

// SetDenominationId gets a reference to the given string and assigns it to the DenominationId field.
func (o *ResponseIPOBalance) SetDenominationId(v string) {
	o.DenominationId = &v
}

// GetIpoDetails returns the IpoDetails field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetIpoDetails() string {
	if o == nil || IsNil(o.IpoDetails) {
		var ret string
		return ret
	}
	return *o.IpoDetails
}

// GetIpoDetailsOk returns a tuple with the IpoDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetIpoDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.IpoDetails) {
		return nil, false
	}
	return o.IpoDetails, true
}

// HasIpoDetails returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasIpoDetails() bool {
	if o != nil && !IsNil(o.IpoDetails) {
		return true
	}

	return false
}

// SetIpoDetails gets a reference to the given string and assigns it to the IpoDetails field.
func (o *ResponseIPOBalance) SetIpoDetails(v string) {
	o.IpoDetails = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetIssues() int32 {
	if o == nil || IsNil(o.Issues) {
		var ret int32
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetIssuesOk() (*int32, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given int32 and assigns it to the Issues field.
func (o *ResponseIPOBalance) SetIssues(v int32) {
	o.Issues = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseIPOBalance) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOpeningBal returns the OpeningBal field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetOpeningBal() int32 {
	if o == nil || IsNil(o.OpeningBal) {
		var ret int32
		return ret
	}
	return *o.OpeningBal
}

// GetOpeningBalOk returns a tuple with the OpeningBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetOpeningBalOk() (*int32, bool) {
	if o == nil || IsNil(o.OpeningBal) {
		return nil, false
	}
	return o.OpeningBal, true
}

// HasOpeningBal returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasOpeningBal() bool {
	if o != nil && !IsNil(o.OpeningBal) {
		return true
	}

	return false
}

// SetOpeningBal gets a reference to the given int32 and assigns it to the OpeningBal field.
func (o *ResponseIPOBalance) SetOpeningBal(v int32) {
	o.OpeningBal = &v
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetReceipts() int32 {
	if o == nil || IsNil(o.Receipts) {
		var ret int32
		return ret
	}
	return *o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetReceiptsOk() (*int32, bool) {
	if o == nil || IsNil(o.Receipts) {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasReceipts() bool {
	if o != nil && !IsNil(o.Receipts) {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given int32 and assigns it to the Receipts field.
func (o *ResponseIPOBalance) SetReceipts(v int32) {
	o.Receipts = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseIPOBalance) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalance) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseIPOBalance) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseIPOBalance) SetTransDate(v string) {
	o.TransDate = &v
}

func (o ResponseIPOBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIPOBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClosingBal) {
		toSerialize["closing_bal"] = o.ClosingBal
	}
	if !IsNil(o.DenominationId) {
		toSerialize["denomination_id"] = o.DenominationId
	}
	if !IsNil(o.IpoDetails) {
		toSerialize["ipo_details"] = o.IpoDetails
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.OpeningBal) {
		toSerialize["opening_bal"] = o.OpeningBal
	}
	if !IsNil(o.Receipts) {
		toSerialize["receipts"] = o.Receipts
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	return toSerialize, nil
}

type NullableResponseIPOBalance struct {
	value *ResponseIPOBalance
	isSet bool
}

func (v NullableResponseIPOBalance) Get() *ResponseIPOBalance {
	return v.value
}

func (v *NullableResponseIPOBalance) Set(val *ResponseIPOBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIPOBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIPOBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIPOBalance(val *ResponseIPOBalance) *NullableResponseIPOBalance {
	return &NullableResponseIPOBalance{value: val, isSet: true}
}

func (v NullableResponseIPOBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIPOBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



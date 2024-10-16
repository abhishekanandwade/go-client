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

// checks if the ResponseCashError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCashError{}

// ResponseCashError struct for ResponseCashError
type ResponseCashError struct {
	CashErrorDetails []ResponseCashErrorDetails `json:"cashErrorDetails,omitempty"`
	CurrOffice *int32 `json:"curr_office,omitempty"`
	DestAmt *float32 `json:"dest_amt,omitempty"`
	DestDetails *map[string]int32 `json:"dest_details,omitempty"`
	DestOffice *int32 `json:"dest_office,omitempty"`
	DestOfficeName *string `json:"dest_office_name,omitempty"`
	ErrorId *string `json:"error_id,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	SrcAmt *float32 `json:"src_amt,omitempty"`
	SrcDetails *map[string]int32 `json:"src_details,omitempty"`
	SrcOffice *int32 `json:"src_office,omitempty"`
	SrcOfficeName *string `json:"src_office_name,omitempty"`
	Status *string `json:"status,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransId *string `json:"trans_id,omitempty"`
	TransType *string `json:"trans_type,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
}

// NewResponseCashError instantiates a new ResponseCashError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCashError() *ResponseCashError {
	this := ResponseCashError{}
	return &this
}

// NewResponseCashErrorWithDefaults instantiates a new ResponseCashError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCashErrorWithDefaults() *ResponseCashError {
	this := ResponseCashError{}
	return &this
}

// GetCashErrorDetails returns the CashErrorDetails field value if set, zero value otherwise.
func (o *ResponseCashError) GetCashErrorDetails() []ResponseCashErrorDetails {
	if o == nil || IsNil(o.CashErrorDetails) {
		var ret []ResponseCashErrorDetails
		return ret
	}
	return o.CashErrorDetails
}

// GetCashErrorDetailsOk returns a tuple with the CashErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetCashErrorDetailsOk() ([]ResponseCashErrorDetails, bool) {
	if o == nil || IsNil(o.CashErrorDetails) {
		return nil, false
	}
	return o.CashErrorDetails, true
}

// HasCashErrorDetails returns a boolean if a field has been set.
func (o *ResponseCashError) HasCashErrorDetails() bool {
	if o != nil && !IsNil(o.CashErrorDetails) {
		return true
	}

	return false
}

// SetCashErrorDetails gets a reference to the given []ResponseCashErrorDetails and assigns it to the CashErrorDetails field.
func (o *ResponseCashError) SetCashErrorDetails(v []ResponseCashErrorDetails) {
	o.CashErrorDetails = v
}

// GetCurrOffice returns the CurrOffice field value if set, zero value otherwise.
func (o *ResponseCashError) GetCurrOffice() int32 {
	if o == nil || IsNil(o.CurrOffice) {
		var ret int32
		return ret
	}
	return *o.CurrOffice
}

// GetCurrOfficeOk returns a tuple with the CurrOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetCurrOfficeOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrOffice) {
		return nil, false
	}
	return o.CurrOffice, true
}

// HasCurrOffice returns a boolean if a field has been set.
func (o *ResponseCashError) HasCurrOffice() bool {
	if o != nil && !IsNil(o.CurrOffice) {
		return true
	}

	return false
}

// SetCurrOffice gets a reference to the given int32 and assigns it to the CurrOffice field.
func (o *ResponseCashError) SetCurrOffice(v int32) {
	o.CurrOffice = &v
}

// GetDestAmt returns the DestAmt field value if set, zero value otherwise.
func (o *ResponseCashError) GetDestAmt() float32 {
	if o == nil || IsNil(o.DestAmt) {
		var ret float32
		return ret
	}
	return *o.DestAmt
}

// GetDestAmtOk returns a tuple with the DestAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetDestAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.DestAmt) {
		return nil, false
	}
	return o.DestAmt, true
}

// HasDestAmt returns a boolean if a field has been set.
func (o *ResponseCashError) HasDestAmt() bool {
	if o != nil && !IsNil(o.DestAmt) {
		return true
	}

	return false
}

// SetDestAmt gets a reference to the given float32 and assigns it to the DestAmt field.
func (o *ResponseCashError) SetDestAmt(v float32) {
	o.DestAmt = &v
}

// GetDestDetails returns the DestDetails field value if set, zero value otherwise.
func (o *ResponseCashError) GetDestDetails() map[string]int32 {
	if o == nil || IsNil(o.DestDetails) {
		var ret map[string]int32
		return ret
	}
	return *o.DestDetails
}

// GetDestDetailsOk returns a tuple with the DestDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetDestDetailsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.DestDetails) {
		return nil, false
	}
	return o.DestDetails, true
}

// HasDestDetails returns a boolean if a field has been set.
func (o *ResponseCashError) HasDestDetails() bool {
	if o != nil && !IsNil(o.DestDetails) {
		return true
	}

	return false
}

// SetDestDetails gets a reference to the given map[string]int32 and assigns it to the DestDetails field.
func (o *ResponseCashError) SetDestDetails(v map[string]int32) {
	o.DestDetails = &v
}

// GetDestOffice returns the DestOffice field value if set, zero value otherwise.
func (o *ResponseCashError) GetDestOffice() int32 {
	if o == nil || IsNil(o.DestOffice) {
		var ret int32
		return ret
	}
	return *o.DestOffice
}

// GetDestOfficeOk returns a tuple with the DestOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetDestOfficeOk() (*int32, bool) {
	if o == nil || IsNil(o.DestOffice) {
		return nil, false
	}
	return o.DestOffice, true
}

// HasDestOffice returns a boolean if a field has been set.
func (o *ResponseCashError) HasDestOffice() bool {
	if o != nil && !IsNil(o.DestOffice) {
		return true
	}

	return false
}

// SetDestOffice gets a reference to the given int32 and assigns it to the DestOffice field.
func (o *ResponseCashError) SetDestOffice(v int32) {
	o.DestOffice = &v
}

// GetDestOfficeName returns the DestOfficeName field value if set, zero value otherwise.
func (o *ResponseCashError) GetDestOfficeName() string {
	if o == nil || IsNil(o.DestOfficeName) {
		var ret string
		return ret
	}
	return *o.DestOfficeName
}

// GetDestOfficeNameOk returns a tuple with the DestOfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetDestOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.DestOfficeName) {
		return nil, false
	}
	return o.DestOfficeName, true
}

// HasDestOfficeName returns a boolean if a field has been set.
func (o *ResponseCashError) HasDestOfficeName() bool {
	if o != nil && !IsNil(o.DestOfficeName) {
		return true
	}

	return false
}

// SetDestOfficeName gets a reference to the given string and assigns it to the DestOfficeName field.
func (o *ResponseCashError) SetDestOfficeName(v string) {
	o.DestOfficeName = &v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *ResponseCashError) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *ResponseCashError) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *ResponseCashError) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseCashError) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseCashError) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseCashError) SetRemarks(v string) {
	o.Remarks = &v
}

// GetSrcAmt returns the SrcAmt field value if set, zero value otherwise.
func (o *ResponseCashError) GetSrcAmt() float32 {
	if o == nil || IsNil(o.SrcAmt) {
		var ret float32
		return ret
	}
	return *o.SrcAmt
}

// GetSrcAmtOk returns a tuple with the SrcAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetSrcAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.SrcAmt) {
		return nil, false
	}
	return o.SrcAmt, true
}

// HasSrcAmt returns a boolean if a field has been set.
func (o *ResponseCashError) HasSrcAmt() bool {
	if o != nil && !IsNil(o.SrcAmt) {
		return true
	}

	return false
}

// SetSrcAmt gets a reference to the given float32 and assigns it to the SrcAmt field.
func (o *ResponseCashError) SetSrcAmt(v float32) {
	o.SrcAmt = &v
}

// GetSrcDetails returns the SrcDetails field value if set, zero value otherwise.
func (o *ResponseCashError) GetSrcDetails() map[string]int32 {
	if o == nil || IsNil(o.SrcDetails) {
		var ret map[string]int32
		return ret
	}
	return *o.SrcDetails
}

// GetSrcDetailsOk returns a tuple with the SrcDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetSrcDetailsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.SrcDetails) {
		return nil, false
	}
	return o.SrcDetails, true
}

// HasSrcDetails returns a boolean if a field has been set.
func (o *ResponseCashError) HasSrcDetails() bool {
	if o != nil && !IsNil(o.SrcDetails) {
		return true
	}

	return false
}

// SetSrcDetails gets a reference to the given map[string]int32 and assigns it to the SrcDetails field.
func (o *ResponseCashError) SetSrcDetails(v map[string]int32) {
	o.SrcDetails = &v
}

// GetSrcOffice returns the SrcOffice field value if set, zero value otherwise.
func (o *ResponseCashError) GetSrcOffice() int32 {
	if o == nil || IsNil(o.SrcOffice) {
		var ret int32
		return ret
	}
	return *o.SrcOffice
}

// GetSrcOfficeOk returns a tuple with the SrcOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetSrcOfficeOk() (*int32, bool) {
	if o == nil || IsNil(o.SrcOffice) {
		return nil, false
	}
	return o.SrcOffice, true
}

// HasSrcOffice returns a boolean if a field has been set.
func (o *ResponseCashError) HasSrcOffice() bool {
	if o != nil && !IsNil(o.SrcOffice) {
		return true
	}

	return false
}

// SetSrcOffice gets a reference to the given int32 and assigns it to the SrcOffice field.
func (o *ResponseCashError) SetSrcOffice(v int32) {
	o.SrcOffice = &v
}

// GetSrcOfficeName returns the SrcOfficeName field value if set, zero value otherwise.
func (o *ResponseCashError) GetSrcOfficeName() string {
	if o == nil || IsNil(o.SrcOfficeName) {
		var ret string
		return ret
	}
	return *o.SrcOfficeName
}

// GetSrcOfficeNameOk returns a tuple with the SrcOfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetSrcOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.SrcOfficeName) {
		return nil, false
	}
	return o.SrcOfficeName, true
}

// HasSrcOfficeName returns a boolean if a field has been set.
func (o *ResponseCashError) HasSrcOfficeName() bool {
	if o != nil && !IsNil(o.SrcOfficeName) {
		return true
	}

	return false
}

// SetSrcOfficeName gets a reference to the given string and assigns it to the SrcOfficeName field.
func (o *ResponseCashError) SetSrcOfficeName(v string) {
	o.SrcOfficeName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ResponseCashError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ResponseCashError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ResponseCashError) SetStatus(v string) {
	o.Status = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseCashError) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseCashError) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseCashError) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransId returns the TransId field value if set, zero value otherwise.
func (o *ResponseCashError) GetTransId() string {
	if o == nil || IsNil(o.TransId) {
		var ret string
		return ret
	}
	return *o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetTransIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransId) {
		return nil, false
	}
	return o.TransId, true
}

// HasTransId returns a boolean if a field has been set.
func (o *ResponseCashError) HasTransId() bool {
	if o != nil && !IsNil(o.TransId) {
		return true
	}

	return false
}

// SetTransId gets a reference to the given string and assigns it to the TransId field.
func (o *ResponseCashError) SetTransId(v string) {
	o.TransId = &v
}

// GetTransType returns the TransType field value if set, zero value otherwise.
func (o *ResponseCashError) GetTransType() string {
	if o == nil || IsNil(o.TransType) {
		var ret string
		return ret
	}
	return *o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetTransTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TransType) {
		return nil, false
	}
	return o.TransType, true
}

// HasTransType returns a boolean if a field has been set.
func (o *ResponseCashError) HasTransType() bool {
	if o != nil && !IsNil(o.TransType) {
		return true
	}

	return false
}

// SetTransType gets a reference to the given string and assigns it to the TransType field.
func (o *ResponseCashError) SetTransType(v string) {
	o.TransType = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ResponseCashError) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashError) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ResponseCashError) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ResponseCashError) SetUserId(v int32) {
	o.UserId = &v
}

func (o ResponseCashError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCashError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CashErrorDetails) {
		toSerialize["cashErrorDetails"] = o.CashErrorDetails
	}
	if !IsNil(o.CurrOffice) {
		toSerialize["curr_office"] = o.CurrOffice
	}
	if !IsNil(o.DestAmt) {
		toSerialize["dest_amt"] = o.DestAmt
	}
	if !IsNil(o.DestDetails) {
		toSerialize["dest_details"] = o.DestDetails
	}
	if !IsNil(o.DestOffice) {
		toSerialize["dest_office"] = o.DestOffice
	}
	if !IsNil(o.DestOfficeName) {
		toSerialize["dest_office_name"] = o.DestOfficeName
	}
	if !IsNil(o.ErrorId) {
		toSerialize["error_id"] = o.ErrorId
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.SrcAmt) {
		toSerialize["src_amt"] = o.SrcAmt
	}
	if !IsNil(o.SrcDetails) {
		toSerialize["src_details"] = o.SrcDetails
	}
	if !IsNil(o.SrcOffice) {
		toSerialize["src_office"] = o.SrcOffice
	}
	if !IsNil(o.SrcOfficeName) {
		toSerialize["src_office_name"] = o.SrcOfficeName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransId) {
		toSerialize["trans_id"] = o.TransId
	}
	if !IsNil(o.TransType) {
		toSerialize["trans_type"] = o.TransType
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableResponseCashError struct {
	value *ResponseCashError
	isSet bool
}

func (v NullableResponseCashError) Get() *ResponseCashError {
	return v.value
}

func (v *NullableResponseCashError) Set(val *ResponseCashError) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCashError) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCashError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCashError(val *ResponseCashError) *NullableResponseCashError {
	return &NullableResponseCashError{value: val, isSet: true}
}

func (v NullableResponseCashError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCashError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ResponseCashTransferLineLimitsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCashTransferLineLimitsResponse{}

// ResponseCashTransferLineLimitsResponse struct for ResponseCashTransferLineLimitsResponse
type ResponseCashTransferLineLimitsResponse struct {
	EnteredByUserid *string `json:"entered_by_userid,omitempty"`
	EntryDate *string `json:"entry_date,omitempty"`
	LimitAmt *float32 `json:"limit_amt,omitempty"`
	LimitDescription *string `json:"limit_description,omitempty"`
	LimitId *string `json:"limit_id,omitempty"`
	UpdateDate *string `json:"update_date,omitempty"`
	UpdatedByUserid *string `json:"updated_by_userid,omitempty"`
	ValidFrom *string `json:"valid_from,omitempty"`
	ValidTo *string `json:"valid_to,omitempty"`
}

// NewResponseCashTransferLineLimitsResponse instantiates a new ResponseCashTransferLineLimitsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCashTransferLineLimitsResponse() *ResponseCashTransferLineLimitsResponse {
	this := ResponseCashTransferLineLimitsResponse{}
	return &this
}

// NewResponseCashTransferLineLimitsResponseWithDefaults instantiates a new ResponseCashTransferLineLimitsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCashTransferLineLimitsResponseWithDefaults() *ResponseCashTransferLineLimitsResponse {
	this := ResponseCashTransferLineLimitsResponse{}
	return &this
}

// GetEnteredByUserid returns the EnteredByUserid field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetEnteredByUserid() string {
	if o == nil || IsNil(o.EnteredByUserid) {
		var ret string
		return ret
	}
	return *o.EnteredByUserid
}

// GetEnteredByUseridOk returns a tuple with the EnteredByUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetEnteredByUseridOk() (*string, bool) {
	if o == nil || IsNil(o.EnteredByUserid) {
		return nil, false
	}
	return o.EnteredByUserid, true
}

// HasEnteredByUserid returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasEnteredByUserid() bool {
	if o != nil && !IsNil(o.EnteredByUserid) {
		return true
	}

	return false
}

// SetEnteredByUserid gets a reference to the given string and assigns it to the EnteredByUserid field.
func (o *ResponseCashTransferLineLimitsResponse) SetEnteredByUserid(v string) {
	o.EnteredByUserid = &v
}

// GetEntryDate returns the EntryDate field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetEntryDate() string {
	if o == nil || IsNil(o.EntryDate) {
		var ret string
		return ret
	}
	return *o.EntryDate
}

// GetEntryDateOk returns a tuple with the EntryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetEntryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EntryDate) {
		return nil, false
	}
	return o.EntryDate, true
}

// HasEntryDate returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasEntryDate() bool {
	if o != nil && !IsNil(o.EntryDate) {
		return true
	}

	return false
}

// SetEntryDate gets a reference to the given string and assigns it to the EntryDate field.
func (o *ResponseCashTransferLineLimitsResponse) SetEntryDate(v string) {
	o.EntryDate = &v
}

// GetLimitAmt returns the LimitAmt field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitAmt() float32 {
	if o == nil || IsNil(o.LimitAmt) {
		var ret float32
		return ret
	}
	return *o.LimitAmt
}

// GetLimitAmtOk returns a tuple with the LimitAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.LimitAmt) {
		return nil, false
	}
	return o.LimitAmt, true
}

// HasLimitAmt returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasLimitAmt() bool {
	if o != nil && !IsNil(o.LimitAmt) {
		return true
	}

	return false
}

// SetLimitAmt gets a reference to the given float32 and assigns it to the LimitAmt field.
func (o *ResponseCashTransferLineLimitsResponse) SetLimitAmt(v float32) {
	o.LimitAmt = &v
}

// GetLimitDescription returns the LimitDescription field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitDescription() string {
	if o == nil || IsNil(o.LimitDescription) {
		var ret string
		return ret
	}
	return *o.LimitDescription
}

// GetLimitDescriptionOk returns a tuple with the LimitDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.LimitDescription) {
		return nil, false
	}
	return o.LimitDescription, true
}

// HasLimitDescription returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasLimitDescription() bool {
	if o != nil && !IsNil(o.LimitDescription) {
		return true
	}

	return false
}

// SetLimitDescription gets a reference to the given string and assigns it to the LimitDescription field.
func (o *ResponseCashTransferLineLimitsResponse) SetLimitDescription(v string) {
	o.LimitDescription = &v
}

// GetLimitId returns the LimitId field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitId() string {
	if o == nil || IsNil(o.LimitId) {
		var ret string
		return ret
	}
	return *o.LimitId
}

// GetLimitIdOk returns a tuple with the LimitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetLimitIdOk() (*string, bool) {
	if o == nil || IsNil(o.LimitId) {
		return nil, false
	}
	return o.LimitId, true
}

// HasLimitId returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasLimitId() bool {
	if o != nil && !IsNil(o.LimitId) {
		return true
	}

	return false
}

// SetLimitId gets a reference to the given string and assigns it to the LimitId field.
func (o *ResponseCashTransferLineLimitsResponse) SetLimitId(v string) {
	o.LimitId = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetUpdateDate() string {
	if o == nil || IsNil(o.UpdateDate) {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetUpdateDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasUpdateDate() bool {
	if o != nil && !IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *ResponseCashTransferLineLimitsResponse) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetUpdatedByUserid returns the UpdatedByUserid field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetUpdatedByUserid() string {
	if o == nil || IsNil(o.UpdatedByUserid) {
		var ret string
		return ret
	}
	return *o.UpdatedByUserid
}

// GetUpdatedByUseridOk returns a tuple with the UpdatedByUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetUpdatedByUseridOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedByUserid) {
		return nil, false
	}
	return o.UpdatedByUserid, true
}

// HasUpdatedByUserid returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasUpdatedByUserid() bool {
	if o != nil && !IsNil(o.UpdatedByUserid) {
		return true
	}

	return false
}

// SetUpdatedByUserid gets a reference to the given string and assigns it to the UpdatedByUserid field.
func (o *ResponseCashTransferLineLimitsResponse) SetUpdatedByUserid(v string) {
	o.UpdatedByUserid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ResponseCashTransferLineLimitsResponse) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *ResponseCashTransferLineLimitsResponse) GetValidTo() string {
	if o == nil || IsNil(o.ValidTo) {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashTransferLineLimitsResponse) GetValidToOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTo) {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *ResponseCashTransferLineLimitsResponse) HasValidTo() bool {
	if o != nil && !IsNil(o.ValidTo) {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *ResponseCashTransferLineLimitsResponse) SetValidTo(v string) {
	o.ValidTo = &v
}

func (o ResponseCashTransferLineLimitsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCashTransferLineLimitsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnteredByUserid) {
		toSerialize["entered_by_userid"] = o.EnteredByUserid
	}
	if !IsNil(o.EntryDate) {
		toSerialize["entry_date"] = o.EntryDate
	}
	if !IsNil(o.LimitAmt) {
		toSerialize["limit_amt"] = o.LimitAmt
	}
	if !IsNil(o.LimitDescription) {
		toSerialize["limit_description"] = o.LimitDescription
	}
	if !IsNil(o.LimitId) {
		toSerialize["limit_id"] = o.LimitId
	}
	if !IsNil(o.UpdateDate) {
		toSerialize["update_date"] = o.UpdateDate
	}
	if !IsNil(o.UpdatedByUserid) {
		toSerialize["updated_by_userid"] = o.UpdatedByUserid
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if !IsNil(o.ValidTo) {
		toSerialize["valid_to"] = o.ValidTo
	}
	return toSerialize, nil
}

type NullableResponseCashTransferLineLimitsResponse struct {
	value *ResponseCashTransferLineLimitsResponse
	isSet bool
}

func (v NullableResponseCashTransferLineLimitsResponse) Get() *ResponseCashTransferLineLimitsResponse {
	return v.value
}

func (v *NullableResponseCashTransferLineLimitsResponse) Set(val *ResponseCashTransferLineLimitsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCashTransferLineLimitsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCashTransferLineLimitsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCashTransferLineLimitsResponse(val *ResponseCashTransferLineLimitsResponse) *NullableResponseCashTransferLineLimitsResponse {
	return &NullableResponseCashTransferLineLimitsResponse{value: val, isSet: true}
}

func (v NullableResponseCashTransferLineLimitsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCashTransferLineLimitsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



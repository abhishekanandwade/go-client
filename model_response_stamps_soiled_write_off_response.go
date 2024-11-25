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

// checks if the ResponseStampsSoiledWriteOffResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStampsSoiledWriteOffResponse{}

// ResponseStampsSoiledWriteOffResponse struct for ResponseStampsSoiledWriteOffResponse
type ResponseStampsSoiledWriteOffResponse struct {
	EnteredByUserid *string `json:"entered_by_userid,omitempty"`
	EntryDate *string `json:"entry_date,omitempty"`
	LimitId *string `json:"limit_id,omitempty"`
	LimitPerAnnum *string `json:"limit_per_annum,omitempty"`
	LimitPerTime *string `json:"limit_per_time,omitempty"`
	UpdateDate *string `json:"update_date,omitempty"`
	UpdatedByUserid *string `json:"updated_by_userid,omitempty"`
	ValidFrom *string `json:"valid_from,omitempty"`
	ValidTo *string `json:"valid_to,omitempty"`
}

// NewResponseStampsSoiledWriteOffResponse instantiates a new ResponseStampsSoiledWriteOffResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStampsSoiledWriteOffResponse() *ResponseStampsSoiledWriteOffResponse {
	this := ResponseStampsSoiledWriteOffResponse{}
	return &this
}

// NewResponseStampsSoiledWriteOffResponseWithDefaults instantiates a new ResponseStampsSoiledWriteOffResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStampsSoiledWriteOffResponseWithDefaults() *ResponseStampsSoiledWriteOffResponse {
	this := ResponseStampsSoiledWriteOffResponse{}
	return &this
}

// GetEnteredByUserid returns the EnteredByUserid field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetEnteredByUserid() string {
	if o == nil || IsNil(o.EnteredByUserid) {
		var ret string
		return ret
	}
	return *o.EnteredByUserid
}

// GetEnteredByUseridOk returns a tuple with the EnteredByUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetEnteredByUseridOk() (*string, bool) {
	if o == nil || IsNil(o.EnteredByUserid) {
		return nil, false
	}
	return o.EnteredByUserid, true
}

// HasEnteredByUserid returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasEnteredByUserid() bool {
	if o != nil && !IsNil(o.EnteredByUserid) {
		return true
	}

	return false
}

// SetEnteredByUserid gets a reference to the given string and assigns it to the EnteredByUserid field.
func (o *ResponseStampsSoiledWriteOffResponse) SetEnteredByUserid(v string) {
	o.EnteredByUserid = &v
}

// GetEntryDate returns the EntryDate field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetEntryDate() string {
	if o == nil || IsNil(o.EntryDate) {
		var ret string
		return ret
	}
	return *o.EntryDate
}

// GetEntryDateOk returns a tuple with the EntryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetEntryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EntryDate) {
		return nil, false
	}
	return o.EntryDate, true
}

// HasEntryDate returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasEntryDate() bool {
	if o != nil && !IsNil(o.EntryDate) {
		return true
	}

	return false
}

// SetEntryDate gets a reference to the given string and assigns it to the EntryDate field.
func (o *ResponseStampsSoiledWriteOffResponse) SetEntryDate(v string) {
	o.EntryDate = &v
}

// GetLimitId returns the LimitId field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitId() string {
	if o == nil || IsNil(o.LimitId) {
		var ret string
		return ret
	}
	return *o.LimitId
}

// GetLimitIdOk returns a tuple with the LimitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitIdOk() (*string, bool) {
	if o == nil || IsNil(o.LimitId) {
		return nil, false
	}
	return o.LimitId, true
}

// HasLimitId returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasLimitId() bool {
	if o != nil && !IsNil(o.LimitId) {
		return true
	}

	return false
}

// SetLimitId gets a reference to the given string and assigns it to the LimitId field.
func (o *ResponseStampsSoiledWriteOffResponse) SetLimitId(v string) {
	o.LimitId = &v
}

// GetLimitPerAnnum returns the LimitPerAnnum field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitPerAnnum() string {
	if o == nil || IsNil(o.LimitPerAnnum) {
		var ret string
		return ret
	}
	return *o.LimitPerAnnum
}

// GetLimitPerAnnumOk returns a tuple with the LimitPerAnnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitPerAnnumOk() (*string, bool) {
	if o == nil || IsNil(o.LimitPerAnnum) {
		return nil, false
	}
	return o.LimitPerAnnum, true
}

// HasLimitPerAnnum returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasLimitPerAnnum() bool {
	if o != nil && !IsNil(o.LimitPerAnnum) {
		return true
	}

	return false
}

// SetLimitPerAnnum gets a reference to the given string and assigns it to the LimitPerAnnum field.
func (o *ResponseStampsSoiledWriteOffResponse) SetLimitPerAnnum(v string) {
	o.LimitPerAnnum = &v
}

// GetLimitPerTime returns the LimitPerTime field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitPerTime() string {
	if o == nil || IsNil(o.LimitPerTime) {
		var ret string
		return ret
	}
	return *o.LimitPerTime
}

// GetLimitPerTimeOk returns a tuple with the LimitPerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetLimitPerTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LimitPerTime) {
		return nil, false
	}
	return o.LimitPerTime, true
}

// HasLimitPerTime returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasLimitPerTime() bool {
	if o != nil && !IsNil(o.LimitPerTime) {
		return true
	}

	return false
}

// SetLimitPerTime gets a reference to the given string and assigns it to the LimitPerTime field.
func (o *ResponseStampsSoiledWriteOffResponse) SetLimitPerTime(v string) {
	o.LimitPerTime = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetUpdateDate() string {
	if o == nil || IsNil(o.UpdateDate) {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetUpdateDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasUpdateDate() bool {
	if o != nil && !IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *ResponseStampsSoiledWriteOffResponse) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetUpdatedByUserid returns the UpdatedByUserid field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetUpdatedByUserid() string {
	if o == nil || IsNil(o.UpdatedByUserid) {
		var ret string
		return ret
	}
	return *o.UpdatedByUserid
}

// GetUpdatedByUseridOk returns a tuple with the UpdatedByUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetUpdatedByUseridOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedByUserid) {
		return nil, false
	}
	return o.UpdatedByUserid, true
}

// HasUpdatedByUserid returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasUpdatedByUserid() bool {
	if o != nil && !IsNil(o.UpdatedByUserid) {
		return true
	}

	return false
}

// SetUpdatedByUserid gets a reference to the given string and assigns it to the UpdatedByUserid field.
func (o *ResponseStampsSoiledWriteOffResponse) SetUpdatedByUserid(v string) {
	o.UpdatedByUserid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ResponseStampsSoiledWriteOffResponse) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *ResponseStampsSoiledWriteOffResponse) GetValidTo() string {
	if o == nil || IsNil(o.ValidTo) {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsSoiledWriteOffResponse) GetValidToOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTo) {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *ResponseStampsSoiledWriteOffResponse) HasValidTo() bool {
	if o != nil && !IsNil(o.ValidTo) {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *ResponseStampsSoiledWriteOffResponse) SetValidTo(v string) {
	o.ValidTo = &v
}

func (o ResponseStampsSoiledWriteOffResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStampsSoiledWriteOffResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnteredByUserid) {
		toSerialize["entered_by_userid"] = o.EnteredByUserid
	}
	if !IsNil(o.EntryDate) {
		toSerialize["entry_date"] = o.EntryDate
	}
	if !IsNil(o.LimitId) {
		toSerialize["limit_id"] = o.LimitId
	}
	if !IsNil(o.LimitPerAnnum) {
		toSerialize["limit_per_annum"] = o.LimitPerAnnum
	}
	if !IsNil(o.LimitPerTime) {
		toSerialize["limit_per_time"] = o.LimitPerTime
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

type NullableResponseStampsSoiledWriteOffResponse struct {
	value *ResponseStampsSoiledWriteOffResponse
	isSet bool
}

func (v NullableResponseStampsSoiledWriteOffResponse) Get() *ResponseStampsSoiledWriteOffResponse {
	return v.value
}

func (v *NullableResponseStampsSoiledWriteOffResponse) Set(val *ResponseStampsSoiledWriteOffResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStampsSoiledWriteOffResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStampsSoiledWriteOffResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStampsSoiledWriteOffResponse(val *ResponseStampsSoiledWriteOffResponse) *NullableResponseStampsSoiledWriteOffResponse {
	return &NullableResponseStampsSoiledWriteOffResponse{value: val, isSet: true}
}

func (v NullableResponseStampsSoiledWriteOffResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStampsSoiledWriteOffResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



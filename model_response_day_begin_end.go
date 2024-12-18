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

// checks if the ResponseDayBeginEnd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseDayBeginEnd{}

// ResponseDayBeginEnd struct for ResponseDayBeginEnd
type ResponseDayBeginEnd struct {
	DayBeginUser *string `json:"day_begin_user,omitempty"`
	DayBeginUserid *int32 `json:"day_begin_userid,omitempty"`
	DayEndTime *string `json:"day_end_time,omitempty"`
	DayEndUser *string `json:"day_end_user,omitempty"`
	DayEndUserid *int32 `json:"day_end_userid,omitempty"`
	IsAdminUnit *bool `json:"is_admin_unit,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	TransactionDate *string `json:"transaction_date,omitempty"`
}

// NewResponseDayBeginEnd instantiates a new ResponseDayBeginEnd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseDayBeginEnd() *ResponseDayBeginEnd {
	this := ResponseDayBeginEnd{}
	return &this
}

// NewResponseDayBeginEndWithDefaults instantiates a new ResponseDayBeginEnd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseDayBeginEndWithDefaults() *ResponseDayBeginEnd {
	this := ResponseDayBeginEnd{}
	return &this
}

// GetDayBeginUser returns the DayBeginUser field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetDayBeginUser() string {
	if o == nil || IsNil(o.DayBeginUser) {
		var ret string
		return ret
	}
	return *o.DayBeginUser
}

// GetDayBeginUserOk returns a tuple with the DayBeginUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetDayBeginUserOk() (*string, bool) {
	if o == nil || IsNil(o.DayBeginUser) {
		return nil, false
	}
	return o.DayBeginUser, true
}

// HasDayBeginUser returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasDayBeginUser() bool {
	if o != nil && !IsNil(o.DayBeginUser) {
		return true
	}

	return false
}

// SetDayBeginUser gets a reference to the given string and assigns it to the DayBeginUser field.
func (o *ResponseDayBeginEnd) SetDayBeginUser(v string) {
	o.DayBeginUser = &v
}

// GetDayBeginUserid returns the DayBeginUserid field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetDayBeginUserid() int32 {
	if o == nil || IsNil(o.DayBeginUserid) {
		var ret int32
		return ret
	}
	return *o.DayBeginUserid
}

// GetDayBeginUseridOk returns a tuple with the DayBeginUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetDayBeginUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.DayBeginUserid) {
		return nil, false
	}
	return o.DayBeginUserid, true
}

// HasDayBeginUserid returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasDayBeginUserid() bool {
	if o != nil && !IsNil(o.DayBeginUserid) {
		return true
	}

	return false
}

// SetDayBeginUserid gets a reference to the given int32 and assigns it to the DayBeginUserid field.
func (o *ResponseDayBeginEnd) SetDayBeginUserid(v int32) {
	o.DayBeginUserid = &v
}

// GetDayEndTime returns the DayEndTime field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetDayEndTime() string {
	if o == nil || IsNil(o.DayEndTime) {
		var ret string
		return ret
	}
	return *o.DayEndTime
}

// GetDayEndTimeOk returns a tuple with the DayEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetDayEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DayEndTime) {
		return nil, false
	}
	return o.DayEndTime, true
}

// HasDayEndTime returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasDayEndTime() bool {
	if o != nil && !IsNil(o.DayEndTime) {
		return true
	}

	return false
}

// SetDayEndTime gets a reference to the given string and assigns it to the DayEndTime field.
func (o *ResponseDayBeginEnd) SetDayEndTime(v string) {
	o.DayEndTime = &v
}

// GetDayEndUser returns the DayEndUser field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetDayEndUser() string {
	if o == nil || IsNil(o.DayEndUser) {
		var ret string
		return ret
	}
	return *o.DayEndUser
}

// GetDayEndUserOk returns a tuple with the DayEndUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetDayEndUserOk() (*string, bool) {
	if o == nil || IsNil(o.DayEndUser) {
		return nil, false
	}
	return o.DayEndUser, true
}

// HasDayEndUser returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasDayEndUser() bool {
	if o != nil && !IsNil(o.DayEndUser) {
		return true
	}

	return false
}

// SetDayEndUser gets a reference to the given string and assigns it to the DayEndUser field.
func (o *ResponseDayBeginEnd) SetDayEndUser(v string) {
	o.DayEndUser = &v
}

// GetDayEndUserid returns the DayEndUserid field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetDayEndUserid() int32 {
	if o == nil || IsNil(o.DayEndUserid) {
		var ret int32
		return ret
	}
	return *o.DayEndUserid
}

// GetDayEndUseridOk returns a tuple with the DayEndUserid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetDayEndUseridOk() (*int32, bool) {
	if o == nil || IsNil(o.DayEndUserid) {
		return nil, false
	}
	return o.DayEndUserid, true
}

// HasDayEndUserid returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasDayEndUserid() bool {
	if o != nil && !IsNil(o.DayEndUserid) {
		return true
	}

	return false
}

// SetDayEndUserid gets a reference to the given int32 and assigns it to the DayEndUserid field.
func (o *ResponseDayBeginEnd) SetDayEndUserid(v int32) {
	o.DayEndUserid = &v
}

// GetIsAdminUnit returns the IsAdminUnit field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetIsAdminUnit() bool {
	if o == nil || IsNil(o.IsAdminUnit) {
		var ret bool
		return ret
	}
	return *o.IsAdminUnit
}

// GetIsAdminUnitOk returns a tuple with the IsAdminUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetIsAdminUnitOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdminUnit) {
		return nil, false
	}
	return o.IsAdminUnit, true
}

// HasIsAdminUnit returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasIsAdminUnit() bool {
	if o != nil && !IsNil(o.IsAdminUnit) {
		return true
	}

	return false
}

// SetIsAdminUnit gets a reference to the given bool and assigns it to the IsAdminUnit field.
func (o *ResponseDayBeginEnd) SetIsAdminUnit(v bool) {
	o.IsAdminUnit = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseDayBeginEnd) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *ResponseDayBeginEnd) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseDayBeginEnd) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *ResponseDayBeginEnd) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *ResponseDayBeginEnd) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

func (o ResponseDayBeginEnd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseDayBeginEnd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayBeginUser) {
		toSerialize["day_begin_user"] = o.DayBeginUser
	}
	if !IsNil(o.DayBeginUserid) {
		toSerialize["day_begin_userid"] = o.DayBeginUserid
	}
	if !IsNil(o.DayEndTime) {
		toSerialize["day_end_time"] = o.DayEndTime
	}
	if !IsNil(o.DayEndUser) {
		toSerialize["day_end_user"] = o.DayEndUser
	}
	if !IsNil(o.DayEndUserid) {
		toSerialize["day_end_userid"] = o.DayEndUserid
	}
	if !IsNil(o.IsAdminUnit) {
		toSerialize["is_admin_unit"] = o.IsAdminUnit
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transaction_date"] = o.TransactionDate
	}
	return toSerialize, nil
}

type NullableResponseDayBeginEnd struct {
	value *ResponseDayBeginEnd
	isSet bool
}

func (v NullableResponseDayBeginEnd) Get() *ResponseDayBeginEnd {
	return v.value
}

func (v *NullableResponseDayBeginEnd) Set(val *ResponseDayBeginEnd) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseDayBeginEnd) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseDayBeginEnd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseDayBeginEnd(val *ResponseDayBeginEnd) *NullableResponseDayBeginEnd {
	return &NullableResponseDayBeginEnd{value: val, isSet: true}
}

func (v NullableResponseDayBeginEnd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseDayBeginEnd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



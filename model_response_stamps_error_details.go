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

// checks if the ResponseStampsErrorDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStampsErrorDetails{}

// ResponseStampsErrorDetails struct for ResponseStampsErrorDetails
type ResponseStampsErrorDetails struct {
	ErrorId *string `json:"error_id,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	OfficeName *string `json:"office_name,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// NewResponseStampsErrorDetails instantiates a new ResponseStampsErrorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStampsErrorDetails() *ResponseStampsErrorDetails {
	this := ResponseStampsErrorDetails{}
	return &this
}

// NewResponseStampsErrorDetailsWithDefaults instantiates a new ResponseStampsErrorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStampsErrorDetailsWithDefaults() *ResponseStampsErrorDetails {
	this := ResponseStampsErrorDetails{}
	return &this
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *ResponseStampsErrorDetails) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseStampsErrorDetails) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOfficeName returns the OfficeName field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetOfficeName() string {
	if o == nil || IsNil(o.OfficeName) {
		var ret string
		return ret
	}
	return *o.OfficeName
}

// GetOfficeNameOk returns a tuple with the OfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.OfficeName) {
		return nil, false
	}
	return o.OfficeName, true
}

// HasOfficeName returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasOfficeName() bool {
	if o != nil && !IsNil(o.OfficeName) {
		return true
	}

	return false
}

// SetOfficeName gets a reference to the given string and assigns it to the OfficeName field.
func (o *ResponseStampsErrorDetails) SetOfficeName(v string) {
	o.OfficeName = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseStampsErrorDetails) SetRemarks(v string) {
	o.Remarks = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseStampsErrorDetails) SetTransDate(v string) {
	o.TransDate = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ResponseStampsErrorDetails) SetUserId(v int32) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ResponseStampsErrorDetails) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsErrorDetails) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ResponseStampsErrorDetails) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ResponseStampsErrorDetails) SetUserName(v string) {
	o.UserName = &v
}

func (o ResponseStampsErrorDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStampsErrorDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorId) {
		toSerialize["error_id"] = o.ErrorId
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.OfficeName) {
		toSerialize["office_name"] = o.OfficeName
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	return toSerialize, nil
}

type NullableResponseStampsErrorDetails struct {
	value *ResponseStampsErrorDetails
	isSet bool
}

func (v NullableResponseStampsErrorDetails) Get() *ResponseStampsErrorDetails {
	return v.value
}

func (v *NullableResponseStampsErrorDetails) Set(val *ResponseStampsErrorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStampsErrorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStampsErrorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStampsErrorDetails(val *ResponseStampsErrorDetails) *NullableResponseStampsErrorDetails {
	return &NullableResponseStampsErrorDetails{value: val, isSet: true}
}

func (v NullableResponseStampsErrorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStampsErrorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



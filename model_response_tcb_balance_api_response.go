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

// checks if the ResponseTcbBalanceAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTcbBalanceAPIResponse{}

// ResponseTcbBalanceAPIResponse struct for ResponseTcbBalanceAPIResponse
type ResponseTcbBalanceAPIResponse struct {
	Data *ResponseTcbBalance `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewResponseTcbBalanceAPIResponse instantiates a new ResponseTcbBalanceAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTcbBalanceAPIResponse() *ResponseTcbBalanceAPIResponse {
	this := ResponseTcbBalanceAPIResponse{}
	return &this
}

// NewResponseTcbBalanceAPIResponseWithDefaults instantiates a new ResponseTcbBalanceAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTcbBalanceAPIResponseWithDefaults() *ResponseTcbBalanceAPIResponse {
	this := ResponseTcbBalanceAPIResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseTcbBalanceAPIResponse) GetData() ResponseTcbBalance {
	if o == nil || IsNil(o.Data) {
		var ret ResponseTcbBalance
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTcbBalanceAPIResponse) GetDataOk() (*ResponseTcbBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseTcbBalanceAPIResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseTcbBalance and assigns it to the Data field.
func (o *ResponseTcbBalanceAPIResponse) SetData(v ResponseTcbBalance) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseTcbBalanceAPIResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTcbBalanceAPIResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseTcbBalanceAPIResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseTcbBalanceAPIResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseTcbBalanceAPIResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTcbBalanceAPIResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseTcbBalanceAPIResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseTcbBalanceAPIResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ResponseTcbBalanceAPIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTcbBalanceAPIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableResponseTcbBalanceAPIResponse struct {
	value *ResponseTcbBalanceAPIResponse
	isSet bool
}

func (v NullableResponseTcbBalanceAPIResponse) Get() *ResponseTcbBalanceAPIResponse {
	return v.value
}

func (v *NullableResponseTcbBalanceAPIResponse) Set(val *ResponseTcbBalanceAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTcbBalanceAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTcbBalanceAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTcbBalanceAPIResponse(val *ResponseTcbBalanceAPIResponse) *NullableResponseTcbBalanceAPIResponse {
	return &NullableResponseTcbBalanceAPIResponse{value: val, isSet: true}
}

func (v NullableResponseTcbBalanceAPIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTcbBalanceAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



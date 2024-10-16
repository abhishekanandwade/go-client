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

// checks if the ResponseListReportsTCBDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListReportsTCBDataResponse{}

// ResponseListReportsTCBDataResponse struct for ResponseListReportsTCBDataResponse
type ResponseListReportsTCBDataResponse struct {
	Data *ResponseTcbBalance `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewResponseListReportsTCBDataResponse instantiates a new ResponseListReportsTCBDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListReportsTCBDataResponse() *ResponseListReportsTCBDataResponse {
	this := ResponseListReportsTCBDataResponse{}
	return &this
}

// NewResponseListReportsTCBDataResponseWithDefaults instantiates a new ResponseListReportsTCBDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListReportsTCBDataResponseWithDefaults() *ResponseListReportsTCBDataResponse {
	this := ResponseListReportsTCBDataResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseListReportsTCBDataResponse) GetData() ResponseTcbBalance {
	if o == nil || IsNil(o.Data) {
		var ret ResponseTcbBalance
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListReportsTCBDataResponse) GetDataOk() (*ResponseTcbBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseListReportsTCBDataResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseTcbBalance and assigns it to the Data field.
func (o *ResponseListReportsTCBDataResponse) SetData(v ResponseTcbBalance) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseListReportsTCBDataResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListReportsTCBDataResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseListReportsTCBDataResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseListReportsTCBDataResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseListReportsTCBDataResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListReportsTCBDataResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseListReportsTCBDataResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseListReportsTCBDataResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ResponseListReportsTCBDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListReportsTCBDataResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseListReportsTCBDataResponse struct {
	value *ResponseListReportsTCBDataResponse
	isSet bool
}

func (v NullableResponseListReportsTCBDataResponse) Get() *ResponseListReportsTCBDataResponse {
	return v.value
}

func (v *NullableResponseListReportsTCBDataResponse) Set(val *ResponseListReportsTCBDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListReportsTCBDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListReportsTCBDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListReportsTCBDataResponse(val *ResponseListReportsTCBDataResponse) *NullableResponseListReportsTCBDataResponse {
	return &NullableResponseListReportsTCBDataResponse{value: val, isSet: true}
}

func (v NullableResponseListReportsTCBDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListReportsTCBDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



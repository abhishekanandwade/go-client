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

// checks if the ResponseCashInTransitAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCashInTransitAPIResponse{}

// ResponseCashInTransitAPIResponse struct for ResponseCashInTransitAPIResponse
type ResponseCashInTransitAPIResponse struct {
	Data []ResponseCashInTransit `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewResponseCashInTransitAPIResponse instantiates a new ResponseCashInTransitAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCashInTransitAPIResponse() *ResponseCashInTransitAPIResponse {
	this := ResponseCashInTransitAPIResponse{}
	return &this
}

// NewResponseCashInTransitAPIResponseWithDefaults instantiates a new ResponseCashInTransitAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCashInTransitAPIResponseWithDefaults() *ResponseCashInTransitAPIResponse {
	this := ResponseCashInTransitAPIResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseCashInTransitAPIResponse) GetData() []ResponseCashInTransit {
	if o == nil || IsNil(o.Data) {
		var ret []ResponseCashInTransit
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransitAPIResponse) GetDataOk() ([]ResponseCashInTransit, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseCashInTransitAPIResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResponseCashInTransit and assigns it to the Data field.
func (o *ResponseCashInTransitAPIResponse) SetData(v []ResponseCashInTransit) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseCashInTransitAPIResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransitAPIResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseCashInTransitAPIResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseCashInTransitAPIResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseCashInTransitAPIResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransitAPIResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseCashInTransitAPIResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseCashInTransitAPIResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ResponseCashInTransitAPIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCashInTransitAPIResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseCashInTransitAPIResponse struct {
	value *ResponseCashInTransitAPIResponse
	isSet bool
}

func (v NullableResponseCashInTransitAPIResponse) Get() *ResponseCashInTransitAPIResponse {
	return v.value
}

func (v *NullableResponseCashInTransitAPIResponse) Set(val *ResponseCashInTransitAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCashInTransitAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCashInTransitAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCashInTransitAPIResponse(val *ResponseCashInTransitAPIResponse) *NullableResponseCashInTransitAPIResponse {
	return &NullableResponseCashInTransitAPIResponse{value: val, isSet: true}
}

func (v NullableResponseCashInTransitAPIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCashInTransitAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



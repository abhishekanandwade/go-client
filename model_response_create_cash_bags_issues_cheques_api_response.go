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

// checks if the ResponseCreateCashBagsIssuesChequesApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCreateCashBagsIssuesChequesApiResponse{}

// ResponseCreateCashBagsIssuesChequesApiResponse struct for ResponseCreateCashBagsIssuesChequesApiResponse
type ResponseCreateCashBagsIssuesChequesApiResponse struct {
	Data *ResponseCashBagTransactions `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewResponseCreateCashBagsIssuesChequesApiResponse instantiates a new ResponseCreateCashBagsIssuesChequesApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCreateCashBagsIssuesChequesApiResponse() *ResponseCreateCashBagsIssuesChequesApiResponse {
	this := ResponseCreateCashBagsIssuesChequesApiResponse{}
	return &this
}

// NewResponseCreateCashBagsIssuesChequesApiResponseWithDefaults instantiates a new ResponseCreateCashBagsIssuesChequesApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCreateCashBagsIssuesChequesApiResponseWithDefaults() *ResponseCreateCashBagsIssuesChequesApiResponse {
	this := ResponseCreateCashBagsIssuesChequesApiResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetData() ResponseCashBagTransactions {
	if o == nil || IsNil(o.Data) {
		var ret ResponseCashBagTransactions
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetDataOk() (*ResponseCashBagTransactions, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseCashBagTransactions and assigns it to the Data field.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) SetData(v ResponseCashBagTransactions) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseCreateCashBagsIssuesChequesApiResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ResponseCreateCashBagsIssuesChequesApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCreateCashBagsIssuesChequesApiResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseCreateCashBagsIssuesChequesApiResponse struct {
	value *ResponseCreateCashBagsIssuesChequesApiResponse
	isSet bool
}

func (v NullableResponseCreateCashBagsIssuesChequesApiResponse) Get() *ResponseCreateCashBagsIssuesChequesApiResponse {
	return v.value
}

func (v *NullableResponseCreateCashBagsIssuesChequesApiResponse) Set(val *ResponseCreateCashBagsIssuesChequesApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCreateCashBagsIssuesChequesApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCreateCashBagsIssuesChequesApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCreateCashBagsIssuesChequesApiResponse(val *ResponseCreateCashBagsIssuesChequesApiResponse) *NullableResponseCreateCashBagsIssuesChequesApiResponse {
	return &NullableResponseCreateCashBagsIssuesChequesApiResponse{value: val, isSet: true}
}

func (v NullableResponseCreateCashBagsIssuesChequesApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCreateCashBagsIssuesChequesApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ResponseStampObCsdReceiptsApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStampObCsdReceiptsApiResponse{}

// ResponseStampObCsdReceiptsApiResponse struct for ResponseStampObCsdReceiptsApiResponse
type ResponseStampObCsdReceiptsApiResponse struct {
	Data *ResponseStampsObReceipts `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewResponseStampObCsdReceiptsApiResponse instantiates a new ResponseStampObCsdReceiptsApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStampObCsdReceiptsApiResponse() *ResponseStampObCsdReceiptsApiResponse {
	this := ResponseStampObCsdReceiptsApiResponse{}
	return &this
}

// NewResponseStampObCsdReceiptsApiResponseWithDefaults instantiates a new ResponseStampObCsdReceiptsApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStampObCsdReceiptsApiResponseWithDefaults() *ResponseStampObCsdReceiptsApiResponse {
	this := ResponseStampObCsdReceiptsApiResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseStampObCsdReceiptsApiResponse) GetData() ResponseStampsObReceipts {
	if o == nil || IsNil(o.Data) {
		var ret ResponseStampsObReceipts
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) GetDataOk() (*ResponseStampsObReceipts, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseStampsObReceipts and assigns it to the Data field.
func (o *ResponseStampObCsdReceiptsApiResponse) SetData(v ResponseStampsObReceipts) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseStampObCsdReceiptsApiResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseStampObCsdReceiptsApiResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseStampObCsdReceiptsApiResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseStampObCsdReceiptsApiResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseStampObCsdReceiptsApiResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ResponseStampObCsdReceiptsApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStampObCsdReceiptsApiResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseStampObCsdReceiptsApiResponse struct {
	value *ResponseStampObCsdReceiptsApiResponse
	isSet bool
}

func (v NullableResponseStampObCsdReceiptsApiResponse) Get() *ResponseStampObCsdReceiptsApiResponse {
	return v.value
}

func (v *NullableResponseStampObCsdReceiptsApiResponse) Set(val *ResponseStampObCsdReceiptsApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStampObCsdReceiptsApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStampObCsdReceiptsApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStampObCsdReceiptsApiResponse(val *ResponseStampObCsdReceiptsApiResponse) *NullableResponseStampObCsdReceiptsApiResponse {
	return &NullableResponseStampObCsdReceiptsApiResponse{value: val, isSet: true}
}

func (v NullableResponseStampObCsdReceiptsApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStampObCsdReceiptsApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



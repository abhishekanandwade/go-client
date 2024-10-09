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

// checks if the ResponseListCashErrorsApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListCashErrorsApiResponse{}

// ResponseListCashErrorsApiResponse struct for ResponseListCashErrorsApiResponse
type ResponseListCashErrorsApiResponse struct {
	Data []ResponseCashError `json:"data,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Message *string `json:"message,omitempty"`
	OrderBy *string `json:"order_by,omitempty"`
	ReturnedRecordsCount *int32 `json:"returned_records_count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	SortType *string `json:"sort_type,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	TotalRecordsCount *int32 `json:"total_records_count,omitempty"`
}

// NewResponseListCashErrorsApiResponse instantiates a new ResponseListCashErrorsApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListCashErrorsApiResponse() *ResponseListCashErrorsApiResponse {
	this := ResponseListCashErrorsApiResponse{}
	return &this
}

// NewResponseListCashErrorsApiResponseWithDefaults instantiates a new ResponseListCashErrorsApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListCashErrorsApiResponseWithDefaults() *ResponseListCashErrorsApiResponse {
	this := ResponseListCashErrorsApiResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetData() []ResponseCashError {
	if o == nil || IsNil(o.Data) {
		var ret []ResponseCashError
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetDataOk() ([]ResponseCashError, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResponseCashError and assigns it to the Data field.
func (o *ResponseListCashErrorsApiResponse) SetData(v []ResponseCashError) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseListCashErrorsApiResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseListCashErrorsApiResponse) SetMessage(v string) {
	o.Message = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *ResponseListCashErrorsApiResponse) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetReturnedRecordsCount returns the ReturnedRecordsCount field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetReturnedRecordsCount() int32 {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		var ret int32
		return ret
	}
	return *o.ReturnedRecordsCount
}

// GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetReturnedRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		return nil, false
	}
	return o.ReturnedRecordsCount, true
}

// HasReturnedRecordsCount returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasReturnedRecordsCount() bool {
	if o != nil && !IsNil(o.ReturnedRecordsCount) {
		return true
	}

	return false
}

// SetReturnedRecordsCount gets a reference to the given int32 and assigns it to the ReturnedRecordsCount field.
func (o *ResponseListCashErrorsApiResponse) SetReturnedRecordsCount(v int32) {
	o.ReturnedRecordsCount = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ResponseListCashErrorsApiResponse) SetSkip(v int32) {
	o.Skip = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *ResponseListCashErrorsApiResponse) SetSortType(v string) {
	o.SortType = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseListCashErrorsApiResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetTotalRecordsCount returns the TotalRecordsCount field value if set, zero value otherwise.
func (o *ResponseListCashErrorsApiResponse) GetTotalRecordsCount() int32 {
	if o == nil || IsNil(o.TotalRecordsCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordsCount
}

// GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListCashErrorsApiResponse) GetTotalRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordsCount) {
		return nil, false
	}
	return o.TotalRecordsCount, true
}

// HasTotalRecordsCount returns a boolean if a field has been set.
func (o *ResponseListCashErrorsApiResponse) HasTotalRecordsCount() bool {
	if o != nil && !IsNil(o.TotalRecordsCount) {
		return true
	}

	return false
}

// SetTotalRecordsCount gets a reference to the given int32 and assigns it to the TotalRecordsCount field.
func (o *ResponseListCashErrorsApiResponse) SetTotalRecordsCount(v int32) {
	o.TotalRecordsCount = &v
}

func (o ResponseListCashErrorsApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListCashErrorsApiResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.OrderBy) {
		toSerialize["order_by"] = o.OrderBy
	}
	if !IsNil(o.ReturnedRecordsCount) {
		toSerialize["returned_records_count"] = o.ReturnedRecordsCount
	}
	if !IsNil(o.Skip) {
		toSerialize["skip"] = o.Skip
	}
	if !IsNil(o.SortType) {
		toSerialize["sort_type"] = o.SortType
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if !IsNil(o.TotalRecordsCount) {
		toSerialize["total_records_count"] = o.TotalRecordsCount
	}
	return toSerialize, nil
}

type NullableResponseListCashErrorsApiResponse struct {
	value *ResponseListCashErrorsApiResponse
	isSet bool
}

func (v NullableResponseListCashErrorsApiResponse) Get() *ResponseListCashErrorsApiResponse {
	return v.value
}

func (v *NullableResponseListCashErrorsApiResponse) Set(val *ResponseListCashErrorsApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListCashErrorsApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListCashErrorsApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListCashErrorsApiResponse(val *ResponseListCashErrorsApiResponse) *NullableResponseListCashErrorsApiResponse {
	return &NullableResponseListCashErrorsApiResponse{value: val, isSet: true}
}

func (v NullableResponseListCashErrorsApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListCashErrorsApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



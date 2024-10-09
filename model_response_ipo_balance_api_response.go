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

// checks if the ResponseIPOBalanceAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIPOBalanceAPIResponse{}

// ResponseIPOBalanceAPIResponse struct for ResponseIPOBalanceAPIResponse
type ResponseIPOBalanceAPIResponse struct {
	Data []ResponseIPOBalance `json:"data,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Message *string `json:"message,omitempty"`
	OrderBy *string `json:"order_by,omitempty"`
	ReturnedRecordsCount *int32 `json:"returned_records_count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	SortType *string `json:"sort_type,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	TotalRecordsCount *int32 `json:"total_records_count,omitempty"`
}

// NewResponseIPOBalanceAPIResponse instantiates a new ResponseIPOBalanceAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIPOBalanceAPIResponse() *ResponseIPOBalanceAPIResponse {
	this := ResponseIPOBalanceAPIResponse{}
	return &this
}

// NewResponseIPOBalanceAPIResponseWithDefaults instantiates a new ResponseIPOBalanceAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIPOBalanceAPIResponseWithDefaults() *ResponseIPOBalanceAPIResponse {
	this := ResponseIPOBalanceAPIResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetData() []ResponseIPOBalance {
	if o == nil || IsNil(o.Data) {
		var ret []ResponseIPOBalance
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetDataOk() ([]ResponseIPOBalance, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResponseIPOBalance and assigns it to the Data field.
func (o *ResponseIPOBalanceAPIResponse) SetData(v []ResponseIPOBalance) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseIPOBalanceAPIResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseIPOBalanceAPIResponse) SetMessage(v string) {
	o.Message = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *ResponseIPOBalanceAPIResponse) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetReturnedRecordsCount returns the ReturnedRecordsCount field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetReturnedRecordsCount() int32 {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		var ret int32
		return ret
	}
	return *o.ReturnedRecordsCount
}

// GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetReturnedRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		return nil, false
	}
	return o.ReturnedRecordsCount, true
}

// HasReturnedRecordsCount returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasReturnedRecordsCount() bool {
	if o != nil && !IsNil(o.ReturnedRecordsCount) {
		return true
	}

	return false
}

// SetReturnedRecordsCount gets a reference to the given int32 and assigns it to the ReturnedRecordsCount field.
func (o *ResponseIPOBalanceAPIResponse) SetReturnedRecordsCount(v int32) {
	o.ReturnedRecordsCount = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ResponseIPOBalanceAPIResponse) SetSkip(v int32) {
	o.Skip = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *ResponseIPOBalanceAPIResponse) SetSortType(v string) {
	o.SortType = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseIPOBalanceAPIResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetTotalRecordsCount returns the TotalRecordsCount field value if set, zero value otherwise.
func (o *ResponseIPOBalanceAPIResponse) GetTotalRecordsCount() int32 {
	if o == nil || IsNil(o.TotalRecordsCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordsCount
}

// GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOBalanceAPIResponse) GetTotalRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordsCount) {
		return nil, false
	}
	return o.TotalRecordsCount, true
}

// HasTotalRecordsCount returns a boolean if a field has been set.
func (o *ResponseIPOBalanceAPIResponse) HasTotalRecordsCount() bool {
	if o != nil && !IsNil(o.TotalRecordsCount) {
		return true
	}

	return false
}

// SetTotalRecordsCount gets a reference to the given int32 and assigns it to the TotalRecordsCount field.
func (o *ResponseIPOBalanceAPIResponse) SetTotalRecordsCount(v int32) {
	o.TotalRecordsCount = &v
}

func (o ResponseIPOBalanceAPIResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIPOBalanceAPIResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseIPOBalanceAPIResponse struct {
	value *ResponseIPOBalanceAPIResponse
	isSet bool
}

func (v NullableResponseIPOBalanceAPIResponse) Get() *ResponseIPOBalanceAPIResponse {
	return v.value
}

func (v *NullableResponseIPOBalanceAPIResponse) Set(val *ResponseIPOBalanceAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIPOBalanceAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIPOBalanceAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIPOBalanceAPIResponse(val *ResponseIPOBalanceAPIResponse) *NullableResponseIPOBalanceAPIResponse {
	return &NullableResponseIPOBalanceAPIResponse{value: val, isSet: true}
}

func (v NullableResponseIPOBalanceAPIResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIPOBalanceAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



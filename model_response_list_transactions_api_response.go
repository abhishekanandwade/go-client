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

// checks if the ResponseListTransactionsApiResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseListTransactionsApiResponse{}

// ResponseListTransactionsApiResponse struct for ResponseListTransactionsApiResponse
type ResponseListTransactionsApiResponse struct {
	Data []ResponseTreasuryTransactions `json:"data,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Message *string `json:"message,omitempty"`
	OrderBy *string `json:"order_by,omitempty"`
	ReturnedRecordsCount *int32 `json:"returned_records_count,omitempty"`
	Skip *int32 `json:"skip,omitempty"`
	SortType *string `json:"sort_type,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	TotalRecordsCount *int32 `json:"total_records_count,omitempty"`
}

// NewResponseListTransactionsApiResponse instantiates a new ResponseListTransactionsApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseListTransactionsApiResponse() *ResponseListTransactionsApiResponse {
	this := ResponseListTransactionsApiResponse{}
	return &this
}

// NewResponseListTransactionsApiResponseWithDefaults instantiates a new ResponseListTransactionsApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseListTransactionsApiResponseWithDefaults() *ResponseListTransactionsApiResponse {
	this := ResponseListTransactionsApiResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetData() []ResponseTreasuryTransactions {
	if o == nil || IsNil(o.Data) {
		var ret []ResponseTreasuryTransactions
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetDataOk() ([]ResponseTreasuryTransactions, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ResponseTreasuryTransactions and assigns it to the Data field.
func (o *ResponseListTransactionsApiResponse) SetData(v []ResponseTreasuryTransactions) {
	o.Data = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ResponseListTransactionsApiResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ResponseListTransactionsApiResponse) SetMessage(v string) {
	o.Message = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *ResponseListTransactionsApiResponse) SetOrderBy(v string) {
	o.OrderBy = &v
}

// GetReturnedRecordsCount returns the ReturnedRecordsCount field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetReturnedRecordsCount() int32 {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		var ret int32
		return ret
	}
	return *o.ReturnedRecordsCount
}

// GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetReturnedRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReturnedRecordsCount) {
		return nil, false
	}
	return o.ReturnedRecordsCount, true
}

// HasReturnedRecordsCount returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasReturnedRecordsCount() bool {
	if o != nil && !IsNil(o.ReturnedRecordsCount) {
		return true
	}

	return false
}

// SetReturnedRecordsCount gets a reference to the given int32 and assigns it to the ReturnedRecordsCount field.
func (o *ResponseListTransactionsApiResponse) SetReturnedRecordsCount(v int32) {
	o.ReturnedRecordsCount = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetSkip() int32 {
	if o == nil || IsNil(o.Skip) {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetSkipOk() (*int32, bool) {
	if o == nil || IsNil(o.Skip) {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasSkip() bool {
	if o != nil && !IsNil(o.Skip) {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *ResponseListTransactionsApiResponse) SetSkip(v int32) {
	o.Skip = &v
}

// GetSortType returns the SortType field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetSortType() string {
	if o == nil || IsNil(o.SortType) {
		var ret string
		return ret
	}
	return *o.SortType
}

// GetSortTypeOk returns a tuple with the SortType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetSortTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SortType) {
		return nil, false
	}
	return o.SortType, true
}

// HasSortType returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasSortType() bool {
	if o != nil && !IsNil(o.SortType) {
		return true
	}

	return false
}

// SetSortType gets a reference to the given string and assigns it to the SortType field.
func (o *ResponseListTransactionsApiResponse) SetSortType(v string) {
	o.SortType = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ResponseListTransactionsApiResponse) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetTotalRecordsCount returns the TotalRecordsCount field value if set, zero value otherwise.
func (o *ResponseListTransactionsApiResponse) GetTotalRecordsCount() int32 {
	if o == nil || IsNil(o.TotalRecordsCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordsCount
}

// GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseListTransactionsApiResponse) GetTotalRecordsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordsCount) {
		return nil, false
	}
	return o.TotalRecordsCount, true
}

// HasTotalRecordsCount returns a boolean if a field has been set.
func (o *ResponseListTransactionsApiResponse) HasTotalRecordsCount() bool {
	if o != nil && !IsNil(o.TotalRecordsCount) {
		return true
	}

	return false
}

// SetTotalRecordsCount gets a reference to the given int32 and assigns it to the TotalRecordsCount field.
func (o *ResponseListTransactionsApiResponse) SetTotalRecordsCount(v int32) {
	o.TotalRecordsCount = &v
}

func (o ResponseListTransactionsApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseListTransactionsApiResponse) ToMap() (map[string]interface{}, error) {
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

type NullableResponseListTransactionsApiResponse struct {
	value *ResponseListTransactionsApiResponse
	isSet bool
}

func (v NullableResponseListTransactionsApiResponse) Get() *ResponseListTransactionsApiResponse {
	return v.value
}

func (v *NullableResponseListTransactionsApiResponse) Set(val *ResponseListTransactionsApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseListTransactionsApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseListTransactionsApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseListTransactionsApiResponse(val *ResponseListTransactionsApiResponse) *NullableResponseListTransactionsApiResponse {
	return &NullableResponseListTransactionsApiResponse{value: val, isSet: true}
}

func (v NullableResponseListTransactionsApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseListTransactionsApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



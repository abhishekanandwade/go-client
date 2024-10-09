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

// checks if the ResponseCashInTransit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCashInTransit{}

// ResponseCashInTransit struct for ResponseCashInTransit
type ResponseCashInTransit struct {
	AmountSent *float32 `json:"amount_sent,omitempty"`
	FromOffice *string `json:"from_office,omitempty"`
	FromUnitId *int32 `json:"from_unit_id,omitempty"`
	IsRsao *bool `json:"is_rsao,omitempty"`
	ToOffice *string `json:"to_office,omitempty"`
	ToUnitId *int32 `json:"to_unit_id,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransId *string `json:"trans_id,omitempty"`
	TransactionDate *string `json:"transaction_date,omitempty"`
}

// NewResponseCashInTransit instantiates a new ResponseCashInTransit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCashInTransit() *ResponseCashInTransit {
	this := ResponseCashInTransit{}
	return &this
}

// NewResponseCashInTransitWithDefaults instantiates a new ResponseCashInTransit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCashInTransitWithDefaults() *ResponseCashInTransit {
	this := ResponseCashInTransit{}
	return &this
}

// GetAmountSent returns the AmountSent field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetAmountSent() float32 {
	if o == nil || IsNil(o.AmountSent) {
		var ret float32
		return ret
	}
	return *o.AmountSent
}

// GetAmountSentOk returns a tuple with the AmountSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetAmountSentOk() (*float32, bool) {
	if o == nil || IsNil(o.AmountSent) {
		return nil, false
	}
	return o.AmountSent, true
}

// HasAmountSent returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasAmountSent() bool {
	if o != nil && !IsNil(o.AmountSent) {
		return true
	}

	return false
}

// SetAmountSent gets a reference to the given float32 and assigns it to the AmountSent field.
func (o *ResponseCashInTransit) SetAmountSent(v float32) {
	o.AmountSent = &v
}

// GetFromOffice returns the FromOffice field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetFromOffice() string {
	if o == nil || IsNil(o.FromOffice) {
		var ret string
		return ret
	}
	return *o.FromOffice
}

// GetFromOfficeOk returns a tuple with the FromOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetFromOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.FromOffice) {
		return nil, false
	}
	return o.FromOffice, true
}

// HasFromOffice returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasFromOffice() bool {
	if o != nil && !IsNil(o.FromOffice) {
		return true
	}

	return false
}

// SetFromOffice gets a reference to the given string and assigns it to the FromOffice field.
func (o *ResponseCashInTransit) SetFromOffice(v string) {
	o.FromOffice = &v
}

// GetFromUnitId returns the FromUnitId field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetFromUnitId() int32 {
	if o == nil || IsNil(o.FromUnitId) {
		var ret int32
		return ret
	}
	return *o.FromUnitId
}

// GetFromUnitIdOk returns a tuple with the FromUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetFromUnitIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FromUnitId) {
		return nil, false
	}
	return o.FromUnitId, true
}

// HasFromUnitId returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasFromUnitId() bool {
	if o != nil && !IsNil(o.FromUnitId) {
		return true
	}

	return false
}

// SetFromUnitId gets a reference to the given int32 and assigns it to the FromUnitId field.
func (o *ResponseCashInTransit) SetFromUnitId(v int32) {
	o.FromUnitId = &v
}

// GetIsRsao returns the IsRsao field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetIsRsao() bool {
	if o == nil || IsNil(o.IsRsao) {
		var ret bool
		return ret
	}
	return *o.IsRsao
}

// GetIsRsaoOk returns a tuple with the IsRsao field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetIsRsaoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRsao) {
		return nil, false
	}
	return o.IsRsao, true
}

// HasIsRsao returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasIsRsao() bool {
	if o != nil && !IsNil(o.IsRsao) {
		return true
	}

	return false
}

// SetIsRsao gets a reference to the given bool and assigns it to the IsRsao field.
func (o *ResponseCashInTransit) SetIsRsao(v bool) {
	o.IsRsao = &v
}

// GetToOffice returns the ToOffice field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetToOffice() string {
	if o == nil || IsNil(o.ToOffice) {
		var ret string
		return ret
	}
	return *o.ToOffice
}

// GetToOfficeOk returns a tuple with the ToOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetToOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.ToOffice) {
		return nil, false
	}
	return o.ToOffice, true
}

// HasToOffice returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasToOffice() bool {
	if o != nil && !IsNil(o.ToOffice) {
		return true
	}

	return false
}

// SetToOffice gets a reference to the given string and assigns it to the ToOffice field.
func (o *ResponseCashInTransit) SetToOffice(v string) {
	o.ToOffice = &v
}

// GetToUnitId returns the ToUnitId field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetToUnitId() int32 {
	if o == nil || IsNil(o.ToUnitId) {
		var ret int32
		return ret
	}
	return *o.ToUnitId
}

// GetToUnitIdOk returns a tuple with the ToUnitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetToUnitIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ToUnitId) {
		return nil, false
	}
	return o.ToUnitId, true
}

// HasToUnitId returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasToUnitId() bool {
	if o != nil && !IsNil(o.ToUnitId) {
		return true
	}

	return false
}

// SetToUnitId gets a reference to the given int32 and assigns it to the ToUnitId field.
func (o *ResponseCashInTransit) SetToUnitId(v int32) {
	o.ToUnitId = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseCashInTransit) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransId returns the TransId field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetTransId() string {
	if o == nil || IsNil(o.TransId) {
		var ret string
		return ret
	}
	return *o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetTransIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransId) {
		return nil, false
	}
	return o.TransId, true
}

// HasTransId returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasTransId() bool {
	if o != nil && !IsNil(o.TransId) {
		return true
	}

	return false
}

// SetTransId gets a reference to the given string and assigns it to the TransId field.
func (o *ResponseCashInTransit) SetTransId(v string) {
	o.TransId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *ResponseCashInTransit) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCashInTransit) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *ResponseCashInTransit) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *ResponseCashInTransit) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

func (o ResponseCashInTransit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCashInTransit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountSent) {
		toSerialize["amount_sent"] = o.AmountSent
	}
	if !IsNil(o.FromOffice) {
		toSerialize["from_office"] = o.FromOffice
	}
	if !IsNil(o.FromUnitId) {
		toSerialize["from_unit_id"] = o.FromUnitId
	}
	if !IsNil(o.IsRsao) {
		toSerialize["is_rsao"] = o.IsRsao
	}
	if !IsNil(o.ToOffice) {
		toSerialize["to_office"] = o.ToOffice
	}
	if !IsNil(o.ToUnitId) {
		toSerialize["to_unit_id"] = o.ToUnitId
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransId) {
		toSerialize["trans_id"] = o.TransId
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transaction_date"] = o.TransactionDate
	}
	return toSerialize, nil
}

type NullableResponseCashInTransit struct {
	value *ResponseCashInTransit
	isSet bool
}

func (v NullableResponseCashInTransit) Get() *ResponseCashInTransit {
	return v.value
}

func (v *NullableResponseCashInTransit) Set(val *ResponseCashInTransit) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCashInTransit) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCashInTransit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCashInTransit(val *ResponseCashInTransit) *NullableResponseCashInTransit {
	return &NullableResponseCashInTransit{value: val, isSet: true}
}

func (v NullableResponseCashInTransit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCashInTransit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


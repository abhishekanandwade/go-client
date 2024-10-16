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

// checks if the ResponseCurrenciesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCurrenciesResponse{}

// ResponseCurrenciesResponse struct for ResponseCurrenciesResponse
type ResponseCurrenciesResponse struct {
	CurrencyId *string `json:"currency_id,omitempty"`
	CurrencyType *string `json:"currency_type,omitempty"`
	Denomination *float32 `json:"denomination,omitempty"`
	EnteredByUser *string `json:"entered_by_user,omitempty"`
	EntryDate *string `json:"entry_date,omitempty"`
	UpdateDate *string `json:"update_date,omitempty"`
	UpdatedByUser *string `json:"updated_by_user,omitempty"`
	ValidFrom *string `json:"valid_from,omitempty"`
	ValidTo *string `json:"valid_to,omitempty"`
}

// NewResponseCurrenciesResponse instantiates a new ResponseCurrenciesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCurrenciesResponse() *ResponseCurrenciesResponse {
	this := ResponseCurrenciesResponse{}
	return &this
}

// NewResponseCurrenciesResponseWithDefaults instantiates a new ResponseCurrenciesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCurrenciesResponseWithDefaults() *ResponseCurrenciesResponse {
	this := ResponseCurrenciesResponse{}
	return &this
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *ResponseCurrenciesResponse) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetCurrencyType returns the CurrencyType field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetCurrencyType() string {
	if o == nil || IsNil(o.CurrencyType) {
		var ret string
		return ret
	}
	return *o.CurrencyType
}

// GetCurrencyTypeOk returns a tuple with the CurrencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetCurrencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyType) {
		return nil, false
	}
	return o.CurrencyType, true
}

// HasCurrencyType returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasCurrencyType() bool {
	if o != nil && !IsNil(o.CurrencyType) {
		return true
	}

	return false
}

// SetCurrencyType gets a reference to the given string and assigns it to the CurrencyType field.
func (o *ResponseCurrenciesResponse) SetCurrencyType(v string) {
	o.CurrencyType = &v
}

// GetDenomination returns the Denomination field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetDenomination() float32 {
	if o == nil || IsNil(o.Denomination) {
		var ret float32
		return ret
	}
	return *o.Denomination
}

// GetDenominationOk returns a tuple with the Denomination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetDenominationOk() (*float32, bool) {
	if o == nil || IsNil(o.Denomination) {
		return nil, false
	}
	return o.Denomination, true
}

// HasDenomination returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasDenomination() bool {
	if o != nil && !IsNil(o.Denomination) {
		return true
	}

	return false
}

// SetDenomination gets a reference to the given float32 and assigns it to the Denomination field.
func (o *ResponseCurrenciesResponse) SetDenomination(v float32) {
	o.Denomination = &v
}

// GetEnteredByUser returns the EnteredByUser field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetEnteredByUser() string {
	if o == nil || IsNil(o.EnteredByUser) {
		var ret string
		return ret
	}
	return *o.EnteredByUser
}

// GetEnteredByUserOk returns a tuple with the EnteredByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetEnteredByUserOk() (*string, bool) {
	if o == nil || IsNil(o.EnteredByUser) {
		return nil, false
	}
	return o.EnteredByUser, true
}

// HasEnteredByUser returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasEnteredByUser() bool {
	if o != nil && !IsNil(o.EnteredByUser) {
		return true
	}

	return false
}

// SetEnteredByUser gets a reference to the given string and assigns it to the EnteredByUser field.
func (o *ResponseCurrenciesResponse) SetEnteredByUser(v string) {
	o.EnteredByUser = &v
}

// GetEntryDate returns the EntryDate field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetEntryDate() string {
	if o == nil || IsNil(o.EntryDate) {
		var ret string
		return ret
	}
	return *o.EntryDate
}

// GetEntryDateOk returns a tuple with the EntryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetEntryDateOk() (*string, bool) {
	if o == nil || IsNil(o.EntryDate) {
		return nil, false
	}
	return o.EntryDate, true
}

// HasEntryDate returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasEntryDate() bool {
	if o != nil && !IsNil(o.EntryDate) {
		return true
	}

	return false
}

// SetEntryDate gets a reference to the given string and assigns it to the EntryDate field.
func (o *ResponseCurrenciesResponse) SetEntryDate(v string) {
	o.EntryDate = &v
}

// GetUpdateDate returns the UpdateDate field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetUpdateDate() string {
	if o == nil || IsNil(o.UpdateDate) {
		var ret string
		return ret
	}
	return *o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetUpdateDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateDate) {
		return nil, false
	}
	return o.UpdateDate, true
}

// HasUpdateDate returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasUpdateDate() bool {
	if o != nil && !IsNil(o.UpdateDate) {
		return true
	}

	return false
}

// SetUpdateDate gets a reference to the given string and assigns it to the UpdateDate field.
func (o *ResponseCurrenciesResponse) SetUpdateDate(v string) {
	o.UpdateDate = &v
}

// GetUpdatedByUser returns the UpdatedByUser field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetUpdatedByUser() string {
	if o == nil || IsNil(o.UpdatedByUser) {
		var ret string
		return ret
	}
	return *o.UpdatedByUser
}

// GetUpdatedByUserOk returns a tuple with the UpdatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetUpdatedByUserOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedByUser) {
		return nil, false
	}
	return o.UpdatedByUser, true
}

// HasUpdatedByUser returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasUpdatedByUser() bool {
	if o != nil && !IsNil(o.UpdatedByUser) {
		return true
	}

	return false
}

// SetUpdatedByUser gets a reference to the given string and assigns it to the UpdatedByUser field.
func (o *ResponseCurrenciesResponse) SetUpdatedByUser(v string) {
	o.UpdatedByUser = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ResponseCurrenciesResponse) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *ResponseCurrenciesResponse) GetValidTo() string {
	if o == nil || IsNil(o.ValidTo) {
		var ret string
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCurrenciesResponse) GetValidToOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTo) {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *ResponseCurrenciesResponse) HasValidTo() bool {
	if o != nil && !IsNil(o.ValidTo) {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given string and assigns it to the ValidTo field.
func (o *ResponseCurrenciesResponse) SetValidTo(v string) {
	o.ValidTo = &v
}

func (o ResponseCurrenciesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCurrenciesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.CurrencyType) {
		toSerialize["currency_type"] = o.CurrencyType
	}
	if !IsNil(o.Denomination) {
		toSerialize["denomination"] = o.Denomination
	}
	if !IsNil(o.EnteredByUser) {
		toSerialize["entered_by_user"] = o.EnteredByUser
	}
	if !IsNil(o.EntryDate) {
		toSerialize["entry_date"] = o.EntryDate
	}
	if !IsNil(o.UpdateDate) {
		toSerialize["update_date"] = o.UpdateDate
	}
	if !IsNil(o.UpdatedByUser) {
		toSerialize["updated_by_user"] = o.UpdatedByUser
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if !IsNil(o.ValidTo) {
		toSerialize["valid_to"] = o.ValidTo
	}
	return toSerialize, nil
}

type NullableResponseCurrenciesResponse struct {
	value *ResponseCurrenciesResponse
	isSet bool
}

func (v NullableResponseCurrenciesResponse) Get() *ResponseCurrenciesResponse {
	return v.value
}

func (v *NullableResponseCurrenciesResponse) Set(val *ResponseCurrenciesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCurrenciesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCurrenciesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCurrenciesResponse(val *ResponseCurrenciesResponse) *NullableResponseCurrenciesResponse {
	return &NullableResponseCurrenciesResponse{value: val, isSet: true}
}

func (v NullableResponseCurrenciesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCurrenciesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



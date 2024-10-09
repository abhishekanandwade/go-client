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

// checks if the ResponseIPOdetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIPOdetails{}

// ResponseIPOdetails struct for ResponseIPOdetails
type ResponseIPOdetails struct {
	CurrQty *int32 `json:"curr_qty,omitempty"`
	CurrSerial *string `json:"curr_serial,omitempty"`
	DenominationDesc *string `json:"denomination_Desc,omitempty"`
	DenominationValue *string `json:"denomination_Value,omitempty"`
	DenominationId *string `json:"denomination_id,omitempty"`
	EndSerial *string `json:"end_serial,omitempty"`
	PrefixId *string `json:"prefix_id,omitempty"`
	StartSerial *string `json:"start_serial,omitempty"`
	TotQuantity *int32 `json:"tot_quantity,omitempty"`
}

// NewResponseIPOdetails instantiates a new ResponseIPOdetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIPOdetails() *ResponseIPOdetails {
	this := ResponseIPOdetails{}
	return &this
}

// NewResponseIPOdetailsWithDefaults instantiates a new ResponseIPOdetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIPOdetailsWithDefaults() *ResponseIPOdetails {
	this := ResponseIPOdetails{}
	return &this
}

// GetCurrQty returns the CurrQty field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetCurrQty() int32 {
	if o == nil || IsNil(o.CurrQty) {
		var ret int32
		return ret
	}
	return *o.CurrQty
}

// GetCurrQtyOk returns a tuple with the CurrQty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetCurrQtyOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrQty) {
		return nil, false
	}
	return o.CurrQty, true
}

// HasCurrQty returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasCurrQty() bool {
	if o != nil && !IsNil(o.CurrQty) {
		return true
	}

	return false
}

// SetCurrQty gets a reference to the given int32 and assigns it to the CurrQty field.
func (o *ResponseIPOdetails) SetCurrQty(v int32) {
	o.CurrQty = &v
}

// GetCurrSerial returns the CurrSerial field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetCurrSerial() string {
	if o == nil || IsNil(o.CurrSerial) {
		var ret string
		return ret
	}
	return *o.CurrSerial
}

// GetCurrSerialOk returns a tuple with the CurrSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetCurrSerialOk() (*string, bool) {
	if o == nil || IsNil(o.CurrSerial) {
		return nil, false
	}
	return o.CurrSerial, true
}

// HasCurrSerial returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasCurrSerial() bool {
	if o != nil && !IsNil(o.CurrSerial) {
		return true
	}

	return false
}

// SetCurrSerial gets a reference to the given string and assigns it to the CurrSerial field.
func (o *ResponseIPOdetails) SetCurrSerial(v string) {
	o.CurrSerial = &v
}

// GetDenominationDesc returns the DenominationDesc field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetDenominationDesc() string {
	if o == nil || IsNil(o.DenominationDesc) {
		var ret string
		return ret
	}
	return *o.DenominationDesc
}

// GetDenominationDescOk returns a tuple with the DenominationDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetDenominationDescOk() (*string, bool) {
	if o == nil || IsNil(o.DenominationDesc) {
		return nil, false
	}
	return o.DenominationDesc, true
}

// HasDenominationDesc returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasDenominationDesc() bool {
	if o != nil && !IsNil(o.DenominationDesc) {
		return true
	}

	return false
}

// SetDenominationDesc gets a reference to the given string and assigns it to the DenominationDesc field.
func (o *ResponseIPOdetails) SetDenominationDesc(v string) {
	o.DenominationDesc = &v
}

// GetDenominationValue returns the DenominationValue field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetDenominationValue() string {
	if o == nil || IsNil(o.DenominationValue) {
		var ret string
		return ret
	}
	return *o.DenominationValue
}

// GetDenominationValueOk returns a tuple with the DenominationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetDenominationValueOk() (*string, bool) {
	if o == nil || IsNil(o.DenominationValue) {
		return nil, false
	}
	return o.DenominationValue, true
}

// HasDenominationValue returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasDenominationValue() bool {
	if o != nil && !IsNil(o.DenominationValue) {
		return true
	}

	return false
}

// SetDenominationValue gets a reference to the given string and assigns it to the DenominationValue field.
func (o *ResponseIPOdetails) SetDenominationValue(v string) {
	o.DenominationValue = &v
}

// GetDenominationId returns the DenominationId field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetDenominationId() string {
	if o == nil || IsNil(o.DenominationId) {
		var ret string
		return ret
	}
	return *o.DenominationId
}

// GetDenominationIdOk returns a tuple with the DenominationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetDenominationIdOk() (*string, bool) {
	if o == nil || IsNil(o.DenominationId) {
		return nil, false
	}
	return o.DenominationId, true
}

// HasDenominationId returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasDenominationId() bool {
	if o != nil && !IsNil(o.DenominationId) {
		return true
	}

	return false
}

// SetDenominationId gets a reference to the given string and assigns it to the DenominationId field.
func (o *ResponseIPOdetails) SetDenominationId(v string) {
	o.DenominationId = &v
}

// GetEndSerial returns the EndSerial field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetEndSerial() string {
	if o == nil || IsNil(o.EndSerial) {
		var ret string
		return ret
	}
	return *o.EndSerial
}

// GetEndSerialOk returns a tuple with the EndSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetEndSerialOk() (*string, bool) {
	if o == nil || IsNil(o.EndSerial) {
		return nil, false
	}
	return o.EndSerial, true
}

// HasEndSerial returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasEndSerial() bool {
	if o != nil && !IsNil(o.EndSerial) {
		return true
	}

	return false
}

// SetEndSerial gets a reference to the given string and assigns it to the EndSerial field.
func (o *ResponseIPOdetails) SetEndSerial(v string) {
	o.EndSerial = &v
}

// GetPrefixId returns the PrefixId field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetPrefixId() string {
	if o == nil || IsNil(o.PrefixId) {
		var ret string
		return ret
	}
	return *o.PrefixId
}

// GetPrefixIdOk returns a tuple with the PrefixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetPrefixIdOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixId) {
		return nil, false
	}
	return o.PrefixId, true
}

// HasPrefixId returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasPrefixId() bool {
	if o != nil && !IsNil(o.PrefixId) {
		return true
	}

	return false
}

// SetPrefixId gets a reference to the given string and assigns it to the PrefixId field.
func (o *ResponseIPOdetails) SetPrefixId(v string) {
	o.PrefixId = &v
}

// GetStartSerial returns the StartSerial field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetStartSerial() string {
	if o == nil || IsNil(o.StartSerial) {
		var ret string
		return ret
	}
	return *o.StartSerial
}

// GetStartSerialOk returns a tuple with the StartSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetStartSerialOk() (*string, bool) {
	if o == nil || IsNil(o.StartSerial) {
		return nil, false
	}
	return o.StartSerial, true
}

// HasStartSerial returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasStartSerial() bool {
	if o != nil && !IsNil(o.StartSerial) {
		return true
	}

	return false
}

// SetStartSerial gets a reference to the given string and assigns it to the StartSerial field.
func (o *ResponseIPOdetails) SetStartSerial(v string) {
	o.StartSerial = &v
}

// GetTotQuantity returns the TotQuantity field value if set, zero value otherwise.
func (o *ResponseIPOdetails) GetTotQuantity() int32 {
	if o == nil || IsNil(o.TotQuantity) {
		var ret int32
		return ret
	}
	return *o.TotQuantity
}

// GetTotQuantityOk returns a tuple with the TotQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIPOdetails) GetTotQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.TotQuantity) {
		return nil, false
	}
	return o.TotQuantity, true
}

// HasTotQuantity returns a boolean if a field has been set.
func (o *ResponseIPOdetails) HasTotQuantity() bool {
	if o != nil && !IsNil(o.TotQuantity) {
		return true
	}

	return false
}

// SetTotQuantity gets a reference to the given int32 and assigns it to the TotQuantity field.
func (o *ResponseIPOdetails) SetTotQuantity(v int32) {
	o.TotQuantity = &v
}

func (o ResponseIPOdetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIPOdetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrQty) {
		toSerialize["curr_qty"] = o.CurrQty
	}
	if !IsNil(o.CurrSerial) {
		toSerialize["curr_serial"] = o.CurrSerial
	}
	if !IsNil(o.DenominationDesc) {
		toSerialize["denomination_Desc"] = o.DenominationDesc
	}
	if !IsNil(o.DenominationValue) {
		toSerialize["denomination_Value"] = o.DenominationValue
	}
	if !IsNil(o.DenominationId) {
		toSerialize["denomination_id"] = o.DenominationId
	}
	if !IsNil(o.EndSerial) {
		toSerialize["end_serial"] = o.EndSerial
	}
	if !IsNil(o.PrefixId) {
		toSerialize["prefix_id"] = o.PrefixId
	}
	if !IsNil(o.StartSerial) {
		toSerialize["start_serial"] = o.StartSerial
	}
	if !IsNil(o.TotQuantity) {
		toSerialize["tot_quantity"] = o.TotQuantity
	}
	return toSerialize, nil
}

type NullableResponseIPOdetails struct {
	value *ResponseIPOdetails
	isSet bool
}

func (v NullableResponseIPOdetails) Get() *ResponseIPOdetails {
	return v.value
}

func (v *NullableResponseIPOdetails) Set(val *ResponseIPOdetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIPOdetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIPOdetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIPOdetails(val *ResponseIPOdetails) *NullableResponseIPOdetails {
	return &NullableResponseIPOdetails{value: val, isSet: true}
}

func (v NullableResponseIPOdetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIPOdetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


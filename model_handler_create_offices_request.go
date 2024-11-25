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
	"bytes"
	"fmt"
)

// checks if the HandlerCreateOfficesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateOfficesRequest{}

// HandlerCreateOfficesRequest struct for HandlerCreateOfficesRequest
type HandlerCreateOfficesRequest struct {
	BankCreditLimit *float32 `json:"bank_credit_limit,omitempty"`
	BankId *string `json:"bank_id,omitempty"`
	CashOfficeId *int32 `json:"cash_office_id,omitempty"`
	CashofficeName *string `json:"cashoffice_name,omitempty"`
	CounterCashTfrLimit *float32 `json:"counter_cash_tfr_limit,omitempty"`
	EnteredBy string `json:"entered_by"`
	MaxAmt float32 `json:"max_amt"`
	MinAmt float32 `json:"min_amt"`
	OfficeId int32 `json:"office_id"`
	OfficeName string `json:"office_name"`
	PostageStampsLimit *float32 `json:"postage_stamps_limit,omitempty"`
	RevenueStampsLimit *float32 `json:"revenue_stamps_limit,omitempty"`
	TransitDuration *int32 `json:"transit_duration,omitempty"`
	ValidFrom string `json:"valid_from"`
}

type _HandlerCreateOfficesRequest HandlerCreateOfficesRequest

// NewHandlerCreateOfficesRequest instantiates a new HandlerCreateOfficesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateOfficesRequest(enteredBy string, maxAmt float32, minAmt float32, officeId int32, officeName string, validFrom string) *HandlerCreateOfficesRequest {
	this := HandlerCreateOfficesRequest{}
	this.EnteredBy = enteredBy
	this.MaxAmt = maxAmt
	this.MinAmt = minAmt
	this.OfficeId = officeId
	this.OfficeName = officeName
	this.ValidFrom = validFrom
	return &this
}

// NewHandlerCreateOfficesRequestWithDefaults instantiates a new HandlerCreateOfficesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateOfficesRequestWithDefaults() *HandlerCreateOfficesRequest {
	this := HandlerCreateOfficesRequest{}
	return &this
}

// GetBankCreditLimit returns the BankCreditLimit field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetBankCreditLimit() float32 {
	if o == nil || IsNil(o.BankCreditLimit) {
		var ret float32
		return ret
	}
	return *o.BankCreditLimit
}

// GetBankCreditLimitOk returns a tuple with the BankCreditLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetBankCreditLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.BankCreditLimit) {
		return nil, false
	}
	return o.BankCreditLimit, true
}

// HasBankCreditLimit returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasBankCreditLimit() bool {
	if o != nil && !IsNil(o.BankCreditLimit) {
		return true
	}

	return false
}

// SetBankCreditLimit gets a reference to the given float32 and assigns it to the BankCreditLimit field.
func (o *HandlerCreateOfficesRequest) SetBankCreditLimit(v float32) {
	o.BankCreditLimit = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetBankId() string {
	if o == nil || IsNil(o.BankId) {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetBankIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankId) {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasBankId() bool {
	if o != nil && !IsNil(o.BankId) {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *HandlerCreateOfficesRequest) SetBankId(v string) {
	o.BankId = &v
}

// GetCashOfficeId returns the CashOfficeId field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetCashOfficeId() int32 {
	if o == nil || IsNil(o.CashOfficeId) {
		var ret int32
		return ret
	}
	return *o.CashOfficeId
}

// GetCashOfficeIdOk returns a tuple with the CashOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetCashOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CashOfficeId) {
		return nil, false
	}
	return o.CashOfficeId, true
}

// HasCashOfficeId returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasCashOfficeId() bool {
	if o != nil && !IsNil(o.CashOfficeId) {
		return true
	}

	return false
}

// SetCashOfficeId gets a reference to the given int32 and assigns it to the CashOfficeId field.
func (o *HandlerCreateOfficesRequest) SetCashOfficeId(v int32) {
	o.CashOfficeId = &v
}

// GetCashofficeName returns the CashofficeName field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetCashofficeName() string {
	if o == nil || IsNil(o.CashofficeName) {
		var ret string
		return ret
	}
	return *o.CashofficeName
}

// GetCashofficeNameOk returns a tuple with the CashofficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetCashofficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.CashofficeName) {
		return nil, false
	}
	return o.CashofficeName, true
}

// HasCashofficeName returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasCashofficeName() bool {
	if o != nil && !IsNil(o.CashofficeName) {
		return true
	}

	return false
}

// SetCashofficeName gets a reference to the given string and assigns it to the CashofficeName field.
func (o *HandlerCreateOfficesRequest) SetCashofficeName(v string) {
	o.CashofficeName = &v
}

// GetCounterCashTfrLimit returns the CounterCashTfrLimit field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetCounterCashTfrLimit() float32 {
	if o == nil || IsNil(o.CounterCashTfrLimit) {
		var ret float32
		return ret
	}
	return *o.CounterCashTfrLimit
}

// GetCounterCashTfrLimitOk returns a tuple with the CounterCashTfrLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetCounterCashTfrLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.CounterCashTfrLimit) {
		return nil, false
	}
	return o.CounterCashTfrLimit, true
}

// HasCounterCashTfrLimit returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasCounterCashTfrLimit() bool {
	if o != nil && !IsNil(o.CounterCashTfrLimit) {
		return true
	}

	return false
}

// SetCounterCashTfrLimit gets a reference to the given float32 and assigns it to the CounterCashTfrLimit field.
func (o *HandlerCreateOfficesRequest) SetCounterCashTfrLimit(v float32) {
	o.CounterCashTfrLimit = &v
}

// GetEnteredBy returns the EnteredBy field value
func (o *HandlerCreateOfficesRequest) GetEnteredBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnteredBy
}

// GetEnteredByOk returns a tuple with the EnteredBy field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetEnteredByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnteredBy, true
}

// SetEnteredBy sets field value
func (o *HandlerCreateOfficesRequest) SetEnteredBy(v string) {
	o.EnteredBy = v
}

// GetMaxAmt returns the MaxAmt field value
func (o *HandlerCreateOfficesRequest) GetMaxAmt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxAmt
}

// GetMaxAmtOk returns a tuple with the MaxAmt field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetMaxAmtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAmt, true
}

// SetMaxAmt sets field value
func (o *HandlerCreateOfficesRequest) SetMaxAmt(v float32) {
	o.MaxAmt = v
}

// GetMinAmt returns the MinAmt field value
func (o *HandlerCreateOfficesRequest) GetMinAmt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinAmt
}

// GetMinAmtOk returns a tuple with the MinAmt field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetMinAmtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinAmt, true
}

// SetMinAmt sets field value
func (o *HandlerCreateOfficesRequest) SetMinAmt(v float32) {
	o.MinAmt = v
}

// GetOfficeId returns the OfficeId field value
func (o *HandlerCreateOfficesRequest) GetOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfficeId, true
}

// SetOfficeId sets field value
func (o *HandlerCreateOfficesRequest) SetOfficeId(v int32) {
	o.OfficeId = v
}

// GetOfficeName returns the OfficeName field value
func (o *HandlerCreateOfficesRequest) GetOfficeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OfficeName
}

// GetOfficeNameOk returns a tuple with the OfficeName field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetOfficeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfficeName, true
}

// SetOfficeName sets field value
func (o *HandlerCreateOfficesRequest) SetOfficeName(v string) {
	o.OfficeName = v
}

// GetPostageStampsLimit returns the PostageStampsLimit field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetPostageStampsLimit() float32 {
	if o == nil || IsNil(o.PostageStampsLimit) {
		var ret float32
		return ret
	}
	return *o.PostageStampsLimit
}

// GetPostageStampsLimitOk returns a tuple with the PostageStampsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetPostageStampsLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.PostageStampsLimit) {
		return nil, false
	}
	return o.PostageStampsLimit, true
}

// HasPostageStampsLimit returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasPostageStampsLimit() bool {
	if o != nil && !IsNil(o.PostageStampsLimit) {
		return true
	}

	return false
}

// SetPostageStampsLimit gets a reference to the given float32 and assigns it to the PostageStampsLimit field.
func (o *HandlerCreateOfficesRequest) SetPostageStampsLimit(v float32) {
	o.PostageStampsLimit = &v
}

// GetRevenueStampsLimit returns the RevenueStampsLimit field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetRevenueStampsLimit() float32 {
	if o == nil || IsNil(o.RevenueStampsLimit) {
		var ret float32
		return ret
	}
	return *o.RevenueStampsLimit
}

// GetRevenueStampsLimitOk returns a tuple with the RevenueStampsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetRevenueStampsLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.RevenueStampsLimit) {
		return nil, false
	}
	return o.RevenueStampsLimit, true
}

// HasRevenueStampsLimit returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasRevenueStampsLimit() bool {
	if o != nil && !IsNil(o.RevenueStampsLimit) {
		return true
	}

	return false
}

// SetRevenueStampsLimit gets a reference to the given float32 and assigns it to the RevenueStampsLimit field.
func (o *HandlerCreateOfficesRequest) SetRevenueStampsLimit(v float32) {
	o.RevenueStampsLimit = &v
}

// GetTransitDuration returns the TransitDuration field value if set, zero value otherwise.
func (o *HandlerCreateOfficesRequest) GetTransitDuration() int32 {
	if o == nil || IsNil(o.TransitDuration) {
		var ret int32
		return ret
	}
	return *o.TransitDuration
}

// GetTransitDurationOk returns a tuple with the TransitDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetTransitDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.TransitDuration) {
		return nil, false
	}
	return o.TransitDuration, true
}

// HasTransitDuration returns a boolean if a field has been set.
func (o *HandlerCreateOfficesRequest) HasTransitDuration() bool {
	if o != nil && !IsNil(o.TransitDuration) {
		return true
	}

	return false
}

// SetTransitDuration gets a reference to the given int32 and assigns it to the TransitDuration field.
func (o *HandlerCreateOfficesRequest) SetTransitDuration(v int32) {
	o.TransitDuration = &v
}

// GetValidFrom returns the ValidFrom field value
func (o *HandlerCreateOfficesRequest) GetValidFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesRequest) GetValidFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *HandlerCreateOfficesRequest) SetValidFrom(v string) {
	o.ValidFrom = v
}

func (o HandlerCreateOfficesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateOfficesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankCreditLimit) {
		toSerialize["bank_credit_limit"] = o.BankCreditLimit
	}
	if !IsNil(o.BankId) {
		toSerialize["bank_id"] = o.BankId
	}
	if !IsNil(o.CashOfficeId) {
		toSerialize["cash_office_id"] = o.CashOfficeId
	}
	if !IsNil(o.CashofficeName) {
		toSerialize["cashoffice_name"] = o.CashofficeName
	}
	if !IsNil(o.CounterCashTfrLimit) {
		toSerialize["counter_cash_tfr_limit"] = o.CounterCashTfrLimit
	}
	toSerialize["entered_by"] = o.EnteredBy
	toSerialize["max_amt"] = o.MaxAmt
	toSerialize["min_amt"] = o.MinAmt
	toSerialize["office_id"] = o.OfficeId
	toSerialize["office_name"] = o.OfficeName
	if !IsNil(o.PostageStampsLimit) {
		toSerialize["postage_stamps_limit"] = o.PostageStampsLimit
	}
	if !IsNil(o.RevenueStampsLimit) {
		toSerialize["revenue_stamps_limit"] = o.RevenueStampsLimit
	}
	if !IsNil(o.TransitDuration) {
		toSerialize["transit_duration"] = o.TransitDuration
	}
	toSerialize["valid_from"] = o.ValidFrom
	return toSerialize, nil
}

func (o *HandlerCreateOfficesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entered_by",
		"max_amt",
		"min_amt",
		"office_id",
		"office_name",
		"valid_from",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHandlerCreateOfficesRequest := _HandlerCreateOfficesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateOfficesRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateOfficesRequest(varHandlerCreateOfficesRequest)

	return err
}

type NullableHandlerCreateOfficesRequest struct {
	value *HandlerCreateOfficesRequest
	isSet bool
}

func (v NullableHandlerCreateOfficesRequest) Get() *HandlerCreateOfficesRequest {
	return v.value
}

func (v *NullableHandlerCreateOfficesRequest) Set(val *HandlerCreateOfficesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateOfficesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateOfficesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateOfficesRequest(val *HandlerCreateOfficesRequest) *NullableHandlerCreateOfficesRequest {
	return &NullableHandlerCreateOfficesRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateOfficesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateOfficesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



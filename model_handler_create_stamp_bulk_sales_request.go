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

// checks if the HandlerCreateStampBulkSalesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateStampBulkSalesRequest{}

// HandlerCreateStampBulkSalesRequest struct for HandlerCreateStampBulkSalesRequest
type HandlerCreateStampBulkSalesRequest struct {
	BankBranch *string `json:"bank_branch,omitempty"`
	InstrumentDate *string `json:"instrument_date,omitempty"`
	InstrumentNo *string `json:"instrument_no,omitempty"`
	IsSingleHand *bool `json:"is_single_hand,omitempty"`
	OfficeId int32 `json:"office_id"`
	PaymentMode string `json:"payment_mode"`
	Remarks *string `json:"remarks,omitempty"`
	SaleAmt float32 `json:"sale_amt"`
	SaleDetails map[string]int32 `json:"sale_details"`
	SoldTo string `json:"sold_to"`
	UserId int32 `json:"user_id"`
}

type _HandlerCreateStampBulkSalesRequest HandlerCreateStampBulkSalesRequest

// NewHandlerCreateStampBulkSalesRequest instantiates a new HandlerCreateStampBulkSalesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateStampBulkSalesRequest(officeId int32, paymentMode string, saleAmt float32, saleDetails map[string]int32, soldTo string, userId int32) *HandlerCreateStampBulkSalesRequest {
	this := HandlerCreateStampBulkSalesRequest{}
	this.OfficeId = officeId
	this.PaymentMode = paymentMode
	this.SaleAmt = saleAmt
	this.SaleDetails = saleDetails
	this.SoldTo = soldTo
	this.UserId = userId
	return &this
}

// NewHandlerCreateStampBulkSalesRequestWithDefaults instantiates a new HandlerCreateStampBulkSalesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateStampBulkSalesRequestWithDefaults() *HandlerCreateStampBulkSalesRequest {
	this := HandlerCreateStampBulkSalesRequest{}
	return &this
}

// GetBankBranch returns the BankBranch field value if set, zero value otherwise.
func (o *HandlerCreateStampBulkSalesRequest) GetBankBranch() string {
	if o == nil || IsNil(o.BankBranch) {
		var ret string
		return ret
	}
	return *o.BankBranch
}

// GetBankBranchOk returns a tuple with the BankBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetBankBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BankBranch) {
		return nil, false
	}
	return o.BankBranch, true
}

// HasBankBranch returns a boolean if a field has been set.
func (o *HandlerCreateStampBulkSalesRequest) HasBankBranch() bool {
	if o != nil && !IsNil(o.BankBranch) {
		return true
	}

	return false
}

// SetBankBranch gets a reference to the given string and assigns it to the BankBranch field.
func (o *HandlerCreateStampBulkSalesRequest) SetBankBranch(v string) {
	o.BankBranch = &v
}

// GetInstrumentDate returns the InstrumentDate field value if set, zero value otherwise.
func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentDate() string {
	if o == nil || IsNil(o.InstrumentDate) {
		var ret string
		return ret
	}
	return *o.InstrumentDate
}

// GetInstrumentDateOk returns a tuple with the InstrumentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentDateOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentDate) {
		return nil, false
	}
	return o.InstrumentDate, true
}

// HasInstrumentDate returns a boolean if a field has been set.
func (o *HandlerCreateStampBulkSalesRequest) HasInstrumentDate() bool {
	if o != nil && !IsNil(o.InstrumentDate) {
		return true
	}

	return false
}

// SetInstrumentDate gets a reference to the given string and assigns it to the InstrumentDate field.
func (o *HandlerCreateStampBulkSalesRequest) SetInstrumentDate(v string) {
	o.InstrumentDate = &v
}

// GetInstrumentNo returns the InstrumentNo field value if set, zero value otherwise.
func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentNo() string {
	if o == nil || IsNil(o.InstrumentNo) {
		var ret string
		return ret
	}
	return *o.InstrumentNo
}

// GetInstrumentNoOk returns a tuple with the InstrumentNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentNoOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentNo) {
		return nil, false
	}
	return o.InstrumentNo, true
}

// HasInstrumentNo returns a boolean if a field has been set.
func (o *HandlerCreateStampBulkSalesRequest) HasInstrumentNo() bool {
	if o != nil && !IsNil(o.InstrumentNo) {
		return true
	}

	return false
}

// SetInstrumentNo gets a reference to the given string and assigns it to the InstrumentNo field.
func (o *HandlerCreateStampBulkSalesRequest) SetInstrumentNo(v string) {
	o.InstrumentNo = &v
}

// GetIsSingleHand returns the IsSingleHand field value if set, zero value otherwise.
func (o *HandlerCreateStampBulkSalesRequest) GetIsSingleHand() bool {
	if o == nil || IsNil(o.IsSingleHand) {
		var ret bool
		return ret
	}
	return *o.IsSingleHand
}

// GetIsSingleHandOk returns a tuple with the IsSingleHand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetIsSingleHandOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSingleHand) {
		return nil, false
	}
	return o.IsSingleHand, true
}

// HasIsSingleHand returns a boolean if a field has been set.
func (o *HandlerCreateStampBulkSalesRequest) HasIsSingleHand() bool {
	if o != nil && !IsNil(o.IsSingleHand) {
		return true
	}

	return false
}

// SetIsSingleHand gets a reference to the given bool and assigns it to the IsSingleHand field.
func (o *HandlerCreateStampBulkSalesRequest) SetIsSingleHand(v bool) {
	o.IsSingleHand = &v
}

// GetOfficeId returns the OfficeId field value
func (o *HandlerCreateStampBulkSalesRequest) GetOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfficeId, true
}

// SetOfficeId sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetOfficeId(v int32) {
	o.OfficeId = v
}

// GetPaymentMode returns the PaymentMode field value
func (o *HandlerCreateStampBulkSalesRequest) GetPaymentMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentMode
}

// GetPaymentModeOk returns a tuple with the PaymentMode field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetPaymentModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PaymentMode, true
}

// SetPaymentMode sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetPaymentMode(v string) {
	o.PaymentMode = v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *HandlerCreateStampBulkSalesRequest) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *HandlerCreateStampBulkSalesRequest) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *HandlerCreateStampBulkSalesRequest) SetRemarks(v string) {
	o.Remarks = &v
}

// GetSaleAmt returns the SaleAmt field value
func (o *HandlerCreateStampBulkSalesRequest) GetSaleAmt() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SaleAmt
}

// GetSaleAmtOk returns a tuple with the SaleAmt field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetSaleAmtOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleAmt, true
}

// SetSaleAmt sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetSaleAmt(v float32) {
	o.SaleAmt = v
}

// GetSaleDetails returns the SaleDetails field value
func (o *HandlerCreateStampBulkSalesRequest) GetSaleDetails() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.SaleDetails
}

// GetSaleDetailsOk returns a tuple with the SaleDetails field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetSaleDetailsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SaleDetails, true
}

// SetSaleDetails sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetSaleDetails(v map[string]int32) {
	o.SaleDetails = v
}

// GetSoldTo returns the SoldTo field value
func (o *HandlerCreateStampBulkSalesRequest) GetSoldTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SoldTo
}

// GetSoldToOk returns a tuple with the SoldTo field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetSoldToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SoldTo, true
}

// SetSoldTo sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetSoldTo(v string) {
	o.SoldTo = v
}

// GetUserId returns the UserId field value
func (o *HandlerCreateStampBulkSalesRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampBulkSalesRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *HandlerCreateStampBulkSalesRequest) SetUserId(v int32) {
	o.UserId = v
}

func (o HandlerCreateStampBulkSalesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateStampBulkSalesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankBranch) {
		toSerialize["bank_branch"] = o.BankBranch
	}
	if !IsNil(o.InstrumentDate) {
		toSerialize["instrument_date"] = o.InstrumentDate
	}
	if !IsNil(o.InstrumentNo) {
		toSerialize["instrument_no"] = o.InstrumentNo
	}
	if !IsNil(o.IsSingleHand) {
		toSerialize["is_single_hand"] = o.IsSingleHand
	}
	toSerialize["office_id"] = o.OfficeId
	toSerialize["payment_mode"] = o.PaymentMode
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	toSerialize["sale_amt"] = o.SaleAmt
	toSerialize["sale_details"] = o.SaleDetails
	toSerialize["sold_to"] = o.SoldTo
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *HandlerCreateStampBulkSalesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"office_id",
		"payment_mode",
		"sale_amt",
		"sale_details",
		"sold_to",
		"user_id",
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

	varHandlerCreateStampBulkSalesRequest := _HandlerCreateStampBulkSalesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateStampBulkSalesRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateStampBulkSalesRequest(varHandlerCreateStampBulkSalesRequest)

	return err
}

type NullableHandlerCreateStampBulkSalesRequest struct {
	value *HandlerCreateStampBulkSalesRequest
	isSet bool
}

func (v NullableHandlerCreateStampBulkSalesRequest) Get() *HandlerCreateStampBulkSalesRequest {
	return v.value
}

func (v *NullableHandlerCreateStampBulkSalesRequest) Set(val *HandlerCreateStampBulkSalesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateStampBulkSalesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateStampBulkSalesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateStampBulkSalesRequest(val *HandlerCreateStampBulkSalesRequest) *NullableHandlerCreateStampBulkSalesRequest {
	return &NullableHandlerCreateStampBulkSalesRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateStampBulkSalesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateStampBulkSalesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



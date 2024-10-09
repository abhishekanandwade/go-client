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

// checks if the HandlerCreateMiscTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateMiscTransactionsRequest{}

// HandlerCreateMiscTransactionsRequest struct for HandlerCreateMiscTransactionsRequest
type HandlerCreateMiscTransactionsRequest struct {
	ChequeDate *string `json:"cheque_date,omitempty"`
	ChequeNo *string `json:"cheque_no,omitempty"`
	OfficeId int32 `json:"office_id"`
	PayeeName *string `json:"payee_name,omitempty"`
	Remarks string `json:"remarks"`
	TransAmount float32 `json:"trans_amount"`
	TransDetails map[string]float32 `json:"trans_details"`
	TransType string `json:"trans_type"`
	TxnMode string `json:"txn_mode"`
	UserId int32 `json:"user_id"`
}

type _HandlerCreateMiscTransactionsRequest HandlerCreateMiscTransactionsRequest

// NewHandlerCreateMiscTransactionsRequest instantiates a new HandlerCreateMiscTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateMiscTransactionsRequest(officeId int32, remarks string, transAmount float32, transDetails map[string]float32, transType string, txnMode string, userId int32) *HandlerCreateMiscTransactionsRequest {
	this := HandlerCreateMiscTransactionsRequest{}
	this.OfficeId = officeId
	this.Remarks = remarks
	this.TransAmount = transAmount
	this.TransDetails = transDetails
	this.TransType = transType
	this.TxnMode = txnMode
	this.UserId = userId
	return &this
}

// NewHandlerCreateMiscTransactionsRequestWithDefaults instantiates a new HandlerCreateMiscTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateMiscTransactionsRequestWithDefaults() *HandlerCreateMiscTransactionsRequest {
	this := HandlerCreateMiscTransactionsRequest{}
	return &this
}

// GetChequeDate returns the ChequeDate field value if set, zero value otherwise.
func (o *HandlerCreateMiscTransactionsRequest) GetChequeDate() string {
	if o == nil || IsNil(o.ChequeDate) {
		var ret string
		return ret
	}
	return *o.ChequeDate
}

// GetChequeDateOk returns a tuple with the ChequeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetChequeDateOk() (*string, bool) {
	if o == nil || IsNil(o.ChequeDate) {
		return nil, false
	}
	return o.ChequeDate, true
}

// HasChequeDate returns a boolean if a field has been set.
func (o *HandlerCreateMiscTransactionsRequest) HasChequeDate() bool {
	if o != nil && !IsNil(o.ChequeDate) {
		return true
	}

	return false
}

// SetChequeDate gets a reference to the given string and assigns it to the ChequeDate field.
func (o *HandlerCreateMiscTransactionsRequest) SetChequeDate(v string) {
	o.ChequeDate = &v
}

// GetChequeNo returns the ChequeNo field value if set, zero value otherwise.
func (o *HandlerCreateMiscTransactionsRequest) GetChequeNo() string {
	if o == nil || IsNil(o.ChequeNo) {
		var ret string
		return ret
	}
	return *o.ChequeNo
}

// GetChequeNoOk returns a tuple with the ChequeNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetChequeNoOk() (*string, bool) {
	if o == nil || IsNil(o.ChequeNo) {
		return nil, false
	}
	return o.ChequeNo, true
}

// HasChequeNo returns a boolean if a field has been set.
func (o *HandlerCreateMiscTransactionsRequest) HasChequeNo() bool {
	if o != nil && !IsNil(o.ChequeNo) {
		return true
	}

	return false
}

// SetChequeNo gets a reference to the given string and assigns it to the ChequeNo field.
func (o *HandlerCreateMiscTransactionsRequest) SetChequeNo(v string) {
	o.ChequeNo = &v
}

// GetOfficeId returns the OfficeId field value
func (o *HandlerCreateMiscTransactionsRequest) GetOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfficeId, true
}

// SetOfficeId sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetOfficeId(v int32) {
	o.OfficeId = v
}

// GetPayeeName returns the PayeeName field value if set, zero value otherwise.
func (o *HandlerCreateMiscTransactionsRequest) GetPayeeName() string {
	if o == nil || IsNil(o.PayeeName) {
		var ret string
		return ret
	}
	return *o.PayeeName
}

// GetPayeeNameOk returns a tuple with the PayeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetPayeeNameOk() (*string, bool) {
	if o == nil || IsNil(o.PayeeName) {
		return nil, false
	}
	return o.PayeeName, true
}

// HasPayeeName returns a boolean if a field has been set.
func (o *HandlerCreateMiscTransactionsRequest) HasPayeeName() bool {
	if o != nil && !IsNil(o.PayeeName) {
		return true
	}

	return false
}

// SetPayeeName gets a reference to the given string and assigns it to the PayeeName field.
func (o *HandlerCreateMiscTransactionsRequest) SetPayeeName(v string) {
	o.PayeeName = &v
}

// GetRemarks returns the Remarks field value
func (o *HandlerCreateMiscTransactionsRequest) GetRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remarks, true
}

// SetRemarks sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetRemarks(v string) {
	o.Remarks = v
}

// GetTransAmount returns the TransAmount field value
func (o *HandlerCreateMiscTransactionsRequest) GetTransAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TransAmount
}

// GetTransAmountOk returns a tuple with the TransAmount field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetTransAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransAmount, true
}

// SetTransAmount sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetTransAmount(v float32) {
	o.TransAmount = v
}

// GetTransDetails returns the TransDetails field value
func (o *HandlerCreateMiscTransactionsRequest) GetTransDetails() map[string]float32 {
	if o == nil {
		var ret map[string]float32
		return ret
	}

	return o.TransDetails
}

// GetTransDetailsOk returns a tuple with the TransDetails field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetTransDetailsOk() (*map[string]float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransDetails, true
}

// SetTransDetails sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetTransDetails(v map[string]float32) {
	o.TransDetails = v
}

// GetTransType returns the TransType field value
func (o *HandlerCreateMiscTransactionsRequest) GetTransType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransType
}

// GetTransTypeOk returns a tuple with the TransType field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetTransTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransType, true
}

// SetTransType sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetTransType(v string) {
	o.TransType = v
}

// GetTxnMode returns the TxnMode field value
func (o *HandlerCreateMiscTransactionsRequest) GetTxnMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxnMode
}

// GetTxnModeOk returns a tuple with the TxnMode field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetTxnModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxnMode, true
}

// SetTxnMode sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetTxnMode(v string) {
	o.TxnMode = v
}

// GetUserId returns the UserId field value
func (o *HandlerCreateMiscTransactionsRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateMiscTransactionsRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *HandlerCreateMiscTransactionsRequest) SetUserId(v int32) {
	o.UserId = v
}

func (o HandlerCreateMiscTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateMiscTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChequeDate) {
		toSerialize["cheque_date"] = o.ChequeDate
	}
	if !IsNil(o.ChequeNo) {
		toSerialize["cheque_no"] = o.ChequeNo
	}
	toSerialize["office_id"] = o.OfficeId
	if !IsNil(o.PayeeName) {
		toSerialize["payee_name"] = o.PayeeName
	}
	toSerialize["remarks"] = o.Remarks
	toSerialize["trans_amount"] = o.TransAmount
	toSerialize["trans_details"] = o.TransDetails
	toSerialize["trans_type"] = o.TransType
	toSerialize["txn_mode"] = o.TxnMode
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *HandlerCreateMiscTransactionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"office_id",
		"remarks",
		"trans_amount",
		"trans_details",
		"trans_type",
		"txn_mode",
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

	varHandlerCreateMiscTransactionsRequest := _HandlerCreateMiscTransactionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateMiscTransactionsRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateMiscTransactionsRequest(varHandlerCreateMiscTransactionsRequest)

	return err
}

type NullableHandlerCreateMiscTransactionsRequest struct {
	value *HandlerCreateMiscTransactionsRequest
	isSet bool
}

func (v NullableHandlerCreateMiscTransactionsRequest) Get() *HandlerCreateMiscTransactionsRequest {
	return v.value
}

func (v *NullableHandlerCreateMiscTransactionsRequest) Set(val *HandlerCreateMiscTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateMiscTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateMiscTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateMiscTransactionsRequest(val *HandlerCreateMiscTransactionsRequest) *NullableHandlerCreateMiscTransactionsRequest {
	return &NullableHandlerCreateMiscTransactionsRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateMiscTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateMiscTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


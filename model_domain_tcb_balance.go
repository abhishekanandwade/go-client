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

// checks if the DomainTcbBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainTcbBalance{}

// DomainTcbBalance struct for DomainTcbBalance
type DomainTcbBalance struct {
	ChequeIssueDetails []DomainChequeDetails `json:"cheque_issue_details,omitempty"`
	ChequeRemittanceDetails []DomainChequeDetails `json:"cheque_remittance_details,omitempty"`
	ClosingBal *float32 `json:"closing_bal,omitempty"`
	CurrencyDetails *string `json:"currency_details,omitempty"`
	CurrencyIdCount *map[string]int32 `json:"currency_id_count,omitempty"`
	DayendStatus *int32 `json:"dayend_status,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	OpeningBal *float32 `json:"opening_bal,omitempty"`
	Payments *float32 `json:"payments,omitempty"`
	Receipts *float32 `json:"receipts,omitempty"`
	TcbData []DomainTCBData `json:"tcb_data,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransId *string `json:"trans_id,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// NewDomainTcbBalance instantiates a new DomainTcbBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainTcbBalance() *DomainTcbBalance {
	this := DomainTcbBalance{}
	return &this
}

// NewDomainTcbBalanceWithDefaults instantiates a new DomainTcbBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainTcbBalanceWithDefaults() *DomainTcbBalance {
	this := DomainTcbBalance{}
	return &this
}

// GetChequeIssueDetails returns the ChequeIssueDetails field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetChequeIssueDetails() []DomainChequeDetails {
	if o == nil || IsNil(o.ChequeIssueDetails) {
		var ret []DomainChequeDetails
		return ret
	}
	return o.ChequeIssueDetails
}

// GetChequeIssueDetailsOk returns a tuple with the ChequeIssueDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetChequeIssueDetailsOk() ([]DomainChequeDetails, bool) {
	if o == nil || IsNil(o.ChequeIssueDetails) {
		return nil, false
	}
	return o.ChequeIssueDetails, true
}

// HasChequeIssueDetails returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasChequeIssueDetails() bool {
	if o != nil && !IsNil(o.ChequeIssueDetails) {
		return true
	}

	return false
}

// SetChequeIssueDetails gets a reference to the given []DomainChequeDetails and assigns it to the ChequeIssueDetails field.
func (o *DomainTcbBalance) SetChequeIssueDetails(v []DomainChequeDetails) {
	o.ChequeIssueDetails = v
}

// GetChequeRemittanceDetails returns the ChequeRemittanceDetails field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetChequeRemittanceDetails() []DomainChequeDetails {
	if o == nil || IsNil(o.ChequeRemittanceDetails) {
		var ret []DomainChequeDetails
		return ret
	}
	return o.ChequeRemittanceDetails
}

// GetChequeRemittanceDetailsOk returns a tuple with the ChequeRemittanceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetChequeRemittanceDetailsOk() ([]DomainChequeDetails, bool) {
	if o == nil || IsNil(o.ChequeRemittanceDetails) {
		return nil, false
	}
	return o.ChequeRemittanceDetails, true
}

// HasChequeRemittanceDetails returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasChequeRemittanceDetails() bool {
	if o != nil && !IsNil(o.ChequeRemittanceDetails) {
		return true
	}

	return false
}

// SetChequeRemittanceDetails gets a reference to the given []DomainChequeDetails and assigns it to the ChequeRemittanceDetails field.
func (o *DomainTcbBalance) SetChequeRemittanceDetails(v []DomainChequeDetails) {
	o.ChequeRemittanceDetails = v
}

// GetClosingBal returns the ClosingBal field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetClosingBal() float32 {
	if o == nil || IsNil(o.ClosingBal) {
		var ret float32
		return ret
	}
	return *o.ClosingBal
}

// GetClosingBalOk returns a tuple with the ClosingBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetClosingBalOk() (*float32, bool) {
	if o == nil || IsNil(o.ClosingBal) {
		return nil, false
	}
	return o.ClosingBal, true
}

// HasClosingBal returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasClosingBal() bool {
	if o != nil && !IsNil(o.ClosingBal) {
		return true
	}

	return false
}

// SetClosingBal gets a reference to the given float32 and assigns it to the ClosingBal field.
func (o *DomainTcbBalance) SetClosingBal(v float32) {
	o.ClosingBal = &v
}

// GetCurrencyDetails returns the CurrencyDetails field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetCurrencyDetails() string {
	if o == nil || IsNil(o.CurrencyDetails) {
		var ret string
		return ret
	}
	return *o.CurrencyDetails
}

// GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetCurrencyDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyDetails) {
		return nil, false
	}
	return o.CurrencyDetails, true
}

// HasCurrencyDetails returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasCurrencyDetails() bool {
	if o != nil && !IsNil(o.CurrencyDetails) {
		return true
	}

	return false
}

// SetCurrencyDetails gets a reference to the given string and assigns it to the CurrencyDetails field.
func (o *DomainTcbBalance) SetCurrencyDetails(v string) {
	o.CurrencyDetails = &v
}

// GetCurrencyIdCount returns the CurrencyIdCount field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetCurrencyIdCount() map[string]int32 {
	if o == nil || IsNil(o.CurrencyIdCount) {
		var ret map[string]int32
		return ret
	}
	return *o.CurrencyIdCount
}

// GetCurrencyIdCountOk returns a tuple with the CurrencyIdCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetCurrencyIdCountOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.CurrencyIdCount) {
		return nil, false
	}
	return o.CurrencyIdCount, true
}

// HasCurrencyIdCount returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasCurrencyIdCount() bool {
	if o != nil && !IsNil(o.CurrencyIdCount) {
		return true
	}

	return false
}

// SetCurrencyIdCount gets a reference to the given map[string]int32 and assigns it to the CurrencyIdCount field.
func (o *DomainTcbBalance) SetCurrencyIdCount(v map[string]int32) {
	o.CurrencyIdCount = &v
}

// GetDayendStatus returns the DayendStatus field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetDayendStatus() int32 {
	if o == nil || IsNil(o.DayendStatus) {
		var ret int32
		return ret
	}
	return *o.DayendStatus
}

// GetDayendStatusOk returns a tuple with the DayendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetDayendStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.DayendStatus) {
		return nil, false
	}
	return o.DayendStatus, true
}

// HasDayendStatus returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasDayendStatus() bool {
	if o != nil && !IsNil(o.DayendStatus) {
		return true
	}

	return false
}

// SetDayendStatus gets a reference to the given int32 and assigns it to the DayendStatus field.
func (o *DomainTcbBalance) SetDayendStatus(v int32) {
	o.DayendStatus = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *DomainTcbBalance) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOpeningBal returns the OpeningBal field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetOpeningBal() float32 {
	if o == nil || IsNil(o.OpeningBal) {
		var ret float32
		return ret
	}
	return *o.OpeningBal
}

// GetOpeningBalOk returns a tuple with the OpeningBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetOpeningBalOk() (*float32, bool) {
	if o == nil || IsNil(o.OpeningBal) {
		return nil, false
	}
	return o.OpeningBal, true
}

// HasOpeningBal returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasOpeningBal() bool {
	if o != nil && !IsNil(o.OpeningBal) {
		return true
	}

	return false
}

// SetOpeningBal gets a reference to the given float32 and assigns it to the OpeningBal field.
func (o *DomainTcbBalance) SetOpeningBal(v float32) {
	o.OpeningBal = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetPayments() float32 {
	if o == nil || IsNil(o.Payments) {
		var ret float32
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetPaymentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given float32 and assigns it to the Payments field.
func (o *DomainTcbBalance) SetPayments(v float32) {
	o.Payments = &v
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetReceipts() float32 {
	if o == nil || IsNil(o.Receipts) {
		var ret float32
		return ret
	}
	return *o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetReceiptsOk() (*float32, bool) {
	if o == nil || IsNil(o.Receipts) {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasReceipts() bool {
	if o != nil && !IsNil(o.Receipts) {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given float32 and assigns it to the Receipts field.
func (o *DomainTcbBalance) SetReceipts(v float32) {
	o.Receipts = &v
}

// GetTcbData returns the TcbData field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetTcbData() []DomainTCBData {
	if o == nil || IsNil(o.TcbData) {
		var ret []DomainTCBData
		return ret
	}
	return o.TcbData
}

// GetTcbDataOk returns a tuple with the TcbData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetTcbDataOk() ([]DomainTCBData, bool) {
	if o == nil || IsNil(o.TcbData) {
		return nil, false
	}
	return o.TcbData, true
}

// HasTcbData returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasTcbData() bool {
	if o != nil && !IsNil(o.TcbData) {
		return true
	}

	return false
}

// SetTcbData gets a reference to the given []DomainTCBData and assigns it to the TcbData field.
func (o *DomainTcbBalance) SetTcbData(v []DomainTCBData) {
	o.TcbData = v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *DomainTcbBalance) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransId returns the TransId field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetTransId() string {
	if o == nil || IsNil(o.TransId) {
		var ret string
		return ret
	}
	return *o.TransId
}

// GetTransIdOk returns a tuple with the TransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetTransIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransId) {
		return nil, false
	}
	return o.TransId, true
}

// HasTransId returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasTransId() bool {
	if o != nil && !IsNil(o.TransId) {
		return true
	}

	return false
}

// SetTransId gets a reference to the given string and assigns it to the TransId field.
func (o *DomainTcbBalance) SetTransId(v string) {
	o.TransId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *DomainTcbBalance) SetUserId(v int32) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *DomainTcbBalance) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainTcbBalance) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *DomainTcbBalance) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *DomainTcbBalance) SetUserName(v string) {
	o.UserName = &v
}

func (o DomainTcbBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainTcbBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChequeIssueDetails) {
		toSerialize["cheque_issue_details"] = o.ChequeIssueDetails
	}
	if !IsNil(o.ChequeRemittanceDetails) {
		toSerialize["cheque_remittance_details"] = o.ChequeRemittanceDetails
	}
	if !IsNil(o.ClosingBal) {
		toSerialize["closing_bal"] = o.ClosingBal
	}
	if !IsNil(o.CurrencyDetails) {
		toSerialize["currency_details"] = o.CurrencyDetails
	}
	if !IsNil(o.CurrencyIdCount) {
		toSerialize["currency_id_count"] = o.CurrencyIdCount
	}
	if !IsNil(o.DayendStatus) {
		toSerialize["dayend_status"] = o.DayendStatus
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.OpeningBal) {
		toSerialize["opening_bal"] = o.OpeningBal
	}
	if !IsNil(o.Payments) {
		toSerialize["payments"] = o.Payments
	}
	if !IsNil(o.Receipts) {
		toSerialize["receipts"] = o.Receipts
	}
	if !IsNil(o.TcbData) {
		toSerialize["tcb_data"] = o.TcbData
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransId) {
		toSerialize["trans_id"] = o.TransId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	return toSerialize, nil
}

type NullableDomainTcbBalance struct {
	value *DomainTcbBalance
	isSet bool
}

func (v NullableDomainTcbBalance) Get() *DomainTcbBalance {
	return v.value
}

func (v *NullableDomainTcbBalance) Set(val *DomainTcbBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainTcbBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainTcbBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainTcbBalance(val *DomainTcbBalance) *NullableDomainTcbBalance {
	return &NullableDomainTcbBalance{value: val, isSet: true}
}

func (v NullableDomainTcbBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainTcbBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the ResponseBodaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseBodaData{}

// ResponseBodaData struct for ResponseBodaData
type ResponseBodaData struct {
	AccountingDetail []ResponseAccountingDetails `json:"accountingDetail,omitempty"`
	CashBagDetail []ResponseCashBagDetail `json:"cashBagDetail,omitempty"`
	ClosingBal *float32 `json:"closing_bal,omitempty"`
	DayendStatus *int32 `json:"dayend_status,omitempty"`
	Ipobalance []ResponseIPOBal `json:"ipobalance,omitempty"`
	MaxBal *float32 `json:"max_bal,omitempty"`
	MinBal *float32 `json:"min_bal,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	OpeningBal *float32 `json:"opening_bal,omitempty"`
	Payments *float32 `json:"payments,omitempty"`
	Receipts *float32 `json:"receipts,omitempty"`
	StampBalance []ResponseStampCatBal `json:"stampBalance,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransitCash *float32 `json:"transit_cash,omitempty"`
	TransitStamps *float32 `json:"transit_stamps,omitempty"`
}

// NewResponseBodaData instantiates a new ResponseBodaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseBodaData() *ResponseBodaData {
	this := ResponseBodaData{}
	return &this
}

// NewResponseBodaDataWithDefaults instantiates a new ResponseBodaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseBodaDataWithDefaults() *ResponseBodaData {
	this := ResponseBodaData{}
	return &this
}

// GetAccountingDetail returns the AccountingDetail field value if set, zero value otherwise.
func (o *ResponseBodaData) GetAccountingDetail() []ResponseAccountingDetails {
	if o == nil || IsNil(o.AccountingDetail) {
		var ret []ResponseAccountingDetails
		return ret
	}
	return o.AccountingDetail
}

// GetAccountingDetailOk returns a tuple with the AccountingDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetAccountingDetailOk() ([]ResponseAccountingDetails, bool) {
	if o == nil || IsNil(o.AccountingDetail) {
		return nil, false
	}
	return o.AccountingDetail, true
}

// HasAccountingDetail returns a boolean if a field has been set.
func (o *ResponseBodaData) HasAccountingDetail() bool {
	if o != nil && !IsNil(o.AccountingDetail) {
		return true
	}

	return false
}

// SetAccountingDetail gets a reference to the given []ResponseAccountingDetails and assigns it to the AccountingDetail field.
func (o *ResponseBodaData) SetAccountingDetail(v []ResponseAccountingDetails) {
	o.AccountingDetail = v
}

// GetCashBagDetail returns the CashBagDetail field value if set, zero value otherwise.
func (o *ResponseBodaData) GetCashBagDetail() []ResponseCashBagDetail {
	if o == nil || IsNil(o.CashBagDetail) {
		var ret []ResponseCashBagDetail
		return ret
	}
	return o.CashBagDetail
}

// GetCashBagDetailOk returns a tuple with the CashBagDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetCashBagDetailOk() ([]ResponseCashBagDetail, bool) {
	if o == nil || IsNil(o.CashBagDetail) {
		return nil, false
	}
	return o.CashBagDetail, true
}

// HasCashBagDetail returns a boolean if a field has been set.
func (o *ResponseBodaData) HasCashBagDetail() bool {
	if o != nil && !IsNil(o.CashBagDetail) {
		return true
	}

	return false
}

// SetCashBagDetail gets a reference to the given []ResponseCashBagDetail and assigns it to the CashBagDetail field.
func (o *ResponseBodaData) SetCashBagDetail(v []ResponseCashBagDetail) {
	o.CashBagDetail = v
}

// GetClosingBal returns the ClosingBal field value if set, zero value otherwise.
func (o *ResponseBodaData) GetClosingBal() float32 {
	if o == nil || IsNil(o.ClosingBal) {
		var ret float32
		return ret
	}
	return *o.ClosingBal
}

// GetClosingBalOk returns a tuple with the ClosingBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetClosingBalOk() (*float32, bool) {
	if o == nil || IsNil(o.ClosingBal) {
		return nil, false
	}
	return o.ClosingBal, true
}

// HasClosingBal returns a boolean if a field has been set.
func (o *ResponseBodaData) HasClosingBal() bool {
	if o != nil && !IsNil(o.ClosingBal) {
		return true
	}

	return false
}

// SetClosingBal gets a reference to the given float32 and assigns it to the ClosingBal field.
func (o *ResponseBodaData) SetClosingBal(v float32) {
	o.ClosingBal = &v
}

// GetDayendStatus returns the DayendStatus field value if set, zero value otherwise.
func (o *ResponseBodaData) GetDayendStatus() int32 {
	if o == nil || IsNil(o.DayendStatus) {
		var ret int32
		return ret
	}
	return *o.DayendStatus
}

// GetDayendStatusOk returns a tuple with the DayendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetDayendStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.DayendStatus) {
		return nil, false
	}
	return o.DayendStatus, true
}

// HasDayendStatus returns a boolean if a field has been set.
func (o *ResponseBodaData) HasDayendStatus() bool {
	if o != nil && !IsNil(o.DayendStatus) {
		return true
	}

	return false
}

// SetDayendStatus gets a reference to the given int32 and assigns it to the DayendStatus field.
func (o *ResponseBodaData) SetDayendStatus(v int32) {
	o.DayendStatus = &v
}

// GetIpobalance returns the Ipobalance field value if set, zero value otherwise.
func (o *ResponseBodaData) GetIpobalance() []ResponseIPOBal {
	if o == nil || IsNil(o.Ipobalance) {
		var ret []ResponseIPOBal
		return ret
	}
	return o.Ipobalance
}

// GetIpobalanceOk returns a tuple with the Ipobalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetIpobalanceOk() ([]ResponseIPOBal, bool) {
	if o == nil || IsNil(o.Ipobalance) {
		return nil, false
	}
	return o.Ipobalance, true
}

// HasIpobalance returns a boolean if a field has been set.
func (o *ResponseBodaData) HasIpobalance() bool {
	if o != nil && !IsNil(o.Ipobalance) {
		return true
	}

	return false
}

// SetIpobalance gets a reference to the given []ResponseIPOBal and assigns it to the Ipobalance field.
func (o *ResponseBodaData) SetIpobalance(v []ResponseIPOBal) {
	o.Ipobalance = v
}

// GetMaxBal returns the MaxBal field value if set, zero value otherwise.
func (o *ResponseBodaData) GetMaxBal() float32 {
	if o == nil || IsNil(o.MaxBal) {
		var ret float32
		return ret
	}
	return *o.MaxBal
}

// GetMaxBalOk returns a tuple with the MaxBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetMaxBalOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxBal) {
		return nil, false
	}
	return o.MaxBal, true
}

// HasMaxBal returns a boolean if a field has been set.
func (o *ResponseBodaData) HasMaxBal() bool {
	if o != nil && !IsNil(o.MaxBal) {
		return true
	}

	return false
}

// SetMaxBal gets a reference to the given float32 and assigns it to the MaxBal field.
func (o *ResponseBodaData) SetMaxBal(v float32) {
	o.MaxBal = &v
}

// GetMinBal returns the MinBal field value if set, zero value otherwise.
func (o *ResponseBodaData) GetMinBal() float32 {
	if o == nil || IsNil(o.MinBal) {
		var ret float32
		return ret
	}
	return *o.MinBal
}

// GetMinBalOk returns a tuple with the MinBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetMinBalOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBal) {
		return nil, false
	}
	return o.MinBal, true
}

// HasMinBal returns a boolean if a field has been set.
func (o *ResponseBodaData) HasMinBal() bool {
	if o != nil && !IsNil(o.MinBal) {
		return true
	}

	return false
}

// SetMinBal gets a reference to the given float32 and assigns it to the MinBal field.
func (o *ResponseBodaData) SetMinBal(v float32) {
	o.MinBal = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseBodaData) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseBodaData) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseBodaData) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOpeningBal returns the OpeningBal field value if set, zero value otherwise.
func (o *ResponseBodaData) GetOpeningBal() float32 {
	if o == nil || IsNil(o.OpeningBal) {
		var ret float32
		return ret
	}
	return *o.OpeningBal
}

// GetOpeningBalOk returns a tuple with the OpeningBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetOpeningBalOk() (*float32, bool) {
	if o == nil || IsNil(o.OpeningBal) {
		return nil, false
	}
	return o.OpeningBal, true
}

// HasOpeningBal returns a boolean if a field has been set.
func (o *ResponseBodaData) HasOpeningBal() bool {
	if o != nil && !IsNil(o.OpeningBal) {
		return true
	}

	return false
}

// SetOpeningBal gets a reference to the given float32 and assigns it to the OpeningBal field.
func (o *ResponseBodaData) SetOpeningBal(v float32) {
	o.OpeningBal = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *ResponseBodaData) GetPayments() float32 {
	if o == nil || IsNil(o.Payments) {
		var ret float32
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetPaymentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *ResponseBodaData) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given float32 and assigns it to the Payments field.
func (o *ResponseBodaData) SetPayments(v float32) {
	o.Payments = &v
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *ResponseBodaData) GetReceipts() float32 {
	if o == nil || IsNil(o.Receipts) {
		var ret float32
		return ret
	}
	return *o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetReceiptsOk() (*float32, bool) {
	if o == nil || IsNil(o.Receipts) {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *ResponseBodaData) HasReceipts() bool {
	if o != nil && !IsNil(o.Receipts) {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given float32 and assigns it to the Receipts field.
func (o *ResponseBodaData) SetReceipts(v float32) {
	o.Receipts = &v
}

// GetStampBalance returns the StampBalance field value if set, zero value otherwise.
func (o *ResponseBodaData) GetStampBalance() []ResponseStampCatBal {
	if o == nil || IsNil(o.StampBalance) {
		var ret []ResponseStampCatBal
		return ret
	}
	return o.StampBalance
}

// GetStampBalanceOk returns a tuple with the StampBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetStampBalanceOk() ([]ResponseStampCatBal, bool) {
	if o == nil || IsNil(o.StampBalance) {
		return nil, false
	}
	return o.StampBalance, true
}

// HasStampBalance returns a boolean if a field has been set.
func (o *ResponseBodaData) HasStampBalance() bool {
	if o != nil && !IsNil(o.StampBalance) {
		return true
	}

	return false
}

// SetStampBalance gets a reference to the given []ResponseStampCatBal and assigns it to the StampBalance field.
func (o *ResponseBodaData) SetStampBalance(v []ResponseStampCatBal) {
	o.StampBalance = v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseBodaData) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseBodaData) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseBodaData) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransitCash returns the TransitCash field value if set, zero value otherwise.
func (o *ResponseBodaData) GetTransitCash() float32 {
	if o == nil || IsNil(o.TransitCash) {
		var ret float32
		return ret
	}
	return *o.TransitCash
}

// GetTransitCashOk returns a tuple with the TransitCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetTransitCashOk() (*float32, bool) {
	if o == nil || IsNil(o.TransitCash) {
		return nil, false
	}
	return o.TransitCash, true
}

// HasTransitCash returns a boolean if a field has been set.
func (o *ResponseBodaData) HasTransitCash() bool {
	if o != nil && !IsNil(o.TransitCash) {
		return true
	}

	return false
}

// SetTransitCash gets a reference to the given float32 and assigns it to the TransitCash field.
func (o *ResponseBodaData) SetTransitCash(v float32) {
	o.TransitCash = &v
}

// GetTransitStamps returns the TransitStamps field value if set, zero value otherwise.
func (o *ResponseBodaData) GetTransitStamps() float32 {
	if o == nil || IsNil(o.TransitStamps) {
		var ret float32
		return ret
	}
	return *o.TransitStamps
}

// GetTransitStampsOk returns a tuple with the TransitStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseBodaData) GetTransitStampsOk() (*float32, bool) {
	if o == nil || IsNil(o.TransitStamps) {
		return nil, false
	}
	return o.TransitStamps, true
}

// HasTransitStamps returns a boolean if a field has been set.
func (o *ResponseBodaData) HasTransitStamps() bool {
	if o != nil && !IsNil(o.TransitStamps) {
		return true
	}

	return false
}

// SetTransitStamps gets a reference to the given float32 and assigns it to the TransitStamps field.
func (o *ResponseBodaData) SetTransitStamps(v float32) {
	o.TransitStamps = &v
}

func (o ResponseBodaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseBodaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountingDetail) {
		toSerialize["accountingDetail"] = o.AccountingDetail
	}
	if !IsNil(o.CashBagDetail) {
		toSerialize["cashBagDetail"] = o.CashBagDetail
	}
	if !IsNil(o.ClosingBal) {
		toSerialize["closing_bal"] = o.ClosingBal
	}
	if !IsNil(o.DayendStatus) {
		toSerialize["dayend_status"] = o.DayendStatus
	}
	if !IsNil(o.Ipobalance) {
		toSerialize["ipobalance"] = o.Ipobalance
	}
	if !IsNil(o.MaxBal) {
		toSerialize["max_bal"] = o.MaxBal
	}
	if !IsNil(o.MinBal) {
		toSerialize["min_bal"] = o.MinBal
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
	if !IsNil(o.StampBalance) {
		toSerialize["stampBalance"] = o.StampBalance
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransitCash) {
		toSerialize["transit_cash"] = o.TransitCash
	}
	if !IsNil(o.TransitStamps) {
		toSerialize["transit_stamps"] = o.TransitStamps
	}
	return toSerialize, nil
}

type NullableResponseBodaData struct {
	value *ResponseBodaData
	isSet bool
}

func (v NullableResponseBodaData) Get() *ResponseBodaData {
	return v.value
}

func (v *NullableResponseBodaData) Set(val *ResponseBodaData) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseBodaData) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseBodaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseBodaData(val *ResponseBodaData) *NullableResponseBodaData {
	return &NullableResponseBodaData{value: val, isSet: true}
}

func (v NullableResponseBodaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseBodaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


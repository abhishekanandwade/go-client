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

// checks if the ResponseSummaryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseSummaryData{}

// ResponseSummaryData struct for ResponseSummaryData
type ResponseSummaryData struct {
	AccountingDetails []ResponseAccountingDetails `json:"accountingDetails,omitempty"`
	ChequeIssueDetails []ResponseChequeDetails `json:"cheque_issue_details,omitempty"`
	ChequeRemittanceDetails []ResponseChequeDetails `json:"cheque_remittance_details,omitempty"`
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
	SubOrdinateOfficeSummarydetails []ResponseSubOrdinateOfficeSummarydetails `json:"subOrdinateOfficeSummarydetails,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransitCash *float32 `json:"transit_cash,omitempty"`
	TransitIpos *float32 `json:"transit_ipos,omitempty"`
	TransitStamps *float32 `json:"transit_stamps,omitempty"`
}

// NewResponseSummaryData instantiates a new ResponseSummaryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseSummaryData() *ResponseSummaryData {
	this := ResponseSummaryData{}
	return &this
}

// NewResponseSummaryDataWithDefaults instantiates a new ResponseSummaryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseSummaryDataWithDefaults() *ResponseSummaryData {
	this := ResponseSummaryData{}
	return &this
}

// GetAccountingDetails returns the AccountingDetails field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetAccountingDetails() []ResponseAccountingDetails {
	if o == nil || IsNil(o.AccountingDetails) {
		var ret []ResponseAccountingDetails
		return ret
	}
	return o.AccountingDetails
}

// GetAccountingDetailsOk returns a tuple with the AccountingDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetAccountingDetailsOk() ([]ResponseAccountingDetails, bool) {
	if o == nil || IsNil(o.AccountingDetails) {
		return nil, false
	}
	return o.AccountingDetails, true
}

// HasAccountingDetails returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasAccountingDetails() bool {
	if o != nil && !IsNil(o.AccountingDetails) {
		return true
	}

	return false
}

// SetAccountingDetails gets a reference to the given []ResponseAccountingDetails and assigns it to the AccountingDetails field.
func (o *ResponseSummaryData) SetAccountingDetails(v []ResponseAccountingDetails) {
	o.AccountingDetails = v
}

// GetChequeIssueDetails returns the ChequeIssueDetails field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetChequeIssueDetails() []ResponseChequeDetails {
	if o == nil || IsNil(o.ChequeIssueDetails) {
		var ret []ResponseChequeDetails
		return ret
	}
	return o.ChequeIssueDetails
}

// GetChequeIssueDetailsOk returns a tuple with the ChequeIssueDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetChequeIssueDetailsOk() ([]ResponseChequeDetails, bool) {
	if o == nil || IsNil(o.ChequeIssueDetails) {
		return nil, false
	}
	return o.ChequeIssueDetails, true
}

// HasChequeIssueDetails returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasChequeIssueDetails() bool {
	if o != nil && !IsNil(o.ChequeIssueDetails) {
		return true
	}

	return false
}

// SetChequeIssueDetails gets a reference to the given []ResponseChequeDetails and assigns it to the ChequeIssueDetails field.
func (o *ResponseSummaryData) SetChequeIssueDetails(v []ResponseChequeDetails) {
	o.ChequeIssueDetails = v
}

// GetChequeRemittanceDetails returns the ChequeRemittanceDetails field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetChequeRemittanceDetails() []ResponseChequeDetails {
	if o == nil || IsNil(o.ChequeRemittanceDetails) {
		var ret []ResponseChequeDetails
		return ret
	}
	return o.ChequeRemittanceDetails
}

// GetChequeRemittanceDetailsOk returns a tuple with the ChequeRemittanceDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetChequeRemittanceDetailsOk() ([]ResponseChequeDetails, bool) {
	if o == nil || IsNil(o.ChequeRemittanceDetails) {
		return nil, false
	}
	return o.ChequeRemittanceDetails, true
}

// HasChequeRemittanceDetails returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasChequeRemittanceDetails() bool {
	if o != nil && !IsNil(o.ChequeRemittanceDetails) {
		return true
	}

	return false
}

// SetChequeRemittanceDetails gets a reference to the given []ResponseChequeDetails and assigns it to the ChequeRemittanceDetails field.
func (o *ResponseSummaryData) SetChequeRemittanceDetails(v []ResponseChequeDetails) {
	o.ChequeRemittanceDetails = v
}

// GetClosingBal returns the ClosingBal field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetClosingBal() float32 {
	if o == nil || IsNil(o.ClosingBal) {
		var ret float32
		return ret
	}
	return *o.ClosingBal
}

// GetClosingBalOk returns a tuple with the ClosingBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetClosingBalOk() (*float32, bool) {
	if o == nil || IsNil(o.ClosingBal) {
		return nil, false
	}
	return o.ClosingBal, true
}

// HasClosingBal returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasClosingBal() bool {
	if o != nil && !IsNil(o.ClosingBal) {
		return true
	}

	return false
}

// SetClosingBal gets a reference to the given float32 and assigns it to the ClosingBal field.
func (o *ResponseSummaryData) SetClosingBal(v float32) {
	o.ClosingBal = &v
}

// GetDayendStatus returns the DayendStatus field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetDayendStatus() int32 {
	if o == nil || IsNil(o.DayendStatus) {
		var ret int32
		return ret
	}
	return *o.DayendStatus
}

// GetDayendStatusOk returns a tuple with the DayendStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetDayendStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.DayendStatus) {
		return nil, false
	}
	return o.DayendStatus, true
}

// HasDayendStatus returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasDayendStatus() bool {
	if o != nil && !IsNil(o.DayendStatus) {
		return true
	}

	return false
}

// SetDayendStatus gets a reference to the given int32 and assigns it to the DayendStatus field.
func (o *ResponseSummaryData) SetDayendStatus(v int32) {
	o.DayendStatus = &v
}

// GetIpobalance returns the Ipobalance field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetIpobalance() []ResponseIPOBal {
	if o == nil || IsNil(o.Ipobalance) {
		var ret []ResponseIPOBal
		return ret
	}
	return o.Ipobalance
}

// GetIpobalanceOk returns a tuple with the Ipobalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetIpobalanceOk() ([]ResponseIPOBal, bool) {
	if o == nil || IsNil(o.Ipobalance) {
		return nil, false
	}
	return o.Ipobalance, true
}

// HasIpobalance returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasIpobalance() bool {
	if o != nil && !IsNil(o.Ipobalance) {
		return true
	}

	return false
}

// SetIpobalance gets a reference to the given []ResponseIPOBal and assigns it to the Ipobalance field.
func (o *ResponseSummaryData) SetIpobalance(v []ResponseIPOBal) {
	o.Ipobalance = v
}

// GetMaxBal returns the MaxBal field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetMaxBal() float32 {
	if o == nil || IsNil(o.MaxBal) {
		var ret float32
		return ret
	}
	return *o.MaxBal
}

// GetMaxBalOk returns a tuple with the MaxBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetMaxBalOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxBal) {
		return nil, false
	}
	return o.MaxBal, true
}

// HasMaxBal returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasMaxBal() bool {
	if o != nil && !IsNil(o.MaxBal) {
		return true
	}

	return false
}

// SetMaxBal gets a reference to the given float32 and assigns it to the MaxBal field.
func (o *ResponseSummaryData) SetMaxBal(v float32) {
	o.MaxBal = &v
}

// GetMinBal returns the MinBal field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetMinBal() float32 {
	if o == nil || IsNil(o.MinBal) {
		var ret float32
		return ret
	}
	return *o.MinBal
}

// GetMinBalOk returns a tuple with the MinBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetMinBalOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBal) {
		return nil, false
	}
	return o.MinBal, true
}

// HasMinBal returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasMinBal() bool {
	if o != nil && !IsNil(o.MinBal) {
		return true
	}

	return false
}

// SetMinBal gets a reference to the given float32 and assigns it to the MinBal field.
func (o *ResponseSummaryData) SetMinBal(v float32) {
	o.MinBal = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseSummaryData) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOpeningBal returns the OpeningBal field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetOpeningBal() float32 {
	if o == nil || IsNil(o.OpeningBal) {
		var ret float32
		return ret
	}
	return *o.OpeningBal
}

// GetOpeningBalOk returns a tuple with the OpeningBal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetOpeningBalOk() (*float32, bool) {
	if o == nil || IsNil(o.OpeningBal) {
		return nil, false
	}
	return o.OpeningBal, true
}

// HasOpeningBal returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasOpeningBal() bool {
	if o != nil && !IsNil(o.OpeningBal) {
		return true
	}

	return false
}

// SetOpeningBal gets a reference to the given float32 and assigns it to the OpeningBal field.
func (o *ResponseSummaryData) SetOpeningBal(v float32) {
	o.OpeningBal = &v
}

// GetPayments returns the Payments field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetPayments() float32 {
	if o == nil || IsNil(o.Payments) {
		var ret float32
		return ret
	}
	return *o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetPaymentsOk() (*float32, bool) {
	if o == nil || IsNil(o.Payments) {
		return nil, false
	}
	return o.Payments, true
}

// HasPayments returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasPayments() bool {
	if o != nil && !IsNil(o.Payments) {
		return true
	}

	return false
}

// SetPayments gets a reference to the given float32 and assigns it to the Payments field.
func (o *ResponseSummaryData) SetPayments(v float32) {
	o.Payments = &v
}

// GetReceipts returns the Receipts field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetReceipts() float32 {
	if o == nil || IsNil(o.Receipts) {
		var ret float32
		return ret
	}
	return *o.Receipts
}

// GetReceiptsOk returns a tuple with the Receipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetReceiptsOk() (*float32, bool) {
	if o == nil || IsNil(o.Receipts) {
		return nil, false
	}
	return o.Receipts, true
}

// HasReceipts returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasReceipts() bool {
	if o != nil && !IsNil(o.Receipts) {
		return true
	}

	return false
}

// SetReceipts gets a reference to the given float32 and assigns it to the Receipts field.
func (o *ResponseSummaryData) SetReceipts(v float32) {
	o.Receipts = &v
}

// GetStampBalance returns the StampBalance field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetStampBalance() []ResponseStampCatBal {
	if o == nil || IsNil(o.StampBalance) {
		var ret []ResponseStampCatBal
		return ret
	}
	return o.StampBalance
}

// GetStampBalanceOk returns a tuple with the StampBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetStampBalanceOk() ([]ResponseStampCatBal, bool) {
	if o == nil || IsNil(o.StampBalance) {
		return nil, false
	}
	return o.StampBalance, true
}

// HasStampBalance returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasStampBalance() bool {
	if o != nil && !IsNil(o.StampBalance) {
		return true
	}

	return false
}

// SetStampBalance gets a reference to the given []ResponseStampCatBal and assigns it to the StampBalance field.
func (o *ResponseSummaryData) SetStampBalance(v []ResponseStampCatBal) {
	o.StampBalance = v
}

// GetSubOrdinateOfficeSummarydetails returns the SubOrdinateOfficeSummarydetails field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetSubOrdinateOfficeSummarydetails() []ResponseSubOrdinateOfficeSummarydetails {
	if o == nil || IsNil(o.SubOrdinateOfficeSummarydetails) {
		var ret []ResponseSubOrdinateOfficeSummarydetails
		return ret
	}
	return o.SubOrdinateOfficeSummarydetails
}

// GetSubOrdinateOfficeSummarydetailsOk returns a tuple with the SubOrdinateOfficeSummarydetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetSubOrdinateOfficeSummarydetailsOk() ([]ResponseSubOrdinateOfficeSummarydetails, bool) {
	if o == nil || IsNil(o.SubOrdinateOfficeSummarydetails) {
		return nil, false
	}
	return o.SubOrdinateOfficeSummarydetails, true
}

// HasSubOrdinateOfficeSummarydetails returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasSubOrdinateOfficeSummarydetails() bool {
	if o != nil && !IsNil(o.SubOrdinateOfficeSummarydetails) {
		return true
	}

	return false
}

// SetSubOrdinateOfficeSummarydetails gets a reference to the given []ResponseSubOrdinateOfficeSummarydetails and assigns it to the SubOrdinateOfficeSummarydetails field.
func (o *ResponseSummaryData) SetSubOrdinateOfficeSummarydetails(v []ResponseSubOrdinateOfficeSummarydetails) {
	o.SubOrdinateOfficeSummarydetails = v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseSummaryData) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransitCash returns the TransitCash field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetTransitCash() float32 {
	if o == nil || IsNil(o.TransitCash) {
		var ret float32
		return ret
	}
	return *o.TransitCash
}

// GetTransitCashOk returns a tuple with the TransitCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetTransitCashOk() (*float32, bool) {
	if o == nil || IsNil(o.TransitCash) {
		return nil, false
	}
	return o.TransitCash, true
}

// HasTransitCash returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasTransitCash() bool {
	if o != nil && !IsNil(o.TransitCash) {
		return true
	}

	return false
}

// SetTransitCash gets a reference to the given float32 and assigns it to the TransitCash field.
func (o *ResponseSummaryData) SetTransitCash(v float32) {
	o.TransitCash = &v
}

// GetTransitIpos returns the TransitIpos field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetTransitIpos() float32 {
	if o == nil || IsNil(o.TransitIpos) {
		var ret float32
		return ret
	}
	return *o.TransitIpos
}

// GetTransitIposOk returns a tuple with the TransitIpos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetTransitIposOk() (*float32, bool) {
	if o == nil || IsNil(o.TransitIpos) {
		return nil, false
	}
	return o.TransitIpos, true
}

// HasTransitIpos returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasTransitIpos() bool {
	if o != nil && !IsNil(o.TransitIpos) {
		return true
	}

	return false
}

// SetTransitIpos gets a reference to the given float32 and assigns it to the TransitIpos field.
func (o *ResponseSummaryData) SetTransitIpos(v float32) {
	o.TransitIpos = &v
}

// GetTransitStamps returns the TransitStamps field value if set, zero value otherwise.
func (o *ResponseSummaryData) GetTransitStamps() float32 {
	if o == nil || IsNil(o.TransitStamps) {
		var ret float32
		return ret
	}
	return *o.TransitStamps
}

// GetTransitStampsOk returns a tuple with the TransitStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseSummaryData) GetTransitStampsOk() (*float32, bool) {
	if o == nil || IsNil(o.TransitStamps) {
		return nil, false
	}
	return o.TransitStamps, true
}

// HasTransitStamps returns a boolean if a field has been set.
func (o *ResponseSummaryData) HasTransitStamps() bool {
	if o != nil && !IsNil(o.TransitStamps) {
		return true
	}

	return false
}

// SetTransitStamps gets a reference to the given float32 and assigns it to the TransitStamps field.
func (o *ResponseSummaryData) SetTransitStamps(v float32) {
	o.TransitStamps = &v
}

func (o ResponseSummaryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseSummaryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountingDetails) {
		toSerialize["accountingDetails"] = o.AccountingDetails
	}
	if !IsNil(o.ChequeIssueDetails) {
		toSerialize["cheque_issue_details"] = o.ChequeIssueDetails
	}
	if !IsNil(o.ChequeRemittanceDetails) {
		toSerialize["cheque_remittance_details"] = o.ChequeRemittanceDetails
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
	if !IsNil(o.SubOrdinateOfficeSummarydetails) {
		toSerialize["subOrdinateOfficeSummarydetails"] = o.SubOrdinateOfficeSummarydetails
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransitCash) {
		toSerialize["transit_cash"] = o.TransitCash
	}
	if !IsNil(o.TransitIpos) {
		toSerialize["transit_ipos"] = o.TransitIpos
	}
	if !IsNil(o.TransitStamps) {
		toSerialize["transit_stamps"] = o.TransitStamps
	}
	return toSerialize, nil
}

type NullableResponseSummaryData struct {
	value *ResponseSummaryData
	isSet bool
}

func (v NullableResponseSummaryData) Get() *ResponseSummaryData {
	return v.value
}

func (v *NullableResponseSummaryData) Set(val *ResponseSummaryData) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseSummaryData) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseSummaryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseSummaryData(val *ResponseSummaryData) *NullableResponseSummaryData {
	return &NullableResponseSummaryData{value: val, isSet: true}
}

func (v NullableResponseSummaryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseSummaryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


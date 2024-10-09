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

// checks if the ResponseAccountingDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAccountingDetails{}

// ResponseAccountingDetails struct for ResponseAccountingDetails
type ResponseAccountingDetails struct {
	AccountCodeDescription *string `json:"account_code_description,omitempty"`
	CreditOrDebit *string `json:"credit_or_debit,omitempty"`
	DigitalTxnsAmount *float32 `json:"digital_txns_amount,omitempty"`
	DigitalTxnsCount *int32 `json:"digital_txns_count,omitempty"`
	GlCode *string `json:"gl_code,omitempty"`
	Hoa *string `json:"hoa,omitempty"`
	NonDigitalTxnsAmount *float32 `json:"non_digital_txns_amount,omitempty"`
	NonDigitalTxnsCount *int32 `json:"non_digital_txns_count,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	Part *string `json:"part,omitempty"`
	PositiveOrNegative *string `json:"positive_or_negative,omitempty"`
	ReceiptOrPmt *string `json:"receipt_or_pmt,omitempty"`
	ReceiptSource *string `json:"receipt_source,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	SrcTranId *string `json:"src_tran_id,omitempty"`
	SrcTransDate *string `json:"src_trans_date,omitempty"`
	TotAmount *float32 `json:"tot_amount,omitempty"`
	TotTxns *int32 `json:"tot_txns,omitempty"`
	TreasuryTranId *string `json:"treasury_tran_id,omitempty"`
	TryPostingDate *string `json:"try_posting_date,omitempty"`
}

// NewResponseAccountingDetails instantiates a new ResponseAccountingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAccountingDetails() *ResponseAccountingDetails {
	this := ResponseAccountingDetails{}
	return &this
}

// NewResponseAccountingDetailsWithDefaults instantiates a new ResponseAccountingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAccountingDetailsWithDefaults() *ResponseAccountingDetails {
	this := ResponseAccountingDetails{}
	return &this
}

// GetAccountCodeDescription returns the AccountCodeDescription field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetAccountCodeDescription() string {
	if o == nil || IsNil(o.AccountCodeDescription) {
		var ret string
		return ret
	}
	return *o.AccountCodeDescription
}

// GetAccountCodeDescriptionOk returns a tuple with the AccountCodeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetAccountCodeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCodeDescription) {
		return nil, false
	}
	return o.AccountCodeDescription, true
}

// HasAccountCodeDescription returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasAccountCodeDescription() bool {
	if o != nil && !IsNil(o.AccountCodeDescription) {
		return true
	}

	return false
}

// SetAccountCodeDescription gets a reference to the given string and assigns it to the AccountCodeDescription field.
func (o *ResponseAccountingDetails) SetAccountCodeDescription(v string) {
	o.AccountCodeDescription = &v
}

// GetCreditOrDebit returns the CreditOrDebit field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetCreditOrDebit() string {
	if o == nil || IsNil(o.CreditOrDebit) {
		var ret string
		return ret
	}
	return *o.CreditOrDebit
}

// GetCreditOrDebitOk returns a tuple with the CreditOrDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetCreditOrDebitOk() (*string, bool) {
	if o == nil || IsNil(o.CreditOrDebit) {
		return nil, false
	}
	return o.CreditOrDebit, true
}

// HasCreditOrDebit returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasCreditOrDebit() bool {
	if o != nil && !IsNil(o.CreditOrDebit) {
		return true
	}

	return false
}

// SetCreditOrDebit gets a reference to the given string and assigns it to the CreditOrDebit field.
func (o *ResponseAccountingDetails) SetCreditOrDebit(v string) {
	o.CreditOrDebit = &v
}

// GetDigitalTxnsAmount returns the DigitalTxnsAmount field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetDigitalTxnsAmount() float32 {
	if o == nil || IsNil(o.DigitalTxnsAmount) {
		var ret float32
		return ret
	}
	return *o.DigitalTxnsAmount
}

// GetDigitalTxnsAmountOk returns a tuple with the DigitalTxnsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetDigitalTxnsAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.DigitalTxnsAmount) {
		return nil, false
	}
	return o.DigitalTxnsAmount, true
}

// HasDigitalTxnsAmount returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasDigitalTxnsAmount() bool {
	if o != nil && !IsNil(o.DigitalTxnsAmount) {
		return true
	}

	return false
}

// SetDigitalTxnsAmount gets a reference to the given float32 and assigns it to the DigitalTxnsAmount field.
func (o *ResponseAccountingDetails) SetDigitalTxnsAmount(v float32) {
	o.DigitalTxnsAmount = &v
}

// GetDigitalTxnsCount returns the DigitalTxnsCount field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetDigitalTxnsCount() int32 {
	if o == nil || IsNil(o.DigitalTxnsCount) {
		var ret int32
		return ret
	}
	return *o.DigitalTxnsCount
}

// GetDigitalTxnsCountOk returns a tuple with the DigitalTxnsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetDigitalTxnsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DigitalTxnsCount) {
		return nil, false
	}
	return o.DigitalTxnsCount, true
}

// HasDigitalTxnsCount returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasDigitalTxnsCount() bool {
	if o != nil && !IsNil(o.DigitalTxnsCount) {
		return true
	}

	return false
}

// SetDigitalTxnsCount gets a reference to the given int32 and assigns it to the DigitalTxnsCount field.
func (o *ResponseAccountingDetails) SetDigitalTxnsCount(v int32) {
	o.DigitalTxnsCount = &v
}

// GetGlCode returns the GlCode field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetGlCode() string {
	if o == nil || IsNil(o.GlCode) {
		var ret string
		return ret
	}
	return *o.GlCode
}

// GetGlCodeOk returns a tuple with the GlCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetGlCodeOk() (*string, bool) {
	if o == nil || IsNil(o.GlCode) {
		return nil, false
	}
	return o.GlCode, true
}

// HasGlCode returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasGlCode() bool {
	if o != nil && !IsNil(o.GlCode) {
		return true
	}

	return false
}

// SetGlCode gets a reference to the given string and assigns it to the GlCode field.
func (o *ResponseAccountingDetails) SetGlCode(v string) {
	o.GlCode = &v
}

// GetHoa returns the Hoa field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetHoa() string {
	if o == nil || IsNil(o.Hoa) {
		var ret string
		return ret
	}
	return *o.Hoa
}

// GetHoaOk returns a tuple with the Hoa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetHoaOk() (*string, bool) {
	if o == nil || IsNil(o.Hoa) {
		return nil, false
	}
	return o.Hoa, true
}

// HasHoa returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasHoa() bool {
	if o != nil && !IsNil(o.Hoa) {
		return true
	}

	return false
}

// SetHoa gets a reference to the given string and assigns it to the Hoa field.
func (o *ResponseAccountingDetails) SetHoa(v string) {
	o.Hoa = &v
}

// GetNonDigitalTxnsAmount returns the NonDigitalTxnsAmount field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetNonDigitalTxnsAmount() float32 {
	if o == nil || IsNil(o.NonDigitalTxnsAmount) {
		var ret float32
		return ret
	}
	return *o.NonDigitalTxnsAmount
}

// GetNonDigitalTxnsAmountOk returns a tuple with the NonDigitalTxnsAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetNonDigitalTxnsAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.NonDigitalTxnsAmount) {
		return nil, false
	}
	return o.NonDigitalTxnsAmount, true
}

// HasNonDigitalTxnsAmount returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasNonDigitalTxnsAmount() bool {
	if o != nil && !IsNil(o.NonDigitalTxnsAmount) {
		return true
	}

	return false
}

// SetNonDigitalTxnsAmount gets a reference to the given float32 and assigns it to the NonDigitalTxnsAmount field.
func (o *ResponseAccountingDetails) SetNonDigitalTxnsAmount(v float32) {
	o.NonDigitalTxnsAmount = &v
}

// GetNonDigitalTxnsCount returns the NonDigitalTxnsCount field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetNonDigitalTxnsCount() int32 {
	if o == nil || IsNil(o.NonDigitalTxnsCount) {
		var ret int32
		return ret
	}
	return *o.NonDigitalTxnsCount
}

// GetNonDigitalTxnsCountOk returns a tuple with the NonDigitalTxnsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetNonDigitalTxnsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NonDigitalTxnsCount) {
		return nil, false
	}
	return o.NonDigitalTxnsCount, true
}

// HasNonDigitalTxnsCount returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasNonDigitalTxnsCount() bool {
	if o != nil && !IsNil(o.NonDigitalTxnsCount) {
		return true
	}

	return false
}

// SetNonDigitalTxnsCount gets a reference to the given int32 and assigns it to the NonDigitalTxnsCount field.
func (o *ResponseAccountingDetails) SetNonDigitalTxnsCount(v int32) {
	o.NonDigitalTxnsCount = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseAccountingDetails) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *ResponseAccountingDetails) SetPart(v string) {
	o.Part = &v
}

// GetPositiveOrNegative returns the PositiveOrNegative field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetPositiveOrNegative() string {
	if o == nil || IsNil(o.PositiveOrNegative) {
		var ret string
		return ret
	}
	return *o.PositiveOrNegative
}

// GetPositiveOrNegativeOk returns a tuple with the PositiveOrNegative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetPositiveOrNegativeOk() (*string, bool) {
	if o == nil || IsNil(o.PositiveOrNegative) {
		return nil, false
	}
	return o.PositiveOrNegative, true
}

// HasPositiveOrNegative returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasPositiveOrNegative() bool {
	if o != nil && !IsNil(o.PositiveOrNegative) {
		return true
	}

	return false
}

// SetPositiveOrNegative gets a reference to the given string and assigns it to the PositiveOrNegative field.
func (o *ResponseAccountingDetails) SetPositiveOrNegative(v string) {
	o.PositiveOrNegative = &v
}

// GetReceiptOrPmt returns the ReceiptOrPmt field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetReceiptOrPmt() string {
	if o == nil || IsNil(o.ReceiptOrPmt) {
		var ret string
		return ret
	}
	return *o.ReceiptOrPmt
}

// GetReceiptOrPmtOk returns a tuple with the ReceiptOrPmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetReceiptOrPmtOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptOrPmt) {
		return nil, false
	}
	return o.ReceiptOrPmt, true
}

// HasReceiptOrPmt returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasReceiptOrPmt() bool {
	if o != nil && !IsNil(o.ReceiptOrPmt) {
		return true
	}

	return false
}

// SetReceiptOrPmt gets a reference to the given string and assigns it to the ReceiptOrPmt field.
func (o *ResponseAccountingDetails) SetReceiptOrPmt(v string) {
	o.ReceiptOrPmt = &v
}

// GetReceiptSource returns the ReceiptSource field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetReceiptSource() string {
	if o == nil || IsNil(o.ReceiptSource) {
		var ret string
		return ret
	}
	return *o.ReceiptSource
}

// GetReceiptSourceOk returns a tuple with the ReceiptSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetReceiptSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptSource) {
		return nil, false
	}
	return o.ReceiptSource, true
}

// HasReceiptSource returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasReceiptSource() bool {
	if o != nil && !IsNil(o.ReceiptSource) {
		return true
	}

	return false
}

// SetReceiptSource gets a reference to the given string and assigns it to the ReceiptSource field.
func (o *ResponseAccountingDetails) SetReceiptSource(v string) {
	o.ReceiptSource = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseAccountingDetails) SetRemarks(v string) {
	o.Remarks = &v
}

// GetSrcTranId returns the SrcTranId field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetSrcTranId() string {
	if o == nil || IsNil(o.SrcTranId) {
		var ret string
		return ret
	}
	return *o.SrcTranId
}

// GetSrcTranIdOk returns a tuple with the SrcTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetSrcTranIdOk() (*string, bool) {
	if o == nil || IsNil(o.SrcTranId) {
		return nil, false
	}
	return o.SrcTranId, true
}

// HasSrcTranId returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasSrcTranId() bool {
	if o != nil && !IsNil(o.SrcTranId) {
		return true
	}

	return false
}

// SetSrcTranId gets a reference to the given string and assigns it to the SrcTranId field.
func (o *ResponseAccountingDetails) SetSrcTranId(v string) {
	o.SrcTranId = &v
}

// GetSrcTransDate returns the SrcTransDate field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetSrcTransDate() string {
	if o == nil || IsNil(o.SrcTransDate) {
		var ret string
		return ret
	}
	return *o.SrcTransDate
}

// GetSrcTransDateOk returns a tuple with the SrcTransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetSrcTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.SrcTransDate) {
		return nil, false
	}
	return o.SrcTransDate, true
}

// HasSrcTransDate returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasSrcTransDate() bool {
	if o != nil && !IsNil(o.SrcTransDate) {
		return true
	}

	return false
}

// SetSrcTransDate gets a reference to the given string and assigns it to the SrcTransDate field.
func (o *ResponseAccountingDetails) SetSrcTransDate(v string) {
	o.SrcTransDate = &v
}

// GetTotAmount returns the TotAmount field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetTotAmount() float32 {
	if o == nil || IsNil(o.TotAmount) {
		var ret float32
		return ret
	}
	return *o.TotAmount
}

// GetTotAmountOk returns a tuple with the TotAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetTotAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotAmount) {
		return nil, false
	}
	return o.TotAmount, true
}

// HasTotAmount returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasTotAmount() bool {
	if o != nil && !IsNil(o.TotAmount) {
		return true
	}

	return false
}

// SetTotAmount gets a reference to the given float32 and assigns it to the TotAmount field.
func (o *ResponseAccountingDetails) SetTotAmount(v float32) {
	o.TotAmount = &v
}

// GetTotTxns returns the TotTxns field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetTotTxns() int32 {
	if o == nil || IsNil(o.TotTxns) {
		var ret int32
		return ret
	}
	return *o.TotTxns
}

// GetTotTxnsOk returns a tuple with the TotTxns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetTotTxnsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotTxns) {
		return nil, false
	}
	return o.TotTxns, true
}

// HasTotTxns returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasTotTxns() bool {
	if o != nil && !IsNil(o.TotTxns) {
		return true
	}

	return false
}

// SetTotTxns gets a reference to the given int32 and assigns it to the TotTxns field.
func (o *ResponseAccountingDetails) SetTotTxns(v int32) {
	o.TotTxns = &v
}

// GetTreasuryTranId returns the TreasuryTranId field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetTreasuryTranId() string {
	if o == nil || IsNil(o.TreasuryTranId) {
		var ret string
		return ret
	}
	return *o.TreasuryTranId
}

// GetTreasuryTranIdOk returns a tuple with the TreasuryTranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetTreasuryTranIdOk() (*string, bool) {
	if o == nil || IsNil(o.TreasuryTranId) {
		return nil, false
	}
	return o.TreasuryTranId, true
}

// HasTreasuryTranId returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasTreasuryTranId() bool {
	if o != nil && !IsNil(o.TreasuryTranId) {
		return true
	}

	return false
}

// SetTreasuryTranId gets a reference to the given string and assigns it to the TreasuryTranId field.
func (o *ResponseAccountingDetails) SetTreasuryTranId(v string) {
	o.TreasuryTranId = &v
}

// GetTryPostingDate returns the TryPostingDate field value if set, zero value otherwise.
func (o *ResponseAccountingDetails) GetTryPostingDate() string {
	if o == nil || IsNil(o.TryPostingDate) {
		var ret string
		return ret
	}
	return *o.TryPostingDate
}

// GetTryPostingDateOk returns a tuple with the TryPostingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountingDetails) GetTryPostingDateOk() (*string, bool) {
	if o == nil || IsNil(o.TryPostingDate) {
		return nil, false
	}
	return o.TryPostingDate, true
}

// HasTryPostingDate returns a boolean if a field has been set.
func (o *ResponseAccountingDetails) HasTryPostingDate() bool {
	if o != nil && !IsNil(o.TryPostingDate) {
		return true
	}

	return false
}

// SetTryPostingDate gets a reference to the given string and assigns it to the TryPostingDate field.
func (o *ResponseAccountingDetails) SetTryPostingDate(v string) {
	o.TryPostingDate = &v
}

func (o ResponseAccountingDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAccountingDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountCodeDescription) {
		toSerialize["account_code_description"] = o.AccountCodeDescription
	}
	if !IsNil(o.CreditOrDebit) {
		toSerialize["credit_or_debit"] = o.CreditOrDebit
	}
	if !IsNil(o.DigitalTxnsAmount) {
		toSerialize["digital_txns_amount"] = o.DigitalTxnsAmount
	}
	if !IsNil(o.DigitalTxnsCount) {
		toSerialize["digital_txns_count"] = o.DigitalTxnsCount
	}
	if !IsNil(o.GlCode) {
		toSerialize["gl_code"] = o.GlCode
	}
	if !IsNil(o.Hoa) {
		toSerialize["hoa"] = o.Hoa
	}
	if !IsNil(o.NonDigitalTxnsAmount) {
		toSerialize["non_digital_txns_amount"] = o.NonDigitalTxnsAmount
	}
	if !IsNil(o.NonDigitalTxnsCount) {
		toSerialize["non_digital_txns_count"] = o.NonDigitalTxnsCount
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !IsNil(o.PositiveOrNegative) {
		toSerialize["positive_or_negative"] = o.PositiveOrNegative
	}
	if !IsNil(o.ReceiptOrPmt) {
		toSerialize["receipt_or_pmt"] = o.ReceiptOrPmt
	}
	if !IsNil(o.ReceiptSource) {
		toSerialize["receipt_source"] = o.ReceiptSource
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.SrcTranId) {
		toSerialize["src_tran_id"] = o.SrcTranId
	}
	if !IsNil(o.SrcTransDate) {
		toSerialize["src_trans_date"] = o.SrcTransDate
	}
	if !IsNil(o.TotAmount) {
		toSerialize["tot_amount"] = o.TotAmount
	}
	if !IsNil(o.TotTxns) {
		toSerialize["tot_txns"] = o.TotTxns
	}
	if !IsNil(o.TreasuryTranId) {
		toSerialize["treasury_tran_id"] = o.TreasuryTranId
	}
	if !IsNil(o.TryPostingDate) {
		toSerialize["try_posting_date"] = o.TryPostingDate
	}
	return toSerialize, nil
}

type NullableResponseAccountingDetails struct {
	value *ResponseAccountingDetails
	isSet bool
}

func (v NullableResponseAccountingDetails) Get() *ResponseAccountingDetails {
	return v.value
}

func (v *NullableResponseAccountingDetails) Set(val *ResponseAccountingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAccountingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAccountingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAccountingDetails(val *ResponseAccountingDetails) *NullableResponseAccountingDetails {
	return &NullableResponseAccountingDetails{value: val, isSet: true}
}

func (v NullableResponseAccountingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAccountingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// checks if the HandlerStampBulkSaleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerStampBulkSaleResponse{}

// HandlerStampBulkSaleResponse struct for HandlerStampBulkSaleResponse
type HandlerStampBulkSaleResponse struct {
	ApprovedDate *string `json:"approved_date,omitempty"`
	ApproverId *int32 `json:"approver_id,omitempty"`
	BankBranch *string `json:"bank_branch,omitempty"`
	ChqRealiseDate *string `json:"chq_realise_date,omitempty"`
	InstrumentDate *string `json:"instrument_date,omitempty"`
	InstrumentNo *string `json:"instrument_no,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	IsRealized *bool `json:"is_realized,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	PaymentMode *string `json:"payment_mode,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	SaleAmt *float32 `json:"sale_amt,omitempty"`
	SaleDetails *map[string]int32 `json:"sale_details,omitempty"`
	SoldTo *string `json:"sold_to,omitempty"`
	StampDetailsDesc []HandlerStampdetails `json:"stampDetailsDesc,omitempty"`
	TranId *string `json:"tran_id,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
}

// NewHandlerStampBulkSaleResponse instantiates a new HandlerStampBulkSaleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerStampBulkSaleResponse() *HandlerStampBulkSaleResponse {
	this := HandlerStampBulkSaleResponse{}
	return &this
}

// NewHandlerStampBulkSaleResponseWithDefaults instantiates a new HandlerStampBulkSaleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerStampBulkSaleResponseWithDefaults() *HandlerStampBulkSaleResponse {
	this := HandlerStampBulkSaleResponse{}
	return &this
}

// GetApprovedDate returns the ApprovedDate field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetApprovedDate() string {
	if o == nil || IsNil(o.ApprovedDate) {
		var ret string
		return ret
	}
	return *o.ApprovedDate
}

// GetApprovedDateOk returns a tuple with the ApprovedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetApprovedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovedDate) {
		return nil, false
	}
	return o.ApprovedDate, true
}

// HasApprovedDate returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasApprovedDate() bool {
	if o != nil && !IsNil(o.ApprovedDate) {
		return true
	}

	return false
}

// SetApprovedDate gets a reference to the given string and assigns it to the ApprovedDate field.
func (o *HandlerStampBulkSaleResponse) SetApprovedDate(v string) {
	o.ApprovedDate = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetApproverId() int32 {
	if o == nil || IsNil(o.ApproverId) {
		var ret int32
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetApproverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproverId) {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasApproverId() bool {
	if o != nil && !IsNil(o.ApproverId) {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int32 and assigns it to the ApproverId field.
func (o *HandlerStampBulkSaleResponse) SetApproverId(v int32) {
	o.ApproverId = &v
}

// GetBankBranch returns the BankBranch field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetBankBranch() string {
	if o == nil || IsNil(o.BankBranch) {
		var ret string
		return ret
	}
	return *o.BankBranch
}

// GetBankBranchOk returns a tuple with the BankBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetBankBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BankBranch) {
		return nil, false
	}
	return o.BankBranch, true
}

// HasBankBranch returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasBankBranch() bool {
	if o != nil && !IsNil(o.BankBranch) {
		return true
	}

	return false
}

// SetBankBranch gets a reference to the given string and assigns it to the BankBranch field.
func (o *HandlerStampBulkSaleResponse) SetBankBranch(v string) {
	o.BankBranch = &v
}

// GetChqRealiseDate returns the ChqRealiseDate field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetChqRealiseDate() string {
	if o == nil || IsNil(o.ChqRealiseDate) {
		var ret string
		return ret
	}
	return *o.ChqRealiseDate
}

// GetChqRealiseDateOk returns a tuple with the ChqRealiseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetChqRealiseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ChqRealiseDate) {
		return nil, false
	}
	return o.ChqRealiseDate, true
}

// HasChqRealiseDate returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasChqRealiseDate() bool {
	if o != nil && !IsNil(o.ChqRealiseDate) {
		return true
	}

	return false
}

// SetChqRealiseDate gets a reference to the given string and assigns it to the ChqRealiseDate field.
func (o *HandlerStampBulkSaleResponse) SetChqRealiseDate(v string) {
	o.ChqRealiseDate = &v
}

// GetInstrumentDate returns the InstrumentDate field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetInstrumentDate() string {
	if o == nil || IsNil(o.InstrumentDate) {
		var ret string
		return ret
	}
	return *o.InstrumentDate
}

// GetInstrumentDateOk returns a tuple with the InstrumentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetInstrumentDateOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentDate) {
		return nil, false
	}
	return o.InstrumentDate, true
}

// HasInstrumentDate returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasInstrumentDate() bool {
	if o != nil && !IsNil(o.InstrumentDate) {
		return true
	}

	return false
}

// SetInstrumentDate gets a reference to the given string and assigns it to the InstrumentDate field.
func (o *HandlerStampBulkSaleResponse) SetInstrumentDate(v string) {
	o.InstrumentDate = &v
}

// GetInstrumentNo returns the InstrumentNo field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetInstrumentNo() string {
	if o == nil || IsNil(o.InstrumentNo) {
		var ret string
		return ret
	}
	return *o.InstrumentNo
}

// GetInstrumentNoOk returns a tuple with the InstrumentNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetInstrumentNoOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentNo) {
		return nil, false
	}
	return o.InstrumentNo, true
}

// HasInstrumentNo returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasInstrumentNo() bool {
	if o != nil && !IsNil(o.InstrumentNo) {
		return true
	}

	return false
}

// SetInstrumentNo gets a reference to the given string and assigns it to the InstrumentNo field.
func (o *HandlerStampBulkSaleResponse) SetInstrumentNo(v string) {
	o.InstrumentNo = &v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *HandlerStampBulkSaleResponse) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetIsRealized returns the IsRealized field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetIsRealized() bool {
	if o == nil || IsNil(o.IsRealized) {
		var ret bool
		return ret
	}
	return *o.IsRealized
}

// GetIsRealizedOk returns a tuple with the IsRealized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetIsRealizedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRealized) {
		return nil, false
	}
	return o.IsRealized, true
}

// HasIsRealized returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasIsRealized() bool {
	if o != nil && !IsNil(o.IsRealized) {
		return true
	}

	return false
}

// SetIsRealized gets a reference to the given bool and assigns it to the IsRealized field.
func (o *HandlerStampBulkSaleResponse) SetIsRealized(v bool) {
	o.IsRealized = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *HandlerStampBulkSaleResponse) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetPaymentMode returns the PaymentMode field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetPaymentMode() string {
	if o == nil || IsNil(o.PaymentMode) {
		var ret string
		return ret
	}
	return *o.PaymentMode
}

// GetPaymentModeOk returns a tuple with the PaymentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetPaymentModeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMode) {
		return nil, false
	}
	return o.PaymentMode, true
}

// HasPaymentMode returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasPaymentMode() bool {
	if o != nil && !IsNil(o.PaymentMode) {
		return true
	}

	return false
}

// SetPaymentMode gets a reference to the given string and assigns it to the PaymentMode field.
func (o *HandlerStampBulkSaleResponse) SetPaymentMode(v string) {
	o.PaymentMode = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *HandlerStampBulkSaleResponse) SetRemarks(v string) {
	o.Remarks = &v
}

// GetSaleAmt returns the SaleAmt field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetSaleAmt() float32 {
	if o == nil || IsNil(o.SaleAmt) {
		var ret float32
		return ret
	}
	return *o.SaleAmt
}

// GetSaleAmtOk returns a tuple with the SaleAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetSaleAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.SaleAmt) {
		return nil, false
	}
	return o.SaleAmt, true
}

// HasSaleAmt returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasSaleAmt() bool {
	if o != nil && !IsNil(o.SaleAmt) {
		return true
	}

	return false
}

// SetSaleAmt gets a reference to the given float32 and assigns it to the SaleAmt field.
func (o *HandlerStampBulkSaleResponse) SetSaleAmt(v float32) {
	o.SaleAmt = &v
}

// GetSaleDetails returns the SaleDetails field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetSaleDetails() map[string]int32 {
	if o == nil || IsNil(o.SaleDetails) {
		var ret map[string]int32
		return ret
	}
	return *o.SaleDetails
}

// GetSaleDetailsOk returns a tuple with the SaleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetSaleDetailsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.SaleDetails) {
		return nil, false
	}
	return o.SaleDetails, true
}

// HasSaleDetails returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasSaleDetails() bool {
	if o != nil && !IsNil(o.SaleDetails) {
		return true
	}

	return false
}

// SetSaleDetails gets a reference to the given map[string]int32 and assigns it to the SaleDetails field.
func (o *HandlerStampBulkSaleResponse) SetSaleDetails(v map[string]int32) {
	o.SaleDetails = &v
}

// GetSoldTo returns the SoldTo field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetSoldTo() string {
	if o == nil || IsNil(o.SoldTo) {
		var ret string
		return ret
	}
	return *o.SoldTo
}

// GetSoldToOk returns a tuple with the SoldTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetSoldToOk() (*string, bool) {
	if o == nil || IsNil(o.SoldTo) {
		return nil, false
	}
	return o.SoldTo, true
}

// HasSoldTo returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasSoldTo() bool {
	if o != nil && !IsNil(o.SoldTo) {
		return true
	}

	return false
}

// SetSoldTo gets a reference to the given string and assigns it to the SoldTo field.
func (o *HandlerStampBulkSaleResponse) SetSoldTo(v string) {
	o.SoldTo = &v
}

// GetStampDetailsDesc returns the StampDetailsDesc field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetStampDetailsDesc() []HandlerStampdetails {
	if o == nil || IsNil(o.StampDetailsDesc) {
		var ret []HandlerStampdetails
		return ret
	}
	return o.StampDetailsDesc
}

// GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetStampDetailsDescOk() ([]HandlerStampdetails, bool) {
	if o == nil || IsNil(o.StampDetailsDesc) {
		return nil, false
	}
	return o.StampDetailsDesc, true
}

// HasStampDetailsDesc returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasStampDetailsDesc() bool {
	if o != nil && !IsNil(o.StampDetailsDesc) {
		return true
	}

	return false
}

// SetStampDetailsDesc gets a reference to the given []HandlerStampdetails and assigns it to the StampDetailsDesc field.
func (o *HandlerStampBulkSaleResponse) SetStampDetailsDesc(v []HandlerStampdetails) {
	o.StampDetailsDesc = v
}

// GetTranId returns the TranId field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetTranId() string {
	if o == nil || IsNil(o.TranId) {
		var ret string
		return ret
	}
	return *o.TranId
}

// GetTranIdOk returns a tuple with the TranId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetTranIdOk() (*string, bool) {
	if o == nil || IsNil(o.TranId) {
		return nil, false
	}
	return o.TranId, true
}

// HasTranId returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasTranId() bool {
	if o != nil && !IsNil(o.TranId) {
		return true
	}

	return false
}

// SetTranId gets a reference to the given string and assigns it to the TranId field.
func (o *HandlerStampBulkSaleResponse) SetTranId(v string) {
	o.TranId = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *HandlerStampBulkSaleResponse) SetTransDate(v string) {
	o.TransDate = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *HandlerStampBulkSaleResponse) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerStampBulkSaleResponse) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *HandlerStampBulkSaleResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *HandlerStampBulkSaleResponse) SetUserId(v int32) {
	o.UserId = &v
}

func (o HandlerStampBulkSaleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerStampBulkSaleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprovedDate) {
		toSerialize["approved_date"] = o.ApprovedDate
	}
	if !IsNil(o.ApproverId) {
		toSerialize["approver_id"] = o.ApproverId
	}
	if !IsNil(o.BankBranch) {
		toSerialize["bank_branch"] = o.BankBranch
	}
	if !IsNil(o.ChqRealiseDate) {
		toSerialize["chq_realise_date"] = o.ChqRealiseDate
	}
	if !IsNil(o.InstrumentDate) {
		toSerialize["instrument_date"] = o.InstrumentDate
	}
	if !IsNil(o.InstrumentNo) {
		toSerialize["instrument_no"] = o.InstrumentNo
	}
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	if !IsNil(o.IsRealized) {
		toSerialize["is_realized"] = o.IsRealized
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.PaymentMode) {
		toSerialize["payment_mode"] = o.PaymentMode
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.SaleAmt) {
		toSerialize["sale_amt"] = o.SaleAmt
	}
	if !IsNil(o.SaleDetails) {
		toSerialize["sale_details"] = o.SaleDetails
	}
	if !IsNil(o.SoldTo) {
		toSerialize["sold_to"] = o.SoldTo
	}
	if !IsNil(o.StampDetailsDesc) {
		toSerialize["stampDetailsDesc"] = o.StampDetailsDesc
	}
	if !IsNil(o.TranId) {
		toSerialize["tran_id"] = o.TranId
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableHandlerStampBulkSaleResponse struct {
	value *HandlerStampBulkSaleResponse
	isSet bool
}

func (v NullableHandlerStampBulkSaleResponse) Get() *HandlerStampBulkSaleResponse {
	return v.value
}

func (v *NullableHandlerStampBulkSaleResponse) Set(val *HandlerStampBulkSaleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerStampBulkSaleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerStampBulkSaleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerStampBulkSaleResponse(val *HandlerStampBulkSaleResponse) *NullableHandlerStampBulkSaleResponse {
	return &NullableHandlerStampBulkSaleResponse{value: val, isSet: true}
}

func (v NullableHandlerStampBulkSaleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerStampBulkSaleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



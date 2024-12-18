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

// checks if the ResponseTreasuryMiscTransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTreasuryMiscTransactionResponse{}

// ResponseTreasuryMiscTransactionResponse struct for ResponseTreasuryMiscTransactionResponse
type ResponseTreasuryMiscTransactionResponse struct {
	AccountDescdetails []ResponseAccountDescdetails `json:"accountDescdetails,omitempty"`
	ApproverDate *string `json:"approver_date,omitempty"`
	ApproverId *int32 `json:"approver_id,omitempty"`
	ApproverRemarks *string `json:"approver_remarks,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	ReqAmount *float32 `json:"req_amount,omitempty"`
	ReqDetails *map[string]float32 `json:"req_details,omitempty"`
	RequestType *string `json:"request_type,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
	TxnMode *string `json:"txn_mode,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
}

// NewResponseTreasuryMiscTransactionResponse instantiates a new ResponseTreasuryMiscTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTreasuryMiscTransactionResponse() *ResponseTreasuryMiscTransactionResponse {
	this := ResponseTreasuryMiscTransactionResponse{}
	return &this
}

// NewResponseTreasuryMiscTransactionResponseWithDefaults instantiates a new ResponseTreasuryMiscTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTreasuryMiscTransactionResponseWithDefaults() *ResponseTreasuryMiscTransactionResponse {
	this := ResponseTreasuryMiscTransactionResponse{}
	return &this
}

// GetAccountDescdetails returns the AccountDescdetails field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetAccountDescdetails() []ResponseAccountDescdetails {
	if o == nil || IsNil(o.AccountDescdetails) {
		var ret []ResponseAccountDescdetails
		return ret
	}
	return o.AccountDescdetails
}

// GetAccountDescdetailsOk returns a tuple with the AccountDescdetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetAccountDescdetailsOk() ([]ResponseAccountDescdetails, bool) {
	if o == nil || IsNil(o.AccountDescdetails) {
		return nil, false
	}
	return o.AccountDescdetails, true
}

// HasAccountDescdetails returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasAccountDescdetails() bool {
	if o != nil && !IsNil(o.AccountDescdetails) {
		return true
	}

	return false
}

// SetAccountDescdetails gets a reference to the given []ResponseAccountDescdetails and assigns it to the AccountDescdetails field.
func (o *ResponseTreasuryMiscTransactionResponse) SetAccountDescdetails(v []ResponseAccountDescdetails) {
	o.AccountDescdetails = v
}

// GetApproverDate returns the ApproverDate field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverDate() string {
	if o == nil || IsNil(o.ApproverDate) {
		var ret string
		return ret
	}
	return *o.ApproverDate
}

// GetApproverDateOk returns a tuple with the ApproverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverDateOk() (*string, bool) {
	if o == nil || IsNil(o.ApproverDate) {
		return nil, false
	}
	return o.ApproverDate, true
}

// HasApproverDate returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasApproverDate() bool {
	if o != nil && !IsNil(o.ApproverDate) {
		return true
	}

	return false
}

// SetApproverDate gets a reference to the given string and assigns it to the ApproverDate field.
func (o *ResponseTreasuryMiscTransactionResponse) SetApproverDate(v string) {
	o.ApproverDate = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverId() int32 {
	if o == nil || IsNil(o.ApproverId) {
		var ret int32
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproverId) {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasApproverId() bool {
	if o != nil && !IsNil(o.ApproverId) {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int32 and assigns it to the ApproverId field.
func (o *ResponseTreasuryMiscTransactionResponse) SetApproverId(v int32) {
	o.ApproverId = &v
}

// GetApproverRemarks returns the ApproverRemarks field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverRemarks() string {
	if o == nil || IsNil(o.ApproverRemarks) {
		var ret string
		return ret
	}
	return *o.ApproverRemarks
}

// GetApproverRemarksOk returns a tuple with the ApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.ApproverRemarks) {
		return nil, false
	}
	return o.ApproverRemarks, true
}

// HasApproverRemarks returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasApproverRemarks() bool {
	if o != nil && !IsNil(o.ApproverRemarks) {
		return true
	}

	return false
}

// SetApproverRemarks gets a reference to the given string and assigns it to the ApproverRemarks field.
func (o *ResponseTreasuryMiscTransactionResponse) SetApproverRemarks(v string) {
	o.ApproverRemarks = &v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *ResponseTreasuryMiscTransactionResponse) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseTreasuryMiscTransactionResponse) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseTreasuryMiscTransactionResponse) SetRemarks(v string) {
	o.Remarks = &v
}

// GetReqAmount returns the ReqAmount field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetReqAmount() float32 {
	if o == nil || IsNil(o.ReqAmount) {
		var ret float32
		return ret
	}
	return *o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetReqAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ReqAmount) {
		return nil, false
	}
	return o.ReqAmount, true
}

// HasReqAmount returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasReqAmount() bool {
	if o != nil && !IsNil(o.ReqAmount) {
		return true
	}

	return false
}

// SetReqAmount gets a reference to the given float32 and assigns it to the ReqAmount field.
func (o *ResponseTreasuryMiscTransactionResponse) SetReqAmount(v float32) {
	o.ReqAmount = &v
}

// GetReqDetails returns the ReqDetails field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetReqDetails() map[string]float32 {
	if o == nil || IsNil(o.ReqDetails) {
		var ret map[string]float32
		return ret
	}
	return *o.ReqDetails
}

// GetReqDetailsOk returns a tuple with the ReqDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetReqDetailsOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.ReqDetails) {
		return nil, false
	}
	return o.ReqDetails, true
}

// HasReqDetails returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasReqDetails() bool {
	if o != nil && !IsNil(o.ReqDetails) {
		return true
	}

	return false
}

// SetReqDetails gets a reference to the given map[string]float32 and assigns it to the ReqDetails field.
func (o *ResponseTreasuryMiscTransactionResponse) SetReqDetails(v map[string]float32) {
	o.ReqDetails = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetRequestType() string {
	if o == nil || IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasRequestType() bool {
	if o != nil && !IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *ResponseTreasuryMiscTransactionResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseTreasuryMiscTransactionResponse) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseTreasuryMiscTransactionResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTxnMode returns the TxnMode field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetTxnMode() string {
	if o == nil || IsNil(o.TxnMode) {
		var ret string
		return ret
	}
	return *o.TxnMode
}

// GetTxnModeOk returns a tuple with the TxnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetTxnModeOk() (*string, bool) {
	if o == nil || IsNil(o.TxnMode) {
		return nil, false
	}
	return o.TxnMode, true
}

// HasTxnMode returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasTxnMode() bool {
	if o != nil && !IsNil(o.TxnMode) {
		return true
	}

	return false
}

// SetTxnMode gets a reference to the given string and assigns it to the TxnMode field.
func (o *ResponseTreasuryMiscTransactionResponse) SetTxnMode(v string) {
	o.TxnMode = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ResponseTreasuryMiscTransactionResponse) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTreasuryMiscTransactionResponse) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ResponseTreasuryMiscTransactionResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ResponseTreasuryMiscTransactionResponse) SetUserId(v int32) {
	o.UserId = &v
}

func (o ResponseTreasuryMiscTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTreasuryMiscTransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountDescdetails) {
		toSerialize["accountDescdetails"] = o.AccountDescdetails
	}
	if !IsNil(o.ApproverDate) {
		toSerialize["approver_date"] = o.ApproverDate
	}
	if !IsNil(o.ApproverId) {
		toSerialize["approver_id"] = o.ApproverId
	}
	if !IsNil(o.ApproverRemarks) {
		toSerialize["approver_remarks"] = o.ApproverRemarks
	}
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.ReqAmount) {
		toSerialize["req_amount"] = o.ReqAmount
	}
	if !IsNil(o.ReqDetails) {
		toSerialize["req_details"] = o.ReqDetails
	}
	if !IsNil(o.RequestType) {
		toSerialize["request_type"] = o.RequestType
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.TxnMode) {
		toSerialize["txn_mode"] = o.TxnMode
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableResponseTreasuryMiscTransactionResponse struct {
	value *ResponseTreasuryMiscTransactionResponse
	isSet bool
}

func (v NullableResponseTreasuryMiscTransactionResponse) Get() *ResponseTreasuryMiscTransactionResponse {
	return v.value
}

func (v *NullableResponseTreasuryMiscTransactionResponse) Set(val *ResponseTreasuryMiscTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTreasuryMiscTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTreasuryMiscTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTreasuryMiscTransactionResponse(val *ResponseTreasuryMiscTransactionResponse) *NullableResponseTreasuryMiscTransactionResponse {
	return &NullableResponseTreasuryMiscTransactionResponse{value: val, isSet: true}
}

func (v NullableResponseTreasuryMiscTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTreasuryMiscTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



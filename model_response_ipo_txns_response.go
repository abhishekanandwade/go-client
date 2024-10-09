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

// checks if the ResponseIpoTxnsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIpoTxnsResponse{}

// ResponseIpoTxnsResponse struct for ResponseIpoTxnsResponse
type ResponseIpoTxnsResponse struct {
	AckAmount *float32 `json:"ack_amount,omitempty"`
	AckApproverRemarks *string `json:"ack_approver_remarks,omitempty"`
	AckDate *string `json:"ack_date,omitempty"`
	AckUserId *int32 `json:"ack_user_id,omitempty"`
	ApprovedAmount *float32 `json:"approved_amount,omitempty"`
	ApprovedDetails *string `json:"approved_details,omitempty"`
	IpodetailsDesc []ResponseIPOdetails `json:"ipodetailsDesc,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	IsSpecialRemittance *bool `json:"is_special_remittance,omitempty"`
	IssApproverDate *string `json:"iss_approver_date,omitempty"`
	IssApproverId *int32 `json:"iss_approver_id,omitempty"`
	IssApproverRemarks *string `json:"iss_approver_remarks,omitempty"`
	IssOfficeId *int32 `json:"iss_office_id,omitempty"`
	IssOfficeName *string `json:"iss_office_name,omitempty"`
	IssUserId *int32 `json:"iss_user_id,omitempty"`
	IssUserProcessedDate *string `json:"iss_user_processed_date,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	ReqAmount *float32 `json:"req_amount,omitempty"`
	ReqApproverAmt *float32 `json:"req_approver_amt,omitempty"`
	ReqApproverDate *string `json:"req_approver_date,omitempty"`
	ReqApproverId *int32 `json:"req_approver_id,omitempty"`
	ReqApproverRemarks *string `json:"req_approver_remarks,omitempty"`
	ReqDetails *string `json:"req_details,omitempty"`
	ReqOfficeId *int32 `json:"req_office_id,omitempty"`
	ReqOfficeName *string `json:"req_office_name,omitempty"`
	ReqUserId *int32 `json:"req_user_id,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	RequestSource *string `json:"request_source,omitempty"`
	RequestType *string `json:"request_type,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
	TxnStatus *string `json:"txn_status,omitempty"`
}

// NewResponseIpoTxnsResponse instantiates a new ResponseIpoTxnsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIpoTxnsResponse() *ResponseIpoTxnsResponse {
	this := ResponseIpoTxnsResponse{}
	return &this
}

// NewResponseIpoTxnsResponseWithDefaults instantiates a new ResponseIpoTxnsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIpoTxnsResponseWithDefaults() *ResponseIpoTxnsResponse {
	this := ResponseIpoTxnsResponse{}
	return &this
}

// GetAckAmount returns the AckAmount field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetAckAmount() float32 {
	if o == nil || IsNil(o.AckAmount) {
		var ret float32
		return ret
	}
	return *o.AckAmount
}

// GetAckAmountOk returns a tuple with the AckAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetAckAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.AckAmount) {
		return nil, false
	}
	return o.AckAmount, true
}

// HasAckAmount returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasAckAmount() bool {
	if o != nil && !IsNil(o.AckAmount) {
		return true
	}

	return false
}

// SetAckAmount gets a reference to the given float32 and assigns it to the AckAmount field.
func (o *ResponseIpoTxnsResponse) SetAckAmount(v float32) {
	o.AckAmount = &v
}

// GetAckApproverRemarks returns the AckApproverRemarks field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetAckApproverRemarks() string {
	if o == nil || IsNil(o.AckApproverRemarks) {
		var ret string
		return ret
	}
	return *o.AckApproverRemarks
}

// GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetAckApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.AckApproverRemarks) {
		return nil, false
	}
	return o.AckApproverRemarks, true
}

// HasAckApproverRemarks returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasAckApproverRemarks() bool {
	if o != nil && !IsNil(o.AckApproverRemarks) {
		return true
	}

	return false
}

// SetAckApproverRemarks gets a reference to the given string and assigns it to the AckApproverRemarks field.
func (o *ResponseIpoTxnsResponse) SetAckApproverRemarks(v string) {
	o.AckApproverRemarks = &v
}

// GetAckDate returns the AckDate field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetAckDate() string {
	if o == nil || IsNil(o.AckDate) {
		var ret string
		return ret
	}
	return *o.AckDate
}

// GetAckDateOk returns a tuple with the AckDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetAckDateOk() (*string, bool) {
	if o == nil || IsNil(o.AckDate) {
		return nil, false
	}
	return o.AckDate, true
}

// HasAckDate returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasAckDate() bool {
	if o != nil && !IsNil(o.AckDate) {
		return true
	}

	return false
}

// SetAckDate gets a reference to the given string and assigns it to the AckDate field.
func (o *ResponseIpoTxnsResponse) SetAckDate(v string) {
	o.AckDate = &v
}

// GetAckUserId returns the AckUserId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetAckUserId() int32 {
	if o == nil || IsNil(o.AckUserId) {
		var ret int32
		return ret
	}
	return *o.AckUserId
}

// GetAckUserIdOk returns a tuple with the AckUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetAckUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AckUserId) {
		return nil, false
	}
	return o.AckUserId, true
}

// HasAckUserId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasAckUserId() bool {
	if o != nil && !IsNil(o.AckUserId) {
		return true
	}

	return false
}

// SetAckUserId gets a reference to the given int32 and assigns it to the AckUserId field.
func (o *ResponseIpoTxnsResponse) SetAckUserId(v int32) {
	o.AckUserId = &v
}

// GetApprovedAmount returns the ApprovedAmount field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetApprovedAmount() float32 {
	if o == nil || IsNil(o.ApprovedAmount) {
		var ret float32
		return ret
	}
	return *o.ApprovedAmount
}

// GetApprovedAmountOk returns a tuple with the ApprovedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetApprovedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ApprovedAmount) {
		return nil, false
	}
	return o.ApprovedAmount, true
}

// HasApprovedAmount returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasApprovedAmount() bool {
	if o != nil && !IsNil(o.ApprovedAmount) {
		return true
	}

	return false
}

// SetApprovedAmount gets a reference to the given float32 and assigns it to the ApprovedAmount field.
func (o *ResponseIpoTxnsResponse) SetApprovedAmount(v float32) {
	o.ApprovedAmount = &v
}

// GetApprovedDetails returns the ApprovedDetails field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetApprovedDetails() string {
	if o == nil || IsNil(o.ApprovedDetails) {
		var ret string
		return ret
	}
	return *o.ApprovedDetails
}

// GetApprovedDetailsOk returns a tuple with the ApprovedDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetApprovedDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovedDetails) {
		return nil, false
	}
	return o.ApprovedDetails, true
}

// HasApprovedDetails returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasApprovedDetails() bool {
	if o != nil && !IsNil(o.ApprovedDetails) {
		return true
	}

	return false
}

// SetApprovedDetails gets a reference to the given string and assigns it to the ApprovedDetails field.
func (o *ResponseIpoTxnsResponse) SetApprovedDetails(v string) {
	o.ApprovedDetails = &v
}

// GetIpodetailsDesc returns the IpodetailsDesc field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIpodetailsDesc() []ResponseIPOdetails {
	if o == nil || IsNil(o.IpodetailsDesc) {
		var ret []ResponseIPOdetails
		return ret
	}
	return o.IpodetailsDesc
}

// GetIpodetailsDescOk returns a tuple with the IpodetailsDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIpodetailsDescOk() ([]ResponseIPOdetails, bool) {
	if o == nil || IsNil(o.IpodetailsDesc) {
		return nil, false
	}
	return o.IpodetailsDesc, true
}

// HasIpodetailsDesc returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIpodetailsDesc() bool {
	if o != nil && !IsNil(o.IpodetailsDesc) {
		return true
	}

	return false
}

// SetIpodetailsDesc gets a reference to the given []ResponseIPOdetails and assigns it to the IpodetailsDesc field.
func (o *ResponseIpoTxnsResponse) SetIpodetailsDesc(v []ResponseIPOdetails) {
	o.IpodetailsDesc = v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *ResponseIpoTxnsResponse) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetIsSpecialRemittance returns the IsSpecialRemittance field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIsSpecialRemittance() bool {
	if o == nil || IsNil(o.IsSpecialRemittance) {
		var ret bool
		return ret
	}
	return *o.IsSpecialRemittance
}

// GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIsSpecialRemittanceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSpecialRemittance) {
		return nil, false
	}
	return o.IsSpecialRemittance, true
}

// HasIsSpecialRemittance returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIsSpecialRemittance() bool {
	if o != nil && !IsNil(o.IsSpecialRemittance) {
		return true
	}

	return false
}

// SetIsSpecialRemittance gets a reference to the given bool and assigns it to the IsSpecialRemittance field.
func (o *ResponseIpoTxnsResponse) SetIsSpecialRemittance(v bool) {
	o.IsSpecialRemittance = &v
}

// GetIssApproverDate returns the IssApproverDate field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssApproverDate() string {
	if o == nil || IsNil(o.IssApproverDate) {
		var ret string
		return ret
	}
	return *o.IssApproverDate
}

// GetIssApproverDateOk returns a tuple with the IssApproverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssApproverDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssApproverDate) {
		return nil, false
	}
	return o.IssApproverDate, true
}

// HasIssApproverDate returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssApproverDate() bool {
	if o != nil && !IsNil(o.IssApproverDate) {
		return true
	}

	return false
}

// SetIssApproverDate gets a reference to the given string and assigns it to the IssApproverDate field.
func (o *ResponseIpoTxnsResponse) SetIssApproverDate(v string) {
	o.IssApproverDate = &v
}

// GetIssApproverId returns the IssApproverId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssApproverId() int32 {
	if o == nil || IsNil(o.IssApproverId) {
		var ret int32
		return ret
	}
	return *o.IssApproverId
}

// GetIssApproverIdOk returns a tuple with the IssApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssApproverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IssApproverId) {
		return nil, false
	}
	return o.IssApproverId, true
}

// HasIssApproverId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssApproverId() bool {
	if o != nil && !IsNil(o.IssApproverId) {
		return true
	}

	return false
}

// SetIssApproverId gets a reference to the given int32 and assigns it to the IssApproverId field.
func (o *ResponseIpoTxnsResponse) SetIssApproverId(v int32) {
	o.IssApproverId = &v
}

// GetIssApproverRemarks returns the IssApproverRemarks field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssApproverRemarks() string {
	if o == nil || IsNil(o.IssApproverRemarks) {
		var ret string
		return ret
	}
	return *o.IssApproverRemarks
}

// GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.IssApproverRemarks) {
		return nil, false
	}
	return o.IssApproverRemarks, true
}

// HasIssApproverRemarks returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssApproverRemarks() bool {
	if o != nil && !IsNil(o.IssApproverRemarks) {
		return true
	}

	return false
}

// SetIssApproverRemarks gets a reference to the given string and assigns it to the IssApproverRemarks field.
func (o *ResponseIpoTxnsResponse) SetIssApproverRemarks(v string) {
	o.IssApproverRemarks = &v
}

// GetIssOfficeId returns the IssOfficeId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssOfficeId() int32 {
	if o == nil || IsNil(o.IssOfficeId) {
		var ret int32
		return ret
	}
	return *o.IssOfficeId
}

// GetIssOfficeIdOk returns a tuple with the IssOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IssOfficeId) {
		return nil, false
	}
	return o.IssOfficeId, true
}

// HasIssOfficeId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssOfficeId() bool {
	if o != nil && !IsNil(o.IssOfficeId) {
		return true
	}

	return false
}

// SetIssOfficeId gets a reference to the given int32 and assigns it to the IssOfficeId field.
func (o *ResponseIpoTxnsResponse) SetIssOfficeId(v int32) {
	o.IssOfficeId = &v
}

// GetIssOfficeName returns the IssOfficeName field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssOfficeName() string {
	if o == nil || IsNil(o.IssOfficeName) {
		var ret string
		return ret
	}
	return *o.IssOfficeName
}

// GetIssOfficeNameOk returns a tuple with the IssOfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.IssOfficeName) {
		return nil, false
	}
	return o.IssOfficeName, true
}

// HasIssOfficeName returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssOfficeName() bool {
	if o != nil && !IsNil(o.IssOfficeName) {
		return true
	}

	return false
}

// SetIssOfficeName gets a reference to the given string and assigns it to the IssOfficeName field.
func (o *ResponseIpoTxnsResponse) SetIssOfficeName(v string) {
	o.IssOfficeName = &v
}

// GetIssUserId returns the IssUserId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssUserId() int32 {
	if o == nil || IsNil(o.IssUserId) {
		var ret int32
		return ret
	}
	return *o.IssUserId
}

// GetIssUserIdOk returns a tuple with the IssUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.IssUserId) {
		return nil, false
	}
	return o.IssUserId, true
}

// HasIssUserId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssUserId() bool {
	if o != nil && !IsNil(o.IssUserId) {
		return true
	}

	return false
}

// SetIssUserId gets a reference to the given int32 and assigns it to the IssUserId field.
func (o *ResponseIpoTxnsResponse) SetIssUserId(v int32) {
	o.IssUserId = &v
}

// GetIssUserProcessedDate returns the IssUserProcessedDate field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetIssUserProcessedDate() string {
	if o == nil || IsNil(o.IssUserProcessedDate) {
		var ret string
		return ret
	}
	return *o.IssUserProcessedDate
}

// GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetIssUserProcessedDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssUserProcessedDate) {
		return nil, false
	}
	return o.IssUserProcessedDate, true
}

// HasIssUserProcessedDate returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasIssUserProcessedDate() bool {
	if o != nil && !IsNil(o.IssUserProcessedDate) {
		return true
	}

	return false
}

// SetIssUserProcessedDate gets a reference to the given string and assigns it to the IssUserProcessedDate field.
func (o *ResponseIpoTxnsResponse) SetIssUserProcessedDate(v string) {
	o.IssUserProcessedDate = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseIpoTxnsResponse) SetRemarks(v string) {
	o.Remarks = &v
}

// GetReqAmount returns the ReqAmount field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqAmount() float32 {
	if o == nil || IsNil(o.ReqAmount) {
		var ret float32
		return ret
	}
	return *o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ReqAmount) {
		return nil, false
	}
	return o.ReqAmount, true
}

// HasReqAmount returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqAmount() bool {
	if o != nil && !IsNil(o.ReqAmount) {
		return true
	}

	return false
}

// SetReqAmount gets a reference to the given float32 and assigns it to the ReqAmount field.
func (o *ResponseIpoTxnsResponse) SetReqAmount(v float32) {
	o.ReqAmount = &v
}

// GetReqApproverAmt returns the ReqApproverAmt field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqApproverAmt() float32 {
	if o == nil || IsNil(o.ReqApproverAmt) {
		var ret float32
		return ret
	}
	return *o.ReqApproverAmt
}

// GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqApproverAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.ReqApproverAmt) {
		return nil, false
	}
	return o.ReqApproverAmt, true
}

// HasReqApproverAmt returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqApproverAmt() bool {
	if o != nil && !IsNil(o.ReqApproverAmt) {
		return true
	}

	return false
}

// SetReqApproverAmt gets a reference to the given float32 and assigns it to the ReqApproverAmt field.
func (o *ResponseIpoTxnsResponse) SetReqApproverAmt(v float32) {
	o.ReqApproverAmt = &v
}

// GetReqApproverDate returns the ReqApproverDate field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqApproverDate() string {
	if o == nil || IsNil(o.ReqApproverDate) {
		var ret string
		return ret
	}
	return *o.ReqApproverDate
}

// GetReqApproverDateOk returns a tuple with the ReqApproverDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqApproverDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReqApproverDate) {
		return nil, false
	}
	return o.ReqApproverDate, true
}

// HasReqApproverDate returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqApproverDate() bool {
	if o != nil && !IsNil(o.ReqApproverDate) {
		return true
	}

	return false
}

// SetReqApproverDate gets a reference to the given string and assigns it to the ReqApproverDate field.
func (o *ResponseIpoTxnsResponse) SetReqApproverDate(v string) {
	o.ReqApproverDate = &v
}

// GetReqApproverId returns the ReqApproverId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqApproverId() int32 {
	if o == nil || IsNil(o.ReqApproverId) {
		var ret int32
		return ret
	}
	return *o.ReqApproverId
}

// GetReqApproverIdOk returns a tuple with the ReqApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqApproverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReqApproverId) {
		return nil, false
	}
	return o.ReqApproverId, true
}

// HasReqApproverId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqApproverId() bool {
	if o != nil && !IsNil(o.ReqApproverId) {
		return true
	}

	return false
}

// SetReqApproverId gets a reference to the given int32 and assigns it to the ReqApproverId field.
func (o *ResponseIpoTxnsResponse) SetReqApproverId(v int32) {
	o.ReqApproverId = &v
}

// GetReqApproverRemarks returns the ReqApproverRemarks field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqApproverRemarks() string {
	if o == nil || IsNil(o.ReqApproverRemarks) {
		var ret string
		return ret
	}
	return *o.ReqApproverRemarks
}

// GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.ReqApproverRemarks) {
		return nil, false
	}
	return o.ReqApproverRemarks, true
}

// HasReqApproverRemarks returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqApproverRemarks() bool {
	if o != nil && !IsNil(o.ReqApproverRemarks) {
		return true
	}

	return false
}

// SetReqApproverRemarks gets a reference to the given string and assigns it to the ReqApproverRemarks field.
func (o *ResponseIpoTxnsResponse) SetReqApproverRemarks(v string) {
	o.ReqApproverRemarks = &v
}

// GetReqDetails returns the ReqDetails field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqDetails() string {
	if o == nil || IsNil(o.ReqDetails) {
		var ret string
		return ret
	}
	return *o.ReqDetails
}

// GetReqDetailsOk returns a tuple with the ReqDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.ReqDetails) {
		return nil, false
	}
	return o.ReqDetails, true
}

// HasReqDetails returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqDetails() bool {
	if o != nil && !IsNil(o.ReqDetails) {
		return true
	}

	return false
}

// SetReqDetails gets a reference to the given string and assigns it to the ReqDetails field.
func (o *ResponseIpoTxnsResponse) SetReqDetails(v string) {
	o.ReqDetails = &v
}

// GetReqOfficeId returns the ReqOfficeId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqOfficeId() int32 {
	if o == nil || IsNil(o.ReqOfficeId) {
		var ret int32
		return ret
	}
	return *o.ReqOfficeId
}

// GetReqOfficeIdOk returns a tuple with the ReqOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReqOfficeId) {
		return nil, false
	}
	return o.ReqOfficeId, true
}

// HasReqOfficeId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqOfficeId() bool {
	if o != nil && !IsNil(o.ReqOfficeId) {
		return true
	}

	return false
}

// SetReqOfficeId gets a reference to the given int32 and assigns it to the ReqOfficeId field.
func (o *ResponseIpoTxnsResponse) SetReqOfficeId(v int32) {
	o.ReqOfficeId = &v
}

// GetReqOfficeName returns the ReqOfficeName field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqOfficeName() string {
	if o == nil || IsNil(o.ReqOfficeName) {
		var ret string
		return ret
	}
	return *o.ReqOfficeName
}

// GetReqOfficeNameOk returns a tuple with the ReqOfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReqOfficeName) {
		return nil, false
	}
	return o.ReqOfficeName, true
}

// HasReqOfficeName returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqOfficeName() bool {
	if o != nil && !IsNil(o.ReqOfficeName) {
		return true
	}

	return false
}

// SetReqOfficeName gets a reference to the given string and assigns it to the ReqOfficeName field.
func (o *ResponseIpoTxnsResponse) SetReqOfficeName(v string) {
	o.ReqOfficeName = &v
}

// GetReqUserId returns the ReqUserId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetReqUserId() int32 {
	if o == nil || IsNil(o.ReqUserId) {
		var ret int32
		return ret
	}
	return *o.ReqUserId
}

// GetReqUserIdOk returns a tuple with the ReqUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetReqUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ReqUserId) {
		return nil, false
	}
	return o.ReqUserId, true
}

// HasReqUserId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasReqUserId() bool {
	if o != nil && !IsNil(o.ReqUserId) {
		return true
	}

	return false
}

// SetReqUserId gets a reference to the given int32 and assigns it to the ReqUserId field.
func (o *ResponseIpoTxnsResponse) SetReqUserId(v int32) {
	o.ReqUserId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ResponseIpoTxnsResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestSource returns the RequestSource field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetRequestSource() string {
	if o == nil || IsNil(o.RequestSource) {
		var ret string
		return ret
	}
	return *o.RequestSource
}

// GetRequestSourceOk returns a tuple with the RequestSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetRequestSourceOk() (*string, bool) {
	if o == nil || IsNil(o.RequestSource) {
		return nil, false
	}
	return o.RequestSource, true
}

// HasRequestSource returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasRequestSource() bool {
	if o != nil && !IsNil(o.RequestSource) {
		return true
	}

	return false
}

// SetRequestSource gets a reference to the given string and assigns it to the RequestSource field.
func (o *ResponseIpoTxnsResponse) SetRequestSource(v string) {
	o.RequestSource = &v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetRequestType() string {
	if o == nil || IsNil(o.RequestType) {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetRequestTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestType) {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasRequestType() bool {
	if o != nil && !IsNil(o.RequestType) {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *ResponseIpoTxnsResponse) SetRequestType(v string) {
	o.RequestType = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseIpoTxnsResponse) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseIpoTxnsResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTxnStatus returns the TxnStatus field value if set, zero value otherwise.
func (o *ResponseIpoTxnsResponse) GetTxnStatus() string {
	if o == nil || IsNil(o.TxnStatus) {
		var ret string
		return ret
	}
	return *o.TxnStatus
}

// GetTxnStatusOk returns a tuple with the TxnStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoTxnsResponse) GetTxnStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TxnStatus) {
		return nil, false
	}
	return o.TxnStatus, true
}

// HasTxnStatus returns a boolean if a field has been set.
func (o *ResponseIpoTxnsResponse) HasTxnStatus() bool {
	if o != nil && !IsNil(o.TxnStatus) {
		return true
	}

	return false
}

// SetTxnStatus gets a reference to the given string and assigns it to the TxnStatus field.
func (o *ResponseIpoTxnsResponse) SetTxnStatus(v string) {
	o.TxnStatus = &v
}

func (o ResponseIpoTxnsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIpoTxnsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AckAmount) {
		toSerialize["ack_amount"] = o.AckAmount
	}
	if !IsNil(o.AckApproverRemarks) {
		toSerialize["ack_approver_remarks"] = o.AckApproverRemarks
	}
	if !IsNil(o.AckDate) {
		toSerialize["ack_date"] = o.AckDate
	}
	if !IsNil(o.AckUserId) {
		toSerialize["ack_user_id"] = o.AckUserId
	}
	if !IsNil(o.ApprovedAmount) {
		toSerialize["approved_amount"] = o.ApprovedAmount
	}
	if !IsNil(o.ApprovedDetails) {
		toSerialize["approved_details"] = o.ApprovedDetails
	}
	if !IsNil(o.IpodetailsDesc) {
		toSerialize["ipodetailsDesc"] = o.IpodetailsDesc
	}
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	if !IsNil(o.IsSpecialRemittance) {
		toSerialize["is_special_remittance"] = o.IsSpecialRemittance
	}
	if !IsNil(o.IssApproverDate) {
		toSerialize["iss_approver_date"] = o.IssApproverDate
	}
	if !IsNil(o.IssApproverId) {
		toSerialize["iss_approver_id"] = o.IssApproverId
	}
	if !IsNil(o.IssApproverRemarks) {
		toSerialize["iss_approver_remarks"] = o.IssApproverRemarks
	}
	if !IsNil(o.IssOfficeId) {
		toSerialize["iss_office_id"] = o.IssOfficeId
	}
	if !IsNil(o.IssOfficeName) {
		toSerialize["iss_office_name"] = o.IssOfficeName
	}
	if !IsNil(o.IssUserId) {
		toSerialize["iss_user_id"] = o.IssUserId
	}
	if !IsNil(o.IssUserProcessedDate) {
		toSerialize["iss_user_processed_date"] = o.IssUserProcessedDate
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.ReqAmount) {
		toSerialize["req_amount"] = o.ReqAmount
	}
	if !IsNil(o.ReqApproverAmt) {
		toSerialize["req_approver_amt"] = o.ReqApproverAmt
	}
	if !IsNil(o.ReqApproverDate) {
		toSerialize["req_approver_date"] = o.ReqApproverDate
	}
	if !IsNil(o.ReqApproverId) {
		toSerialize["req_approver_id"] = o.ReqApproverId
	}
	if !IsNil(o.ReqApproverRemarks) {
		toSerialize["req_approver_remarks"] = o.ReqApproverRemarks
	}
	if !IsNil(o.ReqDetails) {
		toSerialize["req_details"] = o.ReqDetails
	}
	if !IsNil(o.ReqOfficeId) {
		toSerialize["req_office_id"] = o.ReqOfficeId
	}
	if !IsNil(o.ReqOfficeName) {
		toSerialize["req_office_name"] = o.ReqOfficeName
	}
	if !IsNil(o.ReqUserId) {
		toSerialize["req_user_id"] = o.ReqUserId
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.RequestSource) {
		toSerialize["request_source"] = o.RequestSource
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
	if !IsNil(o.TxnStatus) {
		toSerialize["txn_status"] = o.TxnStatus
	}
	return toSerialize, nil
}

type NullableResponseIpoTxnsResponse struct {
	value *ResponseIpoTxnsResponse
	isSet bool
}

func (v NullableResponseIpoTxnsResponse) Get() *ResponseIpoTxnsResponse {
	return v.value
}

func (v *NullableResponseIpoTxnsResponse) Set(val *ResponseIpoTxnsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIpoTxnsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIpoTxnsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIpoTxnsResponse(val *ResponseIpoTxnsResponse) *NullableResponseIpoTxnsResponse {
	return &NullableResponseIpoTxnsResponse{value: val, isSet: true}
}

func (v NullableResponseIpoTxnsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIpoTxnsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



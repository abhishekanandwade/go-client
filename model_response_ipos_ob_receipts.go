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

// checks if the ResponseIposOBReceipts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIposOBReceipts{}

// ResponseIposOBReceipts struct for ResponseIposOBReceipts
type ResponseIposOBReceipts struct {
	ApprUserId *int32 `json:"appr_user_id,omitempty"`
	ApprovedDate *string `json:"approved_date,omitempty"`
	ApproverRemarks *string `json:"approver_remarks,omitempty"`
	InvoiceDate *string `json:"invoice_date,omitempty"`
	IpoDetails []ResponseIPOdetails `json:"ipo_details,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	ObTransactionId *string `json:"ob_transaction_id,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	RecdUserId *int32 `json:"recd_user_id,omitempty"`
	ReceiptInvoice *string `json:"receipt_invoice,omitempty"`
	ReceiptSource *string `json:"receipt_source,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
}

// NewResponseIposOBReceipts instantiates a new ResponseIposOBReceipts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIposOBReceipts() *ResponseIposOBReceipts {
	this := ResponseIposOBReceipts{}
	return &this
}

// NewResponseIposOBReceiptsWithDefaults instantiates a new ResponseIposOBReceipts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIposOBReceiptsWithDefaults() *ResponseIposOBReceipts {
	this := ResponseIposOBReceipts{}
	return &this
}

// GetApprUserId returns the ApprUserId field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetApprUserId() int32 {
	if o == nil || IsNil(o.ApprUserId) {
		var ret int32
		return ret
	}
	return *o.ApprUserId
}

// GetApprUserIdOk returns a tuple with the ApprUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetApprUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApprUserId) {
		return nil, false
	}
	return o.ApprUserId, true
}

// HasApprUserId returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasApprUserId() bool {
	if o != nil && !IsNil(o.ApprUserId) {
		return true
	}

	return false
}

// SetApprUserId gets a reference to the given int32 and assigns it to the ApprUserId field.
func (o *ResponseIposOBReceipts) SetApprUserId(v int32) {
	o.ApprUserId = &v
}

// GetApprovedDate returns the ApprovedDate field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetApprovedDate() string {
	if o == nil || IsNil(o.ApprovedDate) {
		var ret string
		return ret
	}
	return *o.ApprovedDate
}

// GetApprovedDateOk returns a tuple with the ApprovedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetApprovedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovedDate) {
		return nil, false
	}
	return o.ApprovedDate, true
}

// HasApprovedDate returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasApprovedDate() bool {
	if o != nil && !IsNil(o.ApprovedDate) {
		return true
	}

	return false
}

// SetApprovedDate gets a reference to the given string and assigns it to the ApprovedDate field.
func (o *ResponseIposOBReceipts) SetApprovedDate(v string) {
	o.ApprovedDate = &v
}

// GetApproverRemarks returns the ApproverRemarks field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetApproverRemarks() string {
	if o == nil || IsNil(o.ApproverRemarks) {
		var ret string
		return ret
	}
	return *o.ApproverRemarks
}

// GetApproverRemarksOk returns a tuple with the ApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.ApproverRemarks) {
		return nil, false
	}
	return o.ApproverRemarks, true
}

// HasApproverRemarks returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasApproverRemarks() bool {
	if o != nil && !IsNil(o.ApproverRemarks) {
		return true
	}

	return false
}

// SetApproverRemarks gets a reference to the given string and assigns it to the ApproverRemarks field.
func (o *ResponseIposOBReceipts) SetApproverRemarks(v string) {
	o.ApproverRemarks = &v
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *ResponseIposOBReceipts) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetIpoDetails returns the IpoDetails field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetIpoDetails() []ResponseIPOdetails {
	if o == nil || IsNil(o.IpoDetails) {
		var ret []ResponseIPOdetails
		return ret
	}
	return o.IpoDetails
}

// GetIpoDetailsOk returns a tuple with the IpoDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetIpoDetailsOk() ([]ResponseIPOdetails, bool) {
	if o == nil || IsNil(o.IpoDetails) {
		return nil, false
	}
	return o.IpoDetails, true
}

// HasIpoDetails returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasIpoDetails() bool {
	if o != nil && !IsNil(o.IpoDetails) {
		return true
	}

	return false
}

// SetIpoDetails gets a reference to the given []ResponseIPOdetails and assigns it to the IpoDetails field.
func (o *ResponseIposOBReceipts) SetIpoDetails(v []ResponseIPOdetails) {
	o.IpoDetails = v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *ResponseIposOBReceipts) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetObTransactionId returns the ObTransactionId field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetObTransactionId() string {
	if o == nil || IsNil(o.ObTransactionId) {
		var ret string
		return ret
	}
	return *o.ObTransactionId
}

// GetObTransactionIdOk returns a tuple with the ObTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetObTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObTransactionId) {
		return nil, false
	}
	return o.ObTransactionId, true
}

// HasObTransactionId returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasObTransactionId() bool {
	if o != nil && !IsNil(o.ObTransactionId) {
		return true
	}

	return false
}

// SetObTransactionId gets a reference to the given string and assigns it to the ObTransactionId field.
func (o *ResponseIposOBReceipts) SetObTransactionId(v string) {
	o.ObTransactionId = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseIposOBReceipts) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetRecdUserId returns the RecdUserId field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetRecdUserId() int32 {
	if o == nil || IsNil(o.RecdUserId) {
		var ret int32
		return ret
	}
	return *o.RecdUserId
}

// GetRecdUserIdOk returns a tuple with the RecdUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetRecdUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RecdUserId) {
		return nil, false
	}
	return o.RecdUserId, true
}

// HasRecdUserId returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasRecdUserId() bool {
	if o != nil && !IsNil(o.RecdUserId) {
		return true
	}

	return false
}

// SetRecdUserId gets a reference to the given int32 and assigns it to the RecdUserId field.
func (o *ResponseIposOBReceipts) SetRecdUserId(v int32) {
	o.RecdUserId = &v
}

// GetReceiptInvoice returns the ReceiptInvoice field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetReceiptInvoice() string {
	if o == nil || IsNil(o.ReceiptInvoice) {
		var ret string
		return ret
	}
	return *o.ReceiptInvoice
}

// GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetReceiptInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptInvoice) {
		return nil, false
	}
	return o.ReceiptInvoice, true
}

// HasReceiptInvoice returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasReceiptInvoice() bool {
	if o != nil && !IsNil(o.ReceiptInvoice) {
		return true
	}

	return false
}

// SetReceiptInvoice gets a reference to the given string and assigns it to the ReceiptInvoice field.
func (o *ResponseIposOBReceipts) SetReceiptInvoice(v string) {
	o.ReceiptInvoice = &v
}

// GetReceiptSource returns the ReceiptSource field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetReceiptSource() string {
	if o == nil || IsNil(o.ReceiptSource) {
		var ret string
		return ret
	}
	return *o.ReceiptSource
}

// GetReceiptSourceOk returns a tuple with the ReceiptSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetReceiptSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptSource) {
		return nil, false
	}
	return o.ReceiptSource, true
}

// HasReceiptSource returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasReceiptSource() bool {
	if o != nil && !IsNil(o.ReceiptSource) {
		return true
	}

	return false
}

// SetReceiptSource gets a reference to the given string and assigns it to the ReceiptSource field.
func (o *ResponseIposOBReceipts) SetReceiptSource(v string) {
	o.ReceiptSource = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseIposOBReceipts) SetRemarks(v string) {
	o.Remarks = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseIposOBReceipts) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIposOBReceipts) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseIposOBReceipts) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseIposOBReceipts) SetTransDate(v string) {
	o.TransDate = &v
}

func (o ResponseIposOBReceipts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIposOBReceipts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprUserId) {
		toSerialize["appr_user_id"] = o.ApprUserId
	}
	if !IsNil(o.ApprovedDate) {
		toSerialize["approved_date"] = o.ApprovedDate
	}
	if !IsNil(o.ApproverRemarks) {
		toSerialize["approver_remarks"] = o.ApproverRemarks
	}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.IpoDetails) {
		toSerialize["ipo_details"] = o.IpoDetails
	}
	if !IsNil(o.IsApproved) {
		toSerialize["is_approved"] = o.IsApproved
	}
	if !IsNil(o.ObTransactionId) {
		toSerialize["ob_transaction_id"] = o.ObTransactionId
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.RecdUserId) {
		toSerialize["recd_user_id"] = o.RecdUserId
	}
	if !IsNil(o.ReceiptInvoice) {
		toSerialize["receipt_invoice"] = o.ReceiptInvoice
	}
	if !IsNil(o.ReceiptSource) {
		toSerialize["receipt_source"] = o.ReceiptSource
	}
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	return toSerialize, nil
}

type NullableResponseIposOBReceipts struct {
	value *ResponseIposOBReceipts
	isSet bool
}

func (v NullableResponseIposOBReceipts) Get() *ResponseIposOBReceipts {
	return v.value
}

func (v *NullableResponseIposOBReceipts) Set(val *ResponseIposOBReceipts) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIposOBReceipts) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIposOBReceipts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIposOBReceipts(val *ResponseIposOBReceipts) *NullableResponseIposOBReceipts {
	return &NullableResponseIposOBReceipts{value: val, isSet: true}
}

func (v NullableResponseIposOBReceipts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIposOBReceipts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



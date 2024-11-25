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

// checks if the ResponseIpoSale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseIpoSale{}

// ResponseIpoSale struct for ResponseIpoSale
type ResponseIpoSale struct {
	ApprovedDate *string `json:"approved_date,omitempty"`
	ApproverId *int32 `json:"approver_id,omitempty"`
	ApproverName *string `json:"approver_name,omitempty"`
	ApproverRemarks *string `json:"approver_remarks,omitempty"`
	IsApproved *bool `json:"is_approved,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	OfficeName *string `json:"office_name,omitempty"`
	SaleAmount *float32 `json:"sale_amount,omitempty"`
	SaleCommission *float32 `json:"sale_commission,omitempty"`
	SaleDetails []int32 `json:"sale_details,omitempty"`
	SoldTo *string `json:"sold_to,omitempty"`
	TransDate *string `json:"trans_date,omitempty"`
	TransactionId *string `json:"transaction_id,omitempty"`
	Transdate *string `json:"transdate,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty"`
}

// NewResponseIpoSale instantiates a new ResponseIpoSale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseIpoSale() *ResponseIpoSale {
	this := ResponseIpoSale{}
	return &this
}

// NewResponseIpoSaleWithDefaults instantiates a new ResponseIpoSale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseIpoSaleWithDefaults() *ResponseIpoSale {
	this := ResponseIpoSale{}
	return &this
}

// GetApprovedDate returns the ApprovedDate field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetApprovedDate() string {
	if o == nil || IsNil(o.ApprovedDate) {
		var ret string
		return ret
	}
	return *o.ApprovedDate
}

// GetApprovedDateOk returns a tuple with the ApprovedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetApprovedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ApprovedDate) {
		return nil, false
	}
	return o.ApprovedDate, true
}

// HasApprovedDate returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasApprovedDate() bool {
	if o != nil && !IsNil(o.ApprovedDate) {
		return true
	}

	return false
}

// SetApprovedDate gets a reference to the given string and assigns it to the ApprovedDate field.
func (o *ResponseIpoSale) SetApprovedDate(v string) {
	o.ApprovedDate = &v
}

// GetApproverId returns the ApproverId field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetApproverId() int32 {
	if o == nil || IsNil(o.ApproverId) {
		var ret int32
		return ret
	}
	return *o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetApproverIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApproverId) {
		return nil, false
	}
	return o.ApproverId, true
}

// HasApproverId returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasApproverId() bool {
	if o != nil && !IsNil(o.ApproverId) {
		return true
	}

	return false
}

// SetApproverId gets a reference to the given int32 and assigns it to the ApproverId field.
func (o *ResponseIpoSale) SetApproverId(v int32) {
	o.ApproverId = &v
}

// GetApproverName returns the ApproverName field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetApproverName() string {
	if o == nil || IsNil(o.ApproverName) {
		var ret string
		return ret
	}
	return *o.ApproverName
}

// GetApproverNameOk returns a tuple with the ApproverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetApproverNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApproverName) {
		return nil, false
	}
	return o.ApproverName, true
}

// HasApproverName returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasApproverName() bool {
	if o != nil && !IsNil(o.ApproverName) {
		return true
	}

	return false
}

// SetApproverName gets a reference to the given string and assigns it to the ApproverName field.
func (o *ResponseIpoSale) SetApproverName(v string) {
	o.ApproverName = &v
}

// GetApproverRemarks returns the ApproverRemarks field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetApproverRemarks() string {
	if o == nil || IsNil(o.ApproverRemarks) {
		var ret string
		return ret
	}
	return *o.ApproverRemarks
}

// GetApproverRemarksOk returns a tuple with the ApproverRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetApproverRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.ApproverRemarks) {
		return nil, false
	}
	return o.ApproverRemarks, true
}

// HasApproverRemarks returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasApproverRemarks() bool {
	if o != nil && !IsNil(o.ApproverRemarks) {
		return true
	}

	return false
}

// SetApproverRemarks gets a reference to the given string and assigns it to the ApproverRemarks field.
func (o *ResponseIpoSale) SetApproverRemarks(v string) {
	o.ApproverRemarks = &v
}

// GetIsApproved returns the IsApproved field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetIsApproved() bool {
	if o == nil || IsNil(o.IsApproved) {
		var ret bool
		return ret
	}
	return *o.IsApproved
}

// GetIsApprovedOk returns a tuple with the IsApproved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetIsApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsApproved) {
		return nil, false
	}
	return o.IsApproved, true
}

// HasIsApproved returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasIsApproved() bool {
	if o != nil && !IsNil(o.IsApproved) {
		return true
	}

	return false
}

// SetIsApproved gets a reference to the given bool and assigns it to the IsApproved field.
func (o *ResponseIpoSale) SetIsApproved(v bool) {
	o.IsApproved = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseIpoSale) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetOfficeName returns the OfficeName field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetOfficeName() string {
	if o == nil || IsNil(o.OfficeName) {
		var ret string
		return ret
	}
	return *o.OfficeName
}

// GetOfficeNameOk returns a tuple with the OfficeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetOfficeNameOk() (*string, bool) {
	if o == nil || IsNil(o.OfficeName) {
		return nil, false
	}
	return o.OfficeName, true
}

// HasOfficeName returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasOfficeName() bool {
	if o != nil && !IsNil(o.OfficeName) {
		return true
	}

	return false
}

// SetOfficeName gets a reference to the given string and assigns it to the OfficeName field.
func (o *ResponseIpoSale) SetOfficeName(v string) {
	o.OfficeName = &v
}

// GetSaleAmount returns the SaleAmount field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetSaleAmount() float32 {
	if o == nil || IsNil(o.SaleAmount) {
		var ret float32
		return ret
	}
	return *o.SaleAmount
}

// GetSaleAmountOk returns a tuple with the SaleAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetSaleAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.SaleAmount) {
		return nil, false
	}
	return o.SaleAmount, true
}

// HasSaleAmount returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasSaleAmount() bool {
	if o != nil && !IsNil(o.SaleAmount) {
		return true
	}

	return false
}

// SetSaleAmount gets a reference to the given float32 and assigns it to the SaleAmount field.
func (o *ResponseIpoSale) SetSaleAmount(v float32) {
	o.SaleAmount = &v
}

// GetSaleCommission returns the SaleCommission field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetSaleCommission() float32 {
	if o == nil || IsNil(o.SaleCommission) {
		var ret float32
		return ret
	}
	return *o.SaleCommission
}

// GetSaleCommissionOk returns a tuple with the SaleCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetSaleCommissionOk() (*float32, bool) {
	if o == nil || IsNil(o.SaleCommission) {
		return nil, false
	}
	return o.SaleCommission, true
}

// HasSaleCommission returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasSaleCommission() bool {
	if o != nil && !IsNil(o.SaleCommission) {
		return true
	}

	return false
}

// SetSaleCommission gets a reference to the given float32 and assigns it to the SaleCommission field.
func (o *ResponseIpoSale) SetSaleCommission(v float32) {
	o.SaleCommission = &v
}

// GetSaleDetails returns the SaleDetails field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetSaleDetails() []int32 {
	if o == nil || IsNil(o.SaleDetails) {
		var ret []int32
		return ret
	}
	return o.SaleDetails
}

// GetSaleDetailsOk returns a tuple with the SaleDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetSaleDetailsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SaleDetails) {
		return nil, false
	}
	return o.SaleDetails, true
}

// HasSaleDetails returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasSaleDetails() bool {
	if o != nil && !IsNil(o.SaleDetails) {
		return true
	}

	return false
}

// SetSaleDetails gets a reference to the given []int32 and assigns it to the SaleDetails field.
func (o *ResponseIpoSale) SetSaleDetails(v []int32) {
	o.SaleDetails = v
}

// GetSoldTo returns the SoldTo field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetSoldTo() string {
	if o == nil || IsNil(o.SoldTo) {
		var ret string
		return ret
	}
	return *o.SoldTo
}

// GetSoldToOk returns a tuple with the SoldTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetSoldToOk() (*string, bool) {
	if o == nil || IsNil(o.SoldTo) {
		return nil, false
	}
	return o.SoldTo, true
}

// HasSoldTo returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasSoldTo() bool {
	if o != nil && !IsNil(o.SoldTo) {
		return true
	}

	return false
}

// SetSoldTo gets a reference to the given string and assigns it to the SoldTo field.
func (o *ResponseIpoSale) SetSoldTo(v string) {
	o.SoldTo = &v
}

// GetTransDate returns the TransDate field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetTransDate() string {
	if o == nil || IsNil(o.TransDate) {
		var ret string
		return ret
	}
	return *o.TransDate
}

// GetTransDateOk returns a tuple with the TransDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetTransDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransDate) {
		return nil, false
	}
	return o.TransDate, true
}

// HasTransDate returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasTransDate() bool {
	if o != nil && !IsNil(o.TransDate) {
		return true
	}

	return false
}

// SetTransDate gets a reference to the given string and assigns it to the TransDate field.
func (o *ResponseIpoSale) SetTransDate(v string) {
	o.TransDate = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *ResponseIpoSale) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransdate returns the Transdate field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetTransdate() string {
	if o == nil || IsNil(o.Transdate) {
		var ret string
		return ret
	}
	return *o.Transdate
}

// GetTransdateOk returns a tuple with the Transdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetTransdateOk() (*string, bool) {
	if o == nil || IsNil(o.Transdate) {
		return nil, false
	}
	return o.Transdate, true
}

// HasTransdate returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasTransdate() bool {
	if o != nil && !IsNil(o.Transdate) {
		return true
	}

	return false
}

// SetTransdate gets a reference to the given string and assigns it to the Transdate field.
func (o *ResponseIpoSale) SetTransdate(v string) {
	o.Transdate = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ResponseIpoSale) SetUserId(v int32) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ResponseIpoSale) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseIpoSale) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ResponseIpoSale) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ResponseIpoSale) SetUserName(v string) {
	o.UserName = &v
}

func (o ResponseIpoSale) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseIpoSale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApprovedDate) {
		toSerialize["approved_date"] = o.ApprovedDate
	}
	if !IsNil(o.ApproverId) {
		toSerialize["approver_id"] = o.ApproverId
	}
	if !IsNil(o.ApproverName) {
		toSerialize["approver_name"] = o.ApproverName
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
	if !IsNil(o.OfficeName) {
		toSerialize["office_name"] = o.OfficeName
	}
	if !IsNil(o.SaleAmount) {
		toSerialize["sale_amount"] = o.SaleAmount
	}
	if !IsNil(o.SaleCommission) {
		toSerialize["sale_commission"] = o.SaleCommission
	}
	if !IsNil(o.SaleDetails) {
		toSerialize["sale_details"] = o.SaleDetails
	}
	if !IsNil(o.SoldTo) {
		toSerialize["sold_to"] = o.SoldTo
	}
	if !IsNil(o.TransDate) {
		toSerialize["trans_date"] = o.TransDate
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !IsNil(o.Transdate) {
		toSerialize["transdate"] = o.Transdate
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	return toSerialize, nil
}

type NullableResponseIpoSale struct {
	value *ResponseIpoSale
	isSet bool
}

func (v NullableResponseIpoSale) Get() *ResponseIpoSale {
	return v.value
}

func (v *NullableResponseIpoSale) Set(val *ResponseIpoSale) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseIpoSale) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseIpoSale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseIpoSale(val *ResponseIpoSale) *NullableResponseIpoSale {
	return &NullableResponseIpoSale{value: val, isSet: true}
}

func (v NullableResponseIpoSale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseIpoSale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



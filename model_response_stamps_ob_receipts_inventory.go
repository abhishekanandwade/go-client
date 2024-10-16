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

// checks if the ResponseStampsObReceiptsInventory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseStampsObReceiptsInventory{}

// ResponseStampsObReceiptsInventory struct for ResponseStampsObReceiptsInventory
type ResponseStampsObReceiptsInventory struct {
	InvoiceDate *string `json:"invoice_date,omitempty"`
	OfficeId *int32 `json:"office_id,omitempty"`
	RecdUserId *int32 `json:"recd_user_id,omitempty"`
	ReceiptAmt *float32 `json:"receipt_amt,omitempty"`
	ReceiptDetails *map[string]int32 `json:"receipt_details,omitempty"`
	ReceiptInvoice *string `json:"receipt_invoice,omitempty"`
	ReceiptSource *string `json:"receipt_source,omitempty"`
	Remarks *string `json:"remarks,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	SubRequestId *int32 `json:"sub_request_id,omitempty"`
}

// NewResponseStampsObReceiptsInventory instantiates a new ResponseStampsObReceiptsInventory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseStampsObReceiptsInventory() *ResponseStampsObReceiptsInventory {
	this := ResponseStampsObReceiptsInventory{}
	return &this
}

// NewResponseStampsObReceiptsInventoryWithDefaults instantiates a new ResponseStampsObReceiptsInventory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseStampsObReceiptsInventoryWithDefaults() *ResponseStampsObReceiptsInventory {
	this := ResponseStampsObReceiptsInventory{}
	return &this
}

// GetInvoiceDate returns the InvoiceDate field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetInvoiceDate() string {
	if o == nil || IsNil(o.InvoiceDate) {
		var ret string
		return ret
	}
	return *o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetInvoiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceDate) {
		return nil, false
	}
	return o.InvoiceDate, true
}

// HasInvoiceDate returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasInvoiceDate() bool {
	if o != nil && !IsNil(o.InvoiceDate) {
		return true
	}

	return false
}

// SetInvoiceDate gets a reference to the given string and assigns it to the InvoiceDate field.
func (o *ResponseStampsObReceiptsInventory) SetInvoiceDate(v string) {
	o.InvoiceDate = &v
}

// GetOfficeId returns the OfficeId field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetOfficeId() int32 {
	if o == nil || IsNil(o.OfficeId) {
		var ret int32
		return ret
	}
	return *o.OfficeId
}

// GetOfficeIdOk returns a tuple with the OfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OfficeId) {
		return nil, false
	}
	return o.OfficeId, true
}

// HasOfficeId returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasOfficeId() bool {
	if o != nil && !IsNil(o.OfficeId) {
		return true
	}

	return false
}

// SetOfficeId gets a reference to the given int32 and assigns it to the OfficeId field.
func (o *ResponseStampsObReceiptsInventory) SetOfficeId(v int32) {
	o.OfficeId = &v
}

// GetRecdUserId returns the RecdUserId field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetRecdUserId() int32 {
	if o == nil || IsNil(o.RecdUserId) {
		var ret int32
		return ret
	}
	return *o.RecdUserId
}

// GetRecdUserIdOk returns a tuple with the RecdUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetRecdUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RecdUserId) {
		return nil, false
	}
	return o.RecdUserId, true
}

// HasRecdUserId returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasRecdUserId() bool {
	if o != nil && !IsNil(o.RecdUserId) {
		return true
	}

	return false
}

// SetRecdUserId gets a reference to the given int32 and assigns it to the RecdUserId field.
func (o *ResponseStampsObReceiptsInventory) SetRecdUserId(v int32) {
	o.RecdUserId = &v
}

// GetReceiptAmt returns the ReceiptAmt field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetReceiptAmt() float32 {
	if o == nil || IsNil(o.ReceiptAmt) {
		var ret float32
		return ret
	}
	return *o.ReceiptAmt
}

// GetReceiptAmtOk returns a tuple with the ReceiptAmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetReceiptAmtOk() (*float32, bool) {
	if o == nil || IsNil(o.ReceiptAmt) {
		return nil, false
	}
	return o.ReceiptAmt, true
}

// HasReceiptAmt returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasReceiptAmt() bool {
	if o != nil && !IsNil(o.ReceiptAmt) {
		return true
	}

	return false
}

// SetReceiptAmt gets a reference to the given float32 and assigns it to the ReceiptAmt field.
func (o *ResponseStampsObReceiptsInventory) SetReceiptAmt(v float32) {
	o.ReceiptAmt = &v
}

// GetReceiptDetails returns the ReceiptDetails field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetReceiptDetails() map[string]int32 {
	if o == nil || IsNil(o.ReceiptDetails) {
		var ret map[string]int32
		return ret
	}
	return *o.ReceiptDetails
}

// GetReceiptDetailsOk returns a tuple with the ReceiptDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetReceiptDetailsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.ReceiptDetails) {
		return nil, false
	}
	return o.ReceiptDetails, true
}

// HasReceiptDetails returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasReceiptDetails() bool {
	if o != nil && !IsNil(o.ReceiptDetails) {
		return true
	}

	return false
}

// SetReceiptDetails gets a reference to the given map[string]int32 and assigns it to the ReceiptDetails field.
func (o *ResponseStampsObReceiptsInventory) SetReceiptDetails(v map[string]int32) {
	o.ReceiptDetails = &v
}

// GetReceiptInvoice returns the ReceiptInvoice field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetReceiptInvoice() string {
	if o == nil || IsNil(o.ReceiptInvoice) {
		var ret string
		return ret
	}
	return *o.ReceiptInvoice
}

// GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetReceiptInvoiceOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptInvoice) {
		return nil, false
	}
	return o.ReceiptInvoice, true
}

// HasReceiptInvoice returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasReceiptInvoice() bool {
	if o != nil && !IsNil(o.ReceiptInvoice) {
		return true
	}

	return false
}

// SetReceiptInvoice gets a reference to the given string and assigns it to the ReceiptInvoice field.
func (o *ResponseStampsObReceiptsInventory) SetReceiptInvoice(v string) {
	o.ReceiptInvoice = &v
}

// GetReceiptSource returns the ReceiptSource field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetReceiptSource() string {
	if o == nil || IsNil(o.ReceiptSource) {
		var ret string
		return ret
	}
	return *o.ReceiptSource
}

// GetReceiptSourceOk returns a tuple with the ReceiptSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetReceiptSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptSource) {
		return nil, false
	}
	return o.ReceiptSource, true
}

// HasReceiptSource returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasReceiptSource() bool {
	if o != nil && !IsNil(o.ReceiptSource) {
		return true
	}

	return false
}

// SetReceiptSource gets a reference to the given string and assigns it to the ReceiptSource field.
func (o *ResponseStampsObReceiptsInventory) SetReceiptSource(v string) {
	o.ReceiptSource = &v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *ResponseStampsObReceiptsInventory) SetRemarks(v string) {
	o.Remarks = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ResponseStampsObReceiptsInventory) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSubRequestId returns the SubRequestId field value if set, zero value otherwise.
func (o *ResponseStampsObReceiptsInventory) GetSubRequestId() int32 {
	if o == nil || IsNil(o.SubRequestId) {
		var ret int32
		return ret
	}
	return *o.SubRequestId
}

// GetSubRequestIdOk returns a tuple with the SubRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseStampsObReceiptsInventory) GetSubRequestIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubRequestId) {
		return nil, false
	}
	return o.SubRequestId, true
}

// HasSubRequestId returns a boolean if a field has been set.
func (o *ResponseStampsObReceiptsInventory) HasSubRequestId() bool {
	if o != nil && !IsNil(o.SubRequestId) {
		return true
	}

	return false
}

// SetSubRequestId gets a reference to the given int32 and assigns it to the SubRequestId field.
func (o *ResponseStampsObReceiptsInventory) SetSubRequestId(v int32) {
	o.SubRequestId = &v
}

func (o ResponseStampsObReceiptsInventory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseStampsObReceiptsInventory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvoiceDate) {
		toSerialize["invoice_date"] = o.InvoiceDate
	}
	if !IsNil(o.OfficeId) {
		toSerialize["office_id"] = o.OfficeId
	}
	if !IsNil(o.RecdUserId) {
		toSerialize["recd_user_id"] = o.RecdUserId
	}
	if !IsNil(o.ReceiptAmt) {
		toSerialize["receipt_amt"] = o.ReceiptAmt
	}
	if !IsNil(o.ReceiptDetails) {
		toSerialize["receipt_details"] = o.ReceiptDetails
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
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.SubRequestId) {
		toSerialize["sub_request_id"] = o.SubRequestId
	}
	return toSerialize, nil
}

type NullableResponseStampsObReceiptsInventory struct {
	value *ResponseStampsObReceiptsInventory
	isSet bool
}

func (v NullableResponseStampsObReceiptsInventory) Get() *ResponseStampsObReceiptsInventory {
	return v.value
}

func (v *NullableResponseStampsObReceiptsInventory) Set(val *ResponseStampsObReceiptsInventory) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseStampsObReceiptsInventory) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseStampsObReceiptsInventory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseStampsObReceiptsInventory(val *ResponseStampsObReceiptsInventory) *NullableResponseStampsObReceiptsInventory {
	return &NullableResponseStampsObReceiptsInventory{value: val, isSet: true}
}

func (v NullableResponseStampsObReceiptsInventory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseStampsObReceiptsInventory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



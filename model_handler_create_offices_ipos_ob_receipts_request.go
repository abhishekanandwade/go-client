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
	"bytes"
	"fmt"
)

// checks if the HandlerCreateOfficesIPOsOBReceiptsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateOfficesIPOsOBReceiptsRequest{}

// HandlerCreateOfficesIPOsOBReceiptsRequest struct for HandlerCreateOfficesIPOsOBReceiptsRequest
type HandlerCreateOfficesIPOsOBReceiptsRequest struct {
	InvoiceDate string `json:"invoice_date"`
	IpoDetails []HandlerIPODetail `json:"ipo_details,omitempty"`
	RecdUserId int32 `json:"recd_user_id"`
	ReceiptInvoice string `json:"receipt_invoice"`
	ReceiptSource string `json:"receipt_source"`
	Remarks *string `json:"remarks,omitempty"`
}

type _HandlerCreateOfficesIPOsOBReceiptsRequest HandlerCreateOfficesIPOsOBReceiptsRequest

// NewHandlerCreateOfficesIPOsOBReceiptsRequest instantiates a new HandlerCreateOfficesIPOsOBReceiptsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateOfficesIPOsOBReceiptsRequest(invoiceDate string, recdUserId int32, receiptInvoice string, receiptSource string) *HandlerCreateOfficesIPOsOBReceiptsRequest {
	this := HandlerCreateOfficesIPOsOBReceiptsRequest{}
	this.InvoiceDate = invoiceDate
	this.RecdUserId = recdUserId
	this.ReceiptInvoice = receiptInvoice
	this.ReceiptSource = receiptSource
	return &this
}

// NewHandlerCreateOfficesIPOsOBReceiptsRequestWithDefaults instantiates a new HandlerCreateOfficesIPOsOBReceiptsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateOfficesIPOsOBReceiptsRequestWithDefaults() *HandlerCreateOfficesIPOsOBReceiptsRequest {
	this := HandlerCreateOfficesIPOsOBReceiptsRequest{}
	return &this
}

// GetInvoiceDate returns the InvoiceDate field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetInvoiceDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceDate
}

// GetInvoiceDateOk returns a tuple with the InvoiceDate field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetInvoiceDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceDate, true
}

// SetInvoiceDate sets field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetInvoiceDate(v string) {
	o.InvoiceDate = v
}

// GetIpoDetails returns the IpoDetails field value if set, zero value otherwise.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetIpoDetails() []HandlerIPODetail {
	if o == nil || IsNil(o.IpoDetails) {
		var ret []HandlerIPODetail
		return ret
	}
	return o.IpoDetails
}

// GetIpoDetailsOk returns a tuple with the IpoDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetIpoDetailsOk() ([]HandlerIPODetail, bool) {
	if o == nil || IsNil(o.IpoDetails) {
		return nil, false
	}
	return o.IpoDetails, true
}

// HasIpoDetails returns a boolean if a field has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) HasIpoDetails() bool {
	if o != nil && !IsNil(o.IpoDetails) {
		return true
	}

	return false
}

// SetIpoDetails gets a reference to the given []HandlerIPODetail and assigns it to the IpoDetails field.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetIpoDetails(v []HandlerIPODetail) {
	o.IpoDetails = v
}

// GetRecdUserId returns the RecdUserId field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRecdUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecdUserId
}

// GetRecdUserIdOk returns a tuple with the RecdUserId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRecdUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecdUserId, true
}

// SetRecdUserId sets field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetRecdUserId(v int32) {
	o.RecdUserId = v
}

// GetReceiptInvoice returns the ReceiptInvoice field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptInvoice() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiptInvoice
}

// GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptInvoiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptInvoice, true
}

// SetReceiptInvoice sets field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetReceiptInvoice(v string) {
	o.ReceiptInvoice = v
}

// GetReceiptSource returns the ReceiptSource field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiptSource
}

// GetReceiptSourceOk returns a tuple with the ReceiptSource field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiptSource, true
}

// SetReceiptSource sets field value
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetReceiptSource(v string) {
	o.ReceiptSource = v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetRemarks(v string) {
	o.Remarks = &v
}

func (o HandlerCreateOfficesIPOsOBReceiptsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateOfficesIPOsOBReceiptsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["invoice_date"] = o.InvoiceDate
	if !IsNil(o.IpoDetails) {
		toSerialize["ipo_details"] = o.IpoDetails
	}
	toSerialize["recd_user_id"] = o.RecdUserId
	toSerialize["receipt_invoice"] = o.ReceiptInvoice
	toSerialize["receipt_source"] = o.ReceiptSource
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	return toSerialize, nil
}

func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"invoice_date",
		"recd_user_id",
		"receipt_invoice",
		"receipt_source",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHandlerCreateOfficesIPOsOBReceiptsRequest := _HandlerCreateOfficesIPOsOBReceiptsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateOfficesIPOsOBReceiptsRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateOfficesIPOsOBReceiptsRequest(varHandlerCreateOfficesIPOsOBReceiptsRequest)

	return err
}

type NullableHandlerCreateOfficesIPOsOBReceiptsRequest struct {
	value *HandlerCreateOfficesIPOsOBReceiptsRequest
	isSet bool
}

func (v NullableHandlerCreateOfficesIPOsOBReceiptsRequest) Get() *HandlerCreateOfficesIPOsOBReceiptsRequest {
	return v.value
}

func (v *NullableHandlerCreateOfficesIPOsOBReceiptsRequest) Set(val *HandlerCreateOfficesIPOsOBReceiptsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateOfficesIPOsOBReceiptsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateOfficesIPOsOBReceiptsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateOfficesIPOsOBReceiptsRequest(val *HandlerCreateOfficesIPOsOBReceiptsRequest) *NullableHandlerCreateOfficesIPOsOBReceiptsRequest {
	return &NullableHandlerCreateOfficesIPOsOBReceiptsRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateOfficesIPOsOBReceiptsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateOfficesIPOsOBReceiptsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


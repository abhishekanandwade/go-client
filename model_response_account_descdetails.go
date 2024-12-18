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

// checks if the ResponseAccountDescdetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseAccountDescdetails{}

// ResponseAccountDescdetails struct for ResponseAccountDescdetails
type ResponseAccountDescdetails struct {
	AccountCode *string `json:"account_code,omitempty"`
	AccountCodeDescription *string `json:"account_code_description,omitempty"`
	Amount *float32 `json:"amount,omitempty"`
	Hoa *string `json:"hoa,omitempty"`
	Part *string `json:"part,omitempty"`
	PositiveOrNegative *string `json:"positive_or_negative,omitempty"`
	ReceiptOrPayment *string `json:"receipt_or_payment,omitempty"`
}

// NewResponseAccountDescdetails instantiates a new ResponseAccountDescdetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseAccountDescdetails() *ResponseAccountDescdetails {
	this := ResponseAccountDescdetails{}
	return &this
}

// NewResponseAccountDescdetailsWithDefaults instantiates a new ResponseAccountDescdetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseAccountDescdetailsWithDefaults() *ResponseAccountDescdetails {
	this := ResponseAccountDescdetails{}
	return &this
}

// GetAccountCode returns the AccountCode field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetAccountCode() string {
	if o == nil || IsNil(o.AccountCode) {
		var ret string
		return ret
	}
	return *o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetAccountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCode) {
		return nil, false
	}
	return o.AccountCode, true
}

// HasAccountCode returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasAccountCode() bool {
	if o != nil && !IsNil(o.AccountCode) {
		return true
	}

	return false
}

// SetAccountCode gets a reference to the given string and assigns it to the AccountCode field.
func (o *ResponseAccountDescdetails) SetAccountCode(v string) {
	o.AccountCode = &v
}

// GetAccountCodeDescription returns the AccountCodeDescription field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetAccountCodeDescription() string {
	if o == nil || IsNil(o.AccountCodeDescription) {
		var ret string
		return ret
	}
	return *o.AccountCodeDescription
}

// GetAccountCodeDescriptionOk returns a tuple with the AccountCodeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetAccountCodeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.AccountCodeDescription) {
		return nil, false
	}
	return o.AccountCodeDescription, true
}

// HasAccountCodeDescription returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasAccountCodeDescription() bool {
	if o != nil && !IsNil(o.AccountCodeDescription) {
		return true
	}

	return false
}

// SetAccountCodeDescription gets a reference to the given string and assigns it to the AccountCodeDescription field.
func (o *ResponseAccountDescdetails) SetAccountCodeDescription(v string) {
	o.AccountCodeDescription = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *ResponseAccountDescdetails) SetAmount(v float32) {
	o.Amount = &v
}

// GetHoa returns the Hoa field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetHoa() string {
	if o == nil || IsNil(o.Hoa) {
		var ret string
		return ret
	}
	return *o.Hoa
}

// GetHoaOk returns a tuple with the Hoa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetHoaOk() (*string, bool) {
	if o == nil || IsNil(o.Hoa) {
		return nil, false
	}
	return o.Hoa, true
}

// HasHoa returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasHoa() bool {
	if o != nil && !IsNil(o.Hoa) {
		return true
	}

	return false
}

// SetHoa gets a reference to the given string and assigns it to the Hoa field.
func (o *ResponseAccountDescdetails) SetHoa(v string) {
	o.Hoa = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *ResponseAccountDescdetails) SetPart(v string) {
	o.Part = &v
}

// GetPositiveOrNegative returns the PositiveOrNegative field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetPositiveOrNegative() string {
	if o == nil || IsNil(o.PositiveOrNegative) {
		var ret string
		return ret
	}
	return *o.PositiveOrNegative
}

// GetPositiveOrNegativeOk returns a tuple with the PositiveOrNegative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetPositiveOrNegativeOk() (*string, bool) {
	if o == nil || IsNil(o.PositiveOrNegative) {
		return nil, false
	}
	return o.PositiveOrNegative, true
}

// HasPositiveOrNegative returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasPositiveOrNegative() bool {
	if o != nil && !IsNil(o.PositiveOrNegative) {
		return true
	}

	return false
}

// SetPositiveOrNegative gets a reference to the given string and assigns it to the PositiveOrNegative field.
func (o *ResponseAccountDescdetails) SetPositiveOrNegative(v string) {
	o.PositiveOrNegative = &v
}

// GetReceiptOrPayment returns the ReceiptOrPayment field value if set, zero value otherwise.
func (o *ResponseAccountDescdetails) GetReceiptOrPayment() string {
	if o == nil || IsNil(o.ReceiptOrPayment) {
		var ret string
		return ret
	}
	return *o.ReceiptOrPayment
}

// GetReceiptOrPaymentOk returns a tuple with the ReceiptOrPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseAccountDescdetails) GetReceiptOrPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptOrPayment) {
		return nil, false
	}
	return o.ReceiptOrPayment, true
}

// HasReceiptOrPayment returns a boolean if a field has been set.
func (o *ResponseAccountDescdetails) HasReceiptOrPayment() bool {
	if o != nil && !IsNil(o.ReceiptOrPayment) {
		return true
	}

	return false
}

// SetReceiptOrPayment gets a reference to the given string and assigns it to the ReceiptOrPayment field.
func (o *ResponseAccountDescdetails) SetReceiptOrPayment(v string) {
	o.ReceiptOrPayment = &v
}

func (o ResponseAccountDescdetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseAccountDescdetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountCode) {
		toSerialize["account_code"] = o.AccountCode
	}
	if !IsNil(o.AccountCodeDescription) {
		toSerialize["account_code_description"] = o.AccountCodeDescription
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Hoa) {
		toSerialize["hoa"] = o.Hoa
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !IsNil(o.PositiveOrNegative) {
		toSerialize["positive_or_negative"] = o.PositiveOrNegative
	}
	if !IsNil(o.ReceiptOrPayment) {
		toSerialize["receipt_or_payment"] = o.ReceiptOrPayment
	}
	return toSerialize, nil
}

type NullableResponseAccountDescdetails struct {
	value *ResponseAccountDescdetails
	isSet bool
}

func (v NullableResponseAccountDescdetails) Get() *ResponseAccountDescdetails {
	return v.value
}

func (v *NullableResponseAccountDescdetails) Set(val *ResponseAccountDescdetails) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseAccountDescdetails) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseAccountDescdetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseAccountDescdetails(val *ResponseAccountDescdetails) *NullableResponseAccountDescdetails {
	return &NullableResponseAccountDescdetails{value: val, isSet: true}
}

func (v NullableResponseAccountDescdetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseAccountDescdetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



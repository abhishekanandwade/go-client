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

// checks if the HandlerCreateStampsTransactionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerCreateStampsTransactionsRequest{}

// HandlerCreateStampsTransactionsRequest struct for HandlerCreateStampsTransactionsRequest
type HandlerCreateStampsTransactionsRequest struct {
	IsSingleHand *bool `json:"is_single_hand,omitempty"`
	IsSpecialRemittance *bool `json:"is_special_remittance,omitempty"`
	IssOfficeId int32 `json:"iss_office_id"`
	Remarks *string `json:"remarks,omitempty"`
	ReqAmount float32 `json:"req_amount"`
	ReqDetails map[string]int32 `json:"req_details"`
	ReqOfficeId int32 `json:"req_office_id"`
	ReqUserId int32 `json:"req_user_id"`
	RequestId *string `json:"request_id,omitempty"`
	RequestSource string `json:"request_source"`
	RequestType string `json:"request_type"`
}

type _HandlerCreateStampsTransactionsRequest HandlerCreateStampsTransactionsRequest

// NewHandlerCreateStampsTransactionsRequest instantiates a new HandlerCreateStampsTransactionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerCreateStampsTransactionsRequest(issOfficeId int32, reqAmount float32, reqDetails map[string]int32, reqOfficeId int32, reqUserId int32, requestSource string, requestType string) *HandlerCreateStampsTransactionsRequest {
	this := HandlerCreateStampsTransactionsRequest{}
	this.IssOfficeId = issOfficeId
	this.ReqAmount = reqAmount
	this.ReqDetails = reqDetails
	this.ReqOfficeId = reqOfficeId
	this.ReqUserId = reqUserId
	this.RequestSource = requestSource
	this.RequestType = requestType
	return &this
}

// NewHandlerCreateStampsTransactionsRequestWithDefaults instantiates a new HandlerCreateStampsTransactionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerCreateStampsTransactionsRequestWithDefaults() *HandlerCreateStampsTransactionsRequest {
	this := HandlerCreateStampsTransactionsRequest{}
	return &this
}

// GetIsSingleHand returns the IsSingleHand field value if set, zero value otherwise.
func (o *HandlerCreateStampsTransactionsRequest) GetIsSingleHand() bool {
	if o == nil || IsNil(o.IsSingleHand) {
		var ret bool
		return ret
	}
	return *o.IsSingleHand
}

// GetIsSingleHandOk returns a tuple with the IsSingleHand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetIsSingleHandOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSingleHand) {
		return nil, false
	}
	return o.IsSingleHand, true
}

// HasIsSingleHand returns a boolean if a field has been set.
func (o *HandlerCreateStampsTransactionsRequest) HasIsSingleHand() bool {
	if o != nil && !IsNil(o.IsSingleHand) {
		return true
	}

	return false
}

// SetIsSingleHand gets a reference to the given bool and assigns it to the IsSingleHand field.
func (o *HandlerCreateStampsTransactionsRequest) SetIsSingleHand(v bool) {
	o.IsSingleHand = &v
}

// GetIsSpecialRemittance returns the IsSpecialRemittance field value if set, zero value otherwise.
func (o *HandlerCreateStampsTransactionsRequest) GetIsSpecialRemittance() bool {
	if o == nil || IsNil(o.IsSpecialRemittance) {
		var ret bool
		return ret
	}
	return *o.IsSpecialRemittance
}

// GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetIsSpecialRemittanceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSpecialRemittance) {
		return nil, false
	}
	return o.IsSpecialRemittance, true
}

// HasIsSpecialRemittance returns a boolean if a field has been set.
func (o *HandlerCreateStampsTransactionsRequest) HasIsSpecialRemittance() bool {
	if o != nil && !IsNil(o.IsSpecialRemittance) {
		return true
	}

	return false
}

// SetIsSpecialRemittance gets a reference to the given bool and assigns it to the IsSpecialRemittance field.
func (o *HandlerCreateStampsTransactionsRequest) SetIsSpecialRemittance(v bool) {
	o.IsSpecialRemittance = &v
}

// GetIssOfficeId returns the IssOfficeId field value
func (o *HandlerCreateStampsTransactionsRequest) GetIssOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IssOfficeId
}

// GetIssOfficeIdOk returns a tuple with the IssOfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetIssOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssOfficeId, true
}

// SetIssOfficeId sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetIssOfficeId(v int32) {
	o.IssOfficeId = v
}

// GetRemarks returns the Remarks field value if set, zero value otherwise.
func (o *HandlerCreateStampsTransactionsRequest) GetRemarks() string {
	if o == nil || IsNil(o.Remarks) {
		var ret string
		return ret
	}
	return *o.Remarks
}

// GetRemarksOk returns a tuple with the Remarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetRemarksOk() (*string, bool) {
	if o == nil || IsNil(o.Remarks) {
		return nil, false
	}
	return o.Remarks, true
}

// HasRemarks returns a boolean if a field has been set.
func (o *HandlerCreateStampsTransactionsRequest) HasRemarks() bool {
	if o != nil && !IsNil(o.Remarks) {
		return true
	}

	return false
}

// SetRemarks gets a reference to the given string and assigns it to the Remarks field.
func (o *HandlerCreateStampsTransactionsRequest) SetRemarks(v string) {
	o.Remarks = &v
}

// GetReqAmount returns the ReqAmount field value
func (o *HandlerCreateStampsTransactionsRequest) GetReqAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ReqAmount
}

// GetReqAmountOk returns a tuple with the ReqAmount field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetReqAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqAmount, true
}

// SetReqAmount sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetReqAmount(v float32) {
	o.ReqAmount = v
}

// GetReqDetails returns the ReqDetails field value
func (o *HandlerCreateStampsTransactionsRequest) GetReqDetails() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.ReqDetails
}

// GetReqDetailsOk returns a tuple with the ReqDetails field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetReqDetailsOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqDetails, true
}

// SetReqDetails sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetReqDetails(v map[string]int32) {
	o.ReqDetails = v
}

// GetReqOfficeId returns the ReqOfficeId field value
func (o *HandlerCreateStampsTransactionsRequest) GetReqOfficeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReqOfficeId
}

// GetReqOfficeIdOk returns a tuple with the ReqOfficeId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetReqOfficeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqOfficeId, true
}

// SetReqOfficeId sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetReqOfficeId(v int32) {
	o.ReqOfficeId = v
}

// GetReqUserId returns the ReqUserId field value
func (o *HandlerCreateStampsTransactionsRequest) GetReqUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReqUserId
}

// GetReqUserIdOk returns a tuple with the ReqUserId field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetReqUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReqUserId, true
}

// SetReqUserId sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetReqUserId(v int32) {
	o.ReqUserId = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *HandlerCreateStampsTransactionsRequest) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *HandlerCreateStampsTransactionsRequest) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *HandlerCreateStampsTransactionsRequest) SetRequestId(v string) {
	o.RequestId = &v
}

// GetRequestSource returns the RequestSource field value
func (o *HandlerCreateStampsTransactionsRequest) GetRequestSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestSource
}

// GetRequestSourceOk returns a tuple with the RequestSource field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetRequestSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSource, true
}

// SetRequestSource sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetRequestSource(v string) {
	o.RequestSource = v
}

// GetRequestType returns the RequestType field value
func (o *HandlerCreateStampsTransactionsRequest) GetRequestType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *HandlerCreateStampsTransactionsRequest) GetRequestTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *HandlerCreateStampsTransactionsRequest) SetRequestType(v string) {
	o.RequestType = v
}

func (o HandlerCreateStampsTransactionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerCreateStampsTransactionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSingleHand) {
		toSerialize["is_single_hand"] = o.IsSingleHand
	}
	if !IsNil(o.IsSpecialRemittance) {
		toSerialize["is_special_remittance"] = o.IsSpecialRemittance
	}
	toSerialize["iss_office_id"] = o.IssOfficeId
	if !IsNil(o.Remarks) {
		toSerialize["remarks"] = o.Remarks
	}
	toSerialize["req_amount"] = o.ReqAmount
	toSerialize["req_details"] = o.ReqDetails
	toSerialize["req_office_id"] = o.ReqOfficeId
	toSerialize["req_user_id"] = o.ReqUserId
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	toSerialize["request_source"] = o.RequestSource
	toSerialize["request_type"] = o.RequestType
	return toSerialize, nil
}

func (o *HandlerCreateStampsTransactionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iss_office_id",
		"req_amount",
		"req_details",
		"req_office_id",
		"req_user_id",
		"request_source",
		"request_type",
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

	varHandlerCreateStampsTransactionsRequest := _HandlerCreateStampsTransactionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHandlerCreateStampsTransactionsRequest)

	if err != nil {
		return err
	}

	*o = HandlerCreateStampsTransactionsRequest(varHandlerCreateStampsTransactionsRequest)

	return err
}

type NullableHandlerCreateStampsTransactionsRequest struct {
	value *HandlerCreateStampsTransactionsRequest
	isSet bool
}

func (v NullableHandlerCreateStampsTransactionsRequest) Get() *HandlerCreateStampsTransactionsRequest {
	return v.value
}

func (v *NullableHandlerCreateStampsTransactionsRequest) Set(val *HandlerCreateStampsTransactionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerCreateStampsTransactionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerCreateStampsTransactionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerCreateStampsTransactionsRequest(val *HandlerCreateStampsTransactionsRequest) *NullableHandlerCreateStampsTransactionsRequest {
	return &NullableHandlerCreateStampsTransactionsRequest{value: val, isSet: true}
}

func (v NullableHandlerCreateStampsTransactionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerCreateStampsTransactionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


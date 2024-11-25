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

// checks if the HandlerTemporalCombinedRequestbagopenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandlerTemporalCombinedRequestbagopenRequest{}

// HandlerTemporalCombinedRequestbagopenRequest struct for HandlerTemporalCombinedRequestbagopenRequest
type HandlerTemporalCombinedRequestbagopenRequest struct {
	AcBagId *string `json:"ac_bag_id,omitempty"`
	ArticleCount *int32 `json:"article_count,omitempty"`
	ArticleWeight *float32 `json:"article_weight,omitempty"`
	BagCount *int32 `json:"bag_count,omitempty"`
	BagNumber *string `json:"bag_number,omitempty"`
	BagOpenArt []HandlerArticleInfo `json:"bag_open_art,omitempty"`
	BagType *string `json:"bag_type,omitempty"`
	BagUuid *string `json:"bag_uuid,omitempty"`
	BagWeight *float32 `json:"bag_weight,omitempty"`
	DeliveryType *string `json:"delivery_type,omitempty"`
	EventType *string `json:"event_type,omitempty"`
	FromOfficeId *int32 `json:"from_office_id,omitempty"`
	InsuredFlag *bool `json:"insured_flag,omitempty"`
	IsReceived *bool `json:"is_received,omitempty"`
	ReceivedBy *int32 `json:"received_by,omitempty"`
	ScheduleId *int32 `json:"schedule_id,omitempty"`
	SetDate *string `json:"set_date,omitempty"`
	SetNumber *string `json:"set_number,omitempty"`
	ToOfficeId *int32 `json:"to_office_id,omitempty"`
	UserId *int32 `json:"user_id,omitempty"`
}

// NewHandlerTemporalCombinedRequestbagopenRequest instantiates a new HandlerTemporalCombinedRequestbagopenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandlerTemporalCombinedRequestbagopenRequest() *HandlerTemporalCombinedRequestbagopenRequest {
	this := HandlerTemporalCombinedRequestbagopenRequest{}
	return &this
}

// NewHandlerTemporalCombinedRequestbagopenRequestWithDefaults instantiates a new HandlerTemporalCombinedRequestbagopenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandlerTemporalCombinedRequestbagopenRequestWithDefaults() *HandlerTemporalCombinedRequestbagopenRequest {
	this := HandlerTemporalCombinedRequestbagopenRequest{}
	return &this
}

// GetAcBagId returns the AcBagId field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetAcBagId() string {
	if o == nil || IsNil(o.AcBagId) {
		var ret string
		return ret
	}
	return *o.AcBagId
}

// GetAcBagIdOk returns a tuple with the AcBagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetAcBagIdOk() (*string, bool) {
	if o == nil || IsNil(o.AcBagId) {
		return nil, false
	}
	return o.AcBagId, true
}

// HasAcBagId returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasAcBagId() bool {
	if o != nil && !IsNil(o.AcBagId) {
		return true
	}

	return false
}

// SetAcBagId gets a reference to the given string and assigns it to the AcBagId field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetAcBagId(v string) {
	o.AcBagId = &v
}

// GetArticleCount returns the ArticleCount field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetArticleCount() int32 {
	if o == nil || IsNil(o.ArticleCount) {
		var ret int32
		return ret
	}
	return *o.ArticleCount
}

// GetArticleCountOk returns a tuple with the ArticleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetArticleCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ArticleCount) {
		return nil, false
	}
	return o.ArticleCount, true
}

// HasArticleCount returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasArticleCount() bool {
	if o != nil && !IsNil(o.ArticleCount) {
		return true
	}

	return false
}

// SetArticleCount gets a reference to the given int32 and assigns it to the ArticleCount field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetArticleCount(v int32) {
	o.ArticleCount = &v
}

// GetArticleWeight returns the ArticleWeight field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetArticleWeight() float32 {
	if o == nil || IsNil(o.ArticleWeight) {
		var ret float32
		return ret
	}
	return *o.ArticleWeight
}

// GetArticleWeightOk returns a tuple with the ArticleWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetArticleWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.ArticleWeight) {
		return nil, false
	}
	return o.ArticleWeight, true
}

// HasArticleWeight returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasArticleWeight() bool {
	if o != nil && !IsNil(o.ArticleWeight) {
		return true
	}

	return false
}

// SetArticleWeight gets a reference to the given float32 and assigns it to the ArticleWeight field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetArticleWeight(v float32) {
	o.ArticleWeight = &v
}

// GetBagCount returns the BagCount field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagCount() int32 {
	if o == nil || IsNil(o.BagCount) {
		var ret int32
		return ret
	}
	return *o.BagCount
}

// GetBagCountOk returns a tuple with the BagCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BagCount) {
		return nil, false
	}
	return o.BagCount, true
}

// HasBagCount returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagCount() bool {
	if o != nil && !IsNil(o.BagCount) {
		return true
	}

	return false
}

// SetBagCount gets a reference to the given int32 and assigns it to the BagCount field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagCount(v int32) {
	o.BagCount = &v
}

// GetBagNumber returns the BagNumber field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagNumber() string {
	if o == nil || IsNil(o.BagNumber) {
		var ret string
		return ret
	}
	return *o.BagNumber
}

// GetBagNumberOk returns a tuple with the BagNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagNumberOk() (*string, bool) {
	if o == nil || IsNil(o.BagNumber) {
		return nil, false
	}
	return o.BagNumber, true
}

// HasBagNumber returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagNumber() bool {
	if o != nil && !IsNil(o.BagNumber) {
		return true
	}

	return false
}

// SetBagNumber gets a reference to the given string and assigns it to the BagNumber field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagNumber(v string) {
	o.BagNumber = &v
}

// GetBagOpenArt returns the BagOpenArt field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagOpenArt() []HandlerArticleInfo {
	if o == nil || IsNil(o.BagOpenArt) {
		var ret []HandlerArticleInfo
		return ret
	}
	return o.BagOpenArt
}

// GetBagOpenArtOk returns a tuple with the BagOpenArt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagOpenArtOk() ([]HandlerArticleInfo, bool) {
	if o == nil || IsNil(o.BagOpenArt) {
		return nil, false
	}
	return o.BagOpenArt, true
}

// HasBagOpenArt returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagOpenArt() bool {
	if o != nil && !IsNil(o.BagOpenArt) {
		return true
	}

	return false
}

// SetBagOpenArt gets a reference to the given []HandlerArticleInfo and assigns it to the BagOpenArt field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagOpenArt(v []HandlerArticleInfo) {
	o.BagOpenArt = v
}

// GetBagType returns the BagType field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagType() string {
	if o == nil || IsNil(o.BagType) {
		var ret string
		return ret
	}
	return *o.BagType
}

// GetBagTypeOk returns a tuple with the BagType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BagType) {
		return nil, false
	}
	return o.BagType, true
}

// HasBagType returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagType() bool {
	if o != nil && !IsNil(o.BagType) {
		return true
	}

	return false
}

// SetBagType gets a reference to the given string and assigns it to the BagType field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagType(v string) {
	o.BagType = &v
}

// GetBagUuid returns the BagUuid field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagUuid() string {
	if o == nil || IsNil(o.BagUuid) {
		var ret string
		return ret
	}
	return *o.BagUuid
}

// GetBagUuidOk returns a tuple with the BagUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagUuidOk() (*string, bool) {
	if o == nil || IsNil(o.BagUuid) {
		return nil, false
	}
	return o.BagUuid, true
}

// HasBagUuid returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagUuid() bool {
	if o != nil && !IsNil(o.BagUuid) {
		return true
	}

	return false
}

// SetBagUuid gets a reference to the given string and assigns it to the BagUuid field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagUuid(v string) {
	o.BagUuid = &v
}

// GetBagWeight returns the BagWeight field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagWeight() float32 {
	if o == nil || IsNil(o.BagWeight) {
		var ret float32
		return ret
	}
	return *o.BagWeight
}

// GetBagWeightOk returns a tuple with the BagWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetBagWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.BagWeight) {
		return nil, false
	}
	return o.BagWeight, true
}

// HasBagWeight returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasBagWeight() bool {
	if o != nil && !IsNil(o.BagWeight) {
		return true
	}

	return false
}

// SetBagWeight gets a reference to the given float32 and assigns it to the BagWeight field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetBagWeight(v float32) {
	o.BagWeight = &v
}

// GetDeliveryType returns the DeliveryType field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetDeliveryType() string {
	if o == nil || IsNil(o.DeliveryType) {
		var ret string
		return ret
	}
	return *o.DeliveryType
}

// GetDeliveryTypeOk returns a tuple with the DeliveryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetDeliveryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryType) {
		return nil, false
	}
	return o.DeliveryType, true
}

// HasDeliveryType returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasDeliveryType() bool {
	if o != nil && !IsNil(o.DeliveryType) {
		return true
	}

	return false
}

// SetDeliveryType gets a reference to the given string and assigns it to the DeliveryType field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetDeliveryType(v string) {
	o.DeliveryType = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetEventType(v string) {
	o.EventType = &v
}

// GetFromOfficeId returns the FromOfficeId field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetFromOfficeId() int32 {
	if o == nil || IsNil(o.FromOfficeId) {
		var ret int32
		return ret
	}
	return *o.FromOfficeId
}

// GetFromOfficeIdOk returns a tuple with the FromOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetFromOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FromOfficeId) {
		return nil, false
	}
	return o.FromOfficeId, true
}

// HasFromOfficeId returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasFromOfficeId() bool {
	if o != nil && !IsNil(o.FromOfficeId) {
		return true
	}

	return false
}

// SetFromOfficeId gets a reference to the given int32 and assigns it to the FromOfficeId field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetFromOfficeId(v int32) {
	o.FromOfficeId = &v
}

// GetInsuredFlag returns the InsuredFlag field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetInsuredFlag() bool {
	if o == nil || IsNil(o.InsuredFlag) {
		var ret bool
		return ret
	}
	return *o.InsuredFlag
}

// GetInsuredFlagOk returns a tuple with the InsuredFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetInsuredFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.InsuredFlag) {
		return nil, false
	}
	return o.InsuredFlag, true
}

// HasInsuredFlag returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasInsuredFlag() bool {
	if o != nil && !IsNil(o.InsuredFlag) {
		return true
	}

	return false
}

// SetInsuredFlag gets a reference to the given bool and assigns it to the InsuredFlag field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetInsuredFlag(v bool) {
	o.InsuredFlag = &v
}

// GetIsReceived returns the IsReceived field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetIsReceived() bool {
	if o == nil || IsNil(o.IsReceived) {
		var ret bool
		return ret
	}
	return *o.IsReceived
}

// GetIsReceivedOk returns a tuple with the IsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetIsReceivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReceived) {
		return nil, false
	}
	return o.IsReceived, true
}

// HasIsReceived returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasIsReceived() bool {
	if o != nil && !IsNil(o.IsReceived) {
		return true
	}

	return false
}

// SetIsReceived gets a reference to the given bool and assigns it to the IsReceived field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetIsReceived(v bool) {
	o.IsReceived = &v
}

// GetReceivedBy returns the ReceivedBy field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetReceivedBy() int32 {
	if o == nil || IsNil(o.ReceivedBy) {
		var ret int32
		return ret
	}
	return *o.ReceivedBy
}

// GetReceivedByOk returns a tuple with the ReceivedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetReceivedByOk() (*int32, bool) {
	if o == nil || IsNil(o.ReceivedBy) {
		return nil, false
	}
	return o.ReceivedBy, true
}

// HasReceivedBy returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasReceivedBy() bool {
	if o != nil && !IsNil(o.ReceivedBy) {
		return true
	}

	return false
}

// SetReceivedBy gets a reference to the given int32 and assigns it to the ReceivedBy field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetReceivedBy(v int32) {
	o.ReceivedBy = &v
}

// GetScheduleId returns the ScheduleId field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetScheduleId() int32 {
	if o == nil || IsNil(o.ScheduleId) {
		var ret int32
		return ret
	}
	return *o.ScheduleId
}

// GetScheduleIdOk returns a tuple with the ScheduleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetScheduleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ScheduleId) {
		return nil, false
	}
	return o.ScheduleId, true
}

// HasScheduleId returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasScheduleId() bool {
	if o != nil && !IsNil(o.ScheduleId) {
		return true
	}

	return false
}

// SetScheduleId gets a reference to the given int32 and assigns it to the ScheduleId field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetScheduleId(v int32) {
	o.ScheduleId = &v
}

// GetSetDate returns the SetDate field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetSetDate() string {
	if o == nil || IsNil(o.SetDate) {
		var ret string
		return ret
	}
	return *o.SetDate
}

// GetSetDateOk returns a tuple with the SetDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetSetDateOk() (*string, bool) {
	if o == nil || IsNil(o.SetDate) {
		return nil, false
	}
	return o.SetDate, true
}

// HasSetDate returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasSetDate() bool {
	if o != nil && !IsNil(o.SetDate) {
		return true
	}

	return false
}

// SetSetDate gets a reference to the given string and assigns it to the SetDate field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetSetDate(v string) {
	o.SetDate = &v
}

// GetSetNumber returns the SetNumber field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetSetNumber() string {
	if o == nil || IsNil(o.SetNumber) {
		var ret string
		return ret
	}
	return *o.SetNumber
}

// GetSetNumberOk returns a tuple with the SetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetSetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SetNumber) {
		return nil, false
	}
	return o.SetNumber, true
}

// HasSetNumber returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasSetNumber() bool {
	if o != nil && !IsNil(o.SetNumber) {
		return true
	}

	return false
}

// SetSetNumber gets a reference to the given string and assigns it to the SetNumber field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetSetNumber(v string) {
	o.SetNumber = &v
}

// GetToOfficeId returns the ToOfficeId field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetToOfficeId() int32 {
	if o == nil || IsNil(o.ToOfficeId) {
		var ret int32
		return ret
	}
	return *o.ToOfficeId
}

// GetToOfficeIdOk returns a tuple with the ToOfficeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetToOfficeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ToOfficeId) {
		return nil, false
	}
	return o.ToOfficeId, true
}

// HasToOfficeId returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasToOfficeId() bool {
	if o != nil && !IsNil(o.ToOfficeId) {
		return true
	}

	return false
}

// SetToOfficeId gets a reference to the given int32 and assigns it to the ToOfficeId field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetToOfficeId(v int32) {
	o.ToOfficeId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *HandlerTemporalCombinedRequestbagopenRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *HandlerTemporalCombinedRequestbagopenRequest) SetUserId(v int32) {
	o.UserId = &v
}

func (o HandlerTemporalCombinedRequestbagopenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandlerTemporalCombinedRequestbagopenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcBagId) {
		toSerialize["ac_bag_id"] = o.AcBagId
	}
	if !IsNil(o.ArticleCount) {
		toSerialize["article_count"] = o.ArticleCount
	}
	if !IsNil(o.ArticleWeight) {
		toSerialize["article_weight"] = o.ArticleWeight
	}
	if !IsNil(o.BagCount) {
		toSerialize["bag_count"] = o.BagCount
	}
	if !IsNil(o.BagNumber) {
		toSerialize["bag_number"] = o.BagNumber
	}
	if !IsNil(o.BagOpenArt) {
		toSerialize["bag_open_art"] = o.BagOpenArt
	}
	if !IsNil(o.BagType) {
		toSerialize["bag_type"] = o.BagType
	}
	if !IsNil(o.BagUuid) {
		toSerialize["bag_uuid"] = o.BagUuid
	}
	if !IsNil(o.BagWeight) {
		toSerialize["bag_weight"] = o.BagWeight
	}
	if !IsNil(o.DeliveryType) {
		toSerialize["delivery_type"] = o.DeliveryType
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.FromOfficeId) {
		toSerialize["from_office_id"] = o.FromOfficeId
	}
	if !IsNil(o.InsuredFlag) {
		toSerialize["insured_flag"] = o.InsuredFlag
	}
	if !IsNil(o.IsReceived) {
		toSerialize["is_received"] = o.IsReceived
	}
	if !IsNil(o.ReceivedBy) {
		toSerialize["received_by"] = o.ReceivedBy
	}
	if !IsNil(o.ScheduleId) {
		toSerialize["schedule_id"] = o.ScheduleId
	}
	if !IsNil(o.SetDate) {
		toSerialize["set_date"] = o.SetDate
	}
	if !IsNil(o.SetNumber) {
		toSerialize["set_number"] = o.SetNumber
	}
	if !IsNil(o.ToOfficeId) {
		toSerialize["to_office_id"] = o.ToOfficeId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableHandlerTemporalCombinedRequestbagopenRequest struct {
	value *HandlerTemporalCombinedRequestbagopenRequest
	isSet bool
}

func (v NullableHandlerTemporalCombinedRequestbagopenRequest) Get() *HandlerTemporalCombinedRequestbagopenRequest {
	return v.value
}

func (v *NullableHandlerTemporalCombinedRequestbagopenRequest) Set(val *HandlerTemporalCombinedRequestbagopenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHandlerTemporalCombinedRequestbagopenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHandlerTemporalCombinedRequestbagopenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandlerTemporalCombinedRequestbagopenRequest(val *HandlerTemporalCombinedRequestbagopenRequest) *NullableHandlerTemporalCombinedRequestbagopenRequest {
	return &NullableHandlerTemporalCombinedRequestbagopenRequest{value: val, isSet: true}
}

func (v NullableHandlerTemporalCombinedRequestbagopenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandlerTemporalCombinedRequestbagopenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


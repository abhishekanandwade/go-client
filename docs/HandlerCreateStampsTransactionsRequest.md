# HandlerCreateStampsTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssOfficeId** | **int32** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | **float32** |  | 
**ReqDetails** | **map[string]int32** |  | 
**ReqOfficeId** | **int32** |  | 
**ReqUserId** | **int32** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | **string** |  | 
**RequestType** | **string** |  | 

## Methods

### NewHandlerCreateStampsTransactionsRequest

`func NewHandlerCreateStampsTransactionsRequest(issOfficeId int32, reqAmount float32, reqDetails map[string]int32, reqOfficeId int32, reqUserId int32, requestSource string, requestType string, ) *HandlerCreateStampsTransactionsRequest`

NewHandlerCreateStampsTransactionsRequest instantiates a new HandlerCreateStampsTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateStampsTransactionsRequestWithDefaults

`func NewHandlerCreateStampsTransactionsRequestWithDefaults() *HandlerCreateStampsTransactionsRequest`

NewHandlerCreateStampsTransactionsRequestWithDefaults instantiates a new HandlerCreateStampsTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSingleHand

`func (o *HandlerCreateStampsTransactionsRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerCreateStampsTransactionsRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerCreateStampsTransactionsRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerCreateStampsTransactionsRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *HandlerCreateStampsTransactionsRequest) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *HandlerCreateStampsTransactionsRequest) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *HandlerCreateStampsTransactionsRequest) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *HandlerCreateStampsTransactionsRequest) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *HandlerCreateStampsTransactionsRequest) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *HandlerCreateStampsTransactionsRequest) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *HandlerCreateStampsTransactionsRequest) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.


### GetRemarks

`func (o *HandlerCreateStampsTransactionsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateStampsTransactionsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateStampsTransactionsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateStampsTransactionsRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *HandlerCreateStampsTransactionsRequest) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *HandlerCreateStampsTransactionsRequest) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *HandlerCreateStampsTransactionsRequest) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.


### GetReqDetails

`func (o *HandlerCreateStampsTransactionsRequest) GetReqDetails() map[string]int32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *HandlerCreateStampsTransactionsRequest) GetReqDetailsOk() (*map[string]int32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *HandlerCreateStampsTransactionsRequest) SetReqDetails(v map[string]int32)`

SetReqDetails sets ReqDetails field to given value.


### GetReqOfficeId

`func (o *HandlerCreateStampsTransactionsRequest) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *HandlerCreateStampsTransactionsRequest) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *HandlerCreateStampsTransactionsRequest) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.


### GetReqUserId

`func (o *HandlerCreateStampsTransactionsRequest) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *HandlerCreateStampsTransactionsRequest) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *HandlerCreateStampsTransactionsRequest) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.


### GetRequestId

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HandlerCreateStampsTransactionsRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HandlerCreateStampsTransactionsRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *HandlerCreateStampsTransactionsRequest) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.


### GetRequestType

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerCreateStampsTransactionsRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerCreateStampsTransactionsRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



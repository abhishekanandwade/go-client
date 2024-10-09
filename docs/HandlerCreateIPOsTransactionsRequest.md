# HandlerCreateIPOsTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssOfficeId** | **int32** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | **float32** |  | 
**ReqDetails** | Pointer to **string** |  | [optional] 
**ReqOfficeId** | **int32** |  | 
**ReqUserId** | **int32** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | **string** |  | 
**RequestType** | **string** |  | 

## Methods

### NewHandlerCreateIPOsTransactionsRequest

`func NewHandlerCreateIPOsTransactionsRequest(issOfficeId int32, reqAmount float32, reqOfficeId int32, reqUserId int32, requestSource string, requestType string, ) *HandlerCreateIPOsTransactionsRequest`

NewHandlerCreateIPOsTransactionsRequest instantiates a new HandlerCreateIPOsTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateIPOsTransactionsRequestWithDefaults

`func NewHandlerCreateIPOsTransactionsRequestWithDefaults() *HandlerCreateIPOsTransactionsRequest`

NewHandlerCreateIPOsTransactionsRequestWithDefaults instantiates a new HandlerCreateIPOsTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSingleHand

`func (o *HandlerCreateIPOsTransactionsRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerCreateIPOsTransactionsRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerCreateIPOsTransactionsRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *HandlerCreateIPOsTransactionsRequest) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *HandlerCreateIPOsTransactionsRequest) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *HandlerCreateIPOsTransactionsRequest) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *HandlerCreateIPOsTransactionsRequest) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *HandlerCreateIPOsTransactionsRequest) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.


### GetRemarks

`func (o *HandlerCreateIPOsTransactionsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateIPOsTransactionsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateIPOsTransactionsRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *HandlerCreateIPOsTransactionsRequest) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.


### GetReqDetails

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqDetails() string`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqDetailsOk() (*string, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *HandlerCreateIPOsTransactionsRequest) SetReqDetails(v string)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *HandlerCreateIPOsTransactionsRequest) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *HandlerCreateIPOsTransactionsRequest) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.


### GetReqUserId

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *HandlerCreateIPOsTransactionsRequest) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.


### GetRequestId

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HandlerCreateIPOsTransactionsRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HandlerCreateIPOsTransactionsRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *HandlerCreateIPOsTransactionsRequest) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.


### GetRequestType

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerCreateIPOsTransactionsRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerCreateIPOsTransactionsRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerUpdateIPOsTransactionsStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedAmt** | Pointer to **float32** |  | [optional] 
**ApprovedDetails** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsNotApproved** | Pointer to **bool** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**IssUserId** | **int32** |  | 
**RequestType** | Pointer to **string** |  | [optional] 
**Type** | [**HandlerTransactionType**](HandlerTransactionType.md) |  | 

## Methods

### NewHandlerUpdateIPOsTransactionsStatusRequest

`func NewHandlerUpdateIPOsTransactionsStatusRequest(issUserId int32, type_ HandlerTransactionType, ) *HandlerUpdateIPOsTransactionsStatusRequest`

NewHandlerUpdateIPOsTransactionsStatusRequest instantiates a new HandlerUpdateIPOsTransactionsStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateIPOsTransactionsStatusRequestWithDefaults

`func NewHandlerUpdateIPOsTransactionsStatusRequestWithDefaults() *HandlerUpdateIPOsTransactionsStatusRequest`

NewHandlerUpdateIPOsTransactionsStatusRequestWithDefaults instantiates a new HandlerUpdateIPOsTransactionsStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedAmt

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetApprovedAmt() float32`

GetApprovedAmt returns the ApprovedAmt field if non-nil, zero value otherwise.

### GetApprovedAmtOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetApprovedAmtOk() (*float32, bool)`

GetApprovedAmtOk returns a tuple with the ApprovedAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmt

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetApprovedAmt(v float32)`

SetApprovedAmt sets ApprovedAmt field to given value.

### HasApprovedAmt

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasApprovedAmt() bool`

HasApprovedAmt returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetApprovedDetails() string`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetApprovedDetailsOk() (*string, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetApprovedDetails(v string)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsNotApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsNotApproved() bool`

GetIsNotApproved returns the IsNotApproved field if non-nil, zero value otherwise.

### GetIsNotApprovedOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsNotApprovedOk() (*bool, bool)`

GetIsNotApprovedOk returns a tuple with the IsNotApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetIsNotApproved(v bool)`

SetIsNotApproved sets IsNotApproved field to given value.

### HasIsNotApproved

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasIsNotApproved() bool`

HasIsNotApproved returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssUserId

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.


### GetRequestType

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetType

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetType() HandlerTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) GetTypeOk() (*HandlerTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlerUpdateIPOsTransactionsStatusRequest) SetType(v HandlerTransactionType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



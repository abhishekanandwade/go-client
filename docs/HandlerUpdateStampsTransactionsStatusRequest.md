# HandlerUpdateStampsTransactionsStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Details** | Pointer to **map[string]int32** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsNotApproved** | Pointer to **bool** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**ReqType** | Pointer to **string** |  | [optional] 
**UserId** | **int32** |  | 

## Methods

### NewHandlerUpdateStampsTransactionsStatusRequest

`func NewHandlerUpdateStampsTransactionsStatusRequest(userId int32, ) *HandlerUpdateStampsTransactionsStatusRequest`

NewHandlerUpdateStampsTransactionsStatusRequest instantiates a new HandlerUpdateStampsTransactionsStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateStampsTransactionsStatusRequestWithDefaults

`func NewHandlerUpdateStampsTransactionsStatusRequestWithDefaults() *HandlerUpdateStampsTransactionsStatusRequest`

NewHandlerUpdateStampsTransactionsStatusRequestWithDefaults instantiates a new HandlerUpdateStampsTransactionsStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetDetails

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetDetails() map[string]int32`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetDetailsOk() (*map[string]int32, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetDetails(v map[string]int32)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsNotApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsNotApproved() bool`

GetIsNotApproved returns the IsNotApproved field if non-nil, zero value otherwise.

### GetIsNotApprovedOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsNotApprovedOk() (*bool, bool)`

GetIsNotApprovedOk returns a tuple with the IsNotApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetIsNotApproved(v bool)`

SetIsNotApproved sets IsNotApproved field to given value.

### HasIsNotApproved

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasIsNotApproved() bool`

HasIsNotApproved returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetReqType

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetReqType() string`

GetReqType returns the ReqType field if non-nil, zero value otherwise.

### GetReqTypeOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetReqTypeOk() (*string, bool)`

GetReqTypeOk returns a tuple with the ReqType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqType

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetReqType(v string)`

SetReqType sets ReqType field to given value.

### HasReqType

`func (o *HandlerUpdateStampsTransactionsStatusRequest) HasReqType() bool`

HasReqType returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerUpdateStampsTransactionsStatusRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerUpdateStampsTransactionsStatusRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



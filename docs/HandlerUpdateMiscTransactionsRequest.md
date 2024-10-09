# HandlerUpdateMiscTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverId** | **int32** |  | 
**ApproverRemarks** | **string** |  | 
**IsApproved** | Pointer to **bool** |  | [optional] 
**TransactionID** | **string** |  | 

## Methods

### NewHandlerUpdateMiscTransactionsRequest

`func NewHandlerUpdateMiscTransactionsRequest(approverId int32, approverRemarks string, transactionID string, ) *HandlerUpdateMiscTransactionsRequest`

NewHandlerUpdateMiscTransactionsRequest instantiates a new HandlerUpdateMiscTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateMiscTransactionsRequestWithDefaults

`func NewHandlerUpdateMiscTransactionsRequestWithDefaults() *HandlerUpdateMiscTransactionsRequest`

NewHandlerUpdateMiscTransactionsRequestWithDefaults instantiates a new HandlerUpdateMiscTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverId

`func (o *HandlerUpdateMiscTransactionsRequest) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerUpdateMiscTransactionsRequest) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerUpdateMiscTransactionsRequest) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### GetApproverRemarks

`func (o *HandlerUpdateMiscTransactionsRequest) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *HandlerUpdateMiscTransactionsRequest) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *HandlerUpdateMiscTransactionsRequest) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.


### GetIsApproved

`func (o *HandlerUpdateMiscTransactionsRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateMiscTransactionsRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateMiscTransactionsRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateMiscTransactionsRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetTransactionID

`func (o *HandlerUpdateMiscTransactionsRequest) GetTransactionID() string`

GetTransactionID returns the TransactionID field if non-nil, zero value otherwise.

### GetTransactionIDOk

`func (o *HandlerUpdateMiscTransactionsRequest) GetTransactionIDOk() (*string, bool)`

GetTransactionIDOk returns a tuple with the TransactionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionID

`func (o *HandlerUpdateMiscTransactionsRequest) SetTransactionID(v string)`

SetTransactionID sets TransactionID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



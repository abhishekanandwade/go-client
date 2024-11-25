# HandlerUpdateBankAddlCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedAddlCredit** | **float32** |  | 
**ApproverId** | **string** |  | 
**ApproverRemarks** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlerUpdateBankAddlCreditRequest

`func NewHandlerUpdateBankAddlCreditRequest(approvedAddlCredit float32, approverId string, ) *HandlerUpdateBankAddlCreditRequest`

NewHandlerUpdateBankAddlCreditRequest instantiates a new HandlerUpdateBankAddlCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateBankAddlCreditRequestWithDefaults

`func NewHandlerUpdateBankAddlCreditRequestWithDefaults() *HandlerUpdateBankAddlCreditRequest`

NewHandlerUpdateBankAddlCreditRequestWithDefaults instantiates a new HandlerUpdateBankAddlCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedAddlCredit

`func (o *HandlerUpdateBankAddlCreditRequest) GetApprovedAddlCredit() float32`

GetApprovedAddlCredit returns the ApprovedAddlCredit field if non-nil, zero value otherwise.

### GetApprovedAddlCreditOk

`func (o *HandlerUpdateBankAddlCreditRequest) GetApprovedAddlCreditOk() (*float32, bool)`

GetApprovedAddlCreditOk returns a tuple with the ApprovedAddlCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAddlCredit

`func (o *HandlerUpdateBankAddlCreditRequest) SetApprovedAddlCredit(v float32)`

SetApprovedAddlCredit sets ApprovedAddlCredit field to given value.


### GetApproverId

`func (o *HandlerUpdateBankAddlCreditRequest) GetApproverId() string`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerUpdateBankAddlCreditRequest) GetApproverIdOk() (*string, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerUpdateBankAddlCreditRequest) SetApproverId(v string)`

SetApproverId sets ApproverId field to given value.


### GetApproverRemarks

`func (o *HandlerUpdateBankAddlCreditRequest) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *HandlerUpdateBankAddlCreditRequest) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *HandlerUpdateBankAddlCreditRequest) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.

### HasApproverRemarks

`func (o *HandlerUpdateBankAddlCreditRequest) HasApproverRemarks() bool`

HasApproverRemarks returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerUpdateBankAddlCreditRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateBankAddlCreditRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateBankAddlCreditRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateBankAddlCreditRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



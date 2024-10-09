# HandlerUpdateStampAdvanceTxnRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsApproved** | Pointer to **bool** |  | [optional] 
**TryApproverId** | **int32** |  | 

## Methods

### NewHandlerUpdateStampAdvanceTxnRequest

`func NewHandlerUpdateStampAdvanceTxnRequest(tryApproverId int32, ) *HandlerUpdateStampAdvanceTxnRequest`

NewHandlerUpdateStampAdvanceTxnRequest instantiates a new HandlerUpdateStampAdvanceTxnRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateStampAdvanceTxnRequestWithDefaults

`func NewHandlerUpdateStampAdvanceTxnRequestWithDefaults() *HandlerUpdateStampAdvanceTxnRequest`

NewHandlerUpdateStampAdvanceTxnRequestWithDefaults instantiates a new HandlerUpdateStampAdvanceTxnRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsApproved

`func (o *HandlerUpdateStampAdvanceTxnRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateStampAdvanceTxnRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateStampAdvanceTxnRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateStampAdvanceTxnRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetTryApproverId

`func (o *HandlerUpdateStampAdvanceTxnRequest) GetTryApproverId() int32`

GetTryApproverId returns the TryApproverId field if non-nil, zero value otherwise.

### GetTryApproverIdOk

`func (o *HandlerUpdateStampAdvanceTxnRequest) GetTryApproverIdOk() (*int32, bool)`

GetTryApproverIdOk returns a tuple with the TryApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryApproverId

`func (o *HandlerUpdateStampAdvanceTxnRequest) SetTryApproverId(v int32)`

SetTryApproverId sets TryApproverId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



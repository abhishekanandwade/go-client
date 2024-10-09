# HandlerUpdateChequesDisposalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeAmount** | **float32** |  | 
**DisposedBy** | **int32** |  | 
**IsBankDrawl** | Pointer to **bool** |  | [optional] 
**TransactionId** | **string** |  | 

## Methods

### NewHandlerUpdateChequesDisposalRequest

`func NewHandlerUpdateChequesDisposalRequest(chequeAmount float32, disposedBy int32, transactionId string, ) *HandlerUpdateChequesDisposalRequest`

NewHandlerUpdateChequesDisposalRequest instantiates a new HandlerUpdateChequesDisposalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateChequesDisposalRequestWithDefaults

`func NewHandlerUpdateChequesDisposalRequestWithDefaults() *HandlerUpdateChequesDisposalRequest`

NewHandlerUpdateChequesDisposalRequestWithDefaults instantiates a new HandlerUpdateChequesDisposalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeAmount

`func (o *HandlerUpdateChequesDisposalRequest) GetChequeAmount() float32`

GetChequeAmount returns the ChequeAmount field if non-nil, zero value otherwise.

### GetChequeAmountOk

`func (o *HandlerUpdateChequesDisposalRequest) GetChequeAmountOk() (*float32, bool)`

GetChequeAmountOk returns a tuple with the ChequeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeAmount

`func (o *HandlerUpdateChequesDisposalRequest) SetChequeAmount(v float32)`

SetChequeAmount sets ChequeAmount field to given value.


### GetDisposedBy

`func (o *HandlerUpdateChequesDisposalRequest) GetDisposedBy() int32`

GetDisposedBy returns the DisposedBy field if non-nil, zero value otherwise.

### GetDisposedByOk

`func (o *HandlerUpdateChequesDisposalRequest) GetDisposedByOk() (*int32, bool)`

GetDisposedByOk returns a tuple with the DisposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposedBy

`func (o *HandlerUpdateChequesDisposalRequest) SetDisposedBy(v int32)`

SetDisposedBy sets DisposedBy field to given value.


### GetIsBankDrawl

`func (o *HandlerUpdateChequesDisposalRequest) GetIsBankDrawl() bool`

GetIsBankDrawl returns the IsBankDrawl field if non-nil, zero value otherwise.

### GetIsBankDrawlOk

`func (o *HandlerUpdateChequesDisposalRequest) GetIsBankDrawlOk() (*bool, bool)`

GetIsBankDrawlOk returns a tuple with the IsBankDrawl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBankDrawl

`func (o *HandlerUpdateChequesDisposalRequest) SetIsBankDrawl(v bool)`

SetIsBankDrawl sets IsBankDrawl field to given value.

### HasIsBankDrawl

`func (o *HandlerUpdateChequesDisposalRequest) HasIsBankDrawl() bool`

HasIsBankDrawl returns a boolean if a field has been set.

### GetTransactionId

`func (o *HandlerUpdateChequesDisposalRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *HandlerUpdateChequesDisposalRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *HandlerUpdateChequesDisposalRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



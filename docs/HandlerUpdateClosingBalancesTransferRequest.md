# HandlerUpdateClosingBalancesTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverId** | **int32** |  | 
**CashBal** | Pointer to **float32** |  | [optional] 
**FromUserId** | **int32** |  | 
**Remarks** | **string** |  | 
**ToUserId** | **int32** |  | 
**Type** | [**HandlerStampStatus**](HandlerStampStatus.md) |  | 

## Methods

### NewHandlerUpdateClosingBalancesTransferRequest

`func NewHandlerUpdateClosingBalancesTransferRequest(approverId int32, fromUserId int32, remarks string, toUserId int32, type_ HandlerStampStatus, ) *HandlerUpdateClosingBalancesTransferRequest`

NewHandlerUpdateClosingBalancesTransferRequest instantiates a new HandlerUpdateClosingBalancesTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateClosingBalancesTransferRequestWithDefaults

`func NewHandlerUpdateClosingBalancesTransferRequestWithDefaults() *HandlerUpdateClosingBalancesTransferRequest`

NewHandlerUpdateClosingBalancesTransferRequestWithDefaults instantiates a new HandlerUpdateClosingBalancesTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverId

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### GetCashBal

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetCashBal() float32`

GetCashBal returns the CashBal field if non-nil, zero value otherwise.

### GetCashBalOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetCashBalOk() (*float32, bool)`

GetCashBalOk returns a tuple with the CashBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBal

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetCashBal(v float32)`

SetCashBal sets CashBal field to given value.

### HasCashBal

`func (o *HandlerUpdateClosingBalancesTransferRequest) HasCashBal() bool`

HasCashBal returns a boolean if a field has been set.

### GetFromUserId

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetFromUserId() int32`

GetFromUserId returns the FromUserId field if non-nil, zero value otherwise.

### GetFromUserIdOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetFromUserIdOk() (*int32, bool)`

GetFromUserIdOk returns a tuple with the FromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUserId

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetFromUserId(v int32)`

SetFromUserId sets FromUserId field to given value.


### GetRemarks

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.


### GetToUserId

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetToUserId() int32`

GetToUserId returns the ToUserId field if non-nil, zero value otherwise.

### GetToUserIdOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetToUserIdOk() (*int32, bool)`

GetToUserIdOk returns a tuple with the ToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUserId

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetToUserId(v int32)`

SetToUserId sets ToUserId field to given value.


### GetType

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetType() HandlerStampStatus`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlerUpdateClosingBalancesTransferRequest) GetTypeOk() (*HandlerStampStatus, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlerUpdateClosingBalancesTransferRequest) SetType(v HandlerStampStatus)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



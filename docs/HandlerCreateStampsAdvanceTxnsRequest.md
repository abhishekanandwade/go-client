# HandlerCreateStampsAdvanceTxnsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvAmt** | **float32** |  | 
**AdvTxnDetails** | **map[string]int32** |  | 
**AdvUserId** | **int32** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**TransType** | **string** |  | 
**TryUserId** | **int32** |  | 

## Methods

### NewHandlerCreateStampsAdvanceTxnsRequest

`func NewHandlerCreateStampsAdvanceTxnsRequest(advAmt float32, advTxnDetails map[string]int32, advUserId int32, transType string, tryUserId int32, ) *HandlerCreateStampsAdvanceTxnsRequest`

NewHandlerCreateStampsAdvanceTxnsRequest instantiates a new HandlerCreateStampsAdvanceTxnsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateStampsAdvanceTxnsRequestWithDefaults

`func NewHandlerCreateStampsAdvanceTxnsRequestWithDefaults() *HandlerCreateStampsAdvanceTxnsRequest`

NewHandlerCreateStampsAdvanceTxnsRequestWithDefaults instantiates a new HandlerCreateStampsAdvanceTxnsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvAmt

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvAmt() float32`

GetAdvAmt returns the AdvAmt field if non-nil, zero value otherwise.

### GetAdvAmtOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvAmtOk() (*float32, bool)`

GetAdvAmtOk returns a tuple with the AdvAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvAmt

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetAdvAmt(v float32)`

SetAdvAmt sets AdvAmt field to given value.


### GetAdvTxnDetails

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvTxnDetails() map[string]int32`

GetAdvTxnDetails returns the AdvTxnDetails field if non-nil, zero value otherwise.

### GetAdvTxnDetailsOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvTxnDetailsOk() (*map[string]int32, bool)`

GetAdvTxnDetailsOk returns a tuple with the AdvTxnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvTxnDetails

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetAdvTxnDetails(v map[string]int32)`

SetAdvTxnDetails sets AdvTxnDetails field to given value.


### GetAdvUserId

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvUserId() int32`

GetAdvUserId returns the AdvUserId field if non-nil, zero value otherwise.

### GetAdvUserIdOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetAdvUserIdOk() (*int32, bool)`

GetAdvUserIdOk returns a tuple with the AdvUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvUserId

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetAdvUserId(v int32)`

SetAdvUserId sets AdvUserId field to given value.


### GetRemarks

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateStampsAdvanceTxnsRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTransType

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetTransType(v string)`

SetTransType sets TransType field to given value.


### GetTryUserId

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetTryUserId() int32`

GetTryUserId returns the TryUserId field if non-nil, zero value otherwise.

### GetTryUserIdOk

`func (o *HandlerCreateStampsAdvanceTxnsRequest) GetTryUserIdOk() (*int32, bool)`

GetTryUserIdOk returns a tuple with the TryUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryUserId

`func (o *HandlerCreateStampsAdvanceTxnsRequest) SetTryUserId(v int32)`

SetTryUserId sets TryUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



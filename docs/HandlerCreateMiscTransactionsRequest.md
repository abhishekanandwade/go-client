# HandlerCreateMiscTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**OfficeId** | **int32** |  | 
**PayeeName** | Pointer to **string** |  | [optional] 
**Remarks** | **string** |  | 
**TransAmount** | **float32** |  | 
**TransDetails** | **map[string]float32** |  | 
**TransType** | **string** |  | 
**TxnMode** | **string** |  | 
**UserId** | **int32** |  | 

## Methods

### NewHandlerCreateMiscTransactionsRequest

`func NewHandlerCreateMiscTransactionsRequest(officeId int32, remarks string, transAmount float32, transDetails map[string]float32, transType string, txnMode string, userId int32, ) *HandlerCreateMiscTransactionsRequest`

NewHandlerCreateMiscTransactionsRequest instantiates a new HandlerCreateMiscTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateMiscTransactionsRequestWithDefaults

`func NewHandlerCreateMiscTransactionsRequestWithDefaults() *HandlerCreateMiscTransactionsRequest`

NewHandlerCreateMiscTransactionsRequestWithDefaults instantiates a new HandlerCreateMiscTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeDate

`func (o *HandlerCreateMiscTransactionsRequest) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *HandlerCreateMiscTransactionsRequest) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *HandlerCreateMiscTransactionsRequest) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *HandlerCreateMiscTransactionsRequest) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *HandlerCreateMiscTransactionsRequest) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *HandlerCreateMiscTransactionsRequest) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *HandlerCreateMiscTransactionsRequest) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *HandlerCreateMiscTransactionsRequest) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerCreateMiscTransactionsRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateMiscTransactionsRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateMiscTransactionsRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetPayeeName

`func (o *HandlerCreateMiscTransactionsRequest) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *HandlerCreateMiscTransactionsRequest) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *HandlerCreateMiscTransactionsRequest) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *HandlerCreateMiscTransactionsRequest) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerCreateMiscTransactionsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateMiscTransactionsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateMiscTransactionsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.


### GetTransAmount

`func (o *HandlerCreateMiscTransactionsRequest) GetTransAmount() float32`

GetTransAmount returns the TransAmount field if non-nil, zero value otherwise.

### GetTransAmountOk

`func (o *HandlerCreateMiscTransactionsRequest) GetTransAmountOk() (*float32, bool)`

GetTransAmountOk returns a tuple with the TransAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransAmount

`func (o *HandlerCreateMiscTransactionsRequest) SetTransAmount(v float32)`

SetTransAmount sets TransAmount field to given value.


### GetTransDetails

`func (o *HandlerCreateMiscTransactionsRequest) GetTransDetails() map[string]float32`

GetTransDetails returns the TransDetails field if non-nil, zero value otherwise.

### GetTransDetailsOk

`func (o *HandlerCreateMiscTransactionsRequest) GetTransDetailsOk() (*map[string]float32, bool)`

GetTransDetailsOk returns a tuple with the TransDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDetails

`func (o *HandlerCreateMiscTransactionsRequest) SetTransDetails(v map[string]float32)`

SetTransDetails sets TransDetails field to given value.


### GetTransType

`func (o *HandlerCreateMiscTransactionsRequest) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *HandlerCreateMiscTransactionsRequest) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *HandlerCreateMiscTransactionsRequest) SetTransType(v string)`

SetTransType sets TransType field to given value.


### GetTxnMode

`func (o *HandlerCreateMiscTransactionsRequest) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *HandlerCreateMiscTransactionsRequest) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *HandlerCreateMiscTransactionsRequest) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.


### GetUserId

`func (o *HandlerCreateMiscTransactionsRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerCreateMiscTransactionsRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerCreateMiscTransactionsRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerUpdateTransactionsWalletBalanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**OfficeId** | **int32** |  | 
**ReceiptOrPayment** | Pointer to **string** |  | [optional] 
**UserId** | **int32** |  | 

## Methods

### NewHandlerUpdateTransactionsWalletBalanceRequest

`func NewHandlerUpdateTransactionsWalletBalanceRequest(amount float32, officeId int32, userId int32, ) *HandlerUpdateTransactionsWalletBalanceRequest`

NewHandlerUpdateTransactionsWalletBalanceRequest instantiates a new HandlerUpdateTransactionsWalletBalanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateTransactionsWalletBalanceRequestWithDefaults

`func NewHandlerUpdateTransactionsWalletBalanceRequestWithDefaults() *HandlerUpdateTransactionsWalletBalanceRequest`

NewHandlerUpdateTransactionsWalletBalanceRequestWithDefaults instantiates a new HandlerUpdateTransactionsWalletBalanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetOfficeId

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetReceiptOrPayment

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetReceiptOrPayment() string`

GetReceiptOrPayment returns the ReceiptOrPayment field if non-nil, zero value otherwise.

### GetReceiptOrPaymentOk

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetReceiptOrPaymentOk() (*string, bool)`

GetReceiptOrPaymentOk returns a tuple with the ReceiptOrPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOrPayment

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) SetReceiptOrPayment(v string)`

SetReceiptOrPayment sets ReceiptOrPayment field to given value.

### HasReceiptOrPayment

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) HasReceiptOrPayment() bool`

HasReceiptOrPayment returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerUpdateTransactionsWalletBalanceRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseChequeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeAmount** | Pointer to **float32** |  | [optional] 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**DisposedBy** | Pointer to **string** |  | [optional] 
**DisposedDate** | Pointer to **string** |  | [optional] 
**IsBankDrawl** | Pointer to **bool** |  | [optional] 
**PayeeName** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseChequeRequest

`func NewResponseChequeRequest() *ResponseChequeRequest`

NewResponseChequeRequest instantiates a new ResponseChequeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseChequeRequestWithDefaults

`func NewResponseChequeRequestWithDefaults() *ResponseChequeRequest`

NewResponseChequeRequestWithDefaults instantiates a new ResponseChequeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeAmount

`func (o *ResponseChequeRequest) GetChequeAmount() float32`

GetChequeAmount returns the ChequeAmount field if non-nil, zero value otherwise.

### GetChequeAmountOk

`func (o *ResponseChequeRequest) GetChequeAmountOk() (*float32, bool)`

GetChequeAmountOk returns a tuple with the ChequeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeAmount

`func (o *ResponseChequeRequest) SetChequeAmount(v float32)`

SetChequeAmount sets ChequeAmount field to given value.

### HasChequeAmount

`func (o *ResponseChequeRequest) HasChequeAmount() bool`

HasChequeAmount returns a boolean if a field has been set.

### GetChequeDate

`func (o *ResponseChequeRequest) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *ResponseChequeRequest) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *ResponseChequeRequest) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *ResponseChequeRequest) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *ResponseChequeRequest) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *ResponseChequeRequest) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *ResponseChequeRequest) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *ResponseChequeRequest) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetDisposedBy

`func (o *ResponseChequeRequest) GetDisposedBy() string`

GetDisposedBy returns the DisposedBy field if non-nil, zero value otherwise.

### GetDisposedByOk

`func (o *ResponseChequeRequest) GetDisposedByOk() (*string, bool)`

GetDisposedByOk returns a tuple with the DisposedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposedBy

`func (o *ResponseChequeRequest) SetDisposedBy(v string)`

SetDisposedBy sets DisposedBy field to given value.

### HasDisposedBy

`func (o *ResponseChequeRequest) HasDisposedBy() bool`

HasDisposedBy returns a boolean if a field has been set.

### GetDisposedDate

`func (o *ResponseChequeRequest) GetDisposedDate() string`

GetDisposedDate returns the DisposedDate field if non-nil, zero value otherwise.

### GetDisposedDateOk

`func (o *ResponseChequeRequest) GetDisposedDateOk() (*string, bool)`

GetDisposedDateOk returns a tuple with the DisposedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposedDate

`func (o *ResponseChequeRequest) SetDisposedDate(v string)`

SetDisposedDate sets DisposedDate field to given value.

### HasDisposedDate

`func (o *ResponseChequeRequest) HasDisposedDate() bool`

HasDisposedDate returns a boolean if a field has been set.

### GetIsBankDrawl

`func (o *ResponseChequeRequest) GetIsBankDrawl() bool`

GetIsBankDrawl returns the IsBankDrawl field if non-nil, zero value otherwise.

### GetIsBankDrawlOk

`func (o *ResponseChequeRequest) GetIsBankDrawlOk() (*bool, bool)`

GetIsBankDrawlOk returns a tuple with the IsBankDrawl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBankDrawl

`func (o *ResponseChequeRequest) SetIsBankDrawl(v bool)`

SetIsBankDrawl sets IsBankDrawl field to given value.

### HasIsBankDrawl

`func (o *ResponseChequeRequest) HasIsBankDrawl() bool`

HasIsBankDrawl returns a boolean if a field has been set.

### GetPayeeName

`func (o *ResponseChequeRequest) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *ResponseChequeRequest) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *ResponseChequeRequest) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *ResponseChequeRequest) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseChequeRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseChequeRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseChequeRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseChequeRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseChequeRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseChequeRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseChequeRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseChequeRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseChequeRequest) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseChequeRequest) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseChequeRequest) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseChequeRequest) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ResponseChequeRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ResponseChequeRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ResponseChequeRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ResponseChequeRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseChequeRequest) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseChequeRequest) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseChequeRequest) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseChequeRequest) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseCashInTransit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountSent** | Pointer to **float32** |  | [optional] 
**FromOffice** | Pointer to **string** |  | [optional] 
**FromUnitId** | Pointer to **int32** |  | [optional] 
**IsRsao** | Pointer to **bool** |  | [optional] 
**ToOffice** | Pointer to **string** |  | [optional] 
**ToUnitId** | Pointer to **int32** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseCashInTransit

`func NewResponseCashInTransit() *ResponseCashInTransit`

NewResponseCashInTransit instantiates a new ResponseCashInTransit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCashInTransitWithDefaults

`func NewResponseCashInTransitWithDefaults() *ResponseCashInTransit`

NewResponseCashInTransitWithDefaults instantiates a new ResponseCashInTransit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountSent

`func (o *ResponseCashInTransit) GetAmountSent() float32`

GetAmountSent returns the AmountSent field if non-nil, zero value otherwise.

### GetAmountSentOk

`func (o *ResponseCashInTransit) GetAmountSentOk() (*float32, bool)`

GetAmountSentOk returns a tuple with the AmountSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSent

`func (o *ResponseCashInTransit) SetAmountSent(v float32)`

SetAmountSent sets AmountSent field to given value.

### HasAmountSent

`func (o *ResponseCashInTransit) HasAmountSent() bool`

HasAmountSent returns a boolean if a field has been set.

### GetFromOffice

`func (o *ResponseCashInTransit) GetFromOffice() string`

GetFromOffice returns the FromOffice field if non-nil, zero value otherwise.

### GetFromOfficeOk

`func (o *ResponseCashInTransit) GetFromOfficeOk() (*string, bool)`

GetFromOfficeOk returns a tuple with the FromOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOffice

`func (o *ResponseCashInTransit) SetFromOffice(v string)`

SetFromOffice sets FromOffice field to given value.

### HasFromOffice

`func (o *ResponseCashInTransit) HasFromOffice() bool`

HasFromOffice returns a boolean if a field has been set.

### GetFromUnitId

`func (o *ResponseCashInTransit) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *ResponseCashInTransit) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *ResponseCashInTransit) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *ResponseCashInTransit) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetIsRsao

`func (o *ResponseCashInTransit) GetIsRsao() bool`

GetIsRsao returns the IsRsao field if non-nil, zero value otherwise.

### GetIsRsaoOk

`func (o *ResponseCashInTransit) GetIsRsaoOk() (*bool, bool)`

GetIsRsaoOk returns a tuple with the IsRsao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRsao

`func (o *ResponseCashInTransit) SetIsRsao(v bool)`

SetIsRsao sets IsRsao field to given value.

### HasIsRsao

`func (o *ResponseCashInTransit) HasIsRsao() bool`

HasIsRsao returns a boolean if a field has been set.

### GetToOffice

`func (o *ResponseCashInTransit) GetToOffice() string`

GetToOffice returns the ToOffice field if non-nil, zero value otherwise.

### GetToOfficeOk

`func (o *ResponseCashInTransit) GetToOfficeOk() (*string, bool)`

GetToOfficeOk returns a tuple with the ToOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOffice

`func (o *ResponseCashInTransit) SetToOffice(v string)`

SetToOffice sets ToOffice field to given value.

### HasToOffice

`func (o *ResponseCashInTransit) HasToOffice() bool`

HasToOffice returns a boolean if a field has been set.

### GetToUnitId

`func (o *ResponseCashInTransit) GetToUnitId() int32`

GetToUnitId returns the ToUnitId field if non-nil, zero value otherwise.

### GetToUnitIdOk

`func (o *ResponseCashInTransit) GetToUnitIdOk() (*int32, bool)`

GetToUnitIdOk returns a tuple with the ToUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnitId

`func (o *ResponseCashInTransit) SetToUnitId(v int32)`

SetToUnitId sets ToUnitId field to given value.

### HasToUnitId

`func (o *ResponseCashInTransit) HasToUnitId() bool`

HasToUnitId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseCashInTransit) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseCashInTransit) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseCashInTransit) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseCashInTransit) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseCashInTransit) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseCashInTransit) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseCashInTransit) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseCashInTransit) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ResponseCashInTransit) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ResponseCashInTransit) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ResponseCashInTransit) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ResponseCashInTransit) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



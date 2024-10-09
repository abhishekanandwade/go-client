# ResponseChequesInTransit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeAmount** | Pointer to **float32** |  | [optional] 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**FromOffice** | Pointer to **string** |  | [optional] 
**FromUnitId** | Pointer to **int32** |  | [optional] 
**IsRsao** | Pointer to **bool** |  | [optional] 
**ToOffice** | Pointer to **string** |  | [optional] 
**ToUnitId** | Pointer to **int32** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseChequesInTransit

`func NewResponseChequesInTransit() *ResponseChequesInTransit`

NewResponseChequesInTransit instantiates a new ResponseChequesInTransit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseChequesInTransitWithDefaults

`func NewResponseChequesInTransitWithDefaults() *ResponseChequesInTransit`

NewResponseChequesInTransitWithDefaults instantiates a new ResponseChequesInTransit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeAmount

`func (o *ResponseChequesInTransit) GetChequeAmount() float32`

GetChequeAmount returns the ChequeAmount field if non-nil, zero value otherwise.

### GetChequeAmountOk

`func (o *ResponseChequesInTransit) GetChequeAmountOk() (*float32, bool)`

GetChequeAmountOk returns a tuple with the ChequeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeAmount

`func (o *ResponseChequesInTransit) SetChequeAmount(v float32)`

SetChequeAmount sets ChequeAmount field to given value.

### HasChequeAmount

`func (o *ResponseChequesInTransit) HasChequeAmount() bool`

HasChequeAmount returns a boolean if a field has been set.

### GetChequeDate

`func (o *ResponseChequesInTransit) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *ResponseChequesInTransit) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *ResponseChequesInTransit) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *ResponseChequesInTransit) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *ResponseChequesInTransit) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *ResponseChequesInTransit) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *ResponseChequesInTransit) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *ResponseChequesInTransit) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetFromOffice

`func (o *ResponseChequesInTransit) GetFromOffice() string`

GetFromOffice returns the FromOffice field if non-nil, zero value otherwise.

### GetFromOfficeOk

`func (o *ResponseChequesInTransit) GetFromOfficeOk() (*string, bool)`

GetFromOfficeOk returns a tuple with the FromOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOffice

`func (o *ResponseChequesInTransit) SetFromOffice(v string)`

SetFromOffice sets FromOffice field to given value.

### HasFromOffice

`func (o *ResponseChequesInTransit) HasFromOffice() bool`

HasFromOffice returns a boolean if a field has been set.

### GetFromUnitId

`func (o *ResponseChequesInTransit) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *ResponseChequesInTransit) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *ResponseChequesInTransit) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *ResponseChequesInTransit) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetIsRsao

`func (o *ResponseChequesInTransit) GetIsRsao() bool`

GetIsRsao returns the IsRsao field if non-nil, zero value otherwise.

### GetIsRsaoOk

`func (o *ResponseChequesInTransit) GetIsRsaoOk() (*bool, bool)`

GetIsRsaoOk returns a tuple with the IsRsao field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRsao

`func (o *ResponseChequesInTransit) SetIsRsao(v bool)`

SetIsRsao sets IsRsao field to given value.

### HasIsRsao

`func (o *ResponseChequesInTransit) HasIsRsao() bool`

HasIsRsao returns a boolean if a field has been set.

### GetToOffice

`func (o *ResponseChequesInTransit) GetToOffice() string`

GetToOffice returns the ToOffice field if non-nil, zero value otherwise.

### GetToOfficeOk

`func (o *ResponseChequesInTransit) GetToOfficeOk() (*string, bool)`

GetToOfficeOk returns a tuple with the ToOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOffice

`func (o *ResponseChequesInTransit) SetToOffice(v string)`

SetToOffice sets ToOffice field to given value.

### HasToOffice

`func (o *ResponseChequesInTransit) HasToOffice() bool`

HasToOffice returns a boolean if a field has been set.

### GetToUnitId

`func (o *ResponseChequesInTransit) GetToUnitId() int32`

GetToUnitId returns the ToUnitId field if non-nil, zero value otherwise.

### GetToUnitIdOk

`func (o *ResponseChequesInTransit) GetToUnitIdOk() (*int32, bool)`

GetToUnitIdOk returns a tuple with the ToUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnitId

`func (o *ResponseChequesInTransit) SetToUnitId(v int32)`

SetToUnitId sets ToUnitId field to given value.

### HasToUnitId

`func (o *ResponseChequesInTransit) HasToUnitId() bool`

HasToUnitId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseChequesInTransit) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseChequesInTransit) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseChequesInTransit) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseChequesInTransit) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseChequesInTransit) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseChequesInTransit) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseChequesInTransit) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseChequesInTransit) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ResponseChequesInTransit) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ResponseChequesInTransit) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ResponseChequesInTransit) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ResponseChequesInTransit) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerCreateCurrenciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyType** | Pointer to **string** |  | [optional] 
**Denomination** | **float32** |  | 
**EnteredByUser** | **string** |  | 
**ValidFrom** | **string** |  | 

## Methods

### NewHandlerCreateCurrenciesRequest

`func NewHandlerCreateCurrenciesRequest(denomination float32, enteredByUser string, validFrom string, ) *HandlerCreateCurrenciesRequest`

NewHandlerCreateCurrenciesRequest instantiates a new HandlerCreateCurrenciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateCurrenciesRequestWithDefaults

`func NewHandlerCreateCurrenciesRequestWithDefaults() *HandlerCreateCurrenciesRequest`

NewHandlerCreateCurrenciesRequestWithDefaults instantiates a new HandlerCreateCurrenciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyType

`func (o *HandlerCreateCurrenciesRequest) GetCurrencyType() string`

GetCurrencyType returns the CurrencyType field if non-nil, zero value otherwise.

### GetCurrencyTypeOk

`func (o *HandlerCreateCurrenciesRequest) GetCurrencyTypeOk() (*string, bool)`

GetCurrencyTypeOk returns a tuple with the CurrencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyType

`func (o *HandlerCreateCurrenciesRequest) SetCurrencyType(v string)`

SetCurrencyType sets CurrencyType field to given value.

### HasCurrencyType

`func (o *HandlerCreateCurrenciesRequest) HasCurrencyType() bool`

HasCurrencyType returns a boolean if a field has been set.

### GetDenomination

`func (o *HandlerCreateCurrenciesRequest) GetDenomination() float32`

GetDenomination returns the Denomination field if non-nil, zero value otherwise.

### GetDenominationOk

`func (o *HandlerCreateCurrenciesRequest) GetDenominationOk() (*float32, bool)`

GetDenominationOk returns a tuple with the Denomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenomination

`func (o *HandlerCreateCurrenciesRequest) SetDenomination(v float32)`

SetDenomination sets Denomination field to given value.


### GetEnteredByUser

`func (o *HandlerCreateCurrenciesRequest) GetEnteredByUser() string`

GetEnteredByUser returns the EnteredByUser field if non-nil, zero value otherwise.

### GetEnteredByUserOk

`func (o *HandlerCreateCurrenciesRequest) GetEnteredByUserOk() (*string, bool)`

GetEnteredByUserOk returns a tuple with the EnteredByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredByUser

`func (o *HandlerCreateCurrenciesRequest) SetEnteredByUser(v string)`

SetEnteredByUser sets EnteredByUser field to given value.


### GetValidFrom

`func (o *HandlerCreateCurrenciesRequest) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *HandlerCreateCurrenciesRequest) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *HandlerCreateCurrenciesRequest) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



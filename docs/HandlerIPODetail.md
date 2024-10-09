# HandlerIPODetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrQty** | Pointer to **int32** |  | [optional] 
**CurrSerial** | Pointer to **string** |  | [optional] 
**DenominationDesc** | Pointer to **string** |  | [optional] 
**DenominationValue** | Pointer to **string** |  | [optional] 
**DenominationId** | **string** |  | 
**EndSerial** | **string** |  | 
**PrefixId** | **string** |  | 
**StartSerial** | **string** |  | 
**TotQuantity** | **int32** |  | 

## Methods

### NewHandlerIPODetail

`func NewHandlerIPODetail(denominationId string, endSerial string, prefixId string, startSerial string, totQuantity int32, ) *HandlerIPODetail`

NewHandlerIPODetail instantiates a new HandlerIPODetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerIPODetailWithDefaults

`func NewHandlerIPODetailWithDefaults() *HandlerIPODetail`

NewHandlerIPODetailWithDefaults instantiates a new HandlerIPODetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrQty

`func (o *HandlerIPODetail) GetCurrQty() int32`

GetCurrQty returns the CurrQty field if non-nil, zero value otherwise.

### GetCurrQtyOk

`func (o *HandlerIPODetail) GetCurrQtyOk() (*int32, bool)`

GetCurrQtyOk returns a tuple with the CurrQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrQty

`func (o *HandlerIPODetail) SetCurrQty(v int32)`

SetCurrQty sets CurrQty field to given value.

### HasCurrQty

`func (o *HandlerIPODetail) HasCurrQty() bool`

HasCurrQty returns a boolean if a field has been set.

### GetCurrSerial

`func (o *HandlerIPODetail) GetCurrSerial() string`

GetCurrSerial returns the CurrSerial field if non-nil, zero value otherwise.

### GetCurrSerialOk

`func (o *HandlerIPODetail) GetCurrSerialOk() (*string, bool)`

GetCurrSerialOk returns a tuple with the CurrSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrSerial

`func (o *HandlerIPODetail) SetCurrSerial(v string)`

SetCurrSerial sets CurrSerial field to given value.

### HasCurrSerial

`func (o *HandlerIPODetail) HasCurrSerial() bool`

HasCurrSerial returns a boolean if a field has been set.

### GetDenominationDesc

`func (o *HandlerIPODetail) GetDenominationDesc() string`

GetDenominationDesc returns the DenominationDesc field if non-nil, zero value otherwise.

### GetDenominationDescOk

`func (o *HandlerIPODetail) GetDenominationDescOk() (*string, bool)`

GetDenominationDescOk returns a tuple with the DenominationDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationDesc

`func (o *HandlerIPODetail) SetDenominationDesc(v string)`

SetDenominationDesc sets DenominationDesc field to given value.

### HasDenominationDesc

`func (o *HandlerIPODetail) HasDenominationDesc() bool`

HasDenominationDesc returns a boolean if a field has been set.

### GetDenominationValue

`func (o *HandlerIPODetail) GetDenominationValue() string`

GetDenominationValue returns the DenominationValue field if non-nil, zero value otherwise.

### GetDenominationValueOk

`func (o *HandlerIPODetail) GetDenominationValueOk() (*string, bool)`

GetDenominationValueOk returns a tuple with the DenominationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationValue

`func (o *HandlerIPODetail) SetDenominationValue(v string)`

SetDenominationValue sets DenominationValue field to given value.

### HasDenominationValue

`func (o *HandlerIPODetail) HasDenominationValue() bool`

HasDenominationValue returns a boolean if a field has been set.

### GetDenominationId

`func (o *HandlerIPODetail) GetDenominationId() string`

GetDenominationId returns the DenominationId field if non-nil, zero value otherwise.

### GetDenominationIdOk

`func (o *HandlerIPODetail) GetDenominationIdOk() (*string, bool)`

GetDenominationIdOk returns a tuple with the DenominationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationId

`func (o *HandlerIPODetail) SetDenominationId(v string)`

SetDenominationId sets DenominationId field to given value.


### GetEndSerial

`func (o *HandlerIPODetail) GetEndSerial() string`

GetEndSerial returns the EndSerial field if non-nil, zero value otherwise.

### GetEndSerialOk

`func (o *HandlerIPODetail) GetEndSerialOk() (*string, bool)`

GetEndSerialOk returns a tuple with the EndSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndSerial

`func (o *HandlerIPODetail) SetEndSerial(v string)`

SetEndSerial sets EndSerial field to given value.


### GetPrefixId

`func (o *HandlerIPODetail) GetPrefixId() string`

GetPrefixId returns the PrefixId field if non-nil, zero value otherwise.

### GetPrefixIdOk

`func (o *HandlerIPODetail) GetPrefixIdOk() (*string, bool)`

GetPrefixIdOk returns a tuple with the PrefixId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixId

`func (o *HandlerIPODetail) SetPrefixId(v string)`

SetPrefixId sets PrefixId field to given value.


### GetStartSerial

`func (o *HandlerIPODetail) GetStartSerial() string`

GetStartSerial returns the StartSerial field if non-nil, zero value otherwise.

### GetStartSerialOk

`func (o *HandlerIPODetail) GetStartSerialOk() (*string, bool)`

GetStartSerialOk returns a tuple with the StartSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSerial

`func (o *HandlerIPODetail) SetStartSerial(v string)`

SetStartSerial sets StartSerial field to given value.


### GetTotQuantity

`func (o *HandlerIPODetail) GetTotQuantity() int32`

GetTotQuantity returns the TotQuantity field if non-nil, zero value otherwise.

### GetTotQuantityOk

`func (o *HandlerIPODetail) GetTotQuantityOk() (*int32, bool)`

GetTotQuantityOk returns a tuple with the TotQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotQuantity

`func (o *HandlerIPODetail) SetTotQuantity(v int32)`

SetTotQuantity sets TotQuantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseStampsInTransit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountSent** | Pointer to **float32** |  | [optional] 
**FromOffice** | Pointer to **string** |  | [optional] 
**FromUnitId** | Pointer to **int32** |  | [optional] 
**SentDetails** | Pointer to **map[string]int32** |  | [optional] 
**StampDetails** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**ToOffice** | Pointer to **string** |  | [optional] 
**ToUnitId** | Pointer to **int32** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransactionDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseStampsInTransit

`func NewResponseStampsInTransit() *ResponseStampsInTransit`

NewResponseStampsInTransit instantiates a new ResponseStampsInTransit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsInTransitWithDefaults

`func NewResponseStampsInTransitWithDefaults() *ResponseStampsInTransit`

NewResponseStampsInTransitWithDefaults instantiates a new ResponseStampsInTransit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountSent

`func (o *ResponseStampsInTransit) GetAmountSent() float32`

GetAmountSent returns the AmountSent field if non-nil, zero value otherwise.

### GetAmountSentOk

`func (o *ResponseStampsInTransit) GetAmountSentOk() (*float32, bool)`

GetAmountSentOk returns a tuple with the AmountSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSent

`func (o *ResponseStampsInTransit) SetAmountSent(v float32)`

SetAmountSent sets AmountSent field to given value.

### HasAmountSent

`func (o *ResponseStampsInTransit) HasAmountSent() bool`

HasAmountSent returns a boolean if a field has been set.

### GetFromOffice

`func (o *ResponseStampsInTransit) GetFromOffice() string`

GetFromOffice returns the FromOffice field if non-nil, zero value otherwise.

### GetFromOfficeOk

`func (o *ResponseStampsInTransit) GetFromOfficeOk() (*string, bool)`

GetFromOfficeOk returns a tuple with the FromOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOffice

`func (o *ResponseStampsInTransit) SetFromOffice(v string)`

SetFromOffice sets FromOffice field to given value.

### HasFromOffice

`func (o *ResponseStampsInTransit) HasFromOffice() bool`

HasFromOffice returns a boolean if a field has been set.

### GetFromUnitId

`func (o *ResponseStampsInTransit) GetFromUnitId() int32`

GetFromUnitId returns the FromUnitId field if non-nil, zero value otherwise.

### GetFromUnitIdOk

`func (o *ResponseStampsInTransit) GetFromUnitIdOk() (*int32, bool)`

GetFromUnitIdOk returns a tuple with the FromUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromUnitId

`func (o *ResponseStampsInTransit) SetFromUnitId(v int32)`

SetFromUnitId sets FromUnitId field to given value.

### HasFromUnitId

`func (o *ResponseStampsInTransit) HasFromUnitId() bool`

HasFromUnitId returns a boolean if a field has been set.

### GetSentDetails

`func (o *ResponseStampsInTransit) GetSentDetails() map[string]int32`

GetSentDetails returns the SentDetails field if non-nil, zero value otherwise.

### GetSentDetailsOk

`func (o *ResponseStampsInTransit) GetSentDetailsOk() (*map[string]int32, bool)`

GetSentDetailsOk returns a tuple with the SentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentDetails

`func (o *ResponseStampsInTransit) SetSentDetails(v map[string]int32)`

SetSentDetails sets SentDetails field to given value.

### HasSentDetails

`func (o *ResponseStampsInTransit) HasSentDetails() bool`

HasSentDetails returns a boolean if a field has been set.

### GetStampDetails

`func (o *ResponseStampsInTransit) GetStampDetails() []ResponseStampdetails`

GetStampDetails returns the StampDetails field if non-nil, zero value otherwise.

### GetStampDetailsOk

`func (o *ResponseStampsInTransit) GetStampDetailsOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsOk returns a tuple with the StampDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetails

`func (o *ResponseStampsInTransit) SetStampDetails(v []ResponseStampdetails)`

SetStampDetails sets StampDetails field to given value.

### HasStampDetails

`func (o *ResponseStampsInTransit) HasStampDetails() bool`

HasStampDetails returns a boolean if a field has been set.

### GetToOffice

`func (o *ResponseStampsInTransit) GetToOffice() string`

GetToOffice returns the ToOffice field if non-nil, zero value otherwise.

### GetToOfficeOk

`func (o *ResponseStampsInTransit) GetToOfficeOk() (*string, bool)`

GetToOfficeOk returns a tuple with the ToOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOffice

`func (o *ResponseStampsInTransit) SetToOffice(v string)`

SetToOffice sets ToOffice field to given value.

### HasToOffice

`func (o *ResponseStampsInTransit) HasToOffice() bool`

HasToOffice returns a boolean if a field has been set.

### GetToUnitId

`func (o *ResponseStampsInTransit) GetToUnitId() int32`

GetToUnitId returns the ToUnitId field if non-nil, zero value otherwise.

### GetToUnitIdOk

`func (o *ResponseStampsInTransit) GetToUnitIdOk() (*int32, bool)`

GetToUnitIdOk returns a tuple with the ToUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToUnitId

`func (o *ResponseStampsInTransit) SetToUnitId(v int32)`

SetToUnitId sets ToUnitId field to given value.

### HasToUnitId

`func (o *ResponseStampsInTransit) HasToUnitId() bool`

HasToUnitId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsInTransit) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsInTransit) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsInTransit) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsInTransit) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseStampsInTransit) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseStampsInTransit) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseStampsInTransit) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseStampsInTransit) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *ResponseStampsInTransit) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *ResponseStampsInTransit) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *ResponseStampsInTransit) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *ResponseStampsInTransit) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerCreateOfficesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | Pointer to **string** |  | [optional] 
**CashOfficeId** | Pointer to **int32** |  | [optional] 
**CashofficeName** | Pointer to **string** |  | [optional] 
**CounterCashTfrLimit** | Pointer to **float32** |  | [optional] 
**EnteredBy** | **string** |  | 
**MaxAmt** | **float32** |  | 
**MinAmt** | **float32** |  | 
**OfficeId** | **int32** |  | 
**OfficeName** | **string** |  | 
**PostageStampsLimit** | Pointer to **float32** |  | [optional] 
**RevenueStampsLimit** | Pointer to **float32** |  | [optional] 
**TransitDuration** | Pointer to **int32** |  | [optional] 
**ValidFrom** | **string** |  | 

## Methods

### NewHandlerCreateOfficesRequest

`func NewHandlerCreateOfficesRequest(enteredBy string, maxAmt float32, minAmt float32, officeId int32, officeName string, validFrom string, ) *HandlerCreateOfficesRequest`

NewHandlerCreateOfficesRequest instantiates a new HandlerCreateOfficesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateOfficesRequestWithDefaults

`func NewHandlerCreateOfficesRequestWithDefaults() *HandlerCreateOfficesRequest`

NewHandlerCreateOfficesRequestWithDefaults instantiates a new HandlerCreateOfficesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *HandlerCreateOfficesRequest) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *HandlerCreateOfficesRequest) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *HandlerCreateOfficesRequest) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *HandlerCreateOfficesRequest) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetCashOfficeId

`func (o *HandlerCreateOfficesRequest) GetCashOfficeId() int32`

GetCashOfficeId returns the CashOfficeId field if non-nil, zero value otherwise.

### GetCashOfficeIdOk

`func (o *HandlerCreateOfficesRequest) GetCashOfficeIdOk() (*int32, bool)`

GetCashOfficeIdOk returns a tuple with the CashOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOfficeId

`func (o *HandlerCreateOfficesRequest) SetCashOfficeId(v int32)`

SetCashOfficeId sets CashOfficeId field to given value.

### HasCashOfficeId

`func (o *HandlerCreateOfficesRequest) HasCashOfficeId() bool`

HasCashOfficeId returns a boolean if a field has been set.

### GetCashofficeName

`func (o *HandlerCreateOfficesRequest) GetCashofficeName() string`

GetCashofficeName returns the CashofficeName field if non-nil, zero value otherwise.

### GetCashofficeNameOk

`func (o *HandlerCreateOfficesRequest) GetCashofficeNameOk() (*string, bool)`

GetCashofficeNameOk returns a tuple with the CashofficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashofficeName

`func (o *HandlerCreateOfficesRequest) SetCashofficeName(v string)`

SetCashofficeName sets CashofficeName field to given value.

### HasCashofficeName

`func (o *HandlerCreateOfficesRequest) HasCashofficeName() bool`

HasCashofficeName returns a boolean if a field has been set.

### GetCounterCashTfrLimit

`func (o *HandlerCreateOfficesRequest) GetCounterCashTfrLimit() float32`

GetCounterCashTfrLimit returns the CounterCashTfrLimit field if non-nil, zero value otherwise.

### GetCounterCashTfrLimitOk

`func (o *HandlerCreateOfficesRequest) GetCounterCashTfrLimitOk() (*float32, bool)`

GetCounterCashTfrLimitOk returns a tuple with the CounterCashTfrLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterCashTfrLimit

`func (o *HandlerCreateOfficesRequest) SetCounterCashTfrLimit(v float32)`

SetCounterCashTfrLimit sets CounterCashTfrLimit field to given value.

### HasCounterCashTfrLimit

`func (o *HandlerCreateOfficesRequest) HasCounterCashTfrLimit() bool`

HasCounterCashTfrLimit returns a boolean if a field has been set.

### GetEnteredBy

`func (o *HandlerCreateOfficesRequest) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *HandlerCreateOfficesRequest) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *HandlerCreateOfficesRequest) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.


### GetMaxAmt

`func (o *HandlerCreateOfficesRequest) GetMaxAmt() float32`

GetMaxAmt returns the MaxAmt field if non-nil, zero value otherwise.

### GetMaxAmtOk

`func (o *HandlerCreateOfficesRequest) GetMaxAmtOk() (*float32, bool)`

GetMaxAmtOk returns a tuple with the MaxAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmt

`func (o *HandlerCreateOfficesRequest) SetMaxAmt(v float32)`

SetMaxAmt sets MaxAmt field to given value.


### GetMinAmt

`func (o *HandlerCreateOfficesRequest) GetMinAmt() float32`

GetMinAmt returns the MinAmt field if non-nil, zero value otherwise.

### GetMinAmtOk

`func (o *HandlerCreateOfficesRequest) GetMinAmtOk() (*float32, bool)`

GetMinAmtOk returns a tuple with the MinAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmt

`func (o *HandlerCreateOfficesRequest) SetMinAmt(v float32)`

SetMinAmt sets MinAmt field to given value.


### GetOfficeId

`func (o *HandlerCreateOfficesRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateOfficesRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateOfficesRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetOfficeName

`func (o *HandlerCreateOfficesRequest) GetOfficeName() string`

GetOfficeName returns the OfficeName field if non-nil, zero value otherwise.

### GetOfficeNameOk

`func (o *HandlerCreateOfficesRequest) GetOfficeNameOk() (*string, bool)`

GetOfficeNameOk returns a tuple with the OfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeName

`func (o *HandlerCreateOfficesRequest) SetOfficeName(v string)`

SetOfficeName sets OfficeName field to given value.


### GetPostageStampsLimit

`func (o *HandlerCreateOfficesRequest) GetPostageStampsLimit() float32`

GetPostageStampsLimit returns the PostageStampsLimit field if non-nil, zero value otherwise.

### GetPostageStampsLimitOk

`func (o *HandlerCreateOfficesRequest) GetPostageStampsLimitOk() (*float32, bool)`

GetPostageStampsLimitOk returns a tuple with the PostageStampsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostageStampsLimit

`func (o *HandlerCreateOfficesRequest) SetPostageStampsLimit(v float32)`

SetPostageStampsLimit sets PostageStampsLimit field to given value.

### HasPostageStampsLimit

`func (o *HandlerCreateOfficesRequest) HasPostageStampsLimit() bool`

HasPostageStampsLimit returns a boolean if a field has been set.

### GetRevenueStampsLimit

`func (o *HandlerCreateOfficesRequest) GetRevenueStampsLimit() float32`

GetRevenueStampsLimit returns the RevenueStampsLimit field if non-nil, zero value otherwise.

### GetRevenueStampsLimitOk

`func (o *HandlerCreateOfficesRequest) GetRevenueStampsLimitOk() (*float32, bool)`

GetRevenueStampsLimitOk returns a tuple with the RevenueStampsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueStampsLimit

`func (o *HandlerCreateOfficesRequest) SetRevenueStampsLimit(v float32)`

SetRevenueStampsLimit sets RevenueStampsLimit field to given value.

### HasRevenueStampsLimit

`func (o *HandlerCreateOfficesRequest) HasRevenueStampsLimit() bool`

HasRevenueStampsLimit returns a boolean if a field has been set.

### GetTransitDuration

`func (o *HandlerCreateOfficesRequest) GetTransitDuration() int32`

GetTransitDuration returns the TransitDuration field if non-nil, zero value otherwise.

### GetTransitDurationOk

`func (o *HandlerCreateOfficesRequest) GetTransitDurationOk() (*int32, bool)`

GetTransitDurationOk returns a tuple with the TransitDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitDuration

`func (o *HandlerCreateOfficesRequest) SetTransitDuration(v int32)`

SetTransitDuration sets TransitDuration field to given value.

### HasTransitDuration

`func (o *HandlerCreateOfficesRequest) HasTransitDuration() bool`

HasTransitDuration returns a boolean if a field has been set.

### GetValidFrom

`func (o *HandlerCreateOfficesRequest) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *HandlerCreateOfficesRequest) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *HandlerCreateOfficesRequest) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



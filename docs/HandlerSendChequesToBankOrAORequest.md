# HandlerSendChequesToBankOrAORequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoSentDate** | Pointer to **string** |  | [optional] 
**BoUserId** | Pointer to **int32** |  | [optional] 
**ChequeAmount** | **float32** |  | 
**ChequeDate** | **string** |  | 
**ChequeNo** | **string** |  | 
**ClearingSentDate** | Pointer to **string** |  | [optional] 
**CurrOfficeId** | **int32** |  | 
**DdoUserId** | Pointer to **int32** |  | [optional] 
**NcddoSentDate** | Pointer to **string** |  | [optional] 
**NcddoUserId** | Pointer to **int32** |  | [optional] 
**SoSentDate** | Pointer to **string** |  | [optional] 
**SoUserId** | Pointer to **int32** |  | [optional] 
**SrcOfficeId** | **int32** |  | 

## Methods

### NewHandlerSendChequesToBankOrAORequest

`func NewHandlerSendChequesToBankOrAORequest(chequeAmount float32, chequeDate string, chequeNo string, currOfficeId int32, srcOfficeId int32, ) *HandlerSendChequesToBankOrAORequest`

NewHandlerSendChequesToBankOrAORequest instantiates a new HandlerSendChequesToBankOrAORequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerSendChequesToBankOrAORequestWithDefaults

`func NewHandlerSendChequesToBankOrAORequestWithDefaults() *HandlerSendChequesToBankOrAORequest`

NewHandlerSendChequesToBankOrAORequestWithDefaults instantiates a new HandlerSendChequesToBankOrAORequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) GetBoSentDate() string`

GetBoSentDate returns the BoSentDate field if non-nil, zero value otherwise.

### GetBoSentDateOk

`func (o *HandlerSendChequesToBankOrAORequest) GetBoSentDateOk() (*string, bool)`

GetBoSentDateOk returns a tuple with the BoSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) SetBoSentDate(v string)`

SetBoSentDate sets BoSentDate field to given value.

### HasBoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) HasBoSentDate() bool`

HasBoSentDate returns a boolean if a field has been set.

### GetBoUserId

`func (o *HandlerSendChequesToBankOrAORequest) GetBoUserId() int32`

GetBoUserId returns the BoUserId field if non-nil, zero value otherwise.

### GetBoUserIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetBoUserIdOk() (*int32, bool)`

GetBoUserIdOk returns a tuple with the BoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoUserId

`func (o *HandlerSendChequesToBankOrAORequest) SetBoUserId(v int32)`

SetBoUserId sets BoUserId field to given value.

### HasBoUserId

`func (o *HandlerSendChequesToBankOrAORequest) HasBoUserId() bool`

HasBoUserId returns a boolean if a field has been set.

### GetChequeAmount

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeAmount() float32`

GetChequeAmount returns the ChequeAmount field if non-nil, zero value otherwise.

### GetChequeAmountOk

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeAmountOk() (*float32, bool)`

GetChequeAmountOk returns a tuple with the ChequeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeAmount

`func (o *HandlerSendChequesToBankOrAORequest) SetChequeAmount(v float32)`

SetChequeAmount sets ChequeAmount field to given value.


### GetChequeDate

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *HandlerSendChequesToBankOrAORequest) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.


### GetChequeNo

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *HandlerSendChequesToBankOrAORequest) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *HandlerSendChequesToBankOrAORequest) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.


### GetClearingSentDate

`func (o *HandlerSendChequesToBankOrAORequest) GetClearingSentDate() string`

GetClearingSentDate returns the ClearingSentDate field if non-nil, zero value otherwise.

### GetClearingSentDateOk

`func (o *HandlerSendChequesToBankOrAORequest) GetClearingSentDateOk() (*string, bool)`

GetClearingSentDateOk returns a tuple with the ClearingSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingSentDate

`func (o *HandlerSendChequesToBankOrAORequest) SetClearingSentDate(v string)`

SetClearingSentDate sets ClearingSentDate field to given value.

### HasClearingSentDate

`func (o *HandlerSendChequesToBankOrAORequest) HasClearingSentDate() bool`

HasClearingSentDate returns a boolean if a field has been set.

### GetCurrOfficeId

`func (o *HandlerSendChequesToBankOrAORequest) GetCurrOfficeId() int32`

GetCurrOfficeId returns the CurrOfficeId field if non-nil, zero value otherwise.

### GetCurrOfficeIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetCurrOfficeIdOk() (*int32, bool)`

GetCurrOfficeIdOk returns a tuple with the CurrOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrOfficeId

`func (o *HandlerSendChequesToBankOrAORequest) SetCurrOfficeId(v int32)`

SetCurrOfficeId sets CurrOfficeId field to given value.


### GetDdoUserId

`func (o *HandlerSendChequesToBankOrAORequest) GetDdoUserId() int32`

GetDdoUserId returns the DdoUserId field if non-nil, zero value otherwise.

### GetDdoUserIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetDdoUserIdOk() (*int32, bool)`

GetDdoUserIdOk returns a tuple with the DdoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdoUserId

`func (o *HandlerSendChequesToBankOrAORequest) SetDdoUserId(v int32)`

SetDdoUserId sets DdoUserId field to given value.

### HasDdoUserId

`func (o *HandlerSendChequesToBankOrAORequest) HasDdoUserId() bool`

HasDdoUserId returns a boolean if a field has been set.

### GetNcddoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) GetNcddoSentDate() string`

GetNcddoSentDate returns the NcddoSentDate field if non-nil, zero value otherwise.

### GetNcddoSentDateOk

`func (o *HandlerSendChequesToBankOrAORequest) GetNcddoSentDateOk() (*string, bool)`

GetNcddoSentDateOk returns a tuple with the NcddoSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcddoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) SetNcddoSentDate(v string)`

SetNcddoSentDate sets NcddoSentDate field to given value.

### HasNcddoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) HasNcddoSentDate() bool`

HasNcddoSentDate returns a boolean if a field has been set.

### GetNcddoUserId

`func (o *HandlerSendChequesToBankOrAORequest) GetNcddoUserId() int32`

GetNcddoUserId returns the NcddoUserId field if non-nil, zero value otherwise.

### GetNcddoUserIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetNcddoUserIdOk() (*int32, bool)`

GetNcddoUserIdOk returns a tuple with the NcddoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcddoUserId

`func (o *HandlerSendChequesToBankOrAORequest) SetNcddoUserId(v int32)`

SetNcddoUserId sets NcddoUserId field to given value.

### HasNcddoUserId

`func (o *HandlerSendChequesToBankOrAORequest) HasNcddoUserId() bool`

HasNcddoUserId returns a boolean if a field has been set.

### GetSoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) GetSoSentDate() string`

GetSoSentDate returns the SoSentDate field if non-nil, zero value otherwise.

### GetSoSentDateOk

`func (o *HandlerSendChequesToBankOrAORequest) GetSoSentDateOk() (*string, bool)`

GetSoSentDateOk returns a tuple with the SoSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) SetSoSentDate(v string)`

SetSoSentDate sets SoSentDate field to given value.

### HasSoSentDate

`func (o *HandlerSendChequesToBankOrAORequest) HasSoSentDate() bool`

HasSoSentDate returns a boolean if a field has been set.

### GetSoUserId

`func (o *HandlerSendChequesToBankOrAORequest) GetSoUserId() int32`

GetSoUserId returns the SoUserId field if non-nil, zero value otherwise.

### GetSoUserIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetSoUserIdOk() (*int32, bool)`

GetSoUserIdOk returns a tuple with the SoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoUserId

`func (o *HandlerSendChequesToBankOrAORequest) SetSoUserId(v int32)`

SetSoUserId sets SoUserId field to given value.

### HasSoUserId

`func (o *HandlerSendChequesToBankOrAORequest) HasSoUserId() bool`

HasSoUserId returns a boolean if a field has been set.

### GetSrcOfficeId

`func (o *HandlerSendChequesToBankOrAORequest) GetSrcOfficeId() int32`

GetSrcOfficeId returns the SrcOfficeId field if non-nil, zero value otherwise.

### GetSrcOfficeIdOk

`func (o *HandlerSendChequesToBankOrAORequest) GetSrcOfficeIdOk() (*int32, bool)`

GetSrcOfficeIdOk returns a tuple with the SrcOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOfficeId

`func (o *HandlerSendChequesToBankOrAORequest) SetSrcOfficeId(v int32)`

SetSrcOfficeId sets SrcOfficeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



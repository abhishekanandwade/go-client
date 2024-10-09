# ResponseCashError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashErrorDetails** | Pointer to [**[]ResponseCashErrorDetails**](ResponseCashErrorDetails.md) |  | [optional] 
**CurrOffice** | Pointer to **int32** |  | [optional] 
**DestAmt** | Pointer to **float32** |  | [optional] 
**DestDetails** | Pointer to **map[string]int32** |  | [optional] 
**DestOffice** | Pointer to **int32** |  | [optional] 
**DestOfficeName** | Pointer to **string** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SrcAmt** | Pointer to **float32** |  | [optional] 
**SrcDetails** | Pointer to **map[string]int32** |  | [optional] 
**SrcOffice** | Pointer to **int32** |  | [optional] 
**SrcOfficeName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransType** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseCashError

`func NewResponseCashError() *ResponseCashError`

NewResponseCashError instantiates a new ResponseCashError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCashErrorWithDefaults

`func NewResponseCashErrorWithDefaults() *ResponseCashError`

NewResponseCashErrorWithDefaults instantiates a new ResponseCashError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashErrorDetails

`func (o *ResponseCashError) GetCashErrorDetails() []ResponseCashErrorDetails`

GetCashErrorDetails returns the CashErrorDetails field if non-nil, zero value otherwise.

### GetCashErrorDetailsOk

`func (o *ResponseCashError) GetCashErrorDetailsOk() (*[]ResponseCashErrorDetails, bool)`

GetCashErrorDetailsOk returns a tuple with the CashErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashErrorDetails

`func (o *ResponseCashError) SetCashErrorDetails(v []ResponseCashErrorDetails)`

SetCashErrorDetails sets CashErrorDetails field to given value.

### HasCashErrorDetails

`func (o *ResponseCashError) HasCashErrorDetails() bool`

HasCashErrorDetails returns a boolean if a field has been set.

### GetCurrOffice

`func (o *ResponseCashError) GetCurrOffice() int32`

GetCurrOffice returns the CurrOffice field if non-nil, zero value otherwise.

### GetCurrOfficeOk

`func (o *ResponseCashError) GetCurrOfficeOk() (*int32, bool)`

GetCurrOfficeOk returns a tuple with the CurrOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrOffice

`func (o *ResponseCashError) SetCurrOffice(v int32)`

SetCurrOffice sets CurrOffice field to given value.

### HasCurrOffice

`func (o *ResponseCashError) HasCurrOffice() bool`

HasCurrOffice returns a boolean if a field has been set.

### GetDestAmt

`func (o *ResponseCashError) GetDestAmt() float32`

GetDestAmt returns the DestAmt field if non-nil, zero value otherwise.

### GetDestAmtOk

`func (o *ResponseCashError) GetDestAmtOk() (*float32, bool)`

GetDestAmtOk returns a tuple with the DestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAmt

`func (o *ResponseCashError) SetDestAmt(v float32)`

SetDestAmt sets DestAmt field to given value.

### HasDestAmt

`func (o *ResponseCashError) HasDestAmt() bool`

HasDestAmt returns a boolean if a field has been set.

### GetDestDetails

`func (o *ResponseCashError) GetDestDetails() map[string]int32`

GetDestDetails returns the DestDetails field if non-nil, zero value otherwise.

### GetDestDetailsOk

`func (o *ResponseCashError) GetDestDetailsOk() (*map[string]int32, bool)`

GetDestDetailsOk returns a tuple with the DestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestDetails

`func (o *ResponseCashError) SetDestDetails(v map[string]int32)`

SetDestDetails sets DestDetails field to given value.

### HasDestDetails

`func (o *ResponseCashError) HasDestDetails() bool`

HasDestDetails returns a boolean if a field has been set.

### GetDestOffice

`func (o *ResponseCashError) GetDestOffice() int32`

GetDestOffice returns the DestOffice field if non-nil, zero value otherwise.

### GetDestOfficeOk

`func (o *ResponseCashError) GetDestOfficeOk() (*int32, bool)`

GetDestOfficeOk returns a tuple with the DestOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOffice

`func (o *ResponseCashError) SetDestOffice(v int32)`

SetDestOffice sets DestOffice field to given value.

### HasDestOffice

`func (o *ResponseCashError) HasDestOffice() bool`

HasDestOffice returns a boolean if a field has been set.

### GetDestOfficeName

`func (o *ResponseCashError) GetDestOfficeName() string`

GetDestOfficeName returns the DestOfficeName field if non-nil, zero value otherwise.

### GetDestOfficeNameOk

`func (o *ResponseCashError) GetDestOfficeNameOk() (*string, bool)`

GetDestOfficeNameOk returns a tuple with the DestOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOfficeName

`func (o *ResponseCashError) SetDestOfficeName(v string)`

SetDestOfficeName sets DestOfficeName field to given value.

### HasDestOfficeName

`func (o *ResponseCashError) HasDestOfficeName() bool`

HasDestOfficeName returns a boolean if a field has been set.

### GetErrorId

`func (o *ResponseCashError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ResponseCashError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ResponseCashError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ResponseCashError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseCashError) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseCashError) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseCashError) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseCashError) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSrcAmt

`func (o *ResponseCashError) GetSrcAmt() float32`

GetSrcAmt returns the SrcAmt field if non-nil, zero value otherwise.

### GetSrcAmtOk

`func (o *ResponseCashError) GetSrcAmtOk() (*float32, bool)`

GetSrcAmtOk returns a tuple with the SrcAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAmt

`func (o *ResponseCashError) SetSrcAmt(v float32)`

SetSrcAmt sets SrcAmt field to given value.

### HasSrcAmt

`func (o *ResponseCashError) HasSrcAmt() bool`

HasSrcAmt returns a boolean if a field has been set.

### GetSrcDetails

`func (o *ResponseCashError) GetSrcDetails() map[string]int32`

GetSrcDetails returns the SrcDetails field if non-nil, zero value otherwise.

### GetSrcDetailsOk

`func (o *ResponseCashError) GetSrcDetailsOk() (*map[string]int32, bool)`

GetSrcDetailsOk returns a tuple with the SrcDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDetails

`func (o *ResponseCashError) SetSrcDetails(v map[string]int32)`

SetSrcDetails sets SrcDetails field to given value.

### HasSrcDetails

`func (o *ResponseCashError) HasSrcDetails() bool`

HasSrcDetails returns a boolean if a field has been set.

### GetSrcOffice

`func (o *ResponseCashError) GetSrcOffice() int32`

GetSrcOffice returns the SrcOffice field if non-nil, zero value otherwise.

### GetSrcOfficeOk

`func (o *ResponseCashError) GetSrcOfficeOk() (*int32, bool)`

GetSrcOfficeOk returns a tuple with the SrcOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOffice

`func (o *ResponseCashError) SetSrcOffice(v int32)`

SetSrcOffice sets SrcOffice field to given value.

### HasSrcOffice

`func (o *ResponseCashError) HasSrcOffice() bool`

HasSrcOffice returns a boolean if a field has been set.

### GetSrcOfficeName

`func (o *ResponseCashError) GetSrcOfficeName() string`

GetSrcOfficeName returns the SrcOfficeName field if non-nil, zero value otherwise.

### GetSrcOfficeNameOk

`func (o *ResponseCashError) GetSrcOfficeNameOk() (*string, bool)`

GetSrcOfficeNameOk returns a tuple with the SrcOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOfficeName

`func (o *ResponseCashError) SetSrcOfficeName(v string)`

SetSrcOfficeName sets SrcOfficeName field to given value.

### HasSrcOfficeName

`func (o *ResponseCashError) HasSrcOfficeName() bool`

HasSrcOfficeName returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseCashError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseCashError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseCashError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseCashError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseCashError) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseCashError) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseCashError) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseCashError) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseCashError) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseCashError) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseCashError) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseCashError) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransType

`func (o *ResponseCashError) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ResponseCashError) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ResponseCashError) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ResponseCashError) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseCashError) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseCashError) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseCashError) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseCashError) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



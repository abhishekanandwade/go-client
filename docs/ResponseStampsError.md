# ResponseStampsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrOffice** | Pointer to **int32** |  | [optional] 
**DestAmt** | Pointer to **float32** |  | [optional] 
**DestDetails** | Pointer to **map[string]int32** |  | [optional] 
**DestDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**DestOffice** | Pointer to **int32** |  | [optional] 
**DestOfficeName** | Pointer to **string** |  | [optional] 
**DiffDetails** | Pointer to **map[string]int32** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SrcAmt** | Pointer to **float32** |  | [optional] 
**SrcDetails** | Pointer to **map[string]int32** |  | [optional] 
**SrcDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**SrcOffice** | Pointer to **int32** |  | [optional] 
**SrcOfficeName** | Pointer to **string** |  | [optional] 
**StampsErrorDetails** | Pointer to [**[]ResponseStampsErrorDetails**](ResponseStampsErrorDetails.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransType** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStampsError

`func NewResponseStampsError() *ResponseStampsError`

NewResponseStampsError instantiates a new ResponseStampsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsErrorWithDefaults

`func NewResponseStampsErrorWithDefaults() *ResponseStampsError`

NewResponseStampsErrorWithDefaults instantiates a new ResponseStampsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrOffice

`func (o *ResponseStampsError) GetCurrOffice() int32`

GetCurrOffice returns the CurrOffice field if non-nil, zero value otherwise.

### GetCurrOfficeOk

`func (o *ResponseStampsError) GetCurrOfficeOk() (*int32, bool)`

GetCurrOfficeOk returns a tuple with the CurrOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrOffice

`func (o *ResponseStampsError) SetCurrOffice(v int32)`

SetCurrOffice sets CurrOffice field to given value.

### HasCurrOffice

`func (o *ResponseStampsError) HasCurrOffice() bool`

HasCurrOffice returns a boolean if a field has been set.

### GetDestAmt

`func (o *ResponseStampsError) GetDestAmt() float32`

GetDestAmt returns the DestAmt field if non-nil, zero value otherwise.

### GetDestAmtOk

`func (o *ResponseStampsError) GetDestAmtOk() (*float32, bool)`

GetDestAmtOk returns a tuple with the DestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAmt

`func (o *ResponseStampsError) SetDestAmt(v float32)`

SetDestAmt sets DestAmt field to given value.

### HasDestAmt

`func (o *ResponseStampsError) HasDestAmt() bool`

HasDestAmt returns a boolean if a field has been set.

### GetDestDetails

`func (o *ResponseStampsError) GetDestDetails() map[string]int32`

GetDestDetails returns the DestDetails field if non-nil, zero value otherwise.

### GetDestDetailsOk

`func (o *ResponseStampsError) GetDestDetailsOk() (*map[string]int32, bool)`

GetDestDetailsOk returns a tuple with the DestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestDetails

`func (o *ResponseStampsError) SetDestDetails(v map[string]int32)`

SetDestDetails sets DestDetails field to given value.

### HasDestDetails

`func (o *ResponseStampsError) HasDestDetails() bool`

HasDestDetails returns a boolean if a field has been set.

### GetDestDetailsDesc

`func (o *ResponseStampsError) GetDestDetailsDesc() []ResponseStampdetails`

GetDestDetailsDesc returns the DestDetailsDesc field if non-nil, zero value otherwise.

### GetDestDetailsDescOk

`func (o *ResponseStampsError) GetDestDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetDestDetailsDescOk returns a tuple with the DestDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestDetailsDesc

`func (o *ResponseStampsError) SetDestDetailsDesc(v []ResponseStampdetails)`

SetDestDetailsDesc sets DestDetailsDesc field to given value.

### HasDestDetailsDesc

`func (o *ResponseStampsError) HasDestDetailsDesc() bool`

HasDestDetailsDesc returns a boolean if a field has been set.

### GetDestOffice

`func (o *ResponseStampsError) GetDestOffice() int32`

GetDestOffice returns the DestOffice field if non-nil, zero value otherwise.

### GetDestOfficeOk

`func (o *ResponseStampsError) GetDestOfficeOk() (*int32, bool)`

GetDestOfficeOk returns a tuple with the DestOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOffice

`func (o *ResponseStampsError) SetDestOffice(v int32)`

SetDestOffice sets DestOffice field to given value.

### HasDestOffice

`func (o *ResponseStampsError) HasDestOffice() bool`

HasDestOffice returns a boolean if a field has been set.

### GetDestOfficeName

`func (o *ResponseStampsError) GetDestOfficeName() string`

GetDestOfficeName returns the DestOfficeName field if non-nil, zero value otherwise.

### GetDestOfficeNameOk

`func (o *ResponseStampsError) GetDestOfficeNameOk() (*string, bool)`

GetDestOfficeNameOk returns a tuple with the DestOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOfficeName

`func (o *ResponseStampsError) SetDestOfficeName(v string)`

SetDestOfficeName sets DestOfficeName field to given value.

### HasDestOfficeName

`func (o *ResponseStampsError) HasDestOfficeName() bool`

HasDestOfficeName returns a boolean if a field has been set.

### GetDiffDetails

`func (o *ResponseStampsError) GetDiffDetails() map[string]int32`

GetDiffDetails returns the DiffDetails field if non-nil, zero value otherwise.

### GetDiffDetailsOk

`func (o *ResponseStampsError) GetDiffDetailsOk() (*map[string]int32, bool)`

GetDiffDetailsOk returns a tuple with the DiffDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffDetails

`func (o *ResponseStampsError) SetDiffDetails(v map[string]int32)`

SetDiffDetails sets DiffDetails field to given value.

### HasDiffDetails

`func (o *ResponseStampsError) HasDiffDetails() bool`

HasDiffDetails returns a boolean if a field has been set.

### GetErrorId

`func (o *ResponseStampsError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ResponseStampsError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ResponseStampsError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ResponseStampsError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsError) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsError) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsError) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsError) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSrcAmt

`func (o *ResponseStampsError) GetSrcAmt() float32`

GetSrcAmt returns the SrcAmt field if non-nil, zero value otherwise.

### GetSrcAmtOk

`func (o *ResponseStampsError) GetSrcAmtOk() (*float32, bool)`

GetSrcAmtOk returns a tuple with the SrcAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAmt

`func (o *ResponseStampsError) SetSrcAmt(v float32)`

SetSrcAmt sets SrcAmt field to given value.

### HasSrcAmt

`func (o *ResponseStampsError) HasSrcAmt() bool`

HasSrcAmt returns a boolean if a field has been set.

### GetSrcDetails

`func (o *ResponseStampsError) GetSrcDetails() map[string]int32`

GetSrcDetails returns the SrcDetails field if non-nil, zero value otherwise.

### GetSrcDetailsOk

`func (o *ResponseStampsError) GetSrcDetailsOk() (*map[string]int32, bool)`

GetSrcDetailsOk returns a tuple with the SrcDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDetails

`func (o *ResponseStampsError) SetSrcDetails(v map[string]int32)`

SetSrcDetails sets SrcDetails field to given value.

### HasSrcDetails

`func (o *ResponseStampsError) HasSrcDetails() bool`

HasSrcDetails returns a boolean if a field has been set.

### GetSrcDetailsDesc

`func (o *ResponseStampsError) GetSrcDetailsDesc() []ResponseStampdetails`

GetSrcDetailsDesc returns the SrcDetailsDesc field if non-nil, zero value otherwise.

### GetSrcDetailsDescOk

`func (o *ResponseStampsError) GetSrcDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetSrcDetailsDescOk returns a tuple with the SrcDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDetailsDesc

`func (o *ResponseStampsError) SetSrcDetailsDesc(v []ResponseStampdetails)`

SetSrcDetailsDesc sets SrcDetailsDesc field to given value.

### HasSrcDetailsDesc

`func (o *ResponseStampsError) HasSrcDetailsDesc() bool`

HasSrcDetailsDesc returns a boolean if a field has been set.

### GetSrcOffice

`func (o *ResponseStampsError) GetSrcOffice() int32`

GetSrcOffice returns the SrcOffice field if non-nil, zero value otherwise.

### GetSrcOfficeOk

`func (o *ResponseStampsError) GetSrcOfficeOk() (*int32, bool)`

GetSrcOfficeOk returns a tuple with the SrcOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOffice

`func (o *ResponseStampsError) SetSrcOffice(v int32)`

SetSrcOffice sets SrcOffice field to given value.

### HasSrcOffice

`func (o *ResponseStampsError) HasSrcOffice() bool`

HasSrcOffice returns a boolean if a field has been set.

### GetSrcOfficeName

`func (o *ResponseStampsError) GetSrcOfficeName() string`

GetSrcOfficeName returns the SrcOfficeName field if non-nil, zero value otherwise.

### GetSrcOfficeNameOk

`func (o *ResponseStampsError) GetSrcOfficeNameOk() (*string, bool)`

GetSrcOfficeNameOk returns a tuple with the SrcOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOfficeName

`func (o *ResponseStampsError) SetSrcOfficeName(v string)`

SetSrcOfficeName sets SrcOfficeName field to given value.

### HasSrcOfficeName

`func (o *ResponseStampsError) HasSrcOfficeName() bool`

HasSrcOfficeName returns a boolean if a field has been set.

### GetStampsErrorDetails

`func (o *ResponseStampsError) GetStampsErrorDetails() []ResponseStampsErrorDetails`

GetStampsErrorDetails returns the StampsErrorDetails field if non-nil, zero value otherwise.

### GetStampsErrorDetailsOk

`func (o *ResponseStampsError) GetStampsErrorDetailsOk() (*[]ResponseStampsErrorDetails, bool)`

GetStampsErrorDetailsOk returns a tuple with the StampsErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampsErrorDetails

`func (o *ResponseStampsError) SetStampsErrorDetails(v []ResponseStampsErrorDetails)`

SetStampsErrorDetails sets StampsErrorDetails field to given value.

### HasStampsErrorDetails

`func (o *ResponseStampsError) HasStampsErrorDetails() bool`

HasStampsErrorDetails returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseStampsError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseStampsError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseStampsError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseStampsError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsError) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsError) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsError) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsError) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseStampsError) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseStampsError) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseStampsError) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseStampsError) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransType

`func (o *ResponseStampsError) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ResponseStampsError) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ResponseStampsError) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ResponseStampsError) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseStampsError) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseStampsError) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseStampsError) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseStampsError) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



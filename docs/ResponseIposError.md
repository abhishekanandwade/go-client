# ResponseIposError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrOffice** | Pointer to **int32** |  | [optional] 
**DestAmt** | Pointer to **float32** |  | [optional] 
**DestDetails** | Pointer to **string** |  | [optional] 
**DestOffice** | Pointer to **int32** |  | [optional] 
**DestOfficeName** | Pointer to **string** |  | [optional] 
**DiffDetails** | Pointer to **string** |  | [optional] 
**ErrorId** | Pointer to **string** |  | [optional] 
**IposErrorDetails** | Pointer to [**[]ResponseIposErrorDetails**](ResponseIposErrorDetails.md) |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SrcAmt** | Pointer to **float32** |  | [optional] 
**SrcDetails** | Pointer to **string** |  | [optional] 
**SrcOffice** | Pointer to **int32** |  | [optional] 
**SrcOfficeName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TransType** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseIposError

`func NewResponseIposError() *ResponseIposError`

NewResponseIposError instantiates a new ResponseIposError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseIposErrorWithDefaults

`func NewResponseIposErrorWithDefaults() *ResponseIposError`

NewResponseIposErrorWithDefaults instantiates a new ResponseIposError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrOffice

`func (o *ResponseIposError) GetCurrOffice() int32`

GetCurrOffice returns the CurrOffice field if non-nil, zero value otherwise.

### GetCurrOfficeOk

`func (o *ResponseIposError) GetCurrOfficeOk() (*int32, bool)`

GetCurrOfficeOk returns a tuple with the CurrOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrOffice

`func (o *ResponseIposError) SetCurrOffice(v int32)`

SetCurrOffice sets CurrOffice field to given value.

### HasCurrOffice

`func (o *ResponseIposError) HasCurrOffice() bool`

HasCurrOffice returns a boolean if a field has been set.

### GetDestAmt

`func (o *ResponseIposError) GetDestAmt() float32`

GetDestAmt returns the DestAmt field if non-nil, zero value otherwise.

### GetDestAmtOk

`func (o *ResponseIposError) GetDestAmtOk() (*float32, bool)`

GetDestAmtOk returns a tuple with the DestAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestAmt

`func (o *ResponseIposError) SetDestAmt(v float32)`

SetDestAmt sets DestAmt field to given value.

### HasDestAmt

`func (o *ResponseIposError) HasDestAmt() bool`

HasDestAmt returns a boolean if a field has been set.

### GetDestDetails

`func (o *ResponseIposError) GetDestDetails() string`

GetDestDetails returns the DestDetails field if non-nil, zero value otherwise.

### GetDestDetailsOk

`func (o *ResponseIposError) GetDestDetailsOk() (*string, bool)`

GetDestDetailsOk returns a tuple with the DestDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestDetails

`func (o *ResponseIposError) SetDestDetails(v string)`

SetDestDetails sets DestDetails field to given value.

### HasDestDetails

`func (o *ResponseIposError) HasDestDetails() bool`

HasDestDetails returns a boolean if a field has been set.

### GetDestOffice

`func (o *ResponseIposError) GetDestOffice() int32`

GetDestOffice returns the DestOffice field if non-nil, zero value otherwise.

### GetDestOfficeOk

`func (o *ResponseIposError) GetDestOfficeOk() (*int32, bool)`

GetDestOfficeOk returns a tuple with the DestOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOffice

`func (o *ResponseIposError) SetDestOffice(v int32)`

SetDestOffice sets DestOffice field to given value.

### HasDestOffice

`func (o *ResponseIposError) HasDestOffice() bool`

HasDestOffice returns a boolean if a field has been set.

### GetDestOfficeName

`func (o *ResponseIposError) GetDestOfficeName() string`

GetDestOfficeName returns the DestOfficeName field if non-nil, zero value otherwise.

### GetDestOfficeNameOk

`func (o *ResponseIposError) GetDestOfficeNameOk() (*string, bool)`

GetDestOfficeNameOk returns a tuple with the DestOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestOfficeName

`func (o *ResponseIposError) SetDestOfficeName(v string)`

SetDestOfficeName sets DestOfficeName field to given value.

### HasDestOfficeName

`func (o *ResponseIposError) HasDestOfficeName() bool`

HasDestOfficeName returns a boolean if a field has been set.

### GetDiffDetails

`func (o *ResponseIposError) GetDiffDetails() string`

GetDiffDetails returns the DiffDetails field if non-nil, zero value otherwise.

### GetDiffDetailsOk

`func (o *ResponseIposError) GetDiffDetailsOk() (*string, bool)`

GetDiffDetailsOk returns a tuple with the DiffDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffDetails

`func (o *ResponseIposError) SetDiffDetails(v string)`

SetDiffDetails sets DiffDetails field to given value.

### HasDiffDetails

`func (o *ResponseIposError) HasDiffDetails() bool`

HasDiffDetails returns a boolean if a field has been set.

### GetErrorId

`func (o *ResponseIposError) GetErrorId() string`

GetErrorId returns the ErrorId field if non-nil, zero value otherwise.

### GetErrorIdOk

`func (o *ResponseIposError) GetErrorIdOk() (*string, bool)`

GetErrorIdOk returns a tuple with the ErrorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorId

`func (o *ResponseIposError) SetErrorId(v string)`

SetErrorId sets ErrorId field to given value.

### HasErrorId

`func (o *ResponseIposError) HasErrorId() bool`

HasErrorId returns a boolean if a field has been set.

### GetIposErrorDetails

`func (o *ResponseIposError) GetIposErrorDetails() []ResponseIposErrorDetails`

GetIposErrorDetails returns the IposErrorDetails field if non-nil, zero value otherwise.

### GetIposErrorDetailsOk

`func (o *ResponseIposError) GetIposErrorDetailsOk() (*[]ResponseIposErrorDetails, bool)`

GetIposErrorDetailsOk returns a tuple with the IposErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIposErrorDetails

`func (o *ResponseIposError) SetIposErrorDetails(v []ResponseIposErrorDetails)`

SetIposErrorDetails sets IposErrorDetails field to given value.

### HasIposErrorDetails

`func (o *ResponseIposError) HasIposErrorDetails() bool`

HasIposErrorDetails returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseIposError) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseIposError) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseIposError) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseIposError) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSrcAmt

`func (o *ResponseIposError) GetSrcAmt() float32`

GetSrcAmt returns the SrcAmt field if non-nil, zero value otherwise.

### GetSrcAmtOk

`func (o *ResponseIposError) GetSrcAmtOk() (*float32, bool)`

GetSrcAmtOk returns a tuple with the SrcAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAmt

`func (o *ResponseIposError) SetSrcAmt(v float32)`

SetSrcAmt sets SrcAmt field to given value.

### HasSrcAmt

`func (o *ResponseIposError) HasSrcAmt() bool`

HasSrcAmt returns a boolean if a field has been set.

### GetSrcDetails

`func (o *ResponseIposError) GetSrcDetails() string`

GetSrcDetails returns the SrcDetails field if non-nil, zero value otherwise.

### GetSrcDetailsOk

`func (o *ResponseIposError) GetSrcDetailsOk() (*string, bool)`

GetSrcDetailsOk returns a tuple with the SrcDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcDetails

`func (o *ResponseIposError) SetSrcDetails(v string)`

SetSrcDetails sets SrcDetails field to given value.

### HasSrcDetails

`func (o *ResponseIposError) HasSrcDetails() bool`

HasSrcDetails returns a boolean if a field has been set.

### GetSrcOffice

`func (o *ResponseIposError) GetSrcOffice() int32`

GetSrcOffice returns the SrcOffice field if non-nil, zero value otherwise.

### GetSrcOfficeOk

`func (o *ResponseIposError) GetSrcOfficeOk() (*int32, bool)`

GetSrcOfficeOk returns a tuple with the SrcOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOffice

`func (o *ResponseIposError) SetSrcOffice(v int32)`

SetSrcOffice sets SrcOffice field to given value.

### HasSrcOffice

`func (o *ResponseIposError) HasSrcOffice() bool`

HasSrcOffice returns a boolean if a field has been set.

### GetSrcOfficeName

`func (o *ResponseIposError) GetSrcOfficeName() string`

GetSrcOfficeName returns the SrcOfficeName field if non-nil, zero value otherwise.

### GetSrcOfficeNameOk

`func (o *ResponseIposError) GetSrcOfficeNameOk() (*string, bool)`

GetSrcOfficeNameOk returns a tuple with the SrcOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcOfficeName

`func (o *ResponseIposError) SetSrcOfficeName(v string)`

SetSrcOfficeName sets SrcOfficeName field to given value.

### HasSrcOfficeName

`func (o *ResponseIposError) HasSrcOfficeName() bool`

HasSrcOfficeName returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseIposError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseIposError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseIposError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseIposError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseIposError) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseIposError) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseIposError) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseIposError) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseIposError) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseIposError) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseIposError) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseIposError) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTransType

`func (o *ResponseIposError) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ResponseIposError) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ResponseIposError) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ResponseIposError) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseIposError) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseIposError) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseIposError) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseIposError) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



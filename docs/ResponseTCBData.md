# ResponseTCBData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**RctOrPmt** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**TcSerial** | Pointer to **string** |  | [optional] 
**ToName** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseTCBData

`func NewResponseTCBData() *ResponseTCBData`

NewResponseTCBData instantiates a new ResponseTCBData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTCBDataWithDefaults

`func NewResponseTCBDataWithDefaults() *ResponseTCBData`

NewResponseTCBDataWithDefaults instantiates a new ResponseTCBData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponseTCBData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponseTCBData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponseTCBData) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponseTCBData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPurpose

`func (o *ResponseTCBData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ResponseTCBData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ResponseTCBData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *ResponseTCBData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRctOrPmt

`func (o *ResponseTCBData) GetRctOrPmt() string`

GetRctOrPmt returns the RctOrPmt field if non-nil, zero value otherwise.

### GetRctOrPmtOk

`func (o *ResponseTCBData) GetRctOrPmtOk() (*string, bool)`

GetRctOrPmtOk returns a tuple with the RctOrPmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRctOrPmt

`func (o *ResponseTCBData) SetRctOrPmt(v string)`

SetRctOrPmt sets RctOrPmt field to given value.

### HasRctOrPmt

`func (o *ResponseTCBData) HasRctOrPmt() bool`

HasRctOrPmt returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseTCBData) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseTCBData) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseTCBData) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseTCBData) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTcSerial

`func (o *ResponseTCBData) GetTcSerial() string`

GetTcSerial returns the TcSerial field if non-nil, zero value otherwise.

### GetTcSerialOk

`func (o *ResponseTCBData) GetTcSerialOk() (*string, bool)`

GetTcSerialOk returns a tuple with the TcSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcSerial

`func (o *ResponseTCBData) SetTcSerial(v string)`

SetTcSerial sets TcSerial field to given value.

### HasTcSerial

`func (o *ResponseTCBData) HasTcSerial() bool`

HasTcSerial returns a boolean if a field has been set.

### GetToName

`func (o *ResponseTCBData) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *ResponseTCBData) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *ResponseTCBData) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *ResponseTCBData) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseTCBData) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseTCBData) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseTCBData) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseTCBData) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseTCBData) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseTCBData) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseTCBData) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseTCBData) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseTCBData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseTCBData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseTCBData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseTCBData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



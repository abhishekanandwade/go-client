# DomainTCBData

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

### NewDomainTCBData

`func NewDomainTCBData() *DomainTCBData`

NewDomainTCBData instantiates a new DomainTCBData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTCBDataWithDefaults

`func NewDomainTCBDataWithDefaults() *DomainTCBData`

NewDomainTCBDataWithDefaults instantiates a new DomainTCBData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *DomainTCBData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DomainTCBData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DomainTCBData) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *DomainTCBData) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetPurpose

`func (o *DomainTCBData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *DomainTCBData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *DomainTCBData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *DomainTCBData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetRctOrPmt

`func (o *DomainTCBData) GetRctOrPmt() string`

GetRctOrPmt returns the RctOrPmt field if non-nil, zero value otherwise.

### GetRctOrPmtOk

`func (o *DomainTCBData) GetRctOrPmtOk() (*string, bool)`

GetRctOrPmtOk returns a tuple with the RctOrPmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRctOrPmt

`func (o *DomainTCBData) SetRctOrPmt(v string)`

SetRctOrPmt sets RctOrPmt field to given value.

### HasRctOrPmt

`func (o *DomainTCBData) HasRctOrPmt() bool`

HasRctOrPmt returns a boolean if a field has been set.

### GetRemarks

`func (o *DomainTCBData) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *DomainTCBData) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *DomainTCBData) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *DomainTCBData) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTcSerial

`func (o *DomainTCBData) GetTcSerial() string`

GetTcSerial returns the TcSerial field if non-nil, zero value otherwise.

### GetTcSerialOk

`func (o *DomainTCBData) GetTcSerialOk() (*string, bool)`

GetTcSerialOk returns a tuple with the TcSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcSerial

`func (o *DomainTCBData) SetTcSerial(v string)`

SetTcSerial sets TcSerial field to given value.

### HasTcSerial

`func (o *DomainTCBData) HasTcSerial() bool`

HasTcSerial returns a boolean if a field has been set.

### GetToName

`func (o *DomainTCBData) GetToName() string`

GetToName returns the ToName field if non-nil, zero value otherwise.

### GetToNameOk

`func (o *DomainTCBData) GetToNameOk() (*string, bool)`

GetToNameOk returns a tuple with the ToName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToName

`func (o *DomainTCBData) SetToName(v string)`

SetToName sets ToName field to given value.

### HasToName

`func (o *DomainTCBData) HasToName() bool`

HasToName returns a boolean if a field has been set.

### GetTransDate

`func (o *DomainTCBData) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *DomainTCBData) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *DomainTCBData) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *DomainTCBData) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *DomainTCBData) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *DomainTCBData) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *DomainTCBData) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *DomainTCBData) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetUserId

`func (o *DomainTCBData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DomainTCBData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DomainTCBData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DomainTCBData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



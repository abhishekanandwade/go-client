# HandlerUpdateStampErrorTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiffDetails** | Pointer to **map[string]int32** |  | [optional] 
**OfficeId** | **int32** |  | 
**Remarks** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | **int32** |  | 

## Methods

### NewHandlerUpdateStampErrorTransactionsRequest

`func NewHandlerUpdateStampErrorTransactionsRequest(officeId int32, remarks string, userId int32, ) *HandlerUpdateStampErrorTransactionsRequest`

NewHandlerUpdateStampErrorTransactionsRequest instantiates a new HandlerUpdateStampErrorTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateStampErrorTransactionsRequestWithDefaults

`func NewHandlerUpdateStampErrorTransactionsRequestWithDefaults() *HandlerUpdateStampErrorTransactionsRequest`

NewHandlerUpdateStampErrorTransactionsRequestWithDefaults instantiates a new HandlerUpdateStampErrorTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffDetails

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetDiffDetails() map[string]int32`

GetDiffDetails returns the DiffDetails field if non-nil, zero value otherwise.

### GetDiffDetailsOk

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetDiffDetailsOk() (*map[string]int32, bool)`

GetDiffDetailsOk returns a tuple with the DiffDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffDetails

`func (o *HandlerUpdateStampErrorTransactionsRequest) SetDiffDetails(v map[string]int32)`

SetDiffDetails sets DiffDetails field to given value.

### HasDiffDetails

`func (o *HandlerUpdateStampErrorTransactionsRequest) HasDiffDetails() bool`

HasDiffDetails returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerUpdateStampErrorTransactionsRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetRemarks

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerUpdateStampErrorTransactionsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.


### GetStatus

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerUpdateStampErrorTransactionsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlerUpdateStampErrorTransactionsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerUpdateStampErrorTransactionsRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerUpdateStampErrorTransactionsRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



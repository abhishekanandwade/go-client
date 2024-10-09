# HandlerUpdateCashErrorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OfficeId** | **int32** |  | 
**Remarks** | **string** |  | 
**Status** | Pointer to **string** |  | [optional] 
**UserId** | **int32** |  | 

## Methods

### NewHandlerUpdateCashErrorsRequest

`func NewHandlerUpdateCashErrorsRequest(officeId int32, remarks string, userId int32, ) *HandlerUpdateCashErrorsRequest`

NewHandlerUpdateCashErrorsRequest instantiates a new HandlerUpdateCashErrorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateCashErrorsRequestWithDefaults

`func NewHandlerUpdateCashErrorsRequestWithDefaults() *HandlerUpdateCashErrorsRequest`

NewHandlerUpdateCashErrorsRequestWithDefaults instantiates a new HandlerUpdateCashErrorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOfficeId

`func (o *HandlerUpdateCashErrorsRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerUpdateCashErrorsRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerUpdateCashErrorsRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetRemarks

`func (o *HandlerUpdateCashErrorsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerUpdateCashErrorsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerUpdateCashErrorsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.


### GetStatus

`func (o *HandlerUpdateCashErrorsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HandlerUpdateCashErrorsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HandlerUpdateCashErrorsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HandlerUpdateCashErrorsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerUpdateCashErrorsRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerUpdateCashErrorsRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerUpdateCashErrorsRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



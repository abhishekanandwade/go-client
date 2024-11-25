# HandlerUpdateStampSoiledRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverId** | **int32** |  | 
**DisposalDetails** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 

## Methods

### NewHandlerUpdateStampSoiledRequest

`func NewHandlerUpdateStampSoiledRequest(approverId int32, ) *HandlerUpdateStampSoiledRequest`

NewHandlerUpdateStampSoiledRequest instantiates a new HandlerUpdateStampSoiledRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateStampSoiledRequestWithDefaults

`func NewHandlerUpdateStampSoiledRequestWithDefaults() *HandlerUpdateStampSoiledRequest`

NewHandlerUpdateStampSoiledRequestWithDefaults instantiates a new HandlerUpdateStampSoiledRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverId

`func (o *HandlerUpdateStampSoiledRequest) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerUpdateStampSoiledRequest) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerUpdateStampSoiledRequest) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### GetDisposalDetails

`func (o *HandlerUpdateStampSoiledRequest) GetDisposalDetails() string`

GetDisposalDetails returns the DisposalDetails field if non-nil, zero value otherwise.

### GetDisposalDetailsOk

`func (o *HandlerUpdateStampSoiledRequest) GetDisposalDetailsOk() (*string, bool)`

GetDisposalDetailsOk returns a tuple with the DisposalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposalDetails

`func (o *HandlerUpdateStampSoiledRequest) SetDisposalDetails(v string)`

SetDisposalDetails sets DisposalDetails field to given value.

### HasDisposalDetails

`func (o *HandlerUpdateStampSoiledRequest) HasDisposalDetails() bool`

HasDisposalDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerUpdateStampSoiledRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateStampSoiledRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateStampSoiledRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateStampSoiledRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



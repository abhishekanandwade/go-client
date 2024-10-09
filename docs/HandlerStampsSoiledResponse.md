# HandlerStampsSoiledResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** |  | [optional] 
**ApproverDate** | Pointer to **string** |  | [optional] 
**ApproverId** | Pointer to **int32** |  | [optional] 
**DisposalDetails** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SoiledDetails** | Pointer to **map[string]int32** |  | [optional] 
**StampDetailsDesc** | Pointer to [**[]HandlerStampdetails**](HandlerStampdetails.md) |  | [optional] 
**TranId** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlerStampsSoiledResponse

`func NewHandlerStampsSoiledResponse() *HandlerStampsSoiledResponse`

NewHandlerStampsSoiledResponse instantiates a new HandlerStampsSoiledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerStampsSoiledResponseWithDefaults

`func NewHandlerStampsSoiledResponseWithDefaults() *HandlerStampsSoiledResponse`

NewHandlerStampsSoiledResponseWithDefaults instantiates a new HandlerStampsSoiledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *HandlerStampsSoiledResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *HandlerStampsSoiledResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *HandlerStampsSoiledResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *HandlerStampsSoiledResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApproverDate

`func (o *HandlerStampsSoiledResponse) GetApproverDate() string`

GetApproverDate returns the ApproverDate field if non-nil, zero value otherwise.

### GetApproverDateOk

`func (o *HandlerStampsSoiledResponse) GetApproverDateOk() (*string, bool)`

GetApproverDateOk returns a tuple with the ApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverDate

`func (o *HandlerStampsSoiledResponse) SetApproverDate(v string)`

SetApproverDate sets ApproverDate field to given value.

### HasApproverDate

`func (o *HandlerStampsSoiledResponse) HasApproverDate() bool`

HasApproverDate returns a boolean if a field has been set.

### GetApproverId

`func (o *HandlerStampsSoiledResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerStampsSoiledResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerStampsSoiledResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *HandlerStampsSoiledResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDisposalDetails

`func (o *HandlerStampsSoiledResponse) GetDisposalDetails() string`

GetDisposalDetails returns the DisposalDetails field if non-nil, zero value otherwise.

### GetDisposalDetailsOk

`func (o *HandlerStampsSoiledResponse) GetDisposalDetailsOk() (*string, bool)`

GetDisposalDetailsOk returns a tuple with the DisposalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposalDetails

`func (o *HandlerStampsSoiledResponse) SetDisposalDetails(v string)`

SetDisposalDetails sets DisposalDetails field to given value.

### HasDisposalDetails

`func (o *HandlerStampsSoiledResponse) HasDisposalDetails() bool`

HasDisposalDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerStampsSoiledResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerStampsSoiledResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerStampsSoiledResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerStampsSoiledResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerStampsSoiledResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerStampsSoiledResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerStampsSoiledResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *HandlerStampsSoiledResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerStampsSoiledResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerStampsSoiledResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerStampsSoiledResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerStampsSoiledResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSoiledDetails

`func (o *HandlerStampsSoiledResponse) GetSoiledDetails() map[string]int32`

GetSoiledDetails returns the SoiledDetails field if non-nil, zero value otherwise.

### GetSoiledDetailsOk

`func (o *HandlerStampsSoiledResponse) GetSoiledDetailsOk() (*map[string]int32, bool)`

GetSoiledDetailsOk returns a tuple with the SoiledDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoiledDetails

`func (o *HandlerStampsSoiledResponse) SetSoiledDetails(v map[string]int32)`

SetSoiledDetails sets SoiledDetails field to given value.

### HasSoiledDetails

`func (o *HandlerStampsSoiledResponse) HasSoiledDetails() bool`

HasSoiledDetails returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *HandlerStampsSoiledResponse) GetStampDetailsDesc() []HandlerStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *HandlerStampsSoiledResponse) GetStampDetailsDescOk() (*[]HandlerStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *HandlerStampsSoiledResponse) SetStampDetailsDesc(v []HandlerStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *HandlerStampsSoiledResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTranId

`func (o *HandlerStampsSoiledResponse) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *HandlerStampsSoiledResponse) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *HandlerStampsSoiledResponse) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *HandlerStampsSoiledResponse) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransDate

`func (o *HandlerStampsSoiledResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *HandlerStampsSoiledResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *HandlerStampsSoiledResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *HandlerStampsSoiledResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerStampsSoiledResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerStampsSoiledResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerStampsSoiledResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HandlerStampsSoiledResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseStampsSoiledResponse

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
**StampDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TranId** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStampsSoiledResponse

`func NewResponseStampsSoiledResponse() *ResponseStampsSoiledResponse`

NewResponseStampsSoiledResponse instantiates a new ResponseStampsSoiledResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsSoiledResponseWithDefaults

`func NewResponseStampsSoiledResponseWithDefaults() *ResponseStampsSoiledResponse`

NewResponseStampsSoiledResponseWithDefaults instantiates a new ResponseStampsSoiledResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ResponseStampsSoiledResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ResponseStampsSoiledResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ResponseStampsSoiledResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ResponseStampsSoiledResponse) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetApproverDate

`func (o *ResponseStampsSoiledResponse) GetApproverDate() string`

GetApproverDate returns the ApproverDate field if non-nil, zero value otherwise.

### GetApproverDateOk

`func (o *ResponseStampsSoiledResponse) GetApproverDateOk() (*string, bool)`

GetApproverDateOk returns a tuple with the ApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverDate

`func (o *ResponseStampsSoiledResponse) SetApproverDate(v string)`

SetApproverDate sets ApproverDate field to given value.

### HasApproverDate

`func (o *ResponseStampsSoiledResponse) HasApproverDate() bool`

HasApproverDate returns a boolean if a field has been set.

### GetApproverId

`func (o *ResponseStampsSoiledResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ResponseStampsSoiledResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ResponseStampsSoiledResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ResponseStampsSoiledResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetDisposalDetails

`func (o *ResponseStampsSoiledResponse) GetDisposalDetails() string`

GetDisposalDetails returns the DisposalDetails field if non-nil, zero value otherwise.

### GetDisposalDetailsOk

`func (o *ResponseStampsSoiledResponse) GetDisposalDetailsOk() (*string, bool)`

GetDisposalDetailsOk returns a tuple with the DisposalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisposalDetails

`func (o *ResponseStampsSoiledResponse) SetDisposalDetails(v string)`

SetDisposalDetails sets DisposalDetails field to given value.

### HasDisposalDetails

`func (o *ResponseStampsSoiledResponse) HasDisposalDetails() bool`

HasDisposalDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampsSoiledResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampsSoiledResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampsSoiledResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampsSoiledResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseStampsSoiledResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseStampsSoiledResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseStampsSoiledResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseStampsSoiledResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsSoiledResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsSoiledResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsSoiledResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsSoiledResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSoiledDetails

`func (o *ResponseStampsSoiledResponse) GetSoiledDetails() map[string]int32`

GetSoiledDetails returns the SoiledDetails field if non-nil, zero value otherwise.

### GetSoiledDetailsOk

`func (o *ResponseStampsSoiledResponse) GetSoiledDetailsOk() (*map[string]int32, bool)`

GetSoiledDetailsOk returns a tuple with the SoiledDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoiledDetails

`func (o *ResponseStampsSoiledResponse) SetSoiledDetails(v map[string]int32)`

SetSoiledDetails sets SoiledDetails field to given value.

### HasSoiledDetails

`func (o *ResponseStampsSoiledResponse) HasSoiledDetails() bool`

HasSoiledDetails returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *ResponseStampsSoiledResponse) GetStampDetailsDesc() []ResponseStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *ResponseStampsSoiledResponse) GetStampDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *ResponseStampsSoiledResponse) SetStampDetailsDesc(v []ResponseStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *ResponseStampsSoiledResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTranId

`func (o *ResponseStampsSoiledResponse) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *ResponseStampsSoiledResponse) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *ResponseStampsSoiledResponse) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *ResponseStampsSoiledResponse) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsSoiledResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsSoiledResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsSoiledResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsSoiledResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseStampsSoiledResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseStampsSoiledResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseStampsSoiledResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseStampsSoiledResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



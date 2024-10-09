# ResponseStampsAdvanceTxnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvAmt** | Pointer to **float32** |  | [optional] 
**AdvTxnDetails** | Pointer to **map[string]int32** |  | [optional] 
**AdvUserId** | Pointer to **int32** |  | [optional] 
**ApprovedDate** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**StampDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TranId** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransType** | Pointer to **string** |  | [optional] 
**TryApproverId** | Pointer to **int32** |  | [optional] 
**TryUserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStampsAdvanceTxnResponse

`func NewResponseStampsAdvanceTxnResponse() *ResponseStampsAdvanceTxnResponse`

NewResponseStampsAdvanceTxnResponse instantiates a new ResponseStampsAdvanceTxnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsAdvanceTxnResponseWithDefaults

`func NewResponseStampsAdvanceTxnResponseWithDefaults() *ResponseStampsAdvanceTxnResponse`

NewResponseStampsAdvanceTxnResponseWithDefaults instantiates a new ResponseStampsAdvanceTxnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvAmt

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvAmt() float32`

GetAdvAmt returns the AdvAmt field if non-nil, zero value otherwise.

### GetAdvAmtOk

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvAmtOk() (*float32, bool)`

GetAdvAmtOk returns a tuple with the AdvAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvAmt

`func (o *ResponseStampsAdvanceTxnResponse) SetAdvAmt(v float32)`

SetAdvAmt sets AdvAmt field to given value.

### HasAdvAmt

`func (o *ResponseStampsAdvanceTxnResponse) HasAdvAmt() bool`

HasAdvAmt returns a boolean if a field has been set.

### GetAdvTxnDetails

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvTxnDetails() map[string]int32`

GetAdvTxnDetails returns the AdvTxnDetails field if non-nil, zero value otherwise.

### GetAdvTxnDetailsOk

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvTxnDetailsOk() (*map[string]int32, bool)`

GetAdvTxnDetailsOk returns a tuple with the AdvTxnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvTxnDetails

`func (o *ResponseStampsAdvanceTxnResponse) SetAdvTxnDetails(v map[string]int32)`

SetAdvTxnDetails sets AdvTxnDetails field to given value.

### HasAdvTxnDetails

`func (o *ResponseStampsAdvanceTxnResponse) HasAdvTxnDetails() bool`

HasAdvTxnDetails returns a boolean if a field has been set.

### GetAdvUserId

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvUserId() int32`

GetAdvUserId returns the AdvUserId field if non-nil, zero value otherwise.

### GetAdvUserIdOk

`func (o *ResponseStampsAdvanceTxnResponse) GetAdvUserIdOk() (*int32, bool)`

GetAdvUserIdOk returns a tuple with the AdvUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvUserId

`func (o *ResponseStampsAdvanceTxnResponse) SetAdvUserId(v int32)`

SetAdvUserId sets AdvUserId field to given value.

### HasAdvUserId

`func (o *ResponseStampsAdvanceTxnResponse) HasAdvUserId() bool`

HasAdvUserId returns a boolean if a field has been set.

### GetApprovedDate

`func (o *ResponseStampsAdvanceTxnResponse) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ResponseStampsAdvanceTxnResponse) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ResponseStampsAdvanceTxnResponse) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ResponseStampsAdvanceTxnResponse) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampsAdvanceTxnResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampsAdvanceTxnResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampsAdvanceTxnResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampsAdvanceTxnResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseStampsAdvanceTxnResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseStampsAdvanceTxnResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseStampsAdvanceTxnResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseStampsAdvanceTxnResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsAdvanceTxnResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsAdvanceTxnResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsAdvanceTxnResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsAdvanceTxnResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *ResponseStampsAdvanceTxnResponse) GetStampDetailsDesc() []ResponseStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *ResponseStampsAdvanceTxnResponse) GetStampDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *ResponseStampsAdvanceTxnResponse) SetStampDetailsDesc(v []ResponseStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *ResponseStampsAdvanceTxnResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTranId

`func (o *ResponseStampsAdvanceTxnResponse) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *ResponseStampsAdvanceTxnResponse) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *ResponseStampsAdvanceTxnResponse) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *ResponseStampsAdvanceTxnResponse) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsAdvanceTxnResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsAdvanceTxnResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsAdvanceTxnResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsAdvanceTxnResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransType

`func (o *ResponseStampsAdvanceTxnResponse) GetTransType() string`

GetTransType returns the TransType field if non-nil, zero value otherwise.

### GetTransTypeOk

`func (o *ResponseStampsAdvanceTxnResponse) GetTransTypeOk() (*string, bool)`

GetTransTypeOk returns a tuple with the TransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransType

`func (o *ResponseStampsAdvanceTxnResponse) SetTransType(v string)`

SetTransType sets TransType field to given value.

### HasTransType

`func (o *ResponseStampsAdvanceTxnResponse) HasTransType() bool`

HasTransType returns a boolean if a field has been set.

### GetTryApproverId

`func (o *ResponseStampsAdvanceTxnResponse) GetTryApproverId() int32`

GetTryApproverId returns the TryApproverId field if non-nil, zero value otherwise.

### GetTryApproverIdOk

`func (o *ResponseStampsAdvanceTxnResponse) GetTryApproverIdOk() (*int32, bool)`

GetTryApproverIdOk returns a tuple with the TryApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryApproverId

`func (o *ResponseStampsAdvanceTxnResponse) SetTryApproverId(v int32)`

SetTryApproverId sets TryApproverId field to given value.

### HasTryApproverId

`func (o *ResponseStampsAdvanceTxnResponse) HasTryApproverId() bool`

HasTryApproverId returns a boolean if a field has been set.

### GetTryUserId

`func (o *ResponseStampsAdvanceTxnResponse) GetTryUserId() int32`

GetTryUserId returns the TryUserId field if non-nil, zero value otherwise.

### GetTryUserIdOk

`func (o *ResponseStampsAdvanceTxnResponse) GetTryUserIdOk() (*int32, bool)`

GetTryUserIdOk returns a tuple with the TryUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryUserId

`func (o *ResponseStampsAdvanceTxnResponse) SetTryUserId(v int32)`

SetTryUserId sets TryUserId field to given value.

### HasTryUserId

`func (o *ResponseStampsAdvanceTxnResponse) HasTryUserId() bool`

HasTryUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseTreasuryMiscTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDescdetails** | Pointer to [**[]ResponseAccountDescdetails**](ResponseAccountDescdetails.md) |  | [optional] 
**ApproverDate** | Pointer to **string** |  | [optional] 
**ApproverId** | Pointer to **int32** |  | [optional] 
**ApproverRemarks** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | Pointer to **float32** |  | [optional] 
**ReqDetails** | Pointer to **map[string]float32** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TxnMode** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseTreasuryMiscTransactionResponse

`func NewResponseTreasuryMiscTransactionResponse() *ResponseTreasuryMiscTransactionResponse`

NewResponseTreasuryMiscTransactionResponse instantiates a new ResponseTreasuryMiscTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTreasuryMiscTransactionResponseWithDefaults

`func NewResponseTreasuryMiscTransactionResponseWithDefaults() *ResponseTreasuryMiscTransactionResponse`

NewResponseTreasuryMiscTransactionResponseWithDefaults instantiates a new ResponseTreasuryMiscTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountDescdetails

`func (o *ResponseTreasuryMiscTransactionResponse) GetAccountDescdetails() []ResponseAccountDescdetails`

GetAccountDescdetails returns the AccountDescdetails field if non-nil, zero value otherwise.

### GetAccountDescdetailsOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetAccountDescdetailsOk() (*[]ResponseAccountDescdetails, bool)`

GetAccountDescdetailsOk returns a tuple with the AccountDescdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDescdetails

`func (o *ResponseTreasuryMiscTransactionResponse) SetAccountDescdetails(v []ResponseAccountDescdetails)`

SetAccountDescdetails sets AccountDescdetails field to given value.

### HasAccountDescdetails

`func (o *ResponseTreasuryMiscTransactionResponse) HasAccountDescdetails() bool`

HasAccountDescdetails returns a boolean if a field has been set.

### GetApproverDate

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverDate() string`

GetApproverDate returns the ApproverDate field if non-nil, zero value otherwise.

### GetApproverDateOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverDateOk() (*string, bool)`

GetApproverDateOk returns a tuple with the ApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverDate

`func (o *ResponseTreasuryMiscTransactionResponse) SetApproverDate(v string)`

SetApproverDate sets ApproverDate field to given value.

### HasApproverDate

`func (o *ResponseTreasuryMiscTransactionResponse) HasApproverDate() bool`

HasApproverDate returns a boolean if a field has been set.

### GetApproverId

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ResponseTreasuryMiscTransactionResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ResponseTreasuryMiscTransactionResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetApproverRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.

### HasApproverRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) HasApproverRemarks() bool`

HasApproverRemarks returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseTreasuryMiscTransactionResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseTreasuryMiscTransactionResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseTreasuryMiscTransactionResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseTreasuryMiscTransactionResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseTreasuryMiscTransactionResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseTreasuryMiscTransactionResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseTreasuryMiscTransactionResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseTreasuryMiscTransactionResponse) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseTreasuryMiscTransactionResponse) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseTreasuryMiscTransactionResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqDetails

`func (o *ResponseTreasuryMiscTransactionResponse) GetReqDetails() map[string]float32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetReqDetailsOk() (*map[string]float32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *ResponseTreasuryMiscTransactionResponse) SetReqDetails(v map[string]float32)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *ResponseTreasuryMiscTransactionResponse) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseTreasuryMiscTransactionResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseTreasuryMiscTransactionResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseTreasuryMiscTransactionResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseTreasuryMiscTransactionResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseTreasuryMiscTransactionResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseTreasuryMiscTransactionResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseTreasuryMiscTransactionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseTreasuryMiscTransactionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseTreasuryMiscTransactionResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnMode

`func (o *ResponseTreasuryMiscTransactionResponse) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *ResponseTreasuryMiscTransactionResponse) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.

### HasTxnMode

`func (o *ResponseTreasuryMiscTransactionResponse) HasTxnMode() bool`

HasTxnMode returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseTreasuryMiscTransactionResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseTreasuryMiscTransactionResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseTreasuryMiscTransactionResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseTreasuryMiscTransactionResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



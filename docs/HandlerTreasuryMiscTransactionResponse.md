# HandlerTreasuryMiscTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountDescdetails** | Pointer to [**[]HandlerAccountDescdetails**](HandlerAccountDescdetails.md) |  | [optional] 
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

### NewHandlerTreasuryMiscTransactionResponse

`func NewHandlerTreasuryMiscTransactionResponse() *HandlerTreasuryMiscTransactionResponse`

NewHandlerTreasuryMiscTransactionResponse instantiates a new HandlerTreasuryMiscTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerTreasuryMiscTransactionResponseWithDefaults

`func NewHandlerTreasuryMiscTransactionResponseWithDefaults() *HandlerTreasuryMiscTransactionResponse`

NewHandlerTreasuryMiscTransactionResponseWithDefaults instantiates a new HandlerTreasuryMiscTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountDescdetails

`func (o *HandlerTreasuryMiscTransactionResponse) GetAccountDescdetails() []HandlerAccountDescdetails`

GetAccountDescdetails returns the AccountDescdetails field if non-nil, zero value otherwise.

### GetAccountDescdetailsOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetAccountDescdetailsOk() (*[]HandlerAccountDescdetails, bool)`

GetAccountDescdetailsOk returns a tuple with the AccountDescdetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDescdetails

`func (o *HandlerTreasuryMiscTransactionResponse) SetAccountDescdetails(v []HandlerAccountDescdetails)`

SetAccountDescdetails sets AccountDescdetails field to given value.

### HasAccountDescdetails

`func (o *HandlerTreasuryMiscTransactionResponse) HasAccountDescdetails() bool`

HasAccountDescdetails returns a boolean if a field has been set.

### GetApproverDate

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverDate() string`

GetApproverDate returns the ApproverDate field if non-nil, zero value otherwise.

### GetApproverDateOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverDateOk() (*string, bool)`

GetApproverDateOk returns a tuple with the ApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverDate

`func (o *HandlerTreasuryMiscTransactionResponse) SetApproverDate(v string)`

SetApproverDate sets ApproverDate field to given value.

### HasApproverDate

`func (o *HandlerTreasuryMiscTransactionResponse) HasApproverDate() bool`

HasApproverDate returns a boolean if a field has been set.

### GetApproverId

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerTreasuryMiscTransactionResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *HandlerTreasuryMiscTransactionResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetApproverRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.

### HasApproverRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) HasApproverRemarks() bool`

HasApproverRemarks returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerTreasuryMiscTransactionResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerTreasuryMiscTransactionResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerTreasuryMiscTransactionResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerTreasuryMiscTransactionResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerTreasuryMiscTransactionResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *HandlerTreasuryMiscTransactionResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerTreasuryMiscTransactionResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *HandlerTreasuryMiscTransactionResponse) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *HandlerTreasuryMiscTransactionResponse) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *HandlerTreasuryMiscTransactionResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqDetails

`func (o *HandlerTreasuryMiscTransactionResponse) GetReqDetails() map[string]float32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetReqDetailsOk() (*map[string]float32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *HandlerTreasuryMiscTransactionResponse) SetReqDetails(v map[string]float32)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *HandlerTreasuryMiscTransactionResponse) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetRequestType

`func (o *HandlerTreasuryMiscTransactionResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerTreasuryMiscTransactionResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *HandlerTreasuryMiscTransactionResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTransDate

`func (o *HandlerTreasuryMiscTransactionResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *HandlerTreasuryMiscTransactionResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *HandlerTreasuryMiscTransactionResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *HandlerTreasuryMiscTransactionResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *HandlerTreasuryMiscTransactionResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *HandlerTreasuryMiscTransactionResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnMode

`func (o *HandlerTreasuryMiscTransactionResponse) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *HandlerTreasuryMiscTransactionResponse) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.

### HasTxnMode

`func (o *HandlerTreasuryMiscTransactionResponse) HasTxnMode() bool`

HasTxnMode returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerTreasuryMiscTransactionResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerTreasuryMiscTransactionResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerTreasuryMiscTransactionResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HandlerTreasuryMiscTransactionResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



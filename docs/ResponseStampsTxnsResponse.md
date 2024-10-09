# ResponseStampsTxnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAmount** | Pointer to **float32** |  | [optional] 
**AckApproverRemarks** | Pointer to **string** |  | [optional] 
**AckDate** | Pointer to **string** |  | [optional] 
**AckUserId** | Pointer to **int32** |  | [optional] 
**ApprovedAmount** | Pointer to **float32** |  | [optional] 
**ApprovedDetails** | Pointer to **map[string]int32** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssApproverDate** | Pointer to **string** |  | [optional] 
**IssApproverId** | Pointer to **int32** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**IssUserId** | Pointer to **int32** |  | [optional] 
**IssUserProcessedDate** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | Pointer to **float32** |  | [optional] 
**ReqApproverAmt** | Pointer to **float32** |  | [optional] 
**ReqApproverDate** | Pointer to **string** |  | [optional] 
**ReqApproverId** | Pointer to **int32** |  | [optional] 
**ReqApproverRemarks** | Pointer to **string** |  | [optional] 
**ReqDetails** | Pointer to **map[string]int32** |  | [optional] 
**ReqOfficeId** | Pointer to **int32** |  | [optional] 
**ReqUserId** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**StampDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseStampsTxnsResponse

`func NewResponseStampsTxnsResponse() *ResponseStampsTxnsResponse`

NewResponseStampsTxnsResponse instantiates a new ResponseStampsTxnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsTxnsResponseWithDefaults

`func NewResponseStampsTxnsResponseWithDefaults() *ResponseStampsTxnsResponse`

NewResponseStampsTxnsResponseWithDefaults instantiates a new ResponseStampsTxnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseStampsTxnsResponse) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseStampsTxnsResponse) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseStampsTxnsResponse) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseStampsTxnsResponse) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseStampsTxnsResponse) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseStampsTxnsResponse) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseStampsTxnsResponse) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseStampsTxnsResponse) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseStampsTxnsResponse) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseStampsTxnsResponse) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseStampsTxnsResponse) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseStampsTxnsResponse) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseStampsTxnsResponse) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseStampsTxnsResponse) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseStampsTxnsResponse) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseStampsTxnsResponse) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseStampsTxnsResponse) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseStampsTxnsResponse) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseStampsTxnsResponse) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseStampsTxnsResponse) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseStampsTxnsResponse) GetApprovedDetails() map[string]int32`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseStampsTxnsResponse) GetApprovedDetailsOk() (*map[string]int32, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseStampsTxnsResponse) SetApprovedDetails(v map[string]int32)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseStampsTxnsResponse) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseStampsTxnsResponse) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseStampsTxnsResponse) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseStampsTxnsResponse) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseStampsTxnsResponse) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseStampsTxnsResponse) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseStampsTxnsResponse) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseStampsTxnsResponse) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseStampsTxnsResponse) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseStampsTxnsResponse) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseStampsTxnsResponse) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseStampsTxnsResponse) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseStampsTxnsResponse) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseStampsTxnsResponse) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseStampsTxnsResponse) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseStampsTxnsResponse) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseStampsTxnsResponse) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseStampsTxnsResponse) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseStampsTxnsResponse) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseStampsTxnsResponse) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseStampsTxnsResponse) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseStampsTxnsResponse) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseStampsTxnsResponse) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseStampsTxnsResponse) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseStampsTxnsResponse) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseStampsTxnsResponse) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseStampsTxnsResponse) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseStampsTxnsResponse) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseStampsTxnsResponse) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsTxnsResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsTxnsResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsTxnsResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsTxnsResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseStampsTxnsResponse) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseStampsTxnsResponse) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseStampsTxnsResponse) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseStampsTxnsResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseStampsTxnsResponse) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseStampsTxnsResponse) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseStampsTxnsResponse) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseStampsTxnsResponse) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseStampsTxnsResponse) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseStampsTxnsResponse) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseStampsTxnsResponse) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseStampsTxnsResponse) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseStampsTxnsResponse) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseStampsTxnsResponse) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseStampsTxnsResponse) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseStampsTxnsResponse) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseStampsTxnsResponse) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseStampsTxnsResponse) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseStampsTxnsResponse) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseStampsTxnsResponse) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqDetails

`func (o *ResponseStampsTxnsResponse) GetReqDetails() map[string]int32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *ResponseStampsTxnsResponse) GetReqDetailsOk() (*map[string]int32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *ResponseStampsTxnsResponse) SetReqDetails(v map[string]int32)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *ResponseStampsTxnsResponse) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseStampsTxnsResponse) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseStampsTxnsResponse) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseStampsTxnsResponse) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseStampsTxnsResponse) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseStampsTxnsResponse) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseStampsTxnsResponse) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseStampsTxnsResponse) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseStampsTxnsResponse) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseStampsTxnsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseStampsTxnsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseStampsTxnsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseStampsTxnsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseStampsTxnsResponse) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseStampsTxnsResponse) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseStampsTxnsResponse) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseStampsTxnsResponse) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseStampsTxnsResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseStampsTxnsResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseStampsTxnsResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseStampsTxnsResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *ResponseStampsTxnsResponse) GetStampDetailsDesc() []ResponseStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *ResponseStampsTxnsResponse) GetStampDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *ResponseStampsTxnsResponse) SetStampDetailsDesc(v []ResponseStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *ResponseStampsTxnsResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsTxnsResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsTxnsResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsTxnsResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsTxnsResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseStampsTxnsResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseStampsTxnsResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseStampsTxnsResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseStampsTxnsResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



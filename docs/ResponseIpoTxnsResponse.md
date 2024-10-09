# ResponseIpoTxnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAmount** | Pointer to **float32** |  | [optional] 
**AckApproverRemarks** | Pointer to **string** |  | [optional] 
**AckDate** | Pointer to **string** |  | [optional] 
**AckUserId** | Pointer to **int32** |  | [optional] 
**ApprovedAmount** | Pointer to **float32** |  | [optional] 
**ApprovedDetails** | Pointer to **string** |  | [optional] 
**IpodetailsDesc** | Pointer to [**[]ResponseIPOdetails**](ResponseIPOdetails.md) |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssApproverDate** | Pointer to **string** |  | [optional] 
**IssApproverId** | Pointer to **int32** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**IssOfficeName** | Pointer to **string** |  | [optional] 
**IssUserId** | Pointer to **int32** |  | [optional] 
**IssUserProcessedDate** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | Pointer to **float32** |  | [optional] 
**ReqApproverAmt** | Pointer to **float32** |  | [optional] 
**ReqApproverDate** | Pointer to **string** |  | [optional] 
**ReqApproverId** | Pointer to **int32** |  | [optional] 
**ReqApproverRemarks** | Pointer to **string** |  | [optional] 
**ReqDetails** | Pointer to **string** |  | [optional] 
**ReqOfficeId** | Pointer to **int32** |  | [optional] 
**ReqOfficeName** | Pointer to **string** |  | [optional] 
**ReqUserId** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TxnStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseIpoTxnsResponse

`func NewResponseIpoTxnsResponse() *ResponseIpoTxnsResponse`

NewResponseIpoTxnsResponse instantiates a new ResponseIpoTxnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseIpoTxnsResponseWithDefaults

`func NewResponseIpoTxnsResponseWithDefaults() *ResponseIpoTxnsResponse`

NewResponseIpoTxnsResponseWithDefaults instantiates a new ResponseIpoTxnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseIpoTxnsResponse) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseIpoTxnsResponse) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseIpoTxnsResponse) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseIpoTxnsResponse) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseIpoTxnsResponse) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseIpoTxnsResponse) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseIpoTxnsResponse) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseIpoTxnsResponse) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseIpoTxnsResponse) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseIpoTxnsResponse) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseIpoTxnsResponse) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseIpoTxnsResponse) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseIpoTxnsResponse) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseIpoTxnsResponse) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseIpoTxnsResponse) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseIpoTxnsResponse) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseIpoTxnsResponse) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseIpoTxnsResponse) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseIpoTxnsResponse) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseIpoTxnsResponse) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseIpoTxnsResponse) GetApprovedDetails() string`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseIpoTxnsResponse) GetApprovedDetailsOk() (*string, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseIpoTxnsResponse) SetApprovedDetails(v string)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseIpoTxnsResponse) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetIpodetailsDesc

`func (o *ResponseIpoTxnsResponse) GetIpodetailsDesc() []ResponseIPOdetails`

GetIpodetailsDesc returns the IpodetailsDesc field if non-nil, zero value otherwise.

### GetIpodetailsDescOk

`func (o *ResponseIpoTxnsResponse) GetIpodetailsDescOk() (*[]ResponseIPOdetails, bool)`

GetIpodetailsDescOk returns a tuple with the IpodetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpodetailsDesc

`func (o *ResponseIpoTxnsResponse) SetIpodetailsDesc(v []ResponseIPOdetails)`

SetIpodetailsDesc sets IpodetailsDesc field to given value.

### HasIpodetailsDesc

`func (o *ResponseIpoTxnsResponse) HasIpodetailsDesc() bool`

HasIpodetailsDesc returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseIpoTxnsResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseIpoTxnsResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseIpoTxnsResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseIpoTxnsResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseIpoTxnsResponse) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseIpoTxnsResponse) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseIpoTxnsResponse) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseIpoTxnsResponse) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseIpoTxnsResponse) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseIpoTxnsResponse) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseIpoTxnsResponse) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseIpoTxnsResponse) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseIpoTxnsResponse) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseIpoTxnsResponse) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseIpoTxnsResponse) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseIpoTxnsResponse) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseIpoTxnsResponse) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseIpoTxnsResponse) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseIpoTxnsResponse) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseIpoTxnsResponse) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseIpoTxnsResponse) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseIpoTxnsResponse) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseIpoTxnsResponse) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseIpoTxnsResponse) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssOfficeName

`func (o *ResponseIpoTxnsResponse) GetIssOfficeName() string`

GetIssOfficeName returns the IssOfficeName field if non-nil, zero value otherwise.

### GetIssOfficeNameOk

`func (o *ResponseIpoTxnsResponse) GetIssOfficeNameOk() (*string, bool)`

GetIssOfficeNameOk returns a tuple with the IssOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeName

`func (o *ResponseIpoTxnsResponse) SetIssOfficeName(v string)`

SetIssOfficeName sets IssOfficeName field to given value.

### HasIssOfficeName

`func (o *ResponseIpoTxnsResponse) HasIssOfficeName() bool`

HasIssOfficeName returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseIpoTxnsResponse) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseIpoTxnsResponse) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseIpoTxnsResponse) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseIpoTxnsResponse) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseIpoTxnsResponse) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseIpoTxnsResponse) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseIpoTxnsResponse) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseIpoTxnsResponse) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseIpoTxnsResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseIpoTxnsResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseIpoTxnsResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseIpoTxnsResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseIpoTxnsResponse) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseIpoTxnsResponse) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseIpoTxnsResponse) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseIpoTxnsResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseIpoTxnsResponse) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseIpoTxnsResponse) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseIpoTxnsResponse) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseIpoTxnsResponse) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseIpoTxnsResponse) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseIpoTxnsResponse) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseIpoTxnsResponse) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseIpoTxnsResponse) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseIpoTxnsResponse) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseIpoTxnsResponse) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseIpoTxnsResponse) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseIpoTxnsResponse) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseIpoTxnsResponse) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseIpoTxnsResponse) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseIpoTxnsResponse) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseIpoTxnsResponse) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqDetails

`func (o *ResponseIpoTxnsResponse) GetReqDetails() string`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *ResponseIpoTxnsResponse) GetReqDetailsOk() (*string, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *ResponseIpoTxnsResponse) SetReqDetails(v string)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *ResponseIpoTxnsResponse) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseIpoTxnsResponse) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseIpoTxnsResponse) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseIpoTxnsResponse) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseIpoTxnsResponse) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqOfficeName

`func (o *ResponseIpoTxnsResponse) GetReqOfficeName() string`

GetReqOfficeName returns the ReqOfficeName field if non-nil, zero value otherwise.

### GetReqOfficeNameOk

`func (o *ResponseIpoTxnsResponse) GetReqOfficeNameOk() (*string, bool)`

GetReqOfficeNameOk returns a tuple with the ReqOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeName

`func (o *ResponseIpoTxnsResponse) SetReqOfficeName(v string)`

SetReqOfficeName sets ReqOfficeName field to given value.

### HasReqOfficeName

`func (o *ResponseIpoTxnsResponse) HasReqOfficeName() bool`

HasReqOfficeName returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseIpoTxnsResponse) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseIpoTxnsResponse) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseIpoTxnsResponse) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseIpoTxnsResponse) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseIpoTxnsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseIpoTxnsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseIpoTxnsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseIpoTxnsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseIpoTxnsResponse) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseIpoTxnsResponse) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseIpoTxnsResponse) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseIpoTxnsResponse) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseIpoTxnsResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseIpoTxnsResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseIpoTxnsResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseIpoTxnsResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseIpoTxnsResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseIpoTxnsResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseIpoTxnsResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseIpoTxnsResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseIpoTxnsResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseIpoTxnsResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseIpoTxnsResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseIpoTxnsResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnStatus

`func (o *ResponseIpoTxnsResponse) GetTxnStatus() string`

GetTxnStatus returns the TxnStatus field if non-nil, zero value otherwise.

### GetTxnStatusOk

`func (o *ResponseIpoTxnsResponse) GetTxnStatusOk() (*string, bool)`

GetTxnStatusOk returns a tuple with the TxnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnStatus

`func (o *ResponseIpoTxnsResponse) SetTxnStatus(v string)`

SetTxnStatus sets TxnStatus field to given value.

### HasTxnStatus

`func (o *ResponseIpoTxnsResponse) HasTxnStatus() bool`

HasTxnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



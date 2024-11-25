# ResponseStampsTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAmount** | Pointer to **float32** |  | [optional] 
**AckApproverRemarks** | Pointer to **string** |  | [optional] 
**AckDate** | Pointer to **string** |  | [optional] 
**AckDetails** | Pointer to **map[string]int32** |  | [optional] 
**AckUserId** | Pointer to **int32** |  | [optional] 
**ApprovedAmount** | Pointer to **float32** |  | [optional] 
**ApprovedDetails** | Pointer to **map[string]int32** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssApproverDate** | Pointer to **string** |  | [optional] 
**IssApproverId** | Pointer to **int32** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**IssOfficeName** | Pointer to **string** |  | [optional] 
**IssUserId** | Pointer to **int32** |  | [optional] 
**IssUserName** | Pointer to **string** |  | [optional] 
**IssUserProcessedDate** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | Pointer to **float32** |  | [optional] 
**ReqApproverAmt** | Pointer to **float32** |  | [optional] 
**ReqApproverDate** | Pointer to **string** |  | [optional] 
**ReqApproverId** | Pointer to **int32** |  | [optional] 
**ReqApproverRemarks** | Pointer to **string** |  | [optional] 
**ReqDetails** | Pointer to **map[string]int32** |  | [optional] 
**ReqOfficeId** | Pointer to **int32** |  | [optional] 
**ReqOfficeName** | Pointer to **string** |  | [optional] 
**ReqUserId** | Pointer to **int32** |  | [optional] 
**ReqUserName** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**StampDetails** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**StampIssuedDetails** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TxnStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseStampsTransactions

`func NewResponseStampsTransactions() *ResponseStampsTransactions`

NewResponseStampsTransactions instantiates a new ResponseStampsTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsTransactionsWithDefaults

`func NewResponseStampsTransactionsWithDefaults() *ResponseStampsTransactions`

NewResponseStampsTransactionsWithDefaults instantiates a new ResponseStampsTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseStampsTransactions) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseStampsTransactions) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseStampsTransactions) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseStampsTransactions) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseStampsTransactions) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseStampsTransactions) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseStampsTransactions) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseStampsTransactions) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseStampsTransactions) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseStampsTransactions) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseStampsTransactions) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseStampsTransactions) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckDetails

`func (o *ResponseStampsTransactions) GetAckDetails() map[string]int32`

GetAckDetails returns the AckDetails field if non-nil, zero value otherwise.

### GetAckDetailsOk

`func (o *ResponseStampsTransactions) GetAckDetailsOk() (*map[string]int32, bool)`

GetAckDetailsOk returns a tuple with the AckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDetails

`func (o *ResponseStampsTransactions) SetAckDetails(v map[string]int32)`

SetAckDetails sets AckDetails field to given value.

### HasAckDetails

`func (o *ResponseStampsTransactions) HasAckDetails() bool`

HasAckDetails returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseStampsTransactions) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseStampsTransactions) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseStampsTransactions) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseStampsTransactions) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseStampsTransactions) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseStampsTransactions) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseStampsTransactions) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseStampsTransactions) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseStampsTransactions) GetApprovedDetails() map[string]int32`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseStampsTransactions) GetApprovedDetailsOk() (*map[string]int32, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseStampsTransactions) SetApprovedDetails(v map[string]int32)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseStampsTransactions) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampsTransactions) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampsTransactions) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampsTransactions) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampsTransactions) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *ResponseStampsTransactions) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *ResponseStampsTransactions) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *ResponseStampsTransactions) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *ResponseStampsTransactions) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseStampsTransactions) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseStampsTransactions) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseStampsTransactions) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseStampsTransactions) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseStampsTransactions) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseStampsTransactions) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseStampsTransactions) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseStampsTransactions) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseStampsTransactions) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseStampsTransactions) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseStampsTransactions) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseStampsTransactions) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseStampsTransactions) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseStampsTransactions) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseStampsTransactions) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseStampsTransactions) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseStampsTransactions) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseStampsTransactions) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseStampsTransactions) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseStampsTransactions) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssOfficeName

`func (o *ResponseStampsTransactions) GetIssOfficeName() string`

GetIssOfficeName returns the IssOfficeName field if non-nil, zero value otherwise.

### GetIssOfficeNameOk

`func (o *ResponseStampsTransactions) GetIssOfficeNameOk() (*string, bool)`

GetIssOfficeNameOk returns a tuple with the IssOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeName

`func (o *ResponseStampsTransactions) SetIssOfficeName(v string)`

SetIssOfficeName sets IssOfficeName field to given value.

### HasIssOfficeName

`func (o *ResponseStampsTransactions) HasIssOfficeName() bool`

HasIssOfficeName returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseStampsTransactions) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseStampsTransactions) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseStampsTransactions) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseStampsTransactions) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserName

`func (o *ResponseStampsTransactions) GetIssUserName() string`

GetIssUserName returns the IssUserName field if non-nil, zero value otherwise.

### GetIssUserNameOk

`func (o *ResponseStampsTransactions) GetIssUserNameOk() (*string, bool)`

GetIssUserNameOk returns a tuple with the IssUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserName

`func (o *ResponseStampsTransactions) SetIssUserName(v string)`

SetIssUserName sets IssUserName field to given value.

### HasIssUserName

`func (o *ResponseStampsTransactions) HasIssUserName() bool`

HasIssUserName returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseStampsTransactions) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseStampsTransactions) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseStampsTransactions) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseStampsTransactions) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsTransactions) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsTransactions) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsTransactions) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsTransactions) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseStampsTransactions) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseStampsTransactions) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseStampsTransactions) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseStampsTransactions) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseStampsTransactions) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseStampsTransactions) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseStampsTransactions) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseStampsTransactions) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseStampsTransactions) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseStampsTransactions) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseStampsTransactions) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseStampsTransactions) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseStampsTransactions) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseStampsTransactions) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseStampsTransactions) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseStampsTransactions) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseStampsTransactions) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseStampsTransactions) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseStampsTransactions) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseStampsTransactions) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqDetails

`func (o *ResponseStampsTransactions) GetReqDetails() map[string]int32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *ResponseStampsTransactions) GetReqDetailsOk() (*map[string]int32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *ResponseStampsTransactions) SetReqDetails(v map[string]int32)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *ResponseStampsTransactions) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseStampsTransactions) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseStampsTransactions) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseStampsTransactions) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseStampsTransactions) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqOfficeName

`func (o *ResponseStampsTransactions) GetReqOfficeName() string`

GetReqOfficeName returns the ReqOfficeName field if non-nil, zero value otherwise.

### GetReqOfficeNameOk

`func (o *ResponseStampsTransactions) GetReqOfficeNameOk() (*string, bool)`

GetReqOfficeNameOk returns a tuple with the ReqOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeName

`func (o *ResponseStampsTransactions) SetReqOfficeName(v string)`

SetReqOfficeName sets ReqOfficeName field to given value.

### HasReqOfficeName

`func (o *ResponseStampsTransactions) HasReqOfficeName() bool`

HasReqOfficeName returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseStampsTransactions) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseStampsTransactions) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseStampsTransactions) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseStampsTransactions) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetReqUserName

`func (o *ResponseStampsTransactions) GetReqUserName() string`

GetReqUserName returns the ReqUserName field if non-nil, zero value otherwise.

### GetReqUserNameOk

`func (o *ResponseStampsTransactions) GetReqUserNameOk() (*string, bool)`

GetReqUserNameOk returns a tuple with the ReqUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserName

`func (o *ResponseStampsTransactions) SetReqUserName(v string)`

SetReqUserName sets ReqUserName field to given value.

### HasReqUserName

`func (o *ResponseStampsTransactions) HasReqUserName() bool`

HasReqUserName returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseStampsTransactions) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseStampsTransactions) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseStampsTransactions) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseStampsTransactions) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseStampsTransactions) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseStampsTransactions) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseStampsTransactions) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseStampsTransactions) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseStampsTransactions) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseStampsTransactions) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseStampsTransactions) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseStampsTransactions) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetStampDetails

`func (o *ResponseStampsTransactions) GetStampDetails() []ResponseStampdetails`

GetStampDetails returns the StampDetails field if non-nil, zero value otherwise.

### GetStampDetailsOk

`func (o *ResponseStampsTransactions) GetStampDetailsOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsOk returns a tuple with the StampDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetails

`func (o *ResponseStampsTransactions) SetStampDetails(v []ResponseStampdetails)`

SetStampDetails sets StampDetails field to given value.

### HasStampDetails

`func (o *ResponseStampsTransactions) HasStampDetails() bool`

HasStampDetails returns a boolean if a field has been set.

### GetStampIssuedDetails

`func (o *ResponseStampsTransactions) GetStampIssuedDetails() []ResponseStampdetails`

GetStampIssuedDetails returns the StampIssuedDetails field if non-nil, zero value otherwise.

### GetStampIssuedDetailsOk

`func (o *ResponseStampsTransactions) GetStampIssuedDetailsOk() (*[]ResponseStampdetails, bool)`

GetStampIssuedDetailsOk returns a tuple with the StampIssuedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampIssuedDetails

`func (o *ResponseStampsTransactions) SetStampIssuedDetails(v []ResponseStampdetails)`

SetStampIssuedDetails sets StampIssuedDetails field to given value.

### HasStampIssuedDetails

`func (o *ResponseStampsTransactions) HasStampIssuedDetails() bool`

HasStampIssuedDetails returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsTransactions) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsTransactions) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsTransactions) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsTransactions) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseStampsTransactions) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseStampsTransactions) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseStampsTransactions) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseStampsTransactions) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnStatus

`func (o *ResponseStampsTransactions) GetTxnStatus() string`

GetTxnStatus returns the TxnStatus field if non-nil, zero value otherwise.

### GetTxnStatusOk

`func (o *ResponseStampsTransactions) GetTxnStatusOk() (*string, bool)`

GetTxnStatusOk returns a tuple with the TxnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnStatus

`func (o *ResponseStampsTransactions) SetTxnStatus(v string)`

SetTxnStatus sets TxnStatus field to given value.

### HasTxnStatus

`func (o *ResponseStampsTransactions) HasTxnStatus() bool`

HasTxnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



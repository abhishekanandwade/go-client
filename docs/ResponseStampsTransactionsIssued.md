# ResponseStampsTransactionsIssued

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

### NewResponseStampsTransactionsIssued

`func NewResponseStampsTransactionsIssued() *ResponseStampsTransactionsIssued`

NewResponseStampsTransactionsIssued instantiates a new ResponseStampsTransactionsIssued object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsTransactionsIssuedWithDefaults

`func NewResponseStampsTransactionsIssuedWithDefaults() *ResponseStampsTransactionsIssued`

NewResponseStampsTransactionsIssuedWithDefaults instantiates a new ResponseStampsTransactionsIssued object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseStampsTransactionsIssued) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseStampsTransactionsIssued) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseStampsTransactionsIssued) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseStampsTransactionsIssued) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseStampsTransactionsIssued) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseStampsTransactionsIssued) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseStampsTransactionsIssued) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseStampsTransactionsIssued) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseStampsTransactionsIssued) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseStampsTransactionsIssued) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseStampsTransactionsIssued) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseStampsTransactionsIssued) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckDetails

`func (o *ResponseStampsTransactionsIssued) GetAckDetails() map[string]int32`

GetAckDetails returns the AckDetails field if non-nil, zero value otherwise.

### GetAckDetailsOk

`func (o *ResponseStampsTransactionsIssued) GetAckDetailsOk() (*map[string]int32, bool)`

GetAckDetailsOk returns a tuple with the AckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDetails

`func (o *ResponseStampsTransactionsIssued) SetAckDetails(v map[string]int32)`

SetAckDetails sets AckDetails field to given value.

### HasAckDetails

`func (o *ResponseStampsTransactionsIssued) HasAckDetails() bool`

HasAckDetails returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseStampsTransactionsIssued) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseStampsTransactionsIssued) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseStampsTransactionsIssued) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseStampsTransactionsIssued) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseStampsTransactionsIssued) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseStampsTransactionsIssued) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseStampsTransactionsIssued) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseStampsTransactionsIssued) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseStampsTransactionsIssued) GetApprovedDetails() map[string]int32`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseStampsTransactionsIssued) GetApprovedDetailsOk() (*map[string]int32, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseStampsTransactionsIssued) SetApprovedDetails(v map[string]int32)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseStampsTransactionsIssued) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampsTransactionsIssued) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampsTransactionsIssued) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampsTransactionsIssued) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampsTransactionsIssued) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseStampsTransactionsIssued) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseStampsTransactionsIssued) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseStampsTransactionsIssued) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseStampsTransactionsIssued) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseStampsTransactionsIssued) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseStampsTransactionsIssued) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseStampsTransactionsIssued) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseStampsTransactionsIssued) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseStampsTransactionsIssued) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseStampsTransactionsIssued) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseStampsTransactionsIssued) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseStampsTransactionsIssued) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseStampsTransactionsIssued) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseStampsTransactionsIssued) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseStampsTransactionsIssued) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseStampsTransactionsIssued) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseStampsTransactionsIssued) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseStampsTransactionsIssued) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseStampsTransactionsIssued) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseStampsTransactionsIssued) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssOfficeName

`func (o *ResponseStampsTransactionsIssued) GetIssOfficeName() string`

GetIssOfficeName returns the IssOfficeName field if non-nil, zero value otherwise.

### GetIssOfficeNameOk

`func (o *ResponseStampsTransactionsIssued) GetIssOfficeNameOk() (*string, bool)`

GetIssOfficeNameOk returns a tuple with the IssOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeName

`func (o *ResponseStampsTransactionsIssued) SetIssOfficeName(v string)`

SetIssOfficeName sets IssOfficeName field to given value.

### HasIssOfficeName

`func (o *ResponseStampsTransactionsIssued) HasIssOfficeName() bool`

HasIssOfficeName returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseStampsTransactionsIssued) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseStampsTransactionsIssued) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseStampsTransactionsIssued) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseStampsTransactionsIssued) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserName

`func (o *ResponseStampsTransactionsIssued) GetIssUserName() string`

GetIssUserName returns the IssUserName field if non-nil, zero value otherwise.

### GetIssUserNameOk

`func (o *ResponseStampsTransactionsIssued) GetIssUserNameOk() (*string, bool)`

GetIssUserNameOk returns a tuple with the IssUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserName

`func (o *ResponseStampsTransactionsIssued) SetIssUserName(v string)`

SetIssUserName sets IssUserName field to given value.

### HasIssUserName

`func (o *ResponseStampsTransactionsIssued) HasIssUserName() bool`

HasIssUserName returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseStampsTransactionsIssued) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseStampsTransactionsIssued) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseStampsTransactionsIssued) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseStampsTransactionsIssued) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsTransactionsIssued) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsTransactionsIssued) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsTransactionsIssued) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsTransactionsIssued) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseStampsTransactionsIssued) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseStampsTransactionsIssued) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseStampsTransactionsIssued) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseStampsTransactionsIssued) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseStampsTransactionsIssued) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseStampsTransactionsIssued) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseStampsTransactionsIssued) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseStampsTransactionsIssued) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseStampsTransactionsIssued) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseStampsTransactionsIssued) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseStampsTransactionsIssued) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseStampsTransactionsIssued) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseStampsTransactionsIssued) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseStampsTransactionsIssued) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseStampsTransactionsIssued) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseStampsTransactionsIssued) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseStampsTransactionsIssued) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseStampsTransactionsIssued) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseStampsTransactionsIssued) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseStampsTransactionsIssued) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqDetails

`func (o *ResponseStampsTransactionsIssued) GetReqDetails() map[string]int32`

GetReqDetails returns the ReqDetails field if non-nil, zero value otherwise.

### GetReqDetailsOk

`func (o *ResponseStampsTransactionsIssued) GetReqDetailsOk() (*map[string]int32, bool)`

GetReqDetailsOk returns a tuple with the ReqDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqDetails

`func (o *ResponseStampsTransactionsIssued) SetReqDetails(v map[string]int32)`

SetReqDetails sets ReqDetails field to given value.

### HasReqDetails

`func (o *ResponseStampsTransactionsIssued) HasReqDetails() bool`

HasReqDetails returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseStampsTransactionsIssued) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseStampsTransactionsIssued) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseStampsTransactionsIssued) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseStampsTransactionsIssued) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqOfficeName

`func (o *ResponseStampsTransactionsIssued) GetReqOfficeName() string`

GetReqOfficeName returns the ReqOfficeName field if non-nil, zero value otherwise.

### GetReqOfficeNameOk

`func (o *ResponseStampsTransactionsIssued) GetReqOfficeNameOk() (*string, bool)`

GetReqOfficeNameOk returns a tuple with the ReqOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeName

`func (o *ResponseStampsTransactionsIssued) SetReqOfficeName(v string)`

SetReqOfficeName sets ReqOfficeName field to given value.

### HasReqOfficeName

`func (o *ResponseStampsTransactionsIssued) HasReqOfficeName() bool`

HasReqOfficeName returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseStampsTransactionsIssued) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseStampsTransactionsIssued) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseStampsTransactionsIssued) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseStampsTransactionsIssued) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetReqUserName

`func (o *ResponseStampsTransactionsIssued) GetReqUserName() string`

GetReqUserName returns the ReqUserName field if non-nil, zero value otherwise.

### GetReqUserNameOk

`func (o *ResponseStampsTransactionsIssued) GetReqUserNameOk() (*string, bool)`

GetReqUserNameOk returns a tuple with the ReqUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserName

`func (o *ResponseStampsTransactionsIssued) SetReqUserName(v string)`

SetReqUserName sets ReqUserName field to given value.

### HasReqUserName

`func (o *ResponseStampsTransactionsIssued) HasReqUserName() bool`

HasReqUserName returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseStampsTransactionsIssued) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseStampsTransactionsIssued) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseStampsTransactionsIssued) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseStampsTransactionsIssued) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseStampsTransactionsIssued) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseStampsTransactionsIssued) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseStampsTransactionsIssued) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseStampsTransactionsIssued) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseStampsTransactionsIssued) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseStampsTransactionsIssued) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseStampsTransactionsIssued) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseStampsTransactionsIssued) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetStampDetails

`func (o *ResponseStampsTransactionsIssued) GetStampDetails() []ResponseStampdetails`

GetStampDetails returns the StampDetails field if non-nil, zero value otherwise.

### GetStampDetailsOk

`func (o *ResponseStampsTransactionsIssued) GetStampDetailsOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsOk returns a tuple with the StampDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetails

`func (o *ResponseStampsTransactionsIssued) SetStampDetails(v []ResponseStampdetails)`

SetStampDetails sets StampDetails field to given value.

### HasStampDetails

`func (o *ResponseStampsTransactionsIssued) HasStampDetails() bool`

HasStampDetails returns a boolean if a field has been set.

### GetStampIssuedDetails

`func (o *ResponseStampsTransactionsIssued) GetStampIssuedDetails() []ResponseStampdetails`

GetStampIssuedDetails returns the StampIssuedDetails field if non-nil, zero value otherwise.

### GetStampIssuedDetailsOk

`func (o *ResponseStampsTransactionsIssued) GetStampIssuedDetailsOk() (*[]ResponseStampdetails, bool)`

GetStampIssuedDetailsOk returns a tuple with the StampIssuedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampIssuedDetails

`func (o *ResponseStampsTransactionsIssued) SetStampIssuedDetails(v []ResponseStampdetails)`

SetStampIssuedDetails sets StampIssuedDetails field to given value.

### HasStampIssuedDetails

`func (o *ResponseStampsTransactionsIssued) HasStampIssuedDetails() bool`

HasStampIssuedDetails returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsTransactionsIssued) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsTransactionsIssued) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsTransactionsIssued) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsTransactionsIssued) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseStampsTransactionsIssued) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseStampsTransactionsIssued) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseStampsTransactionsIssued) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseStampsTransactionsIssued) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnStatus

`func (o *ResponseStampsTransactionsIssued) GetTxnStatus() string`

GetTxnStatus returns the TxnStatus field if non-nil, zero value otherwise.

### GetTxnStatusOk

`func (o *ResponseStampsTransactionsIssued) GetTxnStatusOk() (*string, bool)`

GetTxnStatusOk returns a tuple with the TxnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnStatus

`func (o *ResponseStampsTransactionsIssued) SetTxnStatus(v string)`

SetTxnStatus sets TxnStatus field to given value.

### HasTxnStatus

`func (o *ResponseStampsTransactionsIssued) HasTxnStatus() bool`

HasTxnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



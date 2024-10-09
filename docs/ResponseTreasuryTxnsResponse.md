# ResponseTreasuryTxnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAmount** | Pointer to **float32** |  | [optional] 
**AckApproverRemarks** | Pointer to **string** |  | [optional] 
**AckDate** | Pointer to **string** |  | [optional] 
**AckUserId** | Pointer to **int32** |  | [optional] 
**ApprovedAmount** | Pointer to **float32** |  | [optional] 
**ApprovedDetails** | Pointer to **map[string]int32** |  | [optional] 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**CurrencyDetails** | Pointer to **map[string]int32** |  | [optional] 
**EmployeeId1** | Pointer to **int32** |  | [optional] 
**EmployeeId2** | Pointer to **int32** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssApproverDate** | Pointer to **string** |  | [optional] 
**IssApproverId** | Pointer to **int32** |  | [optional] 
**IssApproverRemarks** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**IssOfficeName** | Pointer to **string** |  | [optional] 
**IssUserId** | Pointer to **int32** |  | [optional] 
**IssUserProcessedDate** | Pointer to **string** |  | [optional] 
**LiabilityDetails** | Pointer to **string** |  | [optional] 
**LimitId** | Pointer to **string** |  | [optional] 
**PayeeName** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | Pointer to **float32** |  | [optional] 
**ReqApproverAmt** | Pointer to **float32** |  | [optional] 
**ReqApproverDate** | Pointer to **string** |  | [optional] 
**ReqApproverId** | Pointer to **int32** |  | [optional] 
**ReqApproverRemarks** | Pointer to **string** |  | [optional] 
**ReqOfficeId** | Pointer to **int32** |  | [optional] 
**ReqOfficeName** | Pointer to **string** |  | [optional] 
**ReqUserId** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**TxnMode** | Pointer to **string** |  | [optional] 
**TxnStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseTreasuryTxnsResponse

`func NewResponseTreasuryTxnsResponse() *ResponseTreasuryTxnsResponse`

NewResponseTreasuryTxnsResponse instantiates a new ResponseTreasuryTxnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTreasuryTxnsResponseWithDefaults

`func NewResponseTreasuryTxnsResponseWithDefaults() *ResponseTreasuryTxnsResponse`

NewResponseTreasuryTxnsResponseWithDefaults instantiates a new ResponseTreasuryTxnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseTreasuryTxnsResponse) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseTreasuryTxnsResponse) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseTreasuryTxnsResponse) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseTreasuryTxnsResponse) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseTreasuryTxnsResponse) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseTreasuryTxnsResponse) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseTreasuryTxnsResponse) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseTreasuryTxnsResponse) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseTreasuryTxnsResponse) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseTreasuryTxnsResponse) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseTreasuryTxnsResponse) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseTreasuryTxnsResponse) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseTreasuryTxnsResponse) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseTreasuryTxnsResponse) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseTreasuryTxnsResponse) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseTreasuryTxnsResponse) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseTreasuryTxnsResponse) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseTreasuryTxnsResponse) GetApprovedDetails() map[string]int32`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseTreasuryTxnsResponse) GetApprovedDetailsOk() (*map[string]int32, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseTreasuryTxnsResponse) SetApprovedDetails(v map[string]int32)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseTreasuryTxnsResponse) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetChequeDate

`func (o *ResponseTreasuryTxnsResponse) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *ResponseTreasuryTxnsResponse) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *ResponseTreasuryTxnsResponse) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *ResponseTreasuryTxnsResponse) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *ResponseTreasuryTxnsResponse) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *ResponseTreasuryTxnsResponse) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *ResponseTreasuryTxnsResponse) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *ResponseTreasuryTxnsResponse) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetCurrencyDetails

`func (o *ResponseTreasuryTxnsResponse) GetCurrencyDetails() map[string]int32`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *ResponseTreasuryTxnsResponse) GetCurrencyDetailsOk() (*map[string]int32, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *ResponseTreasuryTxnsResponse) SetCurrencyDetails(v map[string]int32)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *ResponseTreasuryTxnsResponse) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetEmployeeId1

`func (o *ResponseTreasuryTxnsResponse) GetEmployeeId1() int32`

GetEmployeeId1 returns the EmployeeId1 field if non-nil, zero value otherwise.

### GetEmployeeId1Ok

`func (o *ResponseTreasuryTxnsResponse) GetEmployeeId1Ok() (*int32, bool)`

GetEmployeeId1Ok returns a tuple with the EmployeeId1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId1

`func (o *ResponseTreasuryTxnsResponse) SetEmployeeId1(v int32)`

SetEmployeeId1 sets EmployeeId1 field to given value.

### HasEmployeeId1

`func (o *ResponseTreasuryTxnsResponse) HasEmployeeId1() bool`

HasEmployeeId1 returns a boolean if a field has been set.

### GetEmployeeId2

`func (o *ResponseTreasuryTxnsResponse) GetEmployeeId2() int32`

GetEmployeeId2 returns the EmployeeId2 field if non-nil, zero value otherwise.

### GetEmployeeId2Ok

`func (o *ResponseTreasuryTxnsResponse) GetEmployeeId2Ok() (*int32, bool)`

GetEmployeeId2Ok returns a tuple with the EmployeeId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId2

`func (o *ResponseTreasuryTxnsResponse) SetEmployeeId2(v int32)`

SetEmployeeId2 sets EmployeeId2 field to given value.

### HasEmployeeId2

`func (o *ResponseTreasuryTxnsResponse) HasEmployeeId2() bool`

HasEmployeeId2 returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseTreasuryTxnsResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseTreasuryTxnsResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseTreasuryTxnsResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseTreasuryTxnsResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseTreasuryTxnsResponse) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseTreasuryTxnsResponse) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseTreasuryTxnsResponse) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseTreasuryTxnsResponse) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseTreasuryTxnsResponse) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseTreasuryTxnsResponse) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseTreasuryTxnsResponse) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseTreasuryTxnsResponse) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseTreasuryTxnsResponse) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseTreasuryTxnsResponse) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseTreasuryTxnsResponse) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseTreasuryTxnsResponse) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseTreasuryTxnsResponse) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssOfficeName

`func (o *ResponseTreasuryTxnsResponse) GetIssOfficeName() string`

GetIssOfficeName returns the IssOfficeName field if non-nil, zero value otherwise.

### GetIssOfficeNameOk

`func (o *ResponseTreasuryTxnsResponse) GetIssOfficeNameOk() (*string, bool)`

GetIssOfficeNameOk returns a tuple with the IssOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeName

`func (o *ResponseTreasuryTxnsResponse) SetIssOfficeName(v string)`

SetIssOfficeName sets IssOfficeName field to given value.

### HasIssOfficeName

`func (o *ResponseTreasuryTxnsResponse) HasIssOfficeName() bool`

HasIssOfficeName returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseTreasuryTxnsResponse) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseTreasuryTxnsResponse) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseTreasuryTxnsResponse) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseTreasuryTxnsResponse) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseTreasuryTxnsResponse) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseTreasuryTxnsResponse) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseTreasuryTxnsResponse) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseTreasuryTxnsResponse) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetLiabilityDetails

`func (o *ResponseTreasuryTxnsResponse) GetLiabilityDetails() string`

GetLiabilityDetails returns the LiabilityDetails field if non-nil, zero value otherwise.

### GetLiabilityDetailsOk

`func (o *ResponseTreasuryTxnsResponse) GetLiabilityDetailsOk() (*string, bool)`

GetLiabilityDetailsOk returns a tuple with the LiabilityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityDetails

`func (o *ResponseTreasuryTxnsResponse) SetLiabilityDetails(v string)`

SetLiabilityDetails sets LiabilityDetails field to given value.

### HasLiabilityDetails

`func (o *ResponseTreasuryTxnsResponse) HasLiabilityDetails() bool`

HasLiabilityDetails returns a boolean if a field has been set.

### GetLimitId

`func (o *ResponseTreasuryTxnsResponse) GetLimitId() string`

GetLimitId returns the LimitId field if non-nil, zero value otherwise.

### GetLimitIdOk

`func (o *ResponseTreasuryTxnsResponse) GetLimitIdOk() (*string, bool)`

GetLimitIdOk returns a tuple with the LimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitId

`func (o *ResponseTreasuryTxnsResponse) SetLimitId(v string)`

SetLimitId sets LimitId field to given value.

### HasLimitId

`func (o *ResponseTreasuryTxnsResponse) HasLimitId() bool`

HasLimitId returns a boolean if a field has been set.

### GetPayeeName

`func (o *ResponseTreasuryTxnsResponse) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *ResponseTreasuryTxnsResponse) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *ResponseTreasuryTxnsResponse) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *ResponseTreasuryTxnsResponse) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseTreasuryTxnsResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseTreasuryTxnsResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseTreasuryTxnsResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseTreasuryTxnsResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseTreasuryTxnsResponse) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseTreasuryTxnsResponse) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseTreasuryTxnsResponse) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseTreasuryTxnsResponse) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseTreasuryTxnsResponse) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseTreasuryTxnsResponse) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseTreasuryTxnsResponse) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseTreasuryTxnsResponse) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseTreasuryTxnsResponse) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseTreasuryTxnsResponse) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseTreasuryTxnsResponse) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseTreasuryTxnsResponse) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseTreasuryTxnsResponse) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseTreasuryTxnsResponse) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseTreasuryTxnsResponse) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseTreasuryTxnsResponse) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqOfficeName

`func (o *ResponseTreasuryTxnsResponse) GetReqOfficeName() string`

GetReqOfficeName returns the ReqOfficeName field if non-nil, zero value otherwise.

### GetReqOfficeNameOk

`func (o *ResponseTreasuryTxnsResponse) GetReqOfficeNameOk() (*string, bool)`

GetReqOfficeNameOk returns a tuple with the ReqOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeName

`func (o *ResponseTreasuryTxnsResponse) SetReqOfficeName(v string)`

SetReqOfficeName sets ReqOfficeName field to given value.

### HasReqOfficeName

`func (o *ResponseTreasuryTxnsResponse) HasReqOfficeName() bool`

HasReqOfficeName returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseTreasuryTxnsResponse) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseTreasuryTxnsResponse) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseTreasuryTxnsResponse) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseTreasuryTxnsResponse) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseTreasuryTxnsResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseTreasuryTxnsResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseTreasuryTxnsResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseTreasuryTxnsResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseTreasuryTxnsResponse) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseTreasuryTxnsResponse) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseTreasuryTxnsResponse) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseTreasuryTxnsResponse) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseTreasuryTxnsResponse) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseTreasuryTxnsResponse) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseTreasuryTxnsResponse) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseTreasuryTxnsResponse) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseTreasuryTxnsResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseTreasuryTxnsResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseTreasuryTxnsResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseTreasuryTxnsResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseTreasuryTxnsResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseTreasuryTxnsResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseTreasuryTxnsResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseTreasuryTxnsResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTxnMode

`func (o *ResponseTreasuryTxnsResponse) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *ResponseTreasuryTxnsResponse) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *ResponseTreasuryTxnsResponse) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.

### HasTxnMode

`func (o *ResponseTreasuryTxnsResponse) HasTxnMode() bool`

HasTxnMode returns a boolean if a field has been set.

### GetTxnStatus

`func (o *ResponseTreasuryTxnsResponse) GetTxnStatus() string`

GetTxnStatus returns the TxnStatus field if non-nil, zero value otherwise.

### GetTxnStatusOk

`func (o *ResponseTreasuryTxnsResponse) GetTxnStatusOk() (*string, bool)`

GetTxnStatusOk returns a tuple with the TxnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnStatus

`func (o *ResponseTreasuryTxnsResponse) SetTxnStatus(v string)`

SetTxnStatus sets TxnStatus field to given value.

### HasTxnStatus

`func (o *ResponseTreasuryTxnsResponse) HasTxnStatus() bool`

HasTxnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



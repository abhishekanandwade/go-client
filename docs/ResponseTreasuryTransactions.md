# ResponseTreasuryTransactions

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
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**CurrencyDetails** | Pointer to **map[string]int32** |  | [optional] 
**EmployeeId1** | Pointer to **int32** |  | [optional] 
**EmployeeId2** | Pointer to **int32** |  | [optional] 
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
**ReqUserName** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | Pointer to **string** |  | [optional] 
**RequestType** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Transdate** | Pointer to **string** |  | [optional] 
**TxnMode** | Pointer to **string** |  | [optional] 
**TxnStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseTreasuryTransactions

`func NewResponseTreasuryTransactions() *ResponseTreasuryTransactions`

NewResponseTreasuryTransactions instantiates a new ResponseTreasuryTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTreasuryTransactionsWithDefaults

`func NewResponseTreasuryTransactionsWithDefaults() *ResponseTreasuryTransactions`

NewResponseTreasuryTransactionsWithDefaults instantiates a new ResponseTreasuryTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAmount

`func (o *ResponseTreasuryTransactions) GetAckAmount() float32`

GetAckAmount returns the AckAmount field if non-nil, zero value otherwise.

### GetAckAmountOk

`func (o *ResponseTreasuryTransactions) GetAckAmountOk() (*float32, bool)`

GetAckAmountOk returns a tuple with the AckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAmount

`func (o *ResponseTreasuryTransactions) SetAckAmount(v float32)`

SetAckAmount sets AckAmount field to given value.

### HasAckAmount

`func (o *ResponseTreasuryTransactions) HasAckAmount() bool`

HasAckAmount returns a boolean if a field has been set.

### GetAckApproverRemarks

`func (o *ResponseTreasuryTransactions) GetAckApproverRemarks() string`

GetAckApproverRemarks returns the AckApproverRemarks field if non-nil, zero value otherwise.

### GetAckApproverRemarksOk

`func (o *ResponseTreasuryTransactions) GetAckApproverRemarksOk() (*string, bool)`

GetAckApproverRemarksOk returns a tuple with the AckApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckApproverRemarks

`func (o *ResponseTreasuryTransactions) SetAckApproverRemarks(v string)`

SetAckApproverRemarks sets AckApproverRemarks field to given value.

### HasAckApproverRemarks

`func (o *ResponseTreasuryTransactions) HasAckApproverRemarks() bool`

HasAckApproverRemarks returns a boolean if a field has been set.

### GetAckDate

`func (o *ResponseTreasuryTransactions) GetAckDate() string`

GetAckDate returns the AckDate field if non-nil, zero value otherwise.

### GetAckDateOk

`func (o *ResponseTreasuryTransactions) GetAckDateOk() (*string, bool)`

GetAckDateOk returns a tuple with the AckDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDate

`func (o *ResponseTreasuryTransactions) SetAckDate(v string)`

SetAckDate sets AckDate field to given value.

### HasAckDate

`func (o *ResponseTreasuryTransactions) HasAckDate() bool`

HasAckDate returns a boolean if a field has been set.

### GetAckDetails

`func (o *ResponseTreasuryTransactions) GetAckDetails() map[string]int32`

GetAckDetails returns the AckDetails field if non-nil, zero value otherwise.

### GetAckDetailsOk

`func (o *ResponseTreasuryTransactions) GetAckDetailsOk() (*map[string]int32, bool)`

GetAckDetailsOk returns a tuple with the AckDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckDetails

`func (o *ResponseTreasuryTransactions) SetAckDetails(v map[string]int32)`

SetAckDetails sets AckDetails field to given value.

### HasAckDetails

`func (o *ResponseTreasuryTransactions) HasAckDetails() bool`

HasAckDetails returns a boolean if a field has been set.

### GetAckUserId

`func (o *ResponseTreasuryTransactions) GetAckUserId() int32`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *ResponseTreasuryTransactions) GetAckUserIdOk() (*int32, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *ResponseTreasuryTransactions) SetAckUserId(v int32)`

SetAckUserId sets AckUserId field to given value.

### HasAckUserId

`func (o *ResponseTreasuryTransactions) HasAckUserId() bool`

HasAckUserId returns a boolean if a field has been set.

### GetApprovedAmount

`func (o *ResponseTreasuryTransactions) GetApprovedAmount() float32`

GetApprovedAmount returns the ApprovedAmount field if non-nil, zero value otherwise.

### GetApprovedAmountOk

`func (o *ResponseTreasuryTransactions) GetApprovedAmountOk() (*float32, bool)`

GetApprovedAmountOk returns a tuple with the ApprovedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAmount

`func (o *ResponseTreasuryTransactions) SetApprovedAmount(v float32)`

SetApprovedAmount sets ApprovedAmount field to given value.

### HasApprovedAmount

`func (o *ResponseTreasuryTransactions) HasApprovedAmount() bool`

HasApprovedAmount returns a boolean if a field has been set.

### GetApprovedDetails

`func (o *ResponseTreasuryTransactions) GetApprovedDetails() map[string]int32`

GetApprovedDetails returns the ApprovedDetails field if non-nil, zero value otherwise.

### GetApprovedDetailsOk

`func (o *ResponseTreasuryTransactions) GetApprovedDetailsOk() (*map[string]int32, bool)`

GetApprovedDetailsOk returns a tuple with the ApprovedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDetails

`func (o *ResponseTreasuryTransactions) SetApprovedDetails(v map[string]int32)`

SetApprovedDetails sets ApprovedDetails field to given value.

### HasApprovedDetails

`func (o *ResponseTreasuryTransactions) HasApprovedDetails() bool`

HasApprovedDetails returns a boolean if a field has been set.

### GetChequeDate

`func (o *ResponseTreasuryTransactions) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *ResponseTreasuryTransactions) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *ResponseTreasuryTransactions) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *ResponseTreasuryTransactions) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *ResponseTreasuryTransactions) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *ResponseTreasuryTransactions) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *ResponseTreasuryTransactions) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *ResponseTreasuryTransactions) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetCurrencyDetails

`func (o *ResponseTreasuryTransactions) GetCurrencyDetails() map[string]int32`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *ResponseTreasuryTransactions) GetCurrencyDetailsOk() (*map[string]int32, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *ResponseTreasuryTransactions) SetCurrencyDetails(v map[string]int32)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *ResponseTreasuryTransactions) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetEmployeeId1

`func (o *ResponseTreasuryTransactions) GetEmployeeId1() int32`

GetEmployeeId1 returns the EmployeeId1 field if non-nil, zero value otherwise.

### GetEmployeeId1Ok

`func (o *ResponseTreasuryTransactions) GetEmployeeId1Ok() (*int32, bool)`

GetEmployeeId1Ok returns a tuple with the EmployeeId1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId1

`func (o *ResponseTreasuryTransactions) SetEmployeeId1(v int32)`

SetEmployeeId1 sets EmployeeId1 field to given value.

### HasEmployeeId1

`func (o *ResponseTreasuryTransactions) HasEmployeeId1() bool`

HasEmployeeId1 returns a boolean if a field has been set.

### GetEmployeeId2

`func (o *ResponseTreasuryTransactions) GetEmployeeId2() int32`

GetEmployeeId2 returns the EmployeeId2 field if non-nil, zero value otherwise.

### GetEmployeeId2Ok

`func (o *ResponseTreasuryTransactions) GetEmployeeId2Ok() (*int32, bool)`

GetEmployeeId2Ok returns a tuple with the EmployeeId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId2

`func (o *ResponseTreasuryTransactions) SetEmployeeId2(v int32)`

SetEmployeeId2 sets EmployeeId2 field to given value.

### HasEmployeeId2

`func (o *ResponseTreasuryTransactions) HasEmployeeId2() bool`

HasEmployeeId2 returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseTreasuryTransactions) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseTreasuryTransactions) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseTreasuryTransactions) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseTreasuryTransactions) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *ResponseTreasuryTransactions) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *ResponseTreasuryTransactions) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *ResponseTreasuryTransactions) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *ResponseTreasuryTransactions) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *ResponseTreasuryTransactions) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *ResponseTreasuryTransactions) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *ResponseTreasuryTransactions) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *ResponseTreasuryTransactions) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssApproverDate

`func (o *ResponseTreasuryTransactions) GetIssApproverDate() string`

GetIssApproverDate returns the IssApproverDate field if non-nil, zero value otherwise.

### GetIssApproverDateOk

`func (o *ResponseTreasuryTransactions) GetIssApproverDateOk() (*string, bool)`

GetIssApproverDateOk returns a tuple with the IssApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverDate

`func (o *ResponseTreasuryTransactions) SetIssApproverDate(v string)`

SetIssApproverDate sets IssApproverDate field to given value.

### HasIssApproverDate

`func (o *ResponseTreasuryTransactions) HasIssApproverDate() bool`

HasIssApproverDate returns a boolean if a field has been set.

### GetIssApproverId

`func (o *ResponseTreasuryTransactions) GetIssApproverId() int32`

GetIssApproverId returns the IssApproverId field if non-nil, zero value otherwise.

### GetIssApproverIdOk

`func (o *ResponseTreasuryTransactions) GetIssApproverIdOk() (*int32, bool)`

GetIssApproverIdOk returns a tuple with the IssApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverId

`func (o *ResponseTreasuryTransactions) SetIssApproverId(v int32)`

SetIssApproverId sets IssApproverId field to given value.

### HasIssApproverId

`func (o *ResponseTreasuryTransactions) HasIssApproverId() bool`

HasIssApproverId returns a boolean if a field has been set.

### GetIssApproverRemarks

`func (o *ResponseTreasuryTransactions) GetIssApproverRemarks() string`

GetIssApproverRemarks returns the IssApproverRemarks field if non-nil, zero value otherwise.

### GetIssApproverRemarksOk

`func (o *ResponseTreasuryTransactions) GetIssApproverRemarksOk() (*string, bool)`

GetIssApproverRemarksOk returns a tuple with the IssApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssApproverRemarks

`func (o *ResponseTreasuryTransactions) SetIssApproverRemarks(v string)`

SetIssApproverRemarks sets IssApproverRemarks field to given value.

### HasIssApproverRemarks

`func (o *ResponseTreasuryTransactions) HasIssApproverRemarks() bool`

HasIssApproverRemarks returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseTreasuryTransactions) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseTreasuryTransactions) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseTreasuryTransactions) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseTreasuryTransactions) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetIssOfficeName

`func (o *ResponseTreasuryTransactions) GetIssOfficeName() string`

GetIssOfficeName returns the IssOfficeName field if non-nil, zero value otherwise.

### GetIssOfficeNameOk

`func (o *ResponseTreasuryTransactions) GetIssOfficeNameOk() (*string, bool)`

GetIssOfficeNameOk returns a tuple with the IssOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeName

`func (o *ResponseTreasuryTransactions) SetIssOfficeName(v string)`

SetIssOfficeName sets IssOfficeName field to given value.

### HasIssOfficeName

`func (o *ResponseTreasuryTransactions) HasIssOfficeName() bool`

HasIssOfficeName returns a boolean if a field has been set.

### GetIssUserId

`func (o *ResponseTreasuryTransactions) GetIssUserId() int32`

GetIssUserId returns the IssUserId field if non-nil, zero value otherwise.

### GetIssUserIdOk

`func (o *ResponseTreasuryTransactions) GetIssUserIdOk() (*int32, bool)`

GetIssUserIdOk returns a tuple with the IssUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserId

`func (o *ResponseTreasuryTransactions) SetIssUserId(v int32)`

SetIssUserId sets IssUserId field to given value.

### HasIssUserId

`func (o *ResponseTreasuryTransactions) HasIssUserId() bool`

HasIssUserId returns a boolean if a field has been set.

### GetIssUserName

`func (o *ResponseTreasuryTransactions) GetIssUserName() string`

GetIssUserName returns the IssUserName field if non-nil, zero value otherwise.

### GetIssUserNameOk

`func (o *ResponseTreasuryTransactions) GetIssUserNameOk() (*string, bool)`

GetIssUserNameOk returns a tuple with the IssUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserName

`func (o *ResponseTreasuryTransactions) SetIssUserName(v string)`

SetIssUserName sets IssUserName field to given value.

### HasIssUserName

`func (o *ResponseTreasuryTransactions) HasIssUserName() bool`

HasIssUserName returns a boolean if a field has been set.

### GetIssUserProcessedDate

`func (o *ResponseTreasuryTransactions) GetIssUserProcessedDate() string`

GetIssUserProcessedDate returns the IssUserProcessedDate field if non-nil, zero value otherwise.

### GetIssUserProcessedDateOk

`func (o *ResponseTreasuryTransactions) GetIssUserProcessedDateOk() (*string, bool)`

GetIssUserProcessedDateOk returns a tuple with the IssUserProcessedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssUserProcessedDate

`func (o *ResponseTreasuryTransactions) SetIssUserProcessedDate(v string)`

SetIssUserProcessedDate sets IssUserProcessedDate field to given value.

### HasIssUserProcessedDate

`func (o *ResponseTreasuryTransactions) HasIssUserProcessedDate() bool`

HasIssUserProcessedDate returns a boolean if a field has been set.

### GetLiabilityDetails

`func (o *ResponseTreasuryTransactions) GetLiabilityDetails() string`

GetLiabilityDetails returns the LiabilityDetails field if non-nil, zero value otherwise.

### GetLiabilityDetailsOk

`func (o *ResponseTreasuryTransactions) GetLiabilityDetailsOk() (*string, bool)`

GetLiabilityDetailsOk returns a tuple with the LiabilityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityDetails

`func (o *ResponseTreasuryTransactions) SetLiabilityDetails(v string)`

SetLiabilityDetails sets LiabilityDetails field to given value.

### HasLiabilityDetails

`func (o *ResponseTreasuryTransactions) HasLiabilityDetails() bool`

HasLiabilityDetails returns a boolean if a field has been set.

### GetLimitId

`func (o *ResponseTreasuryTransactions) GetLimitId() string`

GetLimitId returns the LimitId field if non-nil, zero value otherwise.

### GetLimitIdOk

`func (o *ResponseTreasuryTransactions) GetLimitIdOk() (*string, bool)`

GetLimitIdOk returns a tuple with the LimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitId

`func (o *ResponseTreasuryTransactions) SetLimitId(v string)`

SetLimitId sets LimitId field to given value.

### HasLimitId

`func (o *ResponseTreasuryTransactions) HasLimitId() bool`

HasLimitId returns a boolean if a field has been set.

### GetPayeeName

`func (o *ResponseTreasuryTransactions) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *ResponseTreasuryTransactions) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *ResponseTreasuryTransactions) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *ResponseTreasuryTransactions) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseTreasuryTransactions) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseTreasuryTransactions) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseTreasuryTransactions) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseTreasuryTransactions) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *ResponseTreasuryTransactions) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *ResponseTreasuryTransactions) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *ResponseTreasuryTransactions) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.

### HasReqAmount

`func (o *ResponseTreasuryTransactions) HasReqAmount() bool`

HasReqAmount returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseTreasuryTransactions) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseTreasuryTransactions) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseTreasuryTransactions) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseTreasuryTransactions) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetReqApproverDate

`func (o *ResponseTreasuryTransactions) GetReqApproverDate() string`

GetReqApproverDate returns the ReqApproverDate field if non-nil, zero value otherwise.

### GetReqApproverDateOk

`func (o *ResponseTreasuryTransactions) GetReqApproverDateOk() (*string, bool)`

GetReqApproverDateOk returns a tuple with the ReqApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverDate

`func (o *ResponseTreasuryTransactions) SetReqApproverDate(v string)`

SetReqApproverDate sets ReqApproverDate field to given value.

### HasReqApproverDate

`func (o *ResponseTreasuryTransactions) HasReqApproverDate() bool`

HasReqApproverDate returns a boolean if a field has been set.

### GetReqApproverId

`func (o *ResponseTreasuryTransactions) GetReqApproverId() int32`

GetReqApproverId returns the ReqApproverId field if non-nil, zero value otherwise.

### GetReqApproverIdOk

`func (o *ResponseTreasuryTransactions) GetReqApproverIdOk() (*int32, bool)`

GetReqApproverIdOk returns a tuple with the ReqApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverId

`func (o *ResponseTreasuryTransactions) SetReqApproverId(v int32)`

SetReqApproverId sets ReqApproverId field to given value.

### HasReqApproverId

`func (o *ResponseTreasuryTransactions) HasReqApproverId() bool`

HasReqApproverId returns a boolean if a field has been set.

### GetReqApproverRemarks

`func (o *ResponseTreasuryTransactions) GetReqApproverRemarks() string`

GetReqApproverRemarks returns the ReqApproverRemarks field if non-nil, zero value otherwise.

### GetReqApproverRemarksOk

`func (o *ResponseTreasuryTransactions) GetReqApproverRemarksOk() (*string, bool)`

GetReqApproverRemarksOk returns a tuple with the ReqApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverRemarks

`func (o *ResponseTreasuryTransactions) SetReqApproverRemarks(v string)`

SetReqApproverRemarks sets ReqApproverRemarks field to given value.

### HasReqApproverRemarks

`func (o *ResponseTreasuryTransactions) HasReqApproverRemarks() bool`

HasReqApproverRemarks returns a boolean if a field has been set.

### GetReqOfficeId

`func (o *ResponseTreasuryTransactions) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *ResponseTreasuryTransactions) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *ResponseTreasuryTransactions) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.

### HasReqOfficeId

`func (o *ResponseTreasuryTransactions) HasReqOfficeId() bool`

HasReqOfficeId returns a boolean if a field has been set.

### GetReqOfficeName

`func (o *ResponseTreasuryTransactions) GetReqOfficeName() string`

GetReqOfficeName returns the ReqOfficeName field if non-nil, zero value otherwise.

### GetReqOfficeNameOk

`func (o *ResponseTreasuryTransactions) GetReqOfficeNameOk() (*string, bool)`

GetReqOfficeNameOk returns a tuple with the ReqOfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeName

`func (o *ResponseTreasuryTransactions) SetReqOfficeName(v string)`

SetReqOfficeName sets ReqOfficeName field to given value.

### HasReqOfficeName

`func (o *ResponseTreasuryTransactions) HasReqOfficeName() bool`

HasReqOfficeName returns a boolean if a field has been set.

### GetReqUserId

`func (o *ResponseTreasuryTransactions) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *ResponseTreasuryTransactions) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *ResponseTreasuryTransactions) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.

### HasReqUserId

`func (o *ResponseTreasuryTransactions) HasReqUserId() bool`

HasReqUserId returns a boolean if a field has been set.

### GetReqUserName

`func (o *ResponseTreasuryTransactions) GetReqUserName() string`

GetReqUserName returns the ReqUserName field if non-nil, zero value otherwise.

### GetReqUserNameOk

`func (o *ResponseTreasuryTransactions) GetReqUserNameOk() (*string, bool)`

GetReqUserNameOk returns a tuple with the ReqUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserName

`func (o *ResponseTreasuryTransactions) SetReqUserName(v string)`

SetReqUserName sets ReqUserName field to given value.

### HasReqUserName

`func (o *ResponseTreasuryTransactions) HasReqUserName() bool`

HasReqUserName returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseTreasuryTransactions) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseTreasuryTransactions) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseTreasuryTransactions) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseTreasuryTransactions) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *ResponseTreasuryTransactions) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ResponseTreasuryTransactions) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ResponseTreasuryTransactions) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ResponseTreasuryTransactions) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ResponseTreasuryTransactions) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ResponseTreasuryTransactions) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ResponseTreasuryTransactions) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ResponseTreasuryTransactions) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseTreasuryTransactions) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseTreasuryTransactions) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseTreasuryTransactions) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseTreasuryTransactions) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseTreasuryTransactions) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseTreasuryTransactions) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseTreasuryTransactions) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseTreasuryTransactions) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransdate

`func (o *ResponseTreasuryTransactions) GetTransdate() string`

GetTransdate returns the Transdate field if non-nil, zero value otherwise.

### GetTransdateOk

`func (o *ResponseTreasuryTransactions) GetTransdateOk() (*string, bool)`

GetTransdateOk returns a tuple with the Transdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransdate

`func (o *ResponseTreasuryTransactions) SetTransdate(v string)`

SetTransdate sets Transdate field to given value.

### HasTransdate

`func (o *ResponseTreasuryTransactions) HasTransdate() bool`

HasTransdate returns a boolean if a field has been set.

### GetTxnMode

`func (o *ResponseTreasuryTransactions) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *ResponseTreasuryTransactions) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *ResponseTreasuryTransactions) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.

### HasTxnMode

`func (o *ResponseTreasuryTransactions) HasTxnMode() bool`

HasTxnMode returns a boolean if a field has been set.

### GetTxnStatus

`func (o *ResponseTreasuryTransactions) GetTxnStatus() string`

GetTxnStatus returns the TxnStatus field if non-nil, zero value otherwise.

### GetTxnStatusOk

`func (o *ResponseTreasuryTransactions) GetTxnStatusOk() (*string, bool)`

GetTxnStatusOk returns a tuple with the TxnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnStatus

`func (o *ResponseTreasuryTransactions) SetTxnStatus(v string)`

SetTxnStatus sets TxnStatus field to given value.

### HasTxnStatus

`func (o *ResponseTreasuryTransactions) HasTxnStatus() bool`

HasTxnStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerCreateTransactionsBankRemittancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BagWeight** | Pointer to **float32** |  | [optional] 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**CurrencyDetails** | Pointer to **map[string]int32** |  | [optional] 
**EmployeeId1** | Pointer to **int32** |  | [optional] 
**EmployeeId2** | Pointer to **int32** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**IsSpecialRemittance** | Pointer to **bool** |  | [optional] 
**IssOfficeId** | **int32** |  | 
**LiabilityDetails** | Pointer to **string** |  | [optional] 
**LimitId** | Pointer to **string** |  | [optional] 
**PayeeName** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**ReqAmount** | **float32** |  | 
**ReqOfficeId** | **int32** |  | 
**ReqUserId** | **int32** |  | 
**RequestId** | Pointer to **string** |  | [optional] 
**RequestSource** | **string** |  | 
**RequestType** | Pointer to **string** |  | [optional] 
**TxnMode** | **string** |  | 

## Methods

### NewHandlerCreateTransactionsBankRemittancesRequest

`func NewHandlerCreateTransactionsBankRemittancesRequest(issOfficeId int32, reqAmount float32, reqOfficeId int32, reqUserId int32, requestSource string, txnMode string, ) *HandlerCreateTransactionsBankRemittancesRequest`

NewHandlerCreateTransactionsBankRemittancesRequest instantiates a new HandlerCreateTransactionsBankRemittancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateTransactionsBankRemittancesRequestWithDefaults

`func NewHandlerCreateTransactionsBankRemittancesRequestWithDefaults() *HandlerCreateTransactionsBankRemittancesRequest`

NewHandlerCreateTransactionsBankRemittancesRequestWithDefaults instantiates a new HandlerCreateTransactionsBankRemittancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBagWeight

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetBagWeight() float32`

GetBagWeight returns the BagWeight field if non-nil, zero value otherwise.

### GetBagWeightOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetBagWeightOk() (*float32, bool)`

GetBagWeightOk returns a tuple with the BagWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagWeight

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetBagWeight(v float32)`

SetBagWeight sets BagWeight field to given value.

### HasBagWeight

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasBagWeight() bool`

HasBagWeight returns a boolean if a field has been set.

### GetChequeDate

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetCurrencyDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetCurrencyDetails() map[string]int32`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetCurrencyDetailsOk() (*map[string]int32, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetCurrencyDetails(v map[string]int32)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetEmployeeId1

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetEmployeeId1() int32`

GetEmployeeId1 returns the EmployeeId1 field if non-nil, zero value otherwise.

### GetEmployeeId1Ok

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetEmployeeId1Ok() (*int32, bool)`

GetEmployeeId1Ok returns a tuple with the EmployeeId1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId1

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetEmployeeId1(v int32)`

SetEmployeeId1 sets EmployeeId1 field to given value.

### HasEmployeeId1

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasEmployeeId1() bool`

HasEmployeeId1 returns a boolean if a field has been set.

### GetEmployeeId2

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetEmployeeId2() int32`

GetEmployeeId2 returns the EmployeeId2 field if non-nil, zero value otherwise.

### GetEmployeeId2Ok

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetEmployeeId2Ok() (*int32, bool)`

GetEmployeeId2Ok returns a tuple with the EmployeeId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId2

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetEmployeeId2(v int32)`

SetEmployeeId2 sets EmployeeId2 field to given value.

### HasEmployeeId2

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasEmployeeId2() bool`

HasEmployeeId2 returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetIsSpecialRemittance

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIsSpecialRemittance() bool`

GetIsSpecialRemittance returns the IsSpecialRemittance field if non-nil, zero value otherwise.

### GetIsSpecialRemittanceOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIsSpecialRemittanceOk() (*bool, bool)`

GetIsSpecialRemittanceOk returns a tuple with the IsSpecialRemittance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpecialRemittance

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetIsSpecialRemittance(v bool)`

SetIsSpecialRemittance sets IsSpecialRemittance field to given value.

### HasIsSpecialRemittance

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasIsSpecialRemittance() bool`

HasIsSpecialRemittance returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.


### GetLiabilityDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetLiabilityDetails() string`

GetLiabilityDetails returns the LiabilityDetails field if non-nil, zero value otherwise.

### GetLiabilityDetailsOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetLiabilityDetailsOk() (*string, bool)`

GetLiabilityDetailsOk returns a tuple with the LiabilityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilityDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetLiabilityDetails(v string)`

SetLiabilityDetails sets LiabilityDetails field to given value.

### HasLiabilityDetails

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasLiabilityDetails() bool`

HasLiabilityDetails returns a boolean if a field has been set.

### GetLimitId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetLimitId() string`

GetLimitId returns the LimitId field if non-nil, zero value otherwise.

### GetLimitIdOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetLimitIdOk() (*string, bool)`

GetLimitIdOk returns a tuple with the LimitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetLimitId(v string)`

SetLimitId sets LimitId field to given value.

### HasLimitId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasLimitId() bool`

HasLimitId returns a boolean if a field has been set.

### GetPayeeName

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetPayeeName() string`

GetPayeeName returns the PayeeName field if non-nil, zero value otherwise.

### GetPayeeNameOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetPayeeNameOk() (*string, bool)`

GetPayeeNameOk returns a tuple with the PayeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeName

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetPayeeName(v string)`

SetPayeeName sets PayeeName field to given value.

### HasPayeeName

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasPayeeName() bool`

HasPayeeName returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetReqAmount

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqAmount() float32`

GetReqAmount returns the ReqAmount field if non-nil, zero value otherwise.

### GetReqAmountOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqAmountOk() (*float32, bool)`

GetReqAmountOk returns a tuple with the ReqAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqAmount

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetReqAmount(v float32)`

SetReqAmount sets ReqAmount field to given value.


### GetReqOfficeId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqOfficeId() int32`

GetReqOfficeId returns the ReqOfficeId field if non-nil, zero value otherwise.

### GetReqOfficeIdOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqOfficeIdOk() (*int32, bool)`

GetReqOfficeIdOk returns a tuple with the ReqOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqOfficeId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetReqOfficeId(v int32)`

SetReqOfficeId sets ReqOfficeId field to given value.


### GetReqUserId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqUserId() int32`

GetReqUserId returns the ReqUserId field if non-nil, zero value otherwise.

### GetReqUserIdOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetReqUserIdOk() (*int32, bool)`

GetReqUserIdOk returns a tuple with the ReqUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUserId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetReqUserId(v int32)`

SetReqUserId sets ReqUserId field to given value.


### GetRequestId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetRequestSource

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestSource() string`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestSourceOk() (*string, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetRequestSource(v string)`

SetRequestSource sets RequestSource field to given value.


### GetRequestType

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *HandlerCreateTransactionsBankRemittancesRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetTxnMode

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetTxnMode() string`

GetTxnMode returns the TxnMode field if non-nil, zero value otherwise.

### GetTxnModeOk

`func (o *HandlerCreateTransactionsBankRemittancesRequest) GetTxnModeOk() (*string, bool)`

GetTxnModeOk returns a tuple with the TxnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxnMode

`func (o *HandlerCreateTransactionsBankRemittancesRequest) SetTxnMode(v string)`

SetTxnMode sets TxnMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DomainTcbBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeIssueDetails** | Pointer to [**[]DomainChequeDetails**](DomainChequeDetails.md) |  | [optional] 
**ChequeRemittanceDetails** | Pointer to [**[]DomainChequeDetails**](DomainChequeDetails.md) |  | [optional] 
**ClosingBal** | Pointer to **float32** |  | [optional] 
**CurrencyDetails** | Pointer to **string** |  | [optional] 
**CurrencyIdCount** | Pointer to **map[string]int32** |  | [optional] 
**DayendStatus** | Pointer to **int32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**OpeningBal** | Pointer to **float32** |  | [optional] 
**Payments** | Pointer to **float32** |  | [optional] 
**Receipts** | Pointer to **float32** |  | [optional] 
**TcbData** | Pointer to [**[]DomainTCBData**](DomainTCBData.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainTcbBalance

`func NewDomainTcbBalance() *DomainTcbBalance`

NewDomainTcbBalance instantiates a new DomainTcbBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTcbBalanceWithDefaults

`func NewDomainTcbBalanceWithDefaults() *DomainTcbBalance`

NewDomainTcbBalanceWithDefaults instantiates a new DomainTcbBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeIssueDetails

`func (o *DomainTcbBalance) GetChequeIssueDetails() []DomainChequeDetails`

GetChequeIssueDetails returns the ChequeIssueDetails field if non-nil, zero value otherwise.

### GetChequeIssueDetailsOk

`func (o *DomainTcbBalance) GetChequeIssueDetailsOk() (*[]DomainChequeDetails, bool)`

GetChequeIssueDetailsOk returns a tuple with the ChequeIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeIssueDetails

`func (o *DomainTcbBalance) SetChequeIssueDetails(v []DomainChequeDetails)`

SetChequeIssueDetails sets ChequeIssueDetails field to given value.

### HasChequeIssueDetails

`func (o *DomainTcbBalance) HasChequeIssueDetails() bool`

HasChequeIssueDetails returns a boolean if a field has been set.

### GetChequeRemittanceDetails

`func (o *DomainTcbBalance) GetChequeRemittanceDetails() []DomainChequeDetails`

GetChequeRemittanceDetails returns the ChequeRemittanceDetails field if non-nil, zero value otherwise.

### GetChequeRemittanceDetailsOk

`func (o *DomainTcbBalance) GetChequeRemittanceDetailsOk() (*[]DomainChequeDetails, bool)`

GetChequeRemittanceDetailsOk returns a tuple with the ChequeRemittanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeRemittanceDetails

`func (o *DomainTcbBalance) SetChequeRemittanceDetails(v []DomainChequeDetails)`

SetChequeRemittanceDetails sets ChequeRemittanceDetails field to given value.

### HasChequeRemittanceDetails

`func (o *DomainTcbBalance) HasChequeRemittanceDetails() bool`

HasChequeRemittanceDetails returns a boolean if a field has been set.

### GetClosingBal

`func (o *DomainTcbBalance) GetClosingBal() float32`

GetClosingBal returns the ClosingBal field if non-nil, zero value otherwise.

### GetClosingBalOk

`func (o *DomainTcbBalance) GetClosingBalOk() (*float32, bool)`

GetClosingBalOk returns a tuple with the ClosingBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBal

`func (o *DomainTcbBalance) SetClosingBal(v float32)`

SetClosingBal sets ClosingBal field to given value.

### HasClosingBal

`func (o *DomainTcbBalance) HasClosingBal() bool`

HasClosingBal returns a boolean if a field has been set.

### GetCurrencyDetails

`func (o *DomainTcbBalance) GetCurrencyDetails() string`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *DomainTcbBalance) GetCurrencyDetailsOk() (*string, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *DomainTcbBalance) SetCurrencyDetails(v string)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *DomainTcbBalance) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetCurrencyIdCount

`func (o *DomainTcbBalance) GetCurrencyIdCount() map[string]int32`

GetCurrencyIdCount returns the CurrencyIdCount field if non-nil, zero value otherwise.

### GetCurrencyIdCountOk

`func (o *DomainTcbBalance) GetCurrencyIdCountOk() (*map[string]int32, bool)`

GetCurrencyIdCountOk returns a tuple with the CurrencyIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIdCount

`func (o *DomainTcbBalance) SetCurrencyIdCount(v map[string]int32)`

SetCurrencyIdCount sets CurrencyIdCount field to given value.

### HasCurrencyIdCount

`func (o *DomainTcbBalance) HasCurrencyIdCount() bool`

HasCurrencyIdCount returns a boolean if a field has been set.

### GetDayendStatus

`func (o *DomainTcbBalance) GetDayendStatus() int32`

GetDayendStatus returns the DayendStatus field if non-nil, zero value otherwise.

### GetDayendStatusOk

`func (o *DomainTcbBalance) GetDayendStatusOk() (*int32, bool)`

GetDayendStatusOk returns a tuple with the DayendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayendStatus

`func (o *DomainTcbBalance) SetDayendStatus(v int32)`

SetDayendStatus sets DayendStatus field to given value.

### HasDayendStatus

`func (o *DomainTcbBalance) HasDayendStatus() bool`

HasDayendStatus returns a boolean if a field has been set.

### GetOfficeId

`func (o *DomainTcbBalance) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *DomainTcbBalance) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *DomainTcbBalance) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *DomainTcbBalance) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOpeningBal

`func (o *DomainTcbBalance) GetOpeningBal() float32`

GetOpeningBal returns the OpeningBal field if non-nil, zero value otherwise.

### GetOpeningBalOk

`func (o *DomainTcbBalance) GetOpeningBalOk() (*float32, bool)`

GetOpeningBalOk returns a tuple with the OpeningBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBal

`func (o *DomainTcbBalance) SetOpeningBal(v float32)`

SetOpeningBal sets OpeningBal field to given value.

### HasOpeningBal

`func (o *DomainTcbBalance) HasOpeningBal() bool`

HasOpeningBal returns a boolean if a field has been set.

### GetPayments

`func (o *DomainTcbBalance) GetPayments() float32`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *DomainTcbBalance) GetPaymentsOk() (*float32, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *DomainTcbBalance) SetPayments(v float32)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *DomainTcbBalance) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *DomainTcbBalance) GetReceipts() float32`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *DomainTcbBalance) GetReceiptsOk() (*float32, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *DomainTcbBalance) SetReceipts(v float32)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *DomainTcbBalance) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetTcbData

`func (o *DomainTcbBalance) GetTcbData() []DomainTCBData`

GetTcbData returns the TcbData field if non-nil, zero value otherwise.

### GetTcbDataOk

`func (o *DomainTcbBalance) GetTcbDataOk() (*[]DomainTCBData, bool)`

GetTcbDataOk returns a tuple with the TcbData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcbData

`func (o *DomainTcbBalance) SetTcbData(v []DomainTCBData)`

SetTcbData sets TcbData field to given value.

### HasTcbData

`func (o *DomainTcbBalance) HasTcbData() bool`

HasTcbData returns a boolean if a field has been set.

### GetTransDate

`func (o *DomainTcbBalance) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *DomainTcbBalance) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *DomainTcbBalance) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *DomainTcbBalance) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *DomainTcbBalance) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *DomainTcbBalance) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *DomainTcbBalance) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *DomainTcbBalance) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetUserId

`func (o *DomainTcbBalance) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DomainTcbBalance) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DomainTcbBalance) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DomainTcbBalance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *DomainTcbBalance) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DomainTcbBalance) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DomainTcbBalance) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DomainTcbBalance) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



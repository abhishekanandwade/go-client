# ResponseTcbBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChequeIssueDetails** | Pointer to [**[]ResponseChequeDetails**](ResponseChequeDetails.md) |  | [optional] 
**ChequeRemittanceDetails** | Pointer to [**[]ResponseChequeDetails**](ResponseChequeDetails.md) |  | [optional] 
**ClosingBal** | Pointer to **float32** |  | [optional] 
**CurrencyDetails** | Pointer to **string** |  | [optional] 
**CurrencyIdCount** | Pointer to **map[string]int32** |  | [optional] 
**DayendStatus** | Pointer to **int32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**OpeningBal** | Pointer to **float32** |  | [optional] 
**Payments** | Pointer to **float32** |  | [optional] 
**Receipts** | Pointer to **float32** |  | [optional] 
**TcbData** | Pointer to [**[]ResponseTCBData**](ResponseTCBData.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseTcbBalance

`func NewResponseTcbBalance() *ResponseTcbBalance`

NewResponseTcbBalance instantiates a new ResponseTcbBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTcbBalanceWithDefaults

`func NewResponseTcbBalanceWithDefaults() *ResponseTcbBalance`

NewResponseTcbBalanceWithDefaults instantiates a new ResponseTcbBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChequeIssueDetails

`func (o *ResponseTcbBalance) GetChequeIssueDetails() []ResponseChequeDetails`

GetChequeIssueDetails returns the ChequeIssueDetails field if non-nil, zero value otherwise.

### GetChequeIssueDetailsOk

`func (o *ResponseTcbBalance) GetChequeIssueDetailsOk() (*[]ResponseChequeDetails, bool)`

GetChequeIssueDetailsOk returns a tuple with the ChequeIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeIssueDetails

`func (o *ResponseTcbBalance) SetChequeIssueDetails(v []ResponseChequeDetails)`

SetChequeIssueDetails sets ChequeIssueDetails field to given value.

### HasChequeIssueDetails

`func (o *ResponseTcbBalance) HasChequeIssueDetails() bool`

HasChequeIssueDetails returns a boolean if a field has been set.

### GetChequeRemittanceDetails

`func (o *ResponseTcbBalance) GetChequeRemittanceDetails() []ResponseChequeDetails`

GetChequeRemittanceDetails returns the ChequeRemittanceDetails field if non-nil, zero value otherwise.

### GetChequeRemittanceDetailsOk

`func (o *ResponseTcbBalance) GetChequeRemittanceDetailsOk() (*[]ResponseChequeDetails, bool)`

GetChequeRemittanceDetailsOk returns a tuple with the ChequeRemittanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeRemittanceDetails

`func (o *ResponseTcbBalance) SetChequeRemittanceDetails(v []ResponseChequeDetails)`

SetChequeRemittanceDetails sets ChequeRemittanceDetails field to given value.

### HasChequeRemittanceDetails

`func (o *ResponseTcbBalance) HasChequeRemittanceDetails() bool`

HasChequeRemittanceDetails returns a boolean if a field has been set.

### GetClosingBal

`func (o *ResponseTcbBalance) GetClosingBal() float32`

GetClosingBal returns the ClosingBal field if non-nil, zero value otherwise.

### GetClosingBalOk

`func (o *ResponseTcbBalance) GetClosingBalOk() (*float32, bool)`

GetClosingBalOk returns a tuple with the ClosingBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBal

`func (o *ResponseTcbBalance) SetClosingBal(v float32)`

SetClosingBal sets ClosingBal field to given value.

### HasClosingBal

`func (o *ResponseTcbBalance) HasClosingBal() bool`

HasClosingBal returns a boolean if a field has been set.

### GetCurrencyDetails

`func (o *ResponseTcbBalance) GetCurrencyDetails() string`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *ResponseTcbBalance) GetCurrencyDetailsOk() (*string, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *ResponseTcbBalance) SetCurrencyDetails(v string)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *ResponseTcbBalance) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetCurrencyIdCount

`func (o *ResponseTcbBalance) GetCurrencyIdCount() map[string]int32`

GetCurrencyIdCount returns the CurrencyIdCount field if non-nil, zero value otherwise.

### GetCurrencyIdCountOk

`func (o *ResponseTcbBalance) GetCurrencyIdCountOk() (*map[string]int32, bool)`

GetCurrencyIdCountOk returns a tuple with the CurrencyIdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIdCount

`func (o *ResponseTcbBalance) SetCurrencyIdCount(v map[string]int32)`

SetCurrencyIdCount sets CurrencyIdCount field to given value.

### HasCurrencyIdCount

`func (o *ResponseTcbBalance) HasCurrencyIdCount() bool`

HasCurrencyIdCount returns a boolean if a field has been set.

### GetDayendStatus

`func (o *ResponseTcbBalance) GetDayendStatus() int32`

GetDayendStatus returns the DayendStatus field if non-nil, zero value otherwise.

### GetDayendStatusOk

`func (o *ResponseTcbBalance) GetDayendStatusOk() (*int32, bool)`

GetDayendStatusOk returns a tuple with the DayendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayendStatus

`func (o *ResponseTcbBalance) SetDayendStatus(v int32)`

SetDayendStatus sets DayendStatus field to given value.

### HasDayendStatus

`func (o *ResponseTcbBalance) HasDayendStatus() bool`

HasDayendStatus returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseTcbBalance) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseTcbBalance) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseTcbBalance) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseTcbBalance) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOpeningBal

`func (o *ResponseTcbBalance) GetOpeningBal() float32`

GetOpeningBal returns the OpeningBal field if non-nil, zero value otherwise.

### GetOpeningBalOk

`func (o *ResponseTcbBalance) GetOpeningBalOk() (*float32, bool)`

GetOpeningBalOk returns a tuple with the OpeningBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBal

`func (o *ResponseTcbBalance) SetOpeningBal(v float32)`

SetOpeningBal sets OpeningBal field to given value.

### HasOpeningBal

`func (o *ResponseTcbBalance) HasOpeningBal() bool`

HasOpeningBal returns a boolean if a field has been set.

### GetPayments

`func (o *ResponseTcbBalance) GetPayments() float32`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ResponseTcbBalance) GetPaymentsOk() (*float32, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ResponseTcbBalance) SetPayments(v float32)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *ResponseTcbBalance) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *ResponseTcbBalance) GetReceipts() float32`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *ResponseTcbBalance) GetReceiptsOk() (*float32, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *ResponseTcbBalance) SetReceipts(v float32)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *ResponseTcbBalance) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetTcbData

`func (o *ResponseTcbBalance) GetTcbData() []ResponseTCBData`

GetTcbData returns the TcbData field if non-nil, zero value otherwise.

### GetTcbDataOk

`func (o *ResponseTcbBalance) GetTcbDataOk() (*[]ResponseTCBData, bool)`

GetTcbDataOk returns a tuple with the TcbData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcbData

`func (o *ResponseTcbBalance) SetTcbData(v []ResponseTCBData)`

SetTcbData sets TcbData field to given value.

### HasTcbData

`func (o *ResponseTcbBalance) HasTcbData() bool`

HasTcbData returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseTcbBalance) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseTcbBalance) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseTcbBalance) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseTcbBalance) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseTcbBalance) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseTcbBalance) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseTcbBalance) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseTcbBalance) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseTcbBalance) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseTcbBalance) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseTcbBalance) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseTcbBalance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ResponseTcbBalance) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ResponseTcbBalance) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ResponseTcbBalance) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ResponseTcbBalance) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



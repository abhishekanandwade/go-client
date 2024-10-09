# ResponseSummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingDetails** | Pointer to [**[]ResponseAccountingDetails**](ResponseAccountingDetails.md) |  | [optional] 
**ChequeIssueDetails** | Pointer to [**[]ResponseChequeDetails**](ResponseChequeDetails.md) |  | [optional] 
**ChequeRemittanceDetails** | Pointer to [**[]ResponseChequeDetails**](ResponseChequeDetails.md) |  | [optional] 
**ClosingBal** | Pointer to **float32** |  | [optional] 
**DayendStatus** | Pointer to **int32** |  | [optional] 
**Ipobalance** | Pointer to [**[]ResponseIPOBal**](ResponseIPOBal.md) |  | [optional] 
**MaxBal** | Pointer to **float32** |  | [optional] 
**MinBal** | Pointer to **float32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**OpeningBal** | Pointer to **float32** |  | [optional] 
**Payments** | Pointer to **float32** |  | [optional] 
**Receipts** | Pointer to **float32** |  | [optional] 
**StampBalance** | Pointer to [**[]ResponseStampCatBal**](ResponseStampCatBal.md) |  | [optional] 
**SubOrdinateOfficeSummarydetails** | Pointer to [**[]ResponseSubOrdinateOfficeSummarydetails**](ResponseSubOrdinateOfficeSummarydetails.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransitCash** | Pointer to **float32** |  | [optional] 
**TransitIpos** | Pointer to **float32** |  | [optional] 
**TransitStamps** | Pointer to **float32** |  | [optional] 

## Methods

### NewResponseSummaryData

`func NewResponseSummaryData() *ResponseSummaryData`

NewResponseSummaryData instantiates a new ResponseSummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSummaryDataWithDefaults

`func NewResponseSummaryDataWithDefaults() *ResponseSummaryData`

NewResponseSummaryDataWithDefaults instantiates a new ResponseSummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingDetails

`func (o *ResponseSummaryData) GetAccountingDetails() []ResponseAccountingDetails`

GetAccountingDetails returns the AccountingDetails field if non-nil, zero value otherwise.

### GetAccountingDetailsOk

`func (o *ResponseSummaryData) GetAccountingDetailsOk() (*[]ResponseAccountingDetails, bool)`

GetAccountingDetailsOk returns a tuple with the AccountingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingDetails

`func (o *ResponseSummaryData) SetAccountingDetails(v []ResponseAccountingDetails)`

SetAccountingDetails sets AccountingDetails field to given value.

### HasAccountingDetails

`func (o *ResponseSummaryData) HasAccountingDetails() bool`

HasAccountingDetails returns a boolean if a field has been set.

### GetChequeIssueDetails

`func (o *ResponseSummaryData) GetChequeIssueDetails() []ResponseChequeDetails`

GetChequeIssueDetails returns the ChequeIssueDetails field if non-nil, zero value otherwise.

### GetChequeIssueDetailsOk

`func (o *ResponseSummaryData) GetChequeIssueDetailsOk() (*[]ResponseChequeDetails, bool)`

GetChequeIssueDetailsOk returns a tuple with the ChequeIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeIssueDetails

`func (o *ResponseSummaryData) SetChequeIssueDetails(v []ResponseChequeDetails)`

SetChequeIssueDetails sets ChequeIssueDetails field to given value.

### HasChequeIssueDetails

`func (o *ResponseSummaryData) HasChequeIssueDetails() bool`

HasChequeIssueDetails returns a boolean if a field has been set.

### GetChequeRemittanceDetails

`func (o *ResponseSummaryData) GetChequeRemittanceDetails() []ResponseChequeDetails`

GetChequeRemittanceDetails returns the ChequeRemittanceDetails field if non-nil, zero value otherwise.

### GetChequeRemittanceDetailsOk

`func (o *ResponseSummaryData) GetChequeRemittanceDetailsOk() (*[]ResponseChequeDetails, bool)`

GetChequeRemittanceDetailsOk returns a tuple with the ChequeRemittanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeRemittanceDetails

`func (o *ResponseSummaryData) SetChequeRemittanceDetails(v []ResponseChequeDetails)`

SetChequeRemittanceDetails sets ChequeRemittanceDetails field to given value.

### HasChequeRemittanceDetails

`func (o *ResponseSummaryData) HasChequeRemittanceDetails() bool`

HasChequeRemittanceDetails returns a boolean if a field has been set.

### GetClosingBal

`func (o *ResponseSummaryData) GetClosingBal() float32`

GetClosingBal returns the ClosingBal field if non-nil, zero value otherwise.

### GetClosingBalOk

`func (o *ResponseSummaryData) GetClosingBalOk() (*float32, bool)`

GetClosingBalOk returns a tuple with the ClosingBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBal

`func (o *ResponseSummaryData) SetClosingBal(v float32)`

SetClosingBal sets ClosingBal field to given value.

### HasClosingBal

`func (o *ResponseSummaryData) HasClosingBal() bool`

HasClosingBal returns a boolean if a field has been set.

### GetDayendStatus

`func (o *ResponseSummaryData) GetDayendStatus() int32`

GetDayendStatus returns the DayendStatus field if non-nil, zero value otherwise.

### GetDayendStatusOk

`func (o *ResponseSummaryData) GetDayendStatusOk() (*int32, bool)`

GetDayendStatusOk returns a tuple with the DayendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayendStatus

`func (o *ResponseSummaryData) SetDayendStatus(v int32)`

SetDayendStatus sets DayendStatus field to given value.

### HasDayendStatus

`func (o *ResponseSummaryData) HasDayendStatus() bool`

HasDayendStatus returns a boolean if a field has been set.

### GetIpobalance

`func (o *ResponseSummaryData) GetIpobalance() []ResponseIPOBal`

GetIpobalance returns the Ipobalance field if non-nil, zero value otherwise.

### GetIpobalanceOk

`func (o *ResponseSummaryData) GetIpobalanceOk() (*[]ResponseIPOBal, bool)`

GetIpobalanceOk returns a tuple with the Ipobalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpobalance

`func (o *ResponseSummaryData) SetIpobalance(v []ResponseIPOBal)`

SetIpobalance sets Ipobalance field to given value.

### HasIpobalance

`func (o *ResponseSummaryData) HasIpobalance() bool`

HasIpobalance returns a boolean if a field has been set.

### GetMaxBal

`func (o *ResponseSummaryData) GetMaxBal() float32`

GetMaxBal returns the MaxBal field if non-nil, zero value otherwise.

### GetMaxBalOk

`func (o *ResponseSummaryData) GetMaxBalOk() (*float32, bool)`

GetMaxBalOk returns a tuple with the MaxBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBal

`func (o *ResponseSummaryData) SetMaxBal(v float32)`

SetMaxBal sets MaxBal field to given value.

### HasMaxBal

`func (o *ResponseSummaryData) HasMaxBal() bool`

HasMaxBal returns a boolean if a field has been set.

### GetMinBal

`func (o *ResponseSummaryData) GetMinBal() float32`

GetMinBal returns the MinBal field if non-nil, zero value otherwise.

### GetMinBalOk

`func (o *ResponseSummaryData) GetMinBalOk() (*float32, bool)`

GetMinBalOk returns a tuple with the MinBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBal

`func (o *ResponseSummaryData) SetMinBal(v float32)`

SetMinBal sets MinBal field to given value.

### HasMinBal

`func (o *ResponseSummaryData) HasMinBal() bool`

HasMinBal returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseSummaryData) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseSummaryData) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseSummaryData) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseSummaryData) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOpeningBal

`func (o *ResponseSummaryData) GetOpeningBal() float32`

GetOpeningBal returns the OpeningBal field if non-nil, zero value otherwise.

### GetOpeningBalOk

`func (o *ResponseSummaryData) GetOpeningBalOk() (*float32, bool)`

GetOpeningBalOk returns a tuple with the OpeningBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBal

`func (o *ResponseSummaryData) SetOpeningBal(v float32)`

SetOpeningBal sets OpeningBal field to given value.

### HasOpeningBal

`func (o *ResponseSummaryData) HasOpeningBal() bool`

HasOpeningBal returns a boolean if a field has been set.

### GetPayments

`func (o *ResponseSummaryData) GetPayments() float32`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ResponseSummaryData) GetPaymentsOk() (*float32, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ResponseSummaryData) SetPayments(v float32)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *ResponseSummaryData) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *ResponseSummaryData) GetReceipts() float32`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *ResponseSummaryData) GetReceiptsOk() (*float32, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *ResponseSummaryData) SetReceipts(v float32)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *ResponseSummaryData) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetStampBalance

`func (o *ResponseSummaryData) GetStampBalance() []ResponseStampCatBal`

GetStampBalance returns the StampBalance field if non-nil, zero value otherwise.

### GetStampBalanceOk

`func (o *ResponseSummaryData) GetStampBalanceOk() (*[]ResponseStampCatBal, bool)`

GetStampBalanceOk returns a tuple with the StampBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampBalance

`func (o *ResponseSummaryData) SetStampBalance(v []ResponseStampCatBal)`

SetStampBalance sets StampBalance field to given value.

### HasStampBalance

`func (o *ResponseSummaryData) HasStampBalance() bool`

HasStampBalance returns a boolean if a field has been set.

### GetSubOrdinateOfficeSummarydetails

`func (o *ResponseSummaryData) GetSubOrdinateOfficeSummarydetails() []ResponseSubOrdinateOfficeSummarydetails`

GetSubOrdinateOfficeSummarydetails returns the SubOrdinateOfficeSummarydetails field if non-nil, zero value otherwise.

### GetSubOrdinateOfficeSummarydetailsOk

`func (o *ResponseSummaryData) GetSubOrdinateOfficeSummarydetailsOk() (*[]ResponseSubOrdinateOfficeSummarydetails, bool)`

GetSubOrdinateOfficeSummarydetailsOk returns a tuple with the SubOrdinateOfficeSummarydetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubOrdinateOfficeSummarydetails

`func (o *ResponseSummaryData) SetSubOrdinateOfficeSummarydetails(v []ResponseSubOrdinateOfficeSummarydetails)`

SetSubOrdinateOfficeSummarydetails sets SubOrdinateOfficeSummarydetails field to given value.

### HasSubOrdinateOfficeSummarydetails

`func (o *ResponseSummaryData) HasSubOrdinateOfficeSummarydetails() bool`

HasSubOrdinateOfficeSummarydetails returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseSummaryData) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseSummaryData) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseSummaryData) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseSummaryData) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransitCash

`func (o *ResponseSummaryData) GetTransitCash() float32`

GetTransitCash returns the TransitCash field if non-nil, zero value otherwise.

### GetTransitCashOk

`func (o *ResponseSummaryData) GetTransitCashOk() (*float32, bool)`

GetTransitCashOk returns a tuple with the TransitCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitCash

`func (o *ResponseSummaryData) SetTransitCash(v float32)`

SetTransitCash sets TransitCash field to given value.

### HasTransitCash

`func (o *ResponseSummaryData) HasTransitCash() bool`

HasTransitCash returns a boolean if a field has been set.

### GetTransitIpos

`func (o *ResponseSummaryData) GetTransitIpos() float32`

GetTransitIpos returns the TransitIpos field if non-nil, zero value otherwise.

### GetTransitIposOk

`func (o *ResponseSummaryData) GetTransitIposOk() (*float32, bool)`

GetTransitIposOk returns a tuple with the TransitIpos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitIpos

`func (o *ResponseSummaryData) SetTransitIpos(v float32)`

SetTransitIpos sets TransitIpos field to given value.

### HasTransitIpos

`func (o *ResponseSummaryData) HasTransitIpos() bool`

HasTransitIpos returns a boolean if a field has been set.

### GetTransitStamps

`func (o *ResponseSummaryData) GetTransitStamps() float32`

GetTransitStamps returns the TransitStamps field if non-nil, zero value otherwise.

### GetTransitStampsOk

`func (o *ResponseSummaryData) GetTransitStampsOk() (*float32, bool)`

GetTransitStampsOk returns a tuple with the TransitStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitStamps

`func (o *ResponseSummaryData) SetTransitStamps(v float32)`

SetTransitStamps sets TransitStamps field to given value.

### HasTransitStamps

`func (o *ResponseSummaryData) HasTransitStamps() bool`

HasTransitStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



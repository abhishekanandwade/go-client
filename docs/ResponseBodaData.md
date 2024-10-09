# ResponseBodaData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountingDetail** | Pointer to [**[]ResponseAccountingDetails**](ResponseAccountingDetails.md) |  | [optional] 
**CashBagDetail** | Pointer to [**[]ResponseCashBagDetail**](ResponseCashBagDetail.md) |  | [optional] 
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
**TransDate** | Pointer to **string** |  | [optional] 
**TransitCash** | Pointer to **float32** |  | [optional] 
**TransitStamps** | Pointer to **float32** |  | [optional] 

## Methods

### NewResponseBodaData

`func NewResponseBodaData() *ResponseBodaData`

NewResponseBodaData instantiates a new ResponseBodaData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseBodaDataWithDefaults

`func NewResponseBodaDataWithDefaults() *ResponseBodaData`

NewResponseBodaDataWithDefaults instantiates a new ResponseBodaData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountingDetail

`func (o *ResponseBodaData) GetAccountingDetail() []ResponseAccountingDetails`

GetAccountingDetail returns the AccountingDetail field if non-nil, zero value otherwise.

### GetAccountingDetailOk

`func (o *ResponseBodaData) GetAccountingDetailOk() (*[]ResponseAccountingDetails, bool)`

GetAccountingDetailOk returns a tuple with the AccountingDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingDetail

`func (o *ResponseBodaData) SetAccountingDetail(v []ResponseAccountingDetails)`

SetAccountingDetail sets AccountingDetail field to given value.

### HasAccountingDetail

`func (o *ResponseBodaData) HasAccountingDetail() bool`

HasAccountingDetail returns a boolean if a field has been set.

### GetCashBagDetail

`func (o *ResponseBodaData) GetCashBagDetail() []ResponseCashBagDetail`

GetCashBagDetail returns the CashBagDetail field if non-nil, zero value otherwise.

### GetCashBagDetailOk

`func (o *ResponseBodaData) GetCashBagDetailOk() (*[]ResponseCashBagDetail, bool)`

GetCashBagDetailOk returns a tuple with the CashBagDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBagDetail

`func (o *ResponseBodaData) SetCashBagDetail(v []ResponseCashBagDetail)`

SetCashBagDetail sets CashBagDetail field to given value.

### HasCashBagDetail

`func (o *ResponseBodaData) HasCashBagDetail() bool`

HasCashBagDetail returns a boolean if a field has been set.

### GetClosingBal

`func (o *ResponseBodaData) GetClosingBal() float32`

GetClosingBal returns the ClosingBal field if non-nil, zero value otherwise.

### GetClosingBalOk

`func (o *ResponseBodaData) GetClosingBalOk() (*float32, bool)`

GetClosingBalOk returns a tuple with the ClosingBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBal

`func (o *ResponseBodaData) SetClosingBal(v float32)`

SetClosingBal sets ClosingBal field to given value.

### HasClosingBal

`func (o *ResponseBodaData) HasClosingBal() bool`

HasClosingBal returns a boolean if a field has been set.

### GetDayendStatus

`func (o *ResponseBodaData) GetDayendStatus() int32`

GetDayendStatus returns the DayendStatus field if non-nil, zero value otherwise.

### GetDayendStatusOk

`func (o *ResponseBodaData) GetDayendStatusOk() (*int32, bool)`

GetDayendStatusOk returns a tuple with the DayendStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayendStatus

`func (o *ResponseBodaData) SetDayendStatus(v int32)`

SetDayendStatus sets DayendStatus field to given value.

### HasDayendStatus

`func (o *ResponseBodaData) HasDayendStatus() bool`

HasDayendStatus returns a boolean if a field has been set.

### GetIpobalance

`func (o *ResponseBodaData) GetIpobalance() []ResponseIPOBal`

GetIpobalance returns the Ipobalance field if non-nil, zero value otherwise.

### GetIpobalanceOk

`func (o *ResponseBodaData) GetIpobalanceOk() (*[]ResponseIPOBal, bool)`

GetIpobalanceOk returns a tuple with the Ipobalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpobalance

`func (o *ResponseBodaData) SetIpobalance(v []ResponseIPOBal)`

SetIpobalance sets Ipobalance field to given value.

### HasIpobalance

`func (o *ResponseBodaData) HasIpobalance() bool`

HasIpobalance returns a boolean if a field has been set.

### GetMaxBal

`func (o *ResponseBodaData) GetMaxBal() float32`

GetMaxBal returns the MaxBal field if non-nil, zero value otherwise.

### GetMaxBalOk

`func (o *ResponseBodaData) GetMaxBalOk() (*float32, bool)`

GetMaxBalOk returns a tuple with the MaxBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBal

`func (o *ResponseBodaData) SetMaxBal(v float32)`

SetMaxBal sets MaxBal field to given value.

### HasMaxBal

`func (o *ResponseBodaData) HasMaxBal() bool`

HasMaxBal returns a boolean if a field has been set.

### GetMinBal

`func (o *ResponseBodaData) GetMinBal() float32`

GetMinBal returns the MinBal field if non-nil, zero value otherwise.

### GetMinBalOk

`func (o *ResponseBodaData) GetMinBalOk() (*float32, bool)`

GetMinBalOk returns a tuple with the MinBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBal

`func (o *ResponseBodaData) SetMinBal(v float32)`

SetMinBal sets MinBal field to given value.

### HasMinBal

`func (o *ResponseBodaData) HasMinBal() bool`

HasMinBal returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseBodaData) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseBodaData) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseBodaData) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseBodaData) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOpeningBal

`func (o *ResponseBodaData) GetOpeningBal() float32`

GetOpeningBal returns the OpeningBal field if non-nil, zero value otherwise.

### GetOpeningBalOk

`func (o *ResponseBodaData) GetOpeningBalOk() (*float32, bool)`

GetOpeningBalOk returns a tuple with the OpeningBal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBal

`func (o *ResponseBodaData) SetOpeningBal(v float32)`

SetOpeningBal sets OpeningBal field to given value.

### HasOpeningBal

`func (o *ResponseBodaData) HasOpeningBal() bool`

HasOpeningBal returns a boolean if a field has been set.

### GetPayments

`func (o *ResponseBodaData) GetPayments() float32`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ResponseBodaData) GetPaymentsOk() (*float32, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ResponseBodaData) SetPayments(v float32)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *ResponseBodaData) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetReceipts

`func (o *ResponseBodaData) GetReceipts() float32`

GetReceipts returns the Receipts field if non-nil, zero value otherwise.

### GetReceiptsOk

`func (o *ResponseBodaData) GetReceiptsOk() (*float32, bool)`

GetReceiptsOk returns a tuple with the Receipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceipts

`func (o *ResponseBodaData) SetReceipts(v float32)`

SetReceipts sets Receipts field to given value.

### HasReceipts

`func (o *ResponseBodaData) HasReceipts() bool`

HasReceipts returns a boolean if a field has been set.

### GetStampBalance

`func (o *ResponseBodaData) GetStampBalance() []ResponseStampCatBal`

GetStampBalance returns the StampBalance field if non-nil, zero value otherwise.

### GetStampBalanceOk

`func (o *ResponseBodaData) GetStampBalanceOk() (*[]ResponseStampCatBal, bool)`

GetStampBalanceOk returns a tuple with the StampBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampBalance

`func (o *ResponseBodaData) SetStampBalance(v []ResponseStampCatBal)`

SetStampBalance sets StampBalance field to given value.

### HasStampBalance

`func (o *ResponseBodaData) HasStampBalance() bool`

HasStampBalance returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseBodaData) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseBodaData) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseBodaData) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseBodaData) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransitCash

`func (o *ResponseBodaData) GetTransitCash() float32`

GetTransitCash returns the TransitCash field if non-nil, zero value otherwise.

### GetTransitCashOk

`func (o *ResponseBodaData) GetTransitCashOk() (*float32, bool)`

GetTransitCashOk returns a tuple with the TransitCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitCash

`func (o *ResponseBodaData) SetTransitCash(v float32)`

SetTransitCash sets TransitCash field to given value.

### HasTransitCash

`func (o *ResponseBodaData) HasTransitCash() bool`

HasTransitCash returns a boolean if a field has been set.

### GetTransitStamps

`func (o *ResponseBodaData) GetTransitStamps() float32`

GetTransitStamps returns the TransitStamps field if non-nil, zero value otherwise.

### GetTransitStampsOk

`func (o *ResponseBodaData) GetTransitStampsOk() (*float32, bool)`

GetTransitStampsOk returns a tuple with the TransitStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitStamps

`func (o *ResponseBodaData) SetTransitStamps(v float32)`

SetTransitStamps sets TransitStamps field to given value.

### HasTransitStamps

`func (o *ResponseBodaData) HasTransitStamps() bool`

HasTransitStamps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



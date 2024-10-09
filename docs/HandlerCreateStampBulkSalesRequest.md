# HandlerCreateStampBulkSalesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankBranch** | Pointer to **string** |  | [optional] 
**InstrumentDate** | Pointer to **string** |  | [optional] 
**InstrumentNo** | Pointer to **string** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**OfficeId** | **int32** |  | 
**PaymentMode** | **string** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**SaleAmt** | **float32** |  | 
**SaleDetails** | **map[string]int32** |  | 
**SoldTo** | **string** |  | 
**UserId** | **int32** |  | 

## Methods

### NewHandlerCreateStampBulkSalesRequest

`func NewHandlerCreateStampBulkSalesRequest(officeId int32, paymentMode string, saleAmt float32, saleDetails map[string]int32, soldTo string, userId int32, ) *HandlerCreateStampBulkSalesRequest`

NewHandlerCreateStampBulkSalesRequest instantiates a new HandlerCreateStampBulkSalesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateStampBulkSalesRequestWithDefaults

`func NewHandlerCreateStampBulkSalesRequestWithDefaults() *HandlerCreateStampBulkSalesRequest`

NewHandlerCreateStampBulkSalesRequestWithDefaults instantiates a new HandlerCreateStampBulkSalesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankBranch

`func (o *HandlerCreateStampBulkSalesRequest) GetBankBranch() string`

GetBankBranch returns the BankBranch field if non-nil, zero value otherwise.

### GetBankBranchOk

`func (o *HandlerCreateStampBulkSalesRequest) GetBankBranchOk() (*string, bool)`

GetBankBranchOk returns a tuple with the BankBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranch

`func (o *HandlerCreateStampBulkSalesRequest) SetBankBranch(v string)`

SetBankBranch sets BankBranch field to given value.

### HasBankBranch

`func (o *HandlerCreateStampBulkSalesRequest) HasBankBranch() bool`

HasBankBranch returns a boolean if a field has been set.

### GetInstrumentDate

`func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentDate() string`

GetInstrumentDate returns the InstrumentDate field if non-nil, zero value otherwise.

### GetInstrumentDateOk

`func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentDateOk() (*string, bool)`

GetInstrumentDateOk returns a tuple with the InstrumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDate

`func (o *HandlerCreateStampBulkSalesRequest) SetInstrumentDate(v string)`

SetInstrumentDate sets InstrumentDate field to given value.

### HasInstrumentDate

`func (o *HandlerCreateStampBulkSalesRequest) HasInstrumentDate() bool`

HasInstrumentDate returns a boolean if a field has been set.

### GetInstrumentNo

`func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentNo() string`

GetInstrumentNo returns the InstrumentNo field if non-nil, zero value otherwise.

### GetInstrumentNoOk

`func (o *HandlerCreateStampBulkSalesRequest) GetInstrumentNoOk() (*string, bool)`

GetInstrumentNoOk returns a tuple with the InstrumentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentNo

`func (o *HandlerCreateStampBulkSalesRequest) SetInstrumentNo(v string)`

SetInstrumentNo sets InstrumentNo field to given value.

### HasInstrumentNo

`func (o *HandlerCreateStampBulkSalesRequest) HasInstrumentNo() bool`

HasInstrumentNo returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *HandlerCreateStampBulkSalesRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerCreateStampBulkSalesRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerCreateStampBulkSalesRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerCreateStampBulkSalesRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerCreateStampBulkSalesRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateStampBulkSalesRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateStampBulkSalesRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetPaymentMode

`func (o *HandlerCreateStampBulkSalesRequest) GetPaymentMode() string`

GetPaymentMode returns the PaymentMode field if non-nil, zero value otherwise.

### GetPaymentModeOk

`func (o *HandlerCreateStampBulkSalesRequest) GetPaymentModeOk() (*string, bool)`

GetPaymentModeOk returns a tuple with the PaymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMode

`func (o *HandlerCreateStampBulkSalesRequest) SetPaymentMode(v string)`

SetPaymentMode sets PaymentMode field to given value.


### GetRemarks

`func (o *HandlerCreateStampBulkSalesRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateStampBulkSalesRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateStampBulkSalesRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateStampBulkSalesRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSaleAmt

`func (o *HandlerCreateStampBulkSalesRequest) GetSaleAmt() float32`

GetSaleAmt returns the SaleAmt field if non-nil, zero value otherwise.

### GetSaleAmtOk

`func (o *HandlerCreateStampBulkSalesRequest) GetSaleAmtOk() (*float32, bool)`

GetSaleAmtOk returns a tuple with the SaleAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmt

`func (o *HandlerCreateStampBulkSalesRequest) SetSaleAmt(v float32)`

SetSaleAmt sets SaleAmt field to given value.


### GetSaleDetails

`func (o *HandlerCreateStampBulkSalesRequest) GetSaleDetails() map[string]int32`

GetSaleDetails returns the SaleDetails field if non-nil, zero value otherwise.

### GetSaleDetailsOk

`func (o *HandlerCreateStampBulkSalesRequest) GetSaleDetailsOk() (*map[string]int32, bool)`

GetSaleDetailsOk returns a tuple with the SaleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDetails

`func (o *HandlerCreateStampBulkSalesRequest) SetSaleDetails(v map[string]int32)`

SetSaleDetails sets SaleDetails field to given value.


### GetSoldTo

`func (o *HandlerCreateStampBulkSalesRequest) GetSoldTo() string`

GetSoldTo returns the SoldTo field if non-nil, zero value otherwise.

### GetSoldToOk

`func (o *HandlerCreateStampBulkSalesRequest) GetSoldToOk() (*string, bool)`

GetSoldToOk returns a tuple with the SoldTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldTo

`func (o *HandlerCreateStampBulkSalesRequest) SetSoldTo(v string)`

SetSoldTo sets SoldTo field to given value.


### GetUserId

`func (o *HandlerCreateStampBulkSalesRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerCreateStampBulkSalesRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerCreateStampBulkSalesRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



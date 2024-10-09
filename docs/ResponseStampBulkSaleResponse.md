# ResponseStampBulkSaleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedDate** | Pointer to **string** |  | [optional] 
**ApproverId** | Pointer to **int32** |  | [optional] 
**BankBranch** | Pointer to **string** |  | [optional] 
**ChqRealiseDate** | Pointer to **string** |  | [optional] 
**InstrumentDate** | Pointer to **string** |  | [optional] 
**InstrumentNo** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsRealized** | Pointer to **bool** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**PaymentMode** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SaleAmt** | Pointer to **float32** |  | [optional] 
**SaleDetails** | Pointer to **map[string]int32** |  | [optional] 
**SoldTo** | Pointer to **string** |  | [optional] 
**StampDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TranId** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStampBulkSaleResponse

`func NewResponseStampBulkSaleResponse() *ResponseStampBulkSaleResponse`

NewResponseStampBulkSaleResponse instantiates a new ResponseStampBulkSaleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampBulkSaleResponseWithDefaults

`func NewResponseStampBulkSaleResponseWithDefaults() *ResponseStampBulkSaleResponse`

NewResponseStampBulkSaleResponseWithDefaults instantiates a new ResponseStampBulkSaleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedDate

`func (o *ResponseStampBulkSaleResponse) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ResponseStampBulkSaleResponse) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ResponseStampBulkSaleResponse) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ResponseStampBulkSaleResponse) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetApproverId

`func (o *ResponseStampBulkSaleResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ResponseStampBulkSaleResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ResponseStampBulkSaleResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ResponseStampBulkSaleResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetBankBranch

`func (o *ResponseStampBulkSaleResponse) GetBankBranch() string`

GetBankBranch returns the BankBranch field if non-nil, zero value otherwise.

### GetBankBranchOk

`func (o *ResponseStampBulkSaleResponse) GetBankBranchOk() (*string, bool)`

GetBankBranchOk returns a tuple with the BankBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranch

`func (o *ResponseStampBulkSaleResponse) SetBankBranch(v string)`

SetBankBranch sets BankBranch field to given value.

### HasBankBranch

`func (o *ResponseStampBulkSaleResponse) HasBankBranch() bool`

HasBankBranch returns a boolean if a field has been set.

### GetChqRealiseDate

`func (o *ResponseStampBulkSaleResponse) GetChqRealiseDate() string`

GetChqRealiseDate returns the ChqRealiseDate field if non-nil, zero value otherwise.

### GetChqRealiseDateOk

`func (o *ResponseStampBulkSaleResponse) GetChqRealiseDateOk() (*string, bool)`

GetChqRealiseDateOk returns a tuple with the ChqRealiseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChqRealiseDate

`func (o *ResponseStampBulkSaleResponse) SetChqRealiseDate(v string)`

SetChqRealiseDate sets ChqRealiseDate field to given value.

### HasChqRealiseDate

`func (o *ResponseStampBulkSaleResponse) HasChqRealiseDate() bool`

HasChqRealiseDate returns a boolean if a field has been set.

### GetInstrumentDate

`func (o *ResponseStampBulkSaleResponse) GetInstrumentDate() string`

GetInstrumentDate returns the InstrumentDate field if non-nil, zero value otherwise.

### GetInstrumentDateOk

`func (o *ResponseStampBulkSaleResponse) GetInstrumentDateOk() (*string, bool)`

GetInstrumentDateOk returns a tuple with the InstrumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDate

`func (o *ResponseStampBulkSaleResponse) SetInstrumentDate(v string)`

SetInstrumentDate sets InstrumentDate field to given value.

### HasInstrumentDate

`func (o *ResponseStampBulkSaleResponse) HasInstrumentDate() bool`

HasInstrumentDate returns a boolean if a field has been set.

### GetInstrumentNo

`func (o *ResponseStampBulkSaleResponse) GetInstrumentNo() string`

GetInstrumentNo returns the InstrumentNo field if non-nil, zero value otherwise.

### GetInstrumentNoOk

`func (o *ResponseStampBulkSaleResponse) GetInstrumentNoOk() (*string, bool)`

GetInstrumentNoOk returns a tuple with the InstrumentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentNo

`func (o *ResponseStampBulkSaleResponse) SetInstrumentNo(v string)`

SetInstrumentNo sets InstrumentNo field to given value.

### HasInstrumentNo

`func (o *ResponseStampBulkSaleResponse) HasInstrumentNo() bool`

HasInstrumentNo returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampBulkSaleResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampBulkSaleResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampBulkSaleResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampBulkSaleResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsRealized

`func (o *ResponseStampBulkSaleResponse) GetIsRealized() bool`

GetIsRealized returns the IsRealized field if non-nil, zero value otherwise.

### GetIsRealizedOk

`func (o *ResponseStampBulkSaleResponse) GetIsRealizedOk() (*bool, bool)`

GetIsRealizedOk returns a tuple with the IsRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealized

`func (o *ResponseStampBulkSaleResponse) SetIsRealized(v bool)`

SetIsRealized sets IsRealized field to given value.

### HasIsRealized

`func (o *ResponseStampBulkSaleResponse) HasIsRealized() bool`

HasIsRealized returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseStampBulkSaleResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseStampBulkSaleResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseStampBulkSaleResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseStampBulkSaleResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetPaymentMode

`func (o *ResponseStampBulkSaleResponse) GetPaymentMode() string`

GetPaymentMode returns the PaymentMode field if non-nil, zero value otherwise.

### GetPaymentModeOk

`func (o *ResponseStampBulkSaleResponse) GetPaymentModeOk() (*string, bool)`

GetPaymentModeOk returns a tuple with the PaymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMode

`func (o *ResponseStampBulkSaleResponse) SetPaymentMode(v string)`

SetPaymentMode sets PaymentMode field to given value.

### HasPaymentMode

`func (o *ResponseStampBulkSaleResponse) HasPaymentMode() bool`

HasPaymentMode returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampBulkSaleResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampBulkSaleResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampBulkSaleResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampBulkSaleResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSaleAmt

`func (o *ResponseStampBulkSaleResponse) GetSaleAmt() float32`

GetSaleAmt returns the SaleAmt field if non-nil, zero value otherwise.

### GetSaleAmtOk

`func (o *ResponseStampBulkSaleResponse) GetSaleAmtOk() (*float32, bool)`

GetSaleAmtOk returns a tuple with the SaleAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmt

`func (o *ResponseStampBulkSaleResponse) SetSaleAmt(v float32)`

SetSaleAmt sets SaleAmt field to given value.

### HasSaleAmt

`func (o *ResponseStampBulkSaleResponse) HasSaleAmt() bool`

HasSaleAmt returns a boolean if a field has been set.

### GetSaleDetails

`func (o *ResponseStampBulkSaleResponse) GetSaleDetails() map[string]int32`

GetSaleDetails returns the SaleDetails field if non-nil, zero value otherwise.

### GetSaleDetailsOk

`func (o *ResponseStampBulkSaleResponse) GetSaleDetailsOk() (*map[string]int32, bool)`

GetSaleDetailsOk returns a tuple with the SaleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDetails

`func (o *ResponseStampBulkSaleResponse) SetSaleDetails(v map[string]int32)`

SetSaleDetails sets SaleDetails field to given value.

### HasSaleDetails

`func (o *ResponseStampBulkSaleResponse) HasSaleDetails() bool`

HasSaleDetails returns a boolean if a field has been set.

### GetSoldTo

`func (o *ResponseStampBulkSaleResponse) GetSoldTo() string`

GetSoldTo returns the SoldTo field if non-nil, zero value otherwise.

### GetSoldToOk

`func (o *ResponseStampBulkSaleResponse) GetSoldToOk() (*string, bool)`

GetSoldToOk returns a tuple with the SoldTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldTo

`func (o *ResponseStampBulkSaleResponse) SetSoldTo(v string)`

SetSoldTo sets SoldTo field to given value.

### HasSoldTo

`func (o *ResponseStampBulkSaleResponse) HasSoldTo() bool`

HasSoldTo returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *ResponseStampBulkSaleResponse) GetStampDetailsDesc() []ResponseStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *ResponseStampBulkSaleResponse) GetStampDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *ResponseStampBulkSaleResponse) SetStampDetailsDesc(v []ResponseStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *ResponseStampBulkSaleResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTranId

`func (o *ResponseStampBulkSaleResponse) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *ResponseStampBulkSaleResponse) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *ResponseStampBulkSaleResponse) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *ResponseStampBulkSaleResponse) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampBulkSaleResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampBulkSaleResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampBulkSaleResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampBulkSaleResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseStampBulkSaleResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseStampBulkSaleResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseStampBulkSaleResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseStampBulkSaleResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerStampBulkSaleResponse

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
**StampDetailsDesc** | Pointer to [**[]HandlerStampdetails**](HandlerStampdetails.md) |  | [optional] 
**TranId** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlerStampBulkSaleResponse

`func NewHandlerStampBulkSaleResponse() *HandlerStampBulkSaleResponse`

NewHandlerStampBulkSaleResponse instantiates a new HandlerStampBulkSaleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerStampBulkSaleResponseWithDefaults

`func NewHandlerStampBulkSaleResponseWithDefaults() *HandlerStampBulkSaleResponse`

NewHandlerStampBulkSaleResponseWithDefaults instantiates a new HandlerStampBulkSaleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedDate

`func (o *HandlerStampBulkSaleResponse) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *HandlerStampBulkSaleResponse) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *HandlerStampBulkSaleResponse) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *HandlerStampBulkSaleResponse) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetApproverId

`func (o *HandlerStampBulkSaleResponse) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *HandlerStampBulkSaleResponse) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *HandlerStampBulkSaleResponse) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *HandlerStampBulkSaleResponse) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetBankBranch

`func (o *HandlerStampBulkSaleResponse) GetBankBranch() string`

GetBankBranch returns the BankBranch field if non-nil, zero value otherwise.

### GetBankBranchOk

`func (o *HandlerStampBulkSaleResponse) GetBankBranchOk() (*string, bool)`

GetBankBranchOk returns a tuple with the BankBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBranch

`func (o *HandlerStampBulkSaleResponse) SetBankBranch(v string)`

SetBankBranch sets BankBranch field to given value.

### HasBankBranch

`func (o *HandlerStampBulkSaleResponse) HasBankBranch() bool`

HasBankBranch returns a boolean if a field has been set.

### GetChqRealiseDate

`func (o *HandlerStampBulkSaleResponse) GetChqRealiseDate() string`

GetChqRealiseDate returns the ChqRealiseDate field if non-nil, zero value otherwise.

### GetChqRealiseDateOk

`func (o *HandlerStampBulkSaleResponse) GetChqRealiseDateOk() (*string, bool)`

GetChqRealiseDateOk returns a tuple with the ChqRealiseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChqRealiseDate

`func (o *HandlerStampBulkSaleResponse) SetChqRealiseDate(v string)`

SetChqRealiseDate sets ChqRealiseDate field to given value.

### HasChqRealiseDate

`func (o *HandlerStampBulkSaleResponse) HasChqRealiseDate() bool`

HasChqRealiseDate returns a boolean if a field has been set.

### GetInstrumentDate

`func (o *HandlerStampBulkSaleResponse) GetInstrumentDate() string`

GetInstrumentDate returns the InstrumentDate field if non-nil, zero value otherwise.

### GetInstrumentDateOk

`func (o *HandlerStampBulkSaleResponse) GetInstrumentDateOk() (*string, bool)`

GetInstrumentDateOk returns a tuple with the InstrumentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentDate

`func (o *HandlerStampBulkSaleResponse) SetInstrumentDate(v string)`

SetInstrumentDate sets InstrumentDate field to given value.

### HasInstrumentDate

`func (o *HandlerStampBulkSaleResponse) HasInstrumentDate() bool`

HasInstrumentDate returns a boolean if a field has been set.

### GetInstrumentNo

`func (o *HandlerStampBulkSaleResponse) GetInstrumentNo() string`

GetInstrumentNo returns the InstrumentNo field if non-nil, zero value otherwise.

### GetInstrumentNoOk

`func (o *HandlerStampBulkSaleResponse) GetInstrumentNoOk() (*string, bool)`

GetInstrumentNoOk returns a tuple with the InstrumentNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentNo

`func (o *HandlerStampBulkSaleResponse) SetInstrumentNo(v string)`

SetInstrumentNo sets InstrumentNo field to given value.

### HasInstrumentNo

`func (o *HandlerStampBulkSaleResponse) HasInstrumentNo() bool`

HasInstrumentNo returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerStampBulkSaleResponse) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerStampBulkSaleResponse) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerStampBulkSaleResponse) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerStampBulkSaleResponse) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsRealized

`func (o *HandlerStampBulkSaleResponse) GetIsRealized() bool`

GetIsRealized returns the IsRealized field if non-nil, zero value otherwise.

### GetIsRealizedOk

`func (o *HandlerStampBulkSaleResponse) GetIsRealizedOk() (*bool, bool)`

GetIsRealizedOk returns a tuple with the IsRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealized

`func (o *HandlerStampBulkSaleResponse) SetIsRealized(v bool)`

SetIsRealized sets IsRealized field to given value.

### HasIsRealized

`func (o *HandlerStampBulkSaleResponse) HasIsRealized() bool`

HasIsRealized returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerStampBulkSaleResponse) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerStampBulkSaleResponse) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerStampBulkSaleResponse) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *HandlerStampBulkSaleResponse) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetPaymentMode

`func (o *HandlerStampBulkSaleResponse) GetPaymentMode() string`

GetPaymentMode returns the PaymentMode field if non-nil, zero value otherwise.

### GetPaymentModeOk

`func (o *HandlerStampBulkSaleResponse) GetPaymentModeOk() (*string, bool)`

GetPaymentModeOk returns a tuple with the PaymentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMode

`func (o *HandlerStampBulkSaleResponse) SetPaymentMode(v string)`

SetPaymentMode sets PaymentMode field to given value.

### HasPaymentMode

`func (o *HandlerStampBulkSaleResponse) HasPaymentMode() bool`

HasPaymentMode returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerStampBulkSaleResponse) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerStampBulkSaleResponse) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerStampBulkSaleResponse) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerStampBulkSaleResponse) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSaleAmt

`func (o *HandlerStampBulkSaleResponse) GetSaleAmt() float32`

GetSaleAmt returns the SaleAmt field if non-nil, zero value otherwise.

### GetSaleAmtOk

`func (o *HandlerStampBulkSaleResponse) GetSaleAmtOk() (*float32, bool)`

GetSaleAmtOk returns a tuple with the SaleAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmt

`func (o *HandlerStampBulkSaleResponse) SetSaleAmt(v float32)`

SetSaleAmt sets SaleAmt field to given value.

### HasSaleAmt

`func (o *HandlerStampBulkSaleResponse) HasSaleAmt() bool`

HasSaleAmt returns a boolean if a field has been set.

### GetSaleDetails

`func (o *HandlerStampBulkSaleResponse) GetSaleDetails() map[string]int32`

GetSaleDetails returns the SaleDetails field if non-nil, zero value otherwise.

### GetSaleDetailsOk

`func (o *HandlerStampBulkSaleResponse) GetSaleDetailsOk() (*map[string]int32, bool)`

GetSaleDetailsOk returns a tuple with the SaleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDetails

`func (o *HandlerStampBulkSaleResponse) SetSaleDetails(v map[string]int32)`

SetSaleDetails sets SaleDetails field to given value.

### HasSaleDetails

`func (o *HandlerStampBulkSaleResponse) HasSaleDetails() bool`

HasSaleDetails returns a boolean if a field has been set.

### GetSoldTo

`func (o *HandlerStampBulkSaleResponse) GetSoldTo() string`

GetSoldTo returns the SoldTo field if non-nil, zero value otherwise.

### GetSoldToOk

`func (o *HandlerStampBulkSaleResponse) GetSoldToOk() (*string, bool)`

GetSoldToOk returns a tuple with the SoldTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldTo

`func (o *HandlerStampBulkSaleResponse) SetSoldTo(v string)`

SetSoldTo sets SoldTo field to given value.

### HasSoldTo

`func (o *HandlerStampBulkSaleResponse) HasSoldTo() bool`

HasSoldTo returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *HandlerStampBulkSaleResponse) GetStampDetailsDesc() []HandlerStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *HandlerStampBulkSaleResponse) GetStampDetailsDescOk() (*[]HandlerStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *HandlerStampBulkSaleResponse) SetStampDetailsDesc(v []HandlerStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *HandlerStampBulkSaleResponse) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTranId

`func (o *HandlerStampBulkSaleResponse) GetTranId() string`

GetTranId returns the TranId field if non-nil, zero value otherwise.

### GetTranIdOk

`func (o *HandlerStampBulkSaleResponse) GetTranIdOk() (*string, bool)`

GetTranIdOk returns a tuple with the TranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranId

`func (o *HandlerStampBulkSaleResponse) SetTranId(v string)`

SetTranId sets TranId field to given value.

### HasTranId

`func (o *HandlerStampBulkSaleResponse) HasTranId() bool`

HasTranId returns a boolean if a field has been set.

### GetTransDate

`func (o *HandlerStampBulkSaleResponse) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *HandlerStampBulkSaleResponse) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *HandlerStampBulkSaleResponse) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *HandlerStampBulkSaleResponse) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerStampBulkSaleResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerStampBulkSaleResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerStampBulkSaleResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HandlerStampBulkSaleResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



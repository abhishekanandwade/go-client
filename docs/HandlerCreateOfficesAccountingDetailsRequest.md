# HandlerCreateOfficesAccountingDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditOrDebit** | Pointer to **string** |  | [optional] 
**DigitalTxnsAmount** | Pointer to **float32** |  | [optional] 
**DigitalTxnsCount** | Pointer to **int32** |  | [optional] 
**GlCode** | Pointer to **string** |  | [optional] 
**NonDigitalTxnsAmount** | Pointer to **float32** |  | [optional] 
**NonDigitalTxnsCount** | Pointer to **int32** |  | [optional] 
**ReceiptSource** | **string** |  | 
**Remarks** | Pointer to **string** |  | [optional] 
**SrcTranId** | **string** |  | 
**SrcTransDate** | **string** |  | 

## Methods

### NewHandlerCreateOfficesAccountingDetailsRequest

`func NewHandlerCreateOfficesAccountingDetailsRequest(receiptSource string, srcTranId string, srcTransDate string, ) *HandlerCreateOfficesAccountingDetailsRequest`

NewHandlerCreateOfficesAccountingDetailsRequest instantiates a new HandlerCreateOfficesAccountingDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateOfficesAccountingDetailsRequestWithDefaults

`func NewHandlerCreateOfficesAccountingDetailsRequestWithDefaults() *HandlerCreateOfficesAccountingDetailsRequest`

NewHandlerCreateOfficesAccountingDetailsRequestWithDefaults instantiates a new HandlerCreateOfficesAccountingDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditOrDebit

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetCreditOrDebit() string`

GetCreditOrDebit returns the CreditOrDebit field if non-nil, zero value otherwise.

### GetCreditOrDebitOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetCreditOrDebitOk() (*string, bool)`

GetCreditOrDebitOk returns a tuple with the CreditOrDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditOrDebit

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetCreditOrDebit(v string)`

SetCreditOrDebit sets CreditOrDebit field to given value.

### HasCreditOrDebit

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasCreditOrDebit() bool`

HasCreditOrDebit returns a boolean if a field has been set.

### GetDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetDigitalTxnsAmount() float32`

GetDigitalTxnsAmount returns the DigitalTxnsAmount field if non-nil, zero value otherwise.

### GetDigitalTxnsAmountOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetDigitalTxnsAmountOk() (*float32, bool)`

GetDigitalTxnsAmountOk returns a tuple with the DigitalTxnsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetDigitalTxnsAmount(v float32)`

SetDigitalTxnsAmount sets DigitalTxnsAmount field to given value.

### HasDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasDigitalTxnsAmount() bool`

HasDigitalTxnsAmount returns a boolean if a field has been set.

### GetDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetDigitalTxnsCount() int32`

GetDigitalTxnsCount returns the DigitalTxnsCount field if non-nil, zero value otherwise.

### GetDigitalTxnsCountOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetDigitalTxnsCountOk() (*int32, bool)`

GetDigitalTxnsCountOk returns a tuple with the DigitalTxnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetDigitalTxnsCount(v int32)`

SetDigitalTxnsCount sets DigitalTxnsCount field to given value.

### HasDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasDigitalTxnsCount() bool`

HasDigitalTxnsCount returns a boolean if a field has been set.

### GetGlCode

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetGlCode() string`

GetGlCode returns the GlCode field if non-nil, zero value otherwise.

### GetGlCodeOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetGlCodeOk() (*string, bool)`

GetGlCodeOk returns a tuple with the GlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlCode

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetGlCode(v string)`

SetGlCode sets GlCode field to given value.

### HasGlCode

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasGlCode() bool`

HasGlCode returns a boolean if a field has been set.

### GetNonDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetNonDigitalTxnsAmount() float32`

GetNonDigitalTxnsAmount returns the NonDigitalTxnsAmount field if non-nil, zero value otherwise.

### GetNonDigitalTxnsAmountOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetNonDigitalTxnsAmountOk() (*float32, bool)`

GetNonDigitalTxnsAmountOk returns a tuple with the NonDigitalTxnsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetNonDigitalTxnsAmount(v float32)`

SetNonDigitalTxnsAmount sets NonDigitalTxnsAmount field to given value.

### HasNonDigitalTxnsAmount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasNonDigitalTxnsAmount() bool`

HasNonDigitalTxnsAmount returns a boolean if a field has been set.

### GetNonDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetNonDigitalTxnsCount() int32`

GetNonDigitalTxnsCount returns the NonDigitalTxnsCount field if non-nil, zero value otherwise.

### GetNonDigitalTxnsCountOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetNonDigitalTxnsCountOk() (*int32, bool)`

GetNonDigitalTxnsCountOk returns a tuple with the NonDigitalTxnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetNonDigitalTxnsCount(v int32)`

SetNonDigitalTxnsCount sets NonDigitalTxnsCount field to given value.

### HasNonDigitalTxnsCount

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasNonDigitalTxnsCount() bool`

HasNonDigitalTxnsCount returns a boolean if a field has been set.

### GetReceiptSource

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.


### GetRemarks

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateOfficesAccountingDetailsRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSrcTranId

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetSrcTranId() string`

GetSrcTranId returns the SrcTranId field if non-nil, zero value otherwise.

### GetSrcTranIdOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetSrcTranIdOk() (*string, bool)`

GetSrcTranIdOk returns a tuple with the SrcTranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTranId

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetSrcTranId(v string)`

SetSrcTranId sets SrcTranId field to given value.


### GetSrcTransDate

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetSrcTransDate() string`

GetSrcTransDate returns the SrcTransDate field if non-nil, zero value otherwise.

### GetSrcTransDateOk

`func (o *HandlerCreateOfficesAccountingDetailsRequest) GetSrcTransDateOk() (*string, bool)`

GetSrcTransDateOk returns a tuple with the SrcTransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTransDate

`func (o *HandlerCreateOfficesAccountingDetailsRequest) SetSrcTransDate(v string)`

SetSrcTransDate sets SrcTransDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



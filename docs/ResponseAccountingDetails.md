# ResponseAccountingDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCodeDescription** | Pointer to **string** |  | [optional] 
**CreditOrDebit** | Pointer to **string** |  | [optional] 
**DigitalTxnsAmount** | Pointer to **float32** |  | [optional] 
**DigitalTxnsCount** | Pointer to **int32** |  | [optional] 
**GlCode** | Pointer to **string** |  | [optional] 
**Hoa** | Pointer to **string** |  | [optional] 
**NonDigitalTxnsAmount** | Pointer to **float32** |  | [optional] 
**NonDigitalTxnsCount** | Pointer to **int32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**Part** | Pointer to **string** |  | [optional] 
**PositiveOrNegative** | Pointer to **string** |  | [optional] 
**ReceiptOrPmt** | Pointer to **string** |  | [optional] 
**ReceiptSource** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**SrcTranId** | Pointer to **string** |  | [optional] 
**SrcTransDate** | Pointer to **string** |  | [optional] 
**TotAmount** | Pointer to **float32** |  | [optional] 
**TotTxns** | Pointer to **int32** |  | [optional] 
**TreasuryTranId** | Pointer to **string** |  | [optional] 
**TryPostingDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseAccountingDetails

`func NewResponseAccountingDetails() *ResponseAccountingDetails`

NewResponseAccountingDetails instantiates a new ResponseAccountingDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseAccountingDetailsWithDefaults

`func NewResponseAccountingDetailsWithDefaults() *ResponseAccountingDetails`

NewResponseAccountingDetailsWithDefaults instantiates a new ResponseAccountingDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCodeDescription

`func (o *ResponseAccountingDetails) GetAccountCodeDescription() string`

GetAccountCodeDescription returns the AccountCodeDescription field if non-nil, zero value otherwise.

### GetAccountCodeDescriptionOk

`func (o *ResponseAccountingDetails) GetAccountCodeDescriptionOk() (*string, bool)`

GetAccountCodeDescriptionOk returns a tuple with the AccountCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCodeDescription

`func (o *ResponseAccountingDetails) SetAccountCodeDescription(v string)`

SetAccountCodeDescription sets AccountCodeDescription field to given value.

### HasAccountCodeDescription

`func (o *ResponseAccountingDetails) HasAccountCodeDescription() bool`

HasAccountCodeDescription returns a boolean if a field has been set.

### GetCreditOrDebit

`func (o *ResponseAccountingDetails) GetCreditOrDebit() string`

GetCreditOrDebit returns the CreditOrDebit field if non-nil, zero value otherwise.

### GetCreditOrDebitOk

`func (o *ResponseAccountingDetails) GetCreditOrDebitOk() (*string, bool)`

GetCreditOrDebitOk returns a tuple with the CreditOrDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditOrDebit

`func (o *ResponseAccountingDetails) SetCreditOrDebit(v string)`

SetCreditOrDebit sets CreditOrDebit field to given value.

### HasCreditOrDebit

`func (o *ResponseAccountingDetails) HasCreditOrDebit() bool`

HasCreditOrDebit returns a boolean if a field has been set.

### GetDigitalTxnsAmount

`func (o *ResponseAccountingDetails) GetDigitalTxnsAmount() float32`

GetDigitalTxnsAmount returns the DigitalTxnsAmount field if non-nil, zero value otherwise.

### GetDigitalTxnsAmountOk

`func (o *ResponseAccountingDetails) GetDigitalTxnsAmountOk() (*float32, bool)`

GetDigitalTxnsAmountOk returns a tuple with the DigitalTxnsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalTxnsAmount

`func (o *ResponseAccountingDetails) SetDigitalTxnsAmount(v float32)`

SetDigitalTxnsAmount sets DigitalTxnsAmount field to given value.

### HasDigitalTxnsAmount

`func (o *ResponseAccountingDetails) HasDigitalTxnsAmount() bool`

HasDigitalTxnsAmount returns a boolean if a field has been set.

### GetDigitalTxnsCount

`func (o *ResponseAccountingDetails) GetDigitalTxnsCount() int32`

GetDigitalTxnsCount returns the DigitalTxnsCount field if non-nil, zero value otherwise.

### GetDigitalTxnsCountOk

`func (o *ResponseAccountingDetails) GetDigitalTxnsCountOk() (*int32, bool)`

GetDigitalTxnsCountOk returns a tuple with the DigitalTxnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalTxnsCount

`func (o *ResponseAccountingDetails) SetDigitalTxnsCount(v int32)`

SetDigitalTxnsCount sets DigitalTxnsCount field to given value.

### HasDigitalTxnsCount

`func (o *ResponseAccountingDetails) HasDigitalTxnsCount() bool`

HasDigitalTxnsCount returns a boolean if a field has been set.

### GetGlCode

`func (o *ResponseAccountingDetails) GetGlCode() string`

GetGlCode returns the GlCode field if non-nil, zero value otherwise.

### GetGlCodeOk

`func (o *ResponseAccountingDetails) GetGlCodeOk() (*string, bool)`

GetGlCodeOk returns a tuple with the GlCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlCode

`func (o *ResponseAccountingDetails) SetGlCode(v string)`

SetGlCode sets GlCode field to given value.

### HasGlCode

`func (o *ResponseAccountingDetails) HasGlCode() bool`

HasGlCode returns a boolean if a field has been set.

### GetHoa

`func (o *ResponseAccountingDetails) GetHoa() string`

GetHoa returns the Hoa field if non-nil, zero value otherwise.

### GetHoaOk

`func (o *ResponseAccountingDetails) GetHoaOk() (*string, bool)`

GetHoaOk returns a tuple with the Hoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoa

`func (o *ResponseAccountingDetails) SetHoa(v string)`

SetHoa sets Hoa field to given value.

### HasHoa

`func (o *ResponseAccountingDetails) HasHoa() bool`

HasHoa returns a boolean if a field has been set.

### GetNonDigitalTxnsAmount

`func (o *ResponseAccountingDetails) GetNonDigitalTxnsAmount() float32`

GetNonDigitalTxnsAmount returns the NonDigitalTxnsAmount field if non-nil, zero value otherwise.

### GetNonDigitalTxnsAmountOk

`func (o *ResponseAccountingDetails) GetNonDigitalTxnsAmountOk() (*float32, bool)`

GetNonDigitalTxnsAmountOk returns a tuple with the NonDigitalTxnsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDigitalTxnsAmount

`func (o *ResponseAccountingDetails) SetNonDigitalTxnsAmount(v float32)`

SetNonDigitalTxnsAmount sets NonDigitalTxnsAmount field to given value.

### HasNonDigitalTxnsAmount

`func (o *ResponseAccountingDetails) HasNonDigitalTxnsAmount() bool`

HasNonDigitalTxnsAmount returns a boolean if a field has been set.

### GetNonDigitalTxnsCount

`func (o *ResponseAccountingDetails) GetNonDigitalTxnsCount() int32`

GetNonDigitalTxnsCount returns the NonDigitalTxnsCount field if non-nil, zero value otherwise.

### GetNonDigitalTxnsCountOk

`func (o *ResponseAccountingDetails) GetNonDigitalTxnsCountOk() (*int32, bool)`

GetNonDigitalTxnsCountOk returns a tuple with the NonDigitalTxnsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDigitalTxnsCount

`func (o *ResponseAccountingDetails) SetNonDigitalTxnsCount(v int32)`

SetNonDigitalTxnsCount sets NonDigitalTxnsCount field to given value.

### HasNonDigitalTxnsCount

`func (o *ResponseAccountingDetails) HasNonDigitalTxnsCount() bool`

HasNonDigitalTxnsCount returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseAccountingDetails) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseAccountingDetails) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseAccountingDetails) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseAccountingDetails) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetPart

`func (o *ResponseAccountingDetails) GetPart() string`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ResponseAccountingDetails) GetPartOk() (*string, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ResponseAccountingDetails) SetPart(v string)`

SetPart sets Part field to given value.

### HasPart

`func (o *ResponseAccountingDetails) HasPart() bool`

HasPart returns a boolean if a field has been set.

### GetPositiveOrNegative

`func (o *ResponseAccountingDetails) GetPositiveOrNegative() string`

GetPositiveOrNegative returns the PositiveOrNegative field if non-nil, zero value otherwise.

### GetPositiveOrNegativeOk

`func (o *ResponseAccountingDetails) GetPositiveOrNegativeOk() (*string, bool)`

GetPositiveOrNegativeOk returns a tuple with the PositiveOrNegative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveOrNegative

`func (o *ResponseAccountingDetails) SetPositiveOrNegative(v string)`

SetPositiveOrNegative sets PositiveOrNegative field to given value.

### HasPositiveOrNegative

`func (o *ResponseAccountingDetails) HasPositiveOrNegative() bool`

HasPositiveOrNegative returns a boolean if a field has been set.

### GetReceiptOrPmt

`func (o *ResponseAccountingDetails) GetReceiptOrPmt() string`

GetReceiptOrPmt returns the ReceiptOrPmt field if non-nil, zero value otherwise.

### GetReceiptOrPmtOk

`func (o *ResponseAccountingDetails) GetReceiptOrPmtOk() (*string, bool)`

GetReceiptOrPmtOk returns a tuple with the ReceiptOrPmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptOrPmt

`func (o *ResponseAccountingDetails) SetReceiptOrPmt(v string)`

SetReceiptOrPmt sets ReceiptOrPmt field to given value.

### HasReceiptOrPmt

`func (o *ResponseAccountingDetails) HasReceiptOrPmt() bool`

HasReceiptOrPmt returns a boolean if a field has been set.

### GetReceiptSource

`func (o *ResponseAccountingDetails) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *ResponseAccountingDetails) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *ResponseAccountingDetails) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.

### HasReceiptSource

`func (o *ResponseAccountingDetails) HasReceiptSource() bool`

HasReceiptSource returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseAccountingDetails) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseAccountingDetails) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseAccountingDetails) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseAccountingDetails) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetSrcTranId

`func (o *ResponseAccountingDetails) GetSrcTranId() string`

GetSrcTranId returns the SrcTranId field if non-nil, zero value otherwise.

### GetSrcTranIdOk

`func (o *ResponseAccountingDetails) GetSrcTranIdOk() (*string, bool)`

GetSrcTranIdOk returns a tuple with the SrcTranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTranId

`func (o *ResponseAccountingDetails) SetSrcTranId(v string)`

SetSrcTranId sets SrcTranId field to given value.

### HasSrcTranId

`func (o *ResponseAccountingDetails) HasSrcTranId() bool`

HasSrcTranId returns a boolean if a field has been set.

### GetSrcTransDate

`func (o *ResponseAccountingDetails) GetSrcTransDate() string`

GetSrcTransDate returns the SrcTransDate field if non-nil, zero value otherwise.

### GetSrcTransDateOk

`func (o *ResponseAccountingDetails) GetSrcTransDateOk() (*string, bool)`

GetSrcTransDateOk returns a tuple with the SrcTransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTransDate

`func (o *ResponseAccountingDetails) SetSrcTransDate(v string)`

SetSrcTransDate sets SrcTransDate field to given value.

### HasSrcTransDate

`func (o *ResponseAccountingDetails) HasSrcTransDate() bool`

HasSrcTransDate returns a boolean if a field has been set.

### GetTotAmount

`func (o *ResponseAccountingDetails) GetTotAmount() float32`

GetTotAmount returns the TotAmount field if non-nil, zero value otherwise.

### GetTotAmountOk

`func (o *ResponseAccountingDetails) GetTotAmountOk() (*float32, bool)`

GetTotAmountOk returns a tuple with the TotAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotAmount

`func (o *ResponseAccountingDetails) SetTotAmount(v float32)`

SetTotAmount sets TotAmount field to given value.

### HasTotAmount

`func (o *ResponseAccountingDetails) HasTotAmount() bool`

HasTotAmount returns a boolean if a field has been set.

### GetTotTxns

`func (o *ResponseAccountingDetails) GetTotTxns() int32`

GetTotTxns returns the TotTxns field if non-nil, zero value otherwise.

### GetTotTxnsOk

`func (o *ResponseAccountingDetails) GetTotTxnsOk() (*int32, bool)`

GetTotTxnsOk returns a tuple with the TotTxns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotTxns

`func (o *ResponseAccountingDetails) SetTotTxns(v int32)`

SetTotTxns sets TotTxns field to given value.

### HasTotTxns

`func (o *ResponseAccountingDetails) HasTotTxns() bool`

HasTotTxns returns a boolean if a field has been set.

### GetTreasuryTranId

`func (o *ResponseAccountingDetails) GetTreasuryTranId() string`

GetTreasuryTranId returns the TreasuryTranId field if non-nil, zero value otherwise.

### GetTreasuryTranIdOk

`func (o *ResponseAccountingDetails) GetTreasuryTranIdOk() (*string, bool)`

GetTreasuryTranIdOk returns a tuple with the TreasuryTranId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryTranId

`func (o *ResponseAccountingDetails) SetTreasuryTranId(v string)`

SetTreasuryTranId sets TreasuryTranId field to given value.

### HasTreasuryTranId

`func (o *ResponseAccountingDetails) HasTreasuryTranId() bool`

HasTreasuryTranId returns a boolean if a field has been set.

### GetTryPostingDate

`func (o *ResponseAccountingDetails) GetTryPostingDate() string`

GetTryPostingDate returns the TryPostingDate field if non-nil, zero value otherwise.

### GetTryPostingDateOk

`func (o *ResponseAccountingDetails) GetTryPostingDateOk() (*string, bool)`

GetTryPostingDateOk returns a tuple with the TryPostingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryPostingDate

`func (o *ResponseAccountingDetails) SetTryPostingDate(v string)`

SetTryPostingDate sets TryPostingDate field to given value.

### HasTryPostingDate

`func (o *ResponseAccountingDetails) HasTryPostingDate() bool`

HasTryPostingDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResponseStampsObReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprUserId** | Pointer to **int32** |  | [optional] 
**ApprovedDate** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**ObTransactionId** | Pointer to **string** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**RecdUserId** | Pointer to **int32** |  | [optional] 
**ReceiptAmt** | Pointer to **float32** |  | [optional] 
**ReceiptDetails** | Pointer to **map[string]int32** |  | [optional] 
**ReceiptDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**ReceiptInvoice** | Pointer to **string** |  | [optional] 
**ReceiptSource** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseStampsObReceipts

`func NewResponseStampsObReceipts() *ResponseStampsObReceipts`

NewResponseStampsObReceipts instantiates a new ResponseStampsObReceipts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsObReceiptsWithDefaults

`func NewResponseStampsObReceiptsWithDefaults() *ResponseStampsObReceipts`

NewResponseStampsObReceiptsWithDefaults instantiates a new ResponseStampsObReceipts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprUserId

`func (o *ResponseStampsObReceipts) GetApprUserId() int32`

GetApprUserId returns the ApprUserId field if non-nil, zero value otherwise.

### GetApprUserIdOk

`func (o *ResponseStampsObReceipts) GetApprUserIdOk() (*int32, bool)`

GetApprUserIdOk returns a tuple with the ApprUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprUserId

`func (o *ResponseStampsObReceipts) SetApprUserId(v int32)`

SetApprUserId sets ApprUserId field to given value.

### HasApprUserId

`func (o *ResponseStampsObReceipts) HasApprUserId() bool`

HasApprUserId returns a boolean if a field has been set.

### GetApprovedDate

`func (o *ResponseStampsObReceipts) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ResponseStampsObReceipts) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ResponseStampsObReceipts) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ResponseStampsObReceipts) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *ResponseStampsObReceipts) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ResponseStampsObReceipts) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ResponseStampsObReceipts) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *ResponseStampsObReceipts) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseStampsObReceipts) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseStampsObReceipts) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseStampsObReceipts) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseStampsObReceipts) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetObTransactionId

`func (o *ResponseStampsObReceipts) GetObTransactionId() string`

GetObTransactionId returns the ObTransactionId field if non-nil, zero value otherwise.

### GetObTransactionIdOk

`func (o *ResponseStampsObReceipts) GetObTransactionIdOk() (*string, bool)`

GetObTransactionIdOk returns a tuple with the ObTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObTransactionId

`func (o *ResponseStampsObReceipts) SetObTransactionId(v string)`

SetObTransactionId sets ObTransactionId field to given value.

### HasObTransactionId

`func (o *ResponseStampsObReceipts) HasObTransactionId() bool`

HasObTransactionId returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseStampsObReceipts) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseStampsObReceipts) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseStampsObReceipts) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseStampsObReceipts) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRecdUserId

`func (o *ResponseStampsObReceipts) GetRecdUserId() int32`

GetRecdUserId returns the RecdUserId field if non-nil, zero value otherwise.

### GetRecdUserIdOk

`func (o *ResponseStampsObReceipts) GetRecdUserIdOk() (*int32, bool)`

GetRecdUserIdOk returns a tuple with the RecdUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecdUserId

`func (o *ResponseStampsObReceipts) SetRecdUserId(v int32)`

SetRecdUserId sets RecdUserId field to given value.

### HasRecdUserId

`func (o *ResponseStampsObReceipts) HasRecdUserId() bool`

HasRecdUserId returns a boolean if a field has been set.

### GetReceiptAmt

`func (o *ResponseStampsObReceipts) GetReceiptAmt() float32`

GetReceiptAmt returns the ReceiptAmt field if non-nil, zero value otherwise.

### GetReceiptAmtOk

`func (o *ResponseStampsObReceipts) GetReceiptAmtOk() (*float32, bool)`

GetReceiptAmtOk returns a tuple with the ReceiptAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptAmt

`func (o *ResponseStampsObReceipts) SetReceiptAmt(v float32)`

SetReceiptAmt sets ReceiptAmt field to given value.

### HasReceiptAmt

`func (o *ResponseStampsObReceipts) HasReceiptAmt() bool`

HasReceiptAmt returns a boolean if a field has been set.

### GetReceiptDetails

`func (o *ResponseStampsObReceipts) GetReceiptDetails() map[string]int32`

GetReceiptDetails returns the ReceiptDetails field if non-nil, zero value otherwise.

### GetReceiptDetailsOk

`func (o *ResponseStampsObReceipts) GetReceiptDetailsOk() (*map[string]int32, bool)`

GetReceiptDetailsOk returns a tuple with the ReceiptDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDetails

`func (o *ResponseStampsObReceipts) SetReceiptDetails(v map[string]int32)`

SetReceiptDetails sets ReceiptDetails field to given value.

### HasReceiptDetails

`func (o *ResponseStampsObReceipts) HasReceiptDetails() bool`

HasReceiptDetails returns a boolean if a field has been set.

### GetReceiptDetailsDesc

`func (o *ResponseStampsObReceipts) GetReceiptDetailsDesc() []ResponseStampdetails`

GetReceiptDetailsDesc returns the ReceiptDetailsDesc field if non-nil, zero value otherwise.

### GetReceiptDetailsDescOk

`func (o *ResponseStampsObReceipts) GetReceiptDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetReceiptDetailsDescOk returns a tuple with the ReceiptDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDetailsDesc

`func (o *ResponseStampsObReceipts) SetReceiptDetailsDesc(v []ResponseStampdetails)`

SetReceiptDetailsDesc sets ReceiptDetailsDesc field to given value.

### HasReceiptDetailsDesc

`func (o *ResponseStampsObReceipts) HasReceiptDetailsDesc() bool`

HasReceiptDetailsDesc returns a boolean if a field has been set.

### GetReceiptInvoice

`func (o *ResponseStampsObReceipts) GetReceiptInvoice() string`

GetReceiptInvoice returns the ReceiptInvoice field if non-nil, zero value otherwise.

### GetReceiptInvoiceOk

`func (o *ResponseStampsObReceipts) GetReceiptInvoiceOk() (*string, bool)`

GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptInvoice

`func (o *ResponseStampsObReceipts) SetReceiptInvoice(v string)`

SetReceiptInvoice sets ReceiptInvoice field to given value.

### HasReceiptInvoice

`func (o *ResponseStampsObReceipts) HasReceiptInvoice() bool`

HasReceiptInvoice returns a boolean if a field has been set.

### GetReceiptSource

`func (o *ResponseStampsObReceipts) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *ResponseStampsObReceipts) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *ResponseStampsObReceipts) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.

### HasReceiptSource

`func (o *ResponseStampsObReceipts) HasReceiptSource() bool`

HasReceiptSource returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsObReceipts) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsObReceipts) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsObReceipts) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsObReceipts) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseStampsObReceipts) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseStampsObReceipts) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseStampsObReceipts) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseStampsObReceipts) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



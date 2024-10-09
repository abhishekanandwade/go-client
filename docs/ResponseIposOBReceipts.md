# ResponseIposOBReceipts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprUserId** | Pointer to **int32** |  | [optional] 
**ApprovedDate** | Pointer to **string** |  | [optional] 
**ApproverRemarks** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**IpoDetails** | Pointer to [**[]ResponseIPOdetails**](ResponseIPOdetails.md) |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**ObTransactionId** | Pointer to **string** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**RecdUserId** | Pointer to **int32** |  | [optional] 
**ReceiptInvoice** | Pointer to **string** |  | [optional] 
**ReceiptSource** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseIposOBReceipts

`func NewResponseIposOBReceipts() *ResponseIposOBReceipts`

NewResponseIposOBReceipts instantiates a new ResponseIposOBReceipts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseIposOBReceiptsWithDefaults

`func NewResponseIposOBReceiptsWithDefaults() *ResponseIposOBReceipts`

NewResponseIposOBReceiptsWithDefaults instantiates a new ResponseIposOBReceipts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprUserId

`func (o *ResponseIposOBReceipts) GetApprUserId() int32`

GetApprUserId returns the ApprUserId field if non-nil, zero value otherwise.

### GetApprUserIdOk

`func (o *ResponseIposOBReceipts) GetApprUserIdOk() (*int32, bool)`

GetApprUserIdOk returns a tuple with the ApprUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprUserId

`func (o *ResponseIposOBReceipts) SetApprUserId(v int32)`

SetApprUserId sets ApprUserId field to given value.

### HasApprUserId

`func (o *ResponseIposOBReceipts) HasApprUserId() bool`

HasApprUserId returns a boolean if a field has been set.

### GetApprovedDate

`func (o *ResponseIposOBReceipts) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ResponseIposOBReceipts) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ResponseIposOBReceipts) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ResponseIposOBReceipts) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetApproverRemarks

`func (o *ResponseIposOBReceipts) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *ResponseIposOBReceipts) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *ResponseIposOBReceipts) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.

### HasApproverRemarks

`func (o *ResponseIposOBReceipts) HasApproverRemarks() bool`

HasApproverRemarks returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *ResponseIposOBReceipts) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ResponseIposOBReceipts) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ResponseIposOBReceipts) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *ResponseIposOBReceipts) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetIpoDetails

`func (o *ResponseIposOBReceipts) GetIpoDetails() []ResponseIPOdetails`

GetIpoDetails returns the IpoDetails field if non-nil, zero value otherwise.

### GetIpoDetailsOk

`func (o *ResponseIposOBReceipts) GetIpoDetailsOk() (*[]ResponseIPOdetails, bool)`

GetIpoDetailsOk returns a tuple with the IpoDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpoDetails

`func (o *ResponseIposOBReceipts) SetIpoDetails(v []ResponseIPOdetails)`

SetIpoDetails sets IpoDetails field to given value.

### HasIpoDetails

`func (o *ResponseIposOBReceipts) HasIpoDetails() bool`

HasIpoDetails returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseIposOBReceipts) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseIposOBReceipts) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseIposOBReceipts) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseIposOBReceipts) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetObTransactionId

`func (o *ResponseIposOBReceipts) GetObTransactionId() string`

GetObTransactionId returns the ObTransactionId field if non-nil, zero value otherwise.

### GetObTransactionIdOk

`func (o *ResponseIposOBReceipts) GetObTransactionIdOk() (*string, bool)`

GetObTransactionIdOk returns a tuple with the ObTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObTransactionId

`func (o *ResponseIposOBReceipts) SetObTransactionId(v string)`

SetObTransactionId sets ObTransactionId field to given value.

### HasObTransactionId

`func (o *ResponseIposOBReceipts) HasObTransactionId() bool`

HasObTransactionId returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseIposOBReceipts) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseIposOBReceipts) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseIposOBReceipts) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseIposOBReceipts) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRecdUserId

`func (o *ResponseIposOBReceipts) GetRecdUserId() int32`

GetRecdUserId returns the RecdUserId field if non-nil, zero value otherwise.

### GetRecdUserIdOk

`func (o *ResponseIposOBReceipts) GetRecdUserIdOk() (*int32, bool)`

GetRecdUserIdOk returns a tuple with the RecdUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecdUserId

`func (o *ResponseIposOBReceipts) SetRecdUserId(v int32)`

SetRecdUserId sets RecdUserId field to given value.

### HasRecdUserId

`func (o *ResponseIposOBReceipts) HasRecdUserId() bool`

HasRecdUserId returns a boolean if a field has been set.

### GetReceiptInvoice

`func (o *ResponseIposOBReceipts) GetReceiptInvoice() string`

GetReceiptInvoice returns the ReceiptInvoice field if non-nil, zero value otherwise.

### GetReceiptInvoiceOk

`func (o *ResponseIposOBReceipts) GetReceiptInvoiceOk() (*string, bool)`

GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptInvoice

`func (o *ResponseIposOBReceipts) SetReceiptInvoice(v string)`

SetReceiptInvoice sets ReceiptInvoice field to given value.

### HasReceiptInvoice

`func (o *ResponseIposOBReceipts) HasReceiptInvoice() bool`

HasReceiptInvoice returns a boolean if a field has been set.

### GetReceiptSource

`func (o *ResponseIposOBReceipts) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *ResponseIposOBReceipts) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *ResponseIposOBReceipts) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.

### HasReceiptSource

`func (o *ResponseIposOBReceipts) HasReceiptSource() bool`

HasReceiptSource returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseIposOBReceipts) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseIposOBReceipts) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseIposOBReceipts) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseIposOBReceipts) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseIposOBReceipts) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseIposOBReceipts) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseIposOBReceipts) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseIposOBReceipts) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



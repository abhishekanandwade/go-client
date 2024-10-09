# ResponseStampsObReceiptsInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceDate** | Pointer to **string** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**RecdUserId** | Pointer to **int32** |  | [optional] 
**ReceiptAmt** | Pointer to **float32** |  | [optional] 
**ReceiptDetails** | Pointer to **map[string]int32** |  | [optional] 
**ReceiptInvoice** | Pointer to **string** |  | [optional] 
**ReceiptSource** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 
**SubRequestId** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseStampsObReceiptsInventory

`func NewResponseStampsObReceiptsInventory() *ResponseStampsObReceiptsInventory`

NewResponseStampsObReceiptsInventory instantiates a new ResponseStampsObReceiptsInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampsObReceiptsInventoryWithDefaults

`func NewResponseStampsObReceiptsInventoryWithDefaults() *ResponseStampsObReceiptsInventory`

NewResponseStampsObReceiptsInventoryWithDefaults instantiates a new ResponseStampsObReceiptsInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceDate

`func (o *ResponseStampsObReceiptsInventory) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *ResponseStampsObReceiptsInventory) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *ResponseStampsObReceiptsInventory) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *ResponseStampsObReceiptsInventory) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseStampsObReceiptsInventory) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseStampsObReceiptsInventory) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseStampsObReceiptsInventory) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseStampsObReceiptsInventory) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRecdUserId

`func (o *ResponseStampsObReceiptsInventory) GetRecdUserId() int32`

GetRecdUserId returns the RecdUserId field if non-nil, zero value otherwise.

### GetRecdUserIdOk

`func (o *ResponseStampsObReceiptsInventory) GetRecdUserIdOk() (*int32, bool)`

GetRecdUserIdOk returns a tuple with the RecdUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecdUserId

`func (o *ResponseStampsObReceiptsInventory) SetRecdUserId(v int32)`

SetRecdUserId sets RecdUserId field to given value.

### HasRecdUserId

`func (o *ResponseStampsObReceiptsInventory) HasRecdUserId() bool`

HasRecdUserId returns a boolean if a field has been set.

### GetReceiptAmt

`func (o *ResponseStampsObReceiptsInventory) GetReceiptAmt() float32`

GetReceiptAmt returns the ReceiptAmt field if non-nil, zero value otherwise.

### GetReceiptAmtOk

`func (o *ResponseStampsObReceiptsInventory) GetReceiptAmtOk() (*float32, bool)`

GetReceiptAmtOk returns a tuple with the ReceiptAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptAmt

`func (o *ResponseStampsObReceiptsInventory) SetReceiptAmt(v float32)`

SetReceiptAmt sets ReceiptAmt field to given value.

### HasReceiptAmt

`func (o *ResponseStampsObReceiptsInventory) HasReceiptAmt() bool`

HasReceiptAmt returns a boolean if a field has been set.

### GetReceiptDetails

`func (o *ResponseStampsObReceiptsInventory) GetReceiptDetails() map[string]int32`

GetReceiptDetails returns the ReceiptDetails field if non-nil, zero value otherwise.

### GetReceiptDetailsOk

`func (o *ResponseStampsObReceiptsInventory) GetReceiptDetailsOk() (*map[string]int32, bool)`

GetReceiptDetailsOk returns a tuple with the ReceiptDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptDetails

`func (o *ResponseStampsObReceiptsInventory) SetReceiptDetails(v map[string]int32)`

SetReceiptDetails sets ReceiptDetails field to given value.

### HasReceiptDetails

`func (o *ResponseStampsObReceiptsInventory) HasReceiptDetails() bool`

HasReceiptDetails returns a boolean if a field has been set.

### GetReceiptInvoice

`func (o *ResponseStampsObReceiptsInventory) GetReceiptInvoice() string`

GetReceiptInvoice returns the ReceiptInvoice field if non-nil, zero value otherwise.

### GetReceiptInvoiceOk

`func (o *ResponseStampsObReceiptsInventory) GetReceiptInvoiceOk() (*string, bool)`

GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptInvoice

`func (o *ResponseStampsObReceiptsInventory) SetReceiptInvoice(v string)`

SetReceiptInvoice sets ReceiptInvoice field to given value.

### HasReceiptInvoice

`func (o *ResponseStampsObReceiptsInventory) HasReceiptInvoice() bool`

HasReceiptInvoice returns a boolean if a field has been set.

### GetReceiptSource

`func (o *ResponseStampsObReceiptsInventory) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *ResponseStampsObReceiptsInventory) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *ResponseStampsObReceiptsInventory) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.

### HasReceiptSource

`func (o *ResponseStampsObReceiptsInventory) HasReceiptSource() bool`

HasReceiptSource returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseStampsObReceiptsInventory) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseStampsObReceiptsInventory) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseStampsObReceiptsInventory) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseStampsObReceiptsInventory) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRequestId

`func (o *ResponseStampsObReceiptsInventory) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResponseStampsObReceiptsInventory) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResponseStampsObReceiptsInventory) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResponseStampsObReceiptsInventory) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetSubRequestId

`func (o *ResponseStampsObReceiptsInventory) GetSubRequestId() int32`

GetSubRequestId returns the SubRequestId field if non-nil, zero value otherwise.

### GetSubRequestIdOk

`func (o *ResponseStampsObReceiptsInventory) GetSubRequestIdOk() (*int32, bool)`

GetSubRequestIdOk returns a tuple with the SubRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRequestId

`func (o *ResponseStampsObReceiptsInventory) SetSubRequestId(v int32)`

SetSubRequestId sets SubRequestId field to given value.

### HasSubRequestId

`func (o *ResponseStampsObReceiptsInventory) HasSubRequestId() bool`

HasSubRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



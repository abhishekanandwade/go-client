# HandlerCreateOfficesIPOsOBReceiptsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceDate** | **string** |  | 
**IpoDetails** | Pointer to [**[]HandlerIPODetail**](HandlerIPODetail.md) |  | [optional] 
**RecdUserId** | **int32** |  | 
**ReceiptInvoice** | **string** |  | 
**ReceiptSource** | **string** |  | 
**Remarks** | Pointer to **string** |  | [optional] 

## Methods

### NewHandlerCreateOfficesIPOsOBReceiptsRequest

`func NewHandlerCreateOfficesIPOsOBReceiptsRequest(invoiceDate string, recdUserId int32, receiptInvoice string, receiptSource string, ) *HandlerCreateOfficesIPOsOBReceiptsRequest`

NewHandlerCreateOfficesIPOsOBReceiptsRequest instantiates a new HandlerCreateOfficesIPOsOBReceiptsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateOfficesIPOsOBReceiptsRequestWithDefaults

`func NewHandlerCreateOfficesIPOsOBReceiptsRequestWithDefaults() *HandlerCreateOfficesIPOsOBReceiptsRequest`

NewHandlerCreateOfficesIPOsOBReceiptsRequestWithDefaults instantiates a new HandlerCreateOfficesIPOsOBReceiptsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceDate

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.


### GetIpoDetails

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetIpoDetails() []HandlerIPODetail`

GetIpoDetails returns the IpoDetails field if non-nil, zero value otherwise.

### GetIpoDetailsOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetIpoDetailsOk() (*[]HandlerIPODetail, bool)`

GetIpoDetailsOk returns a tuple with the IpoDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpoDetails

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetIpoDetails(v []HandlerIPODetail)`

SetIpoDetails sets IpoDetails field to given value.

### HasIpoDetails

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) HasIpoDetails() bool`

HasIpoDetails returns a boolean if a field has been set.

### GetRecdUserId

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRecdUserId() int32`

GetRecdUserId returns the RecdUserId field if non-nil, zero value otherwise.

### GetRecdUserIdOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRecdUserIdOk() (*int32, bool)`

GetRecdUserIdOk returns a tuple with the RecdUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecdUserId

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetRecdUserId(v int32)`

SetRecdUserId sets RecdUserId field to given value.


### GetReceiptInvoice

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptInvoice() string`

GetReceiptInvoice returns the ReceiptInvoice field if non-nil, zero value otherwise.

### GetReceiptInvoiceOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptInvoiceOk() (*string, bool)`

GetReceiptInvoiceOk returns a tuple with the ReceiptInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptInvoice

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetReceiptInvoice(v string)`

SetReceiptInvoice sets ReceiptInvoice field to given value.


### GetReceiptSource

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptSource() string`

GetReceiptSource returns the ReceiptSource field if non-nil, zero value otherwise.

### GetReceiptSourceOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetReceiptSourceOk() (*string, bool)`

GetReceiptSourceOk returns a tuple with the ReceiptSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptSource

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetReceiptSource(v string)`

SetReceiptSource sets ReceiptSource field to given value.


### GetRemarks

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerCreateOfficesIPOsOBReceiptsRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



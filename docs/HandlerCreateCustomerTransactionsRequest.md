# HandlerCreateCustomerTransactionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillAmount** | **float32** |  | 
**BillDate** | **string** |  | 
**BillId** | **string** |  | 
**BillRemarks** | **string** |  | 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**CustomerId** | **string** |  | 
**CustomerName** | **string** |  | 
**CustomerTypeId** | Pointer to **string** |  | [optional] 
**CustomerTypeName** | **string** |  | 
**IsAdvanceCustomer** | Pointer to **bool** |  | [optional] 
**IsReceiptPayment** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**OfficeId** | **int32** |  | 
**Remarks** | **string** |  | 
**TxAmount** | **float32** |  | 
**TxMode** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewHandlerCreateCustomerTransactionsRequest

`func NewHandlerCreateCustomerTransactionsRequest(billAmount float32, billDate string, billId string, billRemarks string, customerId string, customerName string, customerTypeName string, officeId int32, remarks string, txAmount float32, txMode string, userId string, ) *HandlerCreateCustomerTransactionsRequest`

NewHandlerCreateCustomerTransactionsRequest instantiates a new HandlerCreateCustomerTransactionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateCustomerTransactionsRequestWithDefaults

`func NewHandlerCreateCustomerTransactionsRequestWithDefaults() *HandlerCreateCustomerTransactionsRequest`

NewHandlerCreateCustomerTransactionsRequestWithDefaults instantiates a new HandlerCreateCustomerTransactionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillAmount

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillAmount() float32`

GetBillAmount returns the BillAmount field if non-nil, zero value otherwise.

### GetBillAmountOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillAmountOk() (*float32, bool)`

GetBillAmountOk returns a tuple with the BillAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAmount

`func (o *HandlerCreateCustomerTransactionsRequest) SetBillAmount(v float32)`

SetBillAmount sets BillAmount field to given value.


### GetBillDate

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillDate() string`

GetBillDate returns the BillDate field if non-nil, zero value otherwise.

### GetBillDateOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillDateOk() (*string, bool)`

GetBillDateOk returns a tuple with the BillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillDate

`func (o *HandlerCreateCustomerTransactionsRequest) SetBillDate(v string)`

SetBillDate sets BillDate field to given value.


### GetBillId

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *HandlerCreateCustomerTransactionsRequest) SetBillId(v string)`

SetBillId sets BillId field to given value.


### GetBillRemarks

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillRemarks() string`

GetBillRemarks returns the BillRemarks field if non-nil, zero value otherwise.

### GetBillRemarksOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetBillRemarksOk() (*string, bool)`

GetBillRemarksOk returns a tuple with the BillRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillRemarks

`func (o *HandlerCreateCustomerTransactionsRequest) SetBillRemarks(v string)`

SetBillRemarks sets BillRemarks field to given value.


### GetChequeDate

`func (o *HandlerCreateCustomerTransactionsRequest) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *HandlerCreateCustomerTransactionsRequest) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *HandlerCreateCustomerTransactionsRequest) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *HandlerCreateCustomerTransactionsRequest) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *HandlerCreateCustomerTransactionsRequest) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *HandlerCreateCustomerTransactionsRequest) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetCustomerId

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *HandlerCreateCustomerTransactionsRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetCustomerName

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *HandlerCreateCustomerTransactionsRequest) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetCustomerTypeId

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerTypeId() string`

GetCustomerTypeId returns the CustomerTypeId field if non-nil, zero value otherwise.

### GetCustomerTypeIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerTypeIdOk() (*string, bool)`

GetCustomerTypeIdOk returns a tuple with the CustomerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTypeId

`func (o *HandlerCreateCustomerTransactionsRequest) SetCustomerTypeId(v string)`

SetCustomerTypeId sets CustomerTypeId field to given value.

### HasCustomerTypeId

`func (o *HandlerCreateCustomerTransactionsRequest) HasCustomerTypeId() bool`

HasCustomerTypeId returns a boolean if a field has been set.

### GetCustomerTypeName

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerTypeName() string`

GetCustomerTypeName returns the CustomerTypeName field if non-nil, zero value otherwise.

### GetCustomerTypeNameOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetCustomerTypeNameOk() (*string, bool)`

GetCustomerTypeNameOk returns a tuple with the CustomerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTypeName

`func (o *HandlerCreateCustomerTransactionsRequest) SetCustomerTypeName(v string)`

SetCustomerTypeName sets CustomerTypeName field to given value.


### GetIsAdvanceCustomer

`func (o *HandlerCreateCustomerTransactionsRequest) GetIsAdvanceCustomer() bool`

GetIsAdvanceCustomer returns the IsAdvanceCustomer field if non-nil, zero value otherwise.

### GetIsAdvanceCustomerOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetIsAdvanceCustomerOk() (*bool, bool)`

GetIsAdvanceCustomerOk returns a tuple with the IsAdvanceCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanceCustomer

`func (o *HandlerCreateCustomerTransactionsRequest) SetIsAdvanceCustomer(v bool)`

SetIsAdvanceCustomer sets IsAdvanceCustomer field to given value.

### HasIsAdvanceCustomer

`func (o *HandlerCreateCustomerTransactionsRequest) HasIsAdvanceCustomer() bool`

HasIsAdvanceCustomer returns a boolean if a field has been set.

### GetIsReceiptPayment

`func (o *HandlerCreateCustomerTransactionsRequest) GetIsReceiptPayment() string`

GetIsReceiptPayment returns the IsReceiptPayment field if non-nil, zero value otherwise.

### GetIsReceiptPaymentOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetIsReceiptPaymentOk() (*string, bool)`

GetIsReceiptPaymentOk returns a tuple with the IsReceiptPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceiptPayment

`func (o *HandlerCreateCustomerTransactionsRequest) SetIsReceiptPayment(v string)`

SetIsReceiptPayment sets IsReceiptPayment field to given value.

### HasIsReceiptPayment

`func (o *HandlerCreateCustomerTransactionsRequest) HasIsReceiptPayment() bool`

HasIsReceiptPayment returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *HandlerCreateCustomerTransactionsRequest) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *HandlerCreateCustomerTransactionsRequest) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *HandlerCreateCustomerTransactionsRequest) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerCreateCustomerTransactionsRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateCustomerTransactionsRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetRemarks

`func (o *HandlerCreateCustomerTransactionsRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerCreateCustomerTransactionsRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.


### GetTxAmount

`func (o *HandlerCreateCustomerTransactionsRequest) GetTxAmount() float32`

GetTxAmount returns the TxAmount field if non-nil, zero value otherwise.

### GetTxAmountOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetTxAmountOk() (*float32, bool)`

GetTxAmountOk returns a tuple with the TxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAmount

`func (o *HandlerCreateCustomerTransactionsRequest) SetTxAmount(v float32)`

SetTxAmount sets TxAmount field to given value.


### GetTxMode

`func (o *HandlerCreateCustomerTransactionsRequest) GetTxMode() string`

GetTxMode returns the TxMode field if non-nil, zero value otherwise.

### GetTxModeOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetTxModeOk() (*string, bool)`

GetTxModeOk returns a tuple with the TxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMode

`func (o *HandlerCreateCustomerTransactionsRequest) SetTxMode(v string)`

SetTxMode sets TxMode field to given value.


### GetUserId

`func (o *HandlerCreateCustomerTransactionsRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerCreateCustomerTransactionsRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerCreateCustomerTransactionsRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



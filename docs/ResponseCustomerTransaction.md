# ResponseCustomerTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverDate** | Pointer to **string** |  | [optional] 
**ApproverId** | Pointer to **string** |  | [optional] 
**BillAmount** | Pointer to **float32** |  | [optional] 
**BillDate** | Pointer to **string** |  | [optional] 
**BillId** | Pointer to **string** |  | [optional] 
**BillRemarks** | Pointer to **string** |  | [optional] 
**ChequeDate** | Pointer to **string** |  | [optional] 
**ChequeNo** | Pointer to **string** |  | [optional] 
**CustomerId** | Pointer to **string** |  | [optional] 
**CustomerName** | Pointer to **string** |  | [optional] 
**CustomerTypeId** | Pointer to **string** |  | [optional] 
**CustomerTypeName** | Pointer to **string** |  | [optional] 
**IsAdvanceCustomer** | Pointer to **bool** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsRealized** | Pointer to **bool** |  | [optional] 
**IsReceiptPayment** | Pointer to **string** |  | [optional] 
**IssOfficeId** | Pointer to **int32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**RealizationDate** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransId** | Pointer to **string** |  | [optional] 
**TxAmount** | Pointer to **float32** |  | [optional] 
**TxMode** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseCustomerTransaction

`func NewResponseCustomerTransaction() *ResponseCustomerTransaction`

NewResponseCustomerTransaction instantiates a new ResponseCustomerTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCustomerTransactionWithDefaults

`func NewResponseCustomerTransactionWithDefaults() *ResponseCustomerTransaction`

NewResponseCustomerTransactionWithDefaults instantiates a new ResponseCustomerTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverDate

`func (o *ResponseCustomerTransaction) GetApproverDate() string`

GetApproverDate returns the ApproverDate field if non-nil, zero value otherwise.

### GetApproverDateOk

`func (o *ResponseCustomerTransaction) GetApproverDateOk() (*string, bool)`

GetApproverDateOk returns a tuple with the ApproverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverDate

`func (o *ResponseCustomerTransaction) SetApproverDate(v string)`

SetApproverDate sets ApproverDate field to given value.

### HasApproverDate

`func (o *ResponseCustomerTransaction) HasApproverDate() bool`

HasApproverDate returns a boolean if a field has been set.

### GetApproverId

`func (o *ResponseCustomerTransaction) GetApproverId() string`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ResponseCustomerTransaction) GetApproverIdOk() (*string, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ResponseCustomerTransaction) SetApproverId(v string)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ResponseCustomerTransaction) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetBillAmount

`func (o *ResponseCustomerTransaction) GetBillAmount() float32`

GetBillAmount returns the BillAmount field if non-nil, zero value otherwise.

### GetBillAmountOk

`func (o *ResponseCustomerTransaction) GetBillAmountOk() (*float32, bool)`

GetBillAmountOk returns a tuple with the BillAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillAmount

`func (o *ResponseCustomerTransaction) SetBillAmount(v float32)`

SetBillAmount sets BillAmount field to given value.

### HasBillAmount

`func (o *ResponseCustomerTransaction) HasBillAmount() bool`

HasBillAmount returns a boolean if a field has been set.

### GetBillDate

`func (o *ResponseCustomerTransaction) GetBillDate() string`

GetBillDate returns the BillDate field if non-nil, zero value otherwise.

### GetBillDateOk

`func (o *ResponseCustomerTransaction) GetBillDateOk() (*string, bool)`

GetBillDateOk returns a tuple with the BillDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillDate

`func (o *ResponseCustomerTransaction) SetBillDate(v string)`

SetBillDate sets BillDate field to given value.

### HasBillDate

`func (o *ResponseCustomerTransaction) HasBillDate() bool`

HasBillDate returns a boolean if a field has been set.

### GetBillId

`func (o *ResponseCustomerTransaction) GetBillId() string`

GetBillId returns the BillId field if non-nil, zero value otherwise.

### GetBillIdOk

`func (o *ResponseCustomerTransaction) GetBillIdOk() (*string, bool)`

GetBillIdOk returns a tuple with the BillId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillId

`func (o *ResponseCustomerTransaction) SetBillId(v string)`

SetBillId sets BillId field to given value.

### HasBillId

`func (o *ResponseCustomerTransaction) HasBillId() bool`

HasBillId returns a boolean if a field has been set.

### GetBillRemarks

`func (o *ResponseCustomerTransaction) GetBillRemarks() string`

GetBillRemarks returns the BillRemarks field if non-nil, zero value otherwise.

### GetBillRemarksOk

`func (o *ResponseCustomerTransaction) GetBillRemarksOk() (*string, bool)`

GetBillRemarksOk returns a tuple with the BillRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillRemarks

`func (o *ResponseCustomerTransaction) SetBillRemarks(v string)`

SetBillRemarks sets BillRemarks field to given value.

### HasBillRemarks

`func (o *ResponseCustomerTransaction) HasBillRemarks() bool`

HasBillRemarks returns a boolean if a field has been set.

### GetChequeDate

`func (o *ResponseCustomerTransaction) GetChequeDate() string`

GetChequeDate returns the ChequeDate field if non-nil, zero value otherwise.

### GetChequeDateOk

`func (o *ResponseCustomerTransaction) GetChequeDateOk() (*string, bool)`

GetChequeDateOk returns a tuple with the ChequeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDate

`func (o *ResponseCustomerTransaction) SetChequeDate(v string)`

SetChequeDate sets ChequeDate field to given value.

### HasChequeDate

`func (o *ResponseCustomerTransaction) HasChequeDate() bool`

HasChequeDate returns a boolean if a field has been set.

### GetChequeNo

`func (o *ResponseCustomerTransaction) GetChequeNo() string`

GetChequeNo returns the ChequeNo field if non-nil, zero value otherwise.

### GetChequeNoOk

`func (o *ResponseCustomerTransaction) GetChequeNoOk() (*string, bool)`

GetChequeNoOk returns a tuple with the ChequeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeNo

`func (o *ResponseCustomerTransaction) SetChequeNo(v string)`

SetChequeNo sets ChequeNo field to given value.

### HasChequeNo

`func (o *ResponseCustomerTransaction) HasChequeNo() bool`

HasChequeNo returns a boolean if a field has been set.

### GetCustomerId

`func (o *ResponseCustomerTransaction) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *ResponseCustomerTransaction) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *ResponseCustomerTransaction) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *ResponseCustomerTransaction) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerName

`func (o *ResponseCustomerTransaction) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *ResponseCustomerTransaction) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *ResponseCustomerTransaction) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *ResponseCustomerTransaction) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerTypeId

`func (o *ResponseCustomerTransaction) GetCustomerTypeId() string`

GetCustomerTypeId returns the CustomerTypeId field if non-nil, zero value otherwise.

### GetCustomerTypeIdOk

`func (o *ResponseCustomerTransaction) GetCustomerTypeIdOk() (*string, bool)`

GetCustomerTypeIdOk returns a tuple with the CustomerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTypeId

`func (o *ResponseCustomerTransaction) SetCustomerTypeId(v string)`

SetCustomerTypeId sets CustomerTypeId field to given value.

### HasCustomerTypeId

`func (o *ResponseCustomerTransaction) HasCustomerTypeId() bool`

HasCustomerTypeId returns a boolean if a field has been set.

### GetCustomerTypeName

`func (o *ResponseCustomerTransaction) GetCustomerTypeName() string`

GetCustomerTypeName returns the CustomerTypeName field if non-nil, zero value otherwise.

### GetCustomerTypeNameOk

`func (o *ResponseCustomerTransaction) GetCustomerTypeNameOk() (*string, bool)`

GetCustomerTypeNameOk returns a tuple with the CustomerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerTypeName

`func (o *ResponseCustomerTransaction) SetCustomerTypeName(v string)`

SetCustomerTypeName sets CustomerTypeName field to given value.

### HasCustomerTypeName

`func (o *ResponseCustomerTransaction) HasCustomerTypeName() bool`

HasCustomerTypeName returns a boolean if a field has been set.

### GetIsAdvanceCustomer

`func (o *ResponseCustomerTransaction) GetIsAdvanceCustomer() bool`

GetIsAdvanceCustomer returns the IsAdvanceCustomer field if non-nil, zero value otherwise.

### GetIsAdvanceCustomerOk

`func (o *ResponseCustomerTransaction) GetIsAdvanceCustomerOk() (*bool, bool)`

GetIsAdvanceCustomerOk returns a tuple with the IsAdvanceCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanceCustomer

`func (o *ResponseCustomerTransaction) SetIsAdvanceCustomer(v bool)`

SetIsAdvanceCustomer sets IsAdvanceCustomer field to given value.

### HasIsAdvanceCustomer

`func (o *ResponseCustomerTransaction) HasIsAdvanceCustomer() bool`

HasIsAdvanceCustomer returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseCustomerTransaction) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseCustomerTransaction) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseCustomerTransaction) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseCustomerTransaction) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsRealized

`func (o *ResponseCustomerTransaction) GetIsRealized() bool`

GetIsRealized returns the IsRealized field if non-nil, zero value otherwise.

### GetIsRealizedOk

`func (o *ResponseCustomerTransaction) GetIsRealizedOk() (*bool, bool)`

GetIsRealizedOk returns a tuple with the IsRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRealized

`func (o *ResponseCustomerTransaction) SetIsRealized(v bool)`

SetIsRealized sets IsRealized field to given value.

### HasIsRealized

`func (o *ResponseCustomerTransaction) HasIsRealized() bool`

HasIsRealized returns a boolean if a field has been set.

### GetIsReceiptPayment

`func (o *ResponseCustomerTransaction) GetIsReceiptPayment() string`

GetIsReceiptPayment returns the IsReceiptPayment field if non-nil, zero value otherwise.

### GetIsReceiptPaymentOk

`func (o *ResponseCustomerTransaction) GetIsReceiptPaymentOk() (*string, bool)`

GetIsReceiptPaymentOk returns a tuple with the IsReceiptPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReceiptPayment

`func (o *ResponseCustomerTransaction) SetIsReceiptPayment(v string)`

SetIsReceiptPayment sets IsReceiptPayment field to given value.

### HasIsReceiptPayment

`func (o *ResponseCustomerTransaction) HasIsReceiptPayment() bool`

HasIsReceiptPayment returns a boolean if a field has been set.

### GetIssOfficeId

`func (o *ResponseCustomerTransaction) GetIssOfficeId() int32`

GetIssOfficeId returns the IssOfficeId field if non-nil, zero value otherwise.

### GetIssOfficeIdOk

`func (o *ResponseCustomerTransaction) GetIssOfficeIdOk() (*int32, bool)`

GetIssOfficeIdOk returns a tuple with the IssOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssOfficeId

`func (o *ResponseCustomerTransaction) SetIssOfficeId(v int32)`

SetIssOfficeId sets IssOfficeId field to given value.

### HasIssOfficeId

`func (o *ResponseCustomerTransaction) HasIssOfficeId() bool`

HasIssOfficeId returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseCustomerTransaction) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseCustomerTransaction) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseCustomerTransaction) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseCustomerTransaction) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetRealizationDate

`func (o *ResponseCustomerTransaction) GetRealizationDate() string`

GetRealizationDate returns the RealizationDate field if non-nil, zero value otherwise.

### GetRealizationDateOk

`func (o *ResponseCustomerTransaction) GetRealizationDateOk() (*string, bool)`

GetRealizationDateOk returns a tuple with the RealizationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizationDate

`func (o *ResponseCustomerTransaction) SetRealizationDate(v string)`

SetRealizationDate sets RealizationDate field to given value.

### HasRealizationDate

`func (o *ResponseCustomerTransaction) HasRealizationDate() bool`

HasRealizationDate returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseCustomerTransaction) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseCustomerTransaction) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseCustomerTransaction) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseCustomerTransaction) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseCustomerTransaction) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseCustomerTransaction) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseCustomerTransaction) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseCustomerTransaction) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransId

`func (o *ResponseCustomerTransaction) GetTransId() string`

GetTransId returns the TransId field if non-nil, zero value otherwise.

### GetTransIdOk

`func (o *ResponseCustomerTransaction) GetTransIdOk() (*string, bool)`

GetTransIdOk returns a tuple with the TransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransId

`func (o *ResponseCustomerTransaction) SetTransId(v string)`

SetTransId sets TransId field to given value.

### HasTransId

`func (o *ResponseCustomerTransaction) HasTransId() bool`

HasTransId returns a boolean if a field has been set.

### GetTxAmount

`func (o *ResponseCustomerTransaction) GetTxAmount() float32`

GetTxAmount returns the TxAmount field if non-nil, zero value otherwise.

### GetTxAmountOk

`func (o *ResponseCustomerTransaction) GetTxAmountOk() (*float32, bool)`

GetTxAmountOk returns a tuple with the TxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxAmount

`func (o *ResponseCustomerTransaction) SetTxAmount(v float32)`

SetTxAmount sets TxAmount field to given value.

### HasTxAmount

`func (o *ResponseCustomerTransaction) HasTxAmount() bool`

HasTxAmount returns a boolean if a field has been set.

### GetTxMode

`func (o *ResponseCustomerTransaction) GetTxMode() string`

GetTxMode returns the TxMode field if non-nil, zero value otherwise.

### GetTxModeOk

`func (o *ResponseCustomerTransaction) GetTxModeOk() (*string, bool)`

GetTxModeOk returns a tuple with the TxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMode

`func (o *ResponseCustomerTransaction) SetTxMode(v string)`

SetTxMode sets TxMode field to given value.

### HasTxMode

`func (o *ResponseCustomerTransaction) HasTxMode() bool`

HasTxMode returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseCustomerTransaction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseCustomerTransaction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseCustomerTransaction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseCustomerTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



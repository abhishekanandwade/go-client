# ResponseIpoSale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovedDate** | Pointer to **string** |  | [optional] 
**ApproverId** | Pointer to **int32** |  | [optional] 
**ApproverName** | Pointer to **string** |  | [optional] 
**ApproverRemarks** | Pointer to **string** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**OfficeName** | Pointer to **string** |  | [optional] 
**SaleAmount** | Pointer to **float32** |  | [optional] 
**SaleCommission** | Pointer to **float32** |  | [optional] 
**SaleDetails** | Pointer to **[]int32** |  | [optional] 
**SoldTo** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**Transdate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseIpoSale

`func NewResponseIpoSale() *ResponseIpoSale`

NewResponseIpoSale instantiates a new ResponseIpoSale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseIpoSaleWithDefaults

`func NewResponseIpoSaleWithDefaults() *ResponseIpoSale`

NewResponseIpoSaleWithDefaults instantiates a new ResponseIpoSale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovedDate

`func (o *ResponseIpoSale) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *ResponseIpoSale) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *ResponseIpoSale) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *ResponseIpoSale) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### GetApproverId

`func (o *ResponseIpoSale) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *ResponseIpoSale) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *ResponseIpoSale) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.

### HasApproverId

`func (o *ResponseIpoSale) HasApproverId() bool`

HasApproverId returns a boolean if a field has been set.

### GetApproverName

`func (o *ResponseIpoSale) GetApproverName() string`

GetApproverName returns the ApproverName field if non-nil, zero value otherwise.

### GetApproverNameOk

`func (o *ResponseIpoSale) GetApproverNameOk() (*string, bool)`

GetApproverNameOk returns a tuple with the ApproverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverName

`func (o *ResponseIpoSale) SetApproverName(v string)`

SetApproverName sets ApproverName field to given value.

### HasApproverName

`func (o *ResponseIpoSale) HasApproverName() bool`

HasApproverName returns a boolean if a field has been set.

### GetApproverRemarks

`func (o *ResponseIpoSale) GetApproverRemarks() string`

GetApproverRemarks returns the ApproverRemarks field if non-nil, zero value otherwise.

### GetApproverRemarksOk

`func (o *ResponseIpoSale) GetApproverRemarksOk() (*string, bool)`

GetApproverRemarksOk returns a tuple with the ApproverRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverRemarks

`func (o *ResponseIpoSale) SetApproverRemarks(v string)`

SetApproverRemarks sets ApproverRemarks field to given value.

### HasApproverRemarks

`func (o *ResponseIpoSale) HasApproverRemarks() bool`

HasApproverRemarks returns a boolean if a field has been set.

### GetIsApproved

`func (o *ResponseIpoSale) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *ResponseIpoSale) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *ResponseIpoSale) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *ResponseIpoSale) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseIpoSale) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseIpoSale) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseIpoSale) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseIpoSale) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOfficeName

`func (o *ResponseIpoSale) GetOfficeName() string`

GetOfficeName returns the OfficeName field if non-nil, zero value otherwise.

### GetOfficeNameOk

`func (o *ResponseIpoSale) GetOfficeNameOk() (*string, bool)`

GetOfficeNameOk returns a tuple with the OfficeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeName

`func (o *ResponseIpoSale) SetOfficeName(v string)`

SetOfficeName sets OfficeName field to given value.

### HasOfficeName

`func (o *ResponseIpoSale) HasOfficeName() bool`

HasOfficeName returns a boolean if a field has been set.

### GetSaleAmount

`func (o *ResponseIpoSale) GetSaleAmount() float32`

GetSaleAmount returns the SaleAmount field if non-nil, zero value otherwise.

### GetSaleAmountOk

`func (o *ResponseIpoSale) GetSaleAmountOk() (*float32, bool)`

GetSaleAmountOk returns a tuple with the SaleAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmount

`func (o *ResponseIpoSale) SetSaleAmount(v float32)`

SetSaleAmount sets SaleAmount field to given value.

### HasSaleAmount

`func (o *ResponseIpoSale) HasSaleAmount() bool`

HasSaleAmount returns a boolean if a field has been set.

### GetSaleCommission

`func (o *ResponseIpoSale) GetSaleCommission() float32`

GetSaleCommission returns the SaleCommission field if non-nil, zero value otherwise.

### GetSaleCommissionOk

`func (o *ResponseIpoSale) GetSaleCommissionOk() (*float32, bool)`

GetSaleCommissionOk returns a tuple with the SaleCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleCommission

`func (o *ResponseIpoSale) SetSaleCommission(v float32)`

SetSaleCommission sets SaleCommission field to given value.

### HasSaleCommission

`func (o *ResponseIpoSale) HasSaleCommission() bool`

HasSaleCommission returns a boolean if a field has been set.

### GetSaleDetails

`func (o *ResponseIpoSale) GetSaleDetails() []int32`

GetSaleDetails returns the SaleDetails field if non-nil, zero value otherwise.

### GetSaleDetailsOk

`func (o *ResponseIpoSale) GetSaleDetailsOk() (*[]int32, bool)`

GetSaleDetailsOk returns a tuple with the SaleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDetails

`func (o *ResponseIpoSale) SetSaleDetails(v []int32)`

SetSaleDetails sets SaleDetails field to given value.

### HasSaleDetails

`func (o *ResponseIpoSale) HasSaleDetails() bool`

HasSaleDetails returns a boolean if a field has been set.

### GetSoldTo

`func (o *ResponseIpoSale) GetSoldTo() string`

GetSoldTo returns the SoldTo field if non-nil, zero value otherwise.

### GetSoldToOk

`func (o *ResponseIpoSale) GetSoldToOk() (*string, bool)`

GetSoldToOk returns a tuple with the SoldTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldTo

`func (o *ResponseIpoSale) SetSoldTo(v string)`

SetSoldTo sets SoldTo field to given value.

### HasSoldTo

`func (o *ResponseIpoSale) HasSoldTo() bool`

HasSoldTo returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseIpoSale) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseIpoSale) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseIpoSale) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseIpoSale) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetTransactionId

`func (o *ResponseIpoSale) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *ResponseIpoSale) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *ResponseIpoSale) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *ResponseIpoSale) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetTransdate

`func (o *ResponseIpoSale) GetTransdate() string`

GetTransdate returns the Transdate field if non-nil, zero value otherwise.

### GetTransdateOk

`func (o *ResponseIpoSale) GetTransdateOk() (*string, bool)`

GetTransdateOk returns a tuple with the Transdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransdate

`func (o *ResponseIpoSale) SetTransdate(v string)`

SetTransdate sets Transdate field to given value.

### HasTransdate

`func (o *ResponseIpoSale) HasTransdate() bool`

HasTransdate returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseIpoSale) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseIpoSale) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseIpoSale) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseIpoSale) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ResponseIpoSale) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ResponseIpoSale) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ResponseIpoSale) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ResponseIpoSale) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



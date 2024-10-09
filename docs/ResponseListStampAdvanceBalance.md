# ResponseListStampAdvanceBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvanceIssued** | Pointer to **float32** |  | [optional] 
**AmountRealized** | Pointer to **float32** |  | [optional] 
**BalanceDetails** | Pointer to **map[string]int32** |  | [optional] 
**ClosingBalance** | Pointer to **float32** |  | [optional] 
**OfficeId** | Pointer to **int32** |  | [optional] 
**OpeningBalance** | Pointer to **float32** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**StampDetailsDesc** | Pointer to [**[]ResponseStampdetails**](ResponseStampdetails.md) |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseListStampAdvanceBalance

`func NewResponseListStampAdvanceBalance() *ResponseListStampAdvanceBalance`

NewResponseListStampAdvanceBalance instantiates a new ResponseListStampAdvanceBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListStampAdvanceBalanceWithDefaults

`func NewResponseListStampAdvanceBalanceWithDefaults() *ResponseListStampAdvanceBalance`

NewResponseListStampAdvanceBalanceWithDefaults instantiates a new ResponseListStampAdvanceBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvanceIssued

`func (o *ResponseListStampAdvanceBalance) GetAdvanceIssued() float32`

GetAdvanceIssued returns the AdvanceIssued field if non-nil, zero value otherwise.

### GetAdvanceIssuedOk

`func (o *ResponseListStampAdvanceBalance) GetAdvanceIssuedOk() (*float32, bool)`

GetAdvanceIssuedOk returns a tuple with the AdvanceIssued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceIssued

`func (o *ResponseListStampAdvanceBalance) SetAdvanceIssued(v float32)`

SetAdvanceIssued sets AdvanceIssued field to given value.

### HasAdvanceIssued

`func (o *ResponseListStampAdvanceBalance) HasAdvanceIssued() bool`

HasAdvanceIssued returns a boolean if a field has been set.

### GetAmountRealized

`func (o *ResponseListStampAdvanceBalance) GetAmountRealized() float32`

GetAmountRealized returns the AmountRealized field if non-nil, zero value otherwise.

### GetAmountRealizedOk

`func (o *ResponseListStampAdvanceBalance) GetAmountRealizedOk() (*float32, bool)`

GetAmountRealizedOk returns a tuple with the AmountRealized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountRealized

`func (o *ResponseListStampAdvanceBalance) SetAmountRealized(v float32)`

SetAmountRealized sets AmountRealized field to given value.

### HasAmountRealized

`func (o *ResponseListStampAdvanceBalance) HasAmountRealized() bool`

HasAmountRealized returns a boolean if a field has been set.

### GetBalanceDetails

`func (o *ResponseListStampAdvanceBalance) GetBalanceDetails() map[string]int32`

GetBalanceDetails returns the BalanceDetails field if non-nil, zero value otherwise.

### GetBalanceDetailsOk

`func (o *ResponseListStampAdvanceBalance) GetBalanceDetailsOk() (*map[string]int32, bool)`

GetBalanceDetailsOk returns a tuple with the BalanceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceDetails

`func (o *ResponseListStampAdvanceBalance) SetBalanceDetails(v map[string]int32)`

SetBalanceDetails sets BalanceDetails field to given value.

### HasBalanceDetails

`func (o *ResponseListStampAdvanceBalance) HasBalanceDetails() bool`

HasBalanceDetails returns a boolean if a field has been set.

### GetClosingBalance

`func (o *ResponseListStampAdvanceBalance) GetClosingBalance() float32`

GetClosingBalance returns the ClosingBalance field if non-nil, zero value otherwise.

### GetClosingBalanceOk

`func (o *ResponseListStampAdvanceBalance) GetClosingBalanceOk() (*float32, bool)`

GetClosingBalanceOk returns a tuple with the ClosingBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosingBalance

`func (o *ResponseListStampAdvanceBalance) SetClosingBalance(v float32)`

SetClosingBalance sets ClosingBalance field to given value.

### HasClosingBalance

`func (o *ResponseListStampAdvanceBalance) HasClosingBalance() bool`

HasClosingBalance returns a boolean if a field has been set.

### GetOfficeId

`func (o *ResponseListStampAdvanceBalance) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *ResponseListStampAdvanceBalance) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *ResponseListStampAdvanceBalance) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *ResponseListStampAdvanceBalance) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetOpeningBalance

`func (o *ResponseListStampAdvanceBalance) GetOpeningBalance() float32`

GetOpeningBalance returns the OpeningBalance field if non-nil, zero value otherwise.

### GetOpeningBalanceOk

`func (o *ResponseListStampAdvanceBalance) GetOpeningBalanceOk() (*float32, bool)`

GetOpeningBalanceOk returns a tuple with the OpeningBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpeningBalance

`func (o *ResponseListStampAdvanceBalance) SetOpeningBalance(v float32)`

SetOpeningBalance sets OpeningBalance field to given value.

### HasOpeningBalance

`func (o *ResponseListStampAdvanceBalance) HasOpeningBalance() bool`

HasOpeningBalance returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseListStampAdvanceBalance) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseListStampAdvanceBalance) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseListStampAdvanceBalance) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseListStampAdvanceBalance) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetStampDetailsDesc

`func (o *ResponseListStampAdvanceBalance) GetStampDetailsDesc() []ResponseStampdetails`

GetStampDetailsDesc returns the StampDetailsDesc field if non-nil, zero value otherwise.

### GetStampDetailsDescOk

`func (o *ResponseListStampAdvanceBalance) GetStampDetailsDescOk() (*[]ResponseStampdetails, bool)`

GetStampDetailsDescOk returns a tuple with the StampDetailsDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetailsDesc

`func (o *ResponseListStampAdvanceBalance) SetStampDetailsDesc(v []ResponseStampdetails)`

SetStampDetailsDesc sets StampDetailsDesc field to given value.

### HasStampDetailsDesc

`func (o *ResponseListStampAdvanceBalance) HasStampDetailsDesc() bool`

HasStampDetailsDesc returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseListStampAdvanceBalance) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseListStampAdvanceBalance) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseListStampAdvanceBalance) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseListStampAdvanceBalance) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseListStampAdvanceBalance) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseListStampAdvanceBalance) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseListStampAdvanceBalance) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseListStampAdvanceBalance) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ResponseListStampAdvanceBalance) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ResponseListStampAdvanceBalance) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ResponseListStampAdvanceBalance) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ResponseListStampAdvanceBalance) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



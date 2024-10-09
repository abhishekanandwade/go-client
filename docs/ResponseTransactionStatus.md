# ResponseTransactionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TransactionCount** | Pointer to **int32** |  | [optional] 
**TransactionType** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseTransactionStatus

`func NewResponseTransactionStatus() *ResponseTransactionStatus`

NewResponseTransactionStatus instantiates a new ResponseTransactionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseTransactionStatusWithDefaults

`func NewResponseTransactionStatusWithDefaults() *ResponseTransactionStatus`

NewResponseTransactionStatusWithDefaults instantiates a new ResponseTransactionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponseTransactionStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseTransactionStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseTransactionStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponseTransactionStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransactionCount

`func (o *ResponseTransactionStatus) GetTransactionCount() int32`

GetTransactionCount returns the TransactionCount field if non-nil, zero value otherwise.

### GetTransactionCountOk

`func (o *ResponseTransactionStatus) GetTransactionCountOk() (*int32, bool)`

GetTransactionCountOk returns a tuple with the TransactionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCount

`func (o *ResponseTransactionStatus) SetTransactionCount(v int32)`

SetTransactionCount sets TransactionCount field to given value.

### HasTransactionCount

`func (o *ResponseTransactionStatus) HasTransactionCount() bool`

HasTransactionCount returns a boolean if a field has been set.

### GetTransactionType

`func (o *ResponseTransactionStatus) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *ResponseTransactionStatus) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *ResponseTransactionStatus) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *ResponseTransactionStatus) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



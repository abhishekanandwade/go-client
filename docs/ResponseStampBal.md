# ResponseStampBal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastStampTransaction** | Pointer to [**ResponseStampsTransactionsIssued**](ResponseStampsTransactionsIssued.md) |  | [optional] 
**Stampbalance** | Pointer to [**[]ResponseStampBalance**](ResponseStampBalance.md) |  | [optional] 

## Methods

### NewResponseStampBal

`func NewResponseStampBal() *ResponseStampBal`

NewResponseStampBal instantiates a new ResponseStampBal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseStampBalWithDefaults

`func NewResponseStampBalWithDefaults() *ResponseStampBal`

NewResponseStampBalWithDefaults instantiates a new ResponseStampBal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastStampTransaction

`func (o *ResponseStampBal) GetLastStampTransaction() ResponseStampsTransactionsIssued`

GetLastStampTransaction returns the LastStampTransaction field if non-nil, zero value otherwise.

### GetLastStampTransactionOk

`func (o *ResponseStampBal) GetLastStampTransactionOk() (*ResponseStampsTransactionsIssued, bool)`

GetLastStampTransactionOk returns a tuple with the LastStampTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStampTransaction

`func (o *ResponseStampBal) SetLastStampTransaction(v ResponseStampsTransactionsIssued)`

SetLastStampTransaction sets LastStampTransaction field to given value.

### HasLastStampTransaction

`func (o *ResponseStampBal) HasLastStampTransaction() bool`

HasLastStampTransaction returns a boolean if a field has been set.

### GetStampbalance

`func (o *ResponseStampBal) GetStampbalance() []ResponseStampBalance`

GetStampbalance returns the Stampbalance field if non-nil, zero value otherwise.

### GetStampbalanceOk

`func (o *ResponseStampBal) GetStampbalanceOk() (*[]ResponseStampBalance, bool)`

GetStampbalanceOk returns a tuple with the Stampbalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampbalance

`func (o *ResponseStampBal) SetStampbalance(v []ResponseStampBalance)`

SetStampbalance sets Stampbalance field to given value.

### HasStampbalance

`func (o *ResponseStampBal) HasStampbalance() bool`

HasStampbalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



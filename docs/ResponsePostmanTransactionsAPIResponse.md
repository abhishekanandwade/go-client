# ResponsePostmanTransactionsAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponsePostmanTransactions**](ResponsePostmanTransactions.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponsePostmanTransactionsAPIResponse

`func NewResponsePostmanTransactionsAPIResponse() *ResponsePostmanTransactionsAPIResponse`

NewResponsePostmanTransactionsAPIResponse instantiates a new ResponsePostmanTransactionsAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePostmanTransactionsAPIResponseWithDefaults

`func NewResponsePostmanTransactionsAPIResponseWithDefaults() *ResponsePostmanTransactionsAPIResponse`

NewResponsePostmanTransactionsAPIResponseWithDefaults instantiates a new ResponsePostmanTransactionsAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponsePostmanTransactionsAPIResponse) GetData() []ResponsePostmanTransactions`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponsePostmanTransactionsAPIResponse) GetDataOk() (*[]ResponsePostmanTransactions, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponsePostmanTransactionsAPIResponse) SetData(v []ResponsePostmanTransactions)`

SetData sets Data field to given value.

### HasData

`func (o *ResponsePostmanTransactionsAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponsePostmanTransactionsAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponsePostmanTransactionsAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponsePostmanTransactionsAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponsePostmanTransactionsAPIResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponsePostmanTransactionsAPIResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponsePostmanTransactionsAPIResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponsePostmanTransactionsAPIResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponsePostmanTransactionsAPIResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



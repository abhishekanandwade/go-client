# ResponseListTcbBalanceAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseTcbBalance**](ResponseTcbBalance.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseListTcbBalanceAPIResponse

`func NewResponseListTcbBalanceAPIResponse() *ResponseListTcbBalanceAPIResponse`

NewResponseListTcbBalanceAPIResponse instantiates a new ResponseListTcbBalanceAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListTcbBalanceAPIResponseWithDefaults

`func NewResponseListTcbBalanceAPIResponseWithDefaults() *ResponseListTcbBalanceAPIResponse`

NewResponseListTcbBalanceAPIResponseWithDefaults instantiates a new ResponseListTcbBalanceAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseListTcbBalanceAPIResponse) GetData() []ResponseTcbBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseListTcbBalanceAPIResponse) GetDataOk() (*[]ResponseTcbBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseListTcbBalanceAPIResponse) SetData(v []ResponseTcbBalance)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseListTcbBalanceAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseListTcbBalanceAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseListTcbBalanceAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseListTcbBalanceAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseListTcbBalanceAPIResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseListTcbBalanceAPIResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseListTcbBalanceAPIResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseListTcbBalanceAPIResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseListTcbBalanceAPIResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



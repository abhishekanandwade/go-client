# ResponseListIPOTransactionsApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseIpoTxnsResponse**](ResponseIpoTxnsResponse.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**ReturnedRecordsCount** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**SortType** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**TotalRecordsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseListIPOTransactionsApiResponse

`func NewResponseListIPOTransactionsApiResponse() *ResponseListIPOTransactionsApiResponse`

NewResponseListIPOTransactionsApiResponse instantiates a new ResponseListIPOTransactionsApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListIPOTransactionsApiResponseWithDefaults

`func NewResponseListIPOTransactionsApiResponseWithDefaults() *ResponseListIPOTransactionsApiResponse`

NewResponseListIPOTransactionsApiResponseWithDefaults instantiates a new ResponseListIPOTransactionsApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseListIPOTransactionsApiResponse) GetData() []ResponseIpoTxnsResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseListIPOTransactionsApiResponse) GetDataOk() (*[]ResponseIpoTxnsResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseListIPOTransactionsApiResponse) SetData(v []ResponseIpoTxnsResponse)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseListIPOTransactionsApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseListIPOTransactionsApiResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseListIPOTransactionsApiResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseListIPOTransactionsApiResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseListIPOTransactionsApiResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseListIPOTransactionsApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseListIPOTransactionsApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseListIPOTransactionsApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseListIPOTransactionsApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrderBy

`func (o *ResponseListIPOTransactionsApiResponse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ResponseListIPOTransactionsApiResponse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ResponseListIPOTransactionsApiResponse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ResponseListIPOTransactionsApiResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetReturnedRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) GetReturnedRecordsCount() int32`

GetReturnedRecordsCount returns the ReturnedRecordsCount field if non-nil, zero value otherwise.

### GetReturnedRecordsCountOk

`func (o *ResponseListIPOTransactionsApiResponse) GetReturnedRecordsCountOk() (*int32, bool)`

GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) SetReturnedRecordsCount(v int32)`

SetReturnedRecordsCount sets ReturnedRecordsCount field to given value.

### HasReturnedRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) HasReturnedRecordsCount() bool`

HasReturnedRecordsCount returns a boolean if a field has been set.

### GetSkip

`func (o *ResponseListIPOTransactionsApiResponse) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ResponseListIPOTransactionsApiResponse) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ResponseListIPOTransactionsApiResponse) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ResponseListIPOTransactionsApiResponse) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSortType

`func (o *ResponseListIPOTransactionsApiResponse) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *ResponseListIPOTransactionsApiResponse) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *ResponseListIPOTransactionsApiResponse) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *ResponseListIPOTransactionsApiResponse) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseListIPOTransactionsApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseListIPOTransactionsApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseListIPOTransactionsApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseListIPOTransactionsApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTotalRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) GetTotalRecordsCount() int32`

GetTotalRecordsCount returns the TotalRecordsCount field if non-nil, zero value otherwise.

### GetTotalRecordsCountOk

`func (o *ResponseListIPOTransactionsApiResponse) GetTotalRecordsCountOk() (*int32, bool)`

GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) SetTotalRecordsCount(v int32)`

SetTotalRecordsCount sets TotalRecordsCount field to given value.

### HasTotalRecordsCount

`func (o *ResponseListIPOTransactionsApiResponse) HasTotalRecordsCount() bool`

HasTotalRecordsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



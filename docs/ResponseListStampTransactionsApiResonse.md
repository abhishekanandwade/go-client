# ResponseListStampTransactionsApiResonse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseStampsTransactionsIssued**](ResponseStampsTransactionsIssued.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**ReturnedRecordsCount** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**SortType** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**TotalRecordsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseListStampTransactionsApiResonse

`func NewResponseListStampTransactionsApiResonse() *ResponseListStampTransactionsApiResonse`

NewResponseListStampTransactionsApiResonse instantiates a new ResponseListStampTransactionsApiResonse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListStampTransactionsApiResonseWithDefaults

`func NewResponseListStampTransactionsApiResonseWithDefaults() *ResponseListStampTransactionsApiResonse`

NewResponseListStampTransactionsApiResonseWithDefaults instantiates a new ResponseListStampTransactionsApiResonse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseListStampTransactionsApiResonse) GetData() []ResponseStampsTransactionsIssued`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseListStampTransactionsApiResonse) GetDataOk() (*[]ResponseStampsTransactionsIssued, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseListStampTransactionsApiResonse) SetData(v []ResponseStampsTransactionsIssued)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseListStampTransactionsApiResonse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseListStampTransactionsApiResonse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseListStampTransactionsApiResonse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseListStampTransactionsApiResonse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseListStampTransactionsApiResonse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseListStampTransactionsApiResonse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseListStampTransactionsApiResonse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseListStampTransactionsApiResonse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseListStampTransactionsApiResonse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrderBy

`func (o *ResponseListStampTransactionsApiResonse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ResponseListStampTransactionsApiResonse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ResponseListStampTransactionsApiResonse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ResponseListStampTransactionsApiResonse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetReturnedRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) GetReturnedRecordsCount() int32`

GetReturnedRecordsCount returns the ReturnedRecordsCount field if non-nil, zero value otherwise.

### GetReturnedRecordsCountOk

`func (o *ResponseListStampTransactionsApiResonse) GetReturnedRecordsCountOk() (*int32, bool)`

GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) SetReturnedRecordsCount(v int32)`

SetReturnedRecordsCount sets ReturnedRecordsCount field to given value.

### HasReturnedRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) HasReturnedRecordsCount() bool`

HasReturnedRecordsCount returns a boolean if a field has been set.

### GetSkip

`func (o *ResponseListStampTransactionsApiResonse) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ResponseListStampTransactionsApiResonse) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ResponseListStampTransactionsApiResonse) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ResponseListStampTransactionsApiResonse) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSortType

`func (o *ResponseListStampTransactionsApiResonse) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *ResponseListStampTransactionsApiResonse) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *ResponseListStampTransactionsApiResonse) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *ResponseListStampTransactionsApiResonse) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseListStampTransactionsApiResonse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseListStampTransactionsApiResonse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseListStampTransactionsApiResonse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseListStampTransactionsApiResonse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTotalRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) GetTotalRecordsCount() int32`

GetTotalRecordsCount returns the TotalRecordsCount field if non-nil, zero value otherwise.

### GetTotalRecordsCountOk

`func (o *ResponseListStampTransactionsApiResonse) GetTotalRecordsCountOk() (*int32, bool)`

GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) SetTotalRecordsCount(v int32)`

SetTotalRecordsCount sets TotalRecordsCount field to given value.

### HasTotalRecordsCount

`func (o *ResponseListStampTransactionsApiResonse) HasTotalRecordsCount() bool`

HasTotalRecordsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



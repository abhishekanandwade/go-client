# ResponseListCashErrorsApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseCashError**](ResponseCashError.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**ReturnedRecordsCount** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**SortType** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**TotalRecordsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseListCashErrorsApiResponse

`func NewResponseListCashErrorsApiResponse() *ResponseListCashErrorsApiResponse`

NewResponseListCashErrorsApiResponse instantiates a new ResponseListCashErrorsApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListCashErrorsApiResponseWithDefaults

`func NewResponseListCashErrorsApiResponseWithDefaults() *ResponseListCashErrorsApiResponse`

NewResponseListCashErrorsApiResponseWithDefaults instantiates a new ResponseListCashErrorsApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseListCashErrorsApiResponse) GetData() []ResponseCashError`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseListCashErrorsApiResponse) GetDataOk() (*[]ResponseCashError, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseListCashErrorsApiResponse) SetData(v []ResponseCashError)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseListCashErrorsApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseListCashErrorsApiResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseListCashErrorsApiResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseListCashErrorsApiResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseListCashErrorsApiResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseListCashErrorsApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseListCashErrorsApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseListCashErrorsApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseListCashErrorsApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrderBy

`func (o *ResponseListCashErrorsApiResponse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ResponseListCashErrorsApiResponse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ResponseListCashErrorsApiResponse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ResponseListCashErrorsApiResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetReturnedRecordsCount

`func (o *ResponseListCashErrorsApiResponse) GetReturnedRecordsCount() int32`

GetReturnedRecordsCount returns the ReturnedRecordsCount field if non-nil, zero value otherwise.

### GetReturnedRecordsCountOk

`func (o *ResponseListCashErrorsApiResponse) GetReturnedRecordsCountOk() (*int32, bool)`

GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedRecordsCount

`func (o *ResponseListCashErrorsApiResponse) SetReturnedRecordsCount(v int32)`

SetReturnedRecordsCount sets ReturnedRecordsCount field to given value.

### HasReturnedRecordsCount

`func (o *ResponseListCashErrorsApiResponse) HasReturnedRecordsCount() bool`

HasReturnedRecordsCount returns a boolean if a field has been set.

### GetSkip

`func (o *ResponseListCashErrorsApiResponse) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ResponseListCashErrorsApiResponse) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ResponseListCashErrorsApiResponse) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ResponseListCashErrorsApiResponse) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSortType

`func (o *ResponseListCashErrorsApiResponse) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *ResponseListCashErrorsApiResponse) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *ResponseListCashErrorsApiResponse) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *ResponseListCashErrorsApiResponse) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseListCashErrorsApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseListCashErrorsApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseListCashErrorsApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseListCashErrorsApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTotalRecordsCount

`func (o *ResponseListCashErrorsApiResponse) GetTotalRecordsCount() int32`

GetTotalRecordsCount returns the TotalRecordsCount field if non-nil, zero value otherwise.

### GetTotalRecordsCountOk

`func (o *ResponseListCashErrorsApiResponse) GetTotalRecordsCountOk() (*int32, bool)`

GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsCount

`func (o *ResponseListCashErrorsApiResponse) SetTotalRecordsCount(v int32)`

SetTotalRecordsCount sets TotalRecordsCount field to given value.

### HasTotalRecordsCount

`func (o *ResponseListCashErrorsApiResponse) HasTotalRecordsCount() bool`

HasTotalRecordsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



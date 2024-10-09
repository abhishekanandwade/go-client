# ResponseListBankCreditLimitAPIResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ResponseBankCreditLimit**](ResponseBankCreditLimit.md) |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**ReturnedRecordsCount** | Pointer to **int32** |  | [optional] 
**Skip** | Pointer to **int32** |  | [optional] 
**SortType** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**TotalRecordsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseListBankCreditLimitAPIResponse

`func NewResponseListBankCreditLimitAPIResponse() *ResponseListBankCreditLimitAPIResponse`

NewResponseListBankCreditLimitAPIResponse instantiates a new ResponseListBankCreditLimitAPIResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseListBankCreditLimitAPIResponseWithDefaults

`func NewResponseListBankCreditLimitAPIResponseWithDefaults() *ResponseListBankCreditLimitAPIResponse`

NewResponseListBankCreditLimitAPIResponseWithDefaults instantiates a new ResponseListBankCreditLimitAPIResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseListBankCreditLimitAPIResponse) GetData() []ResponseBankCreditLimit`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetDataOk() (*[]ResponseBankCreditLimit, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseListBankCreditLimitAPIResponse) SetData(v []ResponseBankCreditLimit)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseListBankCreditLimitAPIResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLimit

`func (o *ResponseListBankCreditLimitAPIResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseListBankCreditLimitAPIResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ResponseListBankCreditLimitAPIResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseListBankCreditLimitAPIResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseListBankCreditLimitAPIResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseListBankCreditLimitAPIResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrderBy

`func (o *ResponseListBankCreditLimitAPIResponse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *ResponseListBankCreditLimitAPIResponse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *ResponseListBankCreditLimitAPIResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetReturnedRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) GetReturnedRecordsCount() int32`

GetReturnedRecordsCount returns the ReturnedRecordsCount field if non-nil, zero value otherwise.

### GetReturnedRecordsCountOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetReturnedRecordsCountOk() (*int32, bool)`

GetReturnedRecordsCountOk returns a tuple with the ReturnedRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) SetReturnedRecordsCount(v int32)`

SetReturnedRecordsCount sets ReturnedRecordsCount field to given value.

### HasReturnedRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) HasReturnedRecordsCount() bool`

HasReturnedRecordsCount returns a boolean if a field has been set.

### GetSkip

`func (o *ResponseListBankCreditLimitAPIResponse) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ResponseListBankCreditLimitAPIResponse) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ResponseListBankCreditLimitAPIResponse) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetSortType

`func (o *ResponseListBankCreditLimitAPIResponse) GetSortType() string`

GetSortType returns the SortType field if non-nil, zero value otherwise.

### GetSortTypeOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetSortTypeOk() (*string, bool)`

GetSortTypeOk returns a tuple with the SortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortType

`func (o *ResponseListBankCreditLimitAPIResponse) SetSortType(v string)`

SetSortType sets SortType field to given value.

### HasSortType

`func (o *ResponseListBankCreditLimitAPIResponse) HasSortType() bool`

HasSortType returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseListBankCreditLimitAPIResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseListBankCreditLimitAPIResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseListBankCreditLimitAPIResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTotalRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) GetTotalRecordsCount() int32`

GetTotalRecordsCount returns the TotalRecordsCount field if non-nil, zero value otherwise.

### GetTotalRecordsCountOk

`func (o *ResponseListBankCreditLimitAPIResponse) GetTotalRecordsCountOk() (*int32, bool)`

GetTotalRecordsCountOk returns a tuple with the TotalRecordsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) SetTotalRecordsCount(v int32)`

SetTotalRecordsCount sets TotalRecordsCount field to given value.

### HasTotalRecordsCount

`func (o *ResponseListBankCreditLimitAPIResponse) HasTotalRecordsCount() bool`

HasTotalRecordsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



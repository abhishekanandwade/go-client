# ResponseFetchSummaryDetailsApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseSummaryData**](ResponseSummaryData.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseFetchSummaryDetailsApiResponse

`func NewResponseFetchSummaryDetailsApiResponse() *ResponseFetchSummaryDetailsApiResponse`

NewResponseFetchSummaryDetailsApiResponse instantiates a new ResponseFetchSummaryDetailsApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseFetchSummaryDetailsApiResponseWithDefaults

`func NewResponseFetchSummaryDetailsApiResponseWithDefaults() *ResponseFetchSummaryDetailsApiResponse`

NewResponseFetchSummaryDetailsApiResponseWithDefaults instantiates a new ResponseFetchSummaryDetailsApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseFetchSummaryDetailsApiResponse) GetData() ResponseSummaryData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseFetchSummaryDetailsApiResponse) GetDataOk() (*ResponseSummaryData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseFetchSummaryDetailsApiResponse) SetData(v ResponseSummaryData)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseFetchSummaryDetailsApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseFetchSummaryDetailsApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseFetchSummaryDetailsApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseFetchSummaryDetailsApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseFetchSummaryDetailsApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseFetchSummaryDetailsApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseFetchSummaryDetailsApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseFetchSummaryDetailsApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseFetchSummaryDetailsApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



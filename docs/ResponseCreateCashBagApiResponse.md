# ResponseCreateCashBagApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseCashBag**](ResponseCashBag.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseCreateCashBagApiResponse

`func NewResponseCreateCashBagApiResponse() *ResponseCreateCashBagApiResponse`

NewResponseCreateCashBagApiResponse instantiates a new ResponseCreateCashBagApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCreateCashBagApiResponseWithDefaults

`func NewResponseCreateCashBagApiResponseWithDefaults() *ResponseCreateCashBagApiResponse`

NewResponseCreateCashBagApiResponseWithDefaults instantiates a new ResponseCreateCashBagApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseCreateCashBagApiResponse) GetData() ResponseCashBag`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseCreateCashBagApiResponse) GetDataOk() (*ResponseCashBag, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseCreateCashBagApiResponse) SetData(v ResponseCashBag)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseCreateCashBagApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseCreateCashBagApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseCreateCashBagApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseCreateCashBagApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseCreateCashBagApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseCreateCashBagApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseCreateCashBagApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseCreateCashBagApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseCreateCashBagApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



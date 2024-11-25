# ResponseSendChequesToBankOrAOApiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ResponseChequeRemittance**](ResponseChequeRemittance.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseSendChequesToBankOrAOApiResponse

`func NewResponseSendChequesToBankOrAOApiResponse() *ResponseSendChequesToBankOrAOApiResponse`

NewResponseSendChequesToBankOrAOApiResponse instantiates a new ResponseSendChequesToBankOrAOApiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSendChequesToBankOrAOApiResponseWithDefaults

`func NewResponseSendChequesToBankOrAOApiResponseWithDefaults() *ResponseSendChequesToBankOrAOApiResponse`

NewResponseSendChequesToBankOrAOApiResponseWithDefaults instantiates a new ResponseSendChequesToBankOrAOApiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetData() ResponseChequeRemittance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetDataOk() (*ResponseChequeRemittance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResponseSendChequesToBankOrAOApiResponse) SetData(v ResponseChequeRemittance)`

SetData sets Data field to given value.

### HasData

`func (o *ResponseSendChequesToBankOrAOApiResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResponseSendChequesToBankOrAOApiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResponseSendChequesToBankOrAOApiResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseSendChequesToBankOrAOApiResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseSendChequesToBankOrAOApiResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseSendChequesToBankOrAOApiResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ApierrorsAppError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**FieldErrors** | Pointer to [**[]ApierrorsFieldError**](ApierrorsFieldError.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewApierrorsAppError

`func NewApierrorsAppError() *ApierrorsAppError`

NewApierrorsAppError instantiates a new ApierrorsAppError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApierrorsAppErrorWithDefaults

`func NewApierrorsAppErrorWithDefaults() *ApierrorsAppError`

NewApierrorsAppErrorWithDefaults instantiates a new ApierrorsAppError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApierrorsAppError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApierrorsAppError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApierrorsAppError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApierrorsAppError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldErrors

`func (o *ApierrorsAppError) GetFieldErrors() []ApierrorsFieldError`

GetFieldErrors returns the FieldErrors field if non-nil, zero value otherwise.

### GetFieldErrorsOk

`func (o *ApierrorsAppError) GetFieldErrorsOk() (*[]ApierrorsFieldError, bool)`

GetFieldErrorsOk returns a tuple with the FieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldErrors

`func (o *ApierrorsAppError) SetFieldErrors(v []ApierrorsFieldError)`

SetFieldErrors sets FieldErrors field to given value.

### HasFieldErrors

`func (o *ApierrorsAppError) HasFieldErrors() bool`

HasFieldErrors returns a boolean if a field has been set.

### GetId

`func (o *ApierrorsAppError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApierrorsAppError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApierrorsAppError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApierrorsAppError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ApierrorsAppError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApierrorsAppError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApierrorsAppError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApierrorsAppError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



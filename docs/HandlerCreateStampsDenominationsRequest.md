# HandlerCreateStampsDenominationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** |  | 
**DenominationDesc** | **string** |  | 
**DenominationValue** | **float32** |  | 
**EnteredBy** | **string** |  | 
**OfficeId** | Pointer to **int32** |  | [optional] 
**ValidFrom** | **string** |  | 

## Methods

### NewHandlerCreateStampsDenominationsRequest

`func NewHandlerCreateStampsDenominationsRequest(categoryId string, denominationDesc string, denominationValue float32, enteredBy string, validFrom string, ) *HandlerCreateStampsDenominationsRequest`

NewHandlerCreateStampsDenominationsRequest instantiates a new HandlerCreateStampsDenominationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateStampsDenominationsRequestWithDefaults

`func NewHandlerCreateStampsDenominationsRequestWithDefaults() *HandlerCreateStampsDenominationsRequest`

NewHandlerCreateStampsDenominationsRequestWithDefaults instantiates a new HandlerCreateStampsDenominationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *HandlerCreateStampsDenominationsRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *HandlerCreateStampsDenominationsRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *HandlerCreateStampsDenominationsRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetDenominationDesc

`func (o *HandlerCreateStampsDenominationsRequest) GetDenominationDesc() string`

GetDenominationDesc returns the DenominationDesc field if non-nil, zero value otherwise.

### GetDenominationDescOk

`func (o *HandlerCreateStampsDenominationsRequest) GetDenominationDescOk() (*string, bool)`

GetDenominationDescOk returns a tuple with the DenominationDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationDesc

`func (o *HandlerCreateStampsDenominationsRequest) SetDenominationDesc(v string)`

SetDenominationDesc sets DenominationDesc field to given value.


### GetDenominationValue

`func (o *HandlerCreateStampsDenominationsRequest) GetDenominationValue() float32`

GetDenominationValue returns the DenominationValue field if non-nil, zero value otherwise.

### GetDenominationValueOk

`func (o *HandlerCreateStampsDenominationsRequest) GetDenominationValueOk() (*float32, bool)`

GetDenominationValueOk returns a tuple with the DenominationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominationValue

`func (o *HandlerCreateStampsDenominationsRequest) SetDenominationValue(v float32)`

SetDenominationValue sets DenominationValue field to given value.


### GetEnteredBy

`func (o *HandlerCreateStampsDenominationsRequest) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *HandlerCreateStampsDenominationsRequest) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *HandlerCreateStampsDenominationsRequest) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.


### GetOfficeId

`func (o *HandlerCreateStampsDenominationsRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateStampsDenominationsRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateStampsDenominationsRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.

### HasOfficeId

`func (o *HandlerCreateStampsDenominationsRequest) HasOfficeId() bool`

HasOfficeId returns a boolean if a field has been set.

### GetValidFrom

`func (o *HandlerCreateStampsDenominationsRequest) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *HandlerCreateStampsDenominationsRequest) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *HandlerCreateStampsDenominationsRequest) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



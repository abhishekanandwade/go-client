# ResponseWorkflowResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SuccessfulResponses** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewResponseWorkflowResult

`func NewResponseWorkflowResult() *ResponseWorkflowResult`

NewResponseWorkflowResult instantiates a new ResponseWorkflowResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseWorkflowResultWithDefaults

`func NewResponseWorkflowResultWithDefaults() *ResponseWorkflowResult`

NewResponseWorkflowResultWithDefaults instantiates a new ResponseWorkflowResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ResponseWorkflowResult) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ResponseWorkflowResult) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ResponseWorkflowResult) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ResponseWorkflowResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccessfulResponses

`func (o *ResponseWorkflowResult) GetSuccessfulResponses() []map[string]interface{}`

GetSuccessfulResponses returns the SuccessfulResponses field if non-nil, zero value otherwise.

### GetSuccessfulResponsesOk

`func (o *ResponseWorkflowResult) GetSuccessfulResponsesOk() (*[]map[string]interface{}, bool)`

GetSuccessfulResponsesOk returns a tuple with the SuccessfulResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulResponses

`func (o *ResponseWorkflowResult) SetSuccessfulResponses(v []map[string]interface{})`

SetSuccessfulResponses sets SuccessfulResponses field to given value.

### HasSuccessfulResponses

`func (o *ResponseWorkflowResult) HasSuccessfulResponses() bool`

HasSuccessfulResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



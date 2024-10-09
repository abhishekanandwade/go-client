# HandlerUpdateTransactionsStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amtount** | Pointer to **float32** |  | [optional] 
**Details** | Pointer to **map[string]int32** |  | [optional] 
**EmployeeId1** | Pointer to **int32** |  | [optional] 
**EmployeeId2** | Pointer to **int32** |  | [optional] 
**IsApproved** | Pointer to **bool** |  | [optional] 
**IsNotApproved** | Pointer to **bool** |  | [optional] 
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RequestType** | **string** |  | 
**Type** | [**HandlerTransactionType**](HandlerTransactionType.md) |  | 
**UserId** | **int32** |  | 

## Methods

### NewHandlerUpdateTransactionsStatusRequest

`func NewHandlerUpdateTransactionsStatusRequest(requestType string, type_ HandlerTransactionType, userId int32, ) *HandlerUpdateTransactionsStatusRequest`

NewHandlerUpdateTransactionsStatusRequest instantiates a new HandlerUpdateTransactionsStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerUpdateTransactionsStatusRequestWithDefaults

`func NewHandlerUpdateTransactionsStatusRequestWithDefaults() *HandlerUpdateTransactionsStatusRequest`

NewHandlerUpdateTransactionsStatusRequestWithDefaults instantiates a new HandlerUpdateTransactionsStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmtount

`func (o *HandlerUpdateTransactionsStatusRequest) GetAmtount() float32`

GetAmtount returns the Amtount field if non-nil, zero value otherwise.

### GetAmtountOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetAmtountOk() (*float32, bool)`

GetAmtountOk returns a tuple with the Amtount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmtount

`func (o *HandlerUpdateTransactionsStatusRequest) SetAmtount(v float32)`

SetAmtount sets Amtount field to given value.

### HasAmtount

`func (o *HandlerUpdateTransactionsStatusRequest) HasAmtount() bool`

HasAmtount returns a boolean if a field has been set.

### GetDetails

`func (o *HandlerUpdateTransactionsStatusRequest) GetDetails() map[string]int32`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetDetailsOk() (*map[string]int32, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HandlerUpdateTransactionsStatusRequest) SetDetails(v map[string]int32)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HandlerUpdateTransactionsStatusRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEmployeeId1

`func (o *HandlerUpdateTransactionsStatusRequest) GetEmployeeId1() int32`

GetEmployeeId1 returns the EmployeeId1 field if non-nil, zero value otherwise.

### GetEmployeeId1Ok

`func (o *HandlerUpdateTransactionsStatusRequest) GetEmployeeId1Ok() (*int32, bool)`

GetEmployeeId1Ok returns a tuple with the EmployeeId1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId1

`func (o *HandlerUpdateTransactionsStatusRequest) SetEmployeeId1(v int32)`

SetEmployeeId1 sets EmployeeId1 field to given value.

### HasEmployeeId1

`func (o *HandlerUpdateTransactionsStatusRequest) HasEmployeeId1() bool`

HasEmployeeId1 returns a boolean if a field has been set.

### GetEmployeeId2

`func (o *HandlerUpdateTransactionsStatusRequest) GetEmployeeId2() int32`

GetEmployeeId2 returns the EmployeeId2 field if non-nil, zero value otherwise.

### GetEmployeeId2Ok

`func (o *HandlerUpdateTransactionsStatusRequest) GetEmployeeId2Ok() (*int32, bool)`

GetEmployeeId2Ok returns a tuple with the EmployeeId2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeId2

`func (o *HandlerUpdateTransactionsStatusRequest) SetEmployeeId2(v int32)`

SetEmployeeId2 sets EmployeeId2 field to given value.

### HasEmployeeId2

`func (o *HandlerUpdateTransactionsStatusRequest) HasEmployeeId2() bool`

HasEmployeeId2 returns a boolean if a field has been set.

### GetIsApproved

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsApproved() bool`

GetIsApproved returns the IsApproved field if non-nil, zero value otherwise.

### GetIsApprovedOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsApprovedOk() (*bool, bool)`

GetIsApprovedOk returns a tuple with the IsApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApproved

`func (o *HandlerUpdateTransactionsStatusRequest) SetIsApproved(v bool)`

SetIsApproved sets IsApproved field to given value.

### HasIsApproved

`func (o *HandlerUpdateTransactionsStatusRequest) HasIsApproved() bool`

HasIsApproved returns a boolean if a field has been set.

### GetIsNotApproved

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsNotApproved() bool`

GetIsNotApproved returns the IsNotApproved field if non-nil, zero value otherwise.

### GetIsNotApprovedOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsNotApprovedOk() (*bool, bool)`

GetIsNotApprovedOk returns a tuple with the IsNotApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNotApproved

`func (o *HandlerUpdateTransactionsStatusRequest) SetIsNotApproved(v bool)`

SetIsNotApproved sets IsNotApproved field to given value.

### HasIsNotApproved

`func (o *HandlerUpdateTransactionsStatusRequest) HasIsNotApproved() bool`

HasIsNotApproved returns a boolean if a field has been set.

### GetIsSingleHand

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerUpdateTransactionsStatusRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerUpdateTransactionsStatusRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetRemarks

`func (o *HandlerUpdateTransactionsStatusRequest) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HandlerUpdateTransactionsStatusRequest) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *HandlerUpdateTransactionsStatusRequest) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRequestType

`func (o *HandlerUpdateTransactionsStatusRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *HandlerUpdateTransactionsStatusRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.


### GetType

`func (o *HandlerUpdateTransactionsStatusRequest) GetType() HandlerTransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetTypeOk() (*HandlerTransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HandlerUpdateTransactionsStatusRequest) SetType(v HandlerTransactionType)`

SetType sets Type field to given value.


### GetUserId

`func (o *HandlerUpdateTransactionsStatusRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerUpdateTransactionsStatusRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerUpdateTransactionsStatusRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



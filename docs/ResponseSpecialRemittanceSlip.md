# ResponseSpecialRemittanceSlip

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyDetails** | Pointer to [**[]ResponseCurrencyDetails**](ResponseCurrencyDetails.md) |  | [optional] 
**EmployeeName1** | Pointer to **string** |  | [optional] 
**EmployeeName2** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**RemittanceId** | Pointer to **string** |  | [optional] 
**ReqApproverAmt** | Pointer to **float32** |  | [optional] 
**ToOffice** | Pointer to **string** |  | [optional] 
**TransDate** | Pointer to **string** |  | [optional] 

## Methods

### NewResponseSpecialRemittanceSlip

`func NewResponseSpecialRemittanceSlip() *ResponseSpecialRemittanceSlip`

NewResponseSpecialRemittanceSlip instantiates a new ResponseSpecialRemittanceSlip object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSpecialRemittanceSlipWithDefaults

`func NewResponseSpecialRemittanceSlipWithDefaults() *ResponseSpecialRemittanceSlip`

NewResponseSpecialRemittanceSlipWithDefaults instantiates a new ResponseSpecialRemittanceSlip object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyDetails

`func (o *ResponseSpecialRemittanceSlip) GetCurrencyDetails() []ResponseCurrencyDetails`

GetCurrencyDetails returns the CurrencyDetails field if non-nil, zero value otherwise.

### GetCurrencyDetailsOk

`func (o *ResponseSpecialRemittanceSlip) GetCurrencyDetailsOk() (*[]ResponseCurrencyDetails, bool)`

GetCurrencyDetailsOk returns a tuple with the CurrencyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyDetails

`func (o *ResponseSpecialRemittanceSlip) SetCurrencyDetails(v []ResponseCurrencyDetails)`

SetCurrencyDetails sets CurrencyDetails field to given value.

### HasCurrencyDetails

`func (o *ResponseSpecialRemittanceSlip) HasCurrencyDetails() bool`

HasCurrencyDetails returns a boolean if a field has been set.

### GetEmployeeName1

`func (o *ResponseSpecialRemittanceSlip) GetEmployeeName1() string`

GetEmployeeName1 returns the EmployeeName1 field if non-nil, zero value otherwise.

### GetEmployeeName1Ok

`func (o *ResponseSpecialRemittanceSlip) GetEmployeeName1Ok() (*string, bool)`

GetEmployeeName1Ok returns a tuple with the EmployeeName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName1

`func (o *ResponseSpecialRemittanceSlip) SetEmployeeName1(v string)`

SetEmployeeName1 sets EmployeeName1 field to given value.

### HasEmployeeName1

`func (o *ResponseSpecialRemittanceSlip) HasEmployeeName1() bool`

HasEmployeeName1 returns a boolean if a field has been set.

### GetEmployeeName2

`func (o *ResponseSpecialRemittanceSlip) GetEmployeeName2() string`

GetEmployeeName2 returns the EmployeeName2 field if non-nil, zero value otherwise.

### GetEmployeeName2Ok

`func (o *ResponseSpecialRemittanceSlip) GetEmployeeName2Ok() (*string, bool)`

GetEmployeeName2Ok returns a tuple with the EmployeeName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName2

`func (o *ResponseSpecialRemittanceSlip) SetEmployeeName2(v string)`

SetEmployeeName2 sets EmployeeName2 field to given value.

### HasEmployeeName2

`func (o *ResponseSpecialRemittanceSlip) HasEmployeeName2() bool`

HasEmployeeName2 returns a boolean if a field has been set.

### GetRemarks

`func (o *ResponseSpecialRemittanceSlip) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ResponseSpecialRemittanceSlip) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ResponseSpecialRemittanceSlip) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ResponseSpecialRemittanceSlip) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetRemittanceId

`func (o *ResponseSpecialRemittanceSlip) GetRemittanceId() string`

GetRemittanceId returns the RemittanceId field if non-nil, zero value otherwise.

### GetRemittanceIdOk

`func (o *ResponseSpecialRemittanceSlip) GetRemittanceIdOk() (*string, bool)`

GetRemittanceIdOk returns a tuple with the RemittanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemittanceId

`func (o *ResponseSpecialRemittanceSlip) SetRemittanceId(v string)`

SetRemittanceId sets RemittanceId field to given value.

### HasRemittanceId

`func (o *ResponseSpecialRemittanceSlip) HasRemittanceId() bool`

HasRemittanceId returns a boolean if a field has been set.

### GetReqApproverAmt

`func (o *ResponseSpecialRemittanceSlip) GetReqApproverAmt() float32`

GetReqApproverAmt returns the ReqApproverAmt field if non-nil, zero value otherwise.

### GetReqApproverAmtOk

`func (o *ResponseSpecialRemittanceSlip) GetReqApproverAmtOk() (*float32, bool)`

GetReqApproverAmtOk returns a tuple with the ReqApproverAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqApproverAmt

`func (o *ResponseSpecialRemittanceSlip) SetReqApproverAmt(v float32)`

SetReqApproverAmt sets ReqApproverAmt field to given value.

### HasReqApproverAmt

`func (o *ResponseSpecialRemittanceSlip) HasReqApproverAmt() bool`

HasReqApproverAmt returns a boolean if a field has been set.

### GetToOffice

`func (o *ResponseSpecialRemittanceSlip) GetToOffice() string`

GetToOffice returns the ToOffice field if non-nil, zero value otherwise.

### GetToOfficeOk

`func (o *ResponseSpecialRemittanceSlip) GetToOfficeOk() (*string, bool)`

GetToOfficeOk returns a tuple with the ToOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOffice

`func (o *ResponseSpecialRemittanceSlip) SetToOffice(v string)`

SetToOffice sets ToOffice field to given value.

### HasToOffice

`func (o *ResponseSpecialRemittanceSlip) HasToOffice() bool`

HasToOffice returns a boolean if a field has been set.

### GetTransDate

`func (o *ResponseSpecialRemittanceSlip) GetTransDate() string`

GetTransDate returns the TransDate field if non-nil, zero value otherwise.

### GetTransDateOk

`func (o *ResponseSpecialRemittanceSlip) GetTransDateOk() (*string, bool)`

GetTransDateOk returns a tuple with the TransDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransDate

`func (o *ResponseSpecialRemittanceSlip) SetTransDate(v string)`

SetTransDate sets TransDate field to given value.

### HasTransDate

`func (o *ResponseSpecialRemittanceSlip) HasTransDate() bool`

HasTransDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# HandlerCreateIPOSalesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSingleHand** | Pointer to **bool** |  | [optional] 
**OfficeId** | **int32** |  | 
**SaleAmount** | **float32** |  | 
**SaleCommission** | **float32** |  | 
**SaleDetails** | **[]int32** |  | 
**SoldTo** | **string** |  | 
**UserId** | **int32** |  | 

## Methods

### NewHandlerCreateIPOSalesRequest

`func NewHandlerCreateIPOSalesRequest(officeId int32, saleAmount float32, saleCommission float32, saleDetails []int32, soldTo string, userId int32, ) *HandlerCreateIPOSalesRequest`

NewHandlerCreateIPOSalesRequest instantiates a new HandlerCreateIPOSalesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerCreateIPOSalesRequestWithDefaults

`func NewHandlerCreateIPOSalesRequestWithDefaults() *HandlerCreateIPOSalesRequest`

NewHandlerCreateIPOSalesRequestWithDefaults instantiates a new HandlerCreateIPOSalesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSingleHand

`func (o *HandlerCreateIPOSalesRequest) GetIsSingleHand() bool`

GetIsSingleHand returns the IsSingleHand field if non-nil, zero value otherwise.

### GetIsSingleHandOk

`func (o *HandlerCreateIPOSalesRequest) GetIsSingleHandOk() (*bool, bool)`

GetIsSingleHandOk returns a tuple with the IsSingleHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleHand

`func (o *HandlerCreateIPOSalesRequest) SetIsSingleHand(v bool)`

SetIsSingleHand sets IsSingleHand field to given value.

### HasIsSingleHand

`func (o *HandlerCreateIPOSalesRequest) HasIsSingleHand() bool`

HasIsSingleHand returns a boolean if a field has been set.

### GetOfficeId

`func (o *HandlerCreateIPOSalesRequest) GetOfficeId() int32`

GetOfficeId returns the OfficeId field if non-nil, zero value otherwise.

### GetOfficeIdOk

`func (o *HandlerCreateIPOSalesRequest) GetOfficeIdOk() (*int32, bool)`

GetOfficeIdOk returns a tuple with the OfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeId

`func (o *HandlerCreateIPOSalesRequest) SetOfficeId(v int32)`

SetOfficeId sets OfficeId field to given value.


### GetSaleAmount

`func (o *HandlerCreateIPOSalesRequest) GetSaleAmount() float32`

GetSaleAmount returns the SaleAmount field if non-nil, zero value otherwise.

### GetSaleAmountOk

`func (o *HandlerCreateIPOSalesRequest) GetSaleAmountOk() (*float32, bool)`

GetSaleAmountOk returns a tuple with the SaleAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleAmount

`func (o *HandlerCreateIPOSalesRequest) SetSaleAmount(v float32)`

SetSaleAmount sets SaleAmount field to given value.


### GetSaleCommission

`func (o *HandlerCreateIPOSalesRequest) GetSaleCommission() float32`

GetSaleCommission returns the SaleCommission field if non-nil, zero value otherwise.

### GetSaleCommissionOk

`func (o *HandlerCreateIPOSalesRequest) GetSaleCommissionOk() (*float32, bool)`

GetSaleCommissionOk returns a tuple with the SaleCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleCommission

`func (o *HandlerCreateIPOSalesRequest) SetSaleCommission(v float32)`

SetSaleCommission sets SaleCommission field to given value.


### GetSaleDetails

`func (o *HandlerCreateIPOSalesRequest) GetSaleDetails() []int32`

GetSaleDetails returns the SaleDetails field if non-nil, zero value otherwise.

### GetSaleDetailsOk

`func (o *HandlerCreateIPOSalesRequest) GetSaleDetailsOk() (*[]int32, bool)`

GetSaleDetailsOk returns a tuple with the SaleDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleDetails

`func (o *HandlerCreateIPOSalesRequest) SetSaleDetails(v []int32)`

SetSaleDetails sets SaleDetails field to given value.


### GetSoldTo

`func (o *HandlerCreateIPOSalesRequest) GetSoldTo() string`

GetSoldTo returns the SoldTo field if non-nil, zero value otherwise.

### GetSoldToOk

`func (o *HandlerCreateIPOSalesRequest) GetSoldToOk() (*string, bool)`

GetSoldToOk returns a tuple with the SoldTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldTo

`func (o *HandlerCreateIPOSalesRequest) SetSoldTo(v string)`

SetSoldTo sets SoldTo field to given value.


### GetUserId

`func (o *HandlerCreateIPOSalesRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerCreateIPOSalesRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerCreateIPOSalesRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



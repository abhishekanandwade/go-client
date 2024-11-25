# HandlerTemporalcombinedCashbagcloseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcBagId** | Pointer to **string** |  | [optional] 
**AcBagWeight** | Pointer to **float32** |  | [optional] 
**ArticleCount** | Pointer to **int32** |  | [optional] 
**ArticleDetail** | Pointer to [**[]HandlerArticleDetail**](HandlerArticleDetail.md) |  | [optional] 
**ArticleWeight** | Pointer to **float32** |  | [optional] 
**BagCount** | Pointer to **int32** |  | [optional] 
**BagNumber** | Pointer to **string** |  | [optional] 
**BagType** | Pointer to **string** |  | [optional] 
**BagWeight** | Pointer to **float32** |  | [optional] 
**BoBagContent** | Pointer to [**[]HandlerBoBagContent**](HandlerBoBagContent.md) |  | [optional] 
**CashBagWeight** | Pointer to **float32** |  | [optional] 
**CashDenomination** | Pointer to [**[]HandlerCashDenomination**](HandlerCashDenomination.md) |  | [optional] 
**CashTotalAmount** | Pointer to **float32** |  | [optional] 
**CashbagId** | Pointer to **string** |  | [optional] 
**ChequeDetails** | Pointer to [**[]HandlerChequeDetail**](HandlerChequeDetail.md) |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**DeliveryType** | Pointer to **string** |  | [optional] 
**DocumentList** | Pointer to **string** |  | [optional] 
**EmoList** | Pointer to [**[]HandlerEmoDetail**](HandlerEmoDetail.md) |  | [optional] 
**EmoTotalAmount** | Pointer to **float32** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**FromOfficeId** | Pointer to **int32** |  | [optional] 
**InsuredFlag** | Pointer to **bool** |  | [optional] 
**IpoDetail** | Pointer to [**[]HandlerIPODetail**](HandlerIPODetail.md) |  | [optional] 
**ReceiverOfficeId** | Pointer to **int32** |  | [optional] 
**Remark** | Pointer to **string** |  | [optional] 
**SenderOfficeId** | Pointer to **int32** |  | [optional] 
**SetDate** | Pointer to **string** |  | [optional] 
**SetNumber** | Pointer to **string** |  | [optional] 
**StampDetail** | Pointer to [**[]HandlerStampDetail**](HandlerStampDetail.md) |  | [optional] 
**StationeryDetail** | Pointer to [**[]HandlerStationeryDetail**](HandlerStationeryDetail.md) |  | [optional] 
**ToOfficeId** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 

## Methods

### NewHandlerTemporalcombinedCashbagcloseRequest

`func NewHandlerTemporalcombinedCashbagcloseRequest() *HandlerTemporalcombinedCashbagcloseRequest`

NewHandlerTemporalcombinedCashbagcloseRequest instantiates a new HandlerTemporalcombinedCashbagcloseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandlerTemporalcombinedCashbagcloseRequestWithDefaults

`func NewHandlerTemporalcombinedCashbagcloseRequestWithDefaults() *HandlerTemporalcombinedCashbagcloseRequest`

NewHandlerTemporalcombinedCashbagcloseRequestWithDefaults instantiates a new HandlerTemporalcombinedCashbagcloseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcBagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetAcBagId() string`

GetAcBagId returns the AcBagId field if non-nil, zero value otherwise.

### GetAcBagIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetAcBagIdOk() (*string, bool)`

GetAcBagIdOk returns a tuple with the AcBagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcBagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetAcBagId(v string)`

SetAcBagId sets AcBagId field to given value.

### HasAcBagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasAcBagId() bool`

HasAcBagId returns a boolean if a field has been set.

### GetAcBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetAcBagWeight() float32`

GetAcBagWeight returns the AcBagWeight field if non-nil, zero value otherwise.

### GetAcBagWeightOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetAcBagWeightOk() (*float32, bool)`

GetAcBagWeightOk returns a tuple with the AcBagWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetAcBagWeight(v float32)`

SetAcBagWeight sets AcBagWeight field to given value.

### HasAcBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasAcBagWeight() bool`

HasAcBagWeight returns a boolean if a field has been set.

### GetArticleCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleCount() int32`

GetArticleCount returns the ArticleCount field if non-nil, zero value otherwise.

### GetArticleCountOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleCountOk() (*int32, bool)`

GetArticleCountOk returns a tuple with the ArticleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetArticleCount(v int32)`

SetArticleCount sets ArticleCount field to given value.

### HasArticleCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasArticleCount() bool`

HasArticleCount returns a boolean if a field has been set.

### GetArticleDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleDetail() []HandlerArticleDetail`

GetArticleDetail returns the ArticleDetail field if non-nil, zero value otherwise.

### GetArticleDetailOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleDetailOk() (*[]HandlerArticleDetail, bool)`

GetArticleDetailOk returns a tuple with the ArticleDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetArticleDetail(v []HandlerArticleDetail)`

SetArticleDetail sets ArticleDetail field to given value.

### HasArticleDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasArticleDetail() bool`

HasArticleDetail returns a boolean if a field has been set.

### GetArticleWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleWeight() float32`

GetArticleWeight returns the ArticleWeight field if non-nil, zero value otherwise.

### GetArticleWeightOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetArticleWeightOk() (*float32, bool)`

GetArticleWeightOk returns a tuple with the ArticleWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArticleWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetArticleWeight(v float32)`

SetArticleWeight sets ArticleWeight field to given value.

### HasArticleWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasArticleWeight() bool`

HasArticleWeight returns a boolean if a field has been set.

### GetBagCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagCount() int32`

GetBagCount returns the BagCount field if non-nil, zero value otherwise.

### GetBagCountOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagCountOk() (*int32, bool)`

GetBagCountOk returns a tuple with the BagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetBagCount(v int32)`

SetBagCount sets BagCount field to given value.

### HasBagCount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasBagCount() bool`

HasBagCount returns a boolean if a field has been set.

### GetBagNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagNumber() string`

GetBagNumber returns the BagNumber field if non-nil, zero value otherwise.

### GetBagNumberOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagNumberOk() (*string, bool)`

GetBagNumberOk returns a tuple with the BagNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetBagNumber(v string)`

SetBagNumber sets BagNumber field to given value.

### HasBagNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasBagNumber() bool`

HasBagNumber returns a boolean if a field has been set.

### GetBagType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagType() string`

GetBagType returns the BagType field if non-nil, zero value otherwise.

### GetBagTypeOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagTypeOk() (*string, bool)`

GetBagTypeOk returns a tuple with the BagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetBagType(v string)`

SetBagType sets BagType field to given value.

### HasBagType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasBagType() bool`

HasBagType returns a boolean if a field has been set.

### GetBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagWeight() float32`

GetBagWeight returns the BagWeight field if non-nil, zero value otherwise.

### GetBagWeightOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBagWeightOk() (*float32, bool)`

GetBagWeightOk returns a tuple with the BagWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetBagWeight(v float32)`

SetBagWeight sets BagWeight field to given value.

### HasBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasBagWeight() bool`

HasBagWeight returns a boolean if a field has been set.

### GetBoBagContent

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBoBagContent() []HandlerBoBagContent`

GetBoBagContent returns the BoBagContent field if non-nil, zero value otherwise.

### GetBoBagContentOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetBoBagContentOk() (*[]HandlerBoBagContent, bool)`

GetBoBagContentOk returns a tuple with the BoBagContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoBagContent

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetBoBagContent(v []HandlerBoBagContent)`

SetBoBagContent sets BoBagContent field to given value.

### HasBoBagContent

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasBoBagContent() bool`

HasBoBagContent returns a boolean if a field has been set.

### GetCashBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashBagWeight() float32`

GetCashBagWeight returns the CashBagWeight field if non-nil, zero value otherwise.

### GetCashBagWeightOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashBagWeightOk() (*float32, bool)`

GetCashBagWeightOk returns a tuple with the CashBagWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCashBagWeight(v float32)`

SetCashBagWeight sets CashBagWeight field to given value.

### HasCashBagWeight

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCashBagWeight() bool`

HasCashBagWeight returns a boolean if a field has been set.

### GetCashDenomination

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashDenomination() []HandlerCashDenomination`

GetCashDenomination returns the CashDenomination field if non-nil, zero value otherwise.

### GetCashDenominationOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashDenominationOk() (*[]HandlerCashDenomination, bool)`

GetCashDenominationOk returns a tuple with the CashDenomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashDenomination

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCashDenomination(v []HandlerCashDenomination)`

SetCashDenomination sets CashDenomination field to given value.

### HasCashDenomination

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCashDenomination() bool`

HasCashDenomination returns a boolean if a field has been set.

### GetCashTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashTotalAmount() float32`

GetCashTotalAmount returns the CashTotalAmount field if non-nil, zero value otherwise.

### GetCashTotalAmountOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashTotalAmountOk() (*float32, bool)`

GetCashTotalAmountOk returns a tuple with the CashTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCashTotalAmount(v float32)`

SetCashTotalAmount sets CashTotalAmount field to given value.

### HasCashTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCashTotalAmount() bool`

HasCashTotalAmount returns a boolean if a field has been set.

### GetCashbagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashbagId() string`

GetCashbagId returns the CashbagId field if non-nil, zero value otherwise.

### GetCashbagIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCashbagIdOk() (*string, bool)`

GetCashbagIdOk returns a tuple with the CashbagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCashbagId(v string)`

SetCashbagId sets CashbagId field to given value.

### HasCashbagId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCashbagId() bool`

HasCashbagId returns a boolean if a field has been set.

### GetChequeDetails

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetChequeDetails() []HandlerChequeDetail`

GetChequeDetails returns the ChequeDetails field if non-nil, zero value otherwise.

### GetChequeDetailsOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetChequeDetailsOk() (*[]HandlerChequeDetail, bool)`

GetChequeDetailsOk returns a tuple with the ChequeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChequeDetails

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetChequeDetails(v []HandlerChequeDetail)`

SetChequeDetails sets ChequeDetails field to given value.

### HasChequeDetails

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasChequeDetails() bool`

HasChequeDetails returns a boolean if a field has been set.

### GetCreatedBy

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDeliveryType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetDeliveryType() string`

GetDeliveryType returns the DeliveryType field if non-nil, zero value otherwise.

### GetDeliveryTypeOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetDeliveryTypeOk() (*string, bool)`

GetDeliveryTypeOk returns a tuple with the DeliveryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetDeliveryType(v string)`

SetDeliveryType sets DeliveryType field to given value.

### HasDeliveryType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasDeliveryType() bool`

HasDeliveryType returns a boolean if a field has been set.

### GetDocumentList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetDocumentList() string`

GetDocumentList returns the DocumentList field if non-nil, zero value otherwise.

### GetDocumentListOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetDocumentListOk() (*string, bool)`

GetDocumentListOk returns a tuple with the DocumentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetDocumentList(v string)`

SetDocumentList sets DocumentList field to given value.

### HasDocumentList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasDocumentList() bool`

HasDocumentList returns a boolean if a field has been set.

### GetEmoList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEmoList() []HandlerEmoDetail`

GetEmoList returns the EmoList field if non-nil, zero value otherwise.

### GetEmoListOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEmoListOk() (*[]HandlerEmoDetail, bool)`

GetEmoListOk returns a tuple with the EmoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetEmoList(v []HandlerEmoDetail)`

SetEmoList sets EmoList field to given value.

### HasEmoList

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasEmoList() bool`

HasEmoList returns a boolean if a field has been set.

### GetEmoTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEmoTotalAmount() float32`

GetEmoTotalAmount returns the EmoTotalAmount field if non-nil, zero value otherwise.

### GetEmoTotalAmountOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEmoTotalAmountOk() (*float32, bool)`

GetEmoTotalAmountOk returns a tuple with the EmoTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetEmoTotalAmount(v float32)`

SetEmoTotalAmount sets EmoTotalAmount field to given value.

### HasEmoTotalAmount

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasEmoTotalAmount() bool`

HasEmoTotalAmount returns a boolean if a field has been set.

### GetEventType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFromOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetFromOfficeId() int32`

GetFromOfficeId returns the FromOfficeId field if non-nil, zero value otherwise.

### GetFromOfficeIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetFromOfficeIdOk() (*int32, bool)`

GetFromOfficeIdOk returns a tuple with the FromOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetFromOfficeId(v int32)`

SetFromOfficeId sets FromOfficeId field to given value.

### HasFromOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasFromOfficeId() bool`

HasFromOfficeId returns a boolean if a field has been set.

### GetInsuredFlag

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetInsuredFlag() bool`

GetInsuredFlag returns the InsuredFlag field if non-nil, zero value otherwise.

### GetInsuredFlagOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetInsuredFlagOk() (*bool, bool)`

GetInsuredFlagOk returns a tuple with the InsuredFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsuredFlag

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetInsuredFlag(v bool)`

SetInsuredFlag sets InsuredFlag field to given value.

### HasInsuredFlag

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasInsuredFlag() bool`

HasInsuredFlag returns a boolean if a field has been set.

### GetIpoDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetIpoDetail() []HandlerIPODetail`

GetIpoDetail returns the IpoDetail field if non-nil, zero value otherwise.

### GetIpoDetailOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetIpoDetailOk() (*[]HandlerIPODetail, bool)`

GetIpoDetailOk returns a tuple with the IpoDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpoDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetIpoDetail(v []HandlerIPODetail)`

SetIpoDetail sets IpoDetail field to given value.

### HasIpoDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasIpoDetail() bool`

HasIpoDetail returns a boolean if a field has been set.

### GetReceiverOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetReceiverOfficeId() int32`

GetReceiverOfficeId returns the ReceiverOfficeId field if non-nil, zero value otherwise.

### GetReceiverOfficeIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetReceiverOfficeIdOk() (*int32, bool)`

GetReceiverOfficeIdOk returns a tuple with the ReceiverOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetReceiverOfficeId(v int32)`

SetReceiverOfficeId sets ReceiverOfficeId field to given value.

### HasReceiverOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasReceiverOfficeId() bool`

HasReceiverOfficeId returns a boolean if a field has been set.

### GetRemark

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetRemark() string`

GetRemark returns the Remark field if non-nil, zero value otherwise.

### GetRemarkOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetRemarkOk() (*string, bool)`

GetRemarkOk returns a tuple with the Remark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemark

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetRemark(v string)`

SetRemark sets Remark field to given value.

### HasRemark

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasRemark() bool`

HasRemark returns a boolean if a field has been set.

### GetSenderOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSenderOfficeId() int32`

GetSenderOfficeId returns the SenderOfficeId field if non-nil, zero value otherwise.

### GetSenderOfficeIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSenderOfficeIdOk() (*int32, bool)`

GetSenderOfficeIdOk returns a tuple with the SenderOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetSenderOfficeId(v int32)`

SetSenderOfficeId sets SenderOfficeId field to given value.

### HasSenderOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasSenderOfficeId() bool`

HasSenderOfficeId returns a boolean if a field has been set.

### GetSetDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSetDate() string`

GetSetDate returns the SetDate field if non-nil, zero value otherwise.

### GetSetDateOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSetDateOk() (*string, bool)`

GetSetDateOk returns a tuple with the SetDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetSetDate(v string)`

SetSetDate sets SetDate field to given value.

### HasSetDate

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasSetDate() bool`

HasSetDate returns a boolean if a field has been set.

### GetSetNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSetNumber() string`

GetSetNumber returns the SetNumber field if non-nil, zero value otherwise.

### GetSetNumberOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetSetNumberOk() (*string, bool)`

GetSetNumberOk returns a tuple with the SetNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetSetNumber(v string)`

SetSetNumber sets SetNumber field to given value.

### HasSetNumber

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasSetNumber() bool`

HasSetNumber returns a boolean if a field has been set.

### GetStampDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetStampDetail() []HandlerStampDetail`

GetStampDetail returns the StampDetail field if non-nil, zero value otherwise.

### GetStampDetailOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetStampDetailOk() (*[]HandlerStampDetail, bool)`

GetStampDetailOk returns a tuple with the StampDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetStampDetail(v []HandlerStampDetail)`

SetStampDetail sets StampDetail field to given value.

### HasStampDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasStampDetail() bool`

HasStampDetail returns a boolean if a field has been set.

### GetStationeryDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetStationeryDetail() []HandlerStationeryDetail`

GetStationeryDetail returns the StationeryDetail field if non-nil, zero value otherwise.

### GetStationeryDetailOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetStationeryDetailOk() (*[]HandlerStationeryDetail, bool)`

GetStationeryDetailOk returns a tuple with the StationeryDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationeryDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetStationeryDetail(v []HandlerStationeryDetail)`

SetStationeryDetail sets StationeryDetail field to given value.

### HasStationeryDetail

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasStationeryDetail() bool`

HasStationeryDetail returns a boolean if a field has been set.

### GetToOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetToOfficeId() int32`

GetToOfficeId returns the ToOfficeId field if non-nil, zero value otherwise.

### GetToOfficeIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetToOfficeIdOk() (*int32, bool)`

GetToOfficeIdOk returns a tuple with the ToOfficeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetToOfficeId(v int32)`

SetToOfficeId sets ToOfficeId field to given value.

### HasToOfficeId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasToOfficeId() bool`

HasToOfficeId returns a boolean if a field has been set.

### GetUserId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *HandlerTemporalcombinedCashbagcloseRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *HandlerTemporalcombinedCashbagcloseRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



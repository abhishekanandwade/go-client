# \ConfigureOfficesAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficeTransactionsTransactionIdBankCreditLimitsPut**](ConfigureOfficesAPI.md#OfficeTransactionsTransactionIdBankCreditLimitsPut) | **Put** /office-transactions/{transaction-id}/bank-credit-limits | Update Office
[**OfficeTransactionsTransactionIdDelete**](ConfigureOfficesAPI.md#OfficeTransactionsTransactionIdDelete) | **Delete** /office-transactions/{transaction-id} | Update Office
[**OfficesGet**](ConfigureOfficesAPI.md#OfficesGet) | **Get** /offices | Get All Offices List
[**OfficesOfficeIdBankCreditLmitsGet**](ConfigureOfficesAPI.md#OfficesOfficeIdBankCreditLmitsGet) | **Get** /offices/{office-id}/bank-credit-lmits | Get Offices for bank credit limit configuration by Division ID
[**OfficesOfficeIdConfigsGet**](ConfigureOfficesAPI.md#OfficesOfficeIdConfigsGet) | **Get** /offices/{office-id}/configs | List Offices for Config by Division ID
[**OfficesOfficeIdGet**](ConfigureOfficesAPI.md#OfficesOfficeIdGet) | **Get** /offices/{office-id} | Get Specific Office Data
[**OfficesOfficeIdLinkedConfiguredGet**](ConfigureOfficesAPI.md#OfficesOfficeIdLinkedConfiguredGet) | **Get** /offices/{office-id}/linked-configured | Get Offices linked to Cash Office
[**OfficesPost**](ConfigureOfficesAPI.md#OfficesPost) | **Post** /offices | Create Office configuration



## OfficeTransactionsTransactionIdBankCreditLimitsPut

> ResponseUpdateBankCreditLimitAPIResponse OfficeTransactionsTransactionIdBankCreditLimitsPut(ctx, transactionId).Body(body).Execute()

Update Office



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "transactionId_example" // string | Transaction ID
	body := *openapiclient.NewHandlerUpdateOfficeTransactionsBankCreditLimitsRequest(float32(10000.0)) // HandlerUpdateOfficeTransactionsBankCreditLimitsRequest | Bank Credit Limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficeTransactionsTransactionIdBankCreditLimitsPut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficeTransactionsTransactionIdBankCreditLimitsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficeTransactionsTransactionIdBankCreditLimitsPut`: ResponseUpdateBankCreditLimitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficeTransactionsTransactionIdBankCreditLimitsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeTransactionsTransactionIdBankCreditLimitsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateOfficeTransactionsBankCreditLimitsRequest**](HandlerUpdateOfficeTransactionsBankCreditLimitsRequest.md) | Bank Credit Limit | 

### Return type

[**ResponseUpdateBankCreditLimitAPIResponse**](ResponseUpdateBankCreditLimitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficeTransactionsTransactionIdDelete

> ResponseOfficeAPIResponse OfficeTransactionsTransactionIdDelete(ctx, transactionId).UpdatedBy(updatedBy).ValidTo(validTo).Execute()

Update Office



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	transactionId := "transactionId_example" // string | Transaction ID
	updatedBy := "updatedBy_example" // string | User who updated the office
	validTo := "validTo_example" // string | Date when the office details are valid till

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficeTransactionsTransactionIdDelete(context.Background(), transactionId).UpdatedBy(updatedBy).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficeTransactionsTransactionIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficeTransactionsTransactionIdDelete`: ResponseOfficeAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficeTransactionsTransactionIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeTransactionsTransactionIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedBy** | **string** | User who updated the office | 
 **validTo** | **string** | Date when the office details are valid till | 

### Return type

[**ResponseOfficeAPIResponse**](ResponseOfficeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesGet

> ResponseListOfficesResponseAPIResponse OfficesGet(ctx).Skip(skip).Limit(limit).Execute()

Get All Offices List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	skip := int32(56) // int32 | Number of records to skip for pagination (optional) (default to 0)
	limit := int32(56) // int32 | Number of records to limit for pagination (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesGet`: ResponseListOfficesResponseAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfficesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip for pagination | [default to 0]
 **limit** | **int32** | Number of records to limit for pagination | [default to 10]

### Return type

[**ResponseListOfficesResponseAPIResponse**](ResponseListOfficesResponseAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdBankCreditLmitsGet

> ResponseListBankCreditLimitAPIResponse OfficesOfficeIdBankCreditLmitsGet(ctx, officeId).Execute()

Get Offices for bank credit limit configuration by Division ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	officeId := int32(56) // int32 | Office ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdBankCreditLmitsGet(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesOfficeIdBankCreditLmitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdBankCreditLmitsGet`: ResponseListBankCreditLimitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesOfficeIdBankCreditLmitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdBankCreditLmitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseListBankCreditLimitAPIResponse**](ResponseListBankCreditLimitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdConfigsGet

> ResponseListOfficesConfigAPIResponse OfficesOfficeIdConfigsGet(ctx, officeId).Type_(type_).Execute()

List Offices for Config by Division ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	officeId := int32(56) // int32 | Get config details of specific office
	type_ := "type__example" // string | Get linked or configured offices

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdConfigsGet(context.Background(), officeId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesOfficeIdConfigsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdConfigsGet`: ResponseListOfficesConfigAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesOfficeIdConfigsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Get config details of specific office | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdConfigsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Get linked or configured offices | 

### Return type

[**ResponseListOfficesConfigAPIResponse**](ResponseListOfficesConfigAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdGet

> ResponseOfficeAPIResponse OfficesOfficeIdGet(ctx, officeId).Execute()

Get Specific Office Data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	officeId := int32(56) // int32 | Get config details of specific office

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdGet(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesOfficeIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdGet`: ResponseOfficeAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesOfficeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Get config details of specific office | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseOfficeAPIResponse**](ResponseOfficeAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdLinkedConfiguredGet

> ResponseListOfficesResponseAPIResponse OfficesOfficeIdLinkedConfiguredGet(ctx, officeId).Type_(type_).Execute()

Get Offices linked to Cash Office



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	officeId := int32(56) // int32 | Get config details of specific office
	type_ := "type__example" // string | Get linked or configured offices

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesOfficeIdLinkedConfiguredGet(context.Background(), officeId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesOfficeIdLinkedConfiguredGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdLinkedConfiguredGet`: ResponseListOfficesResponseAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesOfficeIdLinkedConfiguredGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Get config details of specific office | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdLinkedConfiguredGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Get linked or configured offices | 

### Return type

[**ResponseListOfficesResponseAPIResponse**](ResponseListOfficesResponseAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesPost

> ResponseCreateOfficesResponseApiResponse OfficesPost(ctx).Body(body).Execute()

Create Office configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewHandlerCreateOfficesRequest("10130000", float32(100000.0), float32(5000.0), int32(9000003), "Mysuru HO", "2023-12-08T00:00:00Z") // HandlerCreateOfficesRequest | Information about adding office details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureOfficesAPI.OfficesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureOfficesAPI.OfficesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesPost`: ResponseCreateOfficesResponseApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureOfficesAPI.OfficesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOfficesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateOfficesRequest**](HandlerCreateOfficesRequest.md) | Information about adding office details | 

### Return type

[**ResponseCreateOfficesResponseApiResponse**](ResponseCreateOfficesResponseApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


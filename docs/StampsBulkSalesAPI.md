# \StampsBulkSalesAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampBulkSalesGet**](StampsBulkSalesAPI.md#StampBulkSalesGet) | **Get** /stamp-bulk-sales | Get Stamps Bulk Sales List
[**StampBulkSalesPost**](StampsBulkSalesAPI.md#StampBulkSalesPost) | **Post** /stamp-bulk-sales | Create Stamp Bulk sales
[**StampBulkSalesTransactionIdApprovePut**](StampsBulkSalesAPI.md#StampBulkSalesTransactionIdApprovePut) | **Put** /stamp-bulk-sales/{transaction-id}/approve | Approve Stamp Bulk sales
[**StampBulkSalesTransactionIdGet**](StampsBulkSalesAPI.md#StampBulkSalesTransactionIdGet) | **Get** /stamp-bulk-sales/{transaction-id} | Get Stamp Bulk Sales by ID



## StampBulkSalesGet

> ResponseListStampBulkSalesApiResponse StampBulkSalesGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()

Get Stamps Bulk Sales List



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
	officeId := "officeId_example" // string | Office ID
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	isPdg := true // bool | Is PDG (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsBulkSalesAPI.StampBulkSalesGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsBulkSalesAPI.StampBulkSalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampBulkSalesGet`: ResponseListStampBulkSalesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsBulkSalesAPI.StampBulkSalesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampBulkSalesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **isPdg** | **bool** | Is PDG | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListStampBulkSalesApiResponse**](ResponseListStampBulkSalesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampBulkSalesPost

> ResponseCreateStampBulkSalesApiResponse StampBulkSalesPost(ctx).Body(body).Execute()

Create Stamp Bulk sales



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
	body := *openapiclient.NewHandlerCreateStampBulkSalesRequest(int32(9000003), "Cash/Cheque", float32(300.65), map[string]int32{"key": int32(123)}, "XXX", int32(10130000)) // HandlerCreateStampBulkSalesRequest | Create stamp bulk sales

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsBulkSalesAPI.StampBulkSalesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsBulkSalesAPI.StampBulkSalesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampBulkSalesPost`: ResponseCreateStampBulkSalesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsBulkSalesAPI.StampBulkSalesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampBulkSalesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateStampBulkSalesRequest**](HandlerCreateStampBulkSalesRequest.md) | Create stamp bulk sales | 

### Return type

[**ResponseCreateStampBulkSalesApiResponse**](ResponseCreateStampBulkSalesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampBulkSalesTransactionIdApprovePut

> ResponseStampBulkSalesApiResponse StampBulkSalesTransactionIdApprovePut(ctx, transactionId).Body(body).Execute()

Approve Stamp Bulk sales



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
	body := *openapiclient.NewHandlerUpdateStampBulkSalesRequest(int32(10121212)) // HandlerUpdateStampBulkSalesRequest | Approve stamp bulk sales (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsBulkSalesAPI.StampBulkSalesTransactionIdApprovePut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsBulkSalesAPI.StampBulkSalesTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampBulkSalesTransactionIdApprovePut`: ResponseStampBulkSalesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsBulkSalesAPI.StampBulkSalesTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampBulkSalesTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateStampBulkSalesRequest**](HandlerUpdateStampBulkSalesRequest.md) | Approve stamp bulk sales | 

### Return type

[**ResponseStampBulkSalesApiResponse**](ResponseStampBulkSalesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampBulkSalesTransactionIdGet

> ResponseStampBulkSalesApiResponse StampBulkSalesTransactionIdGet(ctx, transactionId).Execute()

Get Stamp Bulk Sales by ID



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsBulkSalesAPI.StampBulkSalesTransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsBulkSalesAPI.StampBulkSalesTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampBulkSalesTransactionIdGet`: ResponseStampBulkSalesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsBulkSalesAPI.StampBulkSalesTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampBulkSalesTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseStampBulkSalesApiResponse**](ResponseStampBulkSalesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


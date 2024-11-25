# \IPOsApproveSaleAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposSalesTransactionIdApprovePut**](IPOsApproveSaleAPI.md#IposSalesTransactionIdApprovePut) | **Put** /ipos-sales/{transaction-id}/approve | Approve IPO sale



## IposSalesTransactionIdApprovePut

> ResponseUpdateIPOSaleApiResponse IposSalesTransactionIdApprovePut(ctx, transactionId).UpdateIPOSaleRequest(updateIPOSaleRequest).Execute()

Approve IPO sale



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
	updateIPOSaleRequest := *openapiclient.NewHandlerUpdateIPOSaleRequest(int32(10121212), "Approved") // HandlerUpdateIPOSaleRequest | Information about approving IPO sales

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOsApproveSaleAPI.IposSalesTransactionIdApprovePut(context.Background(), transactionId).UpdateIPOSaleRequest(updateIPOSaleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOsApproveSaleAPI.IposSalesTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposSalesTransactionIdApprovePut`: ResponseUpdateIPOSaleApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOsApproveSaleAPI.IposSalesTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposSalesTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIPOSaleRequest** | [**HandlerUpdateIPOSaleRequest**](HandlerUpdateIPOSaleRequest.md) | Information about approving IPO sales | 

### Return type

[**ResponseUpdateIPOSaleApiResponse**](ResponseUpdateIPOSaleApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


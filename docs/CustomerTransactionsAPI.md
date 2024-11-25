# \CustomerTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerTransactionsGet**](CustomerTransactionsAPI.md#CustomerTransactionsGet) | **Get** /customer-transactions | Get Customer Transactions
[**CustomerTransactionsPost**](CustomerTransactionsAPI.md#CustomerTransactionsPost) | **Post** /customer-transactions | Create Customer Transactions
[**CustomerTransactionsTransactionIdApprovePut**](CustomerTransactionsAPI.md#CustomerTransactionsTransactionIdApprovePut) | **Put** /customer-transactions/{transaction-id}/approve | Approve Customer Transactions



## CustomerTransactionsGet

> ResponseListCustomerTransactionAPIResponse CustomerTransactionsGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()

Get Customer Transactions



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
	officeId := int32(56) // int32 | Office ID (example: 90000003)
	fromDate := "fromDate_example" // string | From Date (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | To Date (example: 2024-01-31) (optional)
	isPdg := true // bool | Is PDG (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTransactionsAPI.CustomerTransactionsGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTransactionsAPI.CustomerTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerTransactionsGet`: ResponseListCustomerTransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerTransactionsAPI.CustomerTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **int32** | Office ID (example: 90000003) | 
 **fromDate** | **string** | From Date (example: 2024-01-31) | 
 **toDate** | **string** | To Date (example: 2024-01-31) | 
 **isPdg** | **bool** | Is PDG | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListCustomerTransactionAPIResponse**](ResponseListCustomerTransactionAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerTransactionsPost

> ResponseCreateCustomerTransactionAPIResponse CustomerTransactionsPost(ctx).Body(body).Execute()

Create Customer Transactions



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
	body := *openapiclient.NewHandlerCreateCustomerTransactionsRequest(float32(1000.5), "9999-12-31T00:00:00Z", "BILL123456", "Monthly bill", "12345", "John Doe", "Regular", int32(9000015), "Payment for December", float32(500.75), "Cash", "10998889") // HandlerCreateCustomerTransactionsRequest | Create customer transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTransactionsAPI.CustomerTransactionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTransactionsAPI.CustomerTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerTransactionsPost`: ResponseCreateCustomerTransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerTransactionsAPI.CustomerTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomerTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateCustomerTransactionsRequest**](HandlerCreateCustomerTransactionsRequest.md) | Create customer transaction | 

### Return type

[**ResponseCreateCustomerTransactionAPIResponse**](ResponseCreateCustomerTransactionAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomerTransactionsTransactionIdApprovePut

> ResponseCustomerTransactionAPIResponse CustomerTransactionsTransactionIdApprovePut(ctx, transactionId).Body(body).Execute()

Approve Customer Transactions



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
	transactionId := "transactionId_example" // string | Information about approval of customer transactions
	body := *openapiclient.NewHandlerApproveCustomerTransactionsApproveRequest("10121212") // HandlerApproveCustomerTransactionsApproveRequest | Approver ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomerTransactionsAPI.CustomerTransactionsTransactionIdApprovePut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomerTransactionsAPI.CustomerTransactionsTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomerTransactionsTransactionIdApprovePut`: ResponseCustomerTransactionAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `CustomerTransactionsAPI.CustomerTransactionsTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Information about approval of customer transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomerTransactionsTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerApproveCustomerTransactionsApproveRequest**](HandlerApproveCustomerTransactionsApproveRequest.md) | Approver ID | 

### Return type

[**ResponseCustomerTransactionAPIResponse**](ResponseCustomerTransactionAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


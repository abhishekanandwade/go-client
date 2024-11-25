# \StampsTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsTransactionsGet**](StampsTransactionsAPI.md#StampsTransactionsGet) | **Get** /stamps-transactions | List Stamp Transactions
[**StampsTransactionsPost**](StampsTransactionsAPI.md#StampsTransactionsPost) | **Post** /stamps-transactions | Create a New Stamp Transaction Request
[**StampsTransactionsTransactionIdChangeStatusPut**](StampsTransactionsAPI.md#StampsTransactionsTransactionIdChangeStatusPut) | **Put** /stamps-transactions/{transaction-id}/change-status | Process Stamp Transaction
[**StampsTransactionsTransactionIdGet**](StampsTransactionsAPI.md#StampsTransactionsTransactionIdGet) | **Get** /stamps-transactions/{transaction-id} | List Stamp Transaction by ID



## StampsTransactionsGet

> ResponseListStampTransactionsApiResonse StampsTransactionsGet(ctx).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).TxnStatus(txnStatus).Skip(skip).Limit(limit).Execute()

List Stamp Transactions



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
	type_ := "type__example" // string | Type
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	txnStatus := "txnStatus_example" // string | Transaction Status (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsTransactionsAPI.StampsTransactionsGet(context.Background()).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).TxnStatus(txnStatus).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsTransactionsAPI.StampsTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsTransactionsGet`: ResponseListStampTransactionsApiResonse
	fmt.Fprintf(os.Stdout, "Response from `StampsTransactionsAPI.StampsTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **type_** | **string** | Type | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **txnStatus** | **string** | Transaction Status | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListStampTransactionsApiResonse**](ResponseListStampTransactionsApiResonse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsTransactionsPost

> ResponseCreateStampsTransactionsApiReaponse StampsTransactionsPost(ctx).Body(body).Execute()

Create a New Stamp Transaction Request



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
	body := *openapiclient.NewHandlerCreateStampsTransactionsRequest(int32(9000003), float32(25000.55), map[string]int32{"key": int32(123)}, int32(9000003), int32(10130000), "Treasury/PoS/DARPAN", "Request/ Remittance") // HandlerCreateStampsTransactionsRequest | Information about the new stamp transaction request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsTransactionsAPI.StampsTransactionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsTransactionsAPI.StampsTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsTransactionsPost`: ResponseCreateStampsTransactionsApiReaponse
	fmt.Fprintf(os.Stdout, "Response from `StampsTransactionsAPI.StampsTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateStampsTransactionsRequest**](HandlerCreateStampsTransactionsRequest.md) | Information about the new stamp transaction request | 

### Return type

[**ResponseCreateStampsTransactionsApiReaponse**](ResponseCreateStampsTransactionsApiReaponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsTransactionsTransactionIdChangeStatusPut

> ResponseStampsTransactionsApiReaponse StampsTransactionsTransactionIdChangeStatusPut(ctx, transactionId).Type_(type_).Body(body).Execute()

Process Stamp Transaction



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
	type_ := "type__example" // string | Type
	body := *openapiclient.NewHandlerUpdateStampsTransactionsStatusRequest(int32(10130000)) // HandlerUpdateStampsTransactionsStatusRequest | Process stamp transaction issue request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsTransactionsAPI.StampsTransactionsTransactionIdChangeStatusPut(context.Background(), transactionId).Type_(type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsTransactionsAPI.StampsTransactionsTransactionIdChangeStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsTransactionsTransactionIdChangeStatusPut`: ResponseStampsTransactionsApiReaponse
	fmt.Fprintf(os.Stdout, "Response from `StampsTransactionsAPI.StampsTransactionsTransactionIdChangeStatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsTransactionsTransactionIdChangeStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type | 
 **body** | [**HandlerUpdateStampsTransactionsStatusRequest**](HandlerUpdateStampsTransactionsStatusRequest.md) | Process stamp transaction issue request | 

### Return type

[**ResponseStampsTransactionsApiReaponse**](ResponseStampsTransactionsApiReaponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsTransactionsTransactionIdGet

> ResponseFetchStampsTransactionsApiResponse StampsTransactionsTransactionIdGet(ctx, transactionId).Execute()

List Stamp Transaction by ID



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
	transactionId := "transactionId_example" // string | List of stamp transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsTransactionsAPI.StampsTransactionsTransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsTransactionsAPI.StampsTransactionsTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsTransactionsTransactionIdGet`: ResponseFetchStampsTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsTransactionsAPI.StampsTransactionsTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | List of stamp transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsTransactionsTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseFetchStampsTransactionsApiResponse**](ResponseFetchStampsTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


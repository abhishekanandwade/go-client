# \IPOTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposTransactionsGet**](IPOTransactionsAPI.md#IposTransactionsGet) | **Get** /ipos-transactions | IPO Transaction requests
[**IposTransactionsPost**](IPOTransactionsAPI.md#IposTransactionsPost) | **Post** /ipos-transactions | Create a New IPO Transaction Requests
[**IposTransactionsTransactionIdStatusPut**](IPOTransactionsAPI.md#IposTransactionsTransactionIdStatusPut) | **Put** /ipos-transactions/{transaction-id}/status | Process IPO  Transaction Issue Request



## IposTransactionsGet

> ResponseListIPOTransactionsApiResponse IposTransactionsGet(ctx).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).TxnStatus(txnStatus).Execute()

IPO Transaction requests



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
	officeId := "officeId_example" // string | Office ID of the user
	type_ := "type__example" // string | type
	fromDate := "fromDate_example" // string | From date of IPO transactions (optional)
	toDate := "toDate_example" // string | To date of IPO transactions (optional)
	txnStatus := "txnStatus_example" // string | Status of IPO transactions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOTransactionsAPI.IposTransactionsGet(context.Background()).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).TxnStatus(txnStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOTransactionsAPI.IposTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposTransactionsGet`: ResponseListIPOTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOTransactionsAPI.IposTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID of the user | 
 **type_** | **string** | type | 
 **fromDate** | **string** | From date of IPO transactions | 
 **toDate** | **string** | To date of IPO transactions | 
 **txnStatus** | **string** | Status of IPO transactions | 

### Return type

[**ResponseListIPOTransactionsApiResponse**](ResponseListIPOTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IposTransactionsPost

> ResponseCreateIPOsTransactionsApiResponse IposTransactionsPost(ctx).CreateIPOsTransactionsRequest(createIPOsTransactionsRequest).Execute()

Create a New IPO Transaction Requests



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
	createIPOsTransactionsRequest := *openapiclient.NewHandlerCreateIPOsTransactionsRequest(int32(9000003), float32(25000.55), int32(9000003), int32(10130000), "Treasury/PoS/DARPAN", "Request/ Remittance") // HandlerCreateIPOsTransactionsRequest | Information about the new IPO transaction request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOTransactionsAPI.IposTransactionsPost(context.Background()).CreateIPOsTransactionsRequest(createIPOsTransactionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOTransactionsAPI.IposTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposTransactionsPost`: ResponseCreateIPOsTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOTransactionsAPI.IposTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIPOsTransactionsRequest** | [**HandlerCreateIPOsTransactionsRequest**](HandlerCreateIPOsTransactionsRequest.md) | Information about the new IPO transaction request | 

### Return type

[**ResponseCreateIPOsTransactionsApiResponse**](ResponseCreateIPOsTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IposTransactionsTransactionIdStatusPut

> ResponseUpdateIPOsTransactionsApiResonse IposTransactionsTransactionIdStatusPut(ctx, transactionId).UpdateIPOsTransactionsStatusRequest(updateIPOsTransactionsStatusRequest).Execute()

Process IPO  Transaction Issue Request



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
	transactionId := "transactionId_example" // string | Process IPO txn at Dest Office
	updateIPOsTransactionsStatusRequest := *openapiclient.NewHandlerUpdateIPOsTransactionsStatusRequest(int32(10130000), openapiclient.handler.transactionType("approve-pdg-source")) // HandlerUpdateIPOsTransactionsStatusRequest | Process IPO txn at Dest Office

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOTransactionsAPI.IposTransactionsTransactionIdStatusPut(context.Background(), transactionId).UpdateIPOsTransactionsStatusRequest(updateIPOsTransactionsStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOTransactionsAPI.IposTransactionsTransactionIdStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposTransactionsTransactionIdStatusPut`: ResponseUpdateIPOsTransactionsApiResonse
	fmt.Fprintf(os.Stdout, "Response from `IPOTransactionsAPI.IposTransactionsTransactionIdStatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Process IPO txn at Dest Office | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposTransactionsTransactionIdStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIPOsTransactionsStatusRequest** | [**HandlerUpdateIPOsTransactionsStatusRequest**](HandlerUpdateIPOsTransactionsStatusRequest.md) | Process IPO txn at Dest Office | 

### Return type

[**ResponseUpdateIPOsTransactionsApiResonse**](ResponseUpdateIPOsTransactionsApiResonse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


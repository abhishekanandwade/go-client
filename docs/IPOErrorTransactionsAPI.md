# \IPOErrorTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposErrorTransactionsErrorIdPut**](IPOErrorTransactionsAPI.md#IposErrorTransactionsErrorIdPut) | **Put** /ipos-error-transactions/{error-id} | Update IPO error Transactions
[**IposErrorTransactionsGet**](IPOErrorTransactionsAPI.md#IposErrorTransactionsGet) | **Get** /ipos-error-transactions | Fetch IPO error Transactions



## IposErrorTransactionsErrorIdPut

> ResponseUpdateIPOsErrorTransactionsApiResponse IposErrorTransactionsErrorIdPut(ctx, errorId).Body(body).Execute()

Update IPO error Transactions



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
	errorId := "errorId_example" // string | Error ID
	body := *openapiclient.NewHandlerUpdateIPOsErrorTransactionsRequest(int32(9000001), "remarks of error", int32(10145824)) // HandlerUpdateIPOsErrorTransactionsRequest | Update IPO Error transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOErrorTransactionsAPI.IposErrorTransactionsErrorIdPut(context.Background(), errorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOErrorTransactionsAPI.IposErrorTransactionsErrorIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposErrorTransactionsErrorIdPut`: ResponseUpdateIPOsErrorTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOErrorTransactionsAPI.IposErrorTransactionsErrorIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**errorId** | **string** | Error ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposErrorTransactionsErrorIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateIPOsErrorTransactionsRequest**](HandlerUpdateIPOsErrorTransactionsRequest.md) | Update IPO Error transactions | 

### Return type

[**ResponseUpdateIPOsErrorTransactionsApiResponse**](ResponseUpdateIPOsErrorTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IposErrorTransactionsGet

> ResponseListIPOErrorTransactionsAPIResponse IposErrorTransactionsGet(ctx).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).Execute()

Fetch IPO error Transactions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOErrorTransactionsAPI.IposErrorTransactionsGet(context.Background()).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOErrorTransactionsAPI.IposErrorTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposErrorTransactionsGet`: ResponseListIPOErrorTransactionsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOErrorTransactionsAPI.IposErrorTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposErrorTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **type_** | **string** | Type | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 

### Return type

[**ResponseListIPOErrorTransactionsAPIResponse**](ResponseListIPOErrorTransactionsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


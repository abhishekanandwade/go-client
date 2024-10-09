# \CashErrorTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashErrorsErrorIdPut**](CashErrorTransactionsAPI.md#CashErrorsErrorIdPut) | **Put** /cash-errors/{error-id} | Update Cash error Transactions
[**CashErrorsGet**](CashErrorTransactionsAPI.md#CashErrorsGet) | **Get** /cash-errors | Fetch Cash error Transactions



## CashErrorsErrorIdPut

> ResponseUpdateCashErrorsApiResponse CashErrorsErrorIdPut(ctx, errorId).Body(body).Execute()

Update Cash error Transactions



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
	errorId := "errorId_example" // string | Update Cash Error transactions
	body := *openapiclient.NewHandlerUpdateCashErrorsRequest(int32(9000001), "remarks of error", int32(10145824)) // HandlerUpdateCashErrorsRequest | Status

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashErrorTransactionsAPI.CashErrorsErrorIdPut(context.Background(), errorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashErrorTransactionsAPI.CashErrorsErrorIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashErrorsErrorIdPut`: ResponseUpdateCashErrorsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashErrorTransactionsAPI.CashErrorsErrorIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**errorId** | **string** | Update Cash Error transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiCashErrorsErrorIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateCashErrorsRequest**](HandlerUpdateCashErrorsRequest.md) | Status | 

### Return type

[**ResponseUpdateCashErrorsApiResponse**](ResponseUpdateCashErrorsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashErrorsGet

> ResponseListCashErrorsApiResponse CashErrorsGet(ctx).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).IsRaisingOffice(isRaisingOffice).Execute()

Fetch Cash error Transactions



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
	isRaisingOffice := true // bool | Is Raising Office (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashErrorTransactionsAPI.CashErrorsGet(context.Background()).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).IsRaisingOffice(isRaisingOffice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashErrorTransactionsAPI.CashErrorsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashErrorsGet`: ResponseListCashErrorsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashErrorTransactionsAPI.CashErrorsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashErrorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **type_** | **string** | Type | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **isRaisingOffice** | **bool** | Is Raising Office | 

### Return type

[**ResponseListCashErrorsApiResponse**](ResponseListCashErrorsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


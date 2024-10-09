# \StampErrorTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsErrorTransactionsErrorIdPut**](StampErrorTransactionsAPI.md#StampsErrorTransactionsErrorIdPut) | **Put** /stamps-error-transactions/{error-id} | Update stamp error Transactions
[**StampsErrorTransactionsGet**](StampErrorTransactionsAPI.md#StampsErrorTransactionsGet) | **Get** /stamps-error-transactions | Get stamp error Transactions



## StampsErrorTransactionsErrorIdPut

> ResponseUpdateStampsErrorTransactionsApiResponse StampsErrorTransactionsErrorIdPut(ctx, errorId).Body(body).Execute()

Update stamp error Transactions



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
	body := *openapiclient.NewHandlerUpdateStampErrorTransactionsRequest(int32(9000001), "remarks of error", int32(10145824)) // HandlerUpdateStampErrorTransactionsRequest | Update Stamp Error transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampErrorTransactionsAPI.StampsErrorTransactionsErrorIdPut(context.Background(), errorId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampErrorTransactionsAPI.StampsErrorTransactionsErrorIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsErrorTransactionsErrorIdPut`: ResponseUpdateStampsErrorTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampErrorTransactionsAPI.StampsErrorTransactionsErrorIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**errorId** | **string** | Error ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsErrorTransactionsErrorIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateStampErrorTransactionsRequest**](HandlerUpdateStampErrorTransactionsRequest.md) | Update Stamp Error transactions | 

### Return type

[**ResponseUpdateStampsErrorTransactionsApiResponse**](ResponseUpdateStampsErrorTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsErrorTransactionsGet

> ResponseListStampsErrorTransactionsApiRespnse StampsErrorTransactionsGet(ctx).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).IsRaisingOffice(isRaisingOffice).Execute()

Get stamp error Transactions



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
	resp, r, err := apiClient.StampErrorTransactionsAPI.StampsErrorTransactionsGet(context.Background()).OfficeId(officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).IsRaisingOffice(isRaisingOffice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampErrorTransactionsAPI.StampsErrorTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsErrorTransactionsGet`: ResponseListStampsErrorTransactionsApiRespnse
	fmt.Fprintf(os.Stdout, "Response from `StampErrorTransactionsAPI.StampsErrorTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsErrorTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **type_** | **string** | Type | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **isRaisingOffice** | **bool** | Is Raising Office | 

### Return type

[**ResponseListStampsErrorTransactionsApiRespnse**](ResponseListStampsErrorTransactionsApiRespnse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


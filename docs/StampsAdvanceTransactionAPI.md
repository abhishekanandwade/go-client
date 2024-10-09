# \StampsAdvanceTransactionAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdStampsAdvanceTxnBalancesGet**](StampsAdvanceTransactionAPI.md#OfficesOfficeIdStampsAdvanceTxnBalancesGet) | **Get** /offices/{office-id}/stamps-advance-txn/balances | Get List of All Advance Balances
[**OfficesOfficeIdStampsAdvanceTxnGet**](StampsAdvanceTransactionAPI.md#OfficesOfficeIdStampsAdvanceTxnGet) | **Get** /offices/{office-id}/stamps-advance-txn | Get List of Stamp Advance Transactions
[**OfficesOfficeIdStampsAdvanceTxnPost**](StampsAdvanceTransactionAPI.md#OfficesOfficeIdStampsAdvanceTxnPost) | **Post** /offices/{office-id}/stamps-advance-txn | Create Stamp Advance Transaction
[**StampsAdvanceTxnTransactionIdApprovePut**](StampsAdvanceTransactionAPI.md#StampsAdvanceTxnTransactionIdApprovePut) | **Put** /stamps-advance-txn/{transaction-id}/approve | Update Stamp Advance Transaction



## OfficesOfficeIdStampsAdvanceTxnBalancesGet

> ResponseListStampAdvanceBalanceApiResponse OfficesOfficeIdStampsAdvanceTxnBalancesGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Execute()

Get List of All Advance Balances



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
	fromDate := "fromDate_example" // string | From date (optional)
	toDate := "toDate_example" // string | To date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnBalancesGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdStampsAdvanceTxnBalancesGet`: ResponseListStampAdvanceBalanceApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdStampsAdvanceTxnBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | From date | 
 **toDate** | **string** | To date | 

### Return type

[**ResponseListStampAdvanceBalanceApiResponse**](ResponseListStampAdvanceBalanceApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdStampsAdvanceTxnGet

> ResponseListAdvanceTxnApiResponse OfficesOfficeIdStampsAdvanceTxnGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Execute()

Get List of Stamp Advance Transactions



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
	fromDate := "fromDate_example" // string | From date (optional)
	toDate := "toDate_example" // string | To date (optional)
	isPdg := true // bool | Is pdg (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdStampsAdvanceTxnGet`: ResponseListAdvanceTxnApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdStampsAdvanceTxnGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | From date | 
 **toDate** | **string** | To date | 
 **isPdg** | **bool** | Is pdg | 

### Return type

[**ResponseListAdvanceTxnApiResponse**](ResponseListAdvanceTxnApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdStampsAdvanceTxnPost

> ResponseCreateStampsAdvanceTxnsApiResponse OfficesOfficeIdStampsAdvanceTxnPost(ctx, officeId).Body(body).Execute()

Create Stamp Advance Transaction



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
	body := *openapiclient.NewHandlerCreateStampsAdvanceTxnsRequest(float32(500.0), map[string]int32{"key": int32(123)}, int32(10145824), "Issue/ Sale/ Return", int32(10145826)) // HandlerCreateStampsAdvanceTxnsRequest | Request body for creating a new stamp advance transaction (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdStampsAdvanceTxnPost`: ResponseCreateStampsAdvanceTxnsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsAdvanceTransactionAPI.OfficesOfficeIdStampsAdvanceTxnPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdStampsAdvanceTxnPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerCreateStampsAdvanceTxnsRequest**](HandlerCreateStampsAdvanceTxnsRequest.md) | Request body for creating a new stamp advance transaction | 

### Return type

[**ResponseCreateStampsAdvanceTxnsApiResponse**](ResponseCreateStampsAdvanceTxnsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsAdvanceTxnTransactionIdApprovePut

> ResponseStampsAdvanceTxnsApiResponse StampsAdvanceTxnTransactionIdApprovePut(ctx, transactionId).Body(body).Execute()

Update Stamp Advance Transaction



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
	body := *openapiclient.NewHandlerUpdateStampAdvanceTxnRequest(int32(10145826)) // HandlerUpdateStampAdvanceTxnRequest | Is Approved (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsAdvanceTransactionAPI.StampsAdvanceTxnTransactionIdApprovePut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsAdvanceTransactionAPI.StampsAdvanceTxnTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsAdvanceTxnTransactionIdApprovePut`: ResponseStampsAdvanceTxnsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsAdvanceTransactionAPI.StampsAdvanceTxnTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsAdvanceTxnTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateStampAdvanceTxnRequest**](HandlerUpdateStampAdvanceTxnRequest.md) | Is Approved | 

### Return type

[**ResponseStampsAdvanceTxnsApiResponse**](ResponseStampsAdvanceTxnsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \MiscTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MiscTransactionsAccountCodesGet**](MiscTransactionsAPI.md#MiscTransactionsAccountCodesGet) | **Get** /misc-transactions/account-codes | List account codes
[**MiscTransactionsGet**](MiscTransactionsAPI.md#MiscTransactionsGet) | **Get** /misc-transactions | List Misc Treasury Transaction
[**MiscTransactionsPost**](MiscTransactionsAPI.md#MiscTransactionsPost) | **Post** /misc-transactions | Create a New Misc Treasury Transaction
[**MiscTransactionsTransactionIdApprovePut**](MiscTransactionsAPI.md#MiscTransactionsTransactionIdApprovePut) | **Put** /misc-transactions/{transaction-id}/approve | Approve Misc Treasury Transaction
[**MiscTransactionsTransactionIdGet**](MiscTransactionsAPI.md#MiscTransactionsTransactionIdGet) | **Get** /misc-transactions/{transaction-id} | List Misc Treasury Transaction by ID



## MiscTransactionsAccountCodesGet

> ResponseListAccountCodesApiResponse MiscTransactionsAccountCodesGet(ctx).ReceiptOrPmt(receiptOrPmt).Execute()

List account codes



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
	receiptOrPmt := "receiptOrPmt_example" // string | List account codes for receipt or pmt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscTransactionsAPI.MiscTransactionsAccountCodesGet(context.Background()).ReceiptOrPmt(receiptOrPmt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscTransactionsAPI.MiscTransactionsAccountCodesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MiscTransactionsAccountCodesGet`: ResponseListAccountCodesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `MiscTransactionsAPI.MiscTransactionsAccountCodesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiscTransactionsAccountCodesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **receiptOrPmt** | **string** | List account codes for receipt or pmt | 

### Return type

[**ResponseListAccountCodesApiResponse**](ResponseListAccountCodesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscTransactionsGet

> ResponseListMiscTransactionsApiResponse MiscTransactionsGet(ctx).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Execute()

List Misc Treasury Transaction



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
	isPdg := true // bool | Is PDG (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscTransactionsAPI.MiscTransactionsGet(context.Background()).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscTransactionsAPI.MiscTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MiscTransactionsGet`: ResponseListMiscTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `MiscTransactionsAPI.MiscTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiscTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **isPdg** | **bool** | Is PDG | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 

### Return type

[**ResponseListMiscTransactionsApiResponse**](ResponseListMiscTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscTransactionsPost

> HandlerTreasuryMiscTransactionResponse MiscTransactionsPost(ctx).Body(body).Execute()

Create a New Misc Treasury Transaction



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
	body := *openapiclient.NewHandlerCreateMiscTransactionsRequest(int32(112256), "this is some reamrks", float32(123), map[string]float32{"key": float32(123)}, "Receipt/Payment", "cash", int32(10145826)) // HandlerCreateMiscTransactionsRequest | Information about the new miscellaneous treasury transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscTransactionsAPI.MiscTransactionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscTransactionsAPI.MiscTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MiscTransactionsPost`: HandlerTreasuryMiscTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `MiscTransactionsAPI.MiscTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMiscTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateMiscTransactionsRequest**](HandlerCreateMiscTransactionsRequest.md) | Information about the new miscellaneous treasury transaction | 

### Return type

[**HandlerTreasuryMiscTransactionResponse**](HandlerTreasuryMiscTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscTransactionsTransactionIdApprovePut

> ResponseMiscTransactionsApiResponse MiscTransactionsTransactionIdApprovePut(ctx, transactionId).Body(body).Execute()

Approve Misc Treasury Transaction



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
	transactionId := "transactionId_example" // string | Transaction-id
	body := *openapiclient.NewHandlerUpdateMiscTransactionsRequest(int32(10130000), "Txn Approved", "TX00000001") // HandlerUpdateMiscTransactionsRequest | Is Approved (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscTransactionsAPI.MiscTransactionsTransactionIdApprovePut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscTransactionsAPI.MiscTransactionsTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MiscTransactionsTransactionIdApprovePut`: ResponseMiscTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `MiscTransactionsAPI.MiscTransactionsTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction-id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMiscTransactionsTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateMiscTransactionsRequest**](HandlerUpdateMiscTransactionsRequest.md) | Is Approved | 

### Return type

[**ResponseMiscTransactionsApiResponse**](ResponseMiscTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MiscTransactionsTransactionIdGet

> HandlerTreasuryMiscTransactionResponse MiscTransactionsTransactionIdGet(ctx, transactionId).Execute()

List Misc Treasury Transaction by ID



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
	transactionId := "transactionId_example" // string | transaction-id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MiscTransactionsAPI.MiscTransactionsTransactionIdGet(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MiscTransactionsAPI.MiscTransactionsTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MiscTransactionsTransactionIdGet`: HandlerTreasuryMiscTransactionResponse
	fmt.Fprintf(os.Stdout, "Response from `MiscTransactionsAPI.MiscTransactionsTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | transaction-id | 

### Other Parameters

Other parameters are passed through a pointer to a apiMiscTransactionsTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HandlerTreasuryMiscTransactionResponse**](HandlerTreasuryMiscTransactionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


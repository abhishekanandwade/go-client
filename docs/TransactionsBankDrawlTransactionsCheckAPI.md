# \TransactionsBankDrawlTransactionsCheckAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsBankDrawlCheckGet**](TransactionsBankDrawlTransactionsCheckAPI.md#TransactionsBankDrawlCheckGet) | **Get** /transactions/bank-drawl-check | Get Bank drawl limits checked



## TransactionsBankDrawlCheckGet

> ResponseFetchBankDrawlCheckApiResponse TransactionsBankDrawlCheckGet(ctx).OfficeId(officeId).Execute()

Get Bank drawl limits checked



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsBankDrawlTransactionsCheckAPI.TransactionsBankDrawlCheckGet(context.Background()).OfficeId(officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsBankDrawlTransactionsCheckAPI.TransactionsBankDrawlCheckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsBankDrawlCheckGet`: ResponseFetchBankDrawlCheckApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsBankDrawlTransactionsCheckAPI.TransactionsBankDrawlCheckGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsBankDrawlCheckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 

### Return type

[**ResponseFetchBankDrawlCheckApiResponse**](ResponseFetchBankDrawlCheckApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


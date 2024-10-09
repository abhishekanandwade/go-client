# \CashUpdateWalletAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsUpdateWalletBalancePost**](CashUpdateWalletAPI.md#TransactionsUpdateWalletBalancePost) | **Post** /transactions/update-wallet-balance | Update wallet balance



## TransactionsUpdateWalletBalancePost

> TransactionsUpdateWalletBalancePost(ctx).Body(body).Execute()

Update wallet balance



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
	body := *openapiclient.NewHandlerUpdateTransactionsWalletBalanceRequest(float32(25000.55), int32(9000003), int32(10130000)) // HandlerUpdateTransactionsWalletBalanceRequest | Update wallet balance

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CashUpdateWalletAPI.TransactionsUpdateWalletBalancePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashUpdateWalletAPI.TransactionsUpdateWalletBalancePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsUpdateWalletBalancePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerUpdateTransactionsWalletBalanceRequest**](HandlerUpdateTransactionsWalletBalanceRequest.md) | Update wallet balance | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


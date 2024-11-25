# \TransaactionsCreateBankAddlCreditAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsBankAddlCreditPost**](TransaactionsCreateBankAddlCreditAPI.md#TransactionsBankAddlCreditPost) | **Post** /transactions/bank-addl-credit | Create a bank additional credit



## TransactionsBankAddlCreditPost

> ResponseCreateBankAddlCreditApiResponse TransactionsBankAddlCreditPost(ctx).CreateBankAddlCreditRequest(createBankAddlCreditRequest).Execute()

Create a bank additional credit



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
	createBankAddlCreditRequest := *openapiclient.NewHandlerCreateBankAddlCreditRequest(float32(200009.35), "21280551", "For payments", "10122222") // HandlerCreateBankAddlCreditRequest | Create Bank Addl Credit request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransaactionsCreateBankAddlCreditAPI.TransactionsBankAddlCreditPost(context.Background()).CreateBankAddlCreditRequest(createBankAddlCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransaactionsCreateBankAddlCreditAPI.TransactionsBankAddlCreditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsBankAddlCreditPost`: ResponseCreateBankAddlCreditApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransaactionsCreateBankAddlCreditAPI.TransactionsBankAddlCreditPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsBankAddlCreditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBankAddlCreditRequest** | [**HandlerCreateBankAddlCreditRequest**](HandlerCreateBankAddlCreditRequest.md) | Create Bank Addl Credit request | 

### Return type

[**ResponseCreateBankAddlCreditApiResponse**](ResponseCreateBankAddlCreditApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


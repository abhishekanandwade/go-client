# \TransactionsUpdateBankAddlCreditRequestAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsTransactionIdBankAddlCreditApprovePut**](TransactionsUpdateBankAddlCreditRequestAPI.md#TransactionsTransactionIdBankAddlCreditApprovePut) | **Put** /transactions/{transaction-id}/bank-addl-credit/approve | Update Bank Additinal Credit request



## TransactionsTransactionIdBankAddlCreditApprovePut

> ResponseUpdateBankAddlCreditTxnApiResponse TransactionsTransactionIdBankAddlCreditApprovePut(ctx, transactionId).UpdateBankAddlCreditRequest(updateBankAddlCreditRequest).Execute()

Update Bank Additinal Credit request



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
	updateBankAddlCreditRequest := *openapiclient.NewHandlerUpdateBankAddlCreditRequest(float32(200009.35), "10122222") // HandlerUpdateBankAddlCreditRequest | Update Bank Addl Credit Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsUpdateBankAddlCreditRequestAPI.TransactionsTransactionIdBankAddlCreditApprovePut(context.Background(), transactionId).UpdateBankAddlCreditRequest(updateBankAddlCreditRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsUpdateBankAddlCreditRequestAPI.TransactionsTransactionIdBankAddlCreditApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsTransactionIdBankAddlCreditApprovePut`: ResponseUpdateBankAddlCreditTxnApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsUpdateBankAddlCreditRequestAPI.TransactionsTransactionIdBankAddlCreditApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdBankAddlCreditApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBankAddlCreditRequest** | [**HandlerUpdateBankAddlCreditRequest**](HandlerUpdateBankAddlCreditRequest.md) | Update Bank Addl Credit Request | 

### Return type

[**ResponseUpdateBankAddlCreditTxnApiResponse**](ResponseUpdateBankAddlCreditTxnApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


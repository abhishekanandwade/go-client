# \TransactionsPdgChequeTransactionsForRemittanceAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsChequesToBankOrAoGet**](TransactionsPdgChequeTransactionsForRemittanceAPI.md#TransactionsChequesToBankOrAoGet) | **Get** /transactions/cheques-to-bank-or-ao | Get Pending Cheque Transactions For Remittance to Bank/ Account Office
[**TransactionsTransactionIdSendChequesToBankOrAoPut**](TransactionsPdgChequeTransactionsForRemittanceAPI.md#TransactionsTransactionIdSendChequesToBankOrAoPut) | **Put** /transactions/{transaction-id}/send-cheques-to-bank-or-ao | Get Pending Cheque Transactions For Remittance to Bank/ Account Office



## TransactionsChequesToBankOrAoGet

> ResponseListPdgChequesToBankOrAOApiResponse TransactionsChequesToBankOrAoGet(ctx).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Pending Cheque Transactions For Remittance to Bank/ Account Office



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
	officeId := "officeId_example" // string | Office ID ex: 90000003
	isPdg := true // bool | Is Pending (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsChequesToBankOrAoGet(context.Background()).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsChequesToBankOrAoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsChequesToBankOrAoGet`: ResponseListPdgChequesToBankOrAOApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsChequesToBankOrAoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsChequesToBankOrAoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID ex: 90000003 | 
 **isPdg** | **bool** | Is Pending | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListPdgChequesToBankOrAOApiResponse**](ResponseListPdgChequesToBankOrAOApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdSendChequesToBankOrAoPut

> ResponseSendChequesToBankOrAOApiResponse TransactionsTransactionIdSendChequesToBankOrAoPut(ctx, transactionId).SendChequesToBankOrAORequest(sendChequesToBankOrAORequest).Execute()

Get Pending Cheque Transactions For Remittance to Bank/ Account Office



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
	sendChequesToBankOrAORequest := *openapiclient.NewHandlerSendChequesToBankOrAORequest(float32(15434.75), "01-03-2024", "345654", int32(21280551), int32(21260551)) // HandlerSendChequesToBankOrAORequest | Information about pending Transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsTransactionIdSendChequesToBankOrAoPut(context.Background(), transactionId).SendChequesToBankOrAORequest(sendChequesToBankOrAORequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsTransactionIdSendChequesToBankOrAoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsTransactionIdSendChequesToBankOrAoPut`: ResponseSendChequesToBankOrAOApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsPdgChequeTransactionsForRemittanceAPI.TransactionsTransactionIdSendChequesToBankOrAoPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdSendChequesToBankOrAoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendChequesToBankOrAORequest** | [**HandlerSendChequesToBankOrAORequest**](HandlerSendChequesToBankOrAORequest.md) | Information about pending Transactions | 

### Return type

[**ResponseSendChequesToBankOrAOApiResponse**](ResponseSendChequesToBankOrAOApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


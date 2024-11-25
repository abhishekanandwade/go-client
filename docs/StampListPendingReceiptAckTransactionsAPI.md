# \StampListPendingReceiptAckTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsTransactionsPdgProcessOrAckGet**](StampListPendingReceiptAckTransactionsAPI.md#StampsTransactionsPdgProcessOrAckGet) | **Get** /stamps-transactions/pdg-process-or-ack | Get Pending stamp receipt ack Transactions List



## StampsTransactionsPdgProcessOrAckGet

> ResponseListStamspTransactionsProcessOrAckApiResponse StampsTransactionsPdgProcessOrAckGet(ctx).OfficeId(officeId).Skip(skip).Limit(limit).Execute()

Get Pending stamp receipt ack Transactions List



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
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampListPendingReceiptAckTransactionsAPI.StampsTransactionsPdgProcessOrAckGet(context.Background()).OfficeId(officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampListPendingReceiptAckTransactionsAPI.StampsTransactionsPdgProcessOrAckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsTransactionsPdgProcessOrAckGet`: ResponseListStamspTransactionsProcessOrAckApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampListPendingReceiptAckTransactionsAPI.StampsTransactionsPdgProcessOrAckGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsTransactionsPdgProcessOrAckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListStamspTransactionsProcessOrAckApiResponse**](ResponseListStamspTransactionsProcessOrAckApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


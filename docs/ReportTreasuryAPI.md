# \ReportTreasuryAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsCashRemittanceReportGet**](ReportTreasuryAPI.md#TransactionsCashRemittanceReportGet) | **Get** /transactions/cash-remittance-report | Fetch treasury transactions



## TransactionsCashRemittanceReportGet

> ResponseListTransactionsApiResponse TransactionsCashRemittanceReportGet(ctx).RemittanceId(remittanceId).Execute()

Fetch treasury transactions



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
	remittanceId := "remittanceId_example" // string | Report about list of Treasury Transactions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportTreasuryAPI.TransactionsCashRemittanceReportGet(context.Background()).RemittanceId(remittanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportTreasuryAPI.TransactionsCashRemittanceReportGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsCashRemittanceReportGet`: ResponseListTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportTreasuryAPI.TransactionsCashRemittanceReportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsCashRemittanceReportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remittanceId** | **string** | Report about list of Treasury Transactions | 

### Return type

[**ResponseListTransactionsApiResponse**](ResponseListTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ReportSpecialRemittanceAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsSpecialRemittanceSlipGet**](ReportSpecialRemittanceAPI.md#TransactionsSpecialRemittanceSlipGet) | **Get** /transactions/special-remittance-slip | Fetch special remittance



## TransactionsSpecialRemittanceSlipGet

> ResponseListTransactionsSpecialRemittanceSlipApiResponse TransactionsSpecialRemittanceSlipGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Type_(type_).Execute()

Fetch special remittance



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
	officeId := "officeId_example" // string | Office ID (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	type_ := "type__example" // string | Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportSpecialRemittanceAPI.TransactionsSpecialRemittanceSlipGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportSpecialRemittanceAPI.TransactionsSpecialRemittanceSlipGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsSpecialRemittanceSlipGet`: ResponseListTransactionsSpecialRemittanceSlipApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportSpecialRemittanceAPI.TransactionsSpecialRemittanceSlipGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsSpecialRemittanceSlipGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **type_** | **string** | Type | 

### Return type

[**ResponseListTransactionsSpecialRemittanceSlipApiResponse**](ResponseListTransactionsSpecialRemittanceSlipApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


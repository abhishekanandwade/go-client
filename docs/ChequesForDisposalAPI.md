# \ChequesForDisposalAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdChequesGet**](ChequesForDisposalAPI.md#OfficesOfficeIdChequesGet) | **Get** /offices/{office-id}/cheques | Fetch cheques for disposal



## OfficesOfficeIdChequesGet

> ResponseListChequesApiResponse OfficesOfficeIdChequesGet(ctx, officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).ReportDate(reportDate).Execute()

Fetch cheques for disposal



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
	type_ := "type__example" // string | Type
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	reportDate := "reportDate_example" // string | Report Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChequesForDisposalAPI.OfficesOfficeIdChequesGet(context.Background(), officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChequesForDisposalAPI.OfficesOfficeIdChequesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdChequesGet`: ResponseListChequesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ChequesForDisposalAPI.OfficesOfficeIdChequesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdChequesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **reportDate** | **string** | Report Date | 

### Return type

[**ResponseListChequesApiResponse**](ResponseListChequesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


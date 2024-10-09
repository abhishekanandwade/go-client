# \BODAdetailsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdAccountingDetailsBodaGet**](BODAdetailsAPI.md#OfficesOfficeIdAccountingDetailsBodaGet) | **Get** /offices/{office-id}/accounting-details/boda | List BODA Details



## OfficesOfficeIdAccountingDetailsBodaGet

> ResponseFetchBODADataApiResponse OfficesOfficeIdAccountingDetailsBodaGet(ctx, officeId).ReportDate(reportDate).Execute()

List BODA Details



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
	officeId := "officeId_example" // string | Fetch accounting details of an office for a day godoc (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch accounting details of an office for a day godoc (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BODAdetailsAPI.OfficesOfficeIdAccountingDetailsBodaGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BODAdetailsAPI.OfficesOfficeIdAccountingDetailsBodaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdAccountingDetailsBodaGet`: ResponseFetchBODADataApiResponse
	fmt.Fprintf(os.Stdout, "Response from `BODAdetailsAPI.OfficesOfficeIdAccountingDetailsBodaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch accounting details of an office for a day godoc (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdAccountingDetailsBodaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch accounting details of an office for a day godoc (example: 2024-01-31) | 

### Return type

[**ResponseFetchBODADataApiResponse**](ResponseFetchBODADataApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


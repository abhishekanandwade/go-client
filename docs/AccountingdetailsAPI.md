# \AccountingdetailsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdAccountingDetailsGet**](AccountingdetailsAPI.md#OfficesOfficeIdAccountingDetailsGet) | **Get** /offices/{office-id}/accounting-details | List Accounting Details by account code
[**OfficesOfficeIdAccountingDetailsPost**](AccountingdetailsAPI.md#OfficesOfficeIdAccountingDetailsPost) | **Post** /offices/{office-id}/accounting-details | Create a New Accounting transaction
[**OfficesOfficeIdAccountingDetailsSummaryGet**](AccountingdetailsAPI.md#OfficesOfficeIdAccountingDetailsSummaryGet) | **Get** /offices/{office-id}/accounting-details/summary | List Summary Details



## OfficesOfficeIdAccountingDetailsGet

> ResponseListDetailsByAccountCodeApiResponse OfficesOfficeIdAccountingDetailsGet(ctx, officeId).Type_(type_).AccountCode(accountCode).FromDate(fromDate).ToDate(toDate).ReportDate(reportDate).Skip(skip).Limit(limit).Execute()

List Accounting Details by account code



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
	officeId := "officeId_example" // string | Office Id (example: 90000003)
	type_ := "type__example" // string | Type (example: list)
	accountCode := "accountCode_example" // string | Account-code (example: 8671000500) (optional)
	fromDate := "fromDate_example" // string | From-date  (example: 2024-01-01) (optional)
	toDate := "toDate_example" // string | To-date  (example: 2024-01-31) (optional)
	reportDate := "reportDate_example" // string | Report-date (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | skip (example: 1) (optional)
	limit := int32(56) // int32 | Limit (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsGet(context.Background(), officeId).Type_(type_).AccountCode(accountCode).FromDate(fromDate).ToDate(toDate).ReportDate(reportDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdAccountingDetailsGet`: ResponseListDetailsByAccountCodeApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Office Id (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdAccountingDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type (example: list) | 
 **accountCode** | **string** | Account-code (example: 8671000500) | 
 **fromDate** | **string** | From-date  (example: 2024-01-01) | 
 **toDate** | **string** | To-date  (example: 2024-01-31) | 
 **reportDate** | **string** | Report-date (example: 2024-01-31) | 
 **skip** | **int32** | skip (example: 1) | 
 **limit** | **int32** | Limit (example: 10) | 

### Return type

[**ResponseListDetailsByAccountCodeApiResponse**](ResponseListDetailsByAccountCodeApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdAccountingDetailsPost

> ResponseCreateAccountingTxnAPIResponse OfficesOfficeIdAccountingDetailsPost(ctx, officeId).Body(body).Execute()

Create a New Accounting transaction



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
	officeId := int32(56) // int32 | Office ID (example: 9000003)
	body := *openapiclient.NewHandlerCreateOfficesAccountingDetailsRequest("Treasury/PoS/DARPAN", "POS78967", "2024-01-11") // HandlerCreateOfficesAccountingDetailsRequest | Information about adding new accounting transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdAccountingDetailsPost`: ResponseCreateAccountingTxnAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID (example: 9000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdAccountingDetailsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerCreateOfficesAccountingDetailsRequest**](HandlerCreateOfficesAccountingDetailsRequest.md) | Information about adding new accounting transaction details | 

### Return type

[**ResponseCreateAccountingTxnAPIResponse**](ResponseCreateAccountingTxnAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdAccountingDetailsSummaryGet

> ResponseFetchSummaryDetailsApiResponse OfficesOfficeIdAccountingDetailsSummaryGet(ctx, officeId).ReportDate(reportDate).Execute()

List Summary Details



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
	resp, r, err := apiClient.AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsSummaryGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdAccountingDetailsSummaryGet`: ResponseFetchSummaryDetailsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountingdetailsAPI.OfficesOfficeIdAccountingDetailsSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch accounting details of an office for a day godoc (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdAccountingDetailsSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch accounting details of an office for a day godoc (example: 2024-01-31) | 

### Return type

[**ResponseFetchSummaryDetailsApiResponse**](ResponseFetchSummaryDetailsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


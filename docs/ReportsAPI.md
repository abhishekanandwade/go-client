# \ReportsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdReportsAccountBalancesGet**](ReportsAPI.md#OfficesOfficeIdReportsAccountBalancesGet) | **Get** /offices/{office-id}/reports/account-balances | Get account Balances Report
[**OfficesOfficeIdReportsCashBalancesGet**](ReportsAPI.md#OfficesOfficeIdReportsCashBalancesGet) | **Get** /offices/{office-id}/reports/cash-balances | Get Cash Balances Report
[**OfficesOfficeIdReportsCashInTransitGet**](ReportsAPI.md#OfficesOfficeIdReportsCashInTransitGet) | **Get** /offices/{office-id}/reports/cash-in-transit | Get Cash in Transit Report
[**OfficesOfficeIdReportsChequesInTransitGet**](ReportsAPI.md#OfficesOfficeIdReportsChequesInTransitGet) | **Get** /offices/{office-id}/reports/cheques-in-transit | Get Cheques in Transit Report
[**OfficesOfficeIdReportsDayBeginEndGet**](ReportsAPI.md#OfficesOfficeIdReportsDayBeginEndGet) | **Get** /offices/{office-id}/reports/day-begin-end | Get Day Begin End Report
[**OfficesOfficeIdReportsIpoBalancesGet**](ReportsAPI.md#OfficesOfficeIdReportsIpoBalancesGet) | **Get** /offices/{office-id}/reports/ipo-balances | Get IPO Balances Report
[**OfficesOfficeIdReportsIposInTransitGet**](ReportsAPI.md#OfficesOfficeIdReportsIposInTransitGet) | **Get** /offices/{office-id}/reports/ipos-in-transit | Get IPOs in Transit Report
[**OfficesOfficeIdReportsOutOfStockGet**](ReportsAPI.md#OfficesOfficeIdReportsOutOfStockGet) | **Get** /offices/{office-id}/reports/out-of-stock | Get Out-Of-Stock Inventory by Office ID
[**OfficesOfficeIdReportsPaymentDetailsGet**](ReportsAPI.md#OfficesOfficeIdReportsPaymentDetailsGet) | **Get** /offices/{office-id}/reports/payment-details | Get Payment Details Report
[**OfficesOfficeIdReportsPostmanTransactionsGet**](ReportsAPI.md#OfficesOfficeIdReportsPostmanTransactionsGet) | **Get** /offices/{office-id}/reports/postman-transactions | Get Postman Transactions Report
[**OfficesOfficeIdReportsStampBalancesGet**](ReportsAPI.md#OfficesOfficeIdReportsStampBalancesGet) | **Get** /offices/{office-id}/reports/stamp-balances | Get Stamp Balances Report
[**OfficesOfficeIdReportsStampBalancesLastSupplyGet**](ReportsAPI.md#OfficesOfficeIdReportsStampBalancesLastSupplyGet) | **Get** /offices/{office-id}/reports/stamp-balances-last-supply | Get Stamp Balances Report with last supply
[**OfficesOfficeIdReportsStampsInTransitGet**](ReportsAPI.md#OfficesOfficeIdReportsStampsInTransitGet) | **Get** /offices/{office-id}/reports/stamps-in-transit | Get Stamps in Transit Report
[**OfficesOfficeIdReportsTcbAllGet**](ReportsAPI.md#OfficesOfficeIdReportsTcbAllGet) | **Get** /offices/{office-id}/reports/tcb-all | Get TCB Data Report of all users
[**OfficesOfficeIdReportsTcbDenomDetailsPut**](ReportsAPI.md#OfficesOfficeIdReportsTcbDenomDetailsPut) | **Put** /offices/{office-id}/reports/tcb-denom-details | Update TCB Denomination Details
[**OfficesOfficeIdReportsTcbGet**](ReportsAPI.md#OfficesOfficeIdReportsTcbGet) | **Get** /offices/{office-id}/reports/tcb | Get TCB Data Report
[**OfficesOfficeIdReportsTcbNewGet**](ReportsAPI.md#OfficesOfficeIdReportsTcbNewGet) | **Get** /offices/{office-id}/reports/tcb-new | Get TCB Data Report
[**OfficesOfficeIdReportsTransitDashboardGet**](ReportsAPI.md#OfficesOfficeIdReportsTransitDashboardGet) | **Get** /offices/{office-id}/reports/transit-dashboard | Get Transit dashboard Report



## OfficesOfficeIdReportsAccountBalancesGet

> ResponseCashBalanceAPIResponse OfficesOfficeIdReportsAccountBalancesGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get account Balances Report



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
	officeId := "officeId_example" // string | office ID (example: 90000003)
	fromDate := "fromDate_example" // string | from date (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | to date (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | skip records (example: 0) (optional)
	limit := int32(56) // int32 | limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsAccountBalancesGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsAccountBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsAccountBalancesGet`: ResponseCashBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsAccountBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | office ID (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsAccountBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | from date (example: 2024-01-31) | 
 **toDate** | **string** | to date (example: 2024-01-31) | 
 **skip** | **int32** | skip records (example: 0) | 
 **limit** | **int32** | limit records (example: 10) | 

### Return type

[**ResponseCashBalanceAPIResponse**](ResponseCashBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsCashBalancesGet

> ResponseCashBalanceAPIResponse OfficesOfficeIdReportsCashBalancesGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Cash Balances Report



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
	officeId := "officeId_example" // string | Fetch cash balances Data (example: 90000003)
	fromDate := "fromDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsCashBalancesGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsCashBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsCashBalancesGet`: ResponseCashBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsCashBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch cash balances Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsCashBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **toDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseCashBalanceAPIResponse**](ResponseCashBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsCashInTransitGet

> ResponseCashInTransitAPIResponse OfficesOfficeIdReportsCashInTransitGet(ctx, officeId).ReportDate(reportDate).Execute()

Get Cash in Transit Report



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
	officeId := "officeId_example" // string | Fetch Cash in Tranist transactions (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch Cash in Tranist transactions (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsCashInTransitGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsCashInTransitGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsCashInTransitGet`: ResponseCashInTransitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsCashInTransitGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Cash in Tranist transactions (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsCashInTransitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch Cash in Tranist transactions (example: 2024-01-31) | 

### Return type

[**ResponseCashInTransitAPIResponse**](ResponseCashInTransitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsChequesInTransitGet

> ResponseChequesInTransitAPIResponse OfficesOfficeIdReportsChequesInTransitGet(ctx, officeId).ReportDate(reportDate).Execute()

Get Cheques in Transit Report



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
	officeId := "officeId_example" // string | Fetch Cash in Tranist transactions (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch Cash in Tranist transactions (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsChequesInTransitGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsChequesInTransitGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsChequesInTransitGet`: ResponseChequesInTransitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsChequesInTransitGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Cash in Tranist transactions (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsChequesInTransitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch Cash in Tranist transactions (example: 2024-01-31) | 

### Return type

[**ResponseChequesInTransitAPIResponse**](ResponseChequesInTransitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsDayBeginEndGet

> ResponseDayBeginEndListAPIResponse OfficesOfficeIdReportsDayBeginEndGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Day Begin End Report



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
	officeId := "officeId_example" // string | Fetch Day Begin End Data
	fromDate := "fromDate_example" // string | Fetch Day Begin End Data
	toDate := "toDate_example" // string | Fetch Day Begin End Data
	skip := int32(56) // int32 | Skip records (optional)
	limit := int32(56) // int32 | Limit records (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsDayBeginEndGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsDayBeginEndGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsDayBeginEndGet`: ResponseDayBeginEndListAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsDayBeginEndGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Day Begin End Data | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsDayBeginEndGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch Day Begin End Data | 
 **toDate** | **string** | Fetch Day Begin End Data | 
 **skip** | **int32** | Skip records | 
 **limit** | **int32** | Limit records | 

### Return type

[**ResponseDayBeginEndListAPIResponse**](ResponseDayBeginEndListAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsIpoBalancesGet

> ResponseIPOBalanceAPIResponse OfficesOfficeIdReportsIpoBalancesGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get IPO Balances Report



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
	officeId := "officeId_example" // string | Fetch cash balances Data (example: 90000003)
	fromDate := "fromDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsIpoBalancesGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsIpoBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsIpoBalancesGet`: ResponseIPOBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsIpoBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch cash balances Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsIpoBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **toDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseIPOBalanceAPIResponse**](ResponseIPOBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsIposInTransitGet

> ResponseIPOsInTransitAPIResponse OfficesOfficeIdReportsIposInTransitGet(ctx, officeId).ReportDate(reportDate).Execute()

Get IPOs in Transit Report



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
	officeId := "officeId_example" // string | Fetch Cash in Tranist transactions (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch Cash in Tranist transactions (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsIposInTransitGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsIposInTransitGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsIposInTransitGet`: ResponseIPOsInTransitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsIposInTransitGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Cash in Tranist transactions (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsIposInTransitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch Cash in Tranist transactions (example: 2024-01-31) | 

### Return type

[**ResponseIPOsInTransitAPIResponse**](ResponseIPOsInTransitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsOutOfStockGet

> ResponseOutOfStockInventoryAPIResponse OfficesOfficeIdReportsOutOfStockGet(ctx, officeId).Skip(skip).Limit(limit).Execute()

Get Out-Of-Stock Inventory by Office ID



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
	officeId := "officeId_example" // string | Office ID (example: 90000003)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsOutOfStockGet(context.Background(), officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsOutOfStockGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsOutOfStockGet`: ResponseOutOfStockInventoryAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsOutOfStockGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Office ID (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsOutOfStockGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseOutOfStockInventoryAPIResponse**](ResponseOutOfStockInventoryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsPaymentDetailsGet

> ResponseListPaymentDetailsResponse OfficesOfficeIdReportsPaymentDetailsGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Payment Details Report



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
	officeId := "officeId_example" // string | Office ID (example: 90000003)
	fromDate := "fromDate_example" // string | From date (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | To date (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsPaymentDetailsGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsPaymentDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsPaymentDetailsGet`: ResponseListPaymentDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsPaymentDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Office ID (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsPaymentDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | From date (example: 2024-01-31) | 
 **toDate** | **string** | To date (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseListPaymentDetailsResponse**](ResponseListPaymentDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsPostmanTransactionsGet

> ResponsePostmanTransactionsAPIResponse OfficesOfficeIdReportsPostmanTransactionsGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Postman Transactions Report



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
	officeId := "officeId_example" // string | Fetch cash balances Data (example: 90000003)
	fromDate := "fromDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsPostmanTransactionsGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsPostmanTransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsPostmanTransactionsGet`: ResponsePostmanTransactionsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsPostmanTransactionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch cash balances Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsPostmanTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **toDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponsePostmanTransactionsAPIResponse**](ResponsePostmanTransactionsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsStampBalancesGet

> ResponseListStampBalanceAPIResponse OfficesOfficeIdReportsStampBalancesGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Stamp Balances Report



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
	officeId := "officeId_example" // string | Fetch cash balances Data (example: 90000003)
	fromDate := "fromDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsStampBalancesGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsStampBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsStampBalancesGet`: ResponseListStampBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsStampBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch cash balances Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsStampBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **toDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseListStampBalanceAPIResponse**](ResponseListStampBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsStampBalancesLastSupplyGet

> ResponseFetchReportsStampBalanceLastSupplyResponse OfficesOfficeIdReportsStampBalancesLastSupplyGet(ctx, officeId).FromDate(fromDate).ToDate(toDate).Execute()

Get Stamp Balances Report with last supply



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
	officeId := "officeId_example" // string | Fetch cash balances Data (example: 90000003)
	fromDate := "fromDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)
	toDate := "toDate_example" // string | Fetch cash balances Data (example: 2024-01-31) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsStampBalancesLastSupplyGet(context.Background(), officeId).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsStampBalancesLastSupplyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsStampBalancesLastSupplyGet`: ResponseFetchReportsStampBalanceLastSupplyResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsStampBalancesLastSupplyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch cash balances Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsStampBalancesLastSupplyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 
 **toDate** | **string** | Fetch cash balances Data (example: 2024-01-31) | 

### Return type

[**ResponseFetchReportsStampBalanceLastSupplyResponse**](ResponseFetchReportsStampBalanceLastSupplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsStampsInTransitGet

> ResponseStampsInTransitAPIResponse OfficesOfficeIdReportsStampsInTransitGet(ctx, officeId).ReportDate(reportDate).Execute()

Get Stamps in Transit Report



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
	officeId := "officeId_example" // string | Fetch Cash in Tranist transactions (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch Cash in Tranist transactions (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsStampsInTransitGet(context.Background(), officeId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsStampsInTransitGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsStampsInTransitGet`: ResponseStampsInTransitAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsStampsInTransitGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Cash in Tranist transactions (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsStampsInTransitGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch Cash in Tranist transactions (example: 2024-01-31) | 

### Return type

[**ResponseStampsInTransitAPIResponse**](ResponseStampsInTransitAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsTcbAllGet

> ResponseListTcbBalanceAPIResponse OfficesOfficeIdReportsTcbAllGet(ctx, officeId).ReportDate(reportDate).Skip(skip).Limit(limit).Execute()

Get TCB Data Report of all users



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
	officeId := "officeId_example" // string | Fetch TCB Data of all users (example: 90000003)
	reportDate := "reportDate_example" // string | Fetch TCB Data of all users (example: 2024-01-31)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsTcbAllGet(context.Background(), officeId).ReportDate(reportDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsTcbAllGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsTcbAllGet`: ResponseListTcbBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsTcbAllGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch TCB Data of all users (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsTcbAllGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportDate** | **string** | Fetch TCB Data of all users (example: 2024-01-31) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseListTcbBalanceAPIResponse**](ResponseListTcbBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsTcbDenomDetailsPut

> ResponseUpdateReportsTCBDenomDetailsApiResponse OfficesOfficeIdReportsTcbDenomDetailsPut(ctx, officeId).TransDate(transDate).Execute()

Update TCB Denomination Details



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
	officeId := int32(56) // int32 | Update TCB denomination details to be given (example: 112256)
	transDate := "transDate_example" // string | Update TCB denomination details to be given (example: 2024-12-23)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsTcbDenomDetailsPut(context.Background(), officeId).TransDate(transDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsTcbDenomDetailsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsTcbDenomDetailsPut`: ResponseUpdateReportsTCBDenomDetailsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsTcbDenomDetailsPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Update TCB denomination details to be given (example: 112256) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsTcbDenomDetailsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transDate** | **string** | Update TCB denomination details to be given (example: 2024-12-23) | 

### Return type

[**ResponseUpdateReportsTCBDenomDetailsApiResponse**](ResponseUpdateReportsTCBDenomDetailsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsTcbGet

> ResponseTcbBalanceAPIResponse OfficesOfficeIdReportsTcbGet(ctx, officeId).UserId(userId).ReportDate(reportDate).Execute()

Get TCB Data Report



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
	officeId := "officeId_example" // string | Fetch TCB Data (example: 90000003)
	userId := "userId_example" // string | Fetch TCB Data (example: 10132232)
	reportDate := "reportDate_example" // string | Fetch TCB Data (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsTcbGet(context.Background(), officeId).UserId(userId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsTcbGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsTcbGet`: ResponseTcbBalanceAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsTcbGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch TCB Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsTcbGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Fetch TCB Data (example: 10132232) | 
 **reportDate** | **string** | Fetch TCB Data (example: 2024-01-31) | 

### Return type

[**ResponseTcbBalanceAPIResponse**](ResponseTcbBalanceAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsTcbNewGet

> ResponseListReportsTCBDataResponse OfficesOfficeIdReportsTcbNewGet(ctx, officeId).UserId(userId).ReportDate(reportDate).Execute()

Get TCB Data Report



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
	officeId := "officeId_example" // string | Fetch TCB Data (example: 90000003)
	userId := "userId_example" // string | Fetch TCB Data (example: 10132232)
	reportDate := "reportDate_example" // string | Fetch TCB Data (example: 2024-01-31)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsTcbNewGet(context.Background(), officeId).UserId(userId).ReportDate(reportDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsTcbNewGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsTcbNewGet`: ResponseListReportsTCBDataResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsTcbNewGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch TCB Data (example: 90000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsTcbNewGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Fetch TCB Data (example: 10132232) | 
 **reportDate** | **string** | Fetch TCB Data (example: 2024-01-31) | 

### Return type

[**ResponseListReportsTCBDataResponse**](ResponseListReportsTCBDataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdReportsTransitDashboardGet

> ResponseTransitDetailsAPIResponse OfficesOfficeIdReportsTransitDashboardGet(ctx, officeId).Skip(skip).Limit(limit).Execute()

Get Transit dashboard Report



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
	officeId := "officeId_example" // string | Fetch Tranist Dashboard Data
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportsAPI.OfficesOfficeIdReportsTransitDashboardGet(context.Background(), officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.OfficesOfficeIdReportsTransitDashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdReportsTransitDashboardGet`: ResponseTransitDetailsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.OfficesOfficeIdReportsTransitDashboardGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Fetch Tranist Dashboard Data | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdReportsTransitDashboardGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseTransitDetailsAPIResponse**](ResponseTransitDetailsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


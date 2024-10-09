# \IPOOBCSDReceiptAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposObCsdReceiptsObTransactionIdApprovePost**](IPOOBCSDReceiptAPI.md#IposObCsdReceiptsObTransactionIdApprovePost) | **Post** /ipos-ob-csd-receipts/{ob-transaction-id}/approve | Approve IPO Transaction Request
[**OfficesOfficeIdIposObCsdReceiptsPost**](IPOOBCSDReceiptAPI.md#OfficesOfficeIdIposObCsdReceiptsPost) | **Post** /offices/{office-id}/ipos-ob-csd-receipts | Get All Pending IPO Transaction Request
[**OfficesOfficeIdIposObCsdReceiptsStockPost**](IPOOBCSDReceiptAPI.md#OfficesOfficeIdIposObCsdReceiptsStockPost) | **Post** /offices/{office-id}/ipos-ob-csd-receipts/stock | Get IPO Stock Of Office



## IposObCsdReceiptsObTransactionIdApprovePost

> ResponseIPOsOBReceiptsApiResponse IposObCsdReceiptsObTransactionIdApprovePost(ctx, obTransactionId).Body(body).Execute()

Approve IPO Transaction Request



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
	obTransactionId := "obTransactionId_example" // string | Transaction ID
	body := *openapiclient.NewHandlerUpdateIPOsOBCSDReceiptsApproveRequest(int32(10130000)) // HandlerUpdateIPOsOBCSDReceiptsApproveRequest | Approver Remarks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePost(context.Background(), obTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposObCsdReceiptsObTransactionIdApprovePost`: ResponseIPOsOBReceiptsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obTransactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposObCsdReceiptsObTransactionIdApprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateIPOsOBCSDReceiptsApproveRequest**](HandlerUpdateIPOsOBCSDReceiptsApproveRequest.md) | Approver Remarks | 

### Return type

[**ResponseIPOsOBReceiptsApiResponse**](ResponseIPOsOBReceiptsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdIposObCsdReceiptsPost

> ResponseListIPOOBApprovalPdgApiResponse OfficesOfficeIdIposObCsdReceiptsPost(ctx, officeId).IsApproval(isApproval).FromDate(fromDate).ToDate(toDate).Execute()

Get All Pending IPO Transaction Request



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
	officeId := int32(56) // int32 | Office ID
	isApproval := true // bool | Is Approval (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsPost(context.Background(), officeId).IsApproval(isApproval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdIposObCsdReceiptsPost`: ResponseListIPOOBApprovalPdgApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdIposObCsdReceiptsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isApproval** | **bool** | Is Approval | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 

### Return type

[**ResponseListIPOOBApprovalPdgApiResponse**](ResponseListIPOOBApprovalPdgApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdIposObCsdReceiptsStockPost

> ResponseListOfficeIPOsObcsdReceiptsStockApiResponse OfficesOfficeIdIposObCsdReceiptsStockPost(ctx, officeId).Execute()

Get IPO Stock Of Office



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
	officeId := int32(56) // int32 | Office ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockPost(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdIposObCsdReceiptsStockPost`: ResponseListOfficeIPOsObcsdReceiptsStockApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdIposObCsdReceiptsStockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseListOfficeIPOsObcsdReceiptsStockApiResponse**](ResponseListOfficeIPOsObcsdReceiptsStockApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


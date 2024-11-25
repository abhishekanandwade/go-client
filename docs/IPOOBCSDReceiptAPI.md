# \IPOOBCSDReceiptAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposObCsdReceiptsObTransactionIdApprovePut**](IPOOBCSDReceiptAPI.md#IposObCsdReceiptsObTransactionIdApprovePut) | **Put** /ipos-ob-csd-receipts/{ob-transaction-id}/approve | Approve IPO Transaction Request
[**OfficesOfficeIdIposObCsdReceiptsGet**](IPOOBCSDReceiptAPI.md#OfficesOfficeIdIposObCsdReceiptsGet) | **Get** /offices/{office-id}/ipos-ob-csd-receipts | Get All Pending IPO Transaction Request
[**OfficesOfficeIdIposObCsdReceiptsPost**](IPOOBCSDReceiptAPI.md#OfficesOfficeIdIposObCsdReceiptsPost) | **Post** /offices/{office-id}/ipos-ob-csd-receipts | Create a New IPO OB/ CSD Receipt Transaction Request
[**OfficesOfficeIdIposObCsdReceiptsStockGet**](IPOOBCSDReceiptAPI.md#OfficesOfficeIdIposObCsdReceiptsStockGet) | **Get** /offices/{office-id}/ipos-ob-csd-receipts/stock | Get IPO Stock Of Office



## IposObCsdReceiptsObTransactionIdApprovePut

> ResponseIPOsOBReceiptsApiResponse IposObCsdReceiptsObTransactionIdApprovePut(ctx, obTransactionId).Body(body).Execute()

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
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePut(context.Background(), obTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposObCsdReceiptsObTransactionIdApprovePut`: ResponseIPOsOBReceiptsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.IposObCsdReceiptsObTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obTransactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposObCsdReceiptsObTransactionIdApprovePutRequest struct via the builder pattern


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


## OfficesOfficeIdIposObCsdReceiptsGet

> ResponseListIPOOBApprovalPdgApiResponse OfficesOfficeIdIposObCsdReceiptsGet(ctx, officeId).IsApproval(isApproval).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

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
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsGet(context.Background(), officeId).IsApproval(isApproval).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdIposObCsdReceiptsGet`: ResponseListIPOOBApprovalPdgApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdIposObCsdReceiptsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isApproval** | **bool** | Is Approval | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

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


## OfficesOfficeIdIposObCsdReceiptsPost

> ResponseCreateIPOsOBReceiptsApiResponse OfficesOfficeIdIposObCsdReceiptsPost(ctx, officeId).Body(body).Execute()

Create a New IPO OB/ CSD Receipt Transaction Request



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
	body := *openapiclient.NewHandlerCreateOfficesIPOsOBReceiptsRequest("2023-12-07T00:00:00Z", int32(10130000), "CSDH/2-3/DLGS", "OPENING BALANCE/ CSD") // HandlerCreateOfficesIPOsOBReceiptsRequest | IPO Details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdIposObCsdReceiptsPost`: ResponseCreateIPOsOBReceiptsApiResponse
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

 **body** | [**HandlerCreateOfficesIPOsOBReceiptsRequest**](HandlerCreateOfficesIPOsOBReceiptsRequest.md) | IPO Details | 

### Return type

[**ResponseCreateIPOsOBReceiptsApiResponse**](ResponseCreateIPOsOBReceiptsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdIposObCsdReceiptsStockGet

> ResponseListOfficeIPOsObcsdReceiptsStockApiResponse OfficesOfficeIdIposObCsdReceiptsStockGet(ctx, officeId).Skip(skip).Limit(limit).Execute()

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
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockGet(context.Background(), officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdIposObCsdReceiptsStockGet`: ResponseListOfficeIPOsObcsdReceiptsStockApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOOBCSDReceiptAPI.OfficesOfficeIdIposObCsdReceiptsStockGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdIposObCsdReceiptsStockGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

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


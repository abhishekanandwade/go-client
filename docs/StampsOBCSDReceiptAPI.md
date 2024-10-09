# \StampsOBCSDReceiptAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampObCsdReceiptsGet**](StampsOBCSDReceiptAPI.md#StampObCsdReceiptsGet) | **Get** /stamp-ob-csd-receipts | Get CSD Opening Balance Receipts
[**StampObCsdReceiptsObTransactionIdApprovePut**](StampsOBCSDReceiptAPI.md#StampObCsdReceiptsObTransactionIdApprovePut) | **Put** /stamp-ob-csd-receipts/{ob-transaction-id}/approve | Approve CSD Opening Balance Receipts
[**StampObCsdReceiptsObTransactionIdGet**](StampsOBCSDReceiptAPI.md#StampObCsdReceiptsObTransactionIdGet) | **Get** /stamp-ob-csd-receipts/{ob-transaction-id} | Get CSD Opening Balance Receipts by Transaction ID



## StampObCsdReceiptsGet

> ResponseListStampObCsdReceiptsApiResponse StampObCsdReceiptsGet(ctx).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Execute()

Get CSD Opening Balance Receipts



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
	isPdg := true // bool | Is PDG (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsOBCSDReceiptAPI.StampObCsdReceiptsGet(context.Background()).OfficeId(officeId).IsPdg(isPdg).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsOBCSDReceiptAPI.StampObCsdReceiptsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampObCsdReceiptsGet`: ResponseListStampObCsdReceiptsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsOBCSDReceiptAPI.StampObCsdReceiptsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampObCsdReceiptsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **int32** | Office ID | 
 **isPdg** | **bool** | Is PDG | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 

### Return type

[**ResponseListStampObCsdReceiptsApiResponse**](ResponseListStampObCsdReceiptsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampObCsdReceiptsObTransactionIdApprovePut

> ResponseStampObCsdReceiptsApiResponse StampObCsdReceiptsObTransactionIdApprovePut(ctx, obTransactionId).Body(body).Execute()

Approve CSD Opening Balance Receipts



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
	body := *openapiclient.NewHandlerUpdateStampObCsdReceiptsRequest(int32(10123232)) // HandlerUpdateStampObCsdReceiptsRequest | Approve stamp opening balances received from CSD

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdApprovePut(context.Background(), obTransactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampObCsdReceiptsObTransactionIdApprovePut`: ResponseStampObCsdReceiptsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdApprovePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obTransactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampObCsdReceiptsObTransactionIdApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateStampObCsdReceiptsRequest**](HandlerUpdateStampObCsdReceiptsRequest.md) | Approve stamp opening balances received from CSD | 

### Return type

[**ResponseStampObCsdReceiptsApiResponse**](ResponseStampObCsdReceiptsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampObCsdReceiptsObTransactionIdGet

> ResponseStampObCsdReceiptsApiResponse StampObCsdReceiptsObTransactionIdGet(ctx, obTransactionId).Execute()

Get CSD Opening Balance Receipts by Transaction ID



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
	obTransactionId := "obTransactionId_example" // string | Transaction ID of the opening balance receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdGet(context.Background(), obTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampObCsdReceiptsObTransactionIdGet`: ResponseStampObCsdReceiptsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsOBCSDReceiptAPI.StampObCsdReceiptsObTransactionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**obTransactionId** | **string** | Transaction ID of the opening balance receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampObCsdReceiptsObTransactionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseStampObCsdReceiptsApiResponse**](ResponseStampObCsdReceiptsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


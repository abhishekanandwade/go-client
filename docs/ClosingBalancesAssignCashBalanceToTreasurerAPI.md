# \ClosingBalancesAssignCashBalanceToTreasurerAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdClosingBalancesAssignToTreasurerPut**](ClosingBalancesAssignCashBalanceToTreasurerAPI.md#OfficesOfficeIdClosingBalancesAssignToTreasurerPut) | **Put** /offices/{office-id}/closing-balances-assign-to-treasurer | Assigning of Office cash balance to Treasurer



## OfficesOfficeIdClosingBalancesAssignToTreasurerPut

> ResponseUpdateClosingBalancesTransferApiResponse OfficesOfficeIdClosingBalancesAssignToTreasurerPut(ctx, officeId).Body(body).Execute()

Assigning of Office cash balance to Treasurer



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
	officeId := int32(56) // int32 | Assign Office Cash balance to Treasurer (example: 900001)
	body := *openapiclient.NewHandlerUpdateClosingBalancesAssignToTransferRequest(int32(10012323)) // HandlerUpdateClosingBalancesAssignToTransferRequest | Assign Office Cash balance to Treasurer (example: 10012323)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClosingBalancesAssignCashBalanceToTreasurerAPI.OfficesOfficeIdClosingBalancesAssignToTreasurerPut(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClosingBalancesAssignCashBalanceToTreasurerAPI.OfficesOfficeIdClosingBalancesAssignToTreasurerPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdClosingBalancesAssignToTreasurerPut`: ResponseUpdateClosingBalancesTransferApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ClosingBalancesAssignCashBalanceToTreasurerAPI.OfficesOfficeIdClosingBalancesAssignToTreasurerPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Assign Office Cash balance to Treasurer (example: 900001) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdClosingBalancesAssignToTreasurerPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateClosingBalancesAssignToTransferRequest**](HandlerUpdateClosingBalancesAssignToTransferRequest.md) | Assign Office Cash balance to Treasurer (example: 10012323) | 

### Return type

[**ResponseUpdateClosingBalancesTransferApiResponse**](ResponseUpdateClosingBalancesTransferApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


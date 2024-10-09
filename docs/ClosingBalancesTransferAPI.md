# \ClosingBalancesTransferAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdClosingBalancesTransferPut**](ClosingBalancesTransferAPI.md#OfficesOfficeIdClosingBalancesTransferPut) | **Put** /offices/{office-id}/closing-balances-transfer | Transfer closing balances from one employee to another



## OfficesOfficeIdClosingBalancesTransferPut

> ResponseUpdateClosingBalancesTransferApiResponse OfficesOfficeIdClosingBalancesTransferPut(ctx, officeId).Type_(type_).Execute()

Transfer closing balances from one employee to another



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
	officeId := int32(56) // int32 | Transfer closing balances
	type_ := "type__example" // string | Transfer closing balances from one employee to another

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClosingBalancesTransferAPI.OfficesOfficeIdClosingBalancesTransferPut(context.Background(), officeId).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClosingBalancesTransferAPI.OfficesOfficeIdClosingBalancesTransferPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdClosingBalancesTransferPut`: ResponseUpdateClosingBalancesTransferApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ClosingBalancesTransferAPI.OfficesOfficeIdClosingBalancesTransferPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Transfer closing balances | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdClosingBalancesTransferPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Transfer closing balances from one employee to another | 

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


# \ClosingBalancesFetchAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdClosingBalancesGet**](ClosingBalancesFetchAPI.md#OfficesOfficeIdClosingBalancesGet) | **Get** /offices/{office-id}/closing-balances | Fetch treasury balances of an employee



## OfficesOfficeIdClosingBalancesGet

> ResponseFetchClosingBalancesApiResponse OfficesOfficeIdClosingBalancesGet(ctx, officeId).UserId(userId).Execute()

Fetch treasury balances of an employee



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
	officeId := int32(56) // int32 | Fetch closing balances
	userId := int32(56) // int32 | Fetch closing balances of an employee

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClosingBalancesFetchAPI.OfficesOfficeIdClosingBalancesGet(context.Background(), officeId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClosingBalancesFetchAPI.OfficesOfficeIdClosingBalancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdClosingBalancesGet`: ResponseFetchClosingBalancesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ClosingBalancesFetchAPI.OfficesOfficeIdClosingBalancesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Fetch closing balances | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdClosingBalancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **int32** | Fetch closing balances of an employee | 

### Return type

[**ResponseFetchClosingBalancesApiResponse**](ResponseFetchClosingBalancesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


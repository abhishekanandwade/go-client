# \ToolsCheckDayEndAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdDayEndStatusGet**](ToolsCheckDayEndAPI.md#OfficesOfficeIdDayEndStatusGet) | **Get** /offices/{office-id}/day-end-status | Check Office Day End Status



## OfficesOfficeIdDayEndStatusGet

> ResponseFetchOfficesDayEndStatusResponse OfficesOfficeIdDayEndStatusGet(ctx, officeId).Execute()

Check Office Day End Status



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
	officeId := int32(56) // int32 | Information about day end status details of an office specified by office-id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsCheckDayEndAPI.OfficesOfficeIdDayEndStatusGet(context.Background(), officeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsCheckDayEndAPI.OfficesOfficeIdDayEndStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdDayEndStatusGet`: ResponseFetchOfficesDayEndStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsCheckDayEndAPI.OfficesOfficeIdDayEndStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Information about day end status details of an office specified by office-id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdDayEndStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseFetchOfficesDayEndStatusResponse**](ResponseFetchOfficesDayEndStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


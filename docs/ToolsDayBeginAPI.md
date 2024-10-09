# \ToolsDayBeginAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdPerformDayBeginPost**](ToolsDayBeginAPI.md#OfficesOfficeIdPerformDayBeginPost) | **Post** /offices/{office-id}/perform-day-begin | Perform Day Begin



## OfficesOfficeIdPerformDayBeginPost

> DomainDayBeginEnd OfficesOfficeIdPerformDayBeginPost(ctx, officeId).Body(body).Execute()

Perform Day Begin



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
	officeId := int32(56) // int32 | Information about office id
	body := *openapiclient.NewHandlerCreateOfficesPerformDayBeginRequest(int32(10130000)) // HandlerCreateOfficesPerformDayBeginRequest | Information about creating new day begin details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsDayBeginAPI.OfficesOfficeIdPerformDayBeginPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsDayBeginAPI.OfficesOfficeIdPerformDayBeginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdPerformDayBeginPost`: DomainDayBeginEnd
	fmt.Fprintf(os.Stdout, "Response from `ToolsDayBeginAPI.OfficesOfficeIdPerformDayBeginPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Information about office id | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdPerformDayBeginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerCreateOfficesPerformDayBeginRequest**](HandlerCreateOfficesPerformDayBeginRequest.md) | Information about creating new day begin details | 

### Return type

[**DomainDayBeginEnd**](DomainDayBeginEnd.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


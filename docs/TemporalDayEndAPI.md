# \TemporalDayEndAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemporalDayEndPost**](TemporalDayEndAPI.md#TemporalDayEndPost) | **Post** /temporal/day-end | Temporal Day End



## TemporalDayEndPost

> ResponseTemporalApiResponse TemporalDayEndPost(ctx).Body(body).Execute()

Temporal Day End



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
	body := *openapiclient.NewHandlerTemporalDayEndCombinedRequest() // HandlerTemporalDayEndCombinedRequest | Information about adding new accounting transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalDayEndAPI.TemporalDayEndPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalDayEndAPI.TemporalDayEndPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemporalDayEndPost`: ResponseTemporalApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TemporalDayEndAPI.TemporalDayEndPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemporalDayEndPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerTemporalDayEndCombinedRequest**](HandlerTemporalDayEndCombinedRequest.md) | Information about adding new accounting transaction details | 

### Return type

[**ResponseTemporalApiResponse**](ResponseTemporalApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


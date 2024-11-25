# \TemporalDayBeginAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemporalDayBeginPost**](TemporalDayBeginAPI.md#TemporalDayBeginPost) | **Post** /temporal/day-begin | Temporal Day Begin



## TemporalDayBeginPost

> ResponseTemporalApiResponse TemporalDayBeginPost(ctx).Body(body).Execute()

Temporal Day Begin



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
	body := *openapiclient.NewHandlerTemporalDayBeginCombinedRequest() // HandlerTemporalDayBeginCombinedRequest | Information about adding new accounting transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalDayBeginAPI.TemporalDayBeginPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalDayBeginAPI.TemporalDayBeginPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemporalDayBeginPost`: ResponseTemporalApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TemporalDayBeginAPI.TemporalDayBeginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemporalDayBeginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerTemporalDayBeginCombinedRequest**](HandlerTemporalDayBeginCombinedRequest.md) | Information about adding new accounting transaction details | 

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


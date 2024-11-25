# \TemporalBagOpenAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemporalBagClosePost**](TemporalBagOpenAPI.md#TemporalBagClosePost) | **Post** /temporal/bag-close | Temporal Bag Close
[**TemporalBagOpenPost**](TemporalBagOpenAPI.md#TemporalBagOpenPost) | **Post** /temporal/bag-open | Temporal Bag Open



## TemporalBagClosePost

> ResponseTemporalApiResponse TemporalBagClosePost(ctx).Body(body).Execute()

Temporal Bag Close



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
	body := *openapiclient.NewHandlerTemporalcombinedCashbagcloseRequest() // HandlerTemporalcombinedCashbagcloseRequest | Information about adding new accounting transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalBagOpenAPI.TemporalBagClosePost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalBagOpenAPI.TemporalBagClosePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemporalBagClosePost`: ResponseTemporalApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TemporalBagOpenAPI.TemporalBagClosePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemporalBagClosePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerTemporalcombinedCashbagcloseRequest**](HandlerTemporalcombinedCashbagcloseRequest.md) | Information about adding new accounting transaction details | 

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


## TemporalBagOpenPost

> ResponseTemporalApiResponse TemporalBagOpenPost(ctx).Body(body).Execute()

Temporal Bag Open



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
	body := *openapiclient.NewHandlerTemporalCombinedRequestbagopenRequest() // HandlerTemporalCombinedRequestbagopenRequest | Information about adding new accounting transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemporalBagOpenAPI.TemporalBagOpenPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemporalBagOpenAPI.TemporalBagOpenPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TemporalBagOpenPost`: ResponseTemporalApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TemporalBagOpenAPI.TemporalBagOpenPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTemporalBagOpenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerTemporalCombinedRequestbagopenRequest**](HandlerTemporalCombinedRequestbagopenRequest.md) | Information about adding new accounting transaction details | 

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


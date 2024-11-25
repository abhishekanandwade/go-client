# \IPOsSalesAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposSalesPost**](IPOsSalesAPI.md#IposSalesPost) | **Post** /ipos-sales | Create IPO sales



## IposSalesPost

> ResponseCreateIPOSaleApiResponse IposSalesPost(ctx).CreateIPOSalesRequest(createIPOSalesRequest).Execute()

Create IPO sales



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
	createIPOSalesRequest := *openapiclient.NewHandlerCreateIPOSalesRequest(int32(9000003), float32(300.65), float32(300.65), []int32{int32(123)}, "XXX", int32(10130000)) // HandlerCreateIPOSalesRequest |  None                                                  

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOsSalesAPI.IposSalesPost(context.Background()).CreateIPOSalesRequest(createIPOSalesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOsSalesAPI.IposSalesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposSalesPost`: ResponseCreateIPOSaleApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOsSalesAPI.IposSalesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposSalesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIPOSalesRequest** | [**HandlerCreateIPOSalesRequest**](HandlerCreateIPOSalesRequest.md) |  None                                                   | 

### Return type

[**ResponseCreateIPOSaleApiResponse**](ResponseCreateIPOSaleApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


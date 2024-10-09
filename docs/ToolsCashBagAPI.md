# \ToolsCashBagAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashBagsIssuedChequesPost**](ToolsCashBagAPI.md#CashBagsIssuedChequesPost) | **Post** /cash-bags/issued-cheques | Include issued cheques in cash bag



## CashBagsIssuedChequesPost

> ResponseCreateCashBagsIssuesChequesApiResponse CashBagsIssuedChequesPost(ctx).Body(body).Execute()

Include issued cheques in cash bag



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
	body := *openapiclient.NewHandlerCreateCashBagsIssuesChequesRequest(int32(90000003), "Remarks", int32(90000003), float32(2234.5), "Cheque No. 112223 dtd. 21.07.24", "TX00001") // HandlerCreateCashBagsIssuesChequesRequest | Include cheque issue in cash bag

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsCashBagAPI.CashBagsIssuedChequesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsCashBagAPI.CashBagsIssuedChequesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashBagsIssuedChequesPost`: ResponseCreateCashBagsIssuesChequesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsCashBagAPI.CashBagsIssuedChequesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashBagsIssuedChequesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateCashBagsIssuesChequesRequest**](HandlerCreateCashBagsIssuesChequesRequest.md) | Include cheque issue in cash bag | 

### Return type

[**ResponseCreateCashBagsIssuesChequesApiResponse**](ResponseCreateCashBagsIssuesChequesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \IPOSalesAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposSalesGet**](IPOSalesAPI.md#IposSalesGet) | **Get** /ipos-sales | Get IPO Sales List



## IposSalesGet

> ResponseListIPOSalesApiResponse IposSalesGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()

Get IPO Sales List



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
	officeId := "officeId_example" // string | Office ID
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	isPdg := true // bool | Is PDG (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPOSalesAPI.IposSalesGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).IsPdg(isPdg).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPOSalesAPI.IposSalesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposSalesGet`: ResponseListIPOSalesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `IPOSalesAPI.IposSalesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposSalesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **isPdg** | **bool** | Is PDG | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListIPOSalesApiResponse**](ResponseListIPOSalesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


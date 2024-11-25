# \ConfigureIPOsDenominationAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IposDenominationsDenominationIdGet**](ConfigureIPOsDenominationAPI.md#IposDenominationsDenominationIdGet) | **Get** /ipos-denominations/{denomination-id} | Get Specific IPO Denominations
[**IposDenominationsGet**](ConfigureIPOsDenominationAPI.md#IposDenominationsGet) | **Get** /ipos-denominations | Get All IPO Denominations



## IposDenominationsDenominationIdGet

> ResponseIPOsDenominationsAPIResponse IposDenominationsDenominationIdGet(ctx, denominationId).Execute()

Get Specific IPO Denominations



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
	denominationId := "denominationId_example" // string | Get IPO Denomination by ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureIPOsDenominationAPI.IposDenominationsDenominationIdGet(context.Background(), denominationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureIPOsDenominationAPI.IposDenominationsDenominationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposDenominationsDenominationIdGet`: ResponseIPOsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureIPOsDenominationAPI.IposDenominationsDenominationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**denominationId** | **string** | Get IPO Denomination by ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiIposDenominationsDenominationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseIPOsDenominationsAPIResponse**](ResponseIPOsDenominationsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IposDenominationsGet

> ResponseListIPOsDenominationsAPIResponse IposDenominationsGet(ctx).Skip(skip).Limit(limit).OrderBy(orderBy).SortType(sortType).Execute()

Get All IPO Denominations



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
	skip := int32(56) // int32 | Number of IPO denominations to be skipped. (optional)
	limit := int32(56) // int32 | Limit of the number of IPO denominations to be fetched. (optional)
	orderBy := "orderBy_example" // string | Order by which the IPO denominations are to be sorted. (optional)
	sortType := "sortType_example" // string | Sort type of the IPO denominations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureIPOsDenominationAPI.IposDenominationsGet(context.Background()).Skip(skip).Limit(limit).OrderBy(orderBy).SortType(sortType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureIPOsDenominationAPI.IposDenominationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IposDenominationsGet`: ResponseListIPOsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureIPOsDenominationAPI.IposDenominationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIposDenominationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of IPO denominations to be skipped. | 
 **limit** | **int32** | Limit of the number of IPO denominations to be fetched. | 
 **orderBy** | **string** | Order by which the IPO denominations are to be sorted. | 
 **sortType** | **string** | Sort type of the IPO denominations. | 

### Return type

[**ResponseListIPOsDenominationsAPIResponse**](ResponseListIPOsDenominationsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


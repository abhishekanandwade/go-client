# \ConfigureStampsDenominationAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsDenominationsDenominationIdGet**](ConfigureStampsDenominationAPI.md#StampsDenominationsDenominationIdGet) | **Get** /stamps-denominations/{denomination-id} | Get Specific Stamp Denomination
[**StampsDenominationsGet**](ConfigureStampsDenominationAPI.md#StampsDenominationsGet) | **Get** /stamps-denominations | Get Specific Office Stamp Denominations



## StampsDenominationsDenominationIdGet

> ResponseStampsDenominationsAPIResponse StampsDenominationsDenominationIdGet(ctx, denominationId).Execute()

Get Specific Stamp Denomination



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
	denominationId := "denominationId_example" // string | Get stamp denomination by ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdGet(context.Background(), denominationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsDenominationsDenominationIdGet`: ResponseStampsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**denominationId** | **string** | Get stamp denomination by ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsDenominationsDenominationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseStampsDenominationsAPIResponse**](ResponseStampsDenominationsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsDenominationsGet

> ResponseListStampsDenominationsAPIResponse StampsDenominationsGet(ctx).OfficeId(officeId).Skip(skip).Limit(limit).SortType(sortType).OrderBy(orderBy).Execute()

Get Specific Office Stamp Denominations



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
	officeId := int32(56) // int32 | Get Stamp Denominations by Office ID (optional)
	skip := int32(56) // int32 | Skip records (optional)
	limit := int32(56) // int32 | Limit records (optional)
	sortType := "sortType_example" // string | Sort records (optional)
	orderBy := "orderBy_example" // string | Sort records (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsDenominationAPI.StampsDenominationsGet(context.Background()).OfficeId(officeId).Skip(skip).Limit(limit).SortType(sortType).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsDenominationAPI.StampsDenominationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsDenominationsGet`: ResponseListStampsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsDenominationAPI.StampsDenominationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsDenominationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **int32** | Get Stamp Denominations by Office ID | 
 **skip** | **int32** | Skip records | 
 **limit** | **int32** | Limit records | 
 **sortType** | **string** | Sort records | 
 **orderBy** | **string** | Sort records | 

### Return type

[**ResponseListStampsDenominationsAPIResponse**](ResponseListStampsDenominationsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


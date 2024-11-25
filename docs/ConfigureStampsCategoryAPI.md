# \ConfigureStampsCategoryAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsCategoriesCategoryIdGet**](ConfigureStampsCategoryAPI.md#StampsCategoriesCategoryIdGet) | **Get** /stamps-categories/{category-id} | Get Specific Stamp Category
[**StampsCategoriesGet**](ConfigureStampsCategoryAPI.md#StampsCategoriesGet) | **Get** /stamps-categories | Get All Stamps Category



## StampsCategoriesCategoryIdGet

> ResponseStampscategoriesAPIResponse StampsCategoriesCategoryIdGet(ctx, categoryId).Execute()

Get Specific Stamp Category



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
	categoryId := "categoryId_example" // string | Category ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdGet(context.Background(), categoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsCategoriesCategoryIdGet`: ResponseStampscategoriesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Category ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsCategoriesCategoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseStampscategoriesAPIResponse**](ResponseStampscategoriesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsCategoriesGet

> ResponseListStampsCategoriesAPIResponse StampsCategoriesGet(ctx).Skip(skip).Limit(limit).SortType(sortType).OrderBy(orderBy).Execute()

Get All Stamps Category



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
	skip := int32(56) // int32 | Number of records to skip for pagination (optional) (default to 0)
	limit := int32(56) // int32 | Number of records to limit for pagination (optional) (default to 10)
	sortType := "sortType_example" // string | Sort the records in ascending or descending order (optional)
	orderBy := "orderBy_example" // string | Order the records based on the specified field (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsCategoryAPI.StampsCategoriesGet(context.Background()).Skip(skip).Limit(limit).SortType(sortType).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsCategoryAPI.StampsCategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsCategoriesGet`: ResponseListStampsCategoriesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsCategoryAPI.StampsCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip for pagination | [default to 0]
 **limit** | **int32** | Number of records to limit for pagination | [default to 10]
 **sortType** | **string** | Sort the records in ascending or descending order | 
 **orderBy** | **string** | Order the records based on the specified field | 

### Return type

[**ResponseListStampsCategoriesAPIResponse**](ResponseListStampsCategoriesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ConfigureStampsCategoryAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsCategoriesCategoryIdGet**](ConfigureStampsCategoryAPI.md#StampsCategoriesCategoryIdGet) | **Get** /stamps-categories/{category-id} | Get Specific Stamp Category
[**StampsCategoriesCategoryIdPut**](ConfigureStampsCategoryAPI.md#StampsCategoriesCategoryIdPut) | **Put** /stamps-categories/{category-id} | Update Stamp Category
[**StampsCategoriesGet**](ConfigureStampsCategoryAPI.md#StampsCategoriesGet) | **Get** /stamps-categories | Get All Stamps Category
[**StampsCategoriesPost**](ConfigureStampsCategoryAPI.md#StampsCategoriesPost) | **Post** /stamps-categories | Create Stamps Category



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


## StampsCategoriesCategoryIdPut

> ResponseStampscategoriesAPIResponse StampsCategoriesCategoryIdPut(ctx, categoryId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()

Update Stamp Category



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
	categoryId := "categoryId_example" // string | Details identified by category Id
	updatedByUserid := "updatedByUserid_example" // string | User ID of the user updating the stamp category
	validTo := "validTo_example" // string | Date the stamp category is valid to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdPut(context.Background(), categoryId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsCategoriesCategoryIdPut`: ResponseStampscategoriesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsCategoryAPI.StampsCategoriesCategoryIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryId** | **string** | Details identified by category Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsCategoriesCategoryIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedByUserid** | **string** | User ID of the user updating the stamp category | 
 **validTo** | **string** | Date the stamp category is valid to | 

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

> ResponseListStampsCategoriesAPIResponse StampsCategoriesGet(ctx).Skip(skip).Limit(limit).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsCategoryAPI.StampsCategoriesGet(context.Background()).Skip(skip).Limit(limit).Execute()
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


## StampsCategoriesPost

> ResponseCreateStampscategoriesAPIResponse StampsCategoriesPost(ctx).Body(body).Execute()

Create Stamps Category



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
	body := *openapiclient.NewHandlerCreateStampsCategoriesRequest("Postage Stamps", "10130000", "1010123456", "2023-12-08T00:00:00Z") // HandlerCreateStampsCategoriesRequest | Information about adding new stamp category details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsCategoryAPI.StampsCategoriesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsCategoryAPI.StampsCategoriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsCategoriesPost`: ResponseCreateStampscategoriesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsCategoryAPI.StampsCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateStampsCategoriesRequest**](HandlerCreateStampsCategoriesRequest.md) | Information about adding new stamp category details | 

### Return type

[**ResponseCreateStampscategoriesAPIResponse**](ResponseCreateStampscategoriesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


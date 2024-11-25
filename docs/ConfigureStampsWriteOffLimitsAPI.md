# \ConfigureStampsWriteOffLimitsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SoiledWriteOffLimitsGet**](ConfigureStampsWriteOffLimitsAPI.md#SoiledWriteOffLimitsGet) | **Get** /soiled-write-off-limits | Get Soiled Write Off List
[**SoiledWriteOffLimitsLimitIdDelete**](ConfigureStampsWriteOffLimitsAPI.md#SoiledWriteOffLimitsLimitIdDelete) | **Delete** /soiled-write-off-limits/{limit-id} | Update Soiled Write Off Limit
[**SoiledWriteOffLimitsPost**](ConfigureStampsWriteOffLimitsAPI.md#SoiledWriteOffLimitsPost) | **Post** /soiled-write-off-limits | Create Soiled Write Off



## SoiledWriteOffLimitsGet

> ResponseListStampsSoiledWriteOffAPIResponse SoiledWriteOffLimitsGet(ctx).Skip(skip).Limit(limit).Execute()

Get Soiled Write Off List



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
	resp, r, err := apiClient.ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoiledWriteOffLimitsGet`: ResponseListStampsSoiledWriteOffAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoiledWriteOffLimitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip for pagination | [default to 0]
 **limit** | **int32** | Number of records to limit for pagination | [default to 10]

### Return type

[**ResponseListStampsSoiledWriteOffAPIResponse**](ResponseListStampsSoiledWriteOffAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoiledWriteOffLimitsLimitIdDelete

> ResponseDeleteSoiledWriteOffLimitsAPIResponse SoiledWriteOffLimitsLimitIdDelete(ctx, limitId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()

Update Soiled Write Off Limit



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
	limitId := "limitId_example" // string | Details identified by limit Id
	updatedByUserid := "updatedByUserid_example" // string | Details identified by updated by user Id
	validTo := "validTo_example" // string | Details identified by valid to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsLimitIdDelete(context.Background(), limitId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsLimitIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoiledWriteOffLimitsLimitIdDelete`: ResponseDeleteSoiledWriteOffLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsLimitIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** | Details identified by limit Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoiledWriteOffLimitsLimitIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedByUserid** | **string** | Details identified by updated by user Id | 
 **validTo** | **string** | Details identified by valid to | 

### Return type

[**ResponseDeleteSoiledWriteOffLimitsAPIResponse**](ResponseDeleteSoiledWriteOffLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SoiledWriteOffLimitsPost

> ResponseCreateStampsSoiledWriteOffLimitsAPIResponse SoiledWriteOffLimitsPost(ctx).CreateStampsSoiledWriteOffLimitsRequest(createStampsSoiledWriteOffLimitsRequest).Execute()

Create Soiled Write Off



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
	createStampsSoiledWriteOffLimitsRequest := *openapiclient.NewHandlerCreateStampsSoiledWriteOffLimitsRequest("10130000", int32(5000), "2023-12-08T00:00:00Z") // HandlerCreateStampsSoiledWriteOffLimitsRequest | Information about soiled write off limit

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsPost(context.Background()).CreateStampsSoiledWriteOffLimitsRequest(createStampsSoiledWriteOffLimitsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoiledWriteOffLimitsPost`: ResponseCreateStampsSoiledWriteOffLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsWriteOffLimitsAPI.SoiledWriteOffLimitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSoiledWriteOffLimitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStampsSoiledWriteOffLimitsRequest** | [**HandlerCreateStampsSoiledWriteOffLimitsRequest**](HandlerCreateStampsSoiledWriteOffLimitsRequest.md) | Information about soiled write off limit | 

### Return type

[**ResponseCreateStampsSoiledWriteOffLimitsAPIResponse**](ResponseCreateStampsSoiledWriteOffLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ConfigureCashConveyanceLimitAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashTransferLineLimitsGet**](ConfigureCashConveyanceLimitAPI.md#CashTransferLineLimitsGet) | **Get** /cash-transfer-line-limits | Get All Cash Transfer Line Limits Handled
[**CashTransferLineLimitsLimitIdDelete**](ConfigureCashConveyanceLimitAPI.md#CashTransferLineLimitsLimitIdDelete) | **Delete** /cash-transfer-line-limits/{limit-id} | Update Cash Transfer Line Limit
[**CashTransferLineLimitsLimitIdGet**](ConfigureCashConveyanceLimitAPI.md#CashTransferLineLimitsLimitIdGet) | **Get** /cash-transfer-line-limits/{limit-id} | Get Specific Cash Transfer Line Limit by ID
[**CashTransferLineLimitsPost**](ConfigureCashConveyanceLimitAPI.md#CashTransferLineLimitsPost) | **Post** /cash-transfer-line-limits | Create a new cash conveyance limit



## CashTransferLineLimitsGet

> ResponseListCashTransferLineLimitsAPIResponse CashTransferLineLimitsGet(ctx).Skip(skip).Limit(limit).Execute()

Get All Cash Transfer Line Limits Handled



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
	resp, r, err := apiClient.ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashTransferLineLimitsGet`: ResponseListCashTransferLineLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashTransferLineLimitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip for pagination | [default to 0]
 **limit** | **int32** | Number of records to limit for pagination | [default to 10]

### Return type

[**ResponseListCashTransferLineLimitsAPIResponse**](ResponseListCashTransferLineLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashTransferLineLimitsLimitIdDelete

> ResponseDeleteCashTransferLineLimitsAPIResponse CashTransferLineLimitsLimitIdDelete(ctx, limitId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()

Update Cash Transfer Line Limit



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
	limitId := "limitId_example" // string | limit-id
	updatedByUserid := "updatedByUserid_example" // string | updated-by-userid
	validTo := "validTo_example" // string | valid-to

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdDelete(context.Background(), limitId).UpdatedByUserid(updatedByUserid).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashTransferLineLimitsLimitIdDelete`: ResponseDeleteCashTransferLineLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** | limit-id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCashTransferLineLimitsLimitIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedByUserid** | **string** | updated-by-userid | 
 **validTo** | **string** | valid-to | 

### Return type

[**ResponseDeleteCashTransferLineLimitsAPIResponse**](ResponseDeleteCashTransferLineLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashTransferLineLimitsLimitIdGet

> ResponseCashTransferLineLimitsAPIResponse CashTransferLineLimitsLimitIdGet(ctx, limitId).Execute()

Get Specific Cash Transfer Line Limit by ID



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
	limitId := "limitId_example" // string | Fetch Cash Transfer Line Limit Request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdGet(context.Background(), limitId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashTransferLineLimitsLimitIdGet`: ResponseCashTransferLineLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsLimitIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**limitId** | **string** | Fetch Cash Transfer Line Limit Request | 

### Other Parameters

Other parameters are passed through a pointer to a apiCashTransferLineLimitsLimitIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseCashTransferLineLimitsAPIResponse**](ResponseCashTransferLineLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashTransferLineLimitsPost

> ResponseCreateCashTransferLineLimitsAPIResponse CashTransferLineLimitsPost(ctx).Body(body).Execute()

Create a new cash conveyance limit



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
	body := *openapiclient.NewHandlerCreateCashTransferLineLimitsRequest("10139999", float32(10000.0), "Loose cash through departmental employee", "9999-12-31T00:00:00Z") // HandlerCreateCashTransferLineLimitsRequest | Create Cash conveyance limit request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashTransferLineLimitsPost`: ResponseCreateCashTransferLineLimitsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCashConveyanceLimitAPI.CashTransferLineLimitsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashTransferLineLimitsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateCashTransferLineLimitsRequest**](HandlerCreateCashTransferLineLimitsRequest.md) | Create Cash conveyance limit request | 

### Return type

[**ResponseCreateCashTransferLineLimitsAPIResponse**](ResponseCreateCashTransferLineLimitsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


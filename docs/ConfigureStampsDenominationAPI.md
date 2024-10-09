# \ConfigureStampsDenominationAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsDenominationsDenominationIdGet**](ConfigureStampsDenominationAPI.md#StampsDenominationsDenominationIdGet) | **Get** /stamps-denominations/{denomination-id} | Get Specific Stamp Denomination
[**StampsDenominationsDenominationIdPut**](ConfigureStampsDenominationAPI.md#StampsDenominationsDenominationIdPut) | **Put** /stamps-denominations/{denomination-id} | Update a Stamp Denomination
[**StampsDenominationsGet**](ConfigureStampsDenominationAPI.md#StampsDenominationsGet) | **Get** /stamps-denominations | Get Specific Office Stamp Denominations
[**StampsDenominationsPost**](ConfigureStampsDenominationAPI.md#StampsDenominationsPost) | **Post** /stamps-denominations | Create Stamps Denomination



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


## StampsDenominationsDenominationIdPut

> ResponseStampsDenominationsAPIResponse StampsDenominationsDenominationIdPut(ctx, denominationId).UpdatedBy(updatedBy).ValidTo(validTo).Execute()

Update a Stamp Denomination



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
	denominationId := "denominationId_example" // string | Details identified by denomination Id
	updatedBy := "updatedBy_example" // string | User who updated the stamp denomination
	validTo := "validTo_example" // string | Valid to date

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdPut(context.Background(), denominationId).UpdatedBy(updatedBy).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsDenominationsDenominationIdPut`: ResponseStampsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsDenominationAPI.StampsDenominationsDenominationIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**denominationId** | **string** | Details identified by denomination Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsDenominationsDenominationIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatedBy** | **string** | User who updated the stamp denomination | 
 **validTo** | **string** | Valid to date | 

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

> ResponseListStampsDenominationsAPIResponse StampsDenominationsGet(ctx).OfficeId(officeId).Execute()

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
	officeId := int32(56) // int32 | Get Stamp Denominations by Office ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsDenominationAPI.StampsDenominationsGet(context.Background()).OfficeId(officeId).Execute()
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


## StampsDenominationsPost

> ResponseCreateStampsDenominationsAPIResponse StampsDenominationsPost(ctx).Body(body).Execute()

Create Stamps Denomination



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
	body := *openapiclient.NewHandlerCreateStampsDenominationsRequest("SC01", "Mahatma Gandhi 5 Rupees", float32(5.0), "10130000", "2023-12-08T00:00:00Z") // HandlerCreateStampsDenominationsRequest | Information about adding new stamp denomination details 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureStampsDenominationAPI.StampsDenominationsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureStampsDenominationAPI.StampsDenominationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsDenominationsPost`: ResponseCreateStampsDenominationsAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureStampsDenominationAPI.StampsDenominationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsDenominationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateStampsDenominationsRequest**](HandlerCreateStampsDenominationsRequest.md) | Information about adding new stamp denomination details  | 

### Return type

[**ResponseCreateStampsDenominationsAPIResponse**](ResponseCreateStampsDenominationsAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


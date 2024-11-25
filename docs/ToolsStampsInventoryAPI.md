# \ToolsStampsInventoryAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StampsInventoryGet**](ToolsStampsInventoryAPI.md#StampsInventoryGet) | **Get** /stamps-inventory | Fetch stamps inventory
[**StampsInventoryRequestsRequestIdPut**](ToolsStampsInventoryAPI.md#StampsInventoryRequestsRequestIdPut) | **Put** /stamps-inventory-requests/{request-id} | Update stamps inventory



## StampsInventoryGet

> ResponseListStampsInventoryAPIResponse StampsInventoryGet(ctx).OfficeId(officeId).Skip(skip).Limit(limit).Execute()

Fetch stamps inventory



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
	officeId := int32(56) // int32 | Office ID (example: 9000003)
	skip := int32(56) // int32 | Skip records (example: 0) (optional)
	limit := int32(56) // int32 | Limit records (example: 10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsStampsInventoryAPI.StampsInventoryGet(context.Background()).OfficeId(officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsStampsInventoryAPI.StampsInventoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsInventoryGet`: ResponseListStampsInventoryAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsStampsInventoryAPI.StampsInventoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsInventoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **int32** | Office ID (example: 9000003) | 
 **skip** | **int32** | Skip records (example: 0) | 
 **limit** | **int32** | Limit records (example: 10) | 

### Return type

[**ResponseListStampsInventoryAPIResponse**](ResponseListStampsInventoryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsInventoryRequestsRequestIdPut

> ResponseUpdateStampsInventoryAPIResponse StampsInventoryRequestsRequestIdPut(ctx, requestId).Body(body).Execute()

Update stamps inventory



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
	requestId := "requestId_example" // string | Update stamps inventory (example: REQ_21260551)
	body := *openapiclient.NewHandlerUpdateStampsInventoryRequest(int32(21), int32(9000003)) // HandlerUpdateStampsInventoryRequest | Update stamps inventory (example: REQ_21260551)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsStampsInventoryAPI.StampsInventoryRequestsRequestIdPut(context.Background(), requestId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsStampsInventoryAPI.StampsInventoryRequestsRequestIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsInventoryRequestsRequestIdPut`: ResponseUpdateStampsInventoryAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsStampsInventoryAPI.StampsInventoryRequestsRequestIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | Update stamps inventory (example: REQ_21260551) | 

### Other Parameters

Other parameters are passed through a pointer to a apiStampsInventoryRequestsRequestIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateStampsInventoryRequest**](HandlerUpdateStampsInventoryRequest.md) | Update stamps inventory (example: REQ_21260551) | 

### Return type

[**ResponseUpdateStampsInventoryAPIResponse**](ResponseUpdateStampsInventoryAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ToolsDayEndAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdDayBeginEndStatusGet**](ToolsDayEndAPI.md#OfficesOfficeIdDayBeginEndStatusGet) | **Get** /offices/{office-id}/day-begin-end-status | Check Office Day End Status
[**OfficesOfficeIdPerformDayEndPut**](ToolsDayEndAPI.md#OfficesOfficeIdPerformDayEndPut) | **Put** /offices/{office-id}/perform-day-end | Perform Day End



## OfficesOfficeIdDayBeginEndStatusGet

> ResponseFetchDayBeginEndStatusAPIResponse OfficesOfficeIdDayBeginEndStatusGet(ctx, officeId).TransDate(transDate).Status(status).Execute()

Check Office Day End Status



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
	transDate := "transDate_example" // string | transaction date (optional)
	status := "status_example" // string | Information about status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsDayEndAPI.OfficesOfficeIdDayBeginEndStatusGet(context.Background(), officeId).TransDate(transDate).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsDayEndAPI.OfficesOfficeIdDayBeginEndStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdDayBeginEndStatusGet`: ResponseFetchDayBeginEndStatusAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsDayEndAPI.OfficesOfficeIdDayBeginEndStatusGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID (example: 9000003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdDayBeginEndStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transDate** | **string** | transaction date | 
 **status** | **string** | Information about status | 

### Return type

[**ResponseFetchDayBeginEndStatusAPIResponse**](ResponseFetchDayBeginEndStatusAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdPerformDayEndPut

> ResponseDayBeginEndAPIResponse OfficesOfficeIdPerformDayEndPut(ctx, officeId).Body(body).Execute()

Perform Day End



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
	officeId := int32(56) // int32 | Information about update day end details
	body := *openapiclient.NewHandlerUpdateOfficesPerformDayEndRequest(int32(10130000), "2024-01-01") // HandlerUpdateOfficesPerformDayEndRequest | Information about update day end details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ToolsDayEndAPI.OfficesOfficeIdPerformDayEndPut(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ToolsDayEndAPI.OfficesOfficeIdPerformDayEndPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdPerformDayEndPut`: ResponseDayBeginEndAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ToolsDayEndAPI.OfficesOfficeIdPerformDayEndPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Information about update day end details | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdPerformDayEndPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateOfficesPerformDayEndRequest**](HandlerUpdateOfficesPerformDayEndRequest.md) | Information about update day end details | 

### Return type

[**ResponseDayBeginEndAPIResponse**](ResponseDayBeginEndAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


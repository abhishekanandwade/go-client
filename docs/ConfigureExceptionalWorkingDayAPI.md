# \ConfigureExceptionalWorkingDayAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdExceptionalWorkingDaysDelete**](ConfigureExceptionalWorkingDayAPI.md#OfficesOfficeIdExceptionalWorkingDaysDelete) | **Delete** /offices/{office-id}/exceptional-working-days | Delete Office Exceptional Working
[**OfficesOfficeIdExceptionalWorkingDaysGet**](ConfigureExceptionalWorkingDayAPI.md#OfficesOfficeIdExceptionalWorkingDaysGet) | **Get** /offices/{office-id}/exceptional-working-days | Get All Exceptional Working Days
[**OfficesOfficeIdExceptionalWorkingDaysPost**](ConfigureExceptionalWorkingDayAPI.md#OfficesOfficeIdExceptionalWorkingDaysPost) | **Post** /offices/{office-id}/exceptional-working-days | Create Exceptional Working Day



## OfficesOfficeIdExceptionalWorkingDaysDelete

> ResponseDeleteExceptionalWkgDayAPIResponse OfficesOfficeIdExceptionalWorkingDaysDelete(ctx, officeId).ExceptionalWkgDayDate(exceptionalWkgDayDate).Execute()

Delete Office Exceptional Working



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
	officeId := "officeId_example" // string | Delete exceptional working day (example:900003)
	exceptionalWkgDayDate := "exceptionalWkgDayDate_example" // string | Delete exceptional working day (example:2023-11-30T00:00:00Z)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysDelete(context.Background(), officeId).ExceptionalWkgDayDate(exceptionalWkgDayDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdExceptionalWorkingDaysDelete`: ResponseDeleteExceptionalWkgDayAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Delete exceptional working day (example:900003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdExceptionalWorkingDaysDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exceptionalWkgDayDate** | **string** | Delete exceptional working day (example:2023-11-30T00:00:00Z) | 

### Return type

[**ResponseDeleteExceptionalWkgDayAPIResponse**](ResponseDeleteExceptionalWkgDayAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdExceptionalWorkingDaysGet

> ResponseListExceptionalWkgDayResponseAPIResponse OfficesOfficeIdExceptionalWorkingDaysGet(ctx, officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get All Exceptional Working Days



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
	officeId := "officeId_example" // string | Information about adding new exceptional working day details (example:900003)
	type_ := "type__example" // string | Information about adding new exceptional working day details (example:office)
	fromDate := "fromDate_example" // string | Information about adding new exceptional working day details (example:2024-01-01) (optional)
	toDate := "toDate_example" // string | Information about adding new exceptional working day details (example:2024-01-31) (optional)
	skip := int32(56) // int32 | Information about adding new exceptional working day details (example:0) (optional)
	limit := int32(56) // int32 | Information about adding new exceptional working day details (example:10) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysGet(context.Background(), officeId).Type_(type_).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdExceptionalWorkingDaysGet`: ResponseListExceptionalWkgDayResponseAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Information about adding new exceptional working day details (example:900003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdExceptionalWorkingDaysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Information about adding new exceptional working day details (example:office) | 
 **fromDate** | **string** | Information about adding new exceptional working day details (example:2024-01-01) | 
 **toDate** | **string** | Information about adding new exceptional working day details (example:2024-01-31) | 
 **skip** | **int32** | Information about adding new exceptional working day details (example:0) | 
 **limit** | **int32** | Information about adding new exceptional working day details (example:10) | 

### Return type

[**ResponseListExceptionalWkgDayResponseAPIResponse**](ResponseListExceptionalWkgDayResponseAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdExceptionalWorkingDaysPost

> ResponseCreateExceptionalWkgDayResponseAPIResponse OfficesOfficeIdExceptionalWorkingDaysPost(ctx, officeId).Body(body).Execute()

Create Exceptional Working Day



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
	officeId := int32(56) // int32 | Information about adding new exceptional working day details (example:900003)
	body := *openapiclient.NewHandlerCreateExceptionalWkorkingDaysRequest(int32(10130000), "2023-11-30T00:00:00Z", "Exam material booking") // HandlerCreateExceptionalWkorkingDaysRequest | Information about adding new exceptional working day details (example:900003)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdExceptionalWorkingDaysPost`: ResponseCreateExceptionalWkgDayResponseAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureExceptionalWorkingDayAPI.OfficesOfficeIdExceptionalWorkingDaysPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Information about adding new exceptional working day details (example:900003) | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdExceptionalWorkingDaysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerCreateExceptionalWkorkingDaysRequest**](HandlerCreateExceptionalWkorkingDaysRequest.md) | Information about adding new exceptional working day details (example:900003) | 

### Return type

[**ResponseCreateExceptionalWkgDayResponseAPIResponse**](ResponseCreateExceptionalWkgDayResponseAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


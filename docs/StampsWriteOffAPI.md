# \StampsWriteOffAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdStampsWriteOffGet**](StampsWriteOffAPI.md#OfficesOfficeIdStampsWriteOffGet) | **Get** /offices/{office-id}/stamps-write-off | Get Pending WriteOff Transactions List
[**OfficesOfficeIdStampsWriteOffPost**](StampsWriteOffAPI.md#OfficesOfficeIdStampsWriteOffPost) | **Post** /offices/{office-id}/stamps-write-off | Create a New Stamp WriteOff Transaction
[**StampsWriteoffApprovePut**](StampsWriteOffAPI.md#StampsWriteoffApprovePut) | **Put** /stamps/writeoff/approve | Approve WriteOff Transaction Request



## OfficesOfficeIdStampsWriteOffGet

> ResponseListStampsSoiledApiResponse OfficesOfficeIdStampsWriteOffGet(ctx, officeId).IsPdg(isPdg).CheckerOrApproval(checkerOrApproval).FromDate(fromDate).ToDate(toDate).Execute()

Get Pending WriteOff Transactions List



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
	officeId := "officeId_example" // string | Office ID
	isPdg := true // bool | Is PDG (optional)
	checkerOrApproval := "checkerOrApproval_example" // string | Checker or Approval (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffGet(context.Background(), officeId).IsPdg(isPdg).CheckerOrApproval(checkerOrApproval).FromDate(fromDate).ToDate(toDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdStampsWriteOffGet`: ResponseListStampsSoiledApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdStampsWriteOffGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isPdg** | **bool** | Is PDG | 
 **checkerOrApproval** | **string** | Checker or Approval | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 

### Return type

[**ResponseListStampsSoiledApiResponse**](ResponseListStampsSoiledApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficesOfficeIdStampsWriteOffPost

> ResponseCreateStampsSoiledApiResponse OfficesOfficeIdStampsWriteOffPost(ctx, officeId).Body(body).Execute()

Create a New Stamp WriteOff Transaction



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
	officeId := int32(56) // int32 | Office ID
	body := "body_example" // string | Request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffPost(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdStampsWriteOffPost`: ResponseCreateStampsSoiledApiResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsWriteOffAPI.OfficesOfficeIdStampsWriteOffPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdStampsWriteOffPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Request body | 

### Return type

[**ResponseCreateStampsSoiledApiResponse**](ResponseCreateStampsSoiledApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StampsWriteoffApprovePut

> HandlerStampsSoiledResponse StampsWriteoffApprovePut(ctx).Body(body).Execute()

Approve WriteOff Transaction Request



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
	body := *openapiclient.NewHandlerUpdateStampSoiledRequest(int32(10145824), openapiclient.handler.stampStatus("all")) // HandlerUpdateStampSoiledRequest | Update or approving write off Transaction details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StampsWriteOffAPI.StampsWriteoffApprovePut(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StampsWriteOffAPI.StampsWriteoffApprovePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StampsWriteoffApprovePut`: HandlerStampsSoiledResponse
	fmt.Fprintf(os.Stdout, "Response from `StampsWriteOffAPI.StampsWriteoffApprovePut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStampsWriteoffApprovePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerUpdateStampSoiledRequest**](HandlerUpdateStampSoiledRequest.md) | Update or approving write off Transaction details | 

### Return type

[**HandlerStampsSoiledResponse**](HandlerStampsSoiledResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


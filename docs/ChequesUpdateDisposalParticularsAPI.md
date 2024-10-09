# \ChequesUpdateDisposalParticularsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OfficesOfficeIdChequesDisposalPut**](ChequesUpdateDisposalParticularsAPI.md#OfficesOfficeIdChequesDisposalPut) | **Put** /offices/{office-id}/cheques-disposal | Update cheque disposal particulars



## OfficesOfficeIdChequesDisposalPut

> ResponseUpdateChequesDisposalApiResponse OfficesOfficeIdChequesDisposalPut(ctx, officeId).Body(body).Execute()

Update cheque disposal particulars



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
	body := *openapiclient.NewHandlerUpdateChequesDisposalRequest(float32(5000.9), int32(10112222), "TX00000933455") // HandlerUpdateChequesDisposalRequest | Disposed By

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChequesUpdateDisposalParticularsAPI.OfficesOfficeIdChequesDisposalPut(context.Background(), officeId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChequesUpdateDisposalParticularsAPI.OfficesOfficeIdChequesDisposalPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OfficesOfficeIdChequesDisposalPut`: ResponseUpdateChequesDisposalApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ChequesUpdateDisposalParticularsAPI.OfficesOfficeIdChequesDisposalPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **int32** | Office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficesOfficeIdChequesDisposalPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandlerUpdateChequesDisposalRequest**](HandlerUpdateChequesDisposalRequest.md) | Disposed By | 

### Return type

[**ResponseUpdateChequesDisposalApiResponse**](ResponseUpdateChequesDisposalApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


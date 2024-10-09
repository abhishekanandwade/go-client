# \ConfigureCurrencyDenominationAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrenciesCurrencyIdDelete**](ConfigureCurrencyDenominationAPI.md#CurrenciesCurrencyIdDelete) | **Delete** /currencies/{currency-id} | Update Currency Denomination
[**CurrenciesCurrencyIdGet**](ConfigureCurrencyDenominationAPI.md#CurrenciesCurrencyIdGet) | **Get** /currencies/{currency-id} | Get Specific Currency Denomination
[**CurrenciesGet**](ConfigureCurrencyDenominationAPI.md#CurrenciesGet) | **Get** /currencies | Get All Currency Denominations
[**CurrenciesPost**](ConfigureCurrencyDenominationAPI.md#CurrenciesPost) | **Post** /currencies | Create Currency Denomination



## CurrenciesCurrencyIdDelete

> ResponseCurrenciesAPIResponse CurrenciesCurrencyIdDelete(ctx, currencyId).UserId(userId).ValidTo(validTo).Execute()

Update Currency Denomination



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
	currencyId := "currencyId_example" // string | Details identified by currencyId
	userId := "userId_example" // string | User who updated the currency denomination
	validTo := "validTo_example" // string | Valid to date

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdDelete(context.Background(), currencyId).UserId(userId).ValidTo(validTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrenciesCurrencyIdDelete`: ResponseCurrenciesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Details identified by currencyId | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesCurrencyIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User who updated the currency denomination | 
 **validTo** | **string** | Valid to date | 

### Return type

[**ResponseCurrenciesAPIResponse**](ResponseCurrenciesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrenciesCurrencyIdGet

> ResponseCurrenciesAPIResponse CurrenciesCurrencyIdGet(ctx, currencyId).Execute()

Get Specific Currency Denomination



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
	currencyId := "currencyId_example" // string | Get currencey denomination by ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdGet(context.Background(), currencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrenciesCurrencyIdGet`: ResponseCurrenciesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCurrencyDenominationAPI.CurrenciesCurrencyIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currencyId** | **string** | Get currencey denomination by ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesCurrencyIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResponseCurrenciesAPIResponse**](ResponseCurrenciesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrenciesGet

> ResponseListCurrenciesAPIResponse CurrenciesGet(ctx).Skip(skip).Limit(limit).Execute()

Get All Currency Denominations



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
	resp, r, err := apiClient.ConfigureCurrencyDenominationAPI.CurrenciesGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCurrencyDenominationAPI.CurrenciesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrenciesGet`: ResponseListCurrenciesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCurrencyDenominationAPI.CurrenciesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip for pagination | [default to 0]
 **limit** | **int32** | Number of records to limit for pagination | [default to 10]

### Return type

[**ResponseListCurrenciesAPIResponse**](ResponseListCurrenciesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CurrenciesPost

> ResponseCreateCurrenciesAPIResponse CurrenciesPost(ctx).Body(body).Execute()

Create Currency Denomination



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
	body := *openapiclient.NewHandlerCreateCurrenciesRequest(float32(10.0), "10130000", "9999-12-31T00:00:00Z") // HandlerCreateCurrenciesRequest | Information about adding a currency denomination

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureCurrencyDenominationAPI.CurrenciesPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureCurrencyDenominationAPI.CurrenciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CurrenciesPost`: ResponseCreateCurrenciesAPIResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigureCurrencyDenominationAPI.CurrenciesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCurrenciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateCurrenciesRequest**](HandlerCreateCurrenciesRequest.md) | Information about adding a currency denomination | 

### Return type

[**ResponseCreateCurrenciesAPIResponse**](ResponseCreateCurrenciesAPIResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


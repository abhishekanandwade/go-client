# \CashChequeTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashBagsGet**](CashChequeTransactionsAPI.md#CashBagsGet) | **Get** /cash-bags | List Pending Cash Bag Transaction
[**CashBagsPost**](CashChequeTransactionsAPI.md#CashBagsPost) | **Post** /cash-bags | Create a New Cash Bag
[**TransactionsBankRemittancesGet**](CashChequeTransactionsAPI.md#TransactionsBankRemittancesGet) | **Get** /transactions/bank-remittances | List Bank Remittances
[**TransactionsGet**](CashChequeTransactionsAPI.md#TransactionsGet) | **Get** /transactions | Get Treasury Transactions List
[**TransactionsPdgProcessOrAckGet**](CashChequeTransactionsAPI.md#TransactionsPdgProcessOrAckGet) | **Get** /transactions/pdg-process-or-ack | Get Pending receipt ack Transactions List
[**TransactionsPendingProcessGet**](CashChequeTransactionsAPI.md#TransactionsPendingProcessGet) | **Get** /transactions/pending-process | Get Pending Treasury Transactions List
[**TransactionsPost**](CashChequeTransactionsAPI.md#TransactionsPost) | **Post** /transactions | Create a New Treasury Transaction
[**TransactionsTransactionIdChangeStatusPut**](CashChequeTransactionsAPI.md#TransactionsTransactionIdChangeStatusPut) | **Put** /transactions/{transaction-id}/change-status | Process Treasury Transaction Issue Request



## CashBagsGet

> ResponseListCashBagsApiResponse CashBagsGet(ctx).FromOfficeId(fromOfficeId).ToOfficeId(toOfficeId).Skip(skip).Limit(limit).Execute()

List Pending Cash Bag Transaction



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
	fromOfficeId := "fromOfficeId_example" // string | From Office ID
	toOfficeId := "toOfficeId_example" // string | To Office ID
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.CashBagsGet(context.Background()).FromOfficeId(fromOfficeId).ToOfficeId(toOfficeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.CashBagsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashBagsGet`: ResponseListCashBagsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.CashBagsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashBagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromOfficeId** | **string** | From Office ID | 
 **toOfficeId** | **string** | To Office ID | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListCashBagsApiResponse**](ResponseListCashBagsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CashBagsPost

> ResponseCreateCashBagApiResponse CashBagsPost(ctx).Body(body).Execute()

Create a New Cash Bag



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
	body := *openapiclient.NewHandlerCreateCashBagsRequest(float32(18.9), int32(9000001), int32(10123333), int32(9000003), "TX0001, SX00003, IPOTX00004") // HandlerCreateCashBagsRequest | Information about the new cash bag transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.CashBagsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.CashBagsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashBagsPost`: ResponseCreateCashBagApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.CashBagsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashBagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateCashBagsRequest**](HandlerCreateCashBagsRequest.md) | Information about the new cash bag transaction | 

### Return type

[**ResponseCreateCashBagApiResponse**](ResponseCreateCashBagApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsBankRemittancesGet

> ResponseListBankRemittancesApiResponse TransactionsBankRemittancesGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

List Bank Remittances



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
	officeId := "officeId_example" // string | List bank remittances
	fromDate := "fromDate_example" // string | From Date
	toDate := "toDate_example" // string | To Date
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsBankRemittancesGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsBankRemittancesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsBankRemittancesGet`: ResponseListBankRemittancesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsBankRemittancesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsBankRemittancesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | List bank remittances | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListBankRemittancesApiResponse**](ResponseListBankRemittancesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> ResponseListTransactionsApiResponse TransactionsGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()

Get Treasury Transactions List



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
	fromDate := "fromDate_example" // string | From Date
	toDate := "toDate_example" // string | To Date
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsGet`: ResponseListTransactionsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListTransactionsApiResponse**](ResponseListTransactionsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsPdgProcessOrAckGet

> ResponseListTransactionsPdgProcessOrAckApiResponse TransactionsPdgProcessOrAckGet(ctx).OfficeId(officeId).Skip(skip).Limit(limit).Execute()

Get Pending receipt ack Transactions List



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
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsPdgProcessOrAckGet(context.Background()).OfficeId(officeId).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsPdgProcessOrAckGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsPdgProcessOrAckGet`: ResponseListTransactionsPdgProcessOrAckApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsPdgProcessOrAckGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsPdgProcessOrAckGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListTransactionsPdgProcessOrAckApiResponse**](ResponseListTransactionsPdgProcessOrAckApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsPendingProcessGet

> ResponseListTransactionsPendingProcessApiResponse TransactionsPendingProcessGet(ctx).OfficeId(officeId).TxnStatus(txnStatus).IsPostman(isPostman).Skip(skip).Limit(limit).Execute()

Get Pending Treasury Transactions List



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
	txnStatus := "txnStatus_example" // string | Transaction Status
	isPostman := true // bool | Is Postman (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsPendingProcessGet(context.Background()).OfficeId(officeId).TxnStatus(txnStatus).IsPostman(isPostman).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsPendingProcessGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsPendingProcessGet`: ResponseListTransactionsPendingProcessApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsPendingProcessGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsPendingProcessGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **txnStatus** | **string** | Transaction Status | 
 **isPostman** | **bool** | Is Postman | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 

### Return type

[**ResponseListTransactionsPendingProcessApiResponse**](ResponseListTransactionsPendingProcessApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsPost

> ResponseCreateTransactionsBankRemittancesApiResponse TransactionsPost(ctx).Body(body).Execute()

Create a New Treasury Transaction



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
	body := *openapiclient.NewHandlerCreateTransactionsBankRemittancesRequest(int32(9000003), float32(25000.55), int32(9000003), int32(10130000), "Treasury/PoS/DARPAN", "Cash/ Cheque/ DD") // HandlerCreateTransactionsBankRemittancesRequest | Information about creating a new Treasury Transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsPost`: ResponseCreateTransactionsBankRemittancesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HandlerCreateTransactionsBankRemittancesRequest**](HandlerCreateTransactionsBankRemittancesRequest.md) | Information about creating a new Treasury Transaction | 

### Return type

[**ResponseCreateTransactionsBankRemittancesApiResponse**](ResponseCreateTransactionsBankRemittancesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsTransactionIdChangeStatusPut

> ResponseTransactionsBankRemittancesApiResponse TransactionsTransactionIdChangeStatusPut(ctx, transactionId).Type_(type_).Body(body).Execute()

Process Treasury Transaction Issue Request



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
	transactionId := "transactionId_example" // string | Transaction ID
	type_ := "type__example" // string | Type of transaction
	body := *openapiclient.NewHandlerUpdateTransactionsStatusRequest("Request/ Remittance", int32(10130000)) // HandlerUpdateTransactionsStatusRequest | Information about processing issue request treasury Transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsTransactionIdChangeStatusPut(context.Background(), transactionId).Type_(type_).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsTransactionIdChangeStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsTransactionIdChangeStatusPut`: ResponseTransactionsBankRemittancesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsTransactionIdChangeStatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdChangeStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of transaction | 
 **body** | [**HandlerUpdateTransactionsStatusRequest**](HandlerUpdateTransactionsStatusRequest.md) | Information about processing issue request treasury Transaction | 

### Return type

[**ResponseTransactionsBankRemittancesApiResponse**](ResponseTransactionsBankRemittancesApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \CashChequeTransactionsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CashBagsPost**](CashChequeTransactionsAPI.md#CashBagsPost) | **Post** /cash-bags | Create a New Cash Bag
[**CashBagsPut**](CashChequeTransactionsAPI.md#CashBagsPut) | **Put** /cash-bags | List Pending Cash Bag Transaction
[**TransactionsBankRemittancesGet**](CashChequeTransactionsAPI.md#TransactionsBankRemittancesGet) | **Get** /transactions/bank-remittances | List Bank Remittances
[**TransactionsGet**](CashChequeTransactionsAPI.md#TransactionsGet) | **Get** /transactions | Get Treasury Transactions List
[**TransactionsPdgProcessOrAckGet**](CashChequeTransactionsAPI.md#TransactionsPdgProcessOrAckGet) | **Get** /transactions/pdg-process-or-ack | Get Pending receipt ack Transactions List
[**TransactionsPendingProcessGet**](CashChequeTransactionsAPI.md#TransactionsPendingProcessGet) | **Get** /transactions/pending-process | Get Pending Treasury Transactions List
[**TransactionsPost**](CashChequeTransactionsAPI.md#TransactionsPost) | **Post** /transactions | Create a New Treasury Transaction
[**TransactionsTransactionIdStatusPut**](CashChequeTransactionsAPI.md#TransactionsTransactionIdStatusPut) | **Put** /transactions/{transaction-id}/status | Process Treasury Transaction Issue Request



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


## CashBagsPut

> ResponseListCashBagsApiResponse CashBagsPut(ctx).FromOfficeId(fromOfficeId).ToOfficeId(toOfficeId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.CashBagsPut(context.Background()).FromOfficeId(fromOfficeId).ToOfficeId(toOfficeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.CashBagsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CashBagsPut`: ResponseListCashBagsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.CashBagsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCashBagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromOfficeId** | **string** | From Office ID | 
 **toOfficeId** | **string** | To Office ID | 

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


## TransactionsBankRemittancesGet

> ResponseListBankRemittancesApiResponse TransactionsBankRemittancesGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsBankRemittancesGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Execute()
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

> ResponseListTransactionsApiResponse TransactionsGet(ctx).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsGet(context.Background()).OfficeId(officeId).FromDate(fromDate).ToDate(toDate).Execute()
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

> ResponseListTransactionsPdgProcessOrAckApiResponse TransactionsPdgProcessOrAckGet(ctx).OfficeId(officeId).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsPdgProcessOrAckGet(context.Background()).OfficeId(officeId).Execute()
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

> ResponseListTransactionsPendingProcessApiResponse TransactionsPendingProcessGet(ctx).OfficeId(officeId).TxnStatus(txnStatus).IsPostman(isPostman).Execute()

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsPendingProcessGet(context.Background()).OfficeId(officeId).TxnStatus(txnStatus).IsPostman(isPostman).Execute()
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


## TransactionsTransactionIdStatusPut

> ResponseTransactionsBankRemittancesApiResponse TransactionsTransactionIdStatusPut(ctx, transactionId).Body(body).Execute()

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
	body := *openapiclient.NewHandlerUpdateTransactionsStatusRequest("Request/ Remittance", openapiclient.handler.transactionType("approve-pdg-source"), int32(10130000)) // HandlerUpdateTransactionsStatusRequest | Information about processing issue request treasury Transaction

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CashChequeTransactionsAPI.TransactionsTransactionIdStatusPut(context.Background(), transactionId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CashChequeTransactionsAPI.TransactionsTransactionIdStatusPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsTransactionIdStatusPut`: ResponseTransactionsBankRemittancesApiResponse
	fmt.Fprintf(os.Stdout, "Response from `CashChequeTransactionsAPI.TransactionsTransactionIdStatusPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsTransactionIdStatusPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


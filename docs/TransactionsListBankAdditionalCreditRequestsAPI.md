# \TransactionsListBankAdditionalCreditRequestsAPI

All URIs are relative to *http://localhost:8080/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionsBankAddlCreditGet**](TransactionsListBankAdditionalCreditRequestsAPI.md#TransactionsBankAddlCreditGet) | **Get** /transactions/bank-addl-credit | Get All bank additonal credit requests



## TransactionsBankAddlCreditGet

> ResponseListBankAddlCreditTxnsApiResponse TransactionsBankAddlCreditGet(ctx).OfficeId(officeId).IsDivision(isDivision).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).OrderBy(orderBy).SortType(sortType).Execute()

Get All bank additonal credit requests



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
	isDivision := true // bool | Is Division (optional)
	fromDate := "fromDate_example" // string | From Date (optional)
	toDate := "toDate_example" // string | To Date (optional)
	skip := int32(56) // int32 | Skip (optional)
	limit := int32(56) // int32 | Limit (optional)
	orderBy := "orderBy_example" // string | Order By (optional)
	sortType := "sortType_example" // string | Sort Type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionsListBankAdditionalCreditRequestsAPI.TransactionsBankAddlCreditGet(context.Background()).OfficeId(officeId).IsDivision(isDivision).FromDate(fromDate).ToDate(toDate).Skip(skip).Limit(limit).OrderBy(orderBy).SortType(sortType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsListBankAdditionalCreditRequestsAPI.TransactionsBankAddlCreditGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionsBankAddlCreditGet`: ResponseListBankAddlCreditTxnsApiResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsListBankAdditionalCreditRequestsAPI.TransactionsBankAddlCreditGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsBankAddlCreditGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **officeId** | **string** | Office ID | 
 **isDivision** | **bool** | Is Division | 
 **fromDate** | **string** | From Date | 
 **toDate** | **string** | To Date | 
 **skip** | **int32** | Skip | 
 **limit** | **int32** | Limit | 
 **orderBy** | **string** | Order By | 
 **sortType** | **string** | Sort Type | 

### Return type

[**ResponseListBankAddlCreditTxnsApiResponse**](ResponseListBankAddlCreditTxnsApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


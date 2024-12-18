/*
Treasury Management API

A comprehensive API for managing addresses, offering endpoints for creation, update, deletion, and retrieval of address data.

API version: 1.0
Contact: support_cept@indiapost.gov.in
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// TransactionsListBankAdditionalCreditRequestsAPIService TransactionsListBankAdditionalCreditRequestsAPI service
type TransactionsListBankAdditionalCreditRequestsAPIService service

type ApiTransactionsBankAddlCreditGetRequest struct {
	ctx context.Context
	ApiService *TransactionsListBankAdditionalCreditRequestsAPIService
	officeId *string
	isDivision *bool
	fromDate *string
	toDate *string
	skip *int32
	limit *int32
	orderBy *string
	sortType *string
}

// Office ID
func (r ApiTransactionsBankAddlCreditGetRequest) OfficeId(officeId string) ApiTransactionsBankAddlCreditGetRequest {
	r.officeId = &officeId
	return r
}

// Is Division
func (r ApiTransactionsBankAddlCreditGetRequest) IsDivision(isDivision bool) ApiTransactionsBankAddlCreditGetRequest {
	r.isDivision = &isDivision
	return r
}

// From Date
func (r ApiTransactionsBankAddlCreditGetRequest) FromDate(fromDate string) ApiTransactionsBankAddlCreditGetRequest {
	r.fromDate = &fromDate
	return r
}

// To Date
func (r ApiTransactionsBankAddlCreditGetRequest) ToDate(toDate string) ApiTransactionsBankAddlCreditGetRequest {
	r.toDate = &toDate
	return r
}

// Skip
func (r ApiTransactionsBankAddlCreditGetRequest) Skip(skip int32) ApiTransactionsBankAddlCreditGetRequest {
	r.skip = &skip
	return r
}

// Limit
func (r ApiTransactionsBankAddlCreditGetRequest) Limit(limit int32) ApiTransactionsBankAddlCreditGetRequest {
	r.limit = &limit
	return r
}

// Order By
func (r ApiTransactionsBankAddlCreditGetRequest) OrderBy(orderBy string) ApiTransactionsBankAddlCreditGetRequest {
	r.orderBy = &orderBy
	return r
}

// Sort Type
func (r ApiTransactionsBankAddlCreditGetRequest) SortType(sortType string) ApiTransactionsBankAddlCreditGetRequest {
	r.sortType = &sortType
	return r
}

func (r ApiTransactionsBankAddlCreditGetRequest) Execute() (*ResponseListBankAddlCreditTxnsApiResponse, *http.Response, error) {
	return r.ApiService.TransactionsBankAddlCreditGetExecute(r)
}

/*
TransactionsBankAddlCreditGet Get All bank additonal credit requests

Retrieve a list of all bank additional credit requests

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTransactionsBankAddlCreditGetRequest
*/
func (a *TransactionsListBankAdditionalCreditRequestsAPIService) TransactionsBankAddlCreditGet(ctx context.Context) ApiTransactionsBankAddlCreditGetRequest {
	return ApiTransactionsBankAddlCreditGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResponseListBankAddlCreditTxnsApiResponse
func (a *TransactionsListBankAdditionalCreditRequestsAPIService) TransactionsBankAddlCreditGetExecute(r ApiTransactionsBankAddlCreditGetRequest) (*ResponseListBankAddlCreditTxnsApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseListBankAddlCreditTxnsApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsListBankAdditionalCreditRequestsAPIService.TransactionsBankAddlCreditGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/transactions/bank-addl-credit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.officeId == nil {
		return localVarReturnValue, nil, reportError("officeId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "office-id", r.officeId, "", "")
	if r.isDivision != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is-division", r.isDivision, "", "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from-date", r.fromDate, "", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to-date", r.toDate, "", "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "", "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortType", r.sortType, "", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApierrorsAPIErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

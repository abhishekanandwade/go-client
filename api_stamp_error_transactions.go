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
	"strings"
)


// StampErrorTransactionsAPIService StampErrorTransactionsAPI service
type StampErrorTransactionsAPIService service

type ApiStampsErrorTransactionsErrorIdPutRequest struct {
	ctx context.Context
	ApiService *StampErrorTransactionsAPIService
	errorId string
	body *HandlerUpdateStampErrorTransactionsRequest
}

// Update Stamp Error transactions
func (r ApiStampsErrorTransactionsErrorIdPutRequest) Body(body HandlerUpdateStampErrorTransactionsRequest) ApiStampsErrorTransactionsErrorIdPutRequest {
	r.body = &body
	return r
}

func (r ApiStampsErrorTransactionsErrorIdPutRequest) Execute() (*ResponseUpdateStampsErrorTransactionsApiResponse, *http.Response, error) {
	return r.ApiService.StampsErrorTransactionsErrorIdPutExecute(r)
}

/*
StampsErrorTransactionsErrorIdPut Update stamp error Transactions

Update stamp Error Transactions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param errorId Error ID
 @return ApiStampsErrorTransactionsErrorIdPutRequest
*/
func (a *StampErrorTransactionsAPIService) StampsErrorTransactionsErrorIdPut(ctx context.Context, errorId string) ApiStampsErrorTransactionsErrorIdPutRequest {
	return ApiStampsErrorTransactionsErrorIdPutRequest{
		ApiService: a,
		ctx: ctx,
		errorId: errorId,
	}
}

// Execute executes the request
//  @return ResponseUpdateStampsErrorTransactionsApiResponse
func (a *StampErrorTransactionsAPIService) StampsErrorTransactionsErrorIdPutExecute(r ApiStampsErrorTransactionsErrorIdPutRequest) (*ResponseUpdateStampsErrorTransactionsApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseUpdateStampsErrorTransactionsApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StampErrorTransactionsAPIService.StampsErrorTransactionsErrorIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stamps-error-transactions/{error-id}"
	localVarPath = strings.Replace(localVarPath, "{"+"error-id"+"}", url.PathEscape(parameterValueToString(r.errorId, "errorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
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

type ApiStampsErrorTransactionsGetRequest struct {
	ctx context.Context
	ApiService *StampErrorTransactionsAPIService
	officeId *string
	type_ *string
	fromDate *string
	toDate *string
	isRaisingOffice *bool
	skip *int32
	limit *int32
}

// Office ID
func (r ApiStampsErrorTransactionsGetRequest) OfficeId(officeId string) ApiStampsErrorTransactionsGetRequest {
	r.officeId = &officeId
	return r
}

// Type
func (r ApiStampsErrorTransactionsGetRequest) Type_(type_ string) ApiStampsErrorTransactionsGetRequest {
	r.type_ = &type_
	return r
}

// From Date
func (r ApiStampsErrorTransactionsGetRequest) FromDate(fromDate string) ApiStampsErrorTransactionsGetRequest {
	r.fromDate = &fromDate
	return r
}

// To Date
func (r ApiStampsErrorTransactionsGetRequest) ToDate(toDate string) ApiStampsErrorTransactionsGetRequest {
	r.toDate = &toDate
	return r
}

// Is Raising Office
func (r ApiStampsErrorTransactionsGetRequest) IsRaisingOffice(isRaisingOffice bool) ApiStampsErrorTransactionsGetRequest {
	r.isRaisingOffice = &isRaisingOffice
	return r
}

// Skip
func (r ApiStampsErrorTransactionsGetRequest) Skip(skip int32) ApiStampsErrorTransactionsGetRequest {
	r.skip = &skip
	return r
}

// Limit
func (r ApiStampsErrorTransactionsGetRequest) Limit(limit int32) ApiStampsErrorTransactionsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiStampsErrorTransactionsGetRequest) Execute() (*ResponseListStampsErrorTransactionsApiRespnse, *http.Response, error) {
	return r.ApiService.StampsErrorTransactionsGetExecute(r)
}

/*
StampsErrorTransactionsGet Get stamp error Transactions

Get stamp Error Transactions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStampsErrorTransactionsGetRequest
*/
func (a *StampErrorTransactionsAPIService) StampsErrorTransactionsGet(ctx context.Context) ApiStampsErrorTransactionsGetRequest {
	return ApiStampsErrorTransactionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResponseListStampsErrorTransactionsApiRespnse
func (a *StampErrorTransactionsAPIService) StampsErrorTransactionsGetExecute(r ApiStampsErrorTransactionsGetRequest) (*ResponseListStampsErrorTransactionsApiRespnse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseListStampsErrorTransactionsApiRespnse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StampErrorTransactionsAPIService.StampsErrorTransactionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stamps-error-transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.officeId == nil {
		return localVarReturnValue, nil, reportError("officeId is required and must be specified")
	}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "office-id", r.officeId, "", "")
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from-date", r.fromDate, "", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to-date", r.toDate, "", "")
	}
	if r.isRaisingOffice != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is-raising-office", r.isRaisingOffice, "", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "", "")
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
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

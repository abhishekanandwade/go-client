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


// ChequesForDisposalAPIService ChequesForDisposalAPI service
type ChequesForDisposalAPIService service

type ApiOfficesOfficeIdChequesGetRequest struct {
	ctx context.Context
	ApiService *ChequesForDisposalAPIService
	officeId string
	type_ *string
	fromDate *string
	toDate *string
	reportDate *string
	skip *int32
	limit *int32
}

// Type
func (r ApiOfficesOfficeIdChequesGetRequest) Type_(type_ string) ApiOfficesOfficeIdChequesGetRequest {
	r.type_ = &type_
	return r
}

// From Date
func (r ApiOfficesOfficeIdChequesGetRequest) FromDate(fromDate string) ApiOfficesOfficeIdChequesGetRequest {
	r.fromDate = &fromDate
	return r
}

// To Date
func (r ApiOfficesOfficeIdChequesGetRequest) ToDate(toDate string) ApiOfficesOfficeIdChequesGetRequest {
	r.toDate = &toDate
	return r
}

// Report Date
func (r ApiOfficesOfficeIdChequesGetRequest) ReportDate(reportDate string) ApiOfficesOfficeIdChequesGetRequest {
	r.reportDate = &reportDate
	return r
}

// Skip
func (r ApiOfficesOfficeIdChequesGetRequest) Skip(skip int32) ApiOfficesOfficeIdChequesGetRequest {
	r.skip = &skip
	return r
}

// Limit
func (r ApiOfficesOfficeIdChequesGetRequest) Limit(limit int32) ApiOfficesOfficeIdChequesGetRequest {
	r.limit = &limit
	return r
}

func (r ApiOfficesOfficeIdChequesGetRequest) Execute() (*ResponseListChequesApiResponse, *http.Response, error) {
	return r.ApiService.OfficesOfficeIdChequesGetExecute(r)
}

/*
OfficesOfficeIdChequesGet Fetch cheques for disposal

Fetch cheques for disposal

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param officeId Office ID
 @return ApiOfficesOfficeIdChequesGetRequest
*/
func (a *ChequesForDisposalAPIService) OfficesOfficeIdChequesGet(ctx context.Context, officeId string) ApiOfficesOfficeIdChequesGetRequest {
	return ApiOfficesOfficeIdChequesGetRequest{
		ApiService: a,
		ctx: ctx,
		officeId: officeId,
	}
}

// Execute executes the request
//  @return ResponseListChequesApiResponse
func (a *ChequesForDisposalAPIService) OfficesOfficeIdChequesGetExecute(r ApiOfficesOfficeIdChequesGetRequest) (*ResponseListChequesApiResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseListChequesApiResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChequesForDisposalAPIService.OfficesOfficeIdChequesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/offices/{office-id}/cheques"
	localVarPath = strings.Replace(localVarPath, "{"+"office-id"+"}", url.PathEscape(parameterValueToString(r.officeId, "officeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from-date", r.fromDate, "", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to-date", r.toDate, "", "")
	}
	if r.reportDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "report-date", r.reportDate, "", "")
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

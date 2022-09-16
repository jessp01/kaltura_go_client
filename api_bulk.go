/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// BulkApiService BulkApi service
type BulkApiService service

type ApiBulkAbortRequest struct {
	ctx context.Context
	ApiService *BulkApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiBulkAbortRequest) Ks(ks string) ApiBulkAbortRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiBulkAbortRequest) Format(format int32) ApiBulkAbortRequest {
	r.format = &format
	return r
}

func (r ApiBulkAbortRequest) ClientTag(clientTag string) ApiBulkAbortRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiBulkAbortRequest) PartnerId(partnerId int32) ApiBulkAbortRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiBulkAbortRequest) Body(body AccessControlDeleteRequest) ApiBulkAbortRequest {
	r.body = &body
	return r
}

func (r ApiBulkAbortRequest) Execute() (*KalturaBulkUpload, *http.Response, error) {
	return r.ApiService.BulkAbortExecute(r)
}

/*
BulkAbort Method for BulkAbort

Aborts the bulk upload and all its child jobs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkAbortRequest
*/
func (a *BulkApiService) BulkAbort(ctx context.Context) ApiBulkAbortRequest {
	return ApiBulkAbortRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBulkUpload
func (a *BulkApiService) BulkAbortExecute(r ApiBulkAbortRequest) (*KalturaBulkUpload, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBulkUpload
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkApiService.BulkAbort")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/bulkupload_bulk/action/abort"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.clientTag != nil {
		localVarQueryParams.Add("clientTag", parameterToString(*r.clientTag, ""))
	}
	if r.partnerId != nil {
		localVarQueryParams.Add("partnerId", parameterToString(*r.partnerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ks"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("ks", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiBulkGetRequest struct {
	ctx context.Context
	ApiService *BulkApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiBulkGetRequest) Ks(ks string) ApiBulkGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiBulkGetRequest) Format(format int32) ApiBulkGetRequest {
	r.format = &format
	return r
}

func (r ApiBulkGetRequest) ClientTag(clientTag string) ApiBulkGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiBulkGetRequest) PartnerId(partnerId int32) ApiBulkGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiBulkGetRequest) Body(body AccessControlDeleteRequest) ApiBulkGetRequest {
	r.body = &body
	return r
}

func (r ApiBulkGetRequest) Execute() (*KalturaBulkUpload, *http.Response, error) {
	return r.ApiService.BulkGetExecute(r)
}

/*
BulkGet Method for BulkGet

Get bulk upload batch job by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkGetRequest
*/
func (a *BulkApiService) BulkGet(ctx context.Context) ApiBulkGetRequest {
	return ApiBulkGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBulkUpload
func (a *BulkApiService) BulkGetExecute(r ApiBulkGetRequest) (*KalturaBulkUpload, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBulkUpload
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkApiService.BulkGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/bulkupload_bulk/action/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.clientTag != nil {
		localVarQueryParams.Add("clientTag", parameterToString(*r.clientTag, ""))
	}
	if r.partnerId != nil {
		localVarQueryParams.Add("partnerId", parameterToString(*r.partnerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ks"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("ks", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiBulkListRequest struct {
	ctx context.Context
	ApiService *BulkApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *BulkListRequest
}

func (r ApiBulkListRequest) Ks(ks string) ApiBulkListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiBulkListRequest) Format(format int32) ApiBulkListRequest {
	r.format = &format
	return r
}

func (r ApiBulkListRequest) ClientTag(clientTag string) ApiBulkListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiBulkListRequest) PartnerId(partnerId int32) ApiBulkListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiBulkListRequest) Body(body BulkListRequest) ApiBulkListRequest {
	r.body = &body
	return r
}

func (r ApiBulkListRequest) Execute() (*KalturaBulkUploadListResponse, *http.Response, error) {
	return r.ApiService.BulkListExecute(r)
}

/*
BulkList Method for BulkList

List bulk upload batch jobs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkListRequest
*/
func (a *BulkApiService) BulkList(ctx context.Context) ApiBulkListRequest {
	return ApiBulkListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBulkUploadListResponse
func (a *BulkApiService) BulkListExecute(r ApiBulkListRequest) (*KalturaBulkUploadListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBulkUploadListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkApiService.BulkList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/bulkupload_bulk/action/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.clientTag != nil {
		localVarQueryParams.Add("clientTag", parameterToString(*r.clientTag, ""))
	}
	if r.partnerId != nil {
		localVarQueryParams.Add("partnerId", parameterToString(*r.partnerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ks"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("ks", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiBulkServeRequest struct {
	ctx context.Context
	ApiService *BulkApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiBulkServeRequest) Ks(ks string) ApiBulkServeRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiBulkServeRequest) Format(format int32) ApiBulkServeRequest {
	r.format = &format
	return r
}

func (r ApiBulkServeRequest) ClientTag(clientTag string) ApiBulkServeRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiBulkServeRequest) PartnerId(partnerId int32) ApiBulkServeRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiBulkServeRequest) Body(body AccessControlDeleteRequest) ApiBulkServeRequest {
	r.body = &body
	return r
}

func (r ApiBulkServeRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BulkServeExecute(r)
}

/*
BulkServe Method for BulkServe

serve action returns the original file.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkServeRequest
*/
func (a *BulkApiService) BulkServe(ctx context.Context) ApiBulkServeRequest {
	return ApiBulkServeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *BulkApiService) BulkServeExecute(r ApiBulkServeRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkApiService.BulkServe")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/bulkupload_bulk/action/serve"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.clientTag != nil {
		localVarQueryParams.Add("clientTag", parameterToString(*r.clientTag, ""))
	}
	if r.partnerId != nil {
		localVarQueryParams.Add("partnerId", parameterToString(*r.partnerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ks"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("ks", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiBulkServeLogRequest struct {
	ctx context.Context
	ApiService *BulkApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiBulkServeLogRequest) Ks(ks string) ApiBulkServeLogRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiBulkServeLogRequest) Format(format int32) ApiBulkServeLogRequest {
	r.format = &format
	return r
}

func (r ApiBulkServeLogRequest) ClientTag(clientTag string) ApiBulkServeLogRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiBulkServeLogRequest) PartnerId(partnerId int32) ApiBulkServeLogRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiBulkServeLogRequest) Body(body AccessControlDeleteRequest) ApiBulkServeLogRequest {
	r.body = &body
	return r
}

func (r ApiBulkServeLogRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BulkServeLogExecute(r)
}

/*
BulkServeLog Method for BulkServeLog

serveLog action returns the log file for the bulk-upload job.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBulkServeLogRequest
*/
func (a *BulkApiService) BulkServeLog(ctx context.Context) ApiBulkServeLogRequest {
	return ApiBulkServeLogRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *BulkApiService) BulkServeLogExecute(r ApiBulkServeLogRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkApiService.BulkServeLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/bulkupload_bulk/action/serveLog"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	if r.clientTag != nil {
		localVarQueryParams.Add("clientTag", parameterToString(*r.clientTag, ""))
	}
	if r.partnerId != nil {
		localVarQueryParams.Add("partnerId", parameterToString(*r.partnerId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ks"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("ks", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

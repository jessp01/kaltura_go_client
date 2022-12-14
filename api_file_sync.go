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


// FileSyncApiService FileSyncApi service
type FileSyncApiService service

type ApiFileSyncDeleteLocalFileSyncsRequest struct {
	ctx context.Context
	ApiService *FileSyncApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FileSyncDeleteLocalFileSyncsRequest
}

func (r ApiFileSyncDeleteLocalFileSyncsRequest) Ks(ks string) ApiFileSyncDeleteLocalFileSyncsRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFileSyncDeleteLocalFileSyncsRequest) Format(format int32) ApiFileSyncDeleteLocalFileSyncsRequest {
	r.format = &format
	return r
}

func (r ApiFileSyncDeleteLocalFileSyncsRequest) ClientTag(clientTag string) ApiFileSyncDeleteLocalFileSyncsRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFileSyncDeleteLocalFileSyncsRequest) PartnerId(partnerId int32) ApiFileSyncDeleteLocalFileSyncsRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFileSyncDeleteLocalFileSyncsRequest) Body(body FileSyncDeleteLocalFileSyncsRequest) ApiFileSyncDeleteLocalFileSyncsRequest {
	r.body = &body
	return r
}

func (r ApiFileSyncDeleteLocalFileSyncsRequest) Execute() (*KalturaFileSyncListResponse, *http.Response, error) {
	return r.ApiService.FileSyncDeleteLocalFileSyncsExecute(r)
}

/*
FileSyncDeleteLocalFileSyncs Method for FileSyncDeleteLocalFileSyncs

Delete local file syncs by filter

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFileSyncDeleteLocalFileSyncsRequest
*/
func (a *FileSyncApiService) FileSyncDeleteLocalFileSyncs(ctx context.Context) ApiFileSyncDeleteLocalFileSyncsRequest {
	return ApiFileSyncDeleteLocalFileSyncsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaFileSyncListResponse
func (a *FileSyncApiService) FileSyncDeleteLocalFileSyncsExecute(r ApiFileSyncDeleteLocalFileSyncsRequest) (*KalturaFileSyncListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaFileSyncListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileSyncApiService.FileSyncDeleteLocalFileSyncs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/filesync_filesync/action/deleteLocalFileSyncs"

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

type ApiFileSyncListRequest struct {
	ctx context.Context
	ApiService *FileSyncApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FileSyncListRequest
}

func (r ApiFileSyncListRequest) Ks(ks string) ApiFileSyncListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFileSyncListRequest) Format(format int32) ApiFileSyncListRequest {
	r.format = &format
	return r
}

func (r ApiFileSyncListRequest) ClientTag(clientTag string) ApiFileSyncListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFileSyncListRequest) PartnerId(partnerId int32) ApiFileSyncListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFileSyncListRequest) Body(body FileSyncListRequest) ApiFileSyncListRequest {
	r.body = &body
	return r
}

func (r ApiFileSyncListRequest) Execute() (*KalturaFileSyncListResponse, *http.Response, error) {
	return r.ApiService.FileSyncListExecute(r)
}

/*
FileSyncList Method for FileSyncList

List file syce objects by filter and pager

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFileSyncListRequest
*/
func (a *FileSyncApiService) FileSyncList(ctx context.Context) ApiFileSyncListRequest {
	return ApiFileSyncListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaFileSyncListResponse
func (a *FileSyncApiService) FileSyncListExecute(r ApiFileSyncListRequest) (*KalturaFileSyncListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaFileSyncListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileSyncApiService.FileSyncList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/filesync_filesync/action/list"

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

type ApiFileSyncLockFileSyncsRequest struct {
	ctx context.Context
	ApiService *FileSyncApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FileSyncLockFileSyncsRequest
}

func (r ApiFileSyncLockFileSyncsRequest) Ks(ks string) ApiFileSyncLockFileSyncsRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFileSyncLockFileSyncsRequest) Format(format int32) ApiFileSyncLockFileSyncsRequest {
	r.format = &format
	return r
}

func (r ApiFileSyncLockFileSyncsRequest) ClientTag(clientTag string) ApiFileSyncLockFileSyncsRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFileSyncLockFileSyncsRequest) PartnerId(partnerId int32) ApiFileSyncLockFileSyncsRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFileSyncLockFileSyncsRequest) Body(body FileSyncLockFileSyncsRequest) ApiFileSyncLockFileSyncsRequest {
	r.body = &body
	return r
}

func (r ApiFileSyncLockFileSyncsRequest) Execute() (*KalturaLockFileSyncsResponse, *http.Response, error) {
	return r.ApiService.FileSyncLockFileSyncsExecute(r)
}

/*
FileSyncLockFileSyncs Method for FileSyncLockFileSyncs

lockFileSyncs action locks file syncs for the file sync periodic worker

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFileSyncLockFileSyncsRequest
*/
func (a *FileSyncApiService) FileSyncLockFileSyncs(ctx context.Context) ApiFileSyncLockFileSyncsRequest {
	return ApiFileSyncLockFileSyncsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaLockFileSyncsResponse
func (a *FileSyncApiService) FileSyncLockFileSyncsExecute(r ApiFileSyncLockFileSyncsRequest) (*KalturaLockFileSyncsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaLockFileSyncsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileSyncApiService.FileSyncLockFileSyncs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/filesync_filesync/action/lockFileSyncs"

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

type ApiFileSyncUpdateRequest struct {
	ctx context.Context
	ApiService *FileSyncApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FileSyncUpdateRequest
}

func (r ApiFileSyncUpdateRequest) Ks(ks string) ApiFileSyncUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFileSyncUpdateRequest) Format(format int32) ApiFileSyncUpdateRequest {
	r.format = &format
	return r
}

func (r ApiFileSyncUpdateRequest) ClientTag(clientTag string) ApiFileSyncUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFileSyncUpdateRequest) PartnerId(partnerId int32) ApiFileSyncUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFileSyncUpdateRequest) Body(body FileSyncUpdateRequest) ApiFileSyncUpdateRequest {
	r.body = &body
	return r
}

func (r ApiFileSyncUpdateRequest) Execute() (*KalturaFileSync, *http.Response, error) {
	return r.ApiService.FileSyncUpdateExecute(r)
}

/*
FileSyncUpdate Method for FileSyncUpdate

Update file sync by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFileSyncUpdateRequest
*/
func (a *FileSyncApiService) FileSyncUpdate(ctx context.Context) ApiFileSyncUpdateRequest {
	return ApiFileSyncUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaFileSync
func (a *FileSyncApiService) FileSyncUpdateExecute(r ApiFileSyncUpdateRequest) (*KalturaFileSync, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaFileSync
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FileSyncApiService.FileSyncUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/filesync_filesync/action/update"

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

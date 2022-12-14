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


// AppTokenApiService AppTokenApi service
type AppTokenApiService service

type ApiAppTokenAddRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AppTokenAddRequest
}

func (r ApiAppTokenAddRequest) Ks(ks string) ApiAppTokenAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenAddRequest) Format(format int32) ApiAppTokenAddRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenAddRequest) ClientTag(clientTag string) ApiAppTokenAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenAddRequest) PartnerId(partnerId int32) ApiAppTokenAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenAddRequest) Body(body AppTokenAddRequest) ApiAppTokenAddRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenAddRequest) Execute() (*KalturaAppToken, *http.Response, error) {
	return r.ApiService.AppTokenAddExecute(r)
}

/*
AppTokenAdd Method for AppTokenAdd

Add new application authentication token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenAddRequest
*/
func (a *AppTokenApiService) AppTokenAdd(ctx context.Context) ApiAppTokenAddRequest {
	return ApiAppTokenAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAppToken
func (a *AppTokenApiService) AppTokenAddExecute(r ApiAppTokenAddRequest) (*KalturaAppToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAppToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/add"

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

type ApiAppTokenDeleteRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiAppTokenDeleteRequest) Ks(ks string) ApiAppTokenDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenDeleteRequest) Format(format int32) ApiAppTokenDeleteRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenDeleteRequest) ClientTag(clientTag string) ApiAppTokenDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenDeleteRequest) PartnerId(partnerId int32) ApiAppTokenDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenDeleteRequest) Body(body AnnotationDeleteRequest) ApiAppTokenDeleteRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppTokenDeleteExecute(r)
}

/*
AppTokenDelete Method for AppTokenDelete

Delete application authentication token by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenDeleteRequest
*/
func (a *AppTokenApiService) AppTokenDelete(ctx context.Context) ApiAppTokenDeleteRequest {
	return ApiAppTokenDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AppTokenApiService) AppTokenDeleteExecute(r ApiAppTokenDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/delete"

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAppTokenGetRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiAppTokenGetRequest) Ks(ks string) ApiAppTokenGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenGetRequest) Format(format int32) ApiAppTokenGetRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenGetRequest) ClientTag(clientTag string) ApiAppTokenGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenGetRequest) PartnerId(partnerId int32) ApiAppTokenGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenGetRequest) Body(body AnnotationDeleteRequest) ApiAppTokenGetRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenGetRequest) Execute() (*KalturaAppToken, *http.Response, error) {
	return r.ApiService.AppTokenGetExecute(r)
}

/*
AppTokenGet Method for AppTokenGet

Get application authentication token by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenGetRequest
*/
func (a *AppTokenApiService) AppTokenGet(ctx context.Context) ApiAppTokenGetRequest {
	return ApiAppTokenGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAppToken
func (a *AppTokenApiService) AppTokenGetExecute(r ApiAppTokenGetRequest) (*KalturaAppToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAppToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/get"

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

type ApiAppTokenListRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AppTokenListRequest
}

func (r ApiAppTokenListRequest) Ks(ks string) ApiAppTokenListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenListRequest) Format(format int32) ApiAppTokenListRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenListRequest) ClientTag(clientTag string) ApiAppTokenListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenListRequest) PartnerId(partnerId int32) ApiAppTokenListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenListRequest) Body(body AppTokenListRequest) ApiAppTokenListRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenListRequest) Execute() (*KalturaAppTokenListResponse, *http.Response, error) {
	return r.ApiService.AppTokenListExecute(r)
}

/*
AppTokenList Method for AppTokenList

List application authentication tokens by filter and pager

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenListRequest
*/
func (a *AppTokenApiService) AppTokenList(ctx context.Context) ApiAppTokenListRequest {
	return ApiAppTokenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAppTokenListResponse
func (a *AppTokenApiService) AppTokenListExecute(r ApiAppTokenListRequest) (*KalturaAppTokenListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAppTokenListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/list"

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

type ApiAppTokenStartSessionRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AppTokenStartSessionRequest
}

func (r ApiAppTokenStartSessionRequest) Ks(ks string) ApiAppTokenStartSessionRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenStartSessionRequest) Format(format int32) ApiAppTokenStartSessionRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenStartSessionRequest) ClientTag(clientTag string) ApiAppTokenStartSessionRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenStartSessionRequest) PartnerId(partnerId int32) ApiAppTokenStartSessionRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenStartSessionRequest) Body(body AppTokenStartSessionRequest) ApiAppTokenStartSessionRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenStartSessionRequest) Execute() (*KalturaSessionInfo, *http.Response, error) {
	return r.ApiService.AppTokenStartSessionExecute(r)
}

/*
AppTokenStartSession Method for AppTokenStartSession

Starts a new KS (kaltura Session) based on an application authentication token ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenStartSessionRequest
*/
func (a *AppTokenApiService) AppTokenStartSession(ctx context.Context) ApiAppTokenStartSessionRequest {
	return ApiAppTokenStartSessionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaSessionInfo
func (a *AppTokenApiService) AppTokenStartSessionExecute(r ApiAppTokenStartSessionRequest) (*KalturaSessionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaSessionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenStartSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/startSession"

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

type ApiAppTokenUpdateRequest struct {
	ctx context.Context
	ApiService *AppTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AppTokenUpdateRequest
}

func (r ApiAppTokenUpdateRequest) Ks(ks string) ApiAppTokenUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAppTokenUpdateRequest) Format(format int32) ApiAppTokenUpdateRequest {
	r.format = &format
	return r
}

func (r ApiAppTokenUpdateRequest) ClientTag(clientTag string) ApiAppTokenUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAppTokenUpdateRequest) PartnerId(partnerId int32) ApiAppTokenUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAppTokenUpdateRequest) Body(body AppTokenUpdateRequest) ApiAppTokenUpdateRequest {
	r.body = &body
	return r
}

func (r ApiAppTokenUpdateRequest) Execute() (*KalturaAppToken, *http.Response, error) {
	return r.ApiService.AppTokenUpdateExecute(r)
}

/*
AppTokenUpdate Method for AppTokenUpdate

Update application authentication token by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppTokenUpdateRequest
*/
func (a *AppTokenApiService) AppTokenUpdate(ctx context.Context) ApiAppTokenUpdateRequest {
	return ApiAppTokenUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAppToken
func (a *AppTokenApiService) AppTokenUpdateExecute(r ApiAppTokenUpdateRequest) (*KalturaAppToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAppToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppTokenApiService.AppTokenUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/apptoken/action/update"

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

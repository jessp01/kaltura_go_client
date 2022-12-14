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


// SessionApiService SessionApi service
type SessionApiService service

type ApiSessionEndRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

func (r ApiSessionEndRequest) Ks(ks string) ApiSessionEndRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSessionEndRequest) Format(format int32) ApiSessionEndRequest {
	r.format = &format
	return r
}

func (r ApiSessionEndRequest) ClientTag(clientTag string) ApiSessionEndRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionEndRequest) PartnerId(partnerId int32) ApiSessionEndRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionEndRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSessionEndRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSessionEndRequest) Execute() (*http.Response, error) {
	return r.ApiService.SessionEndExecute(r)
}

/*
SessionEnd Method for SessionEnd

End a session with the Kaltura server, making the current KS invalid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionEndRequest
*/
func (a *SessionApiService) SessionEnd(ctx context.Context) ApiSessionEndRequest {
	return ApiSessionEndRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SessionApiService) SessionEndExecute(r ApiSessionEndRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionEnd")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/end"

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
	localVarPostBody = r.batchCleanExclusiveJobsRequest
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

type ApiSessionGetRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SessionGetRequest
}

func (r ApiSessionGetRequest) Ks(ks string) ApiSessionGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSessionGetRequest) Format(format int32) ApiSessionGetRequest {
	r.format = &format
	return r
}

func (r ApiSessionGetRequest) ClientTag(clientTag string) ApiSessionGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionGetRequest) PartnerId(partnerId int32) ApiSessionGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionGetRequest) Body(body SessionGetRequest) ApiSessionGetRequest {
	r.body = &body
	return r
}

func (r ApiSessionGetRequest) Execute() (*KalturaSessionInfo, *http.Response, error) {
	return r.ApiService.SessionGetExecute(r)
}

/*
SessionGet Method for SessionGet

Parse session key and return its info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionGetRequest
*/
func (a *SessionApiService) SessionGet(ctx context.Context) ApiSessionGetRequest {
	return ApiSessionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaSessionInfo
func (a *SessionApiService) SessionGetExecute(r ApiSessionGetRequest) (*KalturaSessionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaSessionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/get"

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

type ApiSessionImpersonateRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	format *int32
	clientTag *string
	partnerId *int32
	body *SessionImpersonateRequest
}

// The format of the response
func (r ApiSessionImpersonateRequest) Format(format int32) ApiSessionImpersonateRequest {
	r.format = &format
	return r
}

func (r ApiSessionImpersonateRequest) ClientTag(clientTag string) ApiSessionImpersonateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionImpersonateRequest) PartnerId(partnerId int32) ApiSessionImpersonateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionImpersonateRequest) Body(body SessionImpersonateRequest) ApiSessionImpersonateRequest {
	r.body = &body
	return r
}

func (r ApiSessionImpersonateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SessionImpersonateExecute(r)
}

/*
SessionImpersonate Method for SessionImpersonate

Start an impersonated session with Kaltura's server.

The result KS is the session key that you should pass to all services that requires a ticket.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionImpersonateRequest
*/
func (a *SessionApiService) SessionImpersonate(ctx context.Context) ApiSessionImpersonateRequest {
	return ApiSessionImpersonateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *SessionApiService) SessionImpersonateExecute(r ApiSessionImpersonateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionImpersonate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/impersonate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSessionImpersonateByKsRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SessionImpersonateByKsRequest
}

func (r ApiSessionImpersonateByKsRequest) Ks(ks string) ApiSessionImpersonateByKsRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSessionImpersonateByKsRequest) Format(format int32) ApiSessionImpersonateByKsRequest {
	r.format = &format
	return r
}

func (r ApiSessionImpersonateByKsRequest) ClientTag(clientTag string) ApiSessionImpersonateByKsRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionImpersonateByKsRequest) PartnerId(partnerId int32) ApiSessionImpersonateByKsRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionImpersonateByKsRequest) Body(body SessionImpersonateByKsRequest) ApiSessionImpersonateByKsRequest {
	r.body = &body
	return r
}

func (r ApiSessionImpersonateByKsRequest) Execute() (*KalturaSessionInfo, *http.Response, error) {
	return r.ApiService.SessionImpersonateByKsExecute(r)
}

/*
SessionImpersonateByKs Method for SessionImpersonateByKs

Start an impersonated session with Kaltura's server.

The result KS info contains the session key that you should pass to all services that requires a ticket.

Type, expiry and privileges won't be changed if they're not set

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionImpersonateByKsRequest
*/
func (a *SessionApiService) SessionImpersonateByKs(ctx context.Context) ApiSessionImpersonateByKsRequest {
	return ApiSessionImpersonateByKsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaSessionInfo
func (a *SessionApiService) SessionImpersonateByKsExecute(r ApiSessionImpersonateByKsRequest) (*KalturaSessionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaSessionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionImpersonateByKs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/impersonateByKs"

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

type ApiSessionStartRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	format *int32
	clientTag *string
	partnerId *int32
	body *SessionStartRequest
}

// The format of the response
func (r ApiSessionStartRequest) Format(format int32) ApiSessionStartRequest {
	r.format = &format
	return r
}

func (r ApiSessionStartRequest) ClientTag(clientTag string) ApiSessionStartRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionStartRequest) PartnerId(partnerId int32) ApiSessionStartRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionStartRequest) Body(body SessionStartRequest) ApiSessionStartRequest {
	r.body = &body
	return r
}

func (r ApiSessionStartRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SessionStartExecute(r)
}

/*
SessionStart Method for SessionStart

Start a session with Kaltura's server.

The result KS is the session key that you should pass to all services that requires a ticket.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionStartRequest
*/
func (a *SessionApiService) SessionStart(ctx context.Context) ApiSessionStartRequest {
	return ApiSessionStartRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *SessionApiService) SessionStartExecute(r ApiSessionStartRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionStart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/start"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiSessionStartWidgetSessionRequest struct {
	ctx context.Context
	ApiService *SessionApiService
	format *int32
	clientTag *string
	partnerId *int32
	body *SessionStartWidgetSessionRequest
}

// The format of the response
func (r ApiSessionStartWidgetSessionRequest) Format(format int32) ApiSessionStartWidgetSessionRequest {
	r.format = &format
	return r
}

func (r ApiSessionStartWidgetSessionRequest) ClientTag(clientTag string) ApiSessionStartWidgetSessionRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSessionStartWidgetSessionRequest) PartnerId(partnerId int32) ApiSessionStartWidgetSessionRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSessionStartWidgetSessionRequest) Body(body SessionStartWidgetSessionRequest) ApiSessionStartWidgetSessionRequest {
	r.body = &body
	return r
}

func (r ApiSessionStartWidgetSessionRequest) Execute() (*KalturaStartWidgetSessionResponse, *http.Response, error) {
	return r.ApiService.SessionStartWidgetSessionExecute(r)
}

/*
SessionStartWidgetSession Method for SessionStartWidgetSession

Start a session for Kaltura's flash widgets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSessionStartWidgetSessionRequest
*/
func (a *SessionApiService) SessionStartWidgetSession(ctx context.Context) ApiSessionStartWidgetSessionRequest {
	return ApiSessionStartWidgetSessionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaStartWidgetSessionResponse
func (a *SessionApiService) SessionStartWidgetSessionExecute(r ApiSessionStartWidgetSessionRequest) (*KalturaStartWidgetSessionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaStartWidgetSessionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SessionApiService.SessionStartWidgetSession")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/session/action/startWidgetSession"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

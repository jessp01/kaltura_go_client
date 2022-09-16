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


// SystemApiService SystemApi service
type SystemApiService service

type ApiSystemGetHealthCheckRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

// The format of the response
func (r ApiSystemGetHealthCheckRequest) Format(format int32) ApiSystemGetHealthCheckRequest {
	r.format = &format
	return r
}

func (r ApiSystemGetHealthCheckRequest) ClientTag(clientTag string) ApiSystemGetHealthCheckRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSystemGetHealthCheckRequest) PartnerId(partnerId int32) ApiSystemGetHealthCheckRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSystemGetHealthCheckRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSystemGetHealthCheckRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSystemGetHealthCheckRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SystemGetHealthCheckExecute(r)
}

/*
SystemGetHealthCheck Method for SystemGetHealthCheck

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemGetHealthCheckRequest
*/
func (a *SystemApiService) SystemGetHealthCheck(ctx context.Context) ApiSystemGetHealthCheckRequest {
	return ApiSystemGetHealthCheckRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *SystemApiService) SystemGetHealthCheckExecute(r ApiSystemGetHealthCheckRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.SystemGetHealthCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/system/action/getHealthCheck"

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

type ApiSystemGetTimeRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

// The format of the response
func (r ApiSystemGetTimeRequest) Format(format int32) ApiSystemGetTimeRequest {
	r.format = &format
	return r
}

func (r ApiSystemGetTimeRequest) ClientTag(clientTag string) ApiSystemGetTimeRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSystemGetTimeRequest) PartnerId(partnerId int32) ApiSystemGetTimeRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSystemGetTimeRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSystemGetTimeRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSystemGetTimeRequest) Execute() (int32, *http.Response, error) {
	return r.ApiService.SystemGetTimeExecute(r)
}

/*
SystemGetTime Method for SystemGetTime

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemGetTimeRequest
*/
func (a *SystemApiService) SystemGetTime(ctx context.Context) ApiSystemGetTimeRequest {
	return ApiSystemGetTimeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return int32
func (a *SystemApiService) SystemGetTimeExecute(r ApiSystemGetTimeRequest) (int32, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.SystemGetTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/system/action/getTime"

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

type ApiSystemGetVersionRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

// The format of the response
func (r ApiSystemGetVersionRequest) Format(format int32) ApiSystemGetVersionRequest {
	r.format = &format
	return r
}

func (r ApiSystemGetVersionRequest) ClientTag(clientTag string) ApiSystemGetVersionRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSystemGetVersionRequest) PartnerId(partnerId int32) ApiSystemGetVersionRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSystemGetVersionRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSystemGetVersionRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSystemGetVersionRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SystemGetVersionExecute(r)
}

/*
SystemGetVersion Method for SystemGetVersion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemGetVersionRequest
*/
func (a *SystemApiService) SystemGetVersion(ctx context.Context) ApiSystemGetVersionRequest {
	return ApiSystemGetVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *SystemApiService) SystemGetVersionExecute(r ApiSystemGetVersionRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.SystemGetVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/system/action/getVersion"

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

type ApiSystemPingRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

// The format of the response
func (r ApiSystemPingRequest) Format(format int32) ApiSystemPingRequest {
	r.format = &format
	return r
}

func (r ApiSystemPingRequest) ClientTag(clientTag string) ApiSystemPingRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSystemPingRequest) PartnerId(partnerId int32) ApiSystemPingRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSystemPingRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSystemPingRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSystemPingRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.SystemPingExecute(r)
}

/*
SystemPing Method for SystemPing

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemPingRequest
*/
func (a *SystemApiService) SystemPing(ctx context.Context) ApiSystemPingRequest {
	return ApiSystemPingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *SystemApiService) SystemPingExecute(r ApiSystemPingRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.SystemPing")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/system/action/ping"

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

type ApiSystemPingDatabaseRequest struct {
	ctx context.Context
	ApiService *SystemApiService
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

// The format of the response
func (r ApiSystemPingDatabaseRequest) Format(format int32) ApiSystemPingDatabaseRequest {
	r.format = &format
	return r
}

func (r ApiSystemPingDatabaseRequest) ClientTag(clientTag string) ApiSystemPingDatabaseRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSystemPingDatabaseRequest) PartnerId(partnerId int32) ApiSystemPingDatabaseRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSystemPingDatabaseRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiSystemPingDatabaseRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiSystemPingDatabaseRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.SystemPingDatabaseExecute(r)
}

/*
SystemPingDatabase Method for SystemPingDatabase

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSystemPingDatabaseRequest
*/
func (a *SystemApiService) SystemPingDatabase(ctx context.Context) ApiSystemPingDatabaseRequest {
	return ApiSystemPingDatabaseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *SystemApiService) SystemPingDatabaseExecute(r ApiSystemPingDatabaseRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SystemApiService.SystemPingDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/system/action/pingDatabase"

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
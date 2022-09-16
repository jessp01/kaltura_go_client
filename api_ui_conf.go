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


// UiConfApiService UiConfApi service
type UiConfApiService service

type ApiUiConfAddRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UiConfAddRequest
}

func (r ApiUiConfAddRequest) Ks(ks string) ApiUiConfAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfAddRequest) Format(format int32) ApiUiConfAddRequest {
	r.format = &format
	return r
}

func (r ApiUiConfAddRequest) ClientTag(clientTag string) ApiUiConfAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfAddRequest) PartnerId(partnerId int32) ApiUiConfAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfAddRequest) Body(body UiConfAddRequest) ApiUiConfAddRequest {
	r.body = &body
	return r
}

func (r ApiUiConfAddRequest) Execute() (*KalturaUiConf, *http.Response, error) {
	return r.ApiService.UiConfAddExecute(r)
}

/*
UiConfAdd Method for UiConfAdd

UIConf Add action allows you to add a UIConf to Kaltura DB

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfAddRequest
*/
func (a *UiConfApiService) UiConfAdd(ctx context.Context) ApiUiConfAddRequest {
	return ApiUiConfAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConf
func (a *UiConfApiService) UiConfAddExecute(r ApiUiConfAddRequest) (*KalturaUiConf, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/add"

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

type ApiUiConfCloneRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiUiConfCloneRequest) Ks(ks string) ApiUiConfCloneRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfCloneRequest) Format(format int32) ApiUiConfCloneRequest {
	r.format = &format
	return r
}

func (r ApiUiConfCloneRequest) ClientTag(clientTag string) ApiUiConfCloneRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfCloneRequest) PartnerId(partnerId int32) ApiUiConfCloneRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfCloneRequest) Body(body AccessControlDeleteRequest) ApiUiConfCloneRequest {
	r.body = &body
	return r
}

func (r ApiUiConfCloneRequest) Execute() (*KalturaUiConf, *http.Response, error) {
	return r.ApiService.UiConfCloneExecute(r)
}

/*
UiConfClone Method for UiConfClone

Clone an existing UIConf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfCloneRequest
*/
func (a *UiConfApiService) UiConfClone(ctx context.Context) ApiUiConfCloneRequest {
	return ApiUiConfCloneRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConf
func (a *UiConfApiService) UiConfCloneExecute(r ApiUiConfCloneRequest) (*KalturaUiConf, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfClone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/clone"

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

type ApiUiConfDeleteRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiUiConfDeleteRequest) Ks(ks string) ApiUiConfDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfDeleteRequest) Format(format int32) ApiUiConfDeleteRequest {
	r.format = &format
	return r
}

func (r ApiUiConfDeleteRequest) ClientTag(clientTag string) ApiUiConfDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfDeleteRequest) PartnerId(partnerId int32) ApiUiConfDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfDeleteRequest) Body(body AccessControlDeleteRequest) ApiUiConfDeleteRequest {
	r.body = &body
	return r
}

func (r ApiUiConfDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.UiConfDeleteExecute(r)
}

/*
UiConfDelete Method for UiConfDelete

Delete an existing UIConf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfDeleteRequest
*/
func (a *UiConfApiService) UiConfDelete(ctx context.Context) ApiUiConfDeleteRequest {
	return ApiUiConfDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UiConfApiService) UiConfDeleteExecute(r ApiUiConfDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/delete"

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

type ApiUiConfGetRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiUiConfGetRequest) Ks(ks string) ApiUiConfGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfGetRequest) Format(format int32) ApiUiConfGetRequest {
	r.format = &format
	return r
}

func (r ApiUiConfGetRequest) ClientTag(clientTag string) ApiUiConfGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfGetRequest) PartnerId(partnerId int32) ApiUiConfGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfGetRequest) Body(body AccessControlDeleteRequest) ApiUiConfGetRequest {
	r.body = &body
	return r
}

func (r ApiUiConfGetRequest) Execute() (*KalturaUiConf, *http.Response, error) {
	return r.ApiService.UiConfGetExecute(r)
}

/*
UiConfGet Method for UiConfGet

Retrieve a UIConf by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfGetRequest
*/
func (a *UiConfApiService) UiConfGet(ctx context.Context) ApiUiConfGetRequest {
	return ApiUiConfGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConf
func (a *UiConfApiService) UiConfGetExecute(r ApiUiConfGetRequest) (*KalturaUiConf, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/get"

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

type ApiUiConfGetAvailableTypesRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

func (r ApiUiConfGetAvailableTypesRequest) Ks(ks string) ApiUiConfGetAvailableTypesRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfGetAvailableTypesRequest) Format(format int32) ApiUiConfGetAvailableTypesRequest {
	r.format = &format
	return r
}

func (r ApiUiConfGetAvailableTypesRequest) ClientTag(clientTag string) ApiUiConfGetAvailableTypesRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfGetAvailableTypesRequest) PartnerId(partnerId int32) ApiUiConfGetAvailableTypesRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfGetAvailableTypesRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiUiConfGetAvailableTypesRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiUiConfGetAvailableTypesRequest) Execute() ([]KalturaUiConfTypeInfo, *http.Response, error) {
	return r.ApiService.UiConfGetAvailableTypesExecute(r)
}

/*
UiConfGetAvailableTypes Method for UiConfGetAvailableTypes

Retrieve a list of all available versions by object type

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfGetAvailableTypesRequest
*/
func (a *UiConfApiService) UiConfGetAvailableTypes(ctx context.Context) ApiUiConfGetAvailableTypesRequest {
	return ApiUiConfGetAvailableTypesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []KalturaUiConfTypeInfo
func (a *UiConfApiService) UiConfGetAvailableTypesExecute(r ApiUiConfGetAvailableTypesRequest) ([]KalturaUiConfTypeInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []KalturaUiConfTypeInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfGetAvailableTypes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/getAvailableTypes"

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

type ApiUiConfListRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UiConfListRequest
}

func (r ApiUiConfListRequest) Ks(ks string) ApiUiConfListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfListRequest) Format(format int32) ApiUiConfListRequest {
	r.format = &format
	return r
}

func (r ApiUiConfListRequest) ClientTag(clientTag string) ApiUiConfListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfListRequest) PartnerId(partnerId int32) ApiUiConfListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfListRequest) Body(body UiConfListRequest) ApiUiConfListRequest {
	r.body = &body
	return r
}

func (r ApiUiConfListRequest) Execute() (*KalturaUiConfListResponse, *http.Response, error) {
	return r.ApiService.UiConfListExecute(r)
}

/*
UiConfList Method for UiConfList

Retrieve a list of available UIConfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfListRequest
*/
func (a *UiConfApiService) UiConfList(ctx context.Context) ApiUiConfListRequest {
	return ApiUiConfListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConfListResponse
func (a *UiConfApiService) UiConfListExecute(r ApiUiConfListRequest) (*KalturaUiConfListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConfListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/list"

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

type ApiUiConfListTemplatesRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UiConfListRequest
}

func (r ApiUiConfListTemplatesRequest) Ks(ks string) ApiUiConfListTemplatesRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfListTemplatesRequest) Format(format int32) ApiUiConfListTemplatesRequest {
	r.format = &format
	return r
}

func (r ApiUiConfListTemplatesRequest) ClientTag(clientTag string) ApiUiConfListTemplatesRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfListTemplatesRequest) PartnerId(partnerId int32) ApiUiConfListTemplatesRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfListTemplatesRequest) Body(body UiConfListRequest) ApiUiConfListTemplatesRequest {
	r.body = &body
	return r
}

func (r ApiUiConfListTemplatesRequest) Execute() (*KalturaUiConfListResponse, *http.Response, error) {
	return r.ApiService.UiConfListTemplatesExecute(r)
}

/*
UiConfListTemplates Method for UiConfListTemplates

retrieve a list of available template UIConfs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfListTemplatesRequest
*/
func (a *UiConfApiService) UiConfListTemplates(ctx context.Context) ApiUiConfListTemplatesRequest {
	return ApiUiConfListTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConfListResponse
func (a *UiConfApiService) UiConfListTemplatesExecute(r ApiUiConfListTemplatesRequest) (*KalturaUiConfListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConfListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfListTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/listTemplates"

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

type ApiUiConfUpdateRequest struct {
	ctx context.Context
	ApiService *UiConfApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UiConfUpdateRequest
}

func (r ApiUiConfUpdateRequest) Ks(ks string) ApiUiConfUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUiConfUpdateRequest) Format(format int32) ApiUiConfUpdateRequest {
	r.format = &format
	return r
}

func (r ApiUiConfUpdateRequest) ClientTag(clientTag string) ApiUiConfUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUiConfUpdateRequest) PartnerId(partnerId int32) ApiUiConfUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUiConfUpdateRequest) Body(body UiConfUpdateRequest) ApiUiConfUpdateRequest {
	r.body = &body
	return r
}

func (r ApiUiConfUpdateRequest) Execute() (*KalturaUiConf, *http.Response, error) {
	return r.ApiService.UiConfUpdateExecute(r)
}

/*
UiConfUpdate Method for UiConfUpdate

Update an existing UIConf

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUiConfUpdateRequest
*/
func (a *UiConfApiService) UiConfUpdate(ctx context.Context) ApiUiConfUpdateRequest {
	return ApiUiConfUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUiConf
func (a *UiConfApiService) UiConfUpdateExecute(r ApiUiConfUpdateRequest) (*KalturaUiConf, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUiConf
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UiConfApiService.UiConfUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uiconf/action/update"

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

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


// EntryServerNodeApiService EntryServerNodeApi service
type EntryServerNodeApiService service

type ApiEntryServerNodeGetRequest struct {
	ctx context.Context
	ApiService *EntryServerNodeApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiEntryServerNodeGetRequest) Ks(ks string) ApiEntryServerNodeGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEntryServerNodeGetRequest) Format(format int32) ApiEntryServerNodeGetRequest {
	r.format = &format
	return r
}

func (r ApiEntryServerNodeGetRequest) ClientTag(clientTag string) ApiEntryServerNodeGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEntryServerNodeGetRequest) PartnerId(partnerId int32) ApiEntryServerNodeGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEntryServerNodeGetRequest) Body(body AnnotationDeleteRequest) ApiEntryServerNodeGetRequest {
	r.body = &body
	return r
}

func (r ApiEntryServerNodeGetRequest) Execute() (*KalturaEntryServerNode, *http.Response, error) {
	return r.ApiService.EntryServerNodeGetExecute(r)
}

/*
EntryServerNodeGet Method for EntryServerNodeGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntryServerNodeGetRequest
*/
func (a *EntryServerNodeApiService) EntryServerNodeGet(ctx context.Context) ApiEntryServerNodeGetRequest {
	return ApiEntryServerNodeGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEntryServerNode
func (a *EntryServerNodeApiService) EntryServerNodeGetExecute(r ApiEntryServerNodeGetRequest) (*KalturaEntryServerNode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEntryServerNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryServerNodeApiService.EntryServerNodeGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/entryservernode/action/get"

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

type ApiEntryServerNodeListRequest struct {
	ctx context.Context
	ApiService *EntryServerNodeApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EntryServerNodeListRequest
}

func (r ApiEntryServerNodeListRequest) Ks(ks string) ApiEntryServerNodeListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEntryServerNodeListRequest) Format(format int32) ApiEntryServerNodeListRequest {
	r.format = &format
	return r
}

func (r ApiEntryServerNodeListRequest) ClientTag(clientTag string) ApiEntryServerNodeListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEntryServerNodeListRequest) PartnerId(partnerId int32) ApiEntryServerNodeListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEntryServerNodeListRequest) Body(body EntryServerNodeListRequest) ApiEntryServerNodeListRequest {
	r.body = &body
	return r
}

func (r ApiEntryServerNodeListRequest) Execute() (*KalturaEntryServerNodeListResponse, *http.Response, error) {
	return r.ApiService.EntryServerNodeListExecute(r)
}

/*
EntryServerNodeList Method for EntryServerNodeList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntryServerNodeListRequest
*/
func (a *EntryServerNodeApiService) EntryServerNodeList(ctx context.Context) ApiEntryServerNodeListRequest {
	return ApiEntryServerNodeListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEntryServerNodeListResponse
func (a *EntryServerNodeApiService) EntryServerNodeListExecute(r ApiEntryServerNodeListRequest) (*KalturaEntryServerNodeListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEntryServerNodeListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryServerNodeApiService.EntryServerNodeList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/entryservernode/action/list"

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

type ApiEntryServerNodeUpdateRequest struct {
	ctx context.Context
	ApiService *EntryServerNodeApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EntryServerNodeUpdateRequest
}

func (r ApiEntryServerNodeUpdateRequest) Ks(ks string) ApiEntryServerNodeUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEntryServerNodeUpdateRequest) Format(format int32) ApiEntryServerNodeUpdateRequest {
	r.format = &format
	return r
}

func (r ApiEntryServerNodeUpdateRequest) ClientTag(clientTag string) ApiEntryServerNodeUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEntryServerNodeUpdateRequest) PartnerId(partnerId int32) ApiEntryServerNodeUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEntryServerNodeUpdateRequest) Body(body EntryServerNodeUpdateRequest) ApiEntryServerNodeUpdateRequest {
	r.body = &body
	return r
}

func (r ApiEntryServerNodeUpdateRequest) Execute() (*KalturaEntryServerNode, *http.Response, error) {
	return r.ApiService.EntryServerNodeUpdateExecute(r)
}

/*
EntryServerNodeUpdate Method for EntryServerNodeUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntryServerNodeUpdateRequest
*/
func (a *EntryServerNodeApiService) EntryServerNodeUpdate(ctx context.Context) ApiEntryServerNodeUpdateRequest {
	return ApiEntryServerNodeUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEntryServerNode
func (a *EntryServerNodeApiService) EntryServerNodeUpdateExecute(r ApiEntryServerNodeUpdateRequest) (*KalturaEntryServerNode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEntryServerNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryServerNodeApiService.EntryServerNodeUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/entryservernode/action/update"

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

type ApiEntryServerNodeUpdateStatusRequest struct {
	ctx context.Context
	ApiService *EntryServerNodeApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EntryServerNodeUpdateStatusRequest
}

func (r ApiEntryServerNodeUpdateStatusRequest) Ks(ks string) ApiEntryServerNodeUpdateStatusRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEntryServerNodeUpdateStatusRequest) Format(format int32) ApiEntryServerNodeUpdateStatusRequest {
	r.format = &format
	return r
}

func (r ApiEntryServerNodeUpdateStatusRequest) ClientTag(clientTag string) ApiEntryServerNodeUpdateStatusRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEntryServerNodeUpdateStatusRequest) PartnerId(partnerId int32) ApiEntryServerNodeUpdateStatusRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEntryServerNodeUpdateStatusRequest) Body(body EntryServerNodeUpdateStatusRequest) ApiEntryServerNodeUpdateStatusRequest {
	r.body = &body
	return r
}

func (r ApiEntryServerNodeUpdateStatusRequest) Execute() (*KalturaEntryServerNode, *http.Response, error) {
	return r.ApiService.EntryServerNodeUpdateStatusExecute(r)
}

/*
EntryServerNodeUpdateStatus Method for EntryServerNodeUpdateStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntryServerNodeUpdateStatusRequest
*/
func (a *EntryServerNodeApiService) EntryServerNodeUpdateStatus(ctx context.Context) ApiEntryServerNodeUpdateStatusRequest {
	return ApiEntryServerNodeUpdateStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEntryServerNode
func (a *EntryServerNodeApiService) EntryServerNodeUpdateStatusExecute(r ApiEntryServerNodeUpdateStatusRequest) (*KalturaEntryServerNode, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEntryServerNode
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryServerNodeApiService.EntryServerNodeUpdateStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/entryservernode/action/updateStatus"

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

type ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest struct {
	ctx context.Context
	ApiService *EntryServerNodeApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) Ks(ks string) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) Format(format int32) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	r.format = &format
	return r
}

func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) ClientTag(clientTag string) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) PartnerId(partnerId int32) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) Body(body AccessControlDeleteRequest) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	r.body = &body
	return r
}

func (r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) Execute() (*http.Response, error) {
	return r.ApiService.EntryServerNodeValidateRegisteredEntryServerNodeExecute(r)
}

/*
EntryServerNodeValidateRegisteredEntryServerNode Method for EntryServerNodeValidateRegisteredEntryServerNode

Validates server node still registered on entry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest
*/
func (a *EntryServerNodeApiService) EntryServerNodeValidateRegisteredEntryServerNode(ctx context.Context) ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest {
	return ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EntryServerNodeApiService) EntryServerNodeValidateRegisteredEntryServerNodeExecute(r ApiEntryServerNodeValidateRegisteredEntryServerNodeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntryServerNodeApiService.EntryServerNodeValidateRegisteredEntryServerNode")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/entryservernode/action/validateRegisteredEntryServerNode"

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

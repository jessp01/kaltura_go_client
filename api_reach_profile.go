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


// ReachProfileApiService ReachProfileApi service
type ReachProfileApiService service

type ApiReachProfileAddRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ReachProfileAddRequest
}

func (r ApiReachProfileAddRequest) Ks(ks string) ApiReachProfileAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileAddRequest) Format(format int32) ApiReachProfileAddRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileAddRequest) ClientTag(clientTag string) ApiReachProfileAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileAddRequest) PartnerId(partnerId int32) ApiReachProfileAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileAddRequest) Body(body ReachProfileAddRequest) ApiReachProfileAddRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileAddRequest) Execute() (*KalturaReachProfile, *http.Response, error) {
	return r.ApiService.ReachProfileAddExecute(r)
}

/*
ReachProfileAdd Method for ReachProfileAdd

Allows you to add a partner specific reach profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileAddRequest
*/
func (a *ReachProfileApiService) ReachProfileAdd(ctx context.Context) ApiReachProfileAddRequest {
	return ApiReachProfileAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfile
func (a *ReachProfileApiService) ReachProfileAddExecute(r ApiReachProfileAddRequest) (*KalturaReachProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/add"

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

type ApiReachProfileDeleteRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiReachProfileDeleteRequest) Ks(ks string) ApiReachProfileDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileDeleteRequest) Format(format int32) ApiReachProfileDeleteRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileDeleteRequest) ClientTag(clientTag string) ApiReachProfileDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileDeleteRequest) PartnerId(partnerId int32) ApiReachProfileDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileDeleteRequest) Body(body AccessControlDeleteRequest) ApiReachProfileDeleteRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReachProfileDeleteExecute(r)
}

/*
ReachProfileDelete Method for ReachProfileDelete

Delete vednor profile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileDeleteRequest
*/
func (a *ReachProfileApiService) ReachProfileDelete(ctx context.Context) ApiReachProfileDeleteRequest {
	return ApiReachProfileDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ReachProfileApiService) ReachProfileDeleteExecute(r ApiReachProfileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/delete"

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

type ApiReachProfileGetRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiReachProfileGetRequest) Ks(ks string) ApiReachProfileGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileGetRequest) Format(format int32) ApiReachProfileGetRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileGetRequest) ClientTag(clientTag string) ApiReachProfileGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileGetRequest) PartnerId(partnerId int32) ApiReachProfileGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileGetRequest) Body(body AccessControlDeleteRequest) ApiReachProfileGetRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileGetRequest) Execute() (*KalturaReachProfile, *http.Response, error) {
	return r.ApiService.ReachProfileGetExecute(r)
}

/*
ReachProfileGet Method for ReachProfileGet

Retrieve specific reach profile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileGetRequest
*/
func (a *ReachProfileApiService) ReachProfileGet(ctx context.Context) ApiReachProfileGetRequest {
	return ApiReachProfileGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfile
func (a *ReachProfileApiService) ReachProfileGetExecute(r ApiReachProfileGetRequest) (*KalturaReachProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/get"

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

type ApiReachProfileListRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ReachProfileListRequest
}

func (r ApiReachProfileListRequest) Ks(ks string) ApiReachProfileListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileListRequest) Format(format int32) ApiReachProfileListRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileListRequest) ClientTag(clientTag string) ApiReachProfileListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileListRequest) PartnerId(partnerId int32) ApiReachProfileListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileListRequest) Body(body ReachProfileListRequest) ApiReachProfileListRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileListRequest) Execute() (*KalturaReachProfileListResponse, *http.Response, error) {
	return r.ApiService.ReachProfileListExecute(r)
}

/*
ReachProfileList Method for ReachProfileList

List KalturaReachProfile objects

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileListRequest
*/
func (a *ReachProfileApiService) ReachProfileList(ctx context.Context) ApiReachProfileListRequest {
	return ApiReachProfileListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfileListResponse
func (a *ReachProfileApiService) ReachProfileListExecute(r ApiReachProfileListRequest) (*KalturaReachProfileListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfileListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/list"

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

type ApiReachProfileSyncCreditRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ReachProfileSyncCreditRequest
}

func (r ApiReachProfileSyncCreditRequest) Ks(ks string) ApiReachProfileSyncCreditRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileSyncCreditRequest) Format(format int32) ApiReachProfileSyncCreditRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileSyncCreditRequest) ClientTag(clientTag string) ApiReachProfileSyncCreditRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileSyncCreditRequest) PartnerId(partnerId int32) ApiReachProfileSyncCreditRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileSyncCreditRequest) Body(body ReachProfileSyncCreditRequest) ApiReachProfileSyncCreditRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileSyncCreditRequest) Execute() (*KalturaReachProfile, *http.Response, error) {
	return r.ApiService.ReachProfileSyncCreditExecute(r)
}

/*
ReachProfileSyncCredit Method for ReachProfileSyncCredit

sync vendor profile credit

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileSyncCreditRequest
*/
func (a *ReachProfileApiService) ReachProfileSyncCredit(ctx context.Context) ApiReachProfileSyncCreditRequest {
	return ApiReachProfileSyncCreditRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfile
func (a *ReachProfileApiService) ReachProfileSyncCreditExecute(r ApiReachProfileSyncCreditRequest) (*KalturaReachProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileSyncCredit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/syncCredit"

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

type ApiReachProfileUpdateRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ReachProfileUpdateRequest
}

func (r ApiReachProfileUpdateRequest) Ks(ks string) ApiReachProfileUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileUpdateRequest) Format(format int32) ApiReachProfileUpdateRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileUpdateRequest) ClientTag(clientTag string) ApiReachProfileUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileUpdateRequest) PartnerId(partnerId int32) ApiReachProfileUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileUpdateRequest) Body(body ReachProfileUpdateRequest) ApiReachProfileUpdateRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileUpdateRequest) Execute() (*KalturaReachProfile, *http.Response, error) {
	return r.ApiService.ReachProfileUpdateExecute(r)
}

/*
ReachProfileUpdate Method for ReachProfileUpdate

Update an existing reach profile object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileUpdateRequest
*/
func (a *ReachProfileApiService) ReachProfileUpdate(ctx context.Context) ApiReachProfileUpdateRequest {
	return ApiReachProfileUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfile
func (a *ReachProfileApiService) ReachProfileUpdateExecute(r ApiReachProfileUpdateRequest) (*KalturaReachProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/update"

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

type ApiReachProfileUpdateStatusRequest struct {
	ctx context.Context
	ApiService *ReachProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ReachProfileUpdateStatusRequest
}

func (r ApiReachProfileUpdateStatusRequest) Ks(ks string) ApiReachProfileUpdateStatusRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiReachProfileUpdateStatusRequest) Format(format int32) ApiReachProfileUpdateStatusRequest {
	r.format = &format
	return r
}

func (r ApiReachProfileUpdateStatusRequest) ClientTag(clientTag string) ApiReachProfileUpdateStatusRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiReachProfileUpdateStatusRequest) PartnerId(partnerId int32) ApiReachProfileUpdateStatusRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiReachProfileUpdateStatusRequest) Body(body ReachProfileUpdateStatusRequest) ApiReachProfileUpdateStatusRequest {
	r.body = &body
	return r
}

func (r ApiReachProfileUpdateStatusRequest) Execute() (*KalturaReachProfile, *http.Response, error) {
	return r.ApiService.ReachProfileUpdateStatusExecute(r)
}

/*
ReachProfileUpdateStatus Method for ReachProfileUpdateStatus

Update reach profile status by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReachProfileUpdateStatusRequest
*/
func (a *ReachProfileApiService) ReachProfileUpdateStatus(ctx context.Context) ApiReachProfileUpdateStatusRequest {
	return ApiReachProfileUpdateStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaReachProfile
func (a *ReachProfileApiService) ReachProfileUpdateStatusExecute(r ApiReachProfileUpdateStatusRequest) (*KalturaReachProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaReachProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReachProfileApiService.ReachProfileUpdateStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_reachprofile/action/updateStatus"

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

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


// SyndicationFeedApiService SyndicationFeedApi service
type SyndicationFeedApiService service

type ApiSyndicationFeedAddRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SyndicationFeedAddRequest
}

func (r ApiSyndicationFeedAddRequest) Ks(ks string) ApiSyndicationFeedAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedAddRequest) Format(format int32) ApiSyndicationFeedAddRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedAddRequest) ClientTag(clientTag string) ApiSyndicationFeedAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedAddRequest) PartnerId(partnerId int32) ApiSyndicationFeedAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedAddRequest) Body(body SyndicationFeedAddRequest) ApiSyndicationFeedAddRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedAddRequest) Execute() (*KalturaBaseSyndicationFeed, *http.Response, error) {
	return r.ApiService.SyndicationFeedAddExecute(r)
}

/*
SyndicationFeedAdd Method for SyndicationFeedAdd

Add new Syndication Feed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedAddRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedAdd(ctx context.Context) ApiSyndicationFeedAddRequest {
	return ApiSyndicationFeedAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBaseSyndicationFeed
func (a *SyndicationFeedApiService) SyndicationFeedAddExecute(r ApiSyndicationFeedAddRequest) (*KalturaBaseSyndicationFeed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBaseSyndicationFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/add"

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

type ApiSyndicationFeedDeleteRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiSyndicationFeedDeleteRequest) Ks(ks string) ApiSyndicationFeedDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedDeleteRequest) Format(format int32) ApiSyndicationFeedDeleteRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedDeleteRequest) ClientTag(clientTag string) ApiSyndicationFeedDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedDeleteRequest) PartnerId(partnerId int32) ApiSyndicationFeedDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedDeleteRequest) Body(body AnnotationDeleteRequest) ApiSyndicationFeedDeleteRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.SyndicationFeedDeleteExecute(r)
}

/*
SyndicationFeedDelete Method for SyndicationFeedDelete

Delete Syndication Feed by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedDeleteRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedDelete(ctx context.Context) ApiSyndicationFeedDeleteRequest {
	return ApiSyndicationFeedDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SyndicationFeedApiService) SyndicationFeedDeleteExecute(r ApiSyndicationFeedDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/delete"

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

type ApiSyndicationFeedGetRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiSyndicationFeedGetRequest) Ks(ks string) ApiSyndicationFeedGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedGetRequest) Format(format int32) ApiSyndicationFeedGetRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedGetRequest) ClientTag(clientTag string) ApiSyndicationFeedGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedGetRequest) PartnerId(partnerId int32) ApiSyndicationFeedGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedGetRequest) Body(body AnnotationDeleteRequest) ApiSyndicationFeedGetRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedGetRequest) Execute() (*KalturaBaseSyndicationFeed, *http.Response, error) {
	return r.ApiService.SyndicationFeedGetExecute(r)
}

/*
SyndicationFeedGet Method for SyndicationFeedGet

Get Syndication Feed by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedGetRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedGet(ctx context.Context) ApiSyndicationFeedGetRequest {
	return ApiSyndicationFeedGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBaseSyndicationFeed
func (a *SyndicationFeedApiService) SyndicationFeedGetExecute(r ApiSyndicationFeedGetRequest) (*KalturaBaseSyndicationFeed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBaseSyndicationFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/get"

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

type ApiSyndicationFeedGetEntryCountRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SyndicationFeedGetEntryCountRequest
}

func (r ApiSyndicationFeedGetEntryCountRequest) Ks(ks string) ApiSyndicationFeedGetEntryCountRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedGetEntryCountRequest) Format(format int32) ApiSyndicationFeedGetEntryCountRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedGetEntryCountRequest) ClientTag(clientTag string) ApiSyndicationFeedGetEntryCountRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedGetEntryCountRequest) PartnerId(partnerId int32) ApiSyndicationFeedGetEntryCountRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedGetEntryCountRequest) Body(body SyndicationFeedGetEntryCountRequest) ApiSyndicationFeedGetEntryCountRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedGetEntryCountRequest) Execute() (*KalturaSyndicationFeedEntryCount, *http.Response, error) {
	return r.ApiService.SyndicationFeedGetEntryCountExecute(r)
}

/*
SyndicationFeedGetEntryCount Method for SyndicationFeedGetEntryCount

get entry count for a syndication feed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedGetEntryCountRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedGetEntryCount(ctx context.Context) ApiSyndicationFeedGetEntryCountRequest {
	return ApiSyndicationFeedGetEntryCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaSyndicationFeedEntryCount
func (a *SyndicationFeedApiService) SyndicationFeedGetEntryCountExecute(r ApiSyndicationFeedGetEntryCountRequest) (*KalturaSyndicationFeedEntryCount, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaSyndicationFeedEntryCount
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedGetEntryCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/getEntryCount"

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

type ApiSyndicationFeedListRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SyndicationFeedListRequest
}

func (r ApiSyndicationFeedListRequest) Ks(ks string) ApiSyndicationFeedListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedListRequest) Format(format int32) ApiSyndicationFeedListRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedListRequest) ClientTag(clientTag string) ApiSyndicationFeedListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedListRequest) PartnerId(partnerId int32) ApiSyndicationFeedListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedListRequest) Body(body SyndicationFeedListRequest) ApiSyndicationFeedListRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedListRequest) Execute() (*KalturaBaseSyndicationFeedListResponse, *http.Response, error) {
	return r.ApiService.SyndicationFeedListExecute(r)
}

/*
SyndicationFeedList Method for SyndicationFeedList

List Syndication Feeds by filter with paging support

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedListRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedList(ctx context.Context) ApiSyndicationFeedListRequest {
	return ApiSyndicationFeedListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBaseSyndicationFeedListResponse
func (a *SyndicationFeedApiService) SyndicationFeedListExecute(r ApiSyndicationFeedListRequest) (*KalturaBaseSyndicationFeedListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBaseSyndicationFeedListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/list"

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

type ApiSyndicationFeedRequestConversionRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SyndicationFeedGetEntryCountRequest
}

func (r ApiSyndicationFeedRequestConversionRequest) Ks(ks string) ApiSyndicationFeedRequestConversionRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedRequestConversionRequest) Format(format int32) ApiSyndicationFeedRequestConversionRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedRequestConversionRequest) ClientTag(clientTag string) ApiSyndicationFeedRequestConversionRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedRequestConversionRequest) PartnerId(partnerId int32) ApiSyndicationFeedRequestConversionRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedRequestConversionRequest) Body(body SyndicationFeedGetEntryCountRequest) ApiSyndicationFeedRequestConversionRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedRequestConversionRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SyndicationFeedRequestConversionExecute(r)
}

/*
SyndicationFeedRequestConversion Method for SyndicationFeedRequestConversion

request conversion for all entries that doesn't have the required flavor param

returns a comma-separated ids of conversion jobs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedRequestConversionRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedRequestConversion(ctx context.Context) ApiSyndicationFeedRequestConversionRequest {
	return ApiSyndicationFeedRequestConversionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *SyndicationFeedApiService) SyndicationFeedRequestConversionExecute(r ApiSyndicationFeedRequestConversionRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedRequestConversion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/requestConversion"

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

type ApiSyndicationFeedUpdateRequest struct {
	ctx context.Context
	ApiService *SyndicationFeedApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *SyndicationFeedUpdateRequest
}

func (r ApiSyndicationFeedUpdateRequest) Ks(ks string) ApiSyndicationFeedUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiSyndicationFeedUpdateRequest) Format(format int32) ApiSyndicationFeedUpdateRequest {
	r.format = &format
	return r
}

func (r ApiSyndicationFeedUpdateRequest) ClientTag(clientTag string) ApiSyndicationFeedUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiSyndicationFeedUpdateRequest) PartnerId(partnerId int32) ApiSyndicationFeedUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiSyndicationFeedUpdateRequest) Body(body SyndicationFeedUpdateRequest) ApiSyndicationFeedUpdateRequest {
	r.body = &body
	return r
}

func (r ApiSyndicationFeedUpdateRequest) Execute() (*KalturaBaseSyndicationFeed, *http.Response, error) {
	return r.ApiService.SyndicationFeedUpdateExecute(r)
}

/*
SyndicationFeedUpdate Method for SyndicationFeedUpdate

Update Syndication Feed by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSyndicationFeedUpdateRequest
*/
func (a *SyndicationFeedApiService) SyndicationFeedUpdate(ctx context.Context) ApiSyndicationFeedUpdateRequest {
	return ApiSyndicationFeedUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBaseSyndicationFeed
func (a *SyndicationFeedApiService) SyndicationFeedUpdateExecute(r ApiSyndicationFeedUpdateRequest) (*KalturaBaseSyndicationFeed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBaseSyndicationFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SyndicationFeedApiService.SyndicationFeedUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/syndicationfeed/action/update"

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
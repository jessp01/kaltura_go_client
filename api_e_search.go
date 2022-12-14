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


// ESearchApiService ESearchApi service
type ESearchApiService service

type ApiESearchSearchCategoryRequest struct {
	ctx context.Context
	ApiService *ESearchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ESearchSearchCategoryRequest
}

func (r ApiESearchSearchCategoryRequest) Ks(ks string) ApiESearchSearchCategoryRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiESearchSearchCategoryRequest) Format(format int32) ApiESearchSearchCategoryRequest {
	r.format = &format
	return r
}

func (r ApiESearchSearchCategoryRequest) ClientTag(clientTag string) ApiESearchSearchCategoryRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiESearchSearchCategoryRequest) PartnerId(partnerId int32) ApiESearchSearchCategoryRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiESearchSearchCategoryRequest) Body(body ESearchSearchCategoryRequest) ApiESearchSearchCategoryRequest {
	r.body = &body
	return r
}

func (r ApiESearchSearchCategoryRequest) Execute() (*KalturaESearchCategoryResponse, *http.Response, error) {
	return r.ApiService.ESearchSearchCategoryExecute(r)
}

/*
ESearchSearchCategory Method for ESearchSearchCategory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiESearchSearchCategoryRequest
*/
func (a *ESearchApiService) ESearchSearchCategory(ctx context.Context) ApiESearchSearchCategoryRequest {
	return ApiESearchSearchCategoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaESearchCategoryResponse
func (a *ESearchApiService) ESearchSearchCategoryExecute(r ApiESearchSearchCategoryRequest) (*KalturaESearchCategoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaESearchCategoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ESearchApiService.ESearchSearchCategory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/elasticsearch_esearch/action/searchCategory"

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

type ApiESearchSearchEntryRequest struct {
	ctx context.Context
	ApiService *ESearchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ESearchSearchEntryRequest
}

func (r ApiESearchSearchEntryRequest) Ks(ks string) ApiESearchSearchEntryRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiESearchSearchEntryRequest) Format(format int32) ApiESearchSearchEntryRequest {
	r.format = &format
	return r
}

func (r ApiESearchSearchEntryRequest) ClientTag(clientTag string) ApiESearchSearchEntryRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiESearchSearchEntryRequest) PartnerId(partnerId int32) ApiESearchSearchEntryRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiESearchSearchEntryRequest) Body(body ESearchSearchEntryRequest) ApiESearchSearchEntryRequest {
	r.body = &body
	return r
}

func (r ApiESearchSearchEntryRequest) Execute() (*KalturaESearchEntryResponse, *http.Response, error) {
	return r.ApiService.ESearchSearchEntryExecute(r)
}

/*
ESearchSearchEntry Method for ESearchSearchEntry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiESearchSearchEntryRequest
*/
func (a *ESearchApiService) ESearchSearchEntry(ctx context.Context) ApiESearchSearchEntryRequest {
	return ApiESearchSearchEntryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaESearchEntryResponse
func (a *ESearchApiService) ESearchSearchEntryExecute(r ApiESearchSearchEntryRequest) (*KalturaESearchEntryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaESearchEntryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ESearchApiService.ESearchSearchEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/elasticsearch_esearch/action/searchEntry"

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

type ApiESearchSearchGroupRequest struct {
	ctx context.Context
	ApiService *ESearchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ESearchSearchGroupRequest
}

func (r ApiESearchSearchGroupRequest) Ks(ks string) ApiESearchSearchGroupRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiESearchSearchGroupRequest) Format(format int32) ApiESearchSearchGroupRequest {
	r.format = &format
	return r
}

func (r ApiESearchSearchGroupRequest) ClientTag(clientTag string) ApiESearchSearchGroupRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiESearchSearchGroupRequest) PartnerId(partnerId int32) ApiESearchSearchGroupRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiESearchSearchGroupRequest) Body(body ESearchSearchGroupRequest) ApiESearchSearchGroupRequest {
	r.body = &body
	return r
}

func (r ApiESearchSearchGroupRequest) Execute() (*KalturaESearchGroupResponse, *http.Response, error) {
	return r.ApiService.ESearchSearchGroupExecute(r)
}

/*
ESearchSearchGroup Method for ESearchSearchGroup

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiESearchSearchGroupRequest
*/
func (a *ESearchApiService) ESearchSearchGroup(ctx context.Context) ApiESearchSearchGroupRequest {
	return ApiESearchSearchGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaESearchGroupResponse
func (a *ESearchApiService) ESearchSearchGroupExecute(r ApiESearchSearchGroupRequest) (*KalturaESearchGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaESearchGroupResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ESearchApiService.ESearchSearchGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/elasticsearch_esearch/action/searchGroup"

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

type ApiESearchSearchUserRequest struct {
	ctx context.Context
	ApiService *ESearchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ESearchSearchUserRequest
}

func (r ApiESearchSearchUserRequest) Ks(ks string) ApiESearchSearchUserRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiESearchSearchUserRequest) Format(format int32) ApiESearchSearchUserRequest {
	r.format = &format
	return r
}

func (r ApiESearchSearchUserRequest) ClientTag(clientTag string) ApiESearchSearchUserRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiESearchSearchUserRequest) PartnerId(partnerId int32) ApiESearchSearchUserRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiESearchSearchUserRequest) Body(body ESearchSearchUserRequest) ApiESearchSearchUserRequest {
	r.body = &body
	return r
}

func (r ApiESearchSearchUserRequest) Execute() (*KalturaESearchUserResponse, *http.Response, error) {
	return r.ApiService.ESearchSearchUserExecute(r)
}

/*
ESearchSearchUser Method for ESearchSearchUser

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiESearchSearchUserRequest
*/
func (a *ESearchApiService) ESearchSearchUser(ctx context.Context) ApiESearchSearchUserRequest {
	return ApiESearchSearchUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaESearchUserResponse
func (a *ESearchApiService) ESearchSearchUserExecute(r ApiESearchSearchUserRequest) (*KalturaESearchUserResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaESearchUserResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ESearchApiService.ESearchSearchUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/elasticsearch_esearch/action/searchUser"

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

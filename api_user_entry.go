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


// UserEntryApiService UserEntryApi service
type UserEntryApiService service

type ApiUserEntryAddRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserEntryAddRequest
}

func (r ApiUserEntryAddRequest) Ks(ks string) ApiUserEntryAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryAddRequest) Format(format int32) ApiUserEntryAddRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryAddRequest) ClientTag(clientTag string) ApiUserEntryAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryAddRequest) PartnerId(partnerId int32) ApiUserEntryAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryAddRequest) Body(body UserEntryAddRequest) ApiUserEntryAddRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryAddRequest) Execute() (*KalturaUserEntry, *http.Response, error) {
	return r.ApiService.UserEntryAddExecute(r)
}

/*
UserEntryAdd Method for UserEntryAdd

Adds a user_entry to the Kaltura DB.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryAddRequest
*/
func (a *UserEntryApiService) UserEntryAdd(ctx context.Context) ApiUserEntryAddRequest {
	return ApiUserEntryAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserEntry
func (a *UserEntryApiService) UserEntryAddExecute(r ApiUserEntryAddRequest) (*KalturaUserEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/add"

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

type ApiUserEntryBulkDeleteRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserEntryBulkDeleteRequest
}

func (r ApiUserEntryBulkDeleteRequest) Ks(ks string) ApiUserEntryBulkDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryBulkDeleteRequest) Format(format int32) ApiUserEntryBulkDeleteRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryBulkDeleteRequest) ClientTag(clientTag string) ApiUserEntryBulkDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryBulkDeleteRequest) PartnerId(partnerId int32) ApiUserEntryBulkDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryBulkDeleteRequest) Body(body UserEntryBulkDeleteRequest) ApiUserEntryBulkDeleteRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryBulkDeleteRequest) Execute() (int32, *http.Response, error) {
	return r.ApiService.UserEntryBulkDeleteExecute(r)
}

/*
UserEntryBulkDelete Method for UserEntryBulkDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryBulkDeleteRequest
*/
func (a *UserEntryApiService) UserEntryBulkDelete(ctx context.Context) ApiUserEntryBulkDeleteRequest {
	return ApiUserEntryBulkDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return int32
func (a *UserEntryApiService) UserEntryBulkDeleteExecute(r ApiUserEntryBulkDeleteRequest) (int32, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryBulkDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/bulkDelete"

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

type ApiUserEntryDeleteRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiUserEntryDeleteRequest) Ks(ks string) ApiUserEntryDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryDeleteRequest) Format(format int32) ApiUserEntryDeleteRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryDeleteRequest) ClientTag(clientTag string) ApiUserEntryDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryDeleteRequest) PartnerId(partnerId int32) ApiUserEntryDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryDeleteRequest) Body(body AccessControlDeleteRequest) ApiUserEntryDeleteRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryDeleteRequest) Execute() (*KalturaUserEntry, *http.Response, error) {
	return r.ApiService.UserEntryDeleteExecute(r)
}

/*
UserEntryDelete Method for UserEntryDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryDeleteRequest
*/
func (a *UserEntryApiService) UserEntryDelete(ctx context.Context) ApiUserEntryDeleteRequest {
	return ApiUserEntryDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserEntry
func (a *UserEntryApiService) UserEntryDeleteExecute(r ApiUserEntryDeleteRequest) (*KalturaUserEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/delete"

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

type ApiUserEntryGetRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiUserEntryGetRequest) Ks(ks string) ApiUserEntryGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryGetRequest) Format(format int32) ApiUserEntryGetRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryGetRequest) ClientTag(clientTag string) ApiUserEntryGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryGetRequest) PartnerId(partnerId int32) ApiUserEntryGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryGetRequest) Body(body AnnotationDeleteRequest) ApiUserEntryGetRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryGetRequest) Execute() (*KalturaUserEntry, *http.Response, error) {
	return r.ApiService.UserEntryGetExecute(r)
}

/*
UserEntryGet Method for UserEntryGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryGetRequest
*/
func (a *UserEntryApiService) UserEntryGet(ctx context.Context) ApiUserEntryGetRequest {
	return ApiUserEntryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserEntry
func (a *UserEntryApiService) UserEntryGetExecute(r ApiUserEntryGetRequest) (*KalturaUserEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/get"

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

type ApiUserEntryListRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserEntryListRequest
}

func (r ApiUserEntryListRequest) Ks(ks string) ApiUserEntryListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryListRequest) Format(format int32) ApiUserEntryListRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryListRequest) ClientTag(clientTag string) ApiUserEntryListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryListRequest) PartnerId(partnerId int32) ApiUserEntryListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryListRequest) Body(body UserEntryListRequest) ApiUserEntryListRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryListRequest) Execute() (*KalturaUserEntryListResponse, *http.Response, error) {
	return r.ApiService.UserEntryListExecute(r)
}

/*
UserEntryList Method for UserEntryList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryListRequest
*/
func (a *UserEntryApiService) UserEntryList(ctx context.Context) ApiUserEntryListRequest {
	return ApiUserEntryListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserEntryListResponse
func (a *UserEntryApiService) UserEntryListExecute(r ApiUserEntryListRequest) (*KalturaUserEntryListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserEntryListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/list"

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

type ApiUserEntrySubmitQuizRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiUserEntrySubmitQuizRequest) Ks(ks string) ApiUserEntrySubmitQuizRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntrySubmitQuizRequest) Format(format int32) ApiUserEntrySubmitQuizRequest {
	r.format = &format
	return r
}

func (r ApiUserEntrySubmitQuizRequest) ClientTag(clientTag string) ApiUserEntrySubmitQuizRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntrySubmitQuizRequest) PartnerId(partnerId int32) ApiUserEntrySubmitQuizRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntrySubmitQuizRequest) Body(body AccessControlDeleteRequest) ApiUserEntrySubmitQuizRequest {
	r.body = &body
	return r
}

func (r ApiUserEntrySubmitQuizRequest) Execute() (*KalturaQuizUserEntry, *http.Response, error) {
	return r.ApiService.UserEntrySubmitQuizExecute(r)
}

/*
UserEntrySubmitQuiz Method for UserEntrySubmitQuiz

Submits the quiz so that it's status will be submitted and calculates the score for the quiz

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntrySubmitQuizRequest
*/
func (a *UserEntryApiService) UserEntrySubmitQuiz(ctx context.Context) ApiUserEntrySubmitQuizRequest {
	return ApiUserEntrySubmitQuizRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaQuizUserEntry
func (a *UserEntryApiService) UserEntrySubmitQuizExecute(r ApiUserEntrySubmitQuizRequest) (*KalturaQuizUserEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaQuizUserEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntrySubmitQuiz")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/submitQuiz"

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

type ApiUserEntryUpdateRequest struct {
	ctx context.Context
	ApiService *UserEntryApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserEntryUpdateRequest
}

func (r ApiUserEntryUpdateRequest) Ks(ks string) ApiUserEntryUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserEntryUpdateRequest) Format(format int32) ApiUserEntryUpdateRequest {
	r.format = &format
	return r
}

func (r ApiUserEntryUpdateRequest) ClientTag(clientTag string) ApiUserEntryUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserEntryUpdateRequest) PartnerId(partnerId int32) ApiUserEntryUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserEntryUpdateRequest) Body(body UserEntryUpdateRequest) ApiUserEntryUpdateRequest {
	r.body = &body
	return r
}

func (r ApiUserEntryUpdateRequest) Execute() (*KalturaUserEntry, *http.Response, error) {
	return r.ApiService.UserEntryUpdateExecute(r)
}

/*
UserEntryUpdate Method for UserEntryUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserEntryUpdateRequest
*/
func (a *UserEntryApiService) UserEntryUpdate(ctx context.Context) ApiUserEntryUpdateRequest {
	return ApiUserEntryUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserEntry
func (a *UserEntryApiService) UserEntryUpdateExecute(r ApiUserEntryUpdateRequest) (*KalturaUserEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEntryApiService.UserEntryUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userentry/action/update"

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

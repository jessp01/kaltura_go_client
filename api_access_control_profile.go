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


// AccessControlProfileApiService AccessControlProfileApi service
type AccessControlProfileApiService service

type ApiAccessControlProfileAddRequest struct {
	ctx context.Context
	ApiService *AccessControlProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlProfileAddRequest
}

func (r ApiAccessControlProfileAddRequest) Ks(ks string) ApiAccessControlProfileAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAccessControlProfileAddRequest) Format(format int32) ApiAccessControlProfileAddRequest {
	r.format = &format
	return r
}

func (r ApiAccessControlProfileAddRequest) ClientTag(clientTag string) ApiAccessControlProfileAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAccessControlProfileAddRequest) PartnerId(partnerId int32) ApiAccessControlProfileAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAccessControlProfileAddRequest) Body(body AccessControlProfileAddRequest) ApiAccessControlProfileAddRequest {
	r.body = &body
	return r
}

func (r ApiAccessControlProfileAddRequest) Execute() (*KalturaAccessControlProfile, *http.Response, error) {
	return r.ApiService.AccessControlProfileAddExecute(r)
}

/*
AccessControlProfileAdd Method for AccessControlProfileAdd

Add new access control profile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccessControlProfileAddRequest
*/
func (a *AccessControlProfileApiService) AccessControlProfileAdd(ctx context.Context) ApiAccessControlProfileAddRequest {
	return ApiAccessControlProfileAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAccessControlProfile
func (a *AccessControlProfileApiService) AccessControlProfileAddExecute(r ApiAccessControlProfileAddRequest) (*KalturaAccessControlProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAccessControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessControlProfileApiService.AccessControlProfileAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/accesscontrolprofile/action/add"

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

type ApiAccessControlProfileDeleteRequest struct {
	ctx context.Context
	ApiService *AccessControlProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiAccessControlProfileDeleteRequest) Ks(ks string) ApiAccessControlProfileDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAccessControlProfileDeleteRequest) Format(format int32) ApiAccessControlProfileDeleteRequest {
	r.format = &format
	return r
}

func (r ApiAccessControlProfileDeleteRequest) ClientTag(clientTag string) ApiAccessControlProfileDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAccessControlProfileDeleteRequest) PartnerId(partnerId int32) ApiAccessControlProfileDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAccessControlProfileDeleteRequest) Body(body AccessControlDeleteRequest) ApiAccessControlProfileDeleteRequest {
	r.body = &body
	return r
}

func (r ApiAccessControlProfileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.AccessControlProfileDeleteExecute(r)
}

/*
AccessControlProfileDelete Method for AccessControlProfileDelete

Delete access control profile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccessControlProfileDeleteRequest
*/
func (a *AccessControlProfileApiService) AccessControlProfileDelete(ctx context.Context) ApiAccessControlProfileDeleteRequest {
	return ApiAccessControlProfileDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AccessControlProfileApiService) AccessControlProfileDeleteExecute(r ApiAccessControlProfileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessControlProfileApiService.AccessControlProfileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/accesscontrolprofile/action/delete"

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

type ApiAccessControlProfileGetRequest struct {
	ctx context.Context
	ApiService *AccessControlProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiAccessControlProfileGetRequest) Ks(ks string) ApiAccessControlProfileGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAccessControlProfileGetRequest) Format(format int32) ApiAccessControlProfileGetRequest {
	r.format = &format
	return r
}

func (r ApiAccessControlProfileGetRequest) ClientTag(clientTag string) ApiAccessControlProfileGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAccessControlProfileGetRequest) PartnerId(partnerId int32) ApiAccessControlProfileGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAccessControlProfileGetRequest) Body(body AccessControlDeleteRequest) ApiAccessControlProfileGetRequest {
	r.body = &body
	return r
}

func (r ApiAccessControlProfileGetRequest) Execute() (*KalturaAccessControlProfile, *http.Response, error) {
	return r.ApiService.AccessControlProfileGetExecute(r)
}

/*
AccessControlProfileGet Method for AccessControlProfileGet

Get access control profile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccessControlProfileGetRequest
*/
func (a *AccessControlProfileApiService) AccessControlProfileGet(ctx context.Context) ApiAccessControlProfileGetRequest {
	return ApiAccessControlProfileGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAccessControlProfile
func (a *AccessControlProfileApiService) AccessControlProfileGetExecute(r ApiAccessControlProfileGetRequest) (*KalturaAccessControlProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAccessControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessControlProfileApiService.AccessControlProfileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/accesscontrolprofile/action/get"

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

type ApiAccessControlProfileListRequest struct {
	ctx context.Context
	ApiService *AccessControlProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlProfileListRequest
}

func (r ApiAccessControlProfileListRequest) Ks(ks string) ApiAccessControlProfileListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAccessControlProfileListRequest) Format(format int32) ApiAccessControlProfileListRequest {
	r.format = &format
	return r
}

func (r ApiAccessControlProfileListRequest) ClientTag(clientTag string) ApiAccessControlProfileListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAccessControlProfileListRequest) PartnerId(partnerId int32) ApiAccessControlProfileListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAccessControlProfileListRequest) Body(body AccessControlProfileListRequest) ApiAccessControlProfileListRequest {
	r.body = &body
	return r
}

func (r ApiAccessControlProfileListRequest) Execute() (*KalturaAccessControlProfileListResponse, *http.Response, error) {
	return r.ApiService.AccessControlProfileListExecute(r)
}

/*
AccessControlProfileList Method for AccessControlProfileList

List access control profiles by filter and pager

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccessControlProfileListRequest
*/
func (a *AccessControlProfileApiService) AccessControlProfileList(ctx context.Context) ApiAccessControlProfileListRequest {
	return ApiAccessControlProfileListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAccessControlProfileListResponse
func (a *AccessControlProfileApiService) AccessControlProfileListExecute(r ApiAccessControlProfileListRequest) (*KalturaAccessControlProfileListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAccessControlProfileListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessControlProfileApiService.AccessControlProfileList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/accesscontrolprofile/action/list"

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

type ApiAccessControlProfileUpdateRequest struct {
	ctx context.Context
	ApiService *AccessControlProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlProfileUpdateRequest
}

func (r ApiAccessControlProfileUpdateRequest) Ks(ks string) ApiAccessControlProfileUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAccessControlProfileUpdateRequest) Format(format int32) ApiAccessControlProfileUpdateRequest {
	r.format = &format
	return r
}

func (r ApiAccessControlProfileUpdateRequest) ClientTag(clientTag string) ApiAccessControlProfileUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAccessControlProfileUpdateRequest) PartnerId(partnerId int32) ApiAccessControlProfileUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAccessControlProfileUpdateRequest) Body(body AccessControlProfileUpdateRequest) ApiAccessControlProfileUpdateRequest {
	r.body = &body
	return r
}

func (r ApiAccessControlProfileUpdateRequest) Execute() (*KalturaAccessControlProfile, *http.Response, error) {
	return r.ApiService.AccessControlProfileUpdateExecute(r)
}

/*
AccessControlProfileUpdate Method for AccessControlProfileUpdate

Update access control profile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAccessControlProfileUpdateRequest
*/
func (a *AccessControlProfileApiService) AccessControlProfileUpdate(ctx context.Context) ApiAccessControlProfileUpdateRequest {
	return ApiAccessControlProfileUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAccessControlProfile
func (a *AccessControlProfileApiService) AccessControlProfileUpdateExecute(r ApiAccessControlProfileUpdateRequest) (*KalturaAccessControlProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAccessControlProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccessControlProfileApiService.AccessControlProfileUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/accesscontrolprofile/action/update"

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
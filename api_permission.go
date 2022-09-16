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


// PermissionApiService PermissionApi service
type PermissionApiService service

type ApiPermissionAddRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PermissionAddRequest
}

func (r ApiPermissionAddRequest) Ks(ks string) ApiPermissionAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionAddRequest) Format(format int32) ApiPermissionAddRequest {
	r.format = &format
	return r
}

func (r ApiPermissionAddRequest) ClientTag(clientTag string) ApiPermissionAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionAddRequest) PartnerId(partnerId int32) ApiPermissionAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionAddRequest) Body(body PermissionAddRequest) ApiPermissionAddRequest {
	r.body = &body
	return r
}

func (r ApiPermissionAddRequest) Execute() (*KalturaPermission, *http.Response, error) {
	return r.ApiService.PermissionAddExecute(r)
}

/*
PermissionAdd Method for PermissionAdd

Adds a new permission object to the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionAddRequest
*/
func (a *PermissionApiService) PermissionAdd(ctx context.Context) ApiPermissionAddRequest {
	return ApiPermissionAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaPermission
func (a *PermissionApiService) PermissionAddExecute(r ApiPermissionAddRequest) (*KalturaPermission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaPermission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/add"

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

type ApiPermissionDeleteRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PermissionDeleteRequest
}

func (r ApiPermissionDeleteRequest) Ks(ks string) ApiPermissionDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionDeleteRequest) Format(format int32) ApiPermissionDeleteRequest {
	r.format = &format
	return r
}

func (r ApiPermissionDeleteRequest) ClientTag(clientTag string) ApiPermissionDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionDeleteRequest) PartnerId(partnerId int32) ApiPermissionDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionDeleteRequest) Body(body PermissionDeleteRequest) ApiPermissionDeleteRequest {
	r.body = &body
	return r
}

func (r ApiPermissionDeleteRequest) Execute() (*KalturaPermission, *http.Response, error) {
	return r.ApiService.PermissionDeleteExecute(r)
}

/*
PermissionDelete Method for PermissionDelete

Deletes an existing permission object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionDeleteRequest
*/
func (a *PermissionApiService) PermissionDelete(ctx context.Context) ApiPermissionDeleteRequest {
	return ApiPermissionDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaPermission
func (a *PermissionApiService) PermissionDeleteExecute(r ApiPermissionDeleteRequest) (*KalturaPermission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaPermission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/delete"

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

type ApiPermissionGetRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PermissionDeleteRequest
}

func (r ApiPermissionGetRequest) Ks(ks string) ApiPermissionGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionGetRequest) Format(format int32) ApiPermissionGetRequest {
	r.format = &format
	return r
}

func (r ApiPermissionGetRequest) ClientTag(clientTag string) ApiPermissionGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionGetRequest) PartnerId(partnerId int32) ApiPermissionGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionGetRequest) Body(body PermissionDeleteRequest) ApiPermissionGetRequest {
	r.body = &body
	return r
}

func (r ApiPermissionGetRequest) Execute() (*KalturaPermission, *http.Response, error) {
	return r.ApiService.PermissionGetExecute(r)
}

/*
PermissionGet Method for PermissionGet

Retrieves a permission object using its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGetRequest
*/
func (a *PermissionApiService) PermissionGet(ctx context.Context) ApiPermissionGetRequest {
	return ApiPermissionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaPermission
func (a *PermissionApiService) PermissionGetExecute(r ApiPermissionGetRequest) (*KalturaPermission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaPermission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/get"

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

type ApiPermissionGetCurrentPermissionsRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	batchCleanExclusiveJobsRequest *BatchCleanExclusiveJobsRequest
}

func (r ApiPermissionGetCurrentPermissionsRequest) Ks(ks string) ApiPermissionGetCurrentPermissionsRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionGetCurrentPermissionsRequest) Format(format int32) ApiPermissionGetCurrentPermissionsRequest {
	r.format = &format
	return r
}

func (r ApiPermissionGetCurrentPermissionsRequest) ClientTag(clientTag string) ApiPermissionGetCurrentPermissionsRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionGetCurrentPermissionsRequest) PartnerId(partnerId int32) ApiPermissionGetCurrentPermissionsRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionGetCurrentPermissionsRequest) BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest BatchCleanExclusiveJobsRequest) ApiPermissionGetCurrentPermissionsRequest {
	r.batchCleanExclusiveJobsRequest = &batchCleanExclusiveJobsRequest
	return r
}

func (r ApiPermissionGetCurrentPermissionsRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PermissionGetCurrentPermissionsExecute(r)
}

/*
PermissionGetCurrentPermissions Method for PermissionGetCurrentPermissions

Retrieves a list of permissions that apply to the current KS.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionGetCurrentPermissionsRequest
*/
func (a *PermissionApiService) PermissionGetCurrentPermissions(ctx context.Context) ApiPermissionGetCurrentPermissionsRequest {
	return ApiPermissionGetCurrentPermissionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *PermissionApiService) PermissionGetCurrentPermissionsExecute(r ApiPermissionGetCurrentPermissionsRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionGetCurrentPermissions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/getCurrentPermissions"

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

type ApiPermissionListRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PermissionListRequest
}

func (r ApiPermissionListRequest) Ks(ks string) ApiPermissionListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionListRequest) Format(format int32) ApiPermissionListRequest {
	r.format = &format
	return r
}

func (r ApiPermissionListRequest) ClientTag(clientTag string) ApiPermissionListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionListRequest) PartnerId(partnerId int32) ApiPermissionListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionListRequest) Body(body PermissionListRequest) ApiPermissionListRequest {
	r.body = &body
	return r
}

func (r ApiPermissionListRequest) Execute() (*KalturaPermissionListResponse, *http.Response, error) {
	return r.ApiService.PermissionListExecute(r)
}

/*
PermissionList Method for PermissionList

Lists permission objects that are associated with an account.

Blocked permissions are listed unless you use a filter to exclude them.

Blocked permissions are listed unless you use a filter to exclude them.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionListRequest
*/
func (a *PermissionApiService) PermissionList(ctx context.Context) ApiPermissionListRequest {
	return ApiPermissionListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaPermissionListResponse
func (a *PermissionApiService) PermissionListExecute(r ApiPermissionListRequest) (*KalturaPermissionListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaPermissionListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/list"

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

type ApiPermissionUpdateRequest struct {
	ctx context.Context
	ApiService *PermissionApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PermissionUpdateRequest
}

func (r ApiPermissionUpdateRequest) Ks(ks string) ApiPermissionUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPermissionUpdateRequest) Format(format int32) ApiPermissionUpdateRequest {
	r.format = &format
	return r
}

func (r ApiPermissionUpdateRequest) ClientTag(clientTag string) ApiPermissionUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPermissionUpdateRequest) PartnerId(partnerId int32) ApiPermissionUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPermissionUpdateRequest) Body(body PermissionUpdateRequest) ApiPermissionUpdateRequest {
	r.body = &body
	return r
}

func (r ApiPermissionUpdateRequest) Execute() (*KalturaPermission, *http.Response, error) {
	return r.ApiService.PermissionUpdateExecute(r)
}

/*
PermissionUpdate Method for PermissionUpdate

Updates an existing permission object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPermissionUpdateRequest
*/
func (a *PermissionApiService) PermissionUpdate(ctx context.Context) ApiPermissionUpdateRequest {
	return ApiPermissionUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaPermission
func (a *PermissionApiService) PermissionUpdateExecute(r ApiPermissionUpdateRequest) (*KalturaPermission, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaPermission
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PermissionApiService.PermissionUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/permission/action/update"

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

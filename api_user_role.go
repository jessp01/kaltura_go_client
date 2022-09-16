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


// UserRoleApiService UserRoleApi service
type UserRoleApiService service

type ApiUserRoleAddRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleAddRequest
}

func (r ApiUserRoleAddRequest) Ks(ks string) ApiUserRoleAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleAddRequest) Format(format int32) ApiUserRoleAddRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleAddRequest) ClientTag(clientTag string) ApiUserRoleAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleAddRequest) PartnerId(partnerId int32) ApiUserRoleAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleAddRequest) Body(body UserRoleAddRequest) ApiUserRoleAddRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleAddRequest) Execute() (*KalturaUserRole, *http.Response, error) {
	return r.ApiService.UserRoleAddExecute(r)
}

/*
UserRoleAdd Method for UserRoleAdd

Adds a new user role object to the account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleAddRequest
*/
func (a *UserRoleApiService) UserRoleAdd(ctx context.Context) ApiUserRoleAddRequest {
	return ApiUserRoleAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRole
func (a *UserRoleApiService) UserRoleAddExecute(r ApiUserRoleAddRequest) (*KalturaUserRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/add"

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

type ApiUserRoleCloneRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleCloneRequest
}

func (r ApiUserRoleCloneRequest) Ks(ks string) ApiUserRoleCloneRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleCloneRequest) Format(format int32) ApiUserRoleCloneRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleCloneRequest) ClientTag(clientTag string) ApiUserRoleCloneRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleCloneRequest) PartnerId(partnerId int32) ApiUserRoleCloneRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleCloneRequest) Body(body UserRoleCloneRequest) ApiUserRoleCloneRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleCloneRequest) Execute() (*KalturaUserRole, *http.Response, error) {
	return r.ApiService.UserRoleCloneExecute(r)
}

/*
UserRoleClone Method for UserRoleClone

Creates a new user role object that is a duplicate of an existing role.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleCloneRequest
*/
func (a *UserRoleApiService) UserRoleClone(ctx context.Context) ApiUserRoleCloneRequest {
	return ApiUserRoleCloneRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRole
func (a *UserRoleApiService) UserRoleCloneExecute(r ApiUserRoleCloneRequest) (*KalturaUserRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleClone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/clone"

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

type ApiUserRoleDeleteRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleCloneRequest
}

func (r ApiUserRoleDeleteRequest) Ks(ks string) ApiUserRoleDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleDeleteRequest) Format(format int32) ApiUserRoleDeleteRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleDeleteRequest) ClientTag(clientTag string) ApiUserRoleDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleDeleteRequest) PartnerId(partnerId int32) ApiUserRoleDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleDeleteRequest) Body(body UserRoleCloneRequest) ApiUserRoleDeleteRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleDeleteRequest) Execute() (*KalturaUserRole, *http.Response, error) {
	return r.ApiService.UserRoleDeleteExecute(r)
}

/*
UserRoleDelete Method for UserRoleDelete

Deletes an existing user role object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleDeleteRequest
*/
func (a *UserRoleApiService) UserRoleDelete(ctx context.Context) ApiUserRoleDeleteRequest {
	return ApiUserRoleDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRole
func (a *UserRoleApiService) UserRoleDeleteExecute(r ApiUserRoleDeleteRequest) (*KalturaUserRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/delete"

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

type ApiUserRoleGetRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleCloneRequest
}

func (r ApiUserRoleGetRequest) Ks(ks string) ApiUserRoleGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleGetRequest) Format(format int32) ApiUserRoleGetRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleGetRequest) ClientTag(clientTag string) ApiUserRoleGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleGetRequest) PartnerId(partnerId int32) ApiUserRoleGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleGetRequest) Body(body UserRoleCloneRequest) ApiUserRoleGetRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleGetRequest) Execute() (*KalturaUserRole, *http.Response, error) {
	return r.ApiService.UserRoleGetExecute(r)
}

/*
UserRoleGet Method for UserRoleGet

Retrieves a user role object using its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleGetRequest
*/
func (a *UserRoleApiService) UserRoleGet(ctx context.Context) ApiUserRoleGetRequest {
	return ApiUserRoleGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRole
func (a *UserRoleApiService) UserRoleGetExecute(r ApiUserRoleGetRequest) (*KalturaUserRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/get"

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

type ApiUserRoleListRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleListRequest
}

func (r ApiUserRoleListRequest) Ks(ks string) ApiUserRoleListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleListRequest) Format(format int32) ApiUserRoleListRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleListRequest) ClientTag(clientTag string) ApiUserRoleListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleListRequest) PartnerId(partnerId int32) ApiUserRoleListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleListRequest) Body(body UserRoleListRequest) ApiUserRoleListRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleListRequest) Execute() (*KalturaUserRoleListResponse, *http.Response, error) {
	return r.ApiService.UserRoleListExecute(r)
}

/*
UserRoleList Method for UserRoleList

Lists user role objects that are associated with an account.

Blocked user roles are listed unless you use a filter to exclude them.

Deleted user roles are not listed unless you use a filter to include them.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleListRequest
*/
func (a *UserRoleApiService) UserRoleList(ctx context.Context) ApiUserRoleListRequest {
	return ApiUserRoleListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRoleListResponse
func (a *UserRoleApiService) UserRoleListExecute(r ApiUserRoleListRequest) (*KalturaUserRoleListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRoleListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/list"

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

type ApiUserRoleUpdateRequest struct {
	ctx context.Context
	ApiService *UserRoleApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UserRoleUpdateRequest
}

func (r ApiUserRoleUpdateRequest) Ks(ks string) ApiUserRoleUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUserRoleUpdateRequest) Format(format int32) ApiUserRoleUpdateRequest {
	r.format = &format
	return r
}

func (r ApiUserRoleUpdateRequest) ClientTag(clientTag string) ApiUserRoleUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUserRoleUpdateRequest) PartnerId(partnerId int32) ApiUserRoleUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUserRoleUpdateRequest) Body(body UserRoleUpdateRequest) ApiUserRoleUpdateRequest {
	r.body = &body
	return r
}

func (r ApiUserRoleUpdateRequest) Execute() (*KalturaUserRole, *http.Response, error) {
	return r.ApiService.UserRoleUpdateExecute(r)
}

/*
UserRoleUpdate Method for UserRoleUpdate

Updates an existing user role object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserRoleUpdateRequest
*/
func (a *UserRoleApiService) UserRoleUpdate(ctx context.Context) ApiUserRoleUpdateRequest {
	return ApiUserRoleUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUserRole
func (a *UserRoleApiService) UserRoleUpdateExecute(r ApiUserRoleUpdateRequest) (*KalturaUserRole, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUserRole
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserRoleApiService.UserRoleUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/userrole/action/update"

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

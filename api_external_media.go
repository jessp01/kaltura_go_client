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


// ExternalMediaApiService ExternalMediaApi service
type ExternalMediaApiService service

type ApiExternalMediaAddRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ExternalMediaAddRequest
}

func (r ApiExternalMediaAddRequest) Ks(ks string) ApiExternalMediaAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaAddRequest) Format(format int32) ApiExternalMediaAddRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaAddRequest) ClientTag(clientTag string) ApiExternalMediaAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaAddRequest) PartnerId(partnerId int32) ApiExternalMediaAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaAddRequest) Body(body ExternalMediaAddRequest) ApiExternalMediaAddRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaAddRequest) Execute() (*KalturaExternalMediaEntry, *http.Response, error) {
	return r.ApiService.ExternalMediaAddExecute(r)
}

/*
ExternalMediaAdd Method for ExternalMediaAdd

Add external media entry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaAddRequest
*/
func (a *ExternalMediaApiService) ExternalMediaAdd(ctx context.Context) ApiExternalMediaAddRequest {
	return ApiExternalMediaAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaExternalMediaEntry
func (a *ExternalMediaApiService) ExternalMediaAddExecute(r ApiExternalMediaAddRequest) (*KalturaExternalMediaEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaExternalMediaEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/add"

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

type ApiExternalMediaCountRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ExternalMediaCountRequest
}

func (r ApiExternalMediaCountRequest) Ks(ks string) ApiExternalMediaCountRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaCountRequest) Format(format int32) ApiExternalMediaCountRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaCountRequest) ClientTag(clientTag string) ApiExternalMediaCountRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaCountRequest) PartnerId(partnerId int32) ApiExternalMediaCountRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaCountRequest) Body(body ExternalMediaCountRequest) ApiExternalMediaCountRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaCountRequest) Execute() (int32, *http.Response, error) {
	return r.ApiService.ExternalMediaCountExecute(r)
}

/*
ExternalMediaCount Method for ExternalMediaCount

Count media entries by filter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaCountRequest
*/
func (a *ExternalMediaApiService) ExternalMediaCount(ctx context.Context) ApiExternalMediaCountRequest {
	return ApiExternalMediaCountRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return int32
func (a *ExternalMediaApiService) ExternalMediaCountExecute(r ApiExternalMediaCountRequest) (int32, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaCount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/count"

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

type ApiExternalMediaDeleteRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiExternalMediaDeleteRequest) Ks(ks string) ApiExternalMediaDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaDeleteRequest) Format(format int32) ApiExternalMediaDeleteRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaDeleteRequest) ClientTag(clientTag string) ApiExternalMediaDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaDeleteRequest) PartnerId(partnerId int32) ApiExternalMediaDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaDeleteRequest) Body(body AnnotationDeleteRequest) ApiExternalMediaDeleteRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ExternalMediaDeleteExecute(r)
}

/*
ExternalMediaDelete Method for ExternalMediaDelete

Delete a external media entry.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaDeleteRequest
*/
func (a *ExternalMediaApiService) ExternalMediaDelete(ctx context.Context) ApiExternalMediaDeleteRequest {
	return ApiExternalMediaDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ExternalMediaApiService) ExternalMediaDeleteExecute(r ApiExternalMediaDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/delete"

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

type ApiExternalMediaGetRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AnnotationDeleteRequest
}

func (r ApiExternalMediaGetRequest) Ks(ks string) ApiExternalMediaGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaGetRequest) Format(format int32) ApiExternalMediaGetRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaGetRequest) ClientTag(clientTag string) ApiExternalMediaGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaGetRequest) PartnerId(partnerId int32) ApiExternalMediaGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaGetRequest) Body(body AnnotationDeleteRequest) ApiExternalMediaGetRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaGetRequest) Execute() (*KalturaExternalMediaEntry, *http.Response, error) {
	return r.ApiService.ExternalMediaGetExecute(r)
}

/*
ExternalMediaGet Method for ExternalMediaGet

Get external media entry by ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaGetRequest
*/
func (a *ExternalMediaApiService) ExternalMediaGet(ctx context.Context) ApiExternalMediaGetRequest {
	return ApiExternalMediaGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaExternalMediaEntry
func (a *ExternalMediaApiService) ExternalMediaGetExecute(r ApiExternalMediaGetRequest) (*KalturaExternalMediaEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaExternalMediaEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/get"

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

type ApiExternalMediaListRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ExternalMediaListRequest
}

func (r ApiExternalMediaListRequest) Ks(ks string) ApiExternalMediaListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaListRequest) Format(format int32) ApiExternalMediaListRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaListRequest) ClientTag(clientTag string) ApiExternalMediaListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaListRequest) PartnerId(partnerId int32) ApiExternalMediaListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaListRequest) Body(body ExternalMediaListRequest) ApiExternalMediaListRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaListRequest) Execute() (*KalturaExternalMediaEntryListResponse, *http.Response, error) {
	return r.ApiService.ExternalMediaListExecute(r)
}

/*
ExternalMediaList Method for ExternalMediaList

List media entries by filter with paging support.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaListRequest
*/
func (a *ExternalMediaApiService) ExternalMediaList(ctx context.Context) ApiExternalMediaListRequest {
	return ApiExternalMediaListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaExternalMediaEntryListResponse
func (a *ExternalMediaApiService) ExternalMediaListExecute(r ApiExternalMediaListRequest) (*KalturaExternalMediaEntryListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaExternalMediaEntryListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/list"

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

type ApiExternalMediaUpdateRequest struct {
	ctx context.Context
	ApiService *ExternalMediaApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ExternalMediaUpdateRequest
}

func (r ApiExternalMediaUpdateRequest) Ks(ks string) ApiExternalMediaUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiExternalMediaUpdateRequest) Format(format int32) ApiExternalMediaUpdateRequest {
	r.format = &format
	return r
}

func (r ApiExternalMediaUpdateRequest) ClientTag(clientTag string) ApiExternalMediaUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiExternalMediaUpdateRequest) PartnerId(partnerId int32) ApiExternalMediaUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiExternalMediaUpdateRequest) Body(body ExternalMediaUpdateRequest) ApiExternalMediaUpdateRequest {
	r.body = &body
	return r
}

func (r ApiExternalMediaUpdateRequest) Execute() (*KalturaExternalMediaEntry, *http.Response, error) {
	return r.ApiService.ExternalMediaUpdateExecute(r)
}

/*
ExternalMediaUpdate Method for ExternalMediaUpdate

Update external media entry. Only the properties that were set will be updated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiExternalMediaUpdateRequest
*/
func (a *ExternalMediaApiService) ExternalMediaUpdate(ctx context.Context) ApiExternalMediaUpdateRequest {
	return ApiExternalMediaUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaExternalMediaEntry
func (a *ExternalMediaApiService) ExternalMediaUpdateExecute(r ApiExternalMediaUpdateRequest) (*KalturaExternalMediaEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaExternalMediaEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalMediaApiService.ExternalMediaUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/externalmedia_externalmedia/action/update"

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

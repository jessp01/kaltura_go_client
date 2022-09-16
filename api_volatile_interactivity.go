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


// VolatileInteractivityApiService VolatileInteractivityApi service
type VolatileInteractivityApiService service

type ApiVolatileInteractivityAddRequest struct {
	ctx context.Context
	ApiService *VolatileInteractivityApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *VolatileInteractivityAddRequest
}

func (r ApiVolatileInteractivityAddRequest) Ks(ks string) ApiVolatileInteractivityAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiVolatileInteractivityAddRequest) Format(format int32) ApiVolatileInteractivityAddRequest {
	r.format = &format
	return r
}

func (r ApiVolatileInteractivityAddRequest) ClientTag(clientTag string) ApiVolatileInteractivityAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiVolatileInteractivityAddRequest) PartnerId(partnerId int32) ApiVolatileInteractivityAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiVolatileInteractivityAddRequest) Body(body VolatileInteractivityAddRequest) ApiVolatileInteractivityAddRequest {
	r.body = &body
	return r
}

func (r ApiVolatileInteractivityAddRequest) Execute() (*KalturaVolatileInteractivity, *http.Response, error) {
	return r.ApiService.VolatileInteractivityAddExecute(r)
}

/*
VolatileInteractivityAdd Method for VolatileInteractivityAdd

add a volatile interactivity object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVolatileInteractivityAddRequest
*/
func (a *VolatileInteractivityApiService) VolatileInteractivityAdd(ctx context.Context) ApiVolatileInteractivityAddRequest {
	return ApiVolatileInteractivityAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaVolatileInteractivity
func (a *VolatileInteractivityApiService) VolatileInteractivityAddExecute(r ApiVolatileInteractivityAddRequest) (*KalturaVolatileInteractivity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaVolatileInteractivity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolatileInteractivityApiService.VolatileInteractivityAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/interactivity_volatileinteractivity/action/add"

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

type ApiVolatileInteractivityDeleteRequest struct {
	ctx context.Context
	ApiService *VolatileInteractivityApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *BaseEntryApproveRequest
}

func (r ApiVolatileInteractivityDeleteRequest) Ks(ks string) ApiVolatileInteractivityDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiVolatileInteractivityDeleteRequest) Format(format int32) ApiVolatileInteractivityDeleteRequest {
	r.format = &format
	return r
}

func (r ApiVolatileInteractivityDeleteRequest) ClientTag(clientTag string) ApiVolatileInteractivityDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiVolatileInteractivityDeleteRequest) PartnerId(partnerId int32) ApiVolatileInteractivityDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiVolatileInteractivityDeleteRequest) Body(body BaseEntryApproveRequest) ApiVolatileInteractivityDeleteRequest {
	r.body = &body
	return r
}

func (r ApiVolatileInteractivityDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.VolatileInteractivityDeleteExecute(r)
}

/*
VolatileInteractivityDelete Method for VolatileInteractivityDelete

Delete a volatile interactivity object by entry id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVolatileInteractivityDeleteRequest
*/
func (a *VolatileInteractivityApiService) VolatileInteractivityDelete(ctx context.Context) ApiVolatileInteractivityDeleteRequest {
	return ApiVolatileInteractivityDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *VolatileInteractivityApiService) VolatileInteractivityDeleteExecute(r ApiVolatileInteractivityDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolatileInteractivityApiService.VolatileInteractivityDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/interactivity_volatileinteractivity/action/delete"

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

type ApiVolatileInteractivityGetRequest struct {
	ctx context.Context
	ApiService *VolatileInteractivityApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *BaseEntryApproveRequest
}

func (r ApiVolatileInteractivityGetRequest) Ks(ks string) ApiVolatileInteractivityGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiVolatileInteractivityGetRequest) Format(format int32) ApiVolatileInteractivityGetRequest {
	r.format = &format
	return r
}

func (r ApiVolatileInteractivityGetRequest) ClientTag(clientTag string) ApiVolatileInteractivityGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiVolatileInteractivityGetRequest) PartnerId(partnerId int32) ApiVolatileInteractivityGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiVolatileInteractivityGetRequest) Body(body BaseEntryApproveRequest) ApiVolatileInteractivityGetRequest {
	r.body = &body
	return r
}

func (r ApiVolatileInteractivityGetRequest) Execute() (*KalturaVolatileInteractivity, *http.Response, error) {
	return r.ApiService.VolatileInteractivityGetExecute(r)
}

/*
VolatileInteractivityGet Method for VolatileInteractivityGet

Retrieve a volatile interactivity object by entry id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVolatileInteractivityGetRequest
*/
func (a *VolatileInteractivityApiService) VolatileInteractivityGet(ctx context.Context) ApiVolatileInteractivityGetRequest {
	return ApiVolatileInteractivityGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaVolatileInteractivity
func (a *VolatileInteractivityApiService) VolatileInteractivityGetExecute(r ApiVolatileInteractivityGetRequest) (*KalturaVolatileInteractivity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaVolatileInteractivity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolatileInteractivityApiService.VolatileInteractivityGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/interactivity_volatileinteractivity/action/get"

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

type ApiVolatileInteractivityUpdateRequest struct {
	ctx context.Context
	ApiService *VolatileInteractivityApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *VolatileInteractivityUpdateRequest
}

func (r ApiVolatileInteractivityUpdateRequest) Ks(ks string) ApiVolatileInteractivityUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiVolatileInteractivityUpdateRequest) Format(format int32) ApiVolatileInteractivityUpdateRequest {
	r.format = &format
	return r
}

func (r ApiVolatileInteractivityUpdateRequest) ClientTag(clientTag string) ApiVolatileInteractivityUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiVolatileInteractivityUpdateRequest) PartnerId(partnerId int32) ApiVolatileInteractivityUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiVolatileInteractivityUpdateRequest) Body(body VolatileInteractivityUpdateRequest) ApiVolatileInteractivityUpdateRequest {
	r.body = &body
	return r
}

func (r ApiVolatileInteractivityUpdateRequest) Execute() (*KalturaVolatileInteractivity, *http.Response, error) {
	return r.ApiService.VolatileInteractivityUpdateExecute(r)
}

/*
VolatileInteractivityUpdate Method for VolatileInteractivityUpdate

Update a volatile interactivity object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVolatileInteractivityUpdateRequest
*/
func (a *VolatileInteractivityApiService) VolatileInteractivityUpdate(ctx context.Context) ApiVolatileInteractivityUpdateRequest {
	return ApiVolatileInteractivityUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaVolatileInteractivity
func (a *VolatileInteractivityApiService) VolatileInteractivityUpdateExecute(r ApiVolatileInteractivityUpdateRequest) (*KalturaVolatileInteractivity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaVolatileInteractivity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolatileInteractivityApiService.VolatileInteractivityUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/interactivity_volatileinteractivity/action/update"

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

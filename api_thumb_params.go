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


// ThumbParamsApiService ThumbParamsApi service
type ThumbParamsApiService service

type ApiThumbParamsAddRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ThumbParamsAddRequest
}

func (r ApiThumbParamsAddRequest) Ks(ks string) ApiThumbParamsAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsAddRequest) Format(format int32) ApiThumbParamsAddRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsAddRequest) ClientTag(clientTag string) ApiThumbParamsAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsAddRequest) PartnerId(partnerId int32) ApiThumbParamsAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsAddRequest) Body(body ThumbParamsAddRequest) ApiThumbParamsAddRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsAddRequest) Execute() (*KalturaThumbParams, *http.Response, error) {
	return r.ApiService.ThumbParamsAddExecute(r)
}

/*
ThumbParamsAdd Method for ThumbParamsAdd

Add new Thumb Params

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsAddRequest
*/
func (a *ThumbParamsApiService) ThumbParamsAdd(ctx context.Context) ApiThumbParamsAddRequest {
	return ApiThumbParamsAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaThumbParams
func (a *ThumbParamsApiService) ThumbParamsAddExecute(r ApiThumbParamsAddRequest) (*KalturaThumbParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaThumbParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/add"

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

type ApiThumbParamsDeleteRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiThumbParamsDeleteRequest) Ks(ks string) ApiThumbParamsDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsDeleteRequest) Format(format int32) ApiThumbParamsDeleteRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsDeleteRequest) ClientTag(clientTag string) ApiThumbParamsDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsDeleteRequest) PartnerId(partnerId int32) ApiThumbParamsDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsDeleteRequest) Body(body AccessControlDeleteRequest) ApiThumbParamsDeleteRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ThumbParamsDeleteExecute(r)
}

/*
ThumbParamsDelete Method for ThumbParamsDelete

Delete Thumb Params by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsDeleteRequest
*/
func (a *ThumbParamsApiService) ThumbParamsDelete(ctx context.Context) ApiThumbParamsDeleteRequest {
	return ApiThumbParamsDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ThumbParamsApiService) ThumbParamsDeleteExecute(r ApiThumbParamsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/delete"

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

type ApiThumbParamsGetRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiThumbParamsGetRequest) Ks(ks string) ApiThumbParamsGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsGetRequest) Format(format int32) ApiThumbParamsGetRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsGetRequest) ClientTag(clientTag string) ApiThumbParamsGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsGetRequest) PartnerId(partnerId int32) ApiThumbParamsGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsGetRequest) Body(body AccessControlDeleteRequest) ApiThumbParamsGetRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsGetRequest) Execute() (*KalturaThumbParams, *http.Response, error) {
	return r.ApiService.ThumbParamsGetExecute(r)
}

/*
ThumbParamsGet Method for ThumbParamsGet

Get Thumb Params by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsGetRequest
*/
func (a *ThumbParamsApiService) ThumbParamsGet(ctx context.Context) ApiThumbParamsGetRequest {
	return ApiThumbParamsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaThumbParams
func (a *ThumbParamsApiService) ThumbParamsGetExecute(r ApiThumbParamsGetRequest) (*KalturaThumbParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaThumbParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/get"

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

type ApiThumbParamsGetByConversionProfileIdRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FlavorParamsGetByConversionProfileIdRequest
}

func (r ApiThumbParamsGetByConversionProfileIdRequest) Ks(ks string) ApiThumbParamsGetByConversionProfileIdRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsGetByConversionProfileIdRequest) Format(format int32) ApiThumbParamsGetByConversionProfileIdRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsGetByConversionProfileIdRequest) ClientTag(clientTag string) ApiThumbParamsGetByConversionProfileIdRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsGetByConversionProfileIdRequest) PartnerId(partnerId int32) ApiThumbParamsGetByConversionProfileIdRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsGetByConversionProfileIdRequest) Body(body FlavorParamsGetByConversionProfileIdRequest) ApiThumbParamsGetByConversionProfileIdRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsGetByConversionProfileIdRequest) Execute() ([]KalturaThumbParams, *http.Response, error) {
	return r.ApiService.ThumbParamsGetByConversionProfileIdExecute(r)
}

/*
ThumbParamsGetByConversionProfileId Method for ThumbParamsGetByConversionProfileId

Get Thumb Params by Conversion Profile ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsGetByConversionProfileIdRequest
*/
func (a *ThumbParamsApiService) ThumbParamsGetByConversionProfileId(ctx context.Context) ApiThumbParamsGetByConversionProfileIdRequest {
	return ApiThumbParamsGetByConversionProfileIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []KalturaThumbParams
func (a *ThumbParamsApiService) ThumbParamsGetByConversionProfileIdExecute(r ApiThumbParamsGetByConversionProfileIdRequest) ([]KalturaThumbParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []KalturaThumbParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsGetByConversionProfileId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/getByConversionProfileId"

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

type ApiThumbParamsListRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ThumbParamsListRequest
}

func (r ApiThumbParamsListRequest) Ks(ks string) ApiThumbParamsListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsListRequest) Format(format int32) ApiThumbParamsListRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsListRequest) ClientTag(clientTag string) ApiThumbParamsListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsListRequest) PartnerId(partnerId int32) ApiThumbParamsListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsListRequest) Body(body ThumbParamsListRequest) ApiThumbParamsListRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsListRequest) Execute() (*KalturaThumbParamsListResponse, *http.Response, error) {
	return r.ApiService.ThumbParamsListExecute(r)
}

/*
ThumbParamsList Method for ThumbParamsList

List Thumb Params by filter with paging support (By default - all system default params will be listed too)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsListRequest
*/
func (a *ThumbParamsApiService) ThumbParamsList(ctx context.Context) ApiThumbParamsListRequest {
	return ApiThumbParamsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaThumbParamsListResponse
func (a *ThumbParamsApiService) ThumbParamsListExecute(r ApiThumbParamsListRequest) (*KalturaThumbParamsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaThumbParamsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/list"

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

type ApiThumbParamsUpdateRequest struct {
	ctx context.Context
	ApiService *ThumbParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ThumbParamsUpdateRequest
}

func (r ApiThumbParamsUpdateRequest) Ks(ks string) ApiThumbParamsUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiThumbParamsUpdateRequest) Format(format int32) ApiThumbParamsUpdateRequest {
	r.format = &format
	return r
}

func (r ApiThumbParamsUpdateRequest) ClientTag(clientTag string) ApiThumbParamsUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiThumbParamsUpdateRequest) PartnerId(partnerId int32) ApiThumbParamsUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiThumbParamsUpdateRequest) Body(body ThumbParamsUpdateRequest) ApiThumbParamsUpdateRequest {
	r.body = &body
	return r
}

func (r ApiThumbParamsUpdateRequest) Execute() (*KalturaThumbParams, *http.Response, error) {
	return r.ApiService.ThumbParamsUpdateExecute(r)
}

/*
ThumbParamsUpdate Method for ThumbParamsUpdate

Update Thumb Params by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThumbParamsUpdateRequest
*/
func (a *ThumbParamsApiService) ThumbParamsUpdate(ctx context.Context) ApiThumbParamsUpdateRequest {
	return ApiThumbParamsUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaThumbParams
func (a *ThumbParamsApiService) ThumbParamsUpdateExecute(r ApiThumbParamsUpdateRequest) (*KalturaThumbParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaThumbParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThumbParamsApiService.ThumbParamsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/thumbparams/action/update"

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
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


// ConversionProfileAssetParamsApiService ConversionProfileAssetParamsApi service
type ConversionProfileAssetParamsApiService service

type ApiConversionProfileAssetParamsListRequest struct {
	ctx context.Context
	ApiService *ConversionProfileAssetParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ConversionProfileAssetParamsListRequest
}

func (r ApiConversionProfileAssetParamsListRequest) Ks(ks string) ApiConversionProfileAssetParamsListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiConversionProfileAssetParamsListRequest) Format(format int32) ApiConversionProfileAssetParamsListRequest {
	r.format = &format
	return r
}

func (r ApiConversionProfileAssetParamsListRequest) ClientTag(clientTag string) ApiConversionProfileAssetParamsListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiConversionProfileAssetParamsListRequest) PartnerId(partnerId int32) ApiConversionProfileAssetParamsListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiConversionProfileAssetParamsListRequest) Body(body ConversionProfileAssetParamsListRequest) ApiConversionProfileAssetParamsListRequest {
	r.body = &body
	return r
}

func (r ApiConversionProfileAssetParamsListRequest) Execute() (*KalturaConversionProfileAssetParamsListResponse, *http.Response, error) {
	return r.ApiService.ConversionProfileAssetParamsListExecute(r)
}

/*
ConversionProfileAssetParamsList Method for ConversionProfileAssetParamsList

Lists asset parmas of conversion profile by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConversionProfileAssetParamsListRequest
*/
func (a *ConversionProfileAssetParamsApiService) ConversionProfileAssetParamsList(ctx context.Context) ApiConversionProfileAssetParamsListRequest {
	return ApiConversionProfileAssetParamsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaConversionProfileAssetParamsListResponse
func (a *ConversionProfileAssetParamsApiService) ConversionProfileAssetParamsListExecute(r ApiConversionProfileAssetParamsListRequest) (*KalturaConversionProfileAssetParamsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaConversionProfileAssetParamsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConversionProfileAssetParamsApiService.ConversionProfileAssetParamsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/conversionprofileassetparams/action/list"

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

type ApiConversionProfileAssetParamsUpdateRequest struct {
	ctx context.Context
	ApiService *ConversionProfileAssetParamsApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ConversionProfileAssetParamsUpdateRequest
}

func (r ApiConversionProfileAssetParamsUpdateRequest) Ks(ks string) ApiConversionProfileAssetParamsUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiConversionProfileAssetParamsUpdateRequest) Format(format int32) ApiConversionProfileAssetParamsUpdateRequest {
	r.format = &format
	return r
}

func (r ApiConversionProfileAssetParamsUpdateRequest) ClientTag(clientTag string) ApiConversionProfileAssetParamsUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiConversionProfileAssetParamsUpdateRequest) PartnerId(partnerId int32) ApiConversionProfileAssetParamsUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiConversionProfileAssetParamsUpdateRequest) Body(body ConversionProfileAssetParamsUpdateRequest) ApiConversionProfileAssetParamsUpdateRequest {
	r.body = &body
	return r
}

func (r ApiConversionProfileAssetParamsUpdateRequest) Execute() (*KalturaConversionProfileAssetParams, *http.Response, error) {
	return r.ApiService.ConversionProfileAssetParamsUpdateExecute(r)
}

/*
ConversionProfileAssetParamsUpdate Method for ConversionProfileAssetParamsUpdate

Update asset parmas of conversion profile by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConversionProfileAssetParamsUpdateRequest
*/
func (a *ConversionProfileAssetParamsApiService) ConversionProfileAssetParamsUpdate(ctx context.Context) ApiConversionProfileAssetParamsUpdateRequest {
	return ApiConversionProfileAssetParamsUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaConversionProfileAssetParams
func (a *ConversionProfileAssetParamsApiService) ConversionProfileAssetParamsUpdateExecute(r ApiConversionProfileAssetParamsUpdateRequest) (*KalturaConversionProfileAssetParams, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaConversionProfileAssetParams
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConversionProfileAssetParamsApiService.ConversionProfileAssetParamsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/conversionprofileassetparams/action/update"

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

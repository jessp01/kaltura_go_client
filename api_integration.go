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


// IntegrationApiService IntegrationApi service
type IntegrationApiService service

type ApiIntegrationDispatchRequest struct {
	ctx context.Context
	ApiService *IntegrationApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *IntegrationDispatchRequest
}

func (r ApiIntegrationDispatchRequest) Ks(ks string) ApiIntegrationDispatchRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiIntegrationDispatchRequest) Format(format int32) ApiIntegrationDispatchRequest {
	r.format = &format
	return r
}

func (r ApiIntegrationDispatchRequest) ClientTag(clientTag string) ApiIntegrationDispatchRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiIntegrationDispatchRequest) PartnerId(partnerId int32) ApiIntegrationDispatchRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiIntegrationDispatchRequest) Body(body IntegrationDispatchRequest) ApiIntegrationDispatchRequest {
	r.body = &body
	return r
}

func (r ApiIntegrationDispatchRequest) Execute() (int32, *http.Response, error) {
	return r.ApiService.IntegrationDispatchExecute(r)
}

/*
IntegrationDispatch Method for IntegrationDispatch

Dispatch integration task

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIntegrationDispatchRequest
*/
func (a *IntegrationApiService) IntegrationDispatch(ctx context.Context) ApiIntegrationDispatchRequest {
	return ApiIntegrationDispatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return int32
func (a *IntegrationApiService) IntegrationDispatchExecute(r ApiIntegrationDispatchRequest) (int32, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationApiService.IntegrationDispatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/integration_integration/action/dispatch"

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

type ApiIntegrationNotifyRequest struct {
	ctx context.Context
	ApiService *IntegrationApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiIntegrationNotifyRequest) Ks(ks string) ApiIntegrationNotifyRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiIntegrationNotifyRequest) Format(format int32) ApiIntegrationNotifyRequest {
	r.format = &format
	return r
}

func (r ApiIntegrationNotifyRequest) ClientTag(clientTag string) ApiIntegrationNotifyRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiIntegrationNotifyRequest) PartnerId(partnerId int32) ApiIntegrationNotifyRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiIntegrationNotifyRequest) Body(body AccessControlDeleteRequest) ApiIntegrationNotifyRequest {
	r.body = &body
	return r
}

func (r ApiIntegrationNotifyRequest) Execute() (*http.Response, error) {
	return r.ApiService.IntegrationNotifyExecute(r)
}

/*
IntegrationNotify Method for IntegrationNotify

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIntegrationNotifyRequest
*/
func (a *IntegrationApiService) IntegrationNotify(ctx context.Context) ApiIntegrationNotifyRequest {
	return ApiIntegrationNotifyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *IntegrationApiService) IntegrationNotifyExecute(r ApiIntegrationNotifyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationApiService.IntegrationNotify")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/integration_integration/action/notify"

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
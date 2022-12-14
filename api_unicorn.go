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


// UnicornApiService UnicornApi service
type UnicornApiService service

type ApiUnicornNotifyRequest struct {
	ctx context.Context
	ApiService *UnicornApiService
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

// The format of the response
func (r ApiUnicornNotifyRequest) Format(format int32) ApiUnicornNotifyRequest {
	r.format = &format
	return r
}

func (r ApiUnicornNotifyRequest) ClientTag(clientTag string) ApiUnicornNotifyRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUnicornNotifyRequest) PartnerId(partnerId int32) ApiUnicornNotifyRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUnicornNotifyRequest) Body(body AccessControlDeleteRequest) ApiUnicornNotifyRequest {
	r.body = &body
	return r
}

func (r ApiUnicornNotifyRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnicornNotifyExecute(r)
}

/*
UnicornNotify Method for UnicornNotify

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUnicornNotifyRequest
*/
func (a *UnicornApiService) UnicornNotify(ctx context.Context) ApiUnicornNotifyRequest {
	return ApiUnicornNotifyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UnicornApiService) UnicornNotifyExecute(r ApiUnicornNotifyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UnicornApiService.UnicornNotify")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/unicorndistribution_unicorn/action/notify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

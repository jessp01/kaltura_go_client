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


// TvComApiService TvComApi service
type TvComApiService service

type ApiTvComGetFeedRequest struct {
	ctx context.Context
	ApiService *TvComApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AttUverseGetFeedRequest
}

func (r ApiTvComGetFeedRequest) Ks(ks string) ApiTvComGetFeedRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiTvComGetFeedRequest) Format(format int32) ApiTvComGetFeedRequest {
	r.format = &format
	return r
}

func (r ApiTvComGetFeedRequest) ClientTag(clientTag string) ApiTvComGetFeedRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiTvComGetFeedRequest) PartnerId(partnerId int32) ApiTvComGetFeedRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiTvComGetFeedRequest) Body(body AttUverseGetFeedRequest) ApiTvComGetFeedRequest {
	r.body = &body
	return r
}

func (r ApiTvComGetFeedRequest) Execute() (*http.Response, error) {
	return r.ApiService.TvComGetFeedExecute(r)
}

/*
TvComGetFeed Method for TvComGetFeed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTvComGetFeedRequest
*/
func (a *TvComApiService) TvComGetFeed(ctx context.Context) ApiTvComGetFeedRequest {
	return ApiTvComGetFeedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TvComApiService) TvComGetFeedExecute(r ApiTvComGetFeedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvComApiService.TvComGetFeed")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/tvcomdistribution_tvcom/action/getFeed"

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

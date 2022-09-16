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


// AuditTrailApiService AuditTrailApi service
type AuditTrailApiService service

type ApiAuditTrailAddRequest struct {
	ctx context.Context
	ApiService *AuditTrailApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AuditTrailAddRequest
}

func (r ApiAuditTrailAddRequest) Ks(ks string) ApiAuditTrailAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAuditTrailAddRequest) Format(format int32) ApiAuditTrailAddRequest {
	r.format = &format
	return r
}

func (r ApiAuditTrailAddRequest) ClientTag(clientTag string) ApiAuditTrailAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAuditTrailAddRequest) PartnerId(partnerId int32) ApiAuditTrailAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAuditTrailAddRequest) Body(body AuditTrailAddRequest) ApiAuditTrailAddRequest {
	r.body = &body
	return r
}

func (r ApiAuditTrailAddRequest) Execute() (*KalturaAuditTrail, *http.Response, error) {
	return r.ApiService.AuditTrailAddExecute(r)
}

/*
AuditTrailAdd Method for AuditTrailAdd

Allows you to add an audit trail object and audit trail content associated with Kaltura object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditTrailAddRequest
*/
func (a *AuditTrailApiService) AuditTrailAdd(ctx context.Context) ApiAuditTrailAddRequest {
	return ApiAuditTrailAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAuditTrail
func (a *AuditTrailApiService) AuditTrailAddExecute(r ApiAuditTrailAddRequest) (*KalturaAuditTrail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAuditTrail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditTrailApiService.AuditTrailAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/audit_audittrail/action/add"

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

type ApiAuditTrailGetRequest struct {
	ctx context.Context
	ApiService *AuditTrailApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiAuditTrailGetRequest) Ks(ks string) ApiAuditTrailGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAuditTrailGetRequest) Format(format int32) ApiAuditTrailGetRequest {
	r.format = &format
	return r
}

func (r ApiAuditTrailGetRequest) ClientTag(clientTag string) ApiAuditTrailGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAuditTrailGetRequest) PartnerId(partnerId int32) ApiAuditTrailGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAuditTrailGetRequest) Body(body AccessControlDeleteRequest) ApiAuditTrailGetRequest {
	r.body = &body
	return r
}

func (r ApiAuditTrailGetRequest) Execute() (*KalturaAuditTrail, *http.Response, error) {
	return r.ApiService.AuditTrailGetExecute(r)
}

/*
AuditTrailGet Method for AuditTrailGet

Retrieve an audit trail object by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditTrailGetRequest
*/
func (a *AuditTrailApiService) AuditTrailGet(ctx context.Context) ApiAuditTrailGetRequest {
	return ApiAuditTrailGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAuditTrail
func (a *AuditTrailApiService) AuditTrailGetExecute(r ApiAuditTrailGetRequest) (*KalturaAuditTrail, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAuditTrail
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditTrailApiService.AuditTrailGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/audit_audittrail/action/get"

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

type ApiAuditTrailListRequest struct {
	ctx context.Context
	ApiService *AuditTrailApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AuditTrailListRequest
}

func (r ApiAuditTrailListRequest) Ks(ks string) ApiAuditTrailListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiAuditTrailListRequest) Format(format int32) ApiAuditTrailListRequest {
	r.format = &format
	return r
}

func (r ApiAuditTrailListRequest) ClientTag(clientTag string) ApiAuditTrailListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiAuditTrailListRequest) PartnerId(partnerId int32) ApiAuditTrailListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiAuditTrailListRequest) Body(body AuditTrailListRequest) ApiAuditTrailListRequest {
	r.body = &body
	return r
}

func (r ApiAuditTrailListRequest) Execute() (*KalturaAuditTrailListResponse, *http.Response, error) {
	return r.ApiService.AuditTrailListExecute(r)
}

/*
AuditTrailList Method for AuditTrailList

List audit trail objects by filter and pager

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditTrailListRequest
*/
func (a *AuditTrailApiService) AuditTrailList(ctx context.Context) ApiAuditTrailListRequest {
	return ApiAuditTrailListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaAuditTrailListResponse
func (a *AuditTrailApiService) AuditTrailListExecute(r ApiAuditTrailListRequest) (*KalturaAuditTrailListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaAuditTrailListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditTrailApiService.AuditTrailList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/audit_audittrail/action/list"

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
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


// PollApiService PollApi service
type PollApiService service

type ApiPollAddRequest struct {
	ctx context.Context
	ApiService *PollApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PollAddRequest
}

func (r ApiPollAddRequest) Ks(ks string) ApiPollAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPollAddRequest) Format(format int32) ApiPollAddRequest {
	r.format = &format
	return r
}

func (r ApiPollAddRequest) ClientTag(clientTag string) ApiPollAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPollAddRequest) PartnerId(partnerId int32) ApiPollAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPollAddRequest) Body(body PollAddRequest) ApiPollAddRequest {
	r.body = &body
	return r
}

func (r ApiPollAddRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PollAddExecute(r)
}

/*
PollAdd Method for PollAdd

Add Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPollAddRequest
*/
func (a *PollApiService) PollAdd(ctx context.Context) ApiPollAddRequest {
	return ApiPollAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *PollApiService) PollAddExecute(r ApiPollAddRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PollApiService.PollAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/poll_poll/action/add"

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

type ApiPollGetVoteRequest struct {
	ctx context.Context
	ApiService *PollApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PollGetVoteRequest
}

func (r ApiPollGetVoteRequest) Ks(ks string) ApiPollGetVoteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPollGetVoteRequest) Format(format int32) ApiPollGetVoteRequest {
	r.format = &format
	return r
}

func (r ApiPollGetVoteRequest) ClientTag(clientTag string) ApiPollGetVoteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPollGetVoteRequest) PartnerId(partnerId int32) ApiPollGetVoteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPollGetVoteRequest) Body(body PollGetVoteRequest) ApiPollGetVoteRequest {
	r.body = &body
	return r
}

func (r ApiPollGetVoteRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PollGetVoteExecute(r)
}

/*
PollGetVote Method for PollGetVote

Vote Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPollGetVoteRequest
*/
func (a *PollApiService) PollGetVote(ctx context.Context) ApiPollGetVoteRequest {
	return ApiPollGetVoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *PollApiService) PollGetVoteExecute(r ApiPollGetVoteRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PollApiService.PollGetVote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/poll_poll/action/getVote"

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

type ApiPollGetVotesRequest struct {
	ctx context.Context
	ApiService *PollApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PollGetVotesRequest
}

func (r ApiPollGetVotesRequest) Ks(ks string) ApiPollGetVotesRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPollGetVotesRequest) Format(format int32) ApiPollGetVotesRequest {
	r.format = &format
	return r
}

func (r ApiPollGetVotesRequest) ClientTag(clientTag string) ApiPollGetVotesRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPollGetVotesRequest) PartnerId(partnerId int32) ApiPollGetVotesRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPollGetVotesRequest) Body(body PollGetVotesRequest) ApiPollGetVotesRequest {
	r.body = &body
	return r
}

func (r ApiPollGetVotesRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PollGetVotesExecute(r)
}

/*
PollGetVotes Method for PollGetVotes

Get Votes Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPollGetVotesRequest
*/
func (a *PollApiService) PollGetVotes(ctx context.Context) ApiPollGetVotesRequest {
	return ApiPollGetVotesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *PollApiService) PollGetVotesExecute(r ApiPollGetVotesRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PollApiService.PollGetVotes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/poll_poll/action/getVotes"

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

type ApiPollResetVotesRequest struct {
	ctx context.Context
	ApiService *PollApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PollResetVotesRequest
}

func (r ApiPollResetVotesRequest) Ks(ks string) ApiPollResetVotesRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPollResetVotesRequest) Format(format int32) ApiPollResetVotesRequest {
	r.format = &format
	return r
}

func (r ApiPollResetVotesRequest) ClientTag(clientTag string) ApiPollResetVotesRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPollResetVotesRequest) PartnerId(partnerId int32) ApiPollResetVotesRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPollResetVotesRequest) Body(body PollResetVotesRequest) ApiPollResetVotesRequest {
	r.body = &body
	return r
}

func (r ApiPollResetVotesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PollResetVotesExecute(r)
}

/*
PollResetVotes Method for PollResetVotes

Get resetVotes Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPollResetVotesRequest
*/
func (a *PollApiService) PollResetVotes(ctx context.Context) ApiPollResetVotesRequest {
	return ApiPollResetVotesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PollApiService) PollResetVotesExecute(r ApiPollResetVotesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PollApiService.PollResetVotes")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/poll_poll/action/resetVotes"

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

type ApiPollVoteRequest struct {
	ctx context.Context
	ApiService *PollApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *PollVoteRequest
}

func (r ApiPollVoteRequest) Ks(ks string) ApiPollVoteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPollVoteRequest) Format(format int32) ApiPollVoteRequest {
	r.format = &format
	return r
}

func (r ApiPollVoteRequest) ClientTag(clientTag string) ApiPollVoteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPollVoteRequest) PartnerId(partnerId int32) ApiPollVoteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPollVoteRequest) Body(body PollVoteRequest) ApiPollVoteRequest {
	r.body = &body
	return r
}

func (r ApiPollVoteRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.PollVoteExecute(r)
}

/*
PollVote Method for PollVote

Vote Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPollVoteRequest
*/
func (a *PollApiService) PollVote(ctx context.Context) ApiPollVoteRequest {
	return ApiPollVoteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *PollApiService) PollVoteExecute(r ApiPollVoteRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PollApiService.PollVote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/poll_poll/action/vote"

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

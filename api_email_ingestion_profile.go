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


// EmailIngestionProfileApiService EmailIngestionProfileApi service
type EmailIngestionProfileApiService service

type ApiEmailIngestionProfileAddRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EmailIngestionProfileAddRequest
}

func (r ApiEmailIngestionProfileAddRequest) Ks(ks string) ApiEmailIngestionProfileAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileAddRequest) Format(format int32) ApiEmailIngestionProfileAddRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileAddRequest) ClientTag(clientTag string) ApiEmailIngestionProfileAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileAddRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileAddRequest) Body(body EmailIngestionProfileAddRequest) ApiEmailIngestionProfileAddRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileAddRequest) Execute() (*KalturaEmailIngestionProfile, *http.Response, error) {
	return r.ApiService.EmailIngestionProfileAddExecute(r)
}

/*
EmailIngestionProfileAdd Method for EmailIngestionProfileAdd

EmailIngestionProfile Add action allows you to add a EmailIngestionProfile to Kaltura DB

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileAddRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileAdd(ctx context.Context) ApiEmailIngestionProfileAddRequest {
	return ApiEmailIngestionProfileAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEmailIngestionProfile
func (a *EmailIngestionProfileApiService) EmailIngestionProfileAddExecute(r ApiEmailIngestionProfileAddRequest) (*KalturaEmailIngestionProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEmailIngestionProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/add"

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

type ApiEmailIngestionProfileAddMediaEntryRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EmailIngestionProfileAddMediaEntryRequest
}

func (r ApiEmailIngestionProfileAddMediaEntryRequest) Ks(ks string) ApiEmailIngestionProfileAddMediaEntryRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileAddMediaEntryRequest) Format(format int32) ApiEmailIngestionProfileAddMediaEntryRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileAddMediaEntryRequest) ClientTag(clientTag string) ApiEmailIngestionProfileAddMediaEntryRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileAddMediaEntryRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileAddMediaEntryRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileAddMediaEntryRequest) Body(body EmailIngestionProfileAddMediaEntryRequest) ApiEmailIngestionProfileAddMediaEntryRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileAddMediaEntryRequest) Execute() (*KalturaMediaEntry, *http.Response, error) {
	return r.ApiService.EmailIngestionProfileAddMediaEntryExecute(r)
}

/*
EmailIngestionProfileAddMediaEntry Method for EmailIngestionProfileAddMediaEntry

add KalturaMediaEntry from email ingestion

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileAddMediaEntryRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileAddMediaEntry(ctx context.Context) ApiEmailIngestionProfileAddMediaEntryRequest {
	return ApiEmailIngestionProfileAddMediaEntryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaMediaEntry
func (a *EmailIngestionProfileApiService) EmailIngestionProfileAddMediaEntryExecute(r ApiEmailIngestionProfileAddMediaEntryRequest) (*KalturaMediaEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaMediaEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileAddMediaEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/addMediaEntry"

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

type ApiEmailIngestionProfileDeleteRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiEmailIngestionProfileDeleteRequest) Ks(ks string) ApiEmailIngestionProfileDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileDeleteRequest) Format(format int32) ApiEmailIngestionProfileDeleteRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileDeleteRequest) ClientTag(clientTag string) ApiEmailIngestionProfileDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileDeleteRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileDeleteRequest) Body(body AccessControlDeleteRequest) ApiEmailIngestionProfileDeleteRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EmailIngestionProfileDeleteExecute(r)
}

/*
EmailIngestionProfileDelete Method for EmailIngestionProfileDelete

Delete an existing EmailIngestionProfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileDeleteRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileDelete(ctx context.Context) ApiEmailIngestionProfileDeleteRequest {
	return ApiEmailIngestionProfileDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *EmailIngestionProfileApiService) EmailIngestionProfileDeleteExecute(r ApiEmailIngestionProfileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/delete"

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

type ApiEmailIngestionProfileGetRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiEmailIngestionProfileGetRequest) Ks(ks string) ApiEmailIngestionProfileGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileGetRequest) Format(format int32) ApiEmailIngestionProfileGetRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileGetRequest) ClientTag(clientTag string) ApiEmailIngestionProfileGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileGetRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileGetRequest) Body(body AccessControlDeleteRequest) ApiEmailIngestionProfileGetRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileGetRequest) Execute() (*KalturaEmailIngestionProfile, *http.Response, error) {
	return r.ApiService.EmailIngestionProfileGetExecute(r)
}

/*
EmailIngestionProfileGet Method for EmailIngestionProfileGet

Retrieve a EmailIngestionProfile by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileGetRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileGet(ctx context.Context) ApiEmailIngestionProfileGetRequest {
	return ApiEmailIngestionProfileGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEmailIngestionProfile
func (a *EmailIngestionProfileApiService) EmailIngestionProfileGetExecute(r ApiEmailIngestionProfileGetRequest) (*KalturaEmailIngestionProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEmailIngestionProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/get"

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

type ApiEmailIngestionProfileGetByEmailAddressRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EmailIngestionProfileGetByEmailAddressRequest
}

func (r ApiEmailIngestionProfileGetByEmailAddressRequest) Ks(ks string) ApiEmailIngestionProfileGetByEmailAddressRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileGetByEmailAddressRequest) Format(format int32) ApiEmailIngestionProfileGetByEmailAddressRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileGetByEmailAddressRequest) ClientTag(clientTag string) ApiEmailIngestionProfileGetByEmailAddressRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileGetByEmailAddressRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileGetByEmailAddressRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileGetByEmailAddressRequest) Body(body EmailIngestionProfileGetByEmailAddressRequest) ApiEmailIngestionProfileGetByEmailAddressRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileGetByEmailAddressRequest) Execute() (*KalturaEmailIngestionProfile, *http.Response, error) {
	return r.ApiService.EmailIngestionProfileGetByEmailAddressExecute(r)
}

/*
EmailIngestionProfileGetByEmailAddress Method for EmailIngestionProfileGetByEmailAddress

Retrieve a EmailIngestionProfile by email address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileGetByEmailAddressRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileGetByEmailAddress(ctx context.Context) ApiEmailIngestionProfileGetByEmailAddressRequest {
	return ApiEmailIngestionProfileGetByEmailAddressRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEmailIngestionProfile
func (a *EmailIngestionProfileApiService) EmailIngestionProfileGetByEmailAddressExecute(r ApiEmailIngestionProfileGetByEmailAddressRequest) (*KalturaEmailIngestionProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEmailIngestionProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileGetByEmailAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/getByEmailAddress"

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

type ApiEmailIngestionProfileUpdateRequest struct {
	ctx context.Context
	ApiService *EmailIngestionProfileApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *EmailIngestionProfileUpdateRequest
}

func (r ApiEmailIngestionProfileUpdateRequest) Ks(ks string) ApiEmailIngestionProfileUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiEmailIngestionProfileUpdateRequest) Format(format int32) ApiEmailIngestionProfileUpdateRequest {
	r.format = &format
	return r
}

func (r ApiEmailIngestionProfileUpdateRequest) ClientTag(clientTag string) ApiEmailIngestionProfileUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiEmailIngestionProfileUpdateRequest) PartnerId(partnerId int32) ApiEmailIngestionProfileUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiEmailIngestionProfileUpdateRequest) Body(body EmailIngestionProfileUpdateRequest) ApiEmailIngestionProfileUpdateRequest {
	r.body = &body
	return r
}

func (r ApiEmailIngestionProfileUpdateRequest) Execute() (*KalturaEmailIngestionProfile, *http.Response, error) {
	return r.ApiService.EmailIngestionProfileUpdateExecute(r)
}

/*
EmailIngestionProfileUpdate Method for EmailIngestionProfileUpdate

Update an existing EmailIngestionProfile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmailIngestionProfileUpdateRequest
*/
func (a *EmailIngestionProfileApiService) EmailIngestionProfileUpdate(ctx context.Context) ApiEmailIngestionProfileUpdateRequest {
	return ApiEmailIngestionProfileUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaEmailIngestionProfile
func (a *EmailIngestionProfileApiService) EmailIngestionProfileUpdateExecute(r ApiEmailIngestionProfileUpdateRequest) (*KalturaEmailIngestionProfile, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaEmailIngestionProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmailIngestionProfileApiService.EmailIngestionProfileUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/emailingestionprofile/action/update"

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

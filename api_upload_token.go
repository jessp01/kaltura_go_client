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
	"os"
)


// UploadTokenApiService UploadTokenApi service
type UploadTokenApiService service

type ApiUploadTokenAddRequest struct {
	ctx context.Context
	ApiService *UploadTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UploadTokenAddRequest
}

func (r ApiUploadTokenAddRequest) Ks(ks string) ApiUploadTokenAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadTokenAddRequest) Format(format int32) ApiUploadTokenAddRequest {
	r.format = &format
	return r
}

func (r ApiUploadTokenAddRequest) ClientTag(clientTag string) ApiUploadTokenAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUploadTokenAddRequest) PartnerId(partnerId int32) ApiUploadTokenAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUploadTokenAddRequest) Body(body UploadTokenAddRequest) ApiUploadTokenAddRequest {
	r.body = &body
	return r
}

func (r ApiUploadTokenAddRequest) Execute() (*KalturaUploadToken, *http.Response, error) {
	return r.ApiService.UploadTokenAddExecute(r)
}

/*
UploadTokenAdd Method for UploadTokenAdd

Adds new upload token to upload a file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadTokenAddRequest
*/
func (a *UploadTokenApiService) UploadTokenAdd(ctx context.Context) ApiUploadTokenAddRequest {
	return ApiUploadTokenAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUploadToken
func (a *UploadTokenApiService) UploadTokenAddExecute(r ApiUploadTokenAddRequest) (*KalturaUploadToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUploadToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadTokenApiService.UploadTokenAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uploadtoken/action/add"

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

type ApiUploadTokenDeleteRequest struct {
	ctx context.Context
	ApiService *UploadTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UploadTokenDeleteRequest
}

func (r ApiUploadTokenDeleteRequest) Ks(ks string) ApiUploadTokenDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadTokenDeleteRequest) Format(format int32) ApiUploadTokenDeleteRequest {
	r.format = &format
	return r
}

func (r ApiUploadTokenDeleteRequest) ClientTag(clientTag string) ApiUploadTokenDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUploadTokenDeleteRequest) PartnerId(partnerId int32) ApiUploadTokenDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUploadTokenDeleteRequest) Body(body UploadTokenDeleteRequest) ApiUploadTokenDeleteRequest {
	r.body = &body
	return r
}

func (r ApiUploadTokenDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadTokenDeleteExecute(r)
}

/*
UploadTokenDelete Method for UploadTokenDelete

Deletes the upload token by upload token id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadTokenDeleteRequest
*/
func (a *UploadTokenApiService) UploadTokenDelete(ctx context.Context) ApiUploadTokenDeleteRequest {
	return ApiUploadTokenDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UploadTokenApiService) UploadTokenDeleteExecute(r ApiUploadTokenDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadTokenApiService.UploadTokenDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uploadtoken/action/delete"

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

type ApiUploadTokenGetRequest struct {
	ctx context.Context
	ApiService *UploadTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UploadTokenDeleteRequest
}

func (r ApiUploadTokenGetRequest) Ks(ks string) ApiUploadTokenGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadTokenGetRequest) Format(format int32) ApiUploadTokenGetRequest {
	r.format = &format
	return r
}

func (r ApiUploadTokenGetRequest) ClientTag(clientTag string) ApiUploadTokenGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUploadTokenGetRequest) PartnerId(partnerId int32) ApiUploadTokenGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUploadTokenGetRequest) Body(body UploadTokenDeleteRequest) ApiUploadTokenGetRequest {
	r.body = &body
	return r
}

func (r ApiUploadTokenGetRequest) Execute() (*KalturaUploadToken, *http.Response, error) {
	return r.ApiService.UploadTokenGetExecute(r)
}

/*
UploadTokenGet Method for UploadTokenGet

Get upload token by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadTokenGetRequest
*/
func (a *UploadTokenApiService) UploadTokenGet(ctx context.Context) ApiUploadTokenGetRequest {
	return ApiUploadTokenGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUploadToken
func (a *UploadTokenApiService) UploadTokenGetExecute(r ApiUploadTokenGetRequest) (*KalturaUploadToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUploadToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadTokenApiService.UploadTokenGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uploadtoken/action/get"

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

type ApiUploadTokenListRequest struct {
	ctx context.Context
	ApiService *UploadTokenApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UploadTokenListRequest
}

func (r ApiUploadTokenListRequest) Ks(ks string) ApiUploadTokenListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadTokenListRequest) Format(format int32) ApiUploadTokenListRequest {
	r.format = &format
	return r
}

func (r ApiUploadTokenListRequest) ClientTag(clientTag string) ApiUploadTokenListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUploadTokenListRequest) PartnerId(partnerId int32) ApiUploadTokenListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUploadTokenListRequest) Body(body UploadTokenListRequest) ApiUploadTokenListRequest {
	r.body = &body
	return r
}

func (r ApiUploadTokenListRequest) Execute() (*KalturaUploadTokenListResponse, *http.Response, error) {
	return r.ApiService.UploadTokenListExecute(r)
}

/*
UploadTokenList Method for UploadTokenList

List upload token by filter with pager support. 

When using a user session the service will be restricted to users objects only.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadTokenListRequest
*/
func (a *UploadTokenApiService) UploadTokenList(ctx context.Context) ApiUploadTokenListRequest {
	return ApiUploadTokenListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUploadTokenListResponse
func (a *UploadTokenApiService) UploadTokenListExecute(r ApiUploadTokenListRequest) (*KalturaUploadTokenListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUploadTokenListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadTokenApiService.UploadTokenList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uploadtoken/action/list"

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

type ApiUploadTokenUploadRequest struct {
	ctx context.Context
	ApiService *UploadTokenApiService
	uploadTokenId *string
	fileData **os.File
	ks *string
	format *int32
	resume *bool
	finalChunk *bool
	resumeAt *float32
}

func (r ApiUploadTokenUploadRequest) UploadTokenId(uploadTokenId string) ApiUploadTokenUploadRequest {
	r.uploadTokenId = &uploadTokenId
	return r
}

func (r ApiUploadTokenUploadRequest) FileData(fileData *os.File) ApiUploadTokenUploadRequest {
	r.fileData = &fileData
	return r
}

func (r ApiUploadTokenUploadRequest) Ks(ks string) ApiUploadTokenUploadRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadTokenUploadRequest) Format(format int32) ApiUploadTokenUploadRequest {
	r.format = &format
	return r
}

func (r ApiUploadTokenUploadRequest) Resume(resume bool) ApiUploadTokenUploadRequest {
	r.resume = &resume
	return r
}

func (r ApiUploadTokenUploadRequest) FinalChunk(finalChunk bool) ApiUploadTokenUploadRequest {
	r.finalChunk = &finalChunk
	return r
}

func (r ApiUploadTokenUploadRequest) ResumeAt(resumeAt float32) ApiUploadTokenUploadRequest {
	r.resumeAt = &resumeAt
	return r
}

func (r ApiUploadTokenUploadRequest) Execute() (*KalturaUploadToken, *http.Response, error) {
	return r.ApiService.UploadTokenUploadExecute(r)
}

/*
UploadTokenUpload Method for UploadTokenUpload

Upload a file using the upload token id, returns an error on failure (an exception will be thrown when using one of the Kaltura clients)

Chunks can be uploaded in parallel and they will be appended according to their resumeAt position.

A parallel upload session should have three stages:

1. A single upload with resume=false and finalChunk=false

2. Parallel upload requests each with resume=true,finalChunk=false and the expected resumetAt position.

If a chunk fails to upload it can be re-uploaded.

3. After all of the chunks have been uploaded a final chunk (can be of zero size) should be uploaded 

with resume=true, finalChunk=true and the expected resumeAt position. In case an UPLOAD_TOKEN_CANNOT_MATCH_EXPECTED_SIZE exception

has been returned (indicating not all of the chunks were appended yet) the final request can be retried.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadTokenUploadRequest
*/
func (a *UploadTokenApiService) UploadTokenUpload(ctx context.Context) ApiUploadTokenUploadRequest {
	return ApiUploadTokenUploadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUploadToken
func (a *UploadTokenApiService) UploadTokenUploadExecute(r ApiUploadTokenUploadRequest) (*KalturaUploadToken, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUploadToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadTokenApiService.UploadTokenUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/uploadtoken/action/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.uploadTokenId == nil {
		return localVarReturnValue, nil, reportError("uploadTokenId is required and must be specified")
	}
	if r.fileData == nil {
		return localVarReturnValue, nil, reportError("fileData is required and must be specified")
	}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
	}
	localVarQueryParams.Add("uploadTokenId", parameterToString(*r.uploadTokenId, ""))
	if r.resume != nil {
		localVarQueryParams.Add("resume", parameterToString(*r.resume, ""))
	}
	if r.finalChunk != nil {
		localVarQueryParams.Add("finalChunk", parameterToString(*r.finalChunk, ""))
	}
	if r.resumeAt != nil {
		localVarQueryParams.Add("resumeAt", parameterToString(*r.resumeAt, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileDataLocalVarFormFileName string
	var fileDataLocalVarFileName     string
	var fileDataLocalVarFileBytes    []byte

	fileDataLocalVarFormFileName = "fileData"

	fileDataLocalVarFile := *r.fileData
	if fileDataLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileDataLocalVarFile)
		fileDataLocalVarFileBytes = fbs
		fileDataLocalVarFileName = fileDataLocalVarFile.Name()
		fileDataLocalVarFile.Close()
	}
	formFiles = append(formFiles, formFile{fileBytes: fileDataLocalVarFileBytes, fileName: fileDataLocalVarFileName, formFileName: fileDataLocalVarFormFileName})
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
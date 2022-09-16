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


// FilesyncImportBatchApiService FilesyncImportBatchApi service
type FilesyncImportBatchApiService service

type ApiFilesyncImportBatchExtendFileSyncLockRequest struct {
	ctx context.Context
	ApiService *FilesyncImportBatchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) Ks(ks string) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) Format(format int32) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	r.format = &format
	return r
}

func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) ClientTag(clientTag string) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) PartnerId(partnerId int32) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) Body(body AccessControlDeleteRequest) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	r.body = &body
	return r
}

func (r ApiFilesyncImportBatchExtendFileSyncLockRequest) Execute() (*http.Response, error) {
	return r.ApiService.FilesyncImportBatchExtendFileSyncLockExecute(r)
}

/*
FilesyncImportBatchExtendFileSyncLock Method for FilesyncImportBatchExtendFileSyncLock

batch extendFileSyncLock action extends the expiration of a file sync lock

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilesyncImportBatchExtendFileSyncLockRequest
*/
func (a *FilesyncImportBatchApiService) FilesyncImportBatchExtendFileSyncLock(ctx context.Context) ApiFilesyncImportBatchExtendFileSyncLockRequest {
	return ApiFilesyncImportBatchExtendFileSyncLockRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *FilesyncImportBatchApiService) FilesyncImportBatchExtendFileSyncLockExecute(r ApiFilesyncImportBatchExtendFileSyncLockRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesyncImportBatchApiService.FilesyncImportBatchExtendFileSyncLock")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/multicenters_filesyncimportbatch/action/extendFileSyncLock"

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

type ApiFilesyncImportBatchLockPendingFileSyncsRequest struct {
	ctx context.Context
	ApiService *FilesyncImportBatchApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *FilesyncImportBatchLockPendingFileSyncsRequest
}

func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) Ks(ks string) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) Format(format int32) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	r.format = &format
	return r
}

func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) ClientTag(clientTag string) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) PartnerId(partnerId int32) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) Body(body FilesyncImportBatchLockPendingFileSyncsRequest) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	r.body = &body
	return r
}

func (r ApiFilesyncImportBatchLockPendingFileSyncsRequest) Execute() (*KalturaLockFileSyncsResponse, *http.Response, error) {
	return r.ApiService.FilesyncImportBatchLockPendingFileSyncsExecute(r)
}

/*
FilesyncImportBatchLockPendingFileSyncs Method for FilesyncImportBatchLockPendingFileSyncs

batch lockPendingFileSyncs action locks file syncs for import by the file sync periodic worker

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilesyncImportBatchLockPendingFileSyncsRequest
*/
func (a *FilesyncImportBatchApiService) FilesyncImportBatchLockPendingFileSyncs(ctx context.Context) ApiFilesyncImportBatchLockPendingFileSyncsRequest {
	return ApiFilesyncImportBatchLockPendingFileSyncsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaLockFileSyncsResponse
func (a *FilesyncImportBatchApiService) FilesyncImportBatchLockPendingFileSyncsExecute(r ApiFilesyncImportBatchLockPendingFileSyncsRequest) (*KalturaLockFileSyncsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaLockFileSyncsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesyncImportBatchApiService.FilesyncImportBatchLockPendingFileSyncs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/multicenters_filesyncimportbatch/action/lockPendingFileSyncs"

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
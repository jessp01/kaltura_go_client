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


// UploadApiService UploadApi service
type UploadApiService service

type ApiUploadGetUploadedFileTokenByFileNameRequest struct {
	ctx context.Context
	ApiService *UploadApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *UploadGetUploadedFileTokenByFileNameRequest
}

func (r ApiUploadGetUploadedFileTokenByFileNameRequest) Ks(ks string) ApiUploadGetUploadedFileTokenByFileNameRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadGetUploadedFileTokenByFileNameRequest) Format(format int32) ApiUploadGetUploadedFileTokenByFileNameRequest {
	r.format = &format
	return r
}

func (r ApiUploadGetUploadedFileTokenByFileNameRequest) ClientTag(clientTag string) ApiUploadGetUploadedFileTokenByFileNameRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiUploadGetUploadedFileTokenByFileNameRequest) PartnerId(partnerId int32) ApiUploadGetUploadedFileTokenByFileNameRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiUploadGetUploadedFileTokenByFileNameRequest) Body(body UploadGetUploadedFileTokenByFileNameRequest) ApiUploadGetUploadedFileTokenByFileNameRequest {
	r.body = &body
	return r
}

func (r ApiUploadGetUploadedFileTokenByFileNameRequest) Execute() (*KalturaUploadResponse, *http.Response, error) {
	return r.ApiService.UploadGetUploadedFileTokenByFileNameExecute(r)
}

/*
UploadGetUploadedFileTokenByFileName Method for UploadGetUploadedFileTokenByFileName

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadGetUploadedFileTokenByFileNameRequest
*/
func (a *UploadApiService) UploadGetUploadedFileTokenByFileName(ctx context.Context) ApiUploadGetUploadedFileTokenByFileNameRequest {
	return ApiUploadGetUploadedFileTokenByFileNameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaUploadResponse
func (a *UploadApiService) UploadGetUploadedFileTokenByFileNameExecute(r ApiUploadGetUploadedFileTokenByFileNameRequest) (*KalturaUploadResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaUploadResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadApiService.UploadGetUploadedFileTokenByFileName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/upload/action/getUploadedFileTokenByFileName"

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

type ApiUploadUploadRequest struct {
	ctx context.Context
	ApiService *UploadApiService
	fileData **os.File
	ks *string
	format *int32
}

// The file data
func (r ApiUploadUploadRequest) FileData(fileData *os.File) ApiUploadUploadRequest {
	r.fileData = &fileData
	return r
}

func (r ApiUploadUploadRequest) Ks(ks string) ApiUploadUploadRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiUploadUploadRequest) Format(format int32) ApiUploadUploadRequest {
	r.format = &format
	return r
}

func (r ApiUploadUploadRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.UploadUploadExecute(r)
}

/*
UploadUpload Method for UploadUpload

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUploadUploadRequest
*/
func (a *UploadApiService) UploadUpload(ctx context.Context) ApiUploadUploadRequest {
	return ApiUploadUploadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *UploadApiService) UploadUploadExecute(r ApiUploadUploadRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UploadApiService.UploadUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/upload/action/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fileData == nil {
		return localVarReturnValue, nil, reportError("fileData is required and must be specified")
	}

	if r.ks != nil {
		localVarQueryParams.Add("ks", parameterToString(*r.ks, ""))
	}
	if r.format != nil {
		localVarQueryParams.Add("format", parameterToString(*r.format, ""))
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

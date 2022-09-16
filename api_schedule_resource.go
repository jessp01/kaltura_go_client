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


// ScheduleResourceApiService ScheduleResourceApi service
type ScheduleResourceApiService service

type ApiScheduleResourceAddRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ScheduleResourceAddRequest
}

func (r ApiScheduleResourceAddRequest) Ks(ks string) ApiScheduleResourceAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceAddRequest) Format(format int32) ApiScheduleResourceAddRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceAddRequest) ClientTag(clientTag string) ApiScheduleResourceAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiScheduleResourceAddRequest) PartnerId(partnerId int32) ApiScheduleResourceAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiScheduleResourceAddRequest) Body(body ScheduleResourceAddRequest) ApiScheduleResourceAddRequest {
	r.body = &body
	return r
}

func (r ApiScheduleResourceAddRequest) Execute() (*KalturaScheduleResource, *http.Response, error) {
	return r.ApiService.ScheduleResourceAddExecute(r)
}

/*
ScheduleResourceAdd Method for ScheduleResourceAdd

Allows you to add a new KalturaScheduleResource object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceAddRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceAdd(ctx context.Context) ApiScheduleResourceAddRequest {
	return ApiScheduleResourceAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaScheduleResource
func (a *ScheduleResourceApiService) ScheduleResourceAddExecute(r ApiScheduleResourceAddRequest) (*KalturaScheduleResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaScheduleResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/add"

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

type ApiScheduleResourceAddFromBulkUploadRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	fileData **os.File
	ks *string
	format *int32
	bulkUploadDataObjectType *string
	bulkUploadDataFileName *string
	bulkUploadDataObjectDataObjectType *string
	bulkUploadDataEmailRecipients *string
	bulkUploadDataNumOfErrorObjects *int32
	bulkUploadDataPrivileges *string
	bulkUploadDataColumns *[]map[string]interface{}
	bulkUploadDataProcessObjectId *string
	bulkUploadDataProcessObjectType *string
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) FileData(fileData *os.File) ApiScheduleResourceAddFromBulkUploadRequest {
	r.fileData = &fileData
	return r
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) Ks(ks string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceAddFromBulkUploadRequest) Format(format int32) ApiScheduleResourceAddFromBulkUploadRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataObjectType(bulkUploadDataObjectType string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataObjectType = &bulkUploadDataObjectType
	return r
}

// Friendly name of the file, used to be recognized later in the logs.
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataFileName(bulkUploadDataFileName string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataFileName = &bulkUploadDataFileName
	return r
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataObjectDataObjectType = &bulkUploadDataObjectDataObjectType
	return r
}

// Recipients of the email for bulk upload success/failure
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataEmailRecipients = &bulkUploadDataEmailRecipients
	return r
}

// Number of objects that finished on error status
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects int32) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataNumOfErrorObjects = &bulkUploadDataNumOfErrorObjects
	return r
}

// privileges for the job
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataPrivileges(bulkUploadDataPrivileges string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataPrivileges = &bulkUploadDataPrivileges
	return r
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataColumns(bulkUploadDataColumns []map[string]interface{}) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataColumns = &bulkUploadDataColumns
	return r
}

// The object in process
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataProcessObjectId(bulkUploadDataProcessObjectId string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataProcessObjectId = &bulkUploadDataProcessObjectId
	return r
}

// The type of the object in process
func (r ApiScheduleResourceAddFromBulkUploadRequest) BulkUploadDataProcessObjectType(bulkUploadDataProcessObjectType string) ApiScheduleResourceAddFromBulkUploadRequest {
	r.bulkUploadDataProcessObjectType = &bulkUploadDataProcessObjectType
	return r
}

func (r ApiScheduleResourceAddFromBulkUploadRequest) Execute() (*KalturaBulkUpload, *http.Response, error) {
	return r.ApiService.ScheduleResourceAddFromBulkUploadExecute(r)
}

/*
ScheduleResourceAddFromBulkUpload Method for ScheduleResourceAddFromBulkUpload

Add new bulk upload batch job

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceAddFromBulkUploadRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceAddFromBulkUpload(ctx context.Context) ApiScheduleResourceAddFromBulkUploadRequest {
	return ApiScheduleResourceAddFromBulkUploadRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaBulkUpload
func (a *ScheduleResourceApiService) ScheduleResourceAddFromBulkUploadExecute(r ApiScheduleResourceAddFromBulkUploadRequest) (*KalturaBulkUpload, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaBulkUpload
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceAddFromBulkUpload")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/addFromBulkUpload"

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
	if r.bulkUploadDataObjectType != nil {
		localVarQueryParams.Add("bulkUploadData[objectType]", parameterToString(*r.bulkUploadDataObjectType, ""))
	}
	if r.bulkUploadDataFileName != nil {
		localVarQueryParams.Add("bulkUploadData[fileName]", parameterToString(*r.bulkUploadDataFileName, ""))
	}
	if r.bulkUploadDataObjectDataObjectType != nil {
		localVarQueryParams.Add("bulkUploadData[objectData][objectType]", parameterToString(*r.bulkUploadDataObjectDataObjectType, ""))
	}
	if r.bulkUploadDataEmailRecipients != nil {
		localVarQueryParams.Add("bulkUploadData[emailRecipients]", parameterToString(*r.bulkUploadDataEmailRecipients, ""))
	}
	if r.bulkUploadDataNumOfErrorObjects != nil {
		localVarQueryParams.Add("bulkUploadData[numOfErrorObjects]", parameterToString(*r.bulkUploadDataNumOfErrorObjects, ""))
	}
	if r.bulkUploadDataPrivileges != nil {
		localVarQueryParams.Add("bulkUploadData[privileges]", parameterToString(*r.bulkUploadDataPrivileges, ""))
	}
	if r.bulkUploadDataColumns != nil {
		localVarQueryParams.Add("bulkUploadData[columns]", parameterToString(*r.bulkUploadDataColumns, "csv"))
	}
	if r.bulkUploadDataProcessObjectId != nil {
		localVarQueryParams.Add("bulkUploadData[processObjectId]", parameterToString(*r.bulkUploadDataProcessObjectId, ""))
	}
	if r.bulkUploadDataProcessObjectType != nil {
		localVarQueryParams.Add("bulkUploadData[processObjectType]", parameterToString(*r.bulkUploadDataProcessObjectType, ""))
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

type ApiScheduleResourceDeleteRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ScheduleResourceDeleteRequest
}

func (r ApiScheduleResourceDeleteRequest) Ks(ks string) ApiScheduleResourceDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceDeleteRequest) Format(format int32) ApiScheduleResourceDeleteRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceDeleteRequest) ClientTag(clientTag string) ApiScheduleResourceDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiScheduleResourceDeleteRequest) PartnerId(partnerId int32) ApiScheduleResourceDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiScheduleResourceDeleteRequest) Body(body ScheduleResourceDeleteRequest) ApiScheduleResourceDeleteRequest {
	r.body = &body
	return r
}

func (r ApiScheduleResourceDeleteRequest) Execute() (*KalturaScheduleResource, *http.Response, error) {
	return r.ApiService.ScheduleResourceDeleteExecute(r)
}

/*
ScheduleResourceDelete Method for ScheduleResourceDelete

Mark the KalturaScheduleResource object as deleted

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceDeleteRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceDelete(ctx context.Context) ApiScheduleResourceDeleteRequest {
	return ApiScheduleResourceDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaScheduleResource
func (a *ScheduleResourceApiService) ScheduleResourceDeleteExecute(r ApiScheduleResourceDeleteRequest) (*KalturaScheduleResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaScheduleResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/delete"

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

type ApiScheduleResourceGetRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ScheduleResourceDeleteRequest
}

func (r ApiScheduleResourceGetRequest) Ks(ks string) ApiScheduleResourceGetRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceGetRequest) Format(format int32) ApiScheduleResourceGetRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceGetRequest) ClientTag(clientTag string) ApiScheduleResourceGetRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiScheduleResourceGetRequest) PartnerId(partnerId int32) ApiScheduleResourceGetRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiScheduleResourceGetRequest) Body(body ScheduleResourceDeleteRequest) ApiScheduleResourceGetRequest {
	r.body = &body
	return r
}

func (r ApiScheduleResourceGetRequest) Execute() (*KalturaScheduleResource, *http.Response, error) {
	return r.ApiService.ScheduleResourceGetExecute(r)
}

/*
ScheduleResourceGet Method for ScheduleResourceGet

Retrieve a KalturaScheduleResource object by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceGetRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceGet(ctx context.Context) ApiScheduleResourceGetRequest {
	return ApiScheduleResourceGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaScheduleResource
func (a *ScheduleResourceApiService) ScheduleResourceGetExecute(r ApiScheduleResourceGetRequest) (*KalturaScheduleResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaScheduleResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/get"

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

type ApiScheduleResourceListRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ScheduleResourceListRequest
}

func (r ApiScheduleResourceListRequest) Ks(ks string) ApiScheduleResourceListRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceListRequest) Format(format int32) ApiScheduleResourceListRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceListRequest) ClientTag(clientTag string) ApiScheduleResourceListRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiScheduleResourceListRequest) PartnerId(partnerId int32) ApiScheduleResourceListRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiScheduleResourceListRequest) Body(body ScheduleResourceListRequest) ApiScheduleResourceListRequest {
	r.body = &body
	return r
}

func (r ApiScheduleResourceListRequest) Execute() (*KalturaScheduleResourceListResponse, *http.Response, error) {
	return r.ApiService.ScheduleResourceListExecute(r)
}

/*
ScheduleResourceList Method for ScheduleResourceList

List KalturaScheduleResource objects

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceListRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceList(ctx context.Context) ApiScheduleResourceListRequest {
	return ApiScheduleResourceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaScheduleResourceListResponse
func (a *ScheduleResourceApiService) ScheduleResourceListExecute(r ApiScheduleResourceListRequest) (*KalturaScheduleResourceListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaScheduleResourceListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/list"

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

type ApiScheduleResourceUpdateRequest struct {
	ctx context.Context
	ApiService *ScheduleResourceApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *ScheduleResourceUpdateRequest
}

func (r ApiScheduleResourceUpdateRequest) Ks(ks string) ApiScheduleResourceUpdateRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiScheduleResourceUpdateRequest) Format(format int32) ApiScheduleResourceUpdateRequest {
	r.format = &format
	return r
}

func (r ApiScheduleResourceUpdateRequest) ClientTag(clientTag string) ApiScheduleResourceUpdateRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiScheduleResourceUpdateRequest) PartnerId(partnerId int32) ApiScheduleResourceUpdateRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiScheduleResourceUpdateRequest) Body(body ScheduleResourceUpdateRequest) ApiScheduleResourceUpdateRequest {
	r.body = &body
	return r
}

func (r ApiScheduleResourceUpdateRequest) Execute() (*KalturaScheduleResource, *http.Response, error) {
	return r.ApiService.ScheduleResourceUpdateExecute(r)
}

/*
ScheduleResourceUpdate Method for ScheduleResourceUpdate

Update an existing KalturaScheduleResource object

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiScheduleResourceUpdateRequest
*/
func (a *ScheduleResourceApiService) ScheduleResourceUpdate(ctx context.Context) ApiScheduleResourceUpdateRequest {
	return ApiScheduleResourceUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaScheduleResource
func (a *ScheduleResourceApiService) ScheduleResourceUpdateExecute(r ApiScheduleResourceUpdateRequest) (*KalturaScheduleResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaScheduleResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduleResourceApiService.ScheduleResourceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/schedule_scheduleresource/action/update"

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
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


// DrmLicenseAccessApiService DrmLicenseAccessApi service
type DrmLicenseAccessApiService service

type ApiDrmLicenseAccessGetAccessRequest struct {
	ctx context.Context
	ApiService *DrmLicenseAccessApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *DrmLicenseAccessGetAccessRequest
}

func (r ApiDrmLicenseAccessGetAccessRequest) Ks(ks string) ApiDrmLicenseAccessGetAccessRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiDrmLicenseAccessGetAccessRequest) Format(format int32) ApiDrmLicenseAccessGetAccessRequest {
	r.format = &format
	return r
}

func (r ApiDrmLicenseAccessGetAccessRequest) ClientTag(clientTag string) ApiDrmLicenseAccessGetAccessRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiDrmLicenseAccessGetAccessRequest) PartnerId(partnerId int32) ApiDrmLicenseAccessGetAccessRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiDrmLicenseAccessGetAccessRequest) Body(body DrmLicenseAccessGetAccessRequest) ApiDrmLicenseAccessGetAccessRequest {
	r.body = &body
	return r
}

func (r ApiDrmLicenseAccessGetAccessRequest) Execute() (*KalturaDrmLicenseAccessDetails, *http.Response, error) {
	return r.ApiService.DrmLicenseAccessGetAccessExecute(r)
}

/*
DrmLicenseAccessGetAccess Method for DrmLicenseAccessGetAccess

getAccessAction
     input: flavor ids, drmProvider
     Get Access Action

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDrmLicenseAccessGetAccessRequest
*/
func (a *DrmLicenseAccessApiService) DrmLicenseAccessGetAccess(ctx context.Context) ApiDrmLicenseAccessGetAccessRequest {
	return ApiDrmLicenseAccessGetAccessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaDrmLicenseAccessDetails
func (a *DrmLicenseAccessApiService) DrmLicenseAccessGetAccessExecute(r ApiDrmLicenseAccessGetAccessRequest) (*KalturaDrmLicenseAccessDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaDrmLicenseAccessDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrmLicenseAccessApiService.DrmLicenseAccessGetAccess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/drm_drmlicenseaccess/action/getAccess"

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

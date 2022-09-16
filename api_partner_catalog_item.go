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


// PartnerCatalogItemApiService PartnerCatalogItemApi service
type PartnerCatalogItemApiService service

type ApiPartnerCatalogItemAddRequest struct {
	ctx context.Context
	ApiService *PartnerCatalogItemApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiPartnerCatalogItemAddRequest) Ks(ks string) ApiPartnerCatalogItemAddRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPartnerCatalogItemAddRequest) Format(format int32) ApiPartnerCatalogItemAddRequest {
	r.format = &format
	return r
}

func (r ApiPartnerCatalogItemAddRequest) ClientTag(clientTag string) ApiPartnerCatalogItemAddRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPartnerCatalogItemAddRequest) PartnerId(partnerId int32) ApiPartnerCatalogItemAddRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPartnerCatalogItemAddRequest) Body(body AccessControlDeleteRequest) ApiPartnerCatalogItemAddRequest {
	r.body = &body
	return r
}

func (r ApiPartnerCatalogItemAddRequest) Execute() (*KalturaVendorCatalogItem, *http.Response, error) {
	return r.ApiService.PartnerCatalogItemAddExecute(r)
}

/*
PartnerCatalogItemAdd Method for PartnerCatalogItemAdd

Assign existing catalogItem to specific account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPartnerCatalogItemAddRequest
*/
func (a *PartnerCatalogItemApiService) PartnerCatalogItemAdd(ctx context.Context) ApiPartnerCatalogItemAddRequest {
	return ApiPartnerCatalogItemAddRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KalturaVendorCatalogItem
func (a *PartnerCatalogItemApiService) PartnerCatalogItemAddExecute(r ApiPartnerCatalogItemAddRequest) (*KalturaVendorCatalogItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KalturaVendorCatalogItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerCatalogItemApiService.PartnerCatalogItemAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_partnercatalogitem/action/add"

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

type ApiPartnerCatalogItemDeleteRequest struct {
	ctx context.Context
	ApiService *PartnerCatalogItemApiService
	ks *string
	format *int32
	clientTag *string
	partnerId *int32
	body *AccessControlDeleteRequest
}

func (r ApiPartnerCatalogItemDeleteRequest) Ks(ks string) ApiPartnerCatalogItemDeleteRequest {
	r.ks = &ks
	return r
}

// The format of the response
func (r ApiPartnerCatalogItemDeleteRequest) Format(format int32) ApiPartnerCatalogItemDeleteRequest {
	r.format = &format
	return r
}

func (r ApiPartnerCatalogItemDeleteRequest) ClientTag(clientTag string) ApiPartnerCatalogItemDeleteRequest {
	r.clientTag = &clientTag
	return r
}

func (r ApiPartnerCatalogItemDeleteRequest) PartnerId(partnerId int32) ApiPartnerCatalogItemDeleteRequest {
	r.partnerId = &partnerId
	return r
}

func (r ApiPartnerCatalogItemDeleteRequest) Body(body AccessControlDeleteRequest) ApiPartnerCatalogItemDeleteRequest {
	r.body = &body
	return r
}

func (r ApiPartnerCatalogItemDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PartnerCatalogItemDeleteExecute(r)
}

/*
PartnerCatalogItemDelete Method for PartnerCatalogItemDelete

Remove existing catalogItem from specific account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPartnerCatalogItemDeleteRequest
*/
func (a *PartnerCatalogItemApiService) PartnerCatalogItemDelete(ctx context.Context) ApiPartnerCatalogItemDeleteRequest {
	return ApiPartnerCatalogItemDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PartnerCatalogItemApiService) PartnerCatalogItemDeleteExecute(r ApiPartnerCatalogItemDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PartnerCatalogItemApiService.PartnerCatalogItemDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/service/reach_partnercatalogitem/action/delete"

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

# \MetadataBatchApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataBatchGetExclusiveTransformMetadataJobs**](MetadataBatchApi.md#MetadataBatchGetExclusiveTransformMetadataJobs) | **Post** /service/metadata_metadatabatch/action/getExclusiveTransformMetadataJobs | 
[**MetadataBatchGetTransformMetadataObjects**](MetadataBatchApi.md#MetadataBatchGetTransformMetadataObjects) | **Post** /service/metadata_metadatabatch/action/getTransformMetadataObjects | 



## MetadataBatchGetExclusiveTransformMetadataJobs

> []KalturaBatchJob MetadataBatchGetExclusiveTransformMetadataJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)
    clientTag := "clientTag_example" // string |  (optional) (default to "devkcom")
    partnerId := int32(56) // int32 |  (optional)
    body := *openapiclient.NewMetadataBatchGetExclusiveTransformMetadataJobsRequest(*openapiclient.NewKalturaExclusiveLockKey(), int32(123), int32(123)) // MetadataBatchGetExclusiveTransformMetadataJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataBatchApi.MetadataBatchGetExclusiveTransformMetadataJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataBatchApi.MetadataBatchGetExclusiveTransformMetadataJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataBatchGetExclusiveTransformMetadataJobs`: []KalturaBatchJob
    fmt.Fprintf(os.Stdout, "Response from `MetadataBatchApi.MetadataBatchGetExclusiveTransformMetadataJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataBatchGetExclusiveTransformMetadataJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataBatchGetExclusiveTransformMetadataJobsRequest**](MetadataBatchGetExclusiveTransformMetadataJobsRequest.md) |  | 

### Return type

[**[]KalturaBatchJob**](KalturaBatchJob.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataBatchGetTransformMetadataObjects

> KalturaTransformMetadataResponse MetadataBatchGetTransformMetadataObjects(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)
    clientTag := "clientTag_example" // string |  (optional) (default to "devkcom")
    partnerId := int32(56) // int32 |  (optional)
    body := *openapiclient.NewMetadataBatchGetTransformMetadataObjectsRequest(int32(123), int32(123), int32(123)) // MetadataBatchGetTransformMetadataObjectsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataBatchApi.MetadataBatchGetTransformMetadataObjects(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataBatchApi.MetadataBatchGetTransformMetadataObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataBatchGetTransformMetadataObjects`: KalturaTransformMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataBatchApi.MetadataBatchGetTransformMetadataObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataBatchGetTransformMetadataObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataBatchGetTransformMetadataObjectsRequest**](MetadataBatchGetTransformMetadataObjectsRequest.md) |  | 

### Return type

[**KalturaTransformMetadataResponse**](KalturaTransformMetadataResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


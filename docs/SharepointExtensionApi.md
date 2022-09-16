# \SharepointExtensionApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharepointExtensionIsVersionSupported**](SharepointExtensionApi.md#SharepointExtensionIsVersionSupported) | **Post** /service/kalturasharepointextension_sharepointextension/action/isVersionSupported | 
[**SharepointExtensionListUiconfs**](SharepointExtensionApi.md#SharepointExtensionListUiconfs) | **Post** /service/kalturasharepointextension_sharepointextension/action/listUiconfs | 



## SharepointExtensionIsVersionSupported

> bool SharepointExtensionIsVersionSupported(ctx).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    format := int32(56) // int32 | The format of the response (optional)
    clientTag := "clientTag_example" // string |  (optional) (default to "devkcom")
    partnerId := int32(56) // int32 |  (optional)
    body := *openapiclient.NewSharepointExtensionIsVersionSupportedRequest(int32(123), int32(123), int32(123)) // SharepointExtensionIsVersionSupportedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharepointExtensionApi.SharepointExtensionIsVersionSupported(context.Background()).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharepointExtensionApi.SharepointExtensionIsVersionSupported``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharepointExtensionIsVersionSupported`: bool
    fmt.Fprintf(os.Stdout, "Response from `SharepointExtensionApi.SharepointExtensionIsVersionSupported`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSharepointExtensionIsVersionSupportedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SharepointExtensionIsVersionSupportedRequest**](SharepointExtensionIsVersionSupportedRequest.md) |  | 

### Return type

**bool**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharepointExtensionListUiconfs

> KalturaUiConfListResponse SharepointExtensionListUiconfs(ctx).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    format := int32(56) // int32 | The format of the response (optional)
    clientTag := "clientTag_example" // string |  (optional) (default to "devkcom")
    partnerId := int32(56) // int32 |  (optional)
    batchCleanExclusiveJobsRequest := *openapiclient.NewBatchCleanExclusiveJobsRequest() // BatchCleanExclusiveJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SharepointExtensionApi.SharepointExtensionListUiconfs(context.Background()).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SharepointExtensionApi.SharepointExtensionListUiconfs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SharepointExtensionListUiconfs`: KalturaUiConfListResponse
    fmt.Fprintf(os.Stdout, "Response from `SharepointExtensionApi.SharepointExtensionListUiconfs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSharepointExtensionListUiconfsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

[**KalturaUiConfListResponse**](KalturaUiConfListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


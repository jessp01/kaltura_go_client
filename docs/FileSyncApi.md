# \FileSyncApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileSyncDeleteLocalFileSyncs**](FileSyncApi.md#FileSyncDeleteLocalFileSyncs) | **Post** /service/filesync_filesync/action/deleteLocalFileSyncs | 
[**FileSyncList**](FileSyncApi.md#FileSyncList) | **Post** /service/filesync_filesync/action/list | 
[**FileSyncLockFileSyncs**](FileSyncApi.md#FileSyncLockFileSyncs) | **Post** /service/filesync_filesync/action/lockFileSyncs | 
[**FileSyncUpdate**](FileSyncApi.md#FileSyncUpdate) | **Post** /service/filesync_filesync/action/update | 



## FileSyncDeleteLocalFileSyncs

> KalturaFileSyncListResponse FileSyncDeleteLocalFileSyncs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewFileSyncDeleteLocalFileSyncsRequest(*openapiclient.NewKalturaFileSyncFilter(), int32(123), int32(123), int32(123), int32(123)) // FileSyncDeleteLocalFileSyncsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSyncApi.FileSyncDeleteLocalFileSyncs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSyncApi.FileSyncDeleteLocalFileSyncs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileSyncDeleteLocalFileSyncs`: KalturaFileSyncListResponse
    fmt.Fprintf(os.Stdout, "Response from `FileSyncApi.FileSyncDeleteLocalFileSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFileSyncDeleteLocalFileSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**FileSyncDeleteLocalFileSyncsRequest**](FileSyncDeleteLocalFileSyncsRequest.md) |  | 

### Return type

[**KalturaFileSyncListResponse**](KalturaFileSyncListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSyncList

> KalturaFileSyncListResponse FileSyncList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewFileSyncListRequest() // FileSyncListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSyncApi.FileSyncList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSyncApi.FileSyncList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileSyncList`: KalturaFileSyncListResponse
    fmt.Fprintf(os.Stdout, "Response from `FileSyncApi.FileSyncList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFileSyncListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**FileSyncListRequest**](FileSyncListRequest.md) |  | 

### Return type

[**KalturaFileSyncListResponse**](KalturaFileSyncListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSyncLockFileSyncs

> KalturaLockFileSyncsResponse FileSyncLockFileSyncs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewFileSyncLockFileSyncsRequest(*openapiclient.NewKalturaFileSyncFilter(), int32(123), int32(123), int32(123), int32(123), int32(123)) // FileSyncLockFileSyncsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSyncApi.FileSyncLockFileSyncs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSyncApi.FileSyncLockFileSyncs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileSyncLockFileSyncs`: KalturaLockFileSyncsResponse
    fmt.Fprintf(os.Stdout, "Response from `FileSyncApi.FileSyncLockFileSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFileSyncLockFileSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**FileSyncLockFileSyncsRequest**](FileSyncLockFileSyncsRequest.md) |  | 

### Return type

[**KalturaLockFileSyncsResponse**](KalturaLockFileSyncsResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FileSyncUpdate

> KalturaFileSync FileSyncUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewFileSyncUpdateRequest(*openapiclient.NewKalturaFileSync(), int32(123)) // FileSyncUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileSyncApi.FileSyncUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileSyncApi.FileSyncUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FileSyncUpdate`: KalturaFileSync
    fmt.Fprintf(os.Stdout, "Response from `FileSyncApi.FileSyncUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFileSyncUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**FileSyncUpdateRequest**](FileSyncUpdateRequest.md) |  | 

### Return type

[**KalturaFileSync**](KalturaFileSync.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \FilesyncImportBatchApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesyncImportBatchExtendFileSyncLock**](FilesyncImportBatchApi.md#FilesyncImportBatchExtendFileSyncLock) | **Post** /service/multicenters_filesyncimportbatch/action/extendFileSyncLock | 
[**FilesyncImportBatchLockPendingFileSyncs**](FilesyncImportBatchApi.md#FilesyncImportBatchLockPendingFileSyncs) | **Post** /service/multicenters_filesyncimportbatch/action/lockPendingFileSyncs | 



## FilesyncImportBatchExtendFileSyncLock

> FilesyncImportBatchExtendFileSyncLock(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAccessControlDeleteRequest(int32(123)) // AccessControlDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesyncImportBatchApi.FilesyncImportBatchExtendFileSyncLock(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesyncImportBatchApi.FilesyncImportBatchExtendFileSyncLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesyncImportBatchExtendFileSyncLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesyncImportBatchLockPendingFileSyncs

> KalturaLockFileSyncsResponse FilesyncImportBatchLockPendingFileSyncs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewFilesyncImportBatchLockPendingFileSyncsRequest(*openapiclient.NewKalturaFileSyncFilter(), int32(123), int32(123), int32(123)) // FilesyncImportBatchLockPendingFileSyncsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesyncImportBatchApi.FilesyncImportBatchLockPendingFileSyncs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesyncImportBatchApi.FilesyncImportBatchLockPendingFileSyncs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FilesyncImportBatchLockPendingFileSyncs`: KalturaLockFileSyncsResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesyncImportBatchApi.FilesyncImportBatchLockPendingFileSyncs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFilesyncImportBatchLockPendingFileSyncsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**FilesyncImportBatchLockPendingFileSyncsRequest**](FilesyncImportBatchLockPendingFileSyncsRequest.md) |  | 

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


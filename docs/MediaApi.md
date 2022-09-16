# \MediaApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MediaAdd**](MediaApi.md#MediaAdd) | **Post** /service/media/action/add | 
[**MediaAddContent**](MediaApi.md#MediaAddContent) | **Post** /service/media/action/addContent | 
[**MediaAddFromBulk**](MediaApi.md#MediaAddFromBulk) | **Post** /service/media/action/addFromBulk | 
[**MediaAddFromEntry**](MediaApi.md#MediaAddFromEntry) | **Post** /service/media/action/addFromEntry | 
[**MediaAddFromFlavorAsset**](MediaApi.md#MediaAddFromFlavorAsset) | **Post** /service/media/action/addFromFlavorAsset | 
[**MediaAddFromRecordedWebcam**](MediaApi.md#MediaAddFromRecordedWebcam) | **Post** /service/media/action/addFromRecordedWebcam | 
[**MediaAddFromSearchResult**](MediaApi.md#MediaAddFromSearchResult) | **Post** /service/media/action/addFromSearchResult | 
[**MediaAddFromUploadedFile**](MediaApi.md#MediaAddFromUploadedFile) | **Post** /service/media/action/addFromUploadedFile | 
[**MediaAddFromUrl**](MediaApi.md#MediaAddFromUrl) | **Post** /service/media/action/addFromUrl | 
[**MediaAnonymousRank**](MediaApi.md#MediaAnonymousRank) | **Post** /service/media/action/anonymousRank | 
[**MediaApprove**](MediaApi.md#MediaApprove) | **Post** /service/media/action/approve | 
[**MediaApproveReplace**](MediaApi.md#MediaApproveReplace) | **Post** /service/media/action/approveReplace | 
[**MediaBulkUploadAdd**](MediaApi.md#MediaBulkUploadAdd) | **Post** /service/media/action/bulkUploadAdd | 
[**MediaCancelReplace**](MediaApi.md#MediaCancelReplace) | **Post** /service/media/action/cancelReplace | 
[**MediaConvert**](MediaApi.md#MediaConvert) | **Post** /service/media/action/convert | 
[**MediaCount**](MediaApi.md#MediaCount) | **Post** /service/media/action/count | 
[**MediaDelete**](MediaApi.md#MediaDelete) | **Post** /service/media/action/delete | 
[**MediaExportToCsv**](MediaApi.md#MediaExportToCsv) | **Post** /service/media/action/exportToCsv | 
[**MediaFlag**](MediaApi.md#MediaFlag) | **Post** /service/media/action/flag | 
[**MediaGet**](MediaApi.md#MediaGet) | **Post** /service/media/action/get | 
[**MediaGetMrss**](MediaApi.md#MediaGetMrss) | **Post** /service/media/action/getMrss | 
[**MediaGetVolumeMap**](MediaApi.md#MediaGetVolumeMap) | **Post** /service/media/action/getVolumeMap | 
[**MediaList**](MediaApi.md#MediaList) | **Post** /service/media/action/list | 
[**MediaListFlags**](MediaApi.md#MediaListFlags) | **Post** /service/media/action/listFlags | 
[**MediaReject**](MediaApi.md#MediaReject) | **Post** /service/media/action/reject | 
[**MediaRequestConversion**](MediaApi.md#MediaRequestConversion) | **Post** /service/media/action/requestConversion | 
[**MediaUpdate**](MediaApi.md#MediaUpdate) | **Post** /service/media/action/update | 
[**MediaUpdateContent**](MediaApi.md#MediaUpdateContent) | **Post** /service/media/action/updateContent | 
[**MediaUpdateThumbnail**](MediaApi.md#MediaUpdateThumbnail) | **Post** /service/media/action/updateThumbnail | 
[**MediaUpdateThumbnailFromSourceEntry**](MediaApi.md#MediaUpdateThumbnailFromSourceEntry) | **Post** /service/media/action/updateThumbnailFromSourceEntry | 
[**MediaUpdateThumbnailFromUrl**](MediaApi.md#MediaUpdateThumbnailFromUrl) | **Post** /service/media/action/updateThumbnailFromUrl | 
[**MediaUpdateThumbnailJpeg**](MediaApi.md#MediaUpdateThumbnailJpeg) | **Post** /service/media/action/updateThumbnailJpeg | 
[**MediaUpload**](MediaApi.md#MediaUpload) | **Post** /service/media/action/upload | 



## MediaAdd

> KalturaMediaEntry MediaAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddRequest(*openapiclient.NewKalturaMediaEntry()) // MediaAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAdd`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddRequest**](MediaAddRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddContent

> KalturaMediaEntry MediaAddContent(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddContentRequest("EntryId_example") // MediaAddContentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddContent(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddContent`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddContentRequest**](MediaAddContentRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromBulk

> KalturaMediaEntry MediaAddFromBulk(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromBulkRequest(int32(123), *openapiclient.NewKalturaMediaEntry(), "Url_example") // MediaAddFromBulkRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromBulk(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromBulk`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromBulkRequest**](MediaAddFromBulkRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromEntry

> KalturaMediaEntry MediaAddFromEntry(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromEntryRequest("SourceEntryId_example") // MediaAddFromEntryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromEntry(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromEntry`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromEntryRequest**](MediaAddFromEntryRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromFlavorAsset

> KalturaMediaEntry MediaAddFromFlavorAsset(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromFlavorAssetRequest("SourceFlavorAssetId_example") // MediaAddFromFlavorAssetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromFlavorAsset(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromFlavorAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromFlavorAsset`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromFlavorAsset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromFlavorAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromFlavorAssetRequest**](MediaAddFromFlavorAssetRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromRecordedWebcam

> KalturaMediaEntry MediaAddFromRecordedWebcam(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromRecordedWebcamRequest(*openapiclient.NewKalturaMediaEntry(), "WebcamTokenId_example") // MediaAddFromRecordedWebcamRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromRecordedWebcam(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromRecordedWebcam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromRecordedWebcam`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromRecordedWebcam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromRecordedWebcamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromRecordedWebcamRequest**](MediaAddFromRecordedWebcamRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromSearchResult

> KalturaMediaEntry MediaAddFromSearchResult(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromSearchResultRequest() // MediaAddFromSearchResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromSearchResult(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromSearchResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromSearchResult`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromSearchResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromSearchResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromSearchResultRequest**](MediaAddFromSearchResultRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromUploadedFile

> KalturaMediaEntry MediaAddFromUploadedFile(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromUploadedFileRequest(*openapiclient.NewKalturaMediaEntry(), "UploadTokenId_example") // MediaAddFromUploadedFileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromUploadedFile(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromUploadedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromUploadedFile`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromUploadedFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromUploadedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromUploadedFileRequest**](MediaAddFromUploadedFileRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAddFromUrl

> KalturaMediaEntry MediaAddFromUrl(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaAddFromUrlRequest(*openapiclient.NewKalturaMediaEntry(), "Url_example") // MediaAddFromUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAddFromUrl(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAddFromUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaAddFromUrl`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaAddFromUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAddFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaAddFromUrlRequest**](MediaAddFromUrlRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaAnonymousRank

> MediaAnonymousRank(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryAnonymousRankRequest("EntryId_example", int32(123)) // BaseEntryAnonymousRankRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaAnonymousRank(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaAnonymousRank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaAnonymousRankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryAnonymousRankRequest**](BaseEntryAnonymousRankRequest.md) |  | 

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


## MediaApprove

> MediaApprove(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaApprove(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

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


## MediaApproveReplace

> KalturaMediaEntry MediaApproveReplace(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaApproveReplace(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaApproveReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaApproveReplace`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaApproveReplace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaApproveReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaBulkUploadAdd

> KalturaBulkUpload MediaBulkUploadAdd(ctx).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadEntryDataObjectType(bulkUploadEntryDataObjectType).BulkUploadEntryDataConversionProfileId(bulkUploadEntryDataConversionProfileId).Execute()





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
    fileData := os.NewFile(1234, "some_file") // *os.File | 
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)
    bulkUploadDataObjectType := "bulkUploadDataObjectType_example" // string |  (optional)
    bulkUploadDataFileName := "bulkUploadDataFileName_example" // string | Friendly name of the file, used to be recognized later in the logs. (optional)
    bulkUploadDataObjectDataObjectType := "bulkUploadDataObjectDataObjectType_example" // string |  (optional)
    bulkUploadDataEmailRecipients := "bulkUploadDataEmailRecipients_example" // string | Recipients of the email for bulk upload success/failure (optional)
    bulkUploadDataNumOfErrorObjects := int32(56) // int32 | Number of objects that finished on error status (optional)
    bulkUploadDataPrivileges := "bulkUploadDataPrivileges_example" // string | privileges for the job (optional)
    bulkUploadEntryDataObjectType := "bulkUploadEntryDataObjectType_example" // string |  (optional)
    bulkUploadEntryDataConversionProfileId := int32(56) // int32 | Selected profile id for all bulk entries (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaBulkUploadAdd(context.Background()).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadEntryDataObjectType(bulkUploadEntryDataObjectType).BulkUploadEntryDataConversionProfileId(bulkUploadEntryDataConversionProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaBulkUploadAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaBulkUploadAdd`: KalturaBulkUpload
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaBulkUploadAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaBulkUploadAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileData** | ***os.File** |  | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **bulkUploadDataObjectType** | **string** |  | 
 **bulkUploadDataFileName** | **string** | Friendly name of the file, used to be recognized later in the logs. | 
 **bulkUploadDataObjectDataObjectType** | **string** |  | 
 **bulkUploadDataEmailRecipients** | **string** | Recipients of the email for bulk upload success/failure | 
 **bulkUploadDataNumOfErrorObjects** | **int32** | Number of objects that finished on error status | 
 **bulkUploadDataPrivileges** | **string** | privileges for the job | 
 **bulkUploadEntryDataObjectType** | **string** |  | 
 **bulkUploadEntryDataConversionProfileId** | **int32** | Selected profile id for all bulk entries | 

### Return type

[**KalturaBulkUpload**](KalturaBulkUpload.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaCancelReplace

> KalturaMediaEntry MediaCancelReplace(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaCancelReplace(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaCancelReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaCancelReplace`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaCancelReplace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaCancelReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaConvert

> int32 MediaConvert(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewDocumentsConvertRequest("EntryId_example") // DocumentsConvertRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaConvert(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaConvert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaConvert`: int32
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaConvert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**DocumentsConvertRequest**](DocumentsConvertRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaCount

> int32 MediaCount(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaCountRequest() // MediaCountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaCount(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaCount`: int32
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaCountRequest**](MediaCountRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaDelete

> MediaDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

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


## MediaExportToCsv

> string MediaExportToCsv(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaExportToCsvRequest(*openapiclient.NewKalturaMediaEsearchExportToCsvJobData()) // MediaExportToCsvRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaExportToCsv(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaExportToCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaExportToCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaExportToCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaExportToCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaExportToCsvRequest**](MediaExportToCsvRequest.md) |  | 

### Return type

**string**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaFlag

> MediaFlag(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryFlagRequest(*openapiclient.NewKalturaModerationFlag()) // BaseEntryFlagRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaFlag(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaFlag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryFlagRequest**](BaseEntryFlagRequest.md) |  | 

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


## MediaGet

> KalturaMediaEntry MediaGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaGetRequest("EntryId_example") // MediaGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaGet`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaGetRequest**](MediaGetRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaGetMrss

> string MediaGetMrss(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaGetMrssRequest("EntryId_example") // MediaGetMrssRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaGetMrss(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaGetMrss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaGetMrss`: string
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaGetMrss`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaGetMrssRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaGetMrssRequest**](MediaGetMrssRequest.md) |  | 

### Return type

**string**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaGetVolumeMap

> string MediaGetVolumeMap(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaGetVolumeMap(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaGetVolumeMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaGetVolumeMap`: string
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaGetVolumeMap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaGetVolumeMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

### Return type

**string**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaList

> KalturaMediaListResponse MediaList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaListRequest() // MediaListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaList`: KalturaMediaListResponse
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaListRequest**](MediaListRequest.md) |  | 

### Return type

[**KalturaMediaListResponse**](KalturaMediaListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaListFlags

> KalturaModerationFlagListResponse MediaListFlags(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryListFlagsRequest("EntryId_example") // BaseEntryListFlagsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaListFlags(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaListFlags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaListFlags`: KalturaModerationFlagListResponse
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaListFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaListFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryListFlagsRequest**](BaseEntryListFlagsRequest.md) |  | 

### Return type

[**KalturaModerationFlagListResponse**](KalturaModerationFlagListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaReject

> MediaReject(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryApproveRequest("EntryId_example") // BaseEntryApproveRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaReject(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaReject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryApproveRequest**](BaseEntryApproveRequest.md) |  | 

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


## MediaRequestConversion

> int32 MediaRequestConversion(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaRequestConversionRequest("EntryId_example", "FileFormat_example") // MediaRequestConversionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaRequestConversion(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaRequestConversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaRequestConversion`: int32
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaRequestConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaRequestConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaRequestConversionRequest**](MediaRequestConversionRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdate

> KalturaMediaEntry MediaUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaUpdateRequest("EntryId_example", *openapiclient.NewKalturaMediaEntry()) // MediaUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdate`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaUpdateRequest**](MediaUpdateRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdateContent

> KalturaMediaEntry MediaUpdateContent(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryUpdateContentRequest("EntryId_example", *openapiclient.NewKalturaResource()) // BaseEntryUpdateContentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdateContent(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdateContent`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdateContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryUpdateContentRequest**](BaseEntryUpdateContentRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdateThumbnail

> KalturaMediaEntry MediaUpdateThumbnail(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaUpdateThumbnailRequest("EntryId_example", int32(123)) // MediaUpdateThumbnailRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdateThumbnail(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdateThumbnail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdateThumbnail`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdateThumbnail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaUpdateThumbnailRequest**](MediaUpdateThumbnailRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdateThumbnailFromSourceEntry

> KalturaMediaEntry MediaUpdateThumbnailFromSourceEntry(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMediaUpdateThumbnailFromSourceEntryRequest("EntryId_example", "SourceEntryId_example", int32(123)) // MediaUpdateThumbnailFromSourceEntryRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdateThumbnailFromSourceEntry(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdateThumbnailFromSourceEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdateThumbnailFromSourceEntry`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdateThumbnailFromSourceEntry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateThumbnailFromSourceEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MediaUpdateThumbnailFromSourceEntryRequest**](MediaUpdateThumbnailFromSourceEntryRequest.md) |  | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdateThumbnailFromUrl

> KalturaBaseEntry MediaUpdateThumbnailFromUrl(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBaseEntryUpdateThumbnailFromUrlRequest("EntryId_example", "Url_example") // BaseEntryUpdateThumbnailFromUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdateThumbnailFromUrl(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdateThumbnailFromUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdateThumbnailFromUrl`: KalturaBaseEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdateThumbnailFromUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateThumbnailFromUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BaseEntryUpdateThumbnailFromUrlRequest**](BaseEntryUpdateThumbnailFromUrlRequest.md) |  | 

### Return type

[**KalturaBaseEntry**](KalturaBaseEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpdateThumbnailJpeg

> KalturaMediaEntry MediaUpdateThumbnailJpeg(ctx).EntryId(entryId).FileData(fileData).Ks(ks).Format(format).Execute()





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
    entryId := "entryId_example" // string | Media entry id
    fileData := os.NewFile(1234, "some_file") // *os.File | Jpeg file data
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpdateThumbnailJpeg(context.Background()).EntryId(entryId).FileData(fileData).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpdateThumbnailJpeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpdateThumbnailJpeg`: KalturaMediaEntry
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpdateThumbnailJpeg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUpdateThumbnailJpegRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entryId** | **string** | Media entry id | 
 **fileData** | ***os.File** | Jpeg file data | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaMediaEntry**](KalturaMediaEntry.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MediaUpload

> string MediaUpload(ctx).FileData(fileData).Ks(ks).Format(format).Execute()





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
    fileData := os.NewFile(1234, "some_file") // *os.File | The file data
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MediaApi.MediaUpload(context.Background()).FileData(fileData).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MediaApi.MediaUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MediaUpload`: string
    fmt.Fprintf(os.Stdout, "Response from `MediaApi.MediaUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMediaUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileData** | ***os.File** | The file data | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

**string**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


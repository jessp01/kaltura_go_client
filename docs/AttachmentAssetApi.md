# \AttachmentAssetApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachmentAssetAdd**](AttachmentAssetApi.md#AttachmentAssetAdd) | **Post** /service/attachment_attachmentasset/action/add | 
[**AttachmentAssetDelete**](AttachmentAssetApi.md#AttachmentAssetDelete) | **Post** /service/attachment_attachmentasset/action/delete | 
[**AttachmentAssetGet**](AttachmentAssetApi.md#AttachmentAssetGet) | **Post** /service/attachment_attachmentasset/action/get | 
[**AttachmentAssetGetRemotePaths**](AttachmentAssetApi.md#AttachmentAssetGetRemotePaths) | **Post** /service/attachment_attachmentasset/action/getRemotePaths | 
[**AttachmentAssetGetUrl**](AttachmentAssetApi.md#AttachmentAssetGetUrl) | **Post** /service/attachment_attachmentasset/action/getUrl | 
[**AttachmentAssetList**](AttachmentAssetApi.md#AttachmentAssetList) | **Post** /service/attachment_attachmentasset/action/list | 
[**AttachmentAssetServe**](AttachmentAssetApi.md#AttachmentAssetServe) | **Post** /service/attachment_attachmentasset/action/serve | 
[**AttachmentAssetSetContent**](AttachmentAssetApi.md#AttachmentAssetSetContent) | **Post** /service/attachment_attachmentasset/action/setContent | 
[**AttachmentAssetUpdate**](AttachmentAssetApi.md#AttachmentAssetUpdate) | **Post** /service/attachment_attachmentasset/action/update | 



## AttachmentAssetAdd

> KalturaAttachmentAsset AttachmentAssetAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetAddRequest(*openapiclient.NewKalturaAttachmentAsset(), "EntryId_example") // AttachmentAssetAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetAdd`: KalturaAttachmentAsset
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetAddRequest**](AttachmentAssetAddRequest.md) |  | 

### Return type

[**KalturaAttachmentAsset**](KalturaAttachmentAsset.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentAssetDelete

> AttachmentAssetDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewAttachmentAssetDeleteRequest("AttachmentAssetId_example") // AttachmentAssetDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetDeleteRequest**](AttachmentAssetDeleteRequest.md) |  | 

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


## AttachmentAssetGet

> KalturaAttachmentAsset AttachmentAssetGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewAttachmentAssetDeleteRequest("AttachmentAssetId_example") // AttachmentAssetDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetGet`: KalturaAttachmentAsset
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetDeleteRequest**](AttachmentAssetDeleteRequest.md) |  | 

### Return type

[**KalturaAttachmentAsset**](KalturaAttachmentAsset.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentAssetGetRemotePaths

> KalturaRemotePathListResponse AttachmentAssetGetRemotePaths(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAnnotationDeleteRequest("Id_example") // AnnotationDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetGetRemotePaths(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetGetRemotePaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetGetRemotePaths`: KalturaRemotePathListResponse
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetGetRemotePaths`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetGetRemotePathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AnnotationDeleteRequest**](AnnotationDeleteRequest.md) |  | 

### Return type

[**KalturaRemotePathListResponse**](KalturaRemotePathListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentAssetGetUrl

> string AttachmentAssetGetUrl(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetGetUrlRequest("Id_example") // AttachmentAssetGetUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetGetUrl(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetGetUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetGetUrl`: string
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetGetUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetGetUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetGetUrlRequest**](AttachmentAssetGetUrlRequest.md) |  | 

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


## AttachmentAssetList

> KalturaAttachmentAssetListResponse AttachmentAssetList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetListRequest() // AttachmentAssetListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetList`: KalturaAttachmentAssetListResponse
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetListRequest**](AttachmentAssetListRequest.md) |  | 

### Return type

[**KalturaAttachmentAssetListResponse**](KalturaAttachmentAssetListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentAssetServe

> string AttachmentAssetServe(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetServeRequest("AttachmentAssetId_example") // AttachmentAssetServeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetServe(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetServe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetServe`: string
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetServe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetServeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetServeRequest**](AttachmentAssetServeRequest.md) |  | 

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


## AttachmentAssetSetContent

> KalturaAttachmentAsset AttachmentAssetSetContent(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetSetContentRequest(*openapiclient.NewKalturaContentResource(), "Id_example") // AttachmentAssetSetContentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetSetContent(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetSetContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetSetContent`: KalturaAttachmentAsset
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetSetContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetSetContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetSetContentRequest**](AttachmentAssetSetContentRequest.md) |  | 

### Return type

[**KalturaAttachmentAsset**](KalturaAttachmentAsset.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachmentAssetUpdate

> KalturaAttachmentAsset AttachmentAssetUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAttachmentAssetUpdateRequest(*openapiclient.NewKalturaAttachmentAsset(), "Id_example") // AttachmentAssetUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttachmentAssetApi.AttachmentAssetUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttachmentAssetApi.AttachmentAssetUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachmentAssetUpdate`: KalturaAttachmentAsset
    fmt.Fprintf(os.Stdout, "Response from `AttachmentAssetApi.AttachmentAssetUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAttachmentAssetUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AttachmentAssetUpdateRequest**](AttachmentAssetUpdateRequest.md) |  | 

### Return type

[**KalturaAttachmentAsset**](KalturaAttachmentAsset.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SyndicationFeedApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyndicationFeedAdd**](SyndicationFeedApi.md#SyndicationFeedAdd) | **Post** /service/syndicationfeed/action/add | 
[**SyndicationFeedDelete**](SyndicationFeedApi.md#SyndicationFeedDelete) | **Post** /service/syndicationfeed/action/delete | 
[**SyndicationFeedGet**](SyndicationFeedApi.md#SyndicationFeedGet) | **Post** /service/syndicationfeed/action/get | 
[**SyndicationFeedGetEntryCount**](SyndicationFeedApi.md#SyndicationFeedGetEntryCount) | **Post** /service/syndicationfeed/action/getEntryCount | 
[**SyndicationFeedList**](SyndicationFeedApi.md#SyndicationFeedList) | **Post** /service/syndicationfeed/action/list | 
[**SyndicationFeedRequestConversion**](SyndicationFeedApi.md#SyndicationFeedRequestConversion) | **Post** /service/syndicationfeed/action/requestConversion | 
[**SyndicationFeedUpdate**](SyndicationFeedApi.md#SyndicationFeedUpdate) | **Post** /service/syndicationfeed/action/update | 



## SyndicationFeedAdd

> KalturaBaseSyndicationFeed SyndicationFeedAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewSyndicationFeedAddRequest(*openapiclient.NewKalturaBaseSyndicationFeed()) // SyndicationFeedAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedAdd`: KalturaBaseSyndicationFeed
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SyndicationFeedAddRequest**](SyndicationFeedAddRequest.md) |  | 

### Return type

[**KalturaBaseSyndicationFeed**](KalturaBaseSyndicationFeed.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyndicationFeedDelete

> SyndicationFeedDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AnnotationDeleteRequest**](AnnotationDeleteRequest.md) |  | 

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


## SyndicationFeedGet

> KalturaBaseSyndicationFeed SyndicationFeedGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedGet`: KalturaBaseSyndicationFeed
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AnnotationDeleteRequest**](AnnotationDeleteRequest.md) |  | 

### Return type

[**KalturaBaseSyndicationFeed**](KalturaBaseSyndicationFeed.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyndicationFeedGetEntryCount

> KalturaSyndicationFeedEntryCount SyndicationFeedGetEntryCount(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewSyndicationFeedGetEntryCountRequest("FeedId_example") // SyndicationFeedGetEntryCountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedGetEntryCount(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedGetEntryCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedGetEntryCount`: KalturaSyndicationFeedEntryCount
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedGetEntryCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedGetEntryCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SyndicationFeedGetEntryCountRequest**](SyndicationFeedGetEntryCountRequest.md) |  | 

### Return type

[**KalturaSyndicationFeedEntryCount**](KalturaSyndicationFeedEntryCount.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyndicationFeedList

> KalturaBaseSyndicationFeedListResponse SyndicationFeedList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewSyndicationFeedListRequest() // SyndicationFeedListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedList`: KalturaBaseSyndicationFeedListResponse
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SyndicationFeedListRequest**](SyndicationFeedListRequest.md) |  | 

### Return type

[**KalturaBaseSyndicationFeedListResponse**](KalturaBaseSyndicationFeedListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyndicationFeedRequestConversion

> string SyndicationFeedRequestConversion(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewSyndicationFeedGetEntryCountRequest("FeedId_example") // SyndicationFeedGetEntryCountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedRequestConversion(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedRequestConversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedRequestConversion`: string
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedRequestConversion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedRequestConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SyndicationFeedGetEntryCountRequest**](SyndicationFeedGetEntryCountRequest.md) |  | 

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


## SyndicationFeedUpdate

> KalturaBaseSyndicationFeed SyndicationFeedUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewSyndicationFeedUpdateRequest("Id_example", *openapiclient.NewKalturaBaseSyndicationFeed()) // SyndicationFeedUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyndicationFeedApi.SyndicationFeedUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyndicationFeedApi.SyndicationFeedUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SyndicationFeedUpdate`: KalturaBaseSyndicationFeed
    fmt.Fprintf(os.Stdout, "Response from `SyndicationFeedApi.SyndicationFeedUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyndicationFeedUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**SyndicationFeedUpdateRequest**](SyndicationFeedUpdateRequest.md) |  | 

### Return type

[**KalturaBaseSyndicationFeed**](KalturaBaseSyndicationFeed.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


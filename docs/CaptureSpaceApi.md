# \CaptureSpaceApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaptureSpaceClientUpdates**](CaptureSpaceApi.md#CaptureSpaceClientUpdates) | **Post** /service/capturespace_capturespace/action/clientUpdates | 
[**CaptureSpaceServeInstall**](CaptureSpaceApi.md#CaptureSpaceServeInstall) | **Post** /service/capturespace_capturespace/action/serveInstall | 
[**CaptureSpaceServeUpdate**](CaptureSpaceApi.md#CaptureSpaceServeUpdate) | **Post** /service/capturespace_capturespace/action/serveUpdate | 



## CaptureSpaceClientUpdates

> KalturaCaptureSpaceUpdateResponse CaptureSpaceClientUpdates(ctx).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewCaptureSpaceClientUpdatesRequest("Os_example", "Version_example") // CaptureSpaceClientUpdatesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptureSpaceApi.CaptureSpaceClientUpdates(context.Background()).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptureSpaceApi.CaptureSpaceClientUpdates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CaptureSpaceClientUpdates`: KalturaCaptureSpaceUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `CaptureSpaceApi.CaptureSpaceClientUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaptureSpaceClientUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**CaptureSpaceClientUpdatesRequest**](CaptureSpaceClientUpdatesRequest.md) |  | 

### Return type

[**KalturaCaptureSpaceUpdateResponse**](KalturaCaptureSpaceUpdateResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CaptureSpaceServeInstall

> string CaptureSpaceServeInstall(ctx).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewCaptureSpaceServeInstallRequest("Os_example") // CaptureSpaceServeInstallRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptureSpaceApi.CaptureSpaceServeInstall(context.Background()).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptureSpaceApi.CaptureSpaceServeInstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CaptureSpaceServeInstall`: string
    fmt.Fprintf(os.Stdout, "Response from `CaptureSpaceApi.CaptureSpaceServeInstall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaptureSpaceServeInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**CaptureSpaceServeInstallRequest**](CaptureSpaceServeInstallRequest.md) |  | 

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


## CaptureSpaceServeUpdate

> string CaptureSpaceServeUpdate(ctx).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewCaptureSpaceServeUpdateRequest("Os_example", "Version_example") // CaptureSpaceServeUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CaptureSpaceApi.CaptureSpaceServeUpdate(context.Background()).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CaptureSpaceApi.CaptureSpaceServeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CaptureSpaceServeUpdate`: string
    fmt.Fprintf(os.Stdout, "Response from `CaptureSpaceApi.CaptureSpaceServeUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCaptureSpaceServeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**CaptureSpaceServeUpdateRequest**](CaptureSpaceServeUpdateRequest.md) |  | 

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


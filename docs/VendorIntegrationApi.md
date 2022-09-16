# \VendorIntegrationApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VendorIntegrationAdd**](VendorIntegrationApi.md#VendorIntegrationAdd) | **Post** /service/vendor_vendorintegration/action/add | 
[**VendorIntegrationDelete**](VendorIntegrationApi.md#VendorIntegrationDelete) | **Post** /service/vendor_vendorintegration/action/delete | 
[**VendorIntegrationGet**](VendorIntegrationApi.md#VendorIntegrationGet) | **Post** /service/vendor_vendorintegration/action/get | 
[**VendorIntegrationUpdate**](VendorIntegrationApi.md#VendorIntegrationUpdate) | **Post** /service/vendor_vendorintegration/action/update | 
[**VendorIntegrationUpdateStatus**](VendorIntegrationApi.md#VendorIntegrationUpdateStatus) | **Post** /service/vendor_vendorintegration/action/updateStatus | 



## VendorIntegrationAdd

> KalturaIntegrationSetting VendorIntegrationAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorIntegrationAddRequest(*openapiclient.NewKalturaIntegrationSetting(), "RemoteId_example") // VendorIntegrationAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorIntegrationApi.VendorIntegrationAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorIntegrationApi.VendorIntegrationAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorIntegrationAdd`: KalturaIntegrationSetting
    fmt.Fprintf(os.Stdout, "Response from `VendorIntegrationApi.VendorIntegrationAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorIntegrationAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorIntegrationAddRequest**](VendorIntegrationAddRequest.md) |  | 

### Return type

[**KalturaIntegrationSetting**](KalturaIntegrationSetting.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorIntegrationDelete

> KalturaIntegrationSetting VendorIntegrationDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorIntegrationDeleteRequest(int32(123)) // VendorIntegrationDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorIntegrationApi.VendorIntegrationDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorIntegrationApi.VendorIntegrationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorIntegrationDelete`: KalturaIntegrationSetting
    fmt.Fprintf(os.Stdout, "Response from `VendorIntegrationApi.VendorIntegrationDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorIntegrationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorIntegrationDeleteRequest**](VendorIntegrationDeleteRequest.md) |  | 

### Return type

[**KalturaIntegrationSetting**](KalturaIntegrationSetting.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorIntegrationGet

> KalturaIntegrationSetting VendorIntegrationGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorIntegrationDeleteRequest(int32(123)) // VendorIntegrationDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorIntegrationApi.VendorIntegrationGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorIntegrationApi.VendorIntegrationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorIntegrationGet`: KalturaIntegrationSetting
    fmt.Fprintf(os.Stdout, "Response from `VendorIntegrationApi.VendorIntegrationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorIntegrationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorIntegrationDeleteRequest**](VendorIntegrationDeleteRequest.md) |  | 

### Return type

[**KalturaIntegrationSetting**](KalturaIntegrationSetting.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorIntegrationUpdate

> KalturaIntegrationSetting VendorIntegrationUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorIntegrationUpdateRequest(int32(123), *openapiclient.NewKalturaIntegrationSetting()) // VendorIntegrationUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorIntegrationApi.VendorIntegrationUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorIntegrationApi.VendorIntegrationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorIntegrationUpdate`: KalturaIntegrationSetting
    fmt.Fprintf(os.Stdout, "Response from `VendorIntegrationApi.VendorIntegrationUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorIntegrationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorIntegrationUpdateRequest**](VendorIntegrationUpdateRequest.md) |  | 

### Return type

[**KalturaIntegrationSetting**](KalturaIntegrationSetting.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorIntegrationUpdateStatus

> KalturaIntegrationSetting VendorIntegrationUpdateStatus(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorIntegrationUpdateStatusRequest(int32(123), int32(123)) // VendorIntegrationUpdateStatusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorIntegrationApi.VendorIntegrationUpdateStatus(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorIntegrationApi.VendorIntegrationUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorIntegrationUpdateStatus`: KalturaIntegrationSetting
    fmt.Fprintf(os.Stdout, "Response from `VendorIntegrationApi.VendorIntegrationUpdateStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorIntegrationUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorIntegrationUpdateStatusRequest**](VendorIntegrationUpdateStatusRequest.md) |  | 

### Return type

[**KalturaIntegrationSetting**](KalturaIntegrationSetting.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


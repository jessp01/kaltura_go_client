# \PlayReadyDrmApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlayReadyDrmGenerateKey**](PlayReadyDrmApi.md#PlayReadyDrmGenerateKey) | **Post** /service/playready_playreadydrm/action/generateKey | 
[**PlayReadyDrmGetContentKeys**](PlayReadyDrmApi.md#PlayReadyDrmGetContentKeys) | **Post** /service/playready_playreadydrm/action/getContentKeys | 
[**PlayReadyDrmGetEntryContentKey**](PlayReadyDrmApi.md#PlayReadyDrmGetEntryContentKey) | **Post** /service/playready_playreadydrm/action/getEntryContentKey | 
[**PlayReadyDrmGetLicenseDetails**](PlayReadyDrmApi.md#PlayReadyDrmGetLicenseDetails) | **Post** /service/playready_playreadydrm/action/getLicenseDetails | 



## PlayReadyDrmGenerateKey

> KalturaPlayReadyContentKey PlayReadyDrmGenerateKey(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    batchCleanExclusiveJobsRequest := *openapiclient.NewBatchCleanExclusiveJobsRequest() // BatchCleanExclusiveJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayReadyDrmApi.PlayReadyDrmGenerateKey(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayReadyDrmApi.PlayReadyDrmGenerateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayReadyDrmGenerateKey`: KalturaPlayReadyContentKey
    fmt.Fprintf(os.Stdout, "Response from `PlayReadyDrmApi.PlayReadyDrmGenerateKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayReadyDrmGenerateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

[**KalturaPlayReadyContentKey**](KalturaPlayReadyContentKey.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayReadyDrmGetContentKeys

> []KalturaPlayReadyContentKey PlayReadyDrmGetContentKeys(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewPlayReadyDrmGetContentKeysRequest("KeyIds_example") // PlayReadyDrmGetContentKeysRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayReadyDrmApi.PlayReadyDrmGetContentKeys(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayReadyDrmApi.PlayReadyDrmGetContentKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayReadyDrmGetContentKeys`: []KalturaPlayReadyContentKey
    fmt.Fprintf(os.Stdout, "Response from `PlayReadyDrmApi.PlayReadyDrmGetContentKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayReadyDrmGetContentKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**PlayReadyDrmGetContentKeysRequest**](PlayReadyDrmGetContentKeysRequest.md) |  | 

### Return type

[**[]KalturaPlayReadyContentKey**](KalturaPlayReadyContentKey.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayReadyDrmGetEntryContentKey

> KalturaPlayReadyContentKey PlayReadyDrmGetEntryContentKey(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewPlayReadyDrmGetEntryContentKeyRequest("EntryId_example") // PlayReadyDrmGetEntryContentKeyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayReadyDrmApi.PlayReadyDrmGetEntryContentKey(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayReadyDrmApi.PlayReadyDrmGetEntryContentKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayReadyDrmGetEntryContentKey`: KalturaPlayReadyContentKey
    fmt.Fprintf(os.Stdout, "Response from `PlayReadyDrmApi.PlayReadyDrmGetEntryContentKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayReadyDrmGetEntryContentKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**PlayReadyDrmGetEntryContentKeyRequest**](PlayReadyDrmGetEntryContentKeyRequest.md) |  | 

### Return type

[**KalturaPlayReadyContentKey**](KalturaPlayReadyContentKey.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayReadyDrmGetLicenseDetails

> KalturaPlayReadyLicenseDetails PlayReadyDrmGetLicenseDetails(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewPlayReadyDrmGetLicenseDetailsRequest("DeviceId_example", int32(123), "KeyId_example") // PlayReadyDrmGetLicenseDetailsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlayReadyDrmApi.PlayReadyDrmGetLicenseDetails(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlayReadyDrmApi.PlayReadyDrmGetLicenseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayReadyDrmGetLicenseDetails`: KalturaPlayReadyLicenseDetails
    fmt.Fprintf(os.Stdout, "Response from `PlayReadyDrmApi.PlayReadyDrmGetLicenseDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlayReadyDrmGetLicenseDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**PlayReadyDrmGetLicenseDetailsRequest**](PlayReadyDrmGetLicenseDetailsRequest.md) |  | 

### Return type

[**KalturaPlayReadyLicenseDetails**](KalturaPlayReadyLicenseDetails.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


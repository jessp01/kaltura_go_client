# \ConversionProfileAssetParamsApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversionProfileAssetParamsList**](ConversionProfileAssetParamsApi.md#ConversionProfileAssetParamsList) | **Post** /service/conversionprofileassetparams/action/list | 
[**ConversionProfileAssetParamsUpdate**](ConversionProfileAssetParamsApi.md#ConversionProfileAssetParamsUpdate) | **Post** /service/conversionprofileassetparams/action/update | 



## ConversionProfileAssetParamsList

> KalturaConversionProfileAssetParamsListResponse ConversionProfileAssetParamsList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewConversionProfileAssetParamsListRequest() // ConversionProfileAssetParamsListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConversionProfileAssetParamsApi.ConversionProfileAssetParamsList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConversionProfileAssetParamsApi.ConversionProfileAssetParamsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConversionProfileAssetParamsList`: KalturaConversionProfileAssetParamsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConversionProfileAssetParamsApi.ConversionProfileAssetParamsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConversionProfileAssetParamsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ConversionProfileAssetParamsListRequest**](ConversionProfileAssetParamsListRequest.md) |  | 

### Return type

[**KalturaConversionProfileAssetParamsListResponse**](KalturaConversionProfileAssetParamsListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConversionProfileAssetParamsUpdate

> KalturaConversionProfileAssetParams ConversionProfileAssetParamsUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewConversionProfileAssetParamsUpdateRequest(int32(123), *openapiclient.NewKalturaConversionProfileAssetParams(), int32(123)) // ConversionProfileAssetParamsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConversionProfileAssetParamsApi.ConversionProfileAssetParamsUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConversionProfileAssetParamsApi.ConversionProfileAssetParamsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConversionProfileAssetParamsUpdate`: KalturaConversionProfileAssetParams
    fmt.Fprintf(os.Stdout, "Response from `ConversionProfileAssetParamsApi.ConversionProfileAssetParamsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConversionProfileAssetParamsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ConversionProfileAssetParamsUpdateRequest**](ConversionProfileAssetParamsUpdateRequest.md) |  | 

### Return type

[**KalturaConversionProfileAssetParams**](KalturaConversionProfileAssetParams.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


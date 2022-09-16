# \UploadTokenApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadTokenAdd**](UploadTokenApi.md#UploadTokenAdd) | **Post** /service/uploadtoken/action/add | 
[**UploadTokenDelete**](UploadTokenApi.md#UploadTokenDelete) | **Post** /service/uploadtoken/action/delete | 
[**UploadTokenGet**](UploadTokenApi.md#UploadTokenGet) | **Post** /service/uploadtoken/action/get | 
[**UploadTokenList**](UploadTokenApi.md#UploadTokenList) | **Post** /service/uploadtoken/action/list | 
[**UploadTokenUpload**](UploadTokenApi.md#UploadTokenUpload) | **Post** /service/uploadtoken/action/upload | 



## UploadTokenAdd

> KalturaUploadToken UploadTokenAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewUploadTokenAddRequest() // UploadTokenAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadTokenApi.UploadTokenAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokenApi.UploadTokenAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTokenAdd`: KalturaUploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokenApi.UploadTokenAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTokenAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**UploadTokenAddRequest**](UploadTokenAddRequest.md) |  | 

### Return type

[**KalturaUploadToken**](KalturaUploadToken.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTokenDelete

> UploadTokenDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewUploadTokenDeleteRequest("UploadTokenId_example") // UploadTokenDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadTokenApi.UploadTokenDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokenApi.UploadTokenDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTokenDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**UploadTokenDeleteRequest**](UploadTokenDeleteRequest.md) |  | 

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


## UploadTokenGet

> KalturaUploadToken UploadTokenGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewUploadTokenDeleteRequest("UploadTokenId_example") // UploadTokenDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadTokenApi.UploadTokenGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokenApi.UploadTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTokenGet`: KalturaUploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokenApi.UploadTokenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**UploadTokenDeleteRequest**](UploadTokenDeleteRequest.md) |  | 

### Return type

[**KalturaUploadToken**](KalturaUploadToken.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTokenList

> KalturaUploadTokenListResponse UploadTokenList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewUploadTokenListRequest() // UploadTokenListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadTokenApi.UploadTokenList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokenApi.UploadTokenList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTokenList`: KalturaUploadTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `UploadTokenApi.UploadTokenList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTokenListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**UploadTokenListRequest**](UploadTokenListRequest.md) |  | 

### Return type

[**KalturaUploadTokenListResponse**](KalturaUploadTokenListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTokenUpload

> KalturaUploadToken UploadTokenUpload(ctx).UploadTokenId(uploadTokenId).FileData(fileData).Ks(ks).Format(format).Resume(resume).FinalChunk(finalChunk).ResumeAt(resumeAt).Execute()





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
    uploadTokenId := "uploadTokenId_example" // string | 
    fileData := os.NewFile(1234, "some_file") // *os.File | 
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)
    resume := true // bool |  (optional) (default to false)
    finalChunk := true // bool |  (optional) (default to true)
    resumeAt := float32(8.14) // float32 |  (optional) (default to -1.0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UploadTokenApi.UploadTokenUpload(context.Background()).UploadTokenId(uploadTokenId).FileData(fileData).Ks(ks).Format(format).Resume(resume).FinalChunk(finalChunk).ResumeAt(resumeAt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UploadTokenApi.UploadTokenUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadTokenUpload`: KalturaUploadToken
    fmt.Fprintf(os.Stdout, "Response from `UploadTokenApi.UploadTokenUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTokenUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadTokenId** | **string** |  | 
 **fileData** | ***os.File** |  | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **resume** | **bool** |  | [default to false]
 **finalChunk** | **bool** |  | [default to true]
 **resumeAt** | **float32** |  | [default to -1.0]

### Return type

[**KalturaUploadToken**](KalturaUploadToken.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


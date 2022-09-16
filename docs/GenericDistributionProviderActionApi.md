# \GenericDistributionProviderActionApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenericDistributionProviderActionAdd**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAdd) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/add | 
[**GenericDistributionProviderActionAddMrssTransform**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddMrssTransform) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addMrssTransform | 
[**GenericDistributionProviderActionAddMrssTransformFromFile**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddMrssTransformFromFile) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addMrssTransformFromFile | 
[**GenericDistributionProviderActionAddMrssValidate**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddMrssValidate) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addMrssValidate | 
[**GenericDistributionProviderActionAddMrssValidateFromFile**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddMrssValidateFromFile) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addMrssValidateFromFile | 
[**GenericDistributionProviderActionAddResultsTransform**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddResultsTransform) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addResultsTransform | 
[**GenericDistributionProviderActionAddResultsTransformFromFile**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionAddResultsTransformFromFile) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/addResultsTransformFromFile | 
[**GenericDistributionProviderActionDelete**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionDelete) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/delete | 
[**GenericDistributionProviderActionDeleteByProviderId**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionDeleteByProviderId) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/deleteByProviderId | 
[**GenericDistributionProviderActionGet**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionGet) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/get | 
[**GenericDistributionProviderActionGetByProviderId**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionGetByProviderId) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/getByProviderId | 
[**GenericDistributionProviderActionList**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionList) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/list | 
[**GenericDistributionProviderActionUpdate**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionUpdate) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/update | 
[**GenericDistributionProviderActionUpdateByProviderId**](GenericDistributionProviderActionApi.md#GenericDistributionProviderActionUpdateByProviderId) | **Post** /service/contentdistribution_genericdistributionprovideraction/action/updateByProviderId | 



## GenericDistributionProviderActionAdd

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionAddRequest(*openapiclient.NewKalturaGenericDistributionProviderAction()) // GenericDistributionProviderActionAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAdd`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionAddRequest**](GenericDistributionProviderActionAddRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddMrssTransform

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddMrssTransform(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionAddMrssTransformRequest(int32(123), "XslData_example") // GenericDistributionProviderActionAddMrssTransformRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransform(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddMrssTransform`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddMrssTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionAddMrssTransformRequest**](GenericDistributionProviderActionAddMrssTransformRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddMrssTransformFromFile

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddMrssTransformFromFile(ctx).Id(id).XslFile(xslFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | the id of the generic distribution provider action
    xslFile := os.NewFile(1234, "some_file") // *os.File | XSL MRSS transformation file
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransformFromFile(context.Background()).Id(id).XslFile(xslFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransformFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddMrssTransformFromFile`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssTransformFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddMrssTransformFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | the id of the generic distribution provider action | 
 **xslFile** | ***os.File** | XSL MRSS transformation file | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddMrssValidate

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddMrssValidate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionAddMrssValidateRequest(int32(123), "XsdData_example") // GenericDistributionProviderActionAddMrssValidateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddMrssValidate`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddMrssValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionAddMrssValidateRequest**](GenericDistributionProviderActionAddMrssValidateRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddMrssValidateFromFile

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddMrssValidateFromFile(ctx).Id(id).XsdFile(xsdFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | the id of the generic distribution provider action
    xsdFile := os.NewFile(1234, "some_file") // *os.File | XSD MRSS validatation file
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidateFromFile(context.Background()).Id(id).XsdFile(xsdFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidateFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddMrssValidateFromFile`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddMrssValidateFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddMrssValidateFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | the id of the generic distribution provider action | 
 **xsdFile** | ***os.File** | XSD MRSS validatation file | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddResultsTransform

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddResultsTransform(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionAddResultsTransformRequest(int32(123), "TransformData_example") // GenericDistributionProviderActionAddResultsTransformRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransform(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddResultsTransform`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddResultsTransformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionAddResultsTransformRequest**](GenericDistributionProviderActionAddResultsTransformRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionAddResultsTransformFromFile

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionAddResultsTransformFromFile(ctx).Id(id).TransformFile(transformFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | the id of the generic distribution provider action
    transformFile := os.NewFile(1234, "some_file") // *os.File | transformation file xsl, xPath or regex
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransformFromFile(context.Background()).Id(id).TransformFile(transformFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransformFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionAddResultsTransformFromFile`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionAddResultsTransformFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionAddResultsTransformFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** | the id of the generic distribution provider action | 
 **transformFile** | ***os.File** | transformation file xsl, xPath or regex | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionDelete

> GenericDistributionProviderActionDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionDeleteRequest struct via the builder pattern


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


## GenericDistributionProviderActionDeleteByProviderId

> GenericDistributionProviderActionDeleteByProviderId(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionDeleteByProviderIdRequest(int32(123), int32(123)) // GenericDistributionProviderActionDeleteByProviderIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionDeleteByProviderId(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionDeleteByProviderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionDeleteByProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionDeleteByProviderIdRequest**](GenericDistributionProviderActionDeleteByProviderIdRequest.md) |  | 

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


## GenericDistributionProviderActionGet

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionGet`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionGetByProviderId

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionGetByProviderId(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionDeleteByProviderIdRequest(int32(123), int32(123)) // GenericDistributionProviderActionDeleteByProviderIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionGetByProviderId(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionGetByProviderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionGetByProviderId`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionGetByProviderId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionGetByProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionDeleteByProviderIdRequest**](GenericDistributionProviderActionDeleteByProviderIdRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionList

> KalturaGenericDistributionProviderActionListResponse GenericDistributionProviderActionList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionListRequest() // GenericDistributionProviderActionListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionList`: KalturaGenericDistributionProviderActionListResponse
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionListRequest**](GenericDistributionProviderActionListRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderActionListResponse**](KalturaGenericDistributionProviderActionListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionUpdate

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionUpdateRequest(*openapiclient.NewKalturaGenericDistributionProviderAction(), int32(123)) // GenericDistributionProviderActionUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionUpdate`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionUpdateRequest**](GenericDistributionProviderActionUpdateRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericDistributionProviderActionUpdateByProviderId

> KalturaGenericDistributionProviderAction GenericDistributionProviderActionUpdateByProviderId(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewGenericDistributionProviderActionUpdateByProviderIdRequest(int32(123), *openapiclient.NewKalturaGenericDistributionProviderAction(), int32(123)) // GenericDistributionProviderActionUpdateByProviderIdRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdateByProviderId(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdateByProviderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenericDistributionProviderActionUpdateByProviderId`: KalturaGenericDistributionProviderAction
    fmt.Fprintf(os.Stdout, "Response from `GenericDistributionProviderActionApi.GenericDistributionProviderActionUpdateByProviderId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericDistributionProviderActionUpdateByProviderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**GenericDistributionProviderActionUpdateByProviderIdRequest**](GenericDistributionProviderActionUpdateByProviderIdRequest.md) |  | 

### Return type

[**KalturaGenericDistributionProviderAction**](KalturaGenericDistributionProviderAction.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


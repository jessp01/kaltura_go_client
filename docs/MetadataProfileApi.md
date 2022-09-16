# \MetadataProfileApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MetadataProfileAdd**](MetadataProfileApi.md#MetadataProfileAdd) | **Post** /service/metadata_metadataprofile/action/add | 
[**MetadataProfileAddFromFile**](MetadataProfileApi.md#MetadataProfileAddFromFile) | **Post** /service/metadata_metadataprofile/action/addFromFile | 
[**MetadataProfileDelete**](MetadataProfileApi.md#MetadataProfileDelete) | **Post** /service/metadata_metadataprofile/action/delete | 
[**MetadataProfileGet**](MetadataProfileApi.md#MetadataProfileGet) | **Post** /service/metadata_metadataprofile/action/get | 
[**MetadataProfileList**](MetadataProfileApi.md#MetadataProfileList) | **Post** /service/metadata_metadataprofile/action/list | 
[**MetadataProfileListFields**](MetadataProfileApi.md#MetadataProfileListFields) | **Post** /service/metadata_metadataprofile/action/listFields | 
[**MetadataProfileRevert**](MetadataProfileApi.md#MetadataProfileRevert) | **Post** /service/metadata_metadataprofile/action/revert | 
[**MetadataProfileServe**](MetadataProfileApi.md#MetadataProfileServe) | **Post** /service/metadata_metadataprofile/action/serve | 
[**MetadataProfileServeView**](MetadataProfileApi.md#MetadataProfileServeView) | **Post** /service/metadata_metadataprofile/action/serveView | 
[**MetadataProfileUpdate**](MetadataProfileApi.md#MetadataProfileUpdate) | **Post** /service/metadata_metadataprofile/action/update | 
[**MetadataProfileUpdateDefinitionFromFile**](MetadataProfileApi.md#MetadataProfileUpdateDefinitionFromFile) | **Post** /service/metadata_metadataprofile/action/updateDefinitionFromFile | 
[**MetadataProfileUpdateTransformationFromFile**](MetadataProfileApi.md#MetadataProfileUpdateTransformationFromFile) | **Post** /service/metadata_metadataprofile/action/updateTransformationFromFile | 
[**MetadataProfileUpdateViewsFromFile**](MetadataProfileApi.md#MetadataProfileUpdateViewsFromFile) | **Post** /service/metadata_metadataprofile/action/updateViewsFromFile | 



## MetadataProfileAdd

> KalturaMetadataProfile MetadataProfileAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMetadataProfileAddRequest(*openapiclient.NewKalturaMetadataProfile(), "XsdData_example") // MetadataProfileAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileAdd`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataProfileAddRequest**](MetadataProfileAddRequest.md) |  | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileAddFromFile

> KalturaMetadataProfile MetadataProfileAddFromFile(ctx).XsdFile(xsdFile).Ks(ks).Format(format).MetadataProfileMetadataObjectType(metadataProfileMetadataObjectType).MetadataProfileName(metadataProfileName).MetadataProfileSystemName(metadataProfileSystemName).MetadataProfileDescription(metadataProfileDescription).MetadataProfileCreateMode(metadataProfileCreateMode).MetadataProfileDisableReIndexing(metadataProfileDisableReIndexing).ViewsFile(viewsFile).Execute()





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
    xsdFile := os.NewFile(1234, "some_file") // *os.File | XSD metadata definition
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)
    metadataProfileMetadataObjectType := "metadataProfileMetadataObjectType_example" // string | Enum Type: `KalturaMetadataObjectType` (optional)
    metadataProfileName := "metadataProfileName_example" // string |  (optional)
    metadataProfileSystemName := "metadataProfileSystemName_example" // string |  (optional)
    metadataProfileDescription := "metadataProfileDescription_example" // string |  (optional)
    metadataProfileCreateMode := int32(56) // int32 | Enum Type: `KalturaMetadataProfileCreateMode` (optional)
    metadataProfileDisableReIndexing := true // bool |  (optional)
    viewsFile := os.NewFile(1234, "some_file") // *os.File | UI views definition (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileAddFromFile(context.Background()).XsdFile(xsdFile).Ks(ks).Format(format).MetadataProfileMetadataObjectType(metadataProfileMetadataObjectType).MetadataProfileName(metadataProfileName).MetadataProfileSystemName(metadataProfileSystemName).MetadataProfileDescription(metadataProfileDescription).MetadataProfileCreateMode(metadataProfileCreateMode).MetadataProfileDisableReIndexing(metadataProfileDisableReIndexing).ViewsFile(viewsFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileAddFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileAddFromFile`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileAddFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileAddFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xsdFile** | ***os.File** | XSD metadata definition | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **metadataProfileMetadataObjectType** | **string** | Enum Type: &#x60;KalturaMetadataObjectType&#x60; | 
 **metadataProfileName** | **string** |  | 
 **metadataProfileSystemName** | **string** |  | 
 **metadataProfileDescription** | **string** |  | 
 **metadataProfileCreateMode** | **int32** | Enum Type: &#x60;KalturaMetadataProfileCreateMode&#x60; | 
 **metadataProfileDisableReIndexing** | **bool** |  | 
 **viewsFile** | ***os.File** | UI views definition | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileDelete

> MetadataProfileDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileDeleteRequest struct via the builder pattern


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


## MetadataProfileGet

> KalturaMetadataProfile MetadataProfileGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileGet`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileList

> KalturaMetadataProfileListResponse MetadataProfileList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMetadataProfileListRequest() // MetadataProfileListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileList`: KalturaMetadataProfileListResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataProfileListRequest**](MetadataProfileListRequest.md) |  | 

### Return type

[**KalturaMetadataProfileListResponse**](KalturaMetadataProfileListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileListFields

> KalturaMetadataProfileFieldListResponse MetadataProfileListFields(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMetadataProfileListFieldsRequest(int32(123)) // MetadataProfileListFieldsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileListFields(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileListFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileListFields`: KalturaMetadataProfileFieldListResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileListFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileListFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataProfileListFieldsRequest**](MetadataProfileListFieldsRequest.md) |  | 

### Return type

[**KalturaMetadataProfileFieldListResponse**](KalturaMetadataProfileFieldListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileRevert

> KalturaMetadataProfile MetadataProfileRevert(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMetadataProfileRevertRequest(int32(123), int32(123)) // MetadataProfileRevertRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileRevert(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileRevert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileRevert`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileRevert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataProfileRevertRequest**](MetadataProfileRevertRequest.md) |  | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileServe

> string MetadataProfileServe(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileServe(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileServe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileServe`: string
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileServe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileServeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

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


## MetadataProfileServeView

> string MetadataProfileServeView(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileServeView(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileServeView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileServeView`: string
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileServeView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileServeViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

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


## MetadataProfileUpdate

> KalturaMetadataProfile MetadataProfileUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewMetadataProfileUpdateRequest(int32(123), *openapiclient.NewKalturaMetadataProfile()) // MetadataProfileUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileUpdate`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**MetadataProfileUpdateRequest**](MetadataProfileUpdateRequest.md) |  | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileUpdateDefinitionFromFile

> KalturaMetadataProfile MetadataProfileUpdateDefinitionFromFile(ctx).Id(id).XsdFile(xsdFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | 
    xsdFile := os.NewFile(1234, "some_file") // *os.File | XSD metadata definition
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileUpdateDefinitionFromFile(context.Background()).Id(id).XsdFile(xsdFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileUpdateDefinitionFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileUpdateDefinitionFromFile`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileUpdateDefinitionFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileUpdateDefinitionFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **xsdFile** | ***os.File** | XSD metadata definition | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileUpdateTransformationFromFile

> KalturaMetadataProfile MetadataProfileUpdateTransformationFromFile(ctx).Id(id).XsltFile(xsltFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | 
    xsltFile := os.NewFile(1234, "some_file") // *os.File | XSLT file, will be executed on every metadata add/update
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileUpdateTransformationFromFile(context.Background()).Id(id).XsltFile(xsltFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileUpdateTransformationFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileUpdateTransformationFromFile`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileUpdateTransformationFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileUpdateTransformationFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **xsltFile** | ***os.File** | XSLT file, will be executed on every metadata add/update | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MetadataProfileUpdateViewsFromFile

> KalturaMetadataProfile MetadataProfileUpdateViewsFromFile(ctx).Id(id).ViewsFile(viewsFile).Ks(ks).Format(format).Execute()





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
    id := int32(56) // int32 | 
    viewsFile := os.NewFile(1234, "some_file") // *os.File | UI views file
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataProfileApi.MetadataProfileUpdateViewsFromFile(context.Background()).Id(id).ViewsFile(viewsFile).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataProfileApi.MetadataProfileUpdateViewsFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MetadataProfileUpdateViewsFromFile`: KalturaMetadataProfile
    fmt.Fprintf(os.Stdout, "Response from `MetadataProfileApi.MetadataProfileUpdateViewsFromFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMetadataProfileUpdateViewsFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32** |  | 
 **viewsFile** | ***os.File** | UI views file | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

[**KalturaMetadataProfile**](KalturaMetadataProfile.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


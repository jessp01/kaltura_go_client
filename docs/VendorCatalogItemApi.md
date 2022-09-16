# \VendorCatalogItemApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VendorCatalogItemAdd**](VendorCatalogItemApi.md#VendorCatalogItemAdd) | **Post** /service/reach_vendorcatalogitem/action/add | 
[**VendorCatalogItemAddFromBulkUpload**](VendorCatalogItemApi.md#VendorCatalogItemAddFromBulkUpload) | **Post** /service/reach_vendorcatalogitem/action/addFromBulkUpload | 
[**VendorCatalogItemDelete**](VendorCatalogItemApi.md#VendorCatalogItemDelete) | **Post** /service/reach_vendorcatalogitem/action/delete | 
[**VendorCatalogItemGet**](VendorCatalogItemApi.md#VendorCatalogItemGet) | **Post** /service/reach_vendorcatalogitem/action/get | 
[**VendorCatalogItemGetServeUrl**](VendorCatalogItemApi.md#VendorCatalogItemGetServeUrl) | **Post** /service/reach_vendorcatalogitem/action/getServeUrl | 
[**VendorCatalogItemList**](VendorCatalogItemApi.md#VendorCatalogItemList) | **Post** /service/reach_vendorcatalogitem/action/list | 
[**VendorCatalogItemServe**](VendorCatalogItemApi.md#VendorCatalogItemServe) | **Post** /service/reach_vendorcatalogitem/action/serve | 
[**VendorCatalogItemUpdate**](VendorCatalogItemApi.md#VendorCatalogItemUpdate) | **Post** /service/reach_vendorcatalogitem/action/update | 
[**VendorCatalogItemUpdateStatus**](VendorCatalogItemApi.md#VendorCatalogItemUpdateStatus) | **Post** /service/reach_vendorcatalogitem/action/updateStatus | 



## VendorCatalogItemAdd

> KalturaVendorCatalogItem VendorCatalogItemAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorCatalogItemAddRequest(*openapiclient.NewKalturaVendorCatalogItem()) // VendorCatalogItemAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemAdd`: KalturaVendorCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemAddRequest**](VendorCatalogItemAddRequest.md) |  | 

### Return type

[**KalturaVendorCatalogItem**](KalturaVendorCatalogItem.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorCatalogItemAddFromBulkUpload

> KalturaBulkUpload VendorCatalogItemAddFromBulkUpload(ctx).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadVendorCatalogItemDataObjectType(bulkUploadVendorCatalogItemDataObjectType).Execute()



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
    bulkUploadVendorCatalogItemDataObjectType := "bulkUploadVendorCatalogItemDataObjectType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemAddFromBulkUpload(context.Background()).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadVendorCatalogItemDataObjectType(bulkUploadVendorCatalogItemDataObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemAddFromBulkUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemAddFromBulkUpload`: KalturaBulkUpload
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemAddFromBulkUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemAddFromBulkUploadRequest struct via the builder pattern


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
 **bulkUploadVendorCatalogItemDataObjectType** | **string** |  | 

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


## VendorCatalogItemDelete

> VendorCatalogItemDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemDeleteRequest struct via the builder pattern


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


## VendorCatalogItemGet

> KalturaVendorCatalogItem VendorCatalogItemGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemGet`: KalturaVendorCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaVendorCatalogItem**](KalturaVendorCatalogItem.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorCatalogItemGetServeUrl

> string VendorCatalogItemGetServeUrl(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewVendorCatalogItemGetServeUrlRequest() // VendorCatalogItemGetServeUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemGetServeUrl(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemGetServeUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemGetServeUrl`: string
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemGetServeUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemGetServeUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemGetServeUrlRequest**](VendorCatalogItemGetServeUrlRequest.md) |  | 

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


## VendorCatalogItemList

> KalturaVendorCatalogItemListResponse VendorCatalogItemList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorCatalogItemListRequest() // VendorCatalogItemListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemList`: KalturaVendorCatalogItemListResponse
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemListRequest**](VendorCatalogItemListRequest.md) |  | 

### Return type

[**KalturaVendorCatalogItemListResponse**](KalturaVendorCatalogItemListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorCatalogItemServe

> string VendorCatalogItemServe(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewVendorCatalogItemGetServeUrlRequest() // VendorCatalogItemGetServeUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemServe(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemServe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemServe`: string
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemServe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemServeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemGetServeUrlRequest**](VendorCatalogItemGetServeUrlRequest.md) |  | 

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


## VendorCatalogItemUpdate

> KalturaVendorCatalogItem VendorCatalogItemUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorCatalogItemUpdateRequest(int32(123), *openapiclient.NewKalturaVendorCatalogItem()) // VendorCatalogItemUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemUpdate`: KalturaVendorCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemUpdateRequest**](VendorCatalogItemUpdateRequest.md) |  | 

### Return type

[**KalturaVendorCatalogItem**](KalturaVendorCatalogItem.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VendorCatalogItemUpdateStatus

> KalturaVendorCatalogItem VendorCatalogItemUpdateStatus(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewVendorCatalogItemUpdateStatusRequest(int32(123), int32(123)) // VendorCatalogItemUpdateStatusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VendorCatalogItemApi.VendorCatalogItemUpdateStatus(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VendorCatalogItemApi.VendorCatalogItemUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VendorCatalogItemUpdateStatus`: KalturaVendorCatalogItem
    fmt.Fprintf(os.Stdout, "Response from `VendorCatalogItemApi.VendorCatalogItemUpdateStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVendorCatalogItemUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**VendorCatalogItemUpdateStatusRequest**](VendorCatalogItemUpdateStatusRequest.md) |  | 

### Return type

[**KalturaVendorCatalogItem**](KalturaVendorCatalogItem.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


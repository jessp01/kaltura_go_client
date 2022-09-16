# \ScheduleResourceApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduleResourceAdd**](ScheduleResourceApi.md#ScheduleResourceAdd) | **Post** /service/schedule_scheduleresource/action/add | 
[**ScheduleResourceAddFromBulkUpload**](ScheduleResourceApi.md#ScheduleResourceAddFromBulkUpload) | **Post** /service/schedule_scheduleresource/action/addFromBulkUpload | 
[**ScheduleResourceDelete**](ScheduleResourceApi.md#ScheduleResourceDelete) | **Post** /service/schedule_scheduleresource/action/delete | 
[**ScheduleResourceGet**](ScheduleResourceApi.md#ScheduleResourceGet) | **Post** /service/schedule_scheduleresource/action/get | 
[**ScheduleResourceList**](ScheduleResourceApi.md#ScheduleResourceList) | **Post** /service/schedule_scheduleresource/action/list | 
[**ScheduleResourceUpdate**](ScheduleResourceApi.md#ScheduleResourceUpdate) | **Post** /service/schedule_scheduleresource/action/update | 



## ScheduleResourceAdd

> KalturaScheduleResource ScheduleResourceAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleResourceAddRequest(*openapiclient.NewKalturaScheduleResource()) // ScheduleResourceAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceAdd`: KalturaScheduleResource
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleResourceAddRequest**](ScheduleResourceAddRequest.md) |  | 

### Return type

[**KalturaScheduleResource**](KalturaScheduleResource.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleResourceAddFromBulkUpload

> KalturaBulkUpload ScheduleResourceAddFromBulkUpload(ctx).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadDataColumns(bulkUploadDataColumns).BulkUploadDataProcessObjectId(bulkUploadDataProcessObjectId).BulkUploadDataProcessObjectType(bulkUploadDataProcessObjectType).Execute()





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
    bulkUploadDataColumns := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} |  (optional)
    bulkUploadDataProcessObjectId := "bulkUploadDataProcessObjectId_example" // string | The object in process (optional)
    bulkUploadDataProcessObjectType := "bulkUploadDataProcessObjectType_example" // string | The type of the object in process (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceAddFromBulkUpload(context.Background()).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).BulkUploadDataColumns(bulkUploadDataColumns).BulkUploadDataProcessObjectId(bulkUploadDataProcessObjectId).BulkUploadDataProcessObjectType(bulkUploadDataProcessObjectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceAddFromBulkUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceAddFromBulkUpload`: KalturaBulkUpload
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceAddFromBulkUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceAddFromBulkUploadRequest struct via the builder pattern


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
 **bulkUploadDataColumns** | **[]map[string]interface{}** |  | 
 **bulkUploadDataProcessObjectId** | **string** | The object in process | 
 **bulkUploadDataProcessObjectType** | **string** | The type of the object in process | 

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


## ScheduleResourceDelete

> KalturaScheduleResource ScheduleResourceDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleResourceDeleteRequest(int32(123)) // ScheduleResourceDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceDelete`: KalturaScheduleResource
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleResourceDeleteRequest**](ScheduleResourceDeleteRequest.md) |  | 

### Return type

[**KalturaScheduleResource**](KalturaScheduleResource.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleResourceGet

> KalturaScheduleResource ScheduleResourceGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleResourceDeleteRequest(int32(123)) // ScheduleResourceDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceGet`: KalturaScheduleResource
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleResourceDeleteRequest**](ScheduleResourceDeleteRequest.md) |  | 

### Return type

[**KalturaScheduleResource**](KalturaScheduleResource.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleResourceList

> KalturaScheduleResourceListResponse ScheduleResourceList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleResourceListRequest() // ScheduleResourceListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceList`: KalturaScheduleResourceListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleResourceListRequest**](ScheduleResourceListRequest.md) |  | 

### Return type

[**KalturaScheduleResourceListResponse**](KalturaScheduleResourceListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleResourceUpdate

> KalturaScheduleResource ScheduleResourceUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleResourceUpdateRequest(*openapiclient.NewKalturaScheduleResource(), int32(123)) // ScheduleResourceUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleResourceApi.ScheduleResourceUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleResourceApi.ScheduleResourceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleResourceUpdate`: KalturaScheduleResource
    fmt.Fprintf(os.Stdout, "Response from `ScheduleResourceApi.ScheduleResourceUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleResourceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleResourceUpdateRequest**](ScheduleResourceUpdateRequest.md) |  | 

### Return type

[**KalturaScheduleResource**](KalturaScheduleResource.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


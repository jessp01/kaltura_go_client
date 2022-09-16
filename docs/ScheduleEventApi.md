# \ScheduleEventApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScheduleEventAdd**](ScheduleEventApi.md#ScheduleEventAdd) | **Post** /service/schedule_scheduleevent/action/add | 
[**ScheduleEventAddFromBulkUpload**](ScheduleEventApi.md#ScheduleEventAddFromBulkUpload) | **Post** /service/schedule_scheduleevent/action/addFromBulkUpload | 
[**ScheduleEventCancel**](ScheduleEventApi.md#ScheduleEventCancel) | **Post** /service/schedule_scheduleevent/action/cancel | 
[**ScheduleEventDelete**](ScheduleEventApi.md#ScheduleEventDelete) | **Post** /service/schedule_scheduleevent/action/delete | 
[**ScheduleEventGet**](ScheduleEventApi.md#ScheduleEventGet) | **Post** /service/schedule_scheduleevent/action/get | 
[**ScheduleEventGetConflicts**](ScheduleEventApi.md#ScheduleEventGetConflicts) | **Post** /service/schedule_scheduleevent/action/getConflicts | 
[**ScheduleEventList**](ScheduleEventApi.md#ScheduleEventList) | **Post** /service/schedule_scheduleevent/action/list | 
[**ScheduleEventUpdate**](ScheduleEventApi.md#ScheduleEventUpdate) | **Post** /service/schedule_scheduleevent/action/update | 
[**ScheduleEventUpdateLiveFeature**](ScheduleEventApi.md#ScheduleEventUpdateLiveFeature) | **Post** /service/schedule_scheduleevent/action/updateLiveFeature | 



## ScheduleEventAdd

> KalturaScheduleEvent ScheduleEventAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventAddRequest(*openapiclient.NewKalturaScheduleEvent()) // ScheduleEventAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventAdd`: KalturaScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventAddRequest**](ScheduleEventAddRequest.md) |  | 

### Return type

[**KalturaScheduleEvent**](KalturaScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventAddFromBulkUpload

> KalturaBulkUpload ScheduleEventAddFromBulkUpload(ctx).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventAddFromBulkUpload(context.Background()).FileData(fileData).Ks(ks).Format(format).BulkUploadDataObjectType(bulkUploadDataObjectType).BulkUploadDataFileName(bulkUploadDataFileName).BulkUploadDataObjectDataObjectType(bulkUploadDataObjectDataObjectType).BulkUploadDataEmailRecipients(bulkUploadDataEmailRecipients).BulkUploadDataNumOfErrorObjects(bulkUploadDataNumOfErrorObjects).BulkUploadDataPrivileges(bulkUploadDataPrivileges).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventAddFromBulkUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventAddFromBulkUpload`: KalturaBulkUpload
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventAddFromBulkUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventAddFromBulkUploadRequest struct via the builder pattern


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


## ScheduleEventCancel

> KalturaScheduleEvent ScheduleEventCancel(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventCancelRequest(int32(123)) // ScheduleEventCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventCancel(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventCancel`: KalturaScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventCancelRequest**](ScheduleEventCancelRequest.md) |  | 

### Return type

[**KalturaScheduleEvent**](KalturaScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventDelete

> KalturaScheduleEvent ScheduleEventDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventCancelRequest(int32(123)) // ScheduleEventCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventDelete`: KalturaScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventCancelRequest**](ScheduleEventCancelRequest.md) |  | 

### Return type

[**KalturaScheduleEvent**](KalturaScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventGet

> KalturaScheduleEvent ScheduleEventGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventCancelRequest(int32(123)) // ScheduleEventCancelRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventGet`: KalturaScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventCancelRequest**](ScheduleEventCancelRequest.md) |  | 

### Return type

[**KalturaScheduleEvent**](KalturaScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventGetConflicts

> KalturaScheduleEventListResponse ScheduleEventGetConflicts(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventGetConflictsRequest("ResourceIds_example", *openapiclient.NewKalturaScheduleEvent()) // ScheduleEventGetConflictsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventGetConflicts(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventGetConflicts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventGetConflicts`: KalturaScheduleEventListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventGetConflicts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventGetConflictsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventGetConflictsRequest**](ScheduleEventGetConflictsRequest.md) |  | 

### Return type

[**KalturaScheduleEventListResponse**](KalturaScheduleEventListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventList

> KalturaScheduleEventListResponse ScheduleEventList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventListRequest() // ScheduleEventListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventList`: KalturaScheduleEventListResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventListRequest**](ScheduleEventListRequest.md) |  | 

### Return type

[**KalturaScheduleEventListResponse**](KalturaScheduleEventListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventUpdate

> KalturaScheduleEvent ScheduleEventUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventUpdateRequest(*openapiclient.NewKalturaScheduleEvent(), int32(123)) // ScheduleEventUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventUpdate`: KalturaScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventUpdateRequest**](ScheduleEventUpdateRequest.md) |  | 

### Return type

[**KalturaScheduleEvent**](KalturaScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEventUpdateLiveFeature

> KalturaLiveStreamScheduleEvent ScheduleEventUpdateLiveFeature(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewScheduleEventUpdateLiveFeatureRequest("FeatureName_example", *openapiclient.NewKalturaLiveFeature(), int32(123)) // ScheduleEventUpdateLiveFeatureRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScheduleEventApi.ScheduleEventUpdateLiveFeature(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduleEventApi.ScheduleEventUpdateLiveFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleEventUpdateLiveFeature`: KalturaLiveStreamScheduleEvent
    fmt.Fprintf(os.Stdout, "Response from `ScheduleEventApi.ScheduleEventUpdateLiveFeature`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEventUpdateLiveFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**ScheduleEventUpdateLiveFeatureRequest**](ScheduleEventUpdateLiveFeatureRequest.md) |  | 

### Return type

[**KalturaLiveStreamScheduleEvent**](KalturaLiveStreamScheduleEvent.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


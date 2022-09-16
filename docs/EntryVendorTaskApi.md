# \EntryVendorTaskApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntryVendorTaskAbort**](EntryVendorTaskApi.md#EntryVendorTaskAbort) | **Post** /service/reach_entryvendortask/action/abort | 
[**EntryVendorTaskAdd**](EntryVendorTaskApi.md#EntryVendorTaskAdd) | **Post** /service/reach_entryvendortask/action/add | 
[**EntryVendorTaskApprove**](EntryVendorTaskApi.md#EntryVendorTaskApprove) | **Post** /service/reach_entryvendortask/action/approve | 
[**EntryVendorTaskExportToCsv**](EntryVendorTaskApi.md#EntryVendorTaskExportToCsv) | **Post** /service/reach_entryvendortask/action/exportToCsv | 
[**EntryVendorTaskExtendAccessKey**](EntryVendorTaskApi.md#EntryVendorTaskExtendAccessKey) | **Post** /service/reach_entryvendortask/action/extendAccessKey | 
[**EntryVendorTaskGet**](EntryVendorTaskApi.md#EntryVendorTaskGet) | **Post** /service/reach_entryvendortask/action/get | 
[**EntryVendorTaskGetJobs**](EntryVendorTaskApi.md#EntryVendorTaskGetJobs) | **Post** /service/reach_entryvendortask/action/getJobs | 
[**EntryVendorTaskGetServeUrl**](EntryVendorTaskApi.md#EntryVendorTaskGetServeUrl) | **Post** /service/reach_entryvendortask/action/getServeUrl | 
[**EntryVendorTaskList**](EntryVendorTaskApi.md#EntryVendorTaskList) | **Post** /service/reach_entryvendortask/action/list | 
[**EntryVendorTaskReject**](EntryVendorTaskApi.md#EntryVendorTaskReject) | **Post** /service/reach_entryvendortask/action/reject | 
[**EntryVendorTaskServe**](EntryVendorTaskApi.md#EntryVendorTaskServe) | **Post** /service/reach_entryvendortask/action/serve | 
[**EntryVendorTaskServeCsv**](EntryVendorTaskApi.md#EntryVendorTaskServeCsv) | **Post** /service/reach_entryvendortask/action/serveCsv | 
[**EntryVendorTaskUpdate**](EntryVendorTaskApi.md#EntryVendorTaskUpdate) | **Post** /service/reach_entryvendortask/action/update | 
[**EntryVendorTaskUpdateJob**](EntryVendorTaskApi.md#EntryVendorTaskUpdateJob) | **Post** /service/reach_entryvendortask/action/updateJob | 



## EntryVendorTaskAbort

> KalturaEntryVendorTask EntryVendorTaskAbort(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskAbortRequest(int32(123)) // EntryVendorTaskAbortRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskAbort(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskAbort`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskAbort`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskAbortRequest**](EntryVendorTaskAbortRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskAdd

> KalturaEntryVendorTask EntryVendorTaskAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskAddRequest(*openapiclient.NewKalturaEntryVendorTask()) // EntryVendorTaskAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskAdd`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskAddRequest**](EntryVendorTaskAddRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskApprove

> KalturaEntryVendorTask EntryVendorTaskApprove(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskApprove(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskApprove`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskApprove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskExportToCsv

> string EntryVendorTaskExportToCsv(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskExportToCsvRequest(*openapiclient.NewKalturaEntryVendorTaskFilter()) // EntryVendorTaskExportToCsvRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskExportToCsv(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskExportToCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskExportToCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskExportToCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskExportToCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskExportToCsvRequest**](EntryVendorTaskExportToCsvRequest.md) |  | 

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


## EntryVendorTaskExtendAccessKey

> KalturaEntryVendorTask EntryVendorTaskExtendAccessKey(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskExtendAccessKey(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskExtendAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskExtendAccessKey`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskExtendAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskExtendAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskGet

> KalturaEntryVendorTask EntryVendorTaskGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskGet`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskGetJobs

> KalturaEntryVendorTaskListResponse EntryVendorTaskGetJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskGetJobsRequest() // EntryVendorTaskGetJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskGetJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskGetJobs`: KalturaEntryVendorTaskListResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskGetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskGetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskGetJobsRequest**](EntryVendorTaskGetJobsRequest.md) |  | 

### Return type

[**KalturaEntryVendorTaskListResponse**](KalturaEntryVendorTaskListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskGetServeUrl

> string EntryVendorTaskGetServeUrl(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewEntryVendorTaskGetServeUrlRequest() // EntryVendorTaskGetServeUrlRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskGetServeUrl(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskGetServeUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskGetServeUrl`: string
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskGetServeUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskGetServeUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskGetServeUrlRequest**](EntryVendorTaskGetServeUrlRequest.md) |  | 

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


## EntryVendorTaskList

> KalturaEntryVendorTaskListResponse EntryVendorTaskList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskGetJobsRequest() // EntryVendorTaskGetJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskList`: KalturaEntryVendorTaskListResponse
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskGetJobsRequest**](EntryVendorTaskGetJobsRequest.md) |  | 

### Return type

[**KalturaEntryVendorTaskListResponse**](KalturaEntryVendorTaskListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskReject

> KalturaEntryVendorTask EntryVendorTaskReject(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskRejectRequest(int32(123)) // EntryVendorTaskRejectRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskReject(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskReject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskReject`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskReject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskRejectRequest**](EntryVendorTaskRejectRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskServe

> string EntryVendorTaskServe(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewEntryVendorTaskServeRequest() // EntryVendorTaskServeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskServe(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskServe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskServe`: string
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskServe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskServeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskServeRequest**](EntryVendorTaskServeRequest.md) |  | 

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


## EntryVendorTaskServeCsv

> string EntryVendorTaskServeCsv(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewAnnotationDeleteRequest("Id_example") // AnnotationDeleteRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskServeCsv(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskServeCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskServeCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskServeCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskServeCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AnnotationDeleteRequest**](AnnotationDeleteRequest.md) |  | 

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


## EntryVendorTaskUpdate

> KalturaEntryVendorTask EntryVendorTaskUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskUpdateRequest(*openapiclient.NewKalturaEntryVendorTask(), int32(123)) // EntryVendorTaskUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskUpdate`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskUpdateRequest**](EntryVendorTaskUpdateRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryVendorTaskUpdateJob

> KalturaEntryVendorTask EntryVendorTaskUpdateJob(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEntryVendorTaskUpdateRequest(*openapiclient.NewKalturaEntryVendorTask(), int32(123)) // EntryVendorTaskUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EntryVendorTaskApi.EntryVendorTaskUpdateJob(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntryVendorTaskApi.EntryVendorTaskUpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EntryVendorTaskUpdateJob`: KalturaEntryVendorTask
    fmt.Fprintf(os.Stdout, "Response from `EntryVendorTaskApi.EntryVendorTaskUpdateJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryVendorTaskUpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EntryVendorTaskUpdateRequest**](EntryVendorTaskUpdateRequest.md) |  | 

### Return type

[**KalturaEntryVendorTask**](KalturaEntryVendorTask.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


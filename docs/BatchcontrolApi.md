# \BatchcontrolApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchcontrolConfigLoaded**](BatchcontrolApi.md#BatchcontrolConfigLoaded) | **Post** /service/batchcontrol/action/configLoaded | 
[**BatchcontrolGetCommand**](BatchcontrolApi.md#BatchcontrolGetCommand) | **Post** /service/batchcontrol/action/getCommand | 
[**BatchcontrolGetFullStatus**](BatchcontrolApi.md#BatchcontrolGetFullStatus) | **Post** /service/batchcontrol/action/getFullStatus | 
[**BatchcontrolGetOrCreateScheduler**](BatchcontrolApi.md#BatchcontrolGetOrCreateScheduler) | **Post** /service/batchcontrol/action/getOrCreateScheduler | 
[**BatchcontrolKill**](BatchcontrolApi.md#BatchcontrolKill) | **Post** /service/batchcontrol/action/kill | 
[**BatchcontrolListCommands**](BatchcontrolApi.md#BatchcontrolListCommands) | **Post** /service/batchcontrol/action/listCommands | 
[**BatchcontrolListSchedulers**](BatchcontrolApi.md#BatchcontrolListSchedulers) | **Post** /service/batchcontrol/action/listSchedulers | 
[**BatchcontrolListWorkers**](BatchcontrolApi.md#BatchcontrolListWorkers) | **Post** /service/batchcontrol/action/listWorkers | 
[**BatchcontrolReportStatus**](BatchcontrolApi.md#BatchcontrolReportStatus) | **Post** /service/batchcontrol/action/reportStatus | 
[**BatchcontrolSetCommandResult**](BatchcontrolApi.md#BatchcontrolSetCommandResult) | **Post** /service/batchcontrol/action/setCommandResult | 
[**BatchcontrolSetSchedulerConfig**](BatchcontrolApi.md#BatchcontrolSetSchedulerConfig) | **Post** /service/batchcontrol/action/setSchedulerConfig | 
[**BatchcontrolSetWorkerConfig**](BatchcontrolApi.md#BatchcontrolSetWorkerConfig) | **Post** /service/batchcontrol/action/setWorkerConfig | 
[**BatchcontrolStartWorker**](BatchcontrolApi.md#BatchcontrolStartWorker) | **Post** /service/batchcontrol/action/startWorker | 
[**BatchcontrolStopScheduler**](BatchcontrolApi.md#BatchcontrolStopScheduler) | **Post** /service/batchcontrol/action/stopScheduler | 
[**BatchcontrolStopWorker**](BatchcontrolApi.md#BatchcontrolStopWorker) | **Post** /service/batchcontrol/action/stopWorker | 



## BatchcontrolConfigLoaded

> KalturaSchedulerConfig BatchcontrolConfigLoaded(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolConfigLoadedRequest("ConfigParam_example", "ConfigValue_example", *openapiclient.NewKalturaScheduler()) // BatchcontrolConfigLoadedRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolConfigLoaded(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolConfigLoaded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolConfigLoaded`: KalturaSchedulerConfig
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolConfigLoaded`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolConfigLoadedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolConfigLoadedRequest**](BatchcontrolConfigLoadedRequest.md) |  | 

### Return type

[**KalturaSchedulerConfig**](KalturaSchedulerConfig.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolGetCommand

> KalturaControlPanelCommand BatchcontrolGetCommand(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolGetCommandRequest(int32(123)) // BatchcontrolGetCommandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolGetCommand(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolGetCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolGetCommand`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolGetCommand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolGetCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolGetCommandRequest**](BatchcontrolGetCommandRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolGetFullStatus

> KalturaFullStatusResponse BatchcontrolGetFullStatus(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolGetFullStatus(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolGetFullStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolGetFullStatus`: KalturaFullStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolGetFullStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolGetFullStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

[**KalturaFullStatusResponse**](KalturaFullStatusResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolGetOrCreateScheduler

> KalturaScheduler BatchcontrolGetOrCreateScheduler(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewBatchcontrolGetOrCreateSchedulerRequest(*openapiclient.NewKalturaScheduler()) // BatchcontrolGetOrCreateSchedulerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolGetOrCreateScheduler(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolGetOrCreateScheduler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolGetOrCreateScheduler`: KalturaScheduler
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolGetOrCreateScheduler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolGetOrCreateSchedulerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolGetOrCreateSchedulerRequest**](BatchcontrolGetOrCreateSchedulerRequest.md) |  | 

### Return type

[**KalturaScheduler**](KalturaScheduler.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolKill

> KalturaControlPanelCommand BatchcontrolKill(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolKillRequest(int32(123), int32(123), "Cause_example", int32(123)) // BatchcontrolKillRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolKill(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolKill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolKill`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolKill`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolKillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolKillRequest**](BatchcontrolKillRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolListCommands

> KalturaControlPanelCommandListResponse BatchcontrolListCommands(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolListCommandsRequest() // BatchcontrolListCommandsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolListCommands(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolListCommands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolListCommands`: KalturaControlPanelCommandListResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolListCommands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolListCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolListCommandsRequest**](BatchcontrolListCommandsRequest.md) |  | 

### Return type

[**KalturaControlPanelCommandListResponse**](KalturaControlPanelCommandListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolListSchedulers

> KalturaSchedulerListResponse BatchcontrolListSchedulers(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolListSchedulers(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolListSchedulers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolListSchedulers`: KalturaSchedulerListResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolListSchedulers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolListSchedulersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

[**KalturaSchedulerListResponse**](KalturaSchedulerListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolListWorkers

> KalturaSchedulerWorkerListResponse BatchcontrolListWorkers(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolListWorkers(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolListWorkers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolListWorkers`: KalturaSchedulerWorkerListResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolListWorkers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolListWorkersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

[**KalturaSchedulerWorkerListResponse**](KalturaSchedulerWorkerListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolReportStatus

> KalturaSchedulerStatusResponse BatchcontrolReportStatus(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolReportStatusRequest(*openapiclient.NewKalturaScheduler(), []openapiclient.KalturaSchedulerStatus{*openapiclient.NewKalturaSchedulerStatus()}, []openapiclient.KalturaWorkerQueueFilter{*openapiclient.NewKalturaWorkerQueueFilter()}) // BatchcontrolReportStatusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolReportStatus(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolReportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolReportStatus`: KalturaSchedulerStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolReportStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolReportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolReportStatusRequest**](BatchcontrolReportStatusRequest.md) |  | 

### Return type

[**KalturaSchedulerStatusResponse**](KalturaSchedulerStatusResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolSetCommandResult

> KalturaControlPanelCommand BatchcontrolSetCommandResult(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolSetCommandResultRequest(int32(123), int32(123)) // BatchcontrolSetCommandResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolSetCommandResult(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolSetCommandResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolSetCommandResult`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolSetCommandResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolSetCommandResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolSetCommandResultRequest**](BatchcontrolSetCommandResultRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolSetSchedulerConfig

> KalturaControlPanelCommand BatchcontrolSetSchedulerConfig(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolSetSchedulerConfigRequest(int32(123), "ConfigParam_example", "ConfigValue_example", int32(123)) // BatchcontrolSetSchedulerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolSetSchedulerConfig(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolSetSchedulerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolSetSchedulerConfig`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolSetSchedulerConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolSetSchedulerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolSetSchedulerConfigRequest**](BatchcontrolSetSchedulerConfigRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolSetWorkerConfig

> KalturaControlPanelCommand BatchcontrolSetWorkerConfig(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolSetWorkerConfigRequest(int32(123), "ConfigParam_example", "ConfigValue_example", int32(123)) // BatchcontrolSetWorkerConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolSetWorkerConfig(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolSetWorkerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolSetWorkerConfig`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolSetWorkerConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolSetWorkerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolSetWorkerConfigRequest**](BatchcontrolSetWorkerConfigRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolStartWorker

> KalturaControlPanelCommand BatchcontrolStartWorker(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolStartWorkerRequest(int32(123), int32(123)) // BatchcontrolStartWorkerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolStartWorker(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolStartWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolStartWorker`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolStartWorker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolStartWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolStartWorkerRequest**](BatchcontrolStartWorkerRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolStopScheduler

> KalturaControlPanelCommand BatchcontrolStopScheduler(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolStopSchedulerRequest(int32(123), "Cause_example", int32(123)) // BatchcontrolStopSchedulerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolStopScheduler(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolStopScheduler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolStopScheduler`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolStopScheduler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolStopSchedulerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolStopSchedulerRequest**](BatchcontrolStopSchedulerRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchcontrolStopWorker

> KalturaControlPanelCommand BatchcontrolStopWorker(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchcontrolStopWorkerRequest(int32(123), "Cause_example", int32(123)) // BatchcontrolStopWorkerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchcontrolApi.BatchcontrolStopWorker(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchcontrolApi.BatchcontrolStopWorker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchcontrolStopWorker`: KalturaControlPanelCommand
    fmt.Fprintf(os.Stdout, "Response from `BatchcontrolApi.BatchcontrolStopWorker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchcontrolStopWorkerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchcontrolStopWorkerRequest**](BatchcontrolStopWorkerRequest.md) |  | 

### Return type

[**KalturaControlPanelCommand**](KalturaControlPanelCommand.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


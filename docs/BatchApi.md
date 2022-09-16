# \BatchApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchAddBulkUploadResult**](BatchApi.md#BatchAddBulkUploadResult) | **Post** /service/batch/action/addBulkUploadResult | 
[**BatchAddMediaInfo**](BatchApi.md#BatchAddMediaInfo) | **Post** /service/batch/action/addMediaInfo | 
[**BatchCheckEntryIsDone**](BatchApi.md#BatchCheckEntryIsDone) | **Post** /service/batch/action/checkEntryIsDone | 
[**BatchCheckFileExists**](BatchApi.md#BatchCheckFileExists) | **Post** /service/batch/action/checkFileExists | 
[**BatchCleanExclusiveJobs**](BatchApi.md#BatchCleanExclusiveJobs) | **Post** /service/batch/action/cleanExclusiveJobs | 
[**BatchCountBulkUploadEntries**](BatchApi.md#BatchCountBulkUploadEntries) | **Post** /service/batch/action/countBulkUploadEntries | 
[**BatchExtendLockExpiration**](BatchApi.md#BatchExtendLockExpiration) | **Post** /service/batch/action/extendLockExpiration | 
[**BatchFreeExclusiveJob**](BatchApi.md#BatchFreeExclusiveJob) | **Post** /service/batch/action/freeExclusiveJob | 
[**BatchGetBulkUploadLastResult**](BatchApi.md#BatchGetBulkUploadLastResult) | **Post** /service/batch/action/getBulkUploadLastResult | 
[**BatchGetExclusiveAlmostDone**](BatchApi.md#BatchGetExclusiveAlmostDone) | **Post** /service/batch/action/getExclusiveAlmostDone | 
[**BatchGetExclusiveJobs**](BatchApi.md#BatchGetExclusiveJobs) | **Post** /service/batch/action/getExclusiveJobs | 
[**BatchGetExclusiveNotificationJobs**](BatchApi.md#BatchGetExclusiveNotificationJobs) | **Post** /service/batch/action/getExclusiveNotificationJobs | 
[**BatchGetQueueSize**](BatchApi.md#BatchGetQueueSize) | **Post** /service/batch/action/getQueueSize | 
[**BatchLogConversion**](BatchApi.md#BatchLogConversion) | **Post** /service/batch/action/logConversion | 
[**BatchPutFile**](BatchApi.md#BatchPutFile) | **Post** /service/batch/action/putFile | 
[**BatchResetJobExecutionAttempts**](BatchApi.md#BatchResetJobExecutionAttempts) | **Post** /service/batch/action/resetJobExecutionAttempts | 
[**BatchSuspendJobs**](BatchApi.md#BatchSuspendJobs) | **Post** /service/batch/action/suspendJobs | 
[**BatchUpdateBulkUploadResults**](BatchApi.md#BatchUpdateBulkUploadResults) | **Post** /service/batch/action/updateBulkUploadResults | 
[**BatchUpdateExclusiveConvertCollectionJob**](BatchApi.md#BatchUpdateExclusiveConvertCollectionJob) | **Post** /service/batch/action/updateExclusiveConvertCollectionJob | 
[**BatchUpdateExclusiveJob**](BatchApi.md#BatchUpdateExclusiveJob) | **Post** /service/batch/action/updateExclusiveJob | 
[**BatchUpdatePartnerLoadTable**](BatchApi.md#BatchUpdatePartnerLoadTable) | **Post** /service/batch/action/updatePartnerLoadTable | 



## BatchAddBulkUploadResult

> KalturaBulkUploadResult BatchAddBulkUploadResult(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchAddBulkUploadResultRequest(*openapiclient.NewKalturaBulkUploadResult()) // BatchAddBulkUploadResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchAddBulkUploadResult(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchAddBulkUploadResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchAddBulkUploadResult`: KalturaBulkUploadResult
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchAddBulkUploadResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchAddBulkUploadResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchAddBulkUploadResultRequest**](BatchAddBulkUploadResultRequest.md) |  | 

### Return type

[**KalturaBulkUploadResult**](KalturaBulkUploadResult.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchAddMediaInfo

> KalturaMediaInfo BatchAddMediaInfo(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchAddMediaInfoRequest(*openapiclient.NewKalturaMediaInfo()) // BatchAddMediaInfoRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchAddMediaInfo(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchAddMediaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchAddMediaInfo`: KalturaMediaInfo
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchAddMediaInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchAddMediaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchAddMediaInfoRequest**](BatchAddMediaInfoRequest.md) |  | 

### Return type

[**KalturaMediaInfo**](KalturaMediaInfo.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCheckEntryIsDone

> bool BatchCheckEntryIsDone(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchCheckEntryIsDoneRequest("BatchJobId_example") // BatchCheckEntryIsDoneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchCheckEntryIsDone(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCheckEntryIsDone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCheckEntryIsDone`: bool
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCheckEntryIsDone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCheckEntryIsDoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchCheckEntryIsDoneRequest**](BatchCheckEntryIsDoneRequest.md) |  | 

### Return type

**bool**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCheckFileExists

> KalturaFileExistsResponse BatchCheckFileExists(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchCheckFileExistsRequest("LocalPath_example", float32(123)) // BatchCheckFileExistsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchCheckFileExists(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCheckFileExists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCheckFileExists`: KalturaFileExistsResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCheckFileExists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCheckFileExistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchCheckFileExistsRequest**](BatchCheckFileExistsRequest.md) |  | 

### Return type

[**KalturaFileExistsResponse**](KalturaFileExistsResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCleanExclusiveJobs

> int32 BatchCleanExclusiveJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchApi.BatchCleanExclusiveJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCleanExclusiveJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCleanExclusiveJobs`: int32
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCleanExclusiveJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCleanExclusiveJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCountBulkUploadEntries

> []KalturaKeyValue BatchCountBulkUploadEntries(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchCountBulkUploadEntriesRequest(int32(123)) // BatchCountBulkUploadEntriesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchCountBulkUploadEntries(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchCountBulkUploadEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCountBulkUploadEntries`: []KalturaKeyValue
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchCountBulkUploadEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCountBulkUploadEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchCountBulkUploadEntriesRequest**](BatchCountBulkUploadEntriesRequest.md) |  | 

### Return type

[**[]KalturaKeyValue**](KalturaKeyValue.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchExtendLockExpiration

> BatchExtendLockExpiration(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchExtendLockExpirationRequest(int32(123), int32(123)) // BatchExtendLockExpirationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchExtendLockExpiration(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchExtendLockExpiration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchExtendLockExpirationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchExtendLockExpirationRequest**](BatchExtendLockExpirationRequest.md) |  | 

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


## BatchFreeExclusiveJob

> KalturaFreeJobResponse BatchFreeExclusiveJob(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchFreeExclusiveJobRequest(int32(123), "JobType_example", *openapiclient.NewKalturaExclusiveLockKey()) // BatchFreeExclusiveJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchFreeExclusiveJob(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchFreeExclusiveJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchFreeExclusiveJob`: KalturaFreeJobResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchFreeExclusiveJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchFreeExclusiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchFreeExclusiveJobRequest**](BatchFreeExclusiveJobRequest.md) |  | 

### Return type

[**KalturaFreeJobResponse**](KalturaFreeJobResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetBulkUploadLastResult

> KalturaBulkUploadResult BatchGetBulkUploadLastResult(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetBulkUploadLastResultRequest(int32(123)) // BatchGetBulkUploadLastResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchGetBulkUploadLastResult(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchGetBulkUploadLastResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetBulkUploadLastResult`: KalturaBulkUploadResult
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchGetBulkUploadLastResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetBulkUploadLastResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetBulkUploadLastResultRequest**](BatchGetBulkUploadLastResultRequest.md) |  | 

### Return type

[**KalturaBulkUploadResult**](KalturaBulkUploadResult.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetExclusiveAlmostDone

> []KalturaBatchJob BatchGetExclusiveAlmostDone(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetExclusiveAlmostDoneRequest(*openapiclient.NewKalturaExclusiveLockKey(), int32(123), int32(123)) // BatchGetExclusiveAlmostDoneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchGetExclusiveAlmostDone(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchGetExclusiveAlmostDone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetExclusiveAlmostDone`: []KalturaBatchJob
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchGetExclusiveAlmostDone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetExclusiveAlmostDoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetExclusiveAlmostDoneRequest**](BatchGetExclusiveAlmostDoneRequest.md) |  | 

### Return type

[**[]KalturaBatchJob**](KalturaBatchJob.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetExclusiveJobs

> []KalturaBatchJob BatchGetExclusiveJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetExclusiveJobsRequest(*openapiclient.NewKalturaExclusiveLockKey(), int32(123), int32(123)) // BatchGetExclusiveJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchGetExclusiveJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchGetExclusiveJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetExclusiveJobs`: []KalturaBatchJob
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchGetExclusiveJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetExclusiveJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetExclusiveJobsRequest**](BatchGetExclusiveJobsRequest.md) |  | 

### Return type

[**[]KalturaBatchJob**](KalturaBatchJob.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetExclusiveNotificationJobs

> KalturaBatchGetExclusiveNotificationJobsResponse BatchGetExclusiveNotificationJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetExclusiveNotificationJobsRequest(*openapiclient.NewKalturaExclusiveLockKey(), int32(123), int32(123)) // BatchGetExclusiveNotificationJobsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchGetExclusiveNotificationJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchGetExclusiveNotificationJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetExclusiveNotificationJobs`: KalturaBatchGetExclusiveNotificationJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchGetExclusiveNotificationJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetExclusiveNotificationJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetExclusiveNotificationJobsRequest**](BatchGetExclusiveNotificationJobsRequest.md) |  | 

### Return type

[**KalturaBatchGetExclusiveNotificationJobsResponse**](KalturaBatchGetExclusiveNotificationJobsResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchGetQueueSize

> int32 BatchGetQueueSize(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetQueueSizeRequest(*openapiclient.NewKalturaWorkerQueueFilter()) // BatchGetQueueSizeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchGetQueueSize(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchGetQueueSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetQueueSize`: int32
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchGetQueueSize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetQueueSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetQueueSizeRequest**](BatchGetQueueSizeRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchLogConversion

> BatchLogConversion(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchLogConversionRequest("Data_example", "FlavorAssetId_example") // BatchLogConversionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchLogConversion(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchLogConversion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchLogConversionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchLogConversionRequest**](BatchLogConversionRequest.md) |  | 

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


## BatchPutFile

> int32 BatchPutFile(ctx).DestPath(destPath).Data(data).Ks(ks).Format(format).Execute()





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
    destPath := "destPath_example" // string | 
    data := os.NewFile(1234, "some_file") // *os.File | 
    ks := "ks_example" // string |  (optional)
    format := int32(56) // int32 | The format of the response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchPutFile(context.Background()).DestPath(destPath).Data(data).Ks(ks).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchPutFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchPutFile`: int32
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchPutFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destPath** | **string** |  | 
 **data** | ***os.File** |  | 
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchResetJobExecutionAttempts

> BatchResetJobExecutionAttempts(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchResetJobExecutionAttemptsRequest(int32(123), "JobType_example", *openapiclient.NewKalturaExclusiveLockKey()) // BatchResetJobExecutionAttemptsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchResetJobExecutionAttempts(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchResetJobExecutionAttempts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchResetJobExecutionAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchResetJobExecutionAttemptsRequest**](BatchResetJobExecutionAttemptsRequest.md) |  | 

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


## BatchSuspendJobs

> BatchSuspendJobs(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchApi.BatchSuspendJobs(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchSuspendJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchSuspendJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

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


## BatchUpdateBulkUploadResults

> int32 BatchUpdateBulkUploadResults(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchGetBulkUploadLastResultRequest(int32(123)) // BatchGetBulkUploadLastResultRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchUpdateBulkUploadResults(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchUpdateBulkUploadResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdateBulkUploadResults`: int32
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchUpdateBulkUploadResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateBulkUploadResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchGetBulkUploadLastResultRequest**](BatchGetBulkUploadLastResultRequest.md) |  | 

### Return type

**int32**

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdateExclusiveConvertCollectionJob

> KalturaBatchJob BatchUpdateExclusiveConvertCollectionJob(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchUpdateExclusiveConvertCollectionJobRequest(int32(123), *openapiclient.NewKalturaBatchJob(), *openapiclient.NewKalturaExclusiveLockKey()) // BatchUpdateExclusiveConvertCollectionJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchUpdateExclusiveConvertCollectionJob(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchUpdateExclusiveConvertCollectionJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdateExclusiveConvertCollectionJob`: KalturaBatchJob
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchUpdateExclusiveConvertCollectionJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateExclusiveConvertCollectionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchUpdateExclusiveConvertCollectionJobRequest**](BatchUpdateExclusiveConvertCollectionJobRequest.md) |  | 

### Return type

[**KalturaBatchJob**](KalturaBatchJob.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdateExclusiveJob

> KalturaBatchJob BatchUpdateExclusiveJob(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewBatchUpdateExclusiveJobRequest(int32(123), *openapiclient.NewKalturaBatchJob(), *openapiclient.NewKalturaExclusiveLockKey()) // BatchUpdateExclusiveJobRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BatchApi.BatchUpdateExclusiveJob(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchUpdateExclusiveJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchUpdateExclusiveJob`: KalturaBatchJob
    fmt.Fprintf(os.Stdout, "Response from `BatchApi.BatchUpdateExclusiveJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdateExclusiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**BatchUpdateExclusiveJobRequest**](BatchUpdateExclusiveJobRequest.md) |  | 

### Return type

[**KalturaBatchJob**](KalturaBatchJob.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchUpdatePartnerLoadTable

> BatchUpdatePartnerLoadTable(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()





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
    resp, r, err := apiClient.BatchApi.BatchUpdatePartnerLoadTable(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).BatchCleanExclusiveJobsRequest(batchCleanExclusiveJobsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BatchApi.BatchUpdatePartnerLoadTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchUpdatePartnerLoadTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **batchCleanExclusiveJobsRequest** | [**BatchCleanExclusiveJobsRequest**](BatchCleanExclusiveJobsRequest.md) |  | 

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


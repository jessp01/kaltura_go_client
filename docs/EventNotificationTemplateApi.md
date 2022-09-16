# \EventNotificationTemplateApi

All URIs are relative to *https://www.kaltura.com/api_v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventNotificationTemplateAdd**](EventNotificationTemplateApi.md#EventNotificationTemplateAdd) | **Post** /service/eventnotification_eventnotificationtemplate/action/add | 
[**EventNotificationTemplateClone**](EventNotificationTemplateApi.md#EventNotificationTemplateClone) | **Post** /service/eventnotification_eventnotificationtemplate/action/clone | 
[**EventNotificationTemplateDelete**](EventNotificationTemplateApi.md#EventNotificationTemplateDelete) | **Post** /service/eventnotification_eventnotificationtemplate/action/delete | 
[**EventNotificationTemplateDispatch**](EventNotificationTemplateApi.md#EventNotificationTemplateDispatch) | **Post** /service/eventnotification_eventnotificationtemplate/action/dispatch | 
[**EventNotificationTemplateGet**](EventNotificationTemplateApi.md#EventNotificationTemplateGet) | **Post** /service/eventnotification_eventnotificationtemplate/action/get | 
[**EventNotificationTemplateList**](EventNotificationTemplateApi.md#EventNotificationTemplateList) | **Post** /service/eventnotification_eventnotificationtemplate/action/list | 
[**EventNotificationTemplateListByPartner**](EventNotificationTemplateApi.md#EventNotificationTemplateListByPartner) | **Post** /service/eventnotification_eventnotificationtemplate/action/listByPartner | 
[**EventNotificationTemplateListTemplates**](EventNotificationTemplateApi.md#EventNotificationTemplateListTemplates) | **Post** /service/eventnotification_eventnotificationtemplate/action/listTemplates | 
[**EventNotificationTemplateRegister**](EventNotificationTemplateApi.md#EventNotificationTemplateRegister) | **Post** /service/eventnotification_eventnotificationtemplate/action/register | 
[**EventNotificationTemplateSendCommand**](EventNotificationTemplateApi.md#EventNotificationTemplateSendCommand) | **Post** /service/eventnotification_eventnotificationtemplate/action/sendCommand | 
[**EventNotificationTemplateUpdate**](EventNotificationTemplateApi.md#EventNotificationTemplateUpdate) | **Post** /service/eventnotification_eventnotificationtemplate/action/update | 
[**EventNotificationTemplateUpdateStatus**](EventNotificationTemplateApi.md#EventNotificationTemplateUpdateStatus) | **Post** /service/eventnotification_eventnotificationtemplate/action/updateStatus | 



## EventNotificationTemplateAdd

> KalturaEventNotificationTemplate EventNotificationTemplateAdd(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateAddRequest(*openapiclient.NewKalturaEventNotificationTemplate()) // EventNotificationTemplateAddRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateAdd(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateAdd`: KalturaEventNotificationTemplate
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateAddRequest**](EventNotificationTemplateAddRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateClone

> KalturaEventNotificationTemplate EventNotificationTemplateClone(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateCloneRequest(int32(123)) // EventNotificationTemplateCloneRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateClone(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateClone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateClone`: KalturaEventNotificationTemplate
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateCloneRequest**](EventNotificationTemplateCloneRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateDelete

> EventNotificationTemplateDelete(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateDelete(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateDeleteRequest struct via the builder pattern


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


## EventNotificationTemplateDispatch

> int32 EventNotificationTemplateDispatch(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateDispatchRequest(int32(123), *openapiclient.NewKalturaEventNotificationScope()) // EventNotificationTemplateDispatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateDispatch(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateDispatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateDispatch`: int32
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateDispatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateDispatchRequest**](EventNotificationTemplateDispatchRequest.md) |  | 

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


## EventNotificationTemplateGet

> KalturaEventNotificationTemplate EventNotificationTemplateGet(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateGet(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateGet`: KalturaEventNotificationTemplate
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**AccessControlDeleteRequest**](AccessControlDeleteRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateList

> KalturaEventNotificationTemplateListResponse EventNotificationTemplateList(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateListRequest() // EventNotificationTemplateListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateList(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateList`: KalturaEventNotificationTemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateListRequest**](EventNotificationTemplateListRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplateListResponse**](KalturaEventNotificationTemplateListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateListByPartner

> KalturaEventNotificationTemplateListResponse EventNotificationTemplateListByPartner(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()



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
    body := *openapiclient.NewDistributionProfileListByPartnerRequest() // DistributionProfileListByPartnerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateListByPartner(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateListByPartner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateListByPartner`: KalturaEventNotificationTemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateListByPartner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateListByPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**DistributionProfileListByPartnerRequest**](DistributionProfileListByPartnerRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplateListResponse**](KalturaEventNotificationTemplateListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateListTemplates

> KalturaEventNotificationTemplateListResponse EventNotificationTemplateListTemplates(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateListRequest() // EventNotificationTemplateListRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateListTemplates(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateListTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateListTemplates`: KalturaEventNotificationTemplateListResponse
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateListTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateListTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateListRequest**](EventNotificationTemplateListRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplateListResponse**](KalturaEventNotificationTemplateListResponse.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateRegister

> KalturaPushNotificationData EventNotificationTemplateRegister(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateRegisterRequest("NotificationTemplateSystemName_example", *openapiclient.NewKalturaPushNotificationParams()) // EventNotificationTemplateRegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateRegister(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateRegister``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateRegister`: KalturaPushNotificationData
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateRegister`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateRegisterRequest**](EventNotificationTemplateRegisterRequest.md) |  | 

### Return type

[**KalturaPushNotificationData**](KalturaPushNotificationData.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateSendCommand

> EventNotificationTemplateSendCommand(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateSendCommandRequest("Command_example", "NotificationTemplateSystemName_example", *openapiclient.NewKalturaPushNotificationParams()) // EventNotificationTemplateSendCommandRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateSendCommand(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateSendCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateSendCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateSendCommandRequest**](EventNotificationTemplateSendCommandRequest.md) |  | 

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


## EventNotificationTemplateUpdate

> KalturaEventNotificationTemplate EventNotificationTemplateUpdate(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateUpdateRequest(*openapiclient.NewKalturaEventNotificationTemplate(), int32(123)) // EventNotificationTemplateUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateUpdate(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateUpdate`: KalturaEventNotificationTemplate
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateUpdateRequest**](EventNotificationTemplateUpdateRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventNotificationTemplateUpdateStatus

> KalturaEventNotificationTemplate EventNotificationTemplateUpdateStatus(ctx).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()





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
    body := *openapiclient.NewEventNotificationTemplateUpdateStatusRequest(int32(123), int32(123)) // EventNotificationTemplateUpdateStatusRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventNotificationTemplateApi.EventNotificationTemplateUpdateStatus(context.Background()).Ks(ks).Format(format).ClientTag(clientTag).PartnerId(partnerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventNotificationTemplateApi.EventNotificationTemplateUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventNotificationTemplateUpdateStatus`: KalturaEventNotificationTemplate
    fmt.Fprintf(os.Stdout, "Response from `EventNotificationTemplateApi.EventNotificationTemplateUpdateStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventNotificationTemplateUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ks** | **string** |  | 
 **format** | **int32** | The format of the response | 
 **clientTag** | **string** |  | [default to &quot;devkcom&quot;]
 **partnerId** | **int32** |  | 
 **body** | [**EventNotificationTemplateUpdateStatusRequest**](EventNotificationTemplateUpdateStatusRequest.md) |  | 

### Return type

[**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md)

### Authorization

[ks](../README.md#ks)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


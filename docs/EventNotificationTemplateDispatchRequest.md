# EventNotificationTemplateDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Scope** | [**KalturaEventNotificationScope**](KalturaEventNotificationScope.md) |  | 

## Methods

### NewEventNotificationTemplateDispatchRequest

`func NewEventNotificationTemplateDispatchRequest(id int32, scope KalturaEventNotificationScope, ) *EventNotificationTemplateDispatchRequest`

NewEventNotificationTemplateDispatchRequest instantiates a new EventNotificationTemplateDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationTemplateDispatchRequestWithDefaults

`func NewEventNotificationTemplateDispatchRequestWithDefaults() *EventNotificationTemplateDispatchRequest`

NewEventNotificationTemplateDispatchRequestWithDefaults instantiates a new EventNotificationTemplateDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EventNotificationTemplateDispatchRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventNotificationTemplateDispatchRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventNotificationTemplateDispatchRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetScope

`func (o *EventNotificationTemplateDispatchRequest) GetScope() KalturaEventNotificationScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EventNotificationTemplateDispatchRequest) GetScopeOk() (*KalturaEventNotificationScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EventNotificationTemplateDispatchRequest) SetScope(v KalturaEventNotificationScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



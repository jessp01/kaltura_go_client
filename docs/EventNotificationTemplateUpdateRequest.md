# EventNotificationTemplateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotificationTemplate** | [**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewEventNotificationTemplateUpdateRequest

`func NewEventNotificationTemplateUpdateRequest(eventNotificationTemplate KalturaEventNotificationTemplate, id int32, ) *EventNotificationTemplateUpdateRequest`

NewEventNotificationTemplateUpdateRequest instantiates a new EventNotificationTemplateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationTemplateUpdateRequestWithDefaults

`func NewEventNotificationTemplateUpdateRequestWithDefaults() *EventNotificationTemplateUpdateRequest`

NewEventNotificationTemplateUpdateRequestWithDefaults instantiates a new EventNotificationTemplateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotificationTemplate

`func (o *EventNotificationTemplateUpdateRequest) GetEventNotificationTemplate() KalturaEventNotificationTemplate`

GetEventNotificationTemplate returns the EventNotificationTemplate field if non-nil, zero value otherwise.

### GetEventNotificationTemplateOk

`func (o *EventNotificationTemplateUpdateRequest) GetEventNotificationTemplateOk() (*KalturaEventNotificationTemplate, bool)`

GetEventNotificationTemplateOk returns a tuple with the EventNotificationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationTemplate

`func (o *EventNotificationTemplateUpdateRequest) SetEventNotificationTemplate(v KalturaEventNotificationTemplate)`

SetEventNotificationTemplate sets EventNotificationTemplate field to given value.


### GetId

`func (o *EventNotificationTemplateUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventNotificationTemplateUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventNotificationTemplateUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



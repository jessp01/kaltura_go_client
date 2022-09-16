# EventNotificationTemplateCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotificationTemplate** | Pointer to [**KalturaEventNotificationTemplate**](KalturaEventNotificationTemplate.md) |  | [optional] 
**Id** | **int32** |  | 

## Methods

### NewEventNotificationTemplateCloneRequest

`func NewEventNotificationTemplateCloneRequest(id int32, ) *EventNotificationTemplateCloneRequest`

NewEventNotificationTemplateCloneRequest instantiates a new EventNotificationTemplateCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationTemplateCloneRequestWithDefaults

`func NewEventNotificationTemplateCloneRequestWithDefaults() *EventNotificationTemplateCloneRequest`

NewEventNotificationTemplateCloneRequestWithDefaults instantiates a new EventNotificationTemplateCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotificationTemplate

`func (o *EventNotificationTemplateCloneRequest) GetEventNotificationTemplate() KalturaEventNotificationTemplate`

GetEventNotificationTemplate returns the EventNotificationTemplate field if non-nil, zero value otherwise.

### GetEventNotificationTemplateOk

`func (o *EventNotificationTemplateCloneRequest) GetEventNotificationTemplateOk() (*KalturaEventNotificationTemplate, bool)`

GetEventNotificationTemplateOk returns a tuple with the EventNotificationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationTemplate

`func (o *EventNotificationTemplateCloneRequest) SetEventNotificationTemplate(v KalturaEventNotificationTemplate)`

SetEventNotificationTemplate sets EventNotificationTemplate field to given value.

### HasEventNotificationTemplate

`func (o *EventNotificationTemplateCloneRequest) HasEventNotificationTemplate() bool`

HasEventNotificationTemplate returns a boolean if a field has been set.

### GetId

`func (o *EventNotificationTemplateCloneRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventNotificationTemplateCloneRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventNotificationTemplateCloneRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



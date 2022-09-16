# KalturaHttpNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotificationJobId** | Pointer to **int32** | ID of the batch job that execute the notification | [optional] 
**EventObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaEventNotificationEventObjectType&#x60;  Object type that triggered the notification | [optional] 
**EventType** | Pointer to **string** | Enum Type: &#x60;KalturaEventNotificationEventType&#x60;  Ecent type that triggered the notification | [optional] 
**Object** | Pointer to **map[string]interface{}** |  | [optional] 
**TemplateId** | Pointer to **int32** | ID of the template that triggered the notification | [optional] 
**TemplateName** | Pointer to **string** | Name of the template that triggered the notification | [optional] 
**TemplateSystemName** | Pointer to **string** | System name of the template that triggered the notification | [optional] 

## Methods

### NewKalturaHttpNotification

`func NewKalturaHttpNotification() *KalturaHttpNotification`

NewKalturaHttpNotification instantiates a new KalturaHttpNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaHttpNotificationWithDefaults

`func NewKalturaHttpNotificationWithDefaults() *KalturaHttpNotification`

NewKalturaHttpNotificationWithDefaults instantiates a new KalturaHttpNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotificationJobId

`func (o *KalturaHttpNotification) GetEventNotificationJobId() int32`

GetEventNotificationJobId returns the EventNotificationJobId field if non-nil, zero value otherwise.

### GetEventNotificationJobIdOk

`func (o *KalturaHttpNotification) GetEventNotificationJobIdOk() (*int32, bool)`

GetEventNotificationJobIdOk returns a tuple with the EventNotificationJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationJobId

`func (o *KalturaHttpNotification) SetEventNotificationJobId(v int32)`

SetEventNotificationJobId sets EventNotificationJobId field to given value.

### HasEventNotificationJobId

`func (o *KalturaHttpNotification) HasEventNotificationJobId() bool`

HasEventNotificationJobId returns a boolean if a field has been set.

### GetEventObjectType

`func (o *KalturaHttpNotification) GetEventObjectType() string`

GetEventObjectType returns the EventObjectType field if non-nil, zero value otherwise.

### GetEventObjectTypeOk

`func (o *KalturaHttpNotification) GetEventObjectTypeOk() (*string, bool)`

GetEventObjectTypeOk returns a tuple with the EventObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventObjectType

`func (o *KalturaHttpNotification) SetEventObjectType(v string)`

SetEventObjectType sets EventObjectType field to given value.

### HasEventObjectType

`func (o *KalturaHttpNotification) HasEventObjectType() bool`

HasEventObjectType returns a boolean if a field has been set.

### GetEventType

`func (o *KalturaHttpNotification) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *KalturaHttpNotification) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *KalturaHttpNotification) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *KalturaHttpNotification) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetObject

`func (o *KalturaHttpNotification) GetObject() map[string]interface{}`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *KalturaHttpNotification) GetObjectOk() (*map[string]interface{}, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *KalturaHttpNotification) SetObject(v map[string]interface{})`

SetObject sets Object field to given value.

### HasObject

`func (o *KalturaHttpNotification) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetTemplateId

`func (o *KalturaHttpNotification) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *KalturaHttpNotification) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *KalturaHttpNotification) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *KalturaHttpNotification) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateName

`func (o *KalturaHttpNotification) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *KalturaHttpNotification) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *KalturaHttpNotification) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *KalturaHttpNotification) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateSystemName

`func (o *KalturaHttpNotification) GetTemplateSystemName() string`

GetTemplateSystemName returns the TemplateSystemName field if non-nil, zero value otherwise.

### GetTemplateSystemNameOk

`func (o *KalturaHttpNotification) GetTemplateSystemNameOk() (*string, bool)`

GetTemplateSystemNameOk returns a tuple with the TemplateSystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateSystemName

`func (o *KalturaHttpNotification) SetTemplateSystemName(v string)`

SetTemplateSystemName sets TemplateSystemName field to given value.

### HasTemplateSystemName

`func (o *KalturaHttpNotification) HasTemplateSystemName() bool`

HasTemplateSystemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



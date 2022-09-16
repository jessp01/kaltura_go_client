# KalturaEventNotificationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticDispatchEnabled** | Pointer to **bool** | Define that the template could be dispatched automatically by the system | [optional] 
**ContentParameters** | Pointer to [**[]KalturaEventNotificationParameter**](KalturaEventNotificationParameter.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**EventConditions** | Pointer to [**[]KalturaCondition**](KalturaCondition.md) |  | [optional] 
**EventObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaEventNotificationEventObjectType&#x60;  Define the object that raied the event that should trigger this notification | [optional] 
**EventType** | Pointer to **string** | Enum Type: &#x60;KalturaEventNotificationEventType&#x60;  Define the event that should trigger this notification | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ManualDispatchEnabled** | Pointer to **bool** | Define that the template could be dispatched manually from the API | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaEventNotificationTemplateStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaEventNotificationTemplateType&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserParameters** | Pointer to [**[]KalturaEventNotificationParameter**](KalturaEventNotificationParameter.md) |  | [optional] 

## Methods

### NewKalturaEventNotificationTemplate

`func NewKalturaEventNotificationTemplate() *KalturaEventNotificationTemplate`

NewKalturaEventNotificationTemplate instantiates a new KalturaEventNotificationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEventNotificationTemplateWithDefaults

`func NewKalturaEventNotificationTemplateWithDefaults() *KalturaEventNotificationTemplate`

NewKalturaEventNotificationTemplateWithDefaults instantiates a new KalturaEventNotificationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomaticDispatchEnabled

`func (o *KalturaEventNotificationTemplate) GetAutomaticDispatchEnabled() bool`

GetAutomaticDispatchEnabled returns the AutomaticDispatchEnabled field if non-nil, zero value otherwise.

### GetAutomaticDispatchEnabledOk

`func (o *KalturaEventNotificationTemplate) GetAutomaticDispatchEnabledOk() (*bool, bool)`

GetAutomaticDispatchEnabledOk returns a tuple with the AutomaticDispatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticDispatchEnabled

`func (o *KalturaEventNotificationTemplate) SetAutomaticDispatchEnabled(v bool)`

SetAutomaticDispatchEnabled sets AutomaticDispatchEnabled field to given value.

### HasAutomaticDispatchEnabled

`func (o *KalturaEventNotificationTemplate) HasAutomaticDispatchEnabled() bool`

HasAutomaticDispatchEnabled returns a boolean if a field has been set.

### GetContentParameters

`func (o *KalturaEventNotificationTemplate) GetContentParameters() []KalturaEventNotificationParameter`

GetContentParameters returns the ContentParameters field if non-nil, zero value otherwise.

### GetContentParametersOk

`func (o *KalturaEventNotificationTemplate) GetContentParametersOk() (*[]KalturaEventNotificationParameter, bool)`

GetContentParametersOk returns a tuple with the ContentParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParameters

`func (o *KalturaEventNotificationTemplate) SetContentParameters(v []KalturaEventNotificationParameter)`

SetContentParameters sets ContentParameters field to given value.

### HasContentParameters

`func (o *KalturaEventNotificationTemplate) HasContentParameters() bool`

HasContentParameters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaEventNotificationTemplate) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaEventNotificationTemplate) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaEventNotificationTemplate) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaEventNotificationTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaEventNotificationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaEventNotificationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaEventNotificationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaEventNotificationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEventConditions

`func (o *KalturaEventNotificationTemplate) GetEventConditions() []KalturaCondition`

GetEventConditions returns the EventConditions field if non-nil, zero value otherwise.

### GetEventConditionsOk

`func (o *KalturaEventNotificationTemplate) GetEventConditionsOk() (*[]KalturaCondition, bool)`

GetEventConditionsOk returns a tuple with the EventConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventConditions

`func (o *KalturaEventNotificationTemplate) SetEventConditions(v []KalturaCondition)`

SetEventConditions sets EventConditions field to given value.

### HasEventConditions

`func (o *KalturaEventNotificationTemplate) HasEventConditions() bool`

HasEventConditions returns a boolean if a field has been set.

### GetEventObjectType

`func (o *KalturaEventNotificationTemplate) GetEventObjectType() string`

GetEventObjectType returns the EventObjectType field if non-nil, zero value otherwise.

### GetEventObjectTypeOk

`func (o *KalturaEventNotificationTemplate) GetEventObjectTypeOk() (*string, bool)`

GetEventObjectTypeOk returns a tuple with the EventObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventObjectType

`func (o *KalturaEventNotificationTemplate) SetEventObjectType(v string)`

SetEventObjectType sets EventObjectType field to given value.

### HasEventObjectType

`func (o *KalturaEventNotificationTemplate) HasEventObjectType() bool`

HasEventObjectType returns a boolean if a field has been set.

### GetEventType

`func (o *KalturaEventNotificationTemplate) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *KalturaEventNotificationTemplate) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *KalturaEventNotificationTemplate) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *KalturaEventNotificationTemplate) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *KalturaEventNotificationTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaEventNotificationTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaEventNotificationTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaEventNotificationTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManualDispatchEnabled

`func (o *KalturaEventNotificationTemplate) GetManualDispatchEnabled() bool`

GetManualDispatchEnabled returns the ManualDispatchEnabled field if non-nil, zero value otherwise.

### GetManualDispatchEnabledOk

`func (o *KalturaEventNotificationTemplate) GetManualDispatchEnabledOk() (*bool, bool)`

GetManualDispatchEnabledOk returns a tuple with the ManualDispatchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualDispatchEnabled

`func (o *KalturaEventNotificationTemplate) SetManualDispatchEnabled(v bool)`

SetManualDispatchEnabled sets ManualDispatchEnabled field to given value.

### HasManualDispatchEnabled

`func (o *KalturaEventNotificationTemplate) HasManualDispatchEnabled() bool`

HasManualDispatchEnabled returns a boolean if a field has been set.

### GetName

`func (o *KalturaEventNotificationTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaEventNotificationTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaEventNotificationTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaEventNotificationTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaEventNotificationTemplate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaEventNotificationTemplate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaEventNotificationTemplate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaEventNotificationTemplate) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaEventNotificationTemplate) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaEventNotificationTemplate) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaEventNotificationTemplate) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaEventNotificationTemplate) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaEventNotificationTemplate) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaEventNotificationTemplate) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaEventNotificationTemplate) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaEventNotificationTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaEventNotificationTemplate) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaEventNotificationTemplate) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaEventNotificationTemplate) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaEventNotificationTemplate) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetType

`func (o *KalturaEventNotificationTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaEventNotificationTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaEventNotificationTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaEventNotificationTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaEventNotificationTemplate) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaEventNotificationTemplate) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaEventNotificationTemplate) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaEventNotificationTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserParameters

`func (o *KalturaEventNotificationTemplate) GetUserParameters() []KalturaEventNotificationParameter`

GetUserParameters returns the UserParameters field if non-nil, zero value otherwise.

### GetUserParametersOk

`func (o *KalturaEventNotificationTemplate) GetUserParametersOk() (*[]KalturaEventNotificationParameter, bool)`

GetUserParametersOk returns a tuple with the UserParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserParameters

`func (o *KalturaEventNotificationTemplate) SetUserParameters(v []KalturaEventNotificationParameter)`

SetUserParameters sets UserParameters field to given value.

### HasUserParameters

`func (o *KalturaEventNotificationTemplate) HasUserParameters() bool`

HasUserParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



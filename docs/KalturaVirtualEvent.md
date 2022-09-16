# KalturaVirtualEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminsGroupIds** | Pointer to **string** |  | [optional] 
**AgendaScheduleEventId** | Pointer to **int32** |  | [optional] 
**AttendeesGroupIds** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DeletionDueDate** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MainEventScheduleEventId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**RegistrationFormSchema** | Pointer to **string** | JSON-Schema of the Registration Form | [optional] 
**RegistrationScheduleEventId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVirtualEventStatus&#x60; | [optional] [readonly] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaVirtualEvent

`func NewKalturaVirtualEvent() *KalturaVirtualEvent`

NewKalturaVirtualEvent instantiates a new KalturaVirtualEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaVirtualEventWithDefaults

`func NewKalturaVirtualEventWithDefaults() *KalturaVirtualEvent`

NewKalturaVirtualEventWithDefaults instantiates a new KalturaVirtualEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminsGroupIds

`func (o *KalturaVirtualEvent) GetAdminsGroupIds() string`

GetAdminsGroupIds returns the AdminsGroupIds field if non-nil, zero value otherwise.

### GetAdminsGroupIdsOk

`func (o *KalturaVirtualEvent) GetAdminsGroupIdsOk() (*string, bool)`

GetAdminsGroupIdsOk returns a tuple with the AdminsGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminsGroupIds

`func (o *KalturaVirtualEvent) SetAdminsGroupIds(v string)`

SetAdminsGroupIds sets AdminsGroupIds field to given value.

### HasAdminsGroupIds

`func (o *KalturaVirtualEvent) HasAdminsGroupIds() bool`

HasAdminsGroupIds returns a boolean if a field has been set.

### GetAgendaScheduleEventId

`func (o *KalturaVirtualEvent) GetAgendaScheduleEventId() int32`

GetAgendaScheduleEventId returns the AgendaScheduleEventId field if non-nil, zero value otherwise.

### GetAgendaScheduleEventIdOk

`func (o *KalturaVirtualEvent) GetAgendaScheduleEventIdOk() (*int32, bool)`

GetAgendaScheduleEventIdOk returns a tuple with the AgendaScheduleEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgendaScheduleEventId

`func (o *KalturaVirtualEvent) SetAgendaScheduleEventId(v int32)`

SetAgendaScheduleEventId sets AgendaScheduleEventId field to given value.

### HasAgendaScheduleEventId

`func (o *KalturaVirtualEvent) HasAgendaScheduleEventId() bool`

HasAgendaScheduleEventId returns a boolean if a field has been set.

### GetAttendeesGroupIds

`func (o *KalturaVirtualEvent) GetAttendeesGroupIds() string`

GetAttendeesGroupIds returns the AttendeesGroupIds field if non-nil, zero value otherwise.

### GetAttendeesGroupIdsOk

`func (o *KalturaVirtualEvent) GetAttendeesGroupIdsOk() (*string, bool)`

GetAttendeesGroupIdsOk returns a tuple with the AttendeesGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttendeesGroupIds

`func (o *KalturaVirtualEvent) SetAttendeesGroupIds(v string)`

SetAttendeesGroupIds sets AttendeesGroupIds field to given value.

### HasAttendeesGroupIds

`func (o *KalturaVirtualEvent) HasAttendeesGroupIds() bool`

HasAttendeesGroupIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaVirtualEvent) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaVirtualEvent) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaVirtualEvent) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaVirtualEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletionDueDate

`func (o *KalturaVirtualEvent) GetDeletionDueDate() int32`

GetDeletionDueDate returns the DeletionDueDate field if non-nil, zero value otherwise.

### GetDeletionDueDateOk

`func (o *KalturaVirtualEvent) GetDeletionDueDateOk() (*int32, bool)`

GetDeletionDueDateOk returns a tuple with the DeletionDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDueDate

`func (o *KalturaVirtualEvent) SetDeletionDueDate(v int32)`

SetDeletionDueDate sets DeletionDueDate field to given value.

### HasDeletionDueDate

`func (o *KalturaVirtualEvent) HasDeletionDueDate() bool`

HasDeletionDueDate returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaVirtualEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaVirtualEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaVirtualEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaVirtualEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaVirtualEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaVirtualEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaVirtualEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaVirtualEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMainEventScheduleEventId

`func (o *KalturaVirtualEvent) GetMainEventScheduleEventId() int32`

GetMainEventScheduleEventId returns the MainEventScheduleEventId field if non-nil, zero value otherwise.

### GetMainEventScheduleEventIdOk

`func (o *KalturaVirtualEvent) GetMainEventScheduleEventIdOk() (*int32, bool)`

GetMainEventScheduleEventIdOk returns a tuple with the MainEventScheduleEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainEventScheduleEventId

`func (o *KalturaVirtualEvent) SetMainEventScheduleEventId(v int32)`

SetMainEventScheduleEventId sets MainEventScheduleEventId field to given value.

### HasMainEventScheduleEventId

`func (o *KalturaVirtualEvent) HasMainEventScheduleEventId() bool`

HasMainEventScheduleEventId returns a boolean if a field has been set.

### GetName

`func (o *KalturaVirtualEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaVirtualEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaVirtualEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaVirtualEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaVirtualEvent) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaVirtualEvent) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaVirtualEvent) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaVirtualEvent) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetRegistrationFormSchema

`func (o *KalturaVirtualEvent) GetRegistrationFormSchema() string`

GetRegistrationFormSchema returns the RegistrationFormSchema field if non-nil, zero value otherwise.

### GetRegistrationFormSchemaOk

`func (o *KalturaVirtualEvent) GetRegistrationFormSchemaOk() (*string, bool)`

GetRegistrationFormSchemaOk returns a tuple with the RegistrationFormSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationFormSchema

`func (o *KalturaVirtualEvent) SetRegistrationFormSchema(v string)`

SetRegistrationFormSchema sets RegistrationFormSchema field to given value.

### HasRegistrationFormSchema

`func (o *KalturaVirtualEvent) HasRegistrationFormSchema() bool`

HasRegistrationFormSchema returns a boolean if a field has been set.

### GetRegistrationScheduleEventId

`func (o *KalturaVirtualEvent) GetRegistrationScheduleEventId() int32`

GetRegistrationScheduleEventId returns the RegistrationScheduleEventId field if non-nil, zero value otherwise.

### GetRegistrationScheduleEventIdOk

`func (o *KalturaVirtualEvent) GetRegistrationScheduleEventIdOk() (*int32, bool)`

GetRegistrationScheduleEventIdOk returns a tuple with the RegistrationScheduleEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationScheduleEventId

`func (o *KalturaVirtualEvent) SetRegistrationScheduleEventId(v int32)`

SetRegistrationScheduleEventId sets RegistrationScheduleEventId field to given value.

### HasRegistrationScheduleEventId

`func (o *KalturaVirtualEvent) HasRegistrationScheduleEventId() bool`

HasRegistrationScheduleEventId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaVirtualEvent) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaVirtualEvent) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaVirtualEvent) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaVirtualEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaVirtualEvent) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaVirtualEvent) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaVirtualEvent) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaVirtualEvent) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaVirtualEvent) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaVirtualEvent) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaVirtualEvent) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaVirtualEvent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



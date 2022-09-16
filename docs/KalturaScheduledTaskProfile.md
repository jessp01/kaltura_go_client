# KalturaScheduledTaskProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LastExecutionStartedAt** | Pointer to **int32** |  | [optional] 
**MaxTotalCountAllowed** | Pointer to **int32** | The maximum number of result count allowed to be processed by this profile per execution | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectFilter** | Pointer to [**KalturaFilter**](KalturaFilter.md) |  | [optional] 
**ObjectFilterEngineType** | Pointer to **string** | Enum Type: &#x60;KalturaObjectFilterEngineType&#x60;  The type of engine to use to list objects using the given \&quot;objectFilter\&quot; | [optional] 
**ObjectTasks** | Pointer to [**[]KalturaObjectTask**](KalturaObjectTask.md) |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaScheduledTaskProfileStatus&#x60; | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaScheduledTaskProfile

`func NewKalturaScheduledTaskProfile() *KalturaScheduledTaskProfile`

NewKalturaScheduledTaskProfile instantiates a new KalturaScheduledTaskProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaScheduledTaskProfileWithDefaults

`func NewKalturaScheduledTaskProfileWithDefaults() *KalturaScheduledTaskProfile`

NewKalturaScheduledTaskProfileWithDefaults instantiates a new KalturaScheduledTaskProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaScheduledTaskProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaScheduledTaskProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaScheduledTaskProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaScheduledTaskProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaScheduledTaskProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaScheduledTaskProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaScheduledTaskProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaScheduledTaskProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaScheduledTaskProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaScheduledTaskProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaScheduledTaskProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaScheduledTaskProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastExecutionStartedAt

`func (o *KalturaScheduledTaskProfile) GetLastExecutionStartedAt() int32`

GetLastExecutionStartedAt returns the LastExecutionStartedAt field if non-nil, zero value otherwise.

### GetLastExecutionStartedAtOk

`func (o *KalturaScheduledTaskProfile) GetLastExecutionStartedAtOk() (*int32, bool)`

GetLastExecutionStartedAtOk returns a tuple with the LastExecutionStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionStartedAt

`func (o *KalturaScheduledTaskProfile) SetLastExecutionStartedAt(v int32)`

SetLastExecutionStartedAt sets LastExecutionStartedAt field to given value.

### HasLastExecutionStartedAt

`func (o *KalturaScheduledTaskProfile) HasLastExecutionStartedAt() bool`

HasLastExecutionStartedAt returns a boolean if a field has been set.

### GetMaxTotalCountAllowed

`func (o *KalturaScheduledTaskProfile) GetMaxTotalCountAllowed() int32`

GetMaxTotalCountAllowed returns the MaxTotalCountAllowed field if non-nil, zero value otherwise.

### GetMaxTotalCountAllowedOk

`func (o *KalturaScheduledTaskProfile) GetMaxTotalCountAllowedOk() (*int32, bool)`

GetMaxTotalCountAllowedOk returns a tuple with the MaxTotalCountAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalCountAllowed

`func (o *KalturaScheduledTaskProfile) SetMaxTotalCountAllowed(v int32)`

SetMaxTotalCountAllowed sets MaxTotalCountAllowed field to given value.

### HasMaxTotalCountAllowed

`func (o *KalturaScheduledTaskProfile) HasMaxTotalCountAllowed() bool`

HasMaxTotalCountAllowed returns a boolean if a field has been set.

### GetName

`func (o *KalturaScheduledTaskProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaScheduledTaskProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaScheduledTaskProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaScheduledTaskProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectFilter

`func (o *KalturaScheduledTaskProfile) GetObjectFilter() KalturaFilter`

GetObjectFilter returns the ObjectFilter field if non-nil, zero value otherwise.

### GetObjectFilterOk

`func (o *KalturaScheduledTaskProfile) GetObjectFilterOk() (*KalturaFilter, bool)`

GetObjectFilterOk returns a tuple with the ObjectFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFilter

`func (o *KalturaScheduledTaskProfile) SetObjectFilter(v KalturaFilter)`

SetObjectFilter sets ObjectFilter field to given value.

### HasObjectFilter

`func (o *KalturaScheduledTaskProfile) HasObjectFilter() bool`

HasObjectFilter returns a boolean if a field has been set.

### GetObjectFilterEngineType

`func (o *KalturaScheduledTaskProfile) GetObjectFilterEngineType() string`

GetObjectFilterEngineType returns the ObjectFilterEngineType field if non-nil, zero value otherwise.

### GetObjectFilterEngineTypeOk

`func (o *KalturaScheduledTaskProfile) GetObjectFilterEngineTypeOk() (*string, bool)`

GetObjectFilterEngineTypeOk returns a tuple with the ObjectFilterEngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFilterEngineType

`func (o *KalturaScheduledTaskProfile) SetObjectFilterEngineType(v string)`

SetObjectFilterEngineType sets ObjectFilterEngineType field to given value.

### HasObjectFilterEngineType

`func (o *KalturaScheduledTaskProfile) HasObjectFilterEngineType() bool`

HasObjectFilterEngineType returns a boolean if a field has been set.

### GetObjectTasks

`func (o *KalturaScheduledTaskProfile) GetObjectTasks() []KalturaObjectTask`

GetObjectTasks returns the ObjectTasks field if non-nil, zero value otherwise.

### GetObjectTasksOk

`func (o *KalturaScheduledTaskProfile) GetObjectTasksOk() (*[]KalturaObjectTask, bool)`

GetObjectTasksOk returns a tuple with the ObjectTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTasks

`func (o *KalturaScheduledTaskProfile) SetObjectTasks(v []KalturaObjectTask)`

SetObjectTasks sets ObjectTasks field to given value.

### HasObjectTasks

`func (o *KalturaScheduledTaskProfile) HasObjectTasks() bool`

HasObjectTasks returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaScheduledTaskProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaScheduledTaskProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaScheduledTaskProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaScheduledTaskProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaScheduledTaskProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaScheduledTaskProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaScheduledTaskProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaScheduledTaskProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaScheduledTaskProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaScheduledTaskProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaScheduledTaskProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaScheduledTaskProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaScheduledTaskProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaScheduledTaskProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaScheduledTaskProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaScheduledTaskProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



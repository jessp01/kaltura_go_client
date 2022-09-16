# KalturaScheduleResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto-generated unique identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Defines a short name | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **int32** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaScheduleResourceStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** | Defines a unique system-name | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaScheduleResource

`func NewKalturaScheduleResource() *KalturaScheduleResource`

NewKalturaScheduleResource instantiates a new KalturaScheduleResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaScheduleResourceWithDefaults

`func NewKalturaScheduleResourceWithDefaults() *KalturaScheduleResource`

NewKalturaScheduleResourceWithDefaults instantiates a new KalturaScheduleResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaScheduleResource) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaScheduleResource) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaScheduleResource) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaScheduleResource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaScheduleResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaScheduleResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaScheduleResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaScheduleResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaScheduleResource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaScheduleResource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaScheduleResource) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaScheduleResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaScheduleResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaScheduleResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaScheduleResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaScheduleResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaScheduleResource) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaScheduleResource) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaScheduleResource) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaScheduleResource) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetParentId

`func (o *KalturaScheduleResource) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *KalturaScheduleResource) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *KalturaScheduleResource) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *KalturaScheduleResource) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaScheduleResource) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaScheduleResource) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaScheduleResource) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaScheduleResource) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaScheduleResource) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaScheduleResource) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaScheduleResource) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaScheduleResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaScheduleResource) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaScheduleResource) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaScheduleResource) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaScheduleResource) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaScheduleResource) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaScheduleResource) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaScheduleResource) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaScheduleResource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaScheduleResource) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaScheduleResource) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaScheduleResource) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaScheduleResource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



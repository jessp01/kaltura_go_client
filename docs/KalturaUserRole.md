# KalturaUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PermissionNames** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaUserRoleStatus&#x60; | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaUserRole

`func NewKalturaUserRole() *KalturaUserRole`

NewKalturaUserRole instantiates a new KalturaUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaUserRoleWithDefaults

`func NewKalturaUserRoleWithDefaults() *KalturaUserRole`

NewKalturaUserRoleWithDefaults instantiates a new KalturaUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaUserRole) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaUserRole) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaUserRole) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaUserRole) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaUserRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaUserRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaUserRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaUserRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaUserRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaUserRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaUserRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaUserRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaUserRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaUserRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaUserRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaUserRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaUserRole) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaUserRole) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaUserRole) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaUserRole) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPermissionNames

`func (o *KalturaUserRole) GetPermissionNames() string`

GetPermissionNames returns the PermissionNames field if non-nil, zero value otherwise.

### GetPermissionNamesOk

`func (o *KalturaUserRole) GetPermissionNamesOk() (*string, bool)`

GetPermissionNamesOk returns a tuple with the PermissionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionNames

`func (o *KalturaUserRole) SetPermissionNames(v string)`

SetPermissionNames sets PermissionNames field to given value.

### HasPermissionNames

`func (o *KalturaUserRole) HasPermissionNames() bool`

HasPermissionNames returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaUserRole) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaUserRole) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaUserRole) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaUserRole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaUserRole) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaUserRole) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaUserRole) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaUserRole) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaUserRole) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaUserRole) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaUserRole) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaUserRole) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaUserRole) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaUserRole) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaUserRole) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaUserRole) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



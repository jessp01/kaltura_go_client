# KalturaPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DependsOnPermissionNames** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerGroup** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PermissionItemsIds** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaPermissionStatus&#x60; | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaPermissionType&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaPermission

`func NewKalturaPermission() *KalturaPermission`

NewKalturaPermission instantiates a new KalturaPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaPermissionWithDefaults

`func NewKalturaPermissionWithDefaults() *KalturaPermission`

NewKalturaPermissionWithDefaults instantiates a new KalturaPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaPermission) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaPermission) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaPermission) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaPermission) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDependsOnPermissionNames

`func (o *KalturaPermission) GetDependsOnPermissionNames() string`

GetDependsOnPermissionNames returns the DependsOnPermissionNames field if non-nil, zero value otherwise.

### GetDependsOnPermissionNamesOk

`func (o *KalturaPermission) GetDependsOnPermissionNamesOk() (*string, bool)`

GetDependsOnPermissionNamesOk returns a tuple with the DependsOnPermissionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnPermissionNames

`func (o *KalturaPermission) SetDependsOnPermissionNames(v string)`

SetDependsOnPermissionNames sets DependsOnPermissionNames field to given value.

### HasDependsOnPermissionNames

`func (o *KalturaPermission) HasDependsOnPermissionNames() bool`

HasDependsOnPermissionNames returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFriendlyName

`func (o *KalturaPermission) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *KalturaPermission) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *KalturaPermission) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *KalturaPermission) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *KalturaPermission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaPermission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaPermission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaPermission) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerGroup

`func (o *KalturaPermission) GetPartnerGroup() string`

GetPartnerGroup returns the PartnerGroup field if non-nil, zero value otherwise.

### GetPartnerGroupOk

`func (o *KalturaPermission) GetPartnerGroupOk() (*string, bool)`

GetPartnerGroupOk returns a tuple with the PartnerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerGroup

`func (o *KalturaPermission) SetPartnerGroup(v string)`

SetPartnerGroup sets PartnerGroup field to given value.

### HasPartnerGroup

`func (o *KalturaPermission) HasPartnerGroup() bool`

HasPartnerGroup returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaPermission) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaPermission) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaPermission) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaPermission) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPermissionItemsIds

`func (o *KalturaPermission) GetPermissionItemsIds() string`

GetPermissionItemsIds returns the PermissionItemsIds field if non-nil, zero value otherwise.

### GetPermissionItemsIdsOk

`func (o *KalturaPermission) GetPermissionItemsIdsOk() (*string, bool)`

GetPermissionItemsIdsOk returns a tuple with the PermissionItemsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionItemsIds

`func (o *KalturaPermission) SetPermissionItemsIds(v string)`

SetPermissionItemsIds sets PermissionItemsIds field to given value.

### HasPermissionItemsIds

`func (o *KalturaPermission) HasPermissionItemsIds() bool`

HasPermissionItemsIds returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaPermission) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaPermission) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaPermission) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaPermission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaPermission) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaPermission) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaPermission) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaPermission) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *KalturaPermission) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaPermission) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaPermission) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaPermission) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaPermission) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaPermission) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaPermission) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaPermission) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KalturaCategoryUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryFullIds** | Pointer to **string** | &#x60;readOnly&#x60;  The full ids of the Category | [optional] [readonly] 
**CategoryId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  CategoryUser creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60;  Partner id | [optional] [readonly] 
**PermissionLevel** | Pointer to **int32** | Enum Type: &#x60;KalturaCategoryUserPermissionLevel&#x60;  Permission level | [optional] 
**PermissionNames** | Pointer to **string** | Set of category-related permissions for the current category user. | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaCategoryUserStatus&#x60;  Status | [optional] [readonly] 
**UpdateMethod** | Pointer to **int32** | Enum Type: &#x60;KalturaUpdateMethodType&#x60;  Update method can be either manual or automatic to distinguish between manual operations (for example in KMC) on automatic - using bulk upload | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  CategoryUser update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;insertOnly&#x60;  User id | [optional] 

## Methods

### NewKalturaCategoryUser

`func NewKalturaCategoryUser() *KalturaCategoryUser`

NewKalturaCategoryUser instantiates a new KalturaCategoryUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaCategoryUserWithDefaults

`func NewKalturaCategoryUserWithDefaults() *KalturaCategoryUser`

NewKalturaCategoryUserWithDefaults instantiates a new KalturaCategoryUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryFullIds

`func (o *KalturaCategoryUser) GetCategoryFullIds() string`

GetCategoryFullIds returns the CategoryFullIds field if non-nil, zero value otherwise.

### GetCategoryFullIdsOk

`func (o *KalturaCategoryUser) GetCategoryFullIdsOk() (*string, bool)`

GetCategoryFullIdsOk returns a tuple with the CategoryFullIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFullIds

`func (o *KalturaCategoryUser) SetCategoryFullIds(v string)`

SetCategoryFullIds sets CategoryFullIds field to given value.

### HasCategoryFullIds

`func (o *KalturaCategoryUser) HasCategoryFullIds() bool`

HasCategoryFullIds returns a boolean if a field has been set.

### GetCategoryId

`func (o *KalturaCategoryUser) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *KalturaCategoryUser) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *KalturaCategoryUser) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *KalturaCategoryUser) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaCategoryUser) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaCategoryUser) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaCategoryUser) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaCategoryUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaCategoryUser) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaCategoryUser) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaCategoryUser) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaCategoryUser) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPermissionLevel

`func (o *KalturaCategoryUser) GetPermissionLevel() int32`

GetPermissionLevel returns the PermissionLevel field if non-nil, zero value otherwise.

### GetPermissionLevelOk

`func (o *KalturaCategoryUser) GetPermissionLevelOk() (*int32, bool)`

GetPermissionLevelOk returns a tuple with the PermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionLevel

`func (o *KalturaCategoryUser) SetPermissionLevel(v int32)`

SetPermissionLevel sets PermissionLevel field to given value.

### HasPermissionLevel

`func (o *KalturaCategoryUser) HasPermissionLevel() bool`

HasPermissionLevel returns a boolean if a field has been set.

### GetPermissionNames

`func (o *KalturaCategoryUser) GetPermissionNames() string`

GetPermissionNames returns the PermissionNames field if non-nil, zero value otherwise.

### GetPermissionNamesOk

`func (o *KalturaCategoryUser) GetPermissionNamesOk() (*string, bool)`

GetPermissionNamesOk returns a tuple with the PermissionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionNames

`func (o *KalturaCategoryUser) SetPermissionNames(v string)`

SetPermissionNames sets PermissionNames field to given value.

### HasPermissionNames

`func (o *KalturaCategoryUser) HasPermissionNames() bool`

HasPermissionNames returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaCategoryUser) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaCategoryUser) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaCategoryUser) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaCategoryUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateMethod

`func (o *KalturaCategoryUser) GetUpdateMethod() int32`

GetUpdateMethod returns the UpdateMethod field if non-nil, zero value otherwise.

### GetUpdateMethodOk

`func (o *KalturaCategoryUser) GetUpdateMethodOk() (*int32, bool)`

GetUpdateMethodOk returns a tuple with the UpdateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMethod

`func (o *KalturaCategoryUser) SetUpdateMethod(v int32)`

SetUpdateMethod sets UpdateMethod field to given value.

### HasUpdateMethod

`func (o *KalturaCategoryUser) HasUpdateMethod() bool`

HasUpdateMethod returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaCategoryUser) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaCategoryUser) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaCategoryUser) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaCategoryUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaCategoryUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaCategoryUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaCategoryUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaCategoryUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



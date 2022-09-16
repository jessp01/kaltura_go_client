# KalturaGroupUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CreationMode** | Pointer to **int32** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaGroupUserCreationMode&#x60; | [optional] 
**GroupId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaGroupUserStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**UserRole** | Pointer to **int32** | Enum Type: &#x60;KalturaGroupUserRole&#x60; | [optional] 

## Methods

### NewKalturaGroupUser

`func NewKalturaGroupUser() *KalturaGroupUser`

NewKalturaGroupUser instantiates a new KalturaGroupUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaGroupUserWithDefaults

`func NewKalturaGroupUserWithDefaults() *KalturaGroupUser`

NewKalturaGroupUserWithDefaults instantiates a new KalturaGroupUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaGroupUser) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaGroupUser) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaGroupUser) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaGroupUser) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreationMode

`func (o *KalturaGroupUser) GetCreationMode() int32`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *KalturaGroupUser) GetCreationModeOk() (*int32, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *KalturaGroupUser) SetCreationMode(v int32)`

SetCreationMode sets CreationMode field to given value.

### HasCreationMode

`func (o *KalturaGroupUser) HasCreationMode() bool`

HasCreationMode returns a boolean if a field has been set.

### GetGroupId

`func (o *KalturaGroupUser) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *KalturaGroupUser) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *KalturaGroupUser) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *KalturaGroupUser) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *KalturaGroupUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaGroupUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaGroupUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaGroupUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaGroupUser) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaGroupUser) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaGroupUser) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaGroupUser) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaGroupUser) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaGroupUser) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaGroupUser) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaGroupUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaGroupUser) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaGroupUser) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaGroupUser) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaGroupUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaGroupUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaGroupUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaGroupUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaGroupUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserRole

`func (o *KalturaGroupUser) GetUserRole() int32`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *KalturaGroupUser) GetUserRoleOk() (*int32, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *KalturaGroupUser) SetUserRole(v int32)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *KalturaGroupUser) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GroupUserSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateNewGroups** | Pointer to **bool** |  | [optional] [default to true]
**GroupIds** | Pointer to **string** |  | [optional] 
**RemoveFromExistingGroups** | Pointer to **bool** |  | [optional] [default to true]
**UserId** | **string** |  | 

## Methods

### NewGroupUserSyncRequest

`func NewGroupUserSyncRequest(userId string, ) *GroupUserSyncRequest`

NewGroupUserSyncRequest instantiates a new GroupUserSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserSyncRequestWithDefaults

`func NewGroupUserSyncRequestWithDefaults() *GroupUserSyncRequest`

NewGroupUserSyncRequestWithDefaults instantiates a new GroupUserSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateNewGroups

`func (o *GroupUserSyncRequest) GetCreateNewGroups() bool`

GetCreateNewGroups returns the CreateNewGroups field if non-nil, zero value otherwise.

### GetCreateNewGroupsOk

`func (o *GroupUserSyncRequest) GetCreateNewGroupsOk() (*bool, bool)`

GetCreateNewGroupsOk returns a tuple with the CreateNewGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateNewGroups

`func (o *GroupUserSyncRequest) SetCreateNewGroups(v bool)`

SetCreateNewGroups sets CreateNewGroups field to given value.

### HasCreateNewGroups

`func (o *GroupUserSyncRequest) HasCreateNewGroups() bool`

HasCreateNewGroups returns a boolean if a field has been set.

### GetGroupIds

`func (o *GroupUserSyncRequest) GetGroupIds() string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *GroupUserSyncRequest) GetGroupIdsOk() (*string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *GroupUserSyncRequest) SetGroupIds(v string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *GroupUserSyncRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetRemoveFromExistingGroups

`func (o *GroupUserSyncRequest) GetRemoveFromExistingGroups() bool`

GetRemoveFromExistingGroups returns the RemoveFromExistingGroups field if non-nil, zero value otherwise.

### GetRemoveFromExistingGroupsOk

`func (o *GroupUserSyncRequest) GetRemoveFromExistingGroupsOk() (*bool, bool)`

GetRemoveFromExistingGroupsOk returns a tuple with the RemoveFromExistingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveFromExistingGroups

`func (o *GroupUserSyncRequest) SetRemoveFromExistingGroups(v bool)`

SetRemoveFromExistingGroups sets RemoveFromExistingGroups field to given value.

### HasRemoveFromExistingGroups

`func (o *GroupUserSyncRequest) HasRemoveFromExistingGroups() bool`

HasRemoveFromExistingGroups returns a boolean if a field has been set.

### GetUserId

`func (o *GroupUserSyncRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupUserSyncRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupUserSyncRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



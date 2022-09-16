# GroupUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupUser** | [**KalturaGroupUser**](KalturaGroupUser.md) |  | 
**GroupUserId** | **string** |  | 

## Methods

### NewGroupUserUpdateRequest

`func NewGroupUserUpdateRequest(groupUser KalturaGroupUser, groupUserId string, ) *GroupUserUpdateRequest`

NewGroupUserUpdateRequest instantiates a new GroupUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserUpdateRequestWithDefaults

`func NewGroupUserUpdateRequestWithDefaults() *GroupUserUpdateRequest`

NewGroupUserUpdateRequestWithDefaults instantiates a new GroupUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupUser

`func (o *GroupUserUpdateRequest) GetGroupUser() KalturaGroupUser`

GetGroupUser returns the GroupUser field if non-nil, zero value otherwise.

### GetGroupUserOk

`func (o *GroupUserUpdateRequest) GetGroupUserOk() (*KalturaGroupUser, bool)`

GetGroupUserOk returns a tuple with the GroupUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUser

`func (o *GroupUserUpdateRequest) SetGroupUser(v KalturaGroupUser)`

SetGroupUser sets GroupUser field to given value.


### GetGroupUserId

`func (o *GroupUserUpdateRequest) GetGroupUserId() string`

GetGroupUserId returns the GroupUserId field if non-nil, zero value otherwise.

### GetGroupUserIdOk

`func (o *GroupUserUpdateRequest) GetGroupUserIdOk() (*string, bool)`

GetGroupUserIdOk returns a tuple with the GroupUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupUserId

`func (o *GroupUserUpdateRequest) SetGroupUserId(v string)`

SetGroupUserId sets GroupUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



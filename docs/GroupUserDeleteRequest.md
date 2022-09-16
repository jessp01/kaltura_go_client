# GroupUserDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewGroupUserDeleteRequest

`func NewGroupUserDeleteRequest(groupId string, userId string, ) *GroupUserDeleteRequest`

NewGroupUserDeleteRequest instantiates a new GroupUserDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUserDeleteRequestWithDefaults

`func NewGroupUserDeleteRequestWithDefaults() *GroupUserDeleteRequest`

NewGroupUserDeleteRequestWithDefaults instantiates a new GroupUserDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupUserDeleteRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupUserDeleteRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupUserDeleteRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetUserId

`func (o *GroupUserDeleteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GroupUserDeleteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GroupUserDeleteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



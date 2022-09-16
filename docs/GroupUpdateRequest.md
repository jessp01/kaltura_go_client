# GroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**KalturaGroup**](KalturaGroup.md) |  | 
**GroupId** | **string** |  | 

## Methods

### NewGroupUpdateRequest

`func NewGroupUpdateRequest(group KalturaGroup, groupId string, ) *GroupUpdateRequest`

NewGroupUpdateRequest instantiates a new GroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateRequestWithDefaults

`func NewGroupUpdateRequestWithDefaults() *GroupUpdateRequest`

NewGroupUpdateRequestWithDefaults instantiates a new GroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GroupUpdateRequest) GetGroup() KalturaGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GroupUpdateRequest) GetGroupOk() (*KalturaGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GroupUpdateRequest) SetGroup(v KalturaGroup)`

SetGroup sets Group field to given value.


### GetGroupId

`func (o *GroupUpdateRequest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupUpdateRequest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupUpdateRequest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



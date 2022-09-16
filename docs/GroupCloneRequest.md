# GroupCloneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewGroupId** | **string** |  | 
**NewGroupName** | Pointer to **string** |  | [optional] 
**OriginalGroupId** | **string** |  | 

## Methods

### NewGroupCloneRequest

`func NewGroupCloneRequest(newGroupId string, originalGroupId string, ) *GroupCloneRequest`

NewGroupCloneRequest instantiates a new GroupCloneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCloneRequestWithDefaults

`func NewGroupCloneRequestWithDefaults() *GroupCloneRequest`

NewGroupCloneRequestWithDefaults instantiates a new GroupCloneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewGroupId

`func (o *GroupCloneRequest) GetNewGroupId() string`

GetNewGroupId returns the NewGroupId field if non-nil, zero value otherwise.

### GetNewGroupIdOk

`func (o *GroupCloneRequest) GetNewGroupIdOk() (*string, bool)`

GetNewGroupIdOk returns a tuple with the NewGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewGroupId

`func (o *GroupCloneRequest) SetNewGroupId(v string)`

SetNewGroupId sets NewGroupId field to given value.


### GetNewGroupName

`func (o *GroupCloneRequest) GetNewGroupName() string`

GetNewGroupName returns the NewGroupName field if non-nil, zero value otherwise.

### GetNewGroupNameOk

`func (o *GroupCloneRequest) GetNewGroupNameOk() (*string, bool)`

GetNewGroupNameOk returns a tuple with the NewGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewGroupName

`func (o *GroupCloneRequest) SetNewGroupName(v string)`

SetNewGroupName sets NewGroupName field to given value.

### HasNewGroupName

`func (o *GroupCloneRequest) HasNewGroupName() bool`

HasNewGroupName returns a boolean if a field has been set.

### GetOriginalGroupId

`func (o *GroupCloneRequest) GetOriginalGroupId() string`

GetOriginalGroupId returns the OriginalGroupId field if non-nil, zero value otherwise.

### GetOriginalGroupIdOk

`func (o *GroupCloneRequest) GetOriginalGroupIdOk() (*string, bool)`

GetOriginalGroupIdOk returns a tuple with the OriginalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalGroupId

`func (o *GroupCloneRequest) SetOriginalGroupId(v string)`

SetOriginalGroupId sets OriginalGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



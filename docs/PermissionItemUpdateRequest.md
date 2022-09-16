# PermissionItemUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionItem** | [**KalturaPermissionItem**](KalturaPermissionItem.md) |  | 
**PermissionItemId** | **int32** |  | 

## Methods

### NewPermissionItemUpdateRequest

`func NewPermissionItemUpdateRequest(permissionItem KalturaPermissionItem, permissionItemId int32, ) *PermissionItemUpdateRequest`

NewPermissionItemUpdateRequest instantiates a new PermissionItemUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionItemUpdateRequestWithDefaults

`func NewPermissionItemUpdateRequestWithDefaults() *PermissionItemUpdateRequest`

NewPermissionItemUpdateRequestWithDefaults instantiates a new PermissionItemUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionItem

`func (o *PermissionItemUpdateRequest) GetPermissionItem() KalturaPermissionItem`

GetPermissionItem returns the PermissionItem field if non-nil, zero value otherwise.

### GetPermissionItemOk

`func (o *PermissionItemUpdateRequest) GetPermissionItemOk() (*KalturaPermissionItem, bool)`

GetPermissionItemOk returns a tuple with the PermissionItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionItem

`func (o *PermissionItemUpdateRequest) SetPermissionItem(v KalturaPermissionItem)`

SetPermissionItem sets PermissionItem field to given value.


### GetPermissionItemId

`func (o *PermissionItemUpdateRequest) GetPermissionItemId() int32`

GetPermissionItemId returns the PermissionItemId field if non-nil, zero value otherwise.

### GetPermissionItemIdOk

`func (o *PermissionItemUpdateRequest) GetPermissionItemIdOk() (*int32, bool)`

GetPermissionItemIdOk returns a tuple with the PermissionItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionItemId

`func (o *PermissionItemUpdateRequest) SetPermissionItemId(v int32)`

SetPermissionItemId sets PermissionItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



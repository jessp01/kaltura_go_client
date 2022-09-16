# PermissionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | [**KalturaPermission**](KalturaPermission.md) |  | 
**PermissionName** | **string** |  | 

## Methods

### NewPermissionUpdateRequest

`func NewPermissionUpdateRequest(permission KalturaPermission, permissionName string, ) *PermissionUpdateRequest`

NewPermissionUpdateRequest instantiates a new PermissionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionUpdateRequestWithDefaults

`func NewPermissionUpdateRequestWithDefaults() *PermissionUpdateRequest`

NewPermissionUpdateRequestWithDefaults instantiates a new PermissionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionUpdateRequest) GetPermission() KalturaPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionUpdateRequest) GetPermissionOk() (*KalturaPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionUpdateRequest) SetPermission(v KalturaPermission)`

SetPermission sets Permission field to given value.


### GetPermissionName

`func (o *PermissionUpdateRequest) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *PermissionUpdateRequest) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *PermissionUpdateRequest) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PermissionAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | [**KalturaPermission**](KalturaPermission.md) |  | 

## Methods

### NewPermissionAddRequest

`func NewPermissionAddRequest(permission KalturaPermission, ) *PermissionAddRequest`

NewPermissionAddRequest instantiates a new PermissionAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionAddRequestWithDefaults

`func NewPermissionAddRequestWithDefaults() *PermissionAddRequest`

NewPermissionAddRequestWithDefaults instantiates a new PermissionAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *PermissionAddRequest) GetPermission() KalturaPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *PermissionAddRequest) GetPermissionOk() (*KalturaPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *PermissionAddRequest) SetPermission(v KalturaPermission)`

SetPermission sets Permission field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



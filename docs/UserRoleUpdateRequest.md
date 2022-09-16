# UserRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRole** | [**KalturaUserRole**](KalturaUserRole.md) |  | 
**UserRoleId** | **int32** |  | 

## Methods

### NewUserRoleUpdateRequest

`func NewUserRoleUpdateRequest(userRole KalturaUserRole, userRoleId int32, ) *UserRoleUpdateRequest`

NewUserRoleUpdateRequest instantiates a new UserRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRoleUpdateRequestWithDefaults

`func NewUserRoleUpdateRequestWithDefaults() *UserRoleUpdateRequest`

NewUserRoleUpdateRequestWithDefaults instantiates a new UserRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRole

`func (o *UserRoleUpdateRequest) GetUserRole() KalturaUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *UserRoleUpdateRequest) GetUserRoleOk() (*KalturaUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *UserRoleUpdateRequest) SetUserRole(v KalturaUserRole)`

SetUserRole sets UserRole field to given value.


### GetUserRoleId

`func (o *UserRoleUpdateRequest) GetUserRoleId() int32`

GetUserRoleId returns the UserRoleId field if non-nil, zero value otherwise.

### GetUserRoleIdOk

`func (o *UserRoleUpdateRequest) GetUserRoleIdOk() (*int32, bool)`

GetUserRoleIdOk returns a tuple with the UserRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRoleId

`func (o *UserRoleUpdateRequest) SetUserRoleId(v int32)`

SetUserRoleId sets UserRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



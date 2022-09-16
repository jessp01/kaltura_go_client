# AdminUserSetInitialPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashKey** | **string** |  | 
**NewPassword** | **string** |  | 

## Methods

### NewAdminUserSetInitialPasswordRequest

`func NewAdminUserSetInitialPasswordRequest(hashKey string, newPassword string, ) *AdminUserSetInitialPasswordRequest`

NewAdminUserSetInitialPasswordRequest instantiates a new AdminUserSetInitialPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserSetInitialPasswordRequestWithDefaults

`func NewAdminUserSetInitialPasswordRequestWithDefaults() *AdminUserSetInitialPasswordRequest`

NewAdminUserSetInitialPasswordRequestWithDefaults instantiates a new AdminUserSetInitialPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashKey

`func (o *AdminUserSetInitialPasswordRequest) GetHashKey() string`

GetHashKey returns the HashKey field if non-nil, zero value otherwise.

### GetHashKeyOk

`func (o *AdminUserSetInitialPasswordRequest) GetHashKeyOk() (*string, bool)`

GetHashKeyOk returns a tuple with the HashKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashKey

`func (o *AdminUserSetInitialPasswordRequest) SetHashKey(v string)`

SetHashKey sets HashKey field to given value.


### GetNewPassword

`func (o *AdminUserSetInitialPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *AdminUserSetInitialPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *AdminUserSetInitialPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



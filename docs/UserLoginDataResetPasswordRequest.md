# UserLoginDataResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginDataId** | **string** |  | 
**NewPassword** | **string** |  | 

## Methods

### NewUserLoginDataResetPasswordRequest

`func NewUserLoginDataResetPasswordRequest(loginDataId string, newPassword string, ) *UserLoginDataResetPasswordRequest`

NewUserLoginDataResetPasswordRequest instantiates a new UserLoginDataResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginDataResetPasswordRequestWithDefaults

`func NewUserLoginDataResetPasswordRequestWithDefaults() *UserLoginDataResetPasswordRequest`

NewUserLoginDataResetPasswordRequestWithDefaults instantiates a new UserLoginDataResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginDataId

`func (o *UserLoginDataResetPasswordRequest) GetLoginDataId() string`

GetLoginDataId returns the LoginDataId field if non-nil, zero value otherwise.

### GetLoginDataIdOk

`func (o *UserLoginDataResetPasswordRequest) GetLoginDataIdOk() (*string, bool)`

GetLoginDataIdOk returns a tuple with the LoginDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDataId

`func (o *UserLoginDataResetPasswordRequest) SetLoginDataId(v string)`

SetLoginDataId sets LoginDataId field to given value.


### GetNewPassword

`func (o *UserLoginDataResetPasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserLoginDataResetPasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserLoginDataResetPasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



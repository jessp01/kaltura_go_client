# AdminUserUpdatePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**NewEmail** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 
**Otp** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewAdminUserUpdatePasswordRequest

`func NewAdminUserUpdatePasswordRequest(email string, password string, ) *AdminUserUpdatePasswordRequest`

NewAdminUserUpdatePasswordRequest instantiates a new AdminUserUpdatePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserUpdatePasswordRequestWithDefaults

`func NewAdminUserUpdatePasswordRequestWithDefaults() *AdminUserUpdatePasswordRequest`

NewAdminUserUpdatePasswordRequestWithDefaults instantiates a new AdminUserUpdatePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminUserUpdatePasswordRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserUpdatePasswordRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserUpdatePasswordRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewEmail

`func (o *AdminUserUpdatePasswordRequest) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *AdminUserUpdatePasswordRequest) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *AdminUserUpdatePasswordRequest) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.

### HasNewEmail

`func (o *AdminUserUpdatePasswordRequest) HasNewEmail() bool`

HasNewEmail returns a boolean if a field has been set.

### GetNewPassword

`func (o *AdminUserUpdatePasswordRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *AdminUserUpdatePasswordRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *AdminUserUpdatePasswordRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *AdminUserUpdatePasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetOtp

`func (o *AdminUserUpdatePasswordRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *AdminUserUpdatePasswordRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *AdminUserUpdatePasswordRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *AdminUserUpdatePasswordRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPassword

`func (o *AdminUserUpdatePasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminUserUpdatePasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminUserUpdatePasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



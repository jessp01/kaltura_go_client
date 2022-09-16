# UserUpdateLoginDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewFirstName** | Pointer to **string** |  | [optional] 
**NewLastName** | Pointer to **string** |  | [optional] 
**NewLoginId** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 
**OldLoginId** | **string** |  | 
**Otp** | Pointer to **string** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewUserUpdateLoginDataRequest

`func NewUserUpdateLoginDataRequest(oldLoginId string, password string, ) *UserUpdateLoginDataRequest`

NewUserUpdateLoginDataRequest instantiates a new UserUpdateLoginDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateLoginDataRequestWithDefaults

`func NewUserUpdateLoginDataRequestWithDefaults() *UserUpdateLoginDataRequest`

NewUserUpdateLoginDataRequestWithDefaults instantiates a new UserUpdateLoginDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewFirstName

`func (o *UserUpdateLoginDataRequest) GetNewFirstName() string`

GetNewFirstName returns the NewFirstName field if non-nil, zero value otherwise.

### GetNewFirstNameOk

`func (o *UserUpdateLoginDataRequest) GetNewFirstNameOk() (*string, bool)`

GetNewFirstNameOk returns a tuple with the NewFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewFirstName

`func (o *UserUpdateLoginDataRequest) SetNewFirstName(v string)`

SetNewFirstName sets NewFirstName field to given value.

### HasNewFirstName

`func (o *UserUpdateLoginDataRequest) HasNewFirstName() bool`

HasNewFirstName returns a boolean if a field has been set.

### GetNewLastName

`func (o *UserUpdateLoginDataRequest) GetNewLastName() string`

GetNewLastName returns the NewLastName field if non-nil, zero value otherwise.

### GetNewLastNameOk

`func (o *UserUpdateLoginDataRequest) GetNewLastNameOk() (*string, bool)`

GetNewLastNameOk returns a tuple with the NewLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLastName

`func (o *UserUpdateLoginDataRequest) SetNewLastName(v string)`

SetNewLastName sets NewLastName field to given value.

### HasNewLastName

`func (o *UserUpdateLoginDataRequest) HasNewLastName() bool`

HasNewLastName returns a boolean if a field has been set.

### GetNewLoginId

`func (o *UserUpdateLoginDataRequest) GetNewLoginId() string`

GetNewLoginId returns the NewLoginId field if non-nil, zero value otherwise.

### GetNewLoginIdOk

`func (o *UserUpdateLoginDataRequest) GetNewLoginIdOk() (*string, bool)`

GetNewLoginIdOk returns a tuple with the NewLoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLoginId

`func (o *UserUpdateLoginDataRequest) SetNewLoginId(v string)`

SetNewLoginId sets NewLoginId field to given value.

### HasNewLoginId

`func (o *UserUpdateLoginDataRequest) HasNewLoginId() bool`

HasNewLoginId returns a boolean if a field has been set.

### GetNewPassword

`func (o *UserUpdateLoginDataRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UserUpdateLoginDataRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UserUpdateLoginDataRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UserUpdateLoginDataRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetOldLoginId

`func (o *UserUpdateLoginDataRequest) GetOldLoginId() string`

GetOldLoginId returns the OldLoginId field if non-nil, zero value otherwise.

### GetOldLoginIdOk

`func (o *UserUpdateLoginDataRequest) GetOldLoginIdOk() (*string, bool)`

GetOldLoginIdOk returns a tuple with the OldLoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldLoginId

`func (o *UserUpdateLoginDataRequest) SetOldLoginId(v string)`

SetOldLoginId sets OldLoginId field to given value.


### GetOtp

`func (o *UserUpdateLoginDataRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *UserUpdateLoginDataRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *UserUpdateLoginDataRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *UserUpdateLoginDataRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPassword

`func (o *UserUpdateLoginDataRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserUpdateLoginDataRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserUpdateLoginDataRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



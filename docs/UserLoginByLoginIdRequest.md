# UserLoginByLoginIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**LoginId** | **string** |  | 
**Otp** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 
**Privileges** | Pointer to **string** |  | [optional] [default to "*"]

## Methods

### NewUserLoginByLoginIdRequest

`func NewUserLoginByLoginIdRequest(loginId string, password string, ) *UserLoginByLoginIdRequest`

NewUserLoginByLoginIdRequest instantiates a new UserLoginByLoginIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginByLoginIdRequestWithDefaults

`func NewUserLoginByLoginIdRequestWithDefaults() *UserLoginByLoginIdRequest`

NewUserLoginByLoginIdRequestWithDefaults instantiates a new UserLoginByLoginIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *UserLoginByLoginIdRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UserLoginByLoginIdRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UserLoginByLoginIdRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *UserLoginByLoginIdRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLoginId

`func (o *UserLoginByLoginIdRequest) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *UserLoginByLoginIdRequest) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *UserLoginByLoginIdRequest) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.


### GetOtp

`func (o *UserLoginByLoginIdRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *UserLoginByLoginIdRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *UserLoginByLoginIdRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *UserLoginByLoginIdRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetPartnerId

`func (o *UserLoginByLoginIdRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *UserLoginByLoginIdRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *UserLoginByLoginIdRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *UserLoginByLoginIdRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPassword

`func (o *UserLoginByLoginIdRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserLoginByLoginIdRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserLoginByLoginIdRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPrivileges

`func (o *UserLoginByLoginIdRequest) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UserLoginByLoginIdRequest) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UserLoginByLoginIdRequest) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UserLoginByLoginIdRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



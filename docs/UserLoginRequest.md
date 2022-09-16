# UserLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**PartnerId** | **int32** |  | 
**Password** | **string** |  | 
**Privileges** | Pointer to **string** |  | [optional] [default to "*"]
**UserId** | **string** |  | 

## Methods

### NewUserLoginRequest

`func NewUserLoginRequest(partnerId int32, password string, userId string, ) *UserLoginRequest`

NewUserLoginRequest instantiates a new UserLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginRequestWithDefaults

`func NewUserLoginRequestWithDefaults() *UserLoginRequest`

NewUserLoginRequestWithDefaults instantiates a new UserLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *UserLoginRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UserLoginRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UserLoginRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *UserLoginRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetPartnerId

`func (o *UserLoginRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *UserLoginRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *UserLoginRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.


### GetPassword

`func (o *UserLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPrivileges

`func (o *UserLoginRequest) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *UserLoginRequest) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *UserLoginRequest) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *UserLoginRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetUserId

`func (o *UserLoginRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserLoginRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserLoginRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



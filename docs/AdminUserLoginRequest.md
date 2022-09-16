# AdminUserLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**PartnerId** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 

## Methods

### NewAdminUserLoginRequest

`func NewAdminUserLoginRequest(email string, password string, ) *AdminUserLoginRequest`

NewAdminUserLoginRequest instantiates a new AdminUserLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserLoginRequestWithDefaults

`func NewAdminUserLoginRequestWithDefaults() *AdminUserLoginRequest`

NewAdminUserLoginRequestWithDefaults instantiates a new AdminUserLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AdminUserLoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserLoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserLoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPartnerId

`func (o *AdminUserLoginRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *AdminUserLoginRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *AdminUserLoginRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *AdminUserLoginRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPassword

`func (o *AdminUserLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AdminUserLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AdminUserLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SessionImpersonateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**ImpersonatedPartnerId** | **int32** |  | 
**PartnerId** | Pointer to **int32** |  | [optional] 
**Privileges** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 
**Type** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionImpersonateRequest

`func NewSessionImpersonateRequest(impersonatedPartnerId int32, secret string, ) *SessionImpersonateRequest`

NewSessionImpersonateRequest instantiates a new SessionImpersonateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionImpersonateRequestWithDefaults

`func NewSessionImpersonateRequestWithDefaults() *SessionImpersonateRequest`

NewSessionImpersonateRequestWithDefaults instantiates a new SessionImpersonateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *SessionImpersonateRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SessionImpersonateRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SessionImpersonateRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SessionImpersonateRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetImpersonatedPartnerId

`func (o *SessionImpersonateRequest) GetImpersonatedPartnerId() int32`

GetImpersonatedPartnerId returns the ImpersonatedPartnerId field if non-nil, zero value otherwise.

### GetImpersonatedPartnerIdOk

`func (o *SessionImpersonateRequest) GetImpersonatedPartnerIdOk() (*int32, bool)`

GetImpersonatedPartnerIdOk returns a tuple with the ImpersonatedPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonatedPartnerId

`func (o *SessionImpersonateRequest) SetImpersonatedPartnerId(v int32)`

SetImpersonatedPartnerId sets ImpersonatedPartnerId field to given value.


### GetPartnerId

`func (o *SessionImpersonateRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *SessionImpersonateRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *SessionImpersonateRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *SessionImpersonateRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPrivileges

`func (o *SessionImpersonateRequest) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *SessionImpersonateRequest) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *SessionImpersonateRequest) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *SessionImpersonateRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSecret

`func (o *SessionImpersonateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SessionImpersonateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SessionImpersonateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetType

`func (o *SessionImpersonateRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionImpersonateRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionImpersonateRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SessionImpersonateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *SessionImpersonateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionImpersonateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionImpersonateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionImpersonateRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



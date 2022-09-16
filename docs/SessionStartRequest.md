# SessionStartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**Privileges** | Pointer to **string** |  | [optional] 
**Secret** | **string** |  | 
**Type** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionStartRequest

`func NewSessionStartRequest(secret string, ) *SessionStartRequest`

NewSessionStartRequest instantiates a new SessionStartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionStartRequestWithDefaults

`func NewSessionStartRequestWithDefaults() *SessionStartRequest`

NewSessionStartRequestWithDefaults instantiates a new SessionStartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *SessionStartRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SessionStartRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SessionStartRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SessionStartRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetPartnerId

`func (o *SessionStartRequest) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *SessionStartRequest) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *SessionStartRequest) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *SessionStartRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPrivileges

`func (o *SessionStartRequest) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *SessionStartRequest) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *SessionStartRequest) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *SessionStartRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSecret

`func (o *SessionStartRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SessionStartRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SessionStartRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetType

`func (o *SessionStartRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionStartRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionStartRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SessionStartRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *SessionStartRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionStartRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionStartRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionStartRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



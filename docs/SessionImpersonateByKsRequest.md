# SessionImpersonateByKsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**Privileges** | Pointer to **string** |  | [optional] 
**Session** | **string** |  | 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewSessionImpersonateByKsRequest

`func NewSessionImpersonateByKsRequest(session string, ) *SessionImpersonateByKsRequest`

NewSessionImpersonateByKsRequest instantiates a new SessionImpersonateByKsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionImpersonateByKsRequestWithDefaults

`func NewSessionImpersonateByKsRequestWithDefaults() *SessionImpersonateByKsRequest`

NewSessionImpersonateByKsRequestWithDefaults instantiates a new SessionImpersonateByKsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *SessionImpersonateByKsRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SessionImpersonateByKsRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SessionImpersonateByKsRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SessionImpersonateByKsRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetPrivileges

`func (o *SessionImpersonateByKsRequest) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *SessionImpersonateByKsRequest) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *SessionImpersonateByKsRequest) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *SessionImpersonateByKsRequest) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSession

`func (o *SessionImpersonateByKsRequest) GetSession() string`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *SessionImpersonateByKsRequest) GetSessionOk() (*string, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *SessionImpersonateByKsRequest) SetSession(v string)`

SetSession sets Session field to given value.


### GetType

`func (o *SessionImpersonateByKsRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SessionImpersonateByKsRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SessionImpersonateByKsRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SessionImpersonateByKsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



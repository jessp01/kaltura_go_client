# AppTokenStartSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**SessionPrivileges** | Pointer to **string** |  | [optional] 
**TokenHash** | **string** |  | 
**Type** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewAppTokenStartSessionRequest

`func NewAppTokenStartSessionRequest(id string, tokenHash string, ) *AppTokenStartSessionRequest`

NewAppTokenStartSessionRequest instantiates a new AppTokenStartSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppTokenStartSessionRequestWithDefaults

`func NewAppTokenStartSessionRequestWithDefaults() *AppTokenStartSessionRequest`

NewAppTokenStartSessionRequestWithDefaults instantiates a new AppTokenStartSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *AppTokenStartSessionRequest) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AppTokenStartSessionRequest) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AppTokenStartSessionRequest) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AppTokenStartSessionRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetId

`func (o *AppTokenStartSessionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppTokenStartSessionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppTokenStartSessionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSessionPrivileges

`func (o *AppTokenStartSessionRequest) GetSessionPrivileges() string`

GetSessionPrivileges returns the SessionPrivileges field if non-nil, zero value otherwise.

### GetSessionPrivilegesOk

`func (o *AppTokenStartSessionRequest) GetSessionPrivilegesOk() (*string, bool)`

GetSessionPrivilegesOk returns a tuple with the SessionPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPrivileges

`func (o *AppTokenStartSessionRequest) SetSessionPrivileges(v string)`

SetSessionPrivileges sets SessionPrivileges field to given value.

### HasSessionPrivileges

`func (o *AppTokenStartSessionRequest) HasSessionPrivileges() bool`

HasSessionPrivileges returns a boolean if a field has been set.

### GetTokenHash

`func (o *AppTokenStartSessionRequest) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *AppTokenStartSessionRequest) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *AppTokenStartSessionRequest) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.


### GetType

`func (o *AppTokenStartSessionRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppTokenStartSessionRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppTokenStartSessionRequest) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AppTokenStartSessionRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *AppTokenStartSessionRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AppTokenStartSessionRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AppTokenStartSessionRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AppTokenStartSessionRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KalturaSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Ks** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Privileges** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**SessionType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaSessionType&#x60; | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaSessionInfo

`func NewKalturaSessionInfo() *KalturaSessionInfo`

NewKalturaSessionInfo instantiates a new KalturaSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSessionInfoWithDefaults

`func NewKalturaSessionInfoWithDefaults() *KalturaSessionInfo`

NewKalturaSessionInfoWithDefaults instantiates a new KalturaSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *KalturaSessionInfo) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *KalturaSessionInfo) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *KalturaSessionInfo) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *KalturaSessionInfo) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetKs

`func (o *KalturaSessionInfo) GetKs() string`

GetKs returns the Ks field if non-nil, zero value otherwise.

### GetKsOk

`func (o *KalturaSessionInfo) GetKsOk() (*string, bool)`

GetKsOk returns a tuple with the Ks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKs

`func (o *KalturaSessionInfo) SetKs(v string)`

SetKs sets Ks field to given value.

### HasKs

`func (o *KalturaSessionInfo) HasKs() bool`

HasKs returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaSessionInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaSessionInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaSessionInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaSessionInfo) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaSessionInfo) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaSessionInfo) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaSessionInfo) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaSessionInfo) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPrivileges

`func (o *KalturaSessionInfo) GetPrivileges() string`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *KalturaSessionInfo) GetPrivilegesOk() (*string, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *KalturaSessionInfo) SetPrivileges(v string)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *KalturaSessionInfo) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetSessionType

`func (o *KalturaSessionInfo) GetSessionType() int32`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *KalturaSessionInfo) GetSessionTypeOk() (*int32, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *KalturaSessionInfo) SetSessionType(v int32)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *KalturaSessionInfo) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



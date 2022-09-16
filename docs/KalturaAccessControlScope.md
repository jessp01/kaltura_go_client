# KalturaAccessControlScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contexts** | Pointer to [**[]KalturaAccessControlContextTypeHolder**](KalturaAccessControlContextTypeHolder.md) |  | [optional] 
**Hashes** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**Ip** | Pointer to **string** | IP to be used to test geographic location conditions. | [optional] 
**Ks** | Pointer to **string** | Kaltura session to be used to test session and user conditions. | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** | URL to be used to test domain conditions. | [optional] 
**Time** | Pointer to **int32** | Unix timestamp (In seconds) to be used to test entry scheduling, keep null to use now. | [optional] 
**UserAgent** | Pointer to **string** | Browser or client application to be used to test agent conditions. | [optional] 

## Methods

### NewKalturaAccessControlScope

`func NewKalturaAccessControlScope() *KalturaAccessControlScope`

NewKalturaAccessControlScope instantiates a new KalturaAccessControlScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAccessControlScopeWithDefaults

`func NewKalturaAccessControlScopeWithDefaults() *KalturaAccessControlScope`

NewKalturaAccessControlScopeWithDefaults instantiates a new KalturaAccessControlScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContexts

`func (o *KalturaAccessControlScope) GetContexts() []KalturaAccessControlContextTypeHolder`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *KalturaAccessControlScope) GetContextsOk() (*[]KalturaAccessControlContextTypeHolder, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *KalturaAccessControlScope) SetContexts(v []KalturaAccessControlContextTypeHolder)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *KalturaAccessControlScope) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetHashes

`func (o *KalturaAccessControlScope) GetHashes() []KalturaKeyValue`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *KalturaAccessControlScope) GetHashesOk() (*[]KalturaKeyValue, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *KalturaAccessControlScope) SetHashes(v []KalturaKeyValue)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *KalturaAccessControlScope) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetIp

`func (o *KalturaAccessControlScope) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *KalturaAccessControlScope) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *KalturaAccessControlScope) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *KalturaAccessControlScope) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetKs

`func (o *KalturaAccessControlScope) GetKs() string`

GetKs returns the Ks field if non-nil, zero value otherwise.

### GetKsOk

`func (o *KalturaAccessControlScope) GetKsOk() (*string, bool)`

GetKsOk returns a tuple with the Ks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKs

`func (o *KalturaAccessControlScope) SetKs(v string)`

SetKs sets Ks field to given value.

### HasKs

`func (o *KalturaAccessControlScope) HasKs() bool`

HasKs returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaAccessControlScope) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaAccessControlScope) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaAccessControlScope) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaAccessControlScope) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetReferrer

`func (o *KalturaAccessControlScope) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *KalturaAccessControlScope) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *KalturaAccessControlScope) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *KalturaAccessControlScope) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetTime

`func (o *KalturaAccessControlScope) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *KalturaAccessControlScope) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *KalturaAccessControlScope) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *KalturaAccessControlScope) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUserAgent

`func (o *KalturaAccessControlScope) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *KalturaAccessControlScope) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *KalturaAccessControlScope) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *KalturaAccessControlScope) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



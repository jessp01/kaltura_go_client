# KalturaAppToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation time as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Expiry** | Pointer to **int32** | Expiry time of current token (unix timestamp in seconds) | [optional] 
**HashType** | Pointer to **string** | Enum Type: &#x60;KalturaAppTokenHashType&#x60; | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60;  The id of the application token | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**SessionDuration** | Pointer to **int32** | Expiry duration of KS (Kaltura Session) that created using the current token (in seconds) | [optional] 
**SessionPrivileges** | Pointer to **string** | Comma separated privileges to be applied on KS (Kaltura Session) that created using the current token | [optional] 
**SessionType** | Pointer to **int32** | Enum Type: &#x60;KalturaSessionType&#x60;  Type of KS (Kaltura Session) that created using the current token | [optional] 
**SessionUserId** | Pointer to **string** | User id of KS (Kaltura Session) that created using the current token | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaAppTokenStatus&#x60;  Application token status | [optional] [readonly] 
**Token** | Pointer to **string** | &#x60;readOnly&#x60;  The application token | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update time as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaAppToken

`func NewKalturaAppToken() *KalturaAppToken`

NewKalturaAppToken instantiates a new KalturaAppToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAppTokenWithDefaults

`func NewKalturaAppTokenWithDefaults() *KalturaAppToken`

NewKalturaAppTokenWithDefaults instantiates a new KalturaAppToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaAppToken) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAppToken) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAppToken) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAppToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAppToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAppToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAppToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAppToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiry

`func (o *KalturaAppToken) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *KalturaAppToken) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *KalturaAppToken) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *KalturaAppToken) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetHashType

`func (o *KalturaAppToken) GetHashType() string`

GetHashType returns the HashType field if non-nil, zero value otherwise.

### GetHashTypeOk

`func (o *KalturaAppToken) GetHashTypeOk() (*string, bool)`

GetHashTypeOk returns a tuple with the HashType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashType

`func (o *KalturaAppToken) SetHashType(v string)`

SetHashType sets HashType field to given value.

### HasHashType

`func (o *KalturaAppToken) HasHashType() bool`

HasHashType returns a boolean if a field has been set.

### GetId

`func (o *KalturaAppToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAppToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAppToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAppToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAppToken) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAppToken) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAppToken) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAppToken) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSessionDuration

`func (o *KalturaAppToken) GetSessionDuration() int32`

GetSessionDuration returns the SessionDuration field if non-nil, zero value otherwise.

### GetSessionDurationOk

`func (o *KalturaAppToken) GetSessionDurationOk() (*int32, bool)`

GetSessionDurationOk returns a tuple with the SessionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDuration

`func (o *KalturaAppToken) SetSessionDuration(v int32)`

SetSessionDuration sets SessionDuration field to given value.

### HasSessionDuration

`func (o *KalturaAppToken) HasSessionDuration() bool`

HasSessionDuration returns a boolean if a field has been set.

### GetSessionPrivileges

`func (o *KalturaAppToken) GetSessionPrivileges() string`

GetSessionPrivileges returns the SessionPrivileges field if non-nil, zero value otherwise.

### GetSessionPrivilegesOk

`func (o *KalturaAppToken) GetSessionPrivilegesOk() (*string, bool)`

GetSessionPrivilegesOk returns a tuple with the SessionPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPrivileges

`func (o *KalturaAppToken) SetSessionPrivileges(v string)`

SetSessionPrivileges sets SessionPrivileges field to given value.

### HasSessionPrivileges

`func (o *KalturaAppToken) HasSessionPrivileges() bool`

HasSessionPrivileges returns a boolean if a field has been set.

### GetSessionType

`func (o *KalturaAppToken) GetSessionType() int32`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *KalturaAppToken) GetSessionTypeOk() (*int32, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *KalturaAppToken) SetSessionType(v int32)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *KalturaAppToken) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetSessionUserId

`func (o *KalturaAppToken) GetSessionUserId() string`

GetSessionUserId returns the SessionUserId field if non-nil, zero value otherwise.

### GetSessionUserIdOk

`func (o *KalturaAppToken) GetSessionUserIdOk() (*string, bool)`

GetSessionUserIdOk returns a tuple with the SessionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserId

`func (o *KalturaAppToken) SetSessionUserId(v string)`

SetSessionUserId sets SessionUserId field to given value.

### HasSessionUserId

`func (o *KalturaAppToken) HasSessionUserId() bool`

HasSessionUserId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaAppToken) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaAppToken) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaAppToken) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaAppToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *KalturaAppToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KalturaAppToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KalturaAppToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KalturaAppToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaAppToken) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaAppToken) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaAppToken) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaAppToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



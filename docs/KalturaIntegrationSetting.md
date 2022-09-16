# KalturaIntegrationSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ConversionProfileId** | Pointer to **int32** |  | [optional] 
**CreateUserIfNotExist** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**CreatedAt** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefaultUserId** | Pointer to **string** |  | [optional] 
**DeletionPolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**HandleParticipantsMode** | Pointer to **int32** | Enum Type: &#x60;KalturaHandleParticipantsMode&#x60; | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaVendorIntegrationStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaIntegrationSetting

`func NewKalturaIntegrationSetting() *KalturaIntegrationSetting`

NewKalturaIntegrationSetting instantiates a new KalturaIntegrationSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaIntegrationSettingWithDefaults

`func NewKalturaIntegrationSettingWithDefaults() *KalturaIntegrationSetting`

NewKalturaIntegrationSettingWithDefaults instantiates a new KalturaIntegrationSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *KalturaIntegrationSetting) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *KalturaIntegrationSetting) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *KalturaIntegrationSetting) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *KalturaIntegrationSetting) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetConversionProfileId

`func (o *KalturaIntegrationSetting) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *KalturaIntegrationSetting) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *KalturaIntegrationSetting) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *KalturaIntegrationSetting) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetCreateUserIfNotExist

`func (o *KalturaIntegrationSetting) GetCreateUserIfNotExist() int32`

GetCreateUserIfNotExist returns the CreateUserIfNotExist field if non-nil, zero value otherwise.

### GetCreateUserIfNotExistOk

`func (o *KalturaIntegrationSetting) GetCreateUserIfNotExistOk() (*int32, bool)`

GetCreateUserIfNotExistOk returns a tuple with the CreateUserIfNotExist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUserIfNotExist

`func (o *KalturaIntegrationSetting) SetCreateUserIfNotExist(v int32)`

SetCreateUserIfNotExist sets CreateUserIfNotExist field to given value.

### HasCreateUserIfNotExist

`func (o *KalturaIntegrationSetting) HasCreateUserIfNotExist() bool`

HasCreateUserIfNotExist returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaIntegrationSetting) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaIntegrationSetting) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaIntegrationSetting) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaIntegrationSetting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultUserId

`func (o *KalturaIntegrationSetting) GetDefaultUserId() string`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *KalturaIntegrationSetting) GetDefaultUserIdOk() (*string, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *KalturaIntegrationSetting) SetDefaultUserId(v string)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *KalturaIntegrationSetting) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### GetDeletionPolicy

`func (o *KalturaIntegrationSetting) GetDeletionPolicy() int32`

GetDeletionPolicy returns the DeletionPolicy field if non-nil, zero value otherwise.

### GetDeletionPolicyOk

`func (o *KalturaIntegrationSetting) GetDeletionPolicyOk() (*int32, bool)`

GetDeletionPolicyOk returns a tuple with the DeletionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionPolicy

`func (o *KalturaIntegrationSetting) SetDeletionPolicy(v int32)`

SetDeletionPolicy sets DeletionPolicy field to given value.

### HasDeletionPolicy

`func (o *KalturaIntegrationSetting) HasDeletionPolicy() bool`

HasDeletionPolicy returns a boolean if a field has been set.

### GetHandleParticipantsMode

`func (o *KalturaIntegrationSetting) GetHandleParticipantsMode() int32`

GetHandleParticipantsMode returns the HandleParticipantsMode field if non-nil, zero value otherwise.

### GetHandleParticipantsModeOk

`func (o *KalturaIntegrationSetting) GetHandleParticipantsModeOk() (*int32, bool)`

GetHandleParticipantsModeOk returns a tuple with the HandleParticipantsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleParticipantsMode

`func (o *KalturaIntegrationSetting) SetHandleParticipantsMode(v int32)`

SetHandleParticipantsMode sets HandleParticipantsMode field to given value.

### HasHandleParticipantsMode

`func (o *KalturaIntegrationSetting) HasHandleParticipantsMode() bool`

HasHandleParticipantsMode returns a boolean if a field has been set.

### GetId

`func (o *KalturaIntegrationSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaIntegrationSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaIntegrationSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaIntegrationSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaIntegrationSetting) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaIntegrationSetting) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaIntegrationSetting) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaIntegrationSetting) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaIntegrationSetting) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaIntegrationSetting) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaIntegrationSetting) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaIntegrationSetting) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaIntegrationSetting) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaIntegrationSetting) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaIntegrationSetting) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaIntegrationSetting) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaIntegrationSetting) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaIntegrationSetting) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaIntegrationSetting) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaIntegrationSetting) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



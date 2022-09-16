# KalturaDrmPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** | Duration in days the license is effective | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LicenseExpirationPolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaDrmLicenseExpirationPolicy&#x60; | [optional] 
**LicenseParams** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**LicenseType** | Pointer to **string** | Enum Type: &#x60;KalturaDrmLicenseType&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**Provider** | Pointer to **string** | Enum Type: &#x60;KalturaDrmProviderType&#x60; | [optional] 
**Scenario** | Pointer to **string** | Enum Type: &#x60;KalturaDrmLicenseScenario&#x60; | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaDrmPolicyStatus&#x60; | [optional] 
**SystemName** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaDrmPolicy

`func NewKalturaDrmPolicy() *KalturaDrmPolicy`

NewKalturaDrmPolicy instantiates a new KalturaDrmPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDrmPolicyWithDefaults

`func NewKalturaDrmPolicyWithDefaults() *KalturaDrmPolicy`

NewKalturaDrmPolicyWithDefaults instantiates a new KalturaDrmPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaDrmPolicy) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDrmPolicy) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDrmPolicy) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDrmPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaDrmPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaDrmPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaDrmPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaDrmPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *KalturaDrmPolicy) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KalturaDrmPolicy) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KalturaDrmPolicy) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KalturaDrmPolicy) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *KalturaDrmPolicy) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDrmPolicy) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDrmPolicy) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDrmPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseExpirationPolicy

`func (o *KalturaDrmPolicy) GetLicenseExpirationPolicy() int32`

GetLicenseExpirationPolicy returns the LicenseExpirationPolicy field if non-nil, zero value otherwise.

### GetLicenseExpirationPolicyOk

`func (o *KalturaDrmPolicy) GetLicenseExpirationPolicyOk() (*int32, bool)`

GetLicenseExpirationPolicyOk returns a tuple with the LicenseExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpirationPolicy

`func (o *KalturaDrmPolicy) SetLicenseExpirationPolicy(v int32)`

SetLicenseExpirationPolicy sets LicenseExpirationPolicy field to given value.

### HasLicenseExpirationPolicy

`func (o *KalturaDrmPolicy) HasLicenseExpirationPolicy() bool`

HasLicenseExpirationPolicy returns a boolean if a field has been set.

### GetLicenseParams

`func (o *KalturaDrmPolicy) GetLicenseParams() []KalturaKeyValue`

GetLicenseParams returns the LicenseParams field if non-nil, zero value otherwise.

### GetLicenseParamsOk

`func (o *KalturaDrmPolicy) GetLicenseParamsOk() (*[]KalturaKeyValue, bool)`

GetLicenseParamsOk returns a tuple with the LicenseParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseParams

`func (o *KalturaDrmPolicy) SetLicenseParams(v []KalturaKeyValue)`

SetLicenseParams sets LicenseParams field to given value.

### HasLicenseParams

`func (o *KalturaDrmPolicy) HasLicenseParams() bool`

HasLicenseParams returns a boolean if a field has been set.

### GetLicenseType

`func (o *KalturaDrmPolicy) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *KalturaDrmPolicy) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *KalturaDrmPolicy) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *KalturaDrmPolicy) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetName

`func (o *KalturaDrmPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDrmPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDrmPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDrmPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDrmPolicy) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDrmPolicy) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDrmPolicy) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDrmPolicy) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDrmPolicy) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDrmPolicy) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDrmPolicy) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDrmPolicy) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProvider

`func (o *KalturaDrmPolicy) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KalturaDrmPolicy) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KalturaDrmPolicy) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KalturaDrmPolicy) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScenario

`func (o *KalturaDrmPolicy) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *KalturaDrmPolicy) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *KalturaDrmPolicy) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *KalturaDrmPolicy) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDrmPolicy) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDrmPolicy) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDrmPolicy) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDrmPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaDrmPolicy) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaDrmPolicy) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaDrmPolicy) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaDrmPolicy) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDrmPolicy) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDrmPolicy) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDrmPolicy) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDrmPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KalturaDrmProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DefaultPolicy** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LicenseServerUrl** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**Provider** | Pointer to **string** | Enum Type: &#x60;KalturaDrmProviderType&#x60; | [optional] 
**SigningKey** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaDrmProfileStatus&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaDrmProfile

`func NewKalturaDrmProfile() *KalturaDrmProfile`

NewKalturaDrmProfile instantiates a new KalturaDrmProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDrmProfileWithDefaults

`func NewKalturaDrmProfileWithDefaults() *KalturaDrmProfile`

NewKalturaDrmProfileWithDefaults instantiates a new KalturaDrmProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaDrmProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDrmProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDrmProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDrmProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultPolicy

`func (o *KalturaDrmProfile) GetDefaultPolicy() string`

GetDefaultPolicy returns the DefaultPolicy field if non-nil, zero value otherwise.

### GetDefaultPolicyOk

`func (o *KalturaDrmProfile) GetDefaultPolicyOk() (*string, bool)`

GetDefaultPolicyOk returns a tuple with the DefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPolicy

`func (o *KalturaDrmProfile) SetDefaultPolicy(v string)`

SetDefaultPolicy sets DefaultPolicy field to given value.

### HasDefaultPolicy

`func (o *KalturaDrmProfile) HasDefaultPolicy() bool`

HasDefaultPolicy returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaDrmProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaDrmProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaDrmProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaDrmProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaDrmProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDrmProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDrmProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDrmProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseServerUrl

`func (o *KalturaDrmProfile) GetLicenseServerUrl() string`

GetLicenseServerUrl returns the LicenseServerUrl field if non-nil, zero value otherwise.

### GetLicenseServerUrlOk

`func (o *KalturaDrmProfile) GetLicenseServerUrlOk() (*string, bool)`

GetLicenseServerUrlOk returns a tuple with the LicenseServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseServerUrl

`func (o *KalturaDrmProfile) SetLicenseServerUrl(v string)`

SetLicenseServerUrl sets LicenseServerUrl field to given value.

### HasLicenseServerUrl

`func (o *KalturaDrmProfile) HasLicenseServerUrl() bool`

HasLicenseServerUrl returns a boolean if a field has been set.

### GetName

`func (o *KalturaDrmProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDrmProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDrmProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDrmProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDrmProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDrmProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDrmProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDrmProfile) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDrmProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDrmProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDrmProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDrmProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProvider

`func (o *KalturaDrmProfile) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KalturaDrmProfile) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KalturaDrmProfile) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *KalturaDrmProfile) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSigningKey

`func (o *KalturaDrmProfile) GetSigningKey() string`

GetSigningKey returns the SigningKey field if non-nil, zero value otherwise.

### GetSigningKeyOk

`func (o *KalturaDrmProfile) GetSigningKeyOk() (*string, bool)`

GetSigningKeyOk returns a tuple with the SigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningKey

`func (o *KalturaDrmProfile) SetSigningKey(v string)`

SetSigningKey sets SigningKey field to given value.

### HasSigningKey

`func (o *KalturaDrmProfile) HasSigningKey() bool`

HasSigningKey returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDrmProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDrmProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDrmProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDrmProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDrmProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDrmProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDrmProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDrmProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KalturaAccessControlProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation time as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Access Control Profile | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Access Control Profile | [optional] [readonly] 
**IsDefault** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  True if this access control profile is the partner default | [optional] 
**Name** | Pointer to **string** | The name of the Access Control Profile | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Rules** | Pointer to [**[]KalturaRule**](KalturaRule.md) |  | [optional] 
**SystemName** | Pointer to **string** | System name of the Access Control Profile | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update time as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaAccessControlProfile

`func NewKalturaAccessControlProfile() *KalturaAccessControlProfile`

NewKalturaAccessControlProfile instantiates a new KalturaAccessControlProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAccessControlProfileWithDefaults

`func NewKalturaAccessControlProfileWithDefaults() *KalturaAccessControlProfile`

NewKalturaAccessControlProfileWithDefaults instantiates a new KalturaAccessControlProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaAccessControlProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAccessControlProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAccessControlProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAccessControlProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAccessControlProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAccessControlProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAccessControlProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAccessControlProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaAccessControlProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAccessControlProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAccessControlProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAccessControlProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *KalturaAccessControlProfile) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *KalturaAccessControlProfile) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *KalturaAccessControlProfile) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *KalturaAccessControlProfile) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *KalturaAccessControlProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaAccessControlProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaAccessControlProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaAccessControlProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAccessControlProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAccessControlProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAccessControlProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAccessControlProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetRules

`func (o *KalturaAccessControlProfile) GetRules() []KalturaRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *KalturaAccessControlProfile) GetRulesOk() (*[]KalturaRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *KalturaAccessControlProfile) SetRules(v []KalturaRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *KalturaAccessControlProfile) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaAccessControlProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaAccessControlProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaAccessControlProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaAccessControlProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaAccessControlProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaAccessControlProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaAccessControlProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaAccessControlProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



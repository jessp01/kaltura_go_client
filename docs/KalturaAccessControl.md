# KalturaAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainsUnsuportedRestrictions** | Pointer to **bool** | &#x60;readOnly&#x60;  Indicates that the access control profile is new and should be handled using KalturaAccessControlProfile object and accessControlProfile service | [optional] [readonly] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Access Control Profile | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Access Control Profile | [optional] [readonly] 
**IsDefault** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  True if this Conversion Profile is the default | [optional] 
**Name** | Pointer to **string** | The name of the Access Control Profile | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Restrictions** | Pointer to [**[]KalturaBaseRestriction**](KalturaBaseRestriction.md) |  | [optional] 
**SystemName** | Pointer to **string** | System name of the Access Control Profile | [optional] 

## Methods

### NewKalturaAccessControl

`func NewKalturaAccessControl() *KalturaAccessControl`

NewKalturaAccessControl instantiates a new KalturaAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAccessControlWithDefaults

`func NewKalturaAccessControlWithDefaults() *KalturaAccessControl`

NewKalturaAccessControlWithDefaults instantiates a new KalturaAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainsUnsuportedRestrictions

`func (o *KalturaAccessControl) GetContainsUnsuportedRestrictions() bool`

GetContainsUnsuportedRestrictions returns the ContainsUnsuportedRestrictions field if non-nil, zero value otherwise.

### GetContainsUnsuportedRestrictionsOk

`func (o *KalturaAccessControl) GetContainsUnsuportedRestrictionsOk() (*bool, bool)`

GetContainsUnsuportedRestrictionsOk returns a tuple with the ContainsUnsuportedRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsUnsuportedRestrictions

`func (o *KalturaAccessControl) SetContainsUnsuportedRestrictions(v bool)`

SetContainsUnsuportedRestrictions sets ContainsUnsuportedRestrictions field to given value.

### HasContainsUnsuportedRestrictions

`func (o *KalturaAccessControl) HasContainsUnsuportedRestrictions() bool`

HasContainsUnsuportedRestrictions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaAccessControl) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAccessControl) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAccessControl) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAccessControl) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAccessControl) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAccessControl) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAccessControl) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAccessControl) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaAccessControl) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAccessControl) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAccessControl) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAccessControl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *KalturaAccessControl) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *KalturaAccessControl) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *KalturaAccessControl) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *KalturaAccessControl) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *KalturaAccessControl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaAccessControl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaAccessControl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaAccessControl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAccessControl) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAccessControl) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAccessControl) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAccessControl) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetRestrictions

`func (o *KalturaAccessControl) GetRestrictions() []KalturaBaseRestriction`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *KalturaAccessControl) GetRestrictionsOk() (*[]KalturaBaseRestriction, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *KalturaAccessControl) SetRestrictions(v []KalturaBaseRestriction)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *KalturaAccessControl) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaAccessControl) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaAccessControl) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaAccessControl) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaAccessControl) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



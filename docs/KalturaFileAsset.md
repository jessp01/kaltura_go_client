# KalturaFileAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FileAssetObjectType** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaFileAssetObjectType&#x60; | [optional] 
**FileExt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaFileAssetStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaFileAsset

`func NewKalturaFileAsset() *KalturaFileAsset`

NewKalturaFileAsset instantiates a new KalturaFileAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFileAssetWithDefaults

`func NewKalturaFileAssetWithDefaults() *KalturaFileAsset`

NewKalturaFileAssetWithDefaults instantiates a new KalturaFileAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaFileAsset) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaFileAsset) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaFileAsset) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaFileAsset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFileAssetObjectType

`func (o *KalturaFileAsset) GetFileAssetObjectType() string`

GetFileAssetObjectType returns the FileAssetObjectType field if non-nil, zero value otherwise.

### GetFileAssetObjectTypeOk

`func (o *KalturaFileAsset) GetFileAssetObjectTypeOk() (*string, bool)`

GetFileAssetObjectTypeOk returns a tuple with the FileAssetObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAssetObjectType

`func (o *KalturaFileAsset) SetFileAssetObjectType(v string)`

SetFileAssetObjectType sets FileAssetObjectType field to given value.

### HasFileAssetObjectType

`func (o *KalturaFileAsset) HasFileAssetObjectType() bool`

HasFileAssetObjectType returns a boolean if a field has been set.

### GetFileExt

`func (o *KalturaFileAsset) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *KalturaFileAsset) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *KalturaFileAsset) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.

### HasFileExt

`func (o *KalturaFileAsset) HasFileExt() bool`

HasFileExt returns a boolean if a field has been set.

### GetId

`func (o *KalturaFileAsset) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaFileAsset) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaFileAsset) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaFileAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaFileAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaFileAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaFileAsset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaFileAsset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaFileAsset) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaFileAsset) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaFileAsset) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaFileAsset) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaFileAsset) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaFileAsset) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaFileAsset) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaFileAsset) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaFileAsset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaFileAsset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaFileAsset) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaFileAsset) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaFileAsset) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaFileAsset) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaFileAsset) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaFileAsset) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaFileAsset) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaFileAsset) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaFileAsset) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaFileAsset) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaFileAsset) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaFileAsset) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaFileAsset) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaFileAsset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



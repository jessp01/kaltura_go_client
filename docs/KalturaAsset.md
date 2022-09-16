# KalturaAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualSourceAssetParamsIds** | Pointer to **string** | Comma separated list of source flavor params ids | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DeletedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** | &#x60;readOnly&#x60;  System description, error message, warnings and failure cause. | [optional] [readonly] 
**EntryId** | Pointer to **string** | &#x60;readOnly&#x60;  The entry ID of the Flavor Asset | [optional] [readonly] 
**FileExt** | Pointer to **string** | &#x60;insertOnly&#x60;  The file extension | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60;  The ID of the Flavor Asset | [optional] [readonly] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerData** | Pointer to **string** | Partner private data | [optional] 
**PartnerDescription** | Pointer to **string** | Partner friendly description | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Size** | Pointer to **int32** | &#x60;readOnly&#x60;  The size (in KBytes) of the Flavor Asset | [optional] [readonly] 
**SizeInBytes** | Pointer to **int32** | &#x60;readOnly&#x60;  The size (in Bytes) of the asset | [optional] [readonly] 
**Tags** | Pointer to **string** | Tags used to identify the Flavor Asset in various scenarios | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60;  The version of the Flavor Asset | [optional] [readonly] 

## Methods

### NewKalturaAsset

`func NewKalturaAsset() *KalturaAsset`

NewKalturaAsset instantiates a new KalturaAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAssetWithDefaults

`func NewKalturaAssetWithDefaults() *KalturaAsset`

NewKalturaAssetWithDefaults instantiates a new KalturaAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualSourceAssetParamsIds

`func (o *KalturaAsset) GetActualSourceAssetParamsIds() string`

GetActualSourceAssetParamsIds returns the ActualSourceAssetParamsIds field if non-nil, zero value otherwise.

### GetActualSourceAssetParamsIdsOk

`func (o *KalturaAsset) GetActualSourceAssetParamsIdsOk() (*string, bool)`

GetActualSourceAssetParamsIdsOk returns a tuple with the ActualSourceAssetParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSourceAssetParamsIds

`func (o *KalturaAsset) SetActualSourceAssetParamsIds(v string)`

SetActualSourceAssetParamsIds sets ActualSourceAssetParamsIds field to given value.

### HasActualSourceAssetParamsIds

`func (o *KalturaAsset) HasActualSourceAssetParamsIds() bool`

HasActualSourceAssetParamsIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaAsset) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAsset) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAsset) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAsset) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *KalturaAsset) GetDeletedAt() int32`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *KalturaAsset) GetDeletedAtOk() (*int32, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *KalturaAsset) SetDeletedAt(v int32)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *KalturaAsset) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaAsset) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaAsset) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaAsset) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaAsset) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetFileExt

`func (o *KalturaAsset) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *KalturaAsset) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *KalturaAsset) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.

### HasFileExt

`func (o *KalturaAsset) HasFileExt() bool`

HasFileExt returns a boolean if a field has been set.

### GetId

`func (o *KalturaAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAsset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAsset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaAsset) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaAsset) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaAsset) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaAsset) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaAsset) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaAsset) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaAsset) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaAsset) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerDescription

`func (o *KalturaAsset) GetPartnerDescription() string`

GetPartnerDescription returns the PartnerDescription field if non-nil, zero value otherwise.

### GetPartnerDescriptionOk

`func (o *KalturaAsset) GetPartnerDescriptionOk() (*string, bool)`

GetPartnerDescriptionOk returns a tuple with the PartnerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerDescription

`func (o *KalturaAsset) SetPartnerDescription(v string)`

SetPartnerDescription sets PartnerDescription field to given value.

### HasPartnerDescription

`func (o *KalturaAsset) HasPartnerDescription() bool`

HasPartnerDescription returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAsset) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAsset) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAsset) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAsset) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetSize

`func (o *KalturaAsset) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *KalturaAsset) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *KalturaAsset) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *KalturaAsset) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeInBytes

`func (o *KalturaAsset) GetSizeInBytes() int32`

GetSizeInBytes returns the SizeInBytes field if non-nil, zero value otherwise.

### GetSizeInBytesOk

`func (o *KalturaAsset) GetSizeInBytesOk() (*int32, bool)`

GetSizeInBytesOk returns a tuple with the SizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeInBytes

`func (o *KalturaAsset) SetSizeInBytes(v int32)`

SetSizeInBytes sets SizeInBytes field to given value.

### HasSizeInBytes

`func (o *KalturaAsset) HasSizeInBytes() bool`

HasSizeInBytes returns a boolean if a field has been set.

### GetTags

`func (o *KalturaAsset) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaAsset) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaAsset) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaAsset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaAsset) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaAsset) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaAsset) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaAsset) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaAsset) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaAsset) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaAsset) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaAsset) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



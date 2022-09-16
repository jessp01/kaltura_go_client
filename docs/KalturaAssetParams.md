# KalturaAssetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Flavor Params | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Flavor Params | [optional] [readonly] 
**IsSystemDefault** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60;  True if those Flavor Params are part of system defaults | [optional] [readonly] 
**MediaParserType** | Pointer to **string** | Enum Type: &#x60;KalturaMediaParserType&#x60;  Media parser type to be used for post-conversion validation | [optional] 
**Name** | Pointer to **string** | The name of the Flavor Params | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** |  | [optional] 
**RemoteStorageProfileIds** | Pointer to **int32** | Comma seperated ids of remote storage profiles that the flavor distributed to, the distribution done by the conversion engine | [optional] 
**RequiredPermissions** | Pointer to [**[]KalturaString**](KalturaString.md) |  | [optional] 
**SourceAssetParamsIds** | Pointer to **string** | Comma seperated ids of source flavor params this flavor is created from | [optional] 
**SourceRemoteStorageProfileId** | Pointer to **int32** | Id of remote storage profile that used to get the source, zero indicates Kaltura data center | [optional] 
**SystemName** | Pointer to **string** | System name of the Flavor Params | [optional] 
**Tags** | Pointer to **string** | The Flavor Params tags are used to identify the flavor for different usage (e.g. web, hd, mobile) | [optional] 

## Methods

### NewKalturaAssetParams

`func NewKalturaAssetParams() *KalturaAssetParams`

NewKalturaAssetParams instantiates a new KalturaAssetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAssetParamsWithDefaults

`func NewKalturaAssetParamsWithDefaults() *KalturaAssetParams`

NewKalturaAssetParamsWithDefaults instantiates a new KalturaAssetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaAssetParams) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAssetParams) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAssetParams) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAssetParams) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAssetParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAssetParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAssetParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAssetParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaAssetParams) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAssetParams) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAssetParams) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAssetParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSystemDefault

`func (o *KalturaAssetParams) GetIsSystemDefault() int32`

GetIsSystemDefault returns the IsSystemDefault field if non-nil, zero value otherwise.

### GetIsSystemDefaultOk

`func (o *KalturaAssetParams) GetIsSystemDefaultOk() (*int32, bool)`

GetIsSystemDefaultOk returns a tuple with the IsSystemDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefault

`func (o *KalturaAssetParams) SetIsSystemDefault(v int32)`

SetIsSystemDefault sets IsSystemDefault field to given value.

### HasIsSystemDefault

`func (o *KalturaAssetParams) HasIsSystemDefault() bool`

HasIsSystemDefault returns a boolean if a field has been set.

### GetMediaParserType

`func (o *KalturaAssetParams) GetMediaParserType() string`

GetMediaParserType returns the MediaParserType field if non-nil, zero value otherwise.

### GetMediaParserTypeOk

`func (o *KalturaAssetParams) GetMediaParserTypeOk() (*string, bool)`

GetMediaParserTypeOk returns a tuple with the MediaParserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaParserType

`func (o *KalturaAssetParams) SetMediaParserType(v string)`

SetMediaParserType sets MediaParserType field to given value.

### HasMediaParserType

`func (o *KalturaAssetParams) HasMediaParserType() bool`

HasMediaParserType returns a boolean if a field has been set.

### GetName

`func (o *KalturaAssetParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaAssetParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaAssetParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaAssetParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaAssetParams) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaAssetParams) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaAssetParams) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaAssetParams) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAssetParams) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAssetParams) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAssetParams) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAssetParams) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetRemoteStorageProfileIds

`func (o *KalturaAssetParams) GetRemoteStorageProfileIds() int32`

GetRemoteStorageProfileIds returns the RemoteStorageProfileIds field if non-nil, zero value otherwise.

### GetRemoteStorageProfileIdsOk

`func (o *KalturaAssetParams) GetRemoteStorageProfileIdsOk() (*int32, bool)`

GetRemoteStorageProfileIdsOk returns a tuple with the RemoteStorageProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteStorageProfileIds

`func (o *KalturaAssetParams) SetRemoteStorageProfileIds(v int32)`

SetRemoteStorageProfileIds sets RemoteStorageProfileIds field to given value.

### HasRemoteStorageProfileIds

`func (o *KalturaAssetParams) HasRemoteStorageProfileIds() bool`

HasRemoteStorageProfileIds returns a boolean if a field has been set.

### GetRequiredPermissions

`func (o *KalturaAssetParams) GetRequiredPermissions() []KalturaString`

GetRequiredPermissions returns the RequiredPermissions field if non-nil, zero value otherwise.

### GetRequiredPermissionsOk

`func (o *KalturaAssetParams) GetRequiredPermissionsOk() (*[]KalturaString, bool)`

GetRequiredPermissionsOk returns a tuple with the RequiredPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPermissions

`func (o *KalturaAssetParams) SetRequiredPermissions(v []KalturaString)`

SetRequiredPermissions sets RequiredPermissions field to given value.

### HasRequiredPermissions

`func (o *KalturaAssetParams) HasRequiredPermissions() bool`

HasRequiredPermissions returns a boolean if a field has been set.

### GetSourceAssetParamsIds

`func (o *KalturaAssetParams) GetSourceAssetParamsIds() string`

GetSourceAssetParamsIds returns the SourceAssetParamsIds field if non-nil, zero value otherwise.

### GetSourceAssetParamsIdsOk

`func (o *KalturaAssetParams) GetSourceAssetParamsIdsOk() (*string, bool)`

GetSourceAssetParamsIdsOk returns a tuple with the SourceAssetParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAssetParamsIds

`func (o *KalturaAssetParams) SetSourceAssetParamsIds(v string)`

SetSourceAssetParamsIds sets SourceAssetParamsIds field to given value.

### HasSourceAssetParamsIds

`func (o *KalturaAssetParams) HasSourceAssetParamsIds() bool`

HasSourceAssetParamsIds returns a boolean if a field has been set.

### GetSourceRemoteStorageProfileId

`func (o *KalturaAssetParams) GetSourceRemoteStorageProfileId() int32`

GetSourceRemoteStorageProfileId returns the SourceRemoteStorageProfileId field if non-nil, zero value otherwise.

### GetSourceRemoteStorageProfileIdOk

`func (o *KalturaAssetParams) GetSourceRemoteStorageProfileIdOk() (*int32, bool)`

GetSourceRemoteStorageProfileIdOk returns a tuple with the SourceRemoteStorageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRemoteStorageProfileId

`func (o *KalturaAssetParams) SetSourceRemoteStorageProfileId(v int32)`

SetSourceRemoteStorageProfileId sets SourceRemoteStorageProfileId field to given value.

### HasSourceRemoteStorageProfileId

`func (o *KalturaAssetParams) HasSourceRemoteStorageProfileId() bool`

HasSourceRemoteStorageProfileId returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaAssetParams) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaAssetParams) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaAssetParams) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaAssetParams) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaAssetParams) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaAssetParams) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaAssetParams) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaAssetParams) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



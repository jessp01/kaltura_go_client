# KalturaMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MetadataObjectType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaMetadataObjectType&#x60; | [optional] [readonly] 
**MetadataProfileId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MetadataProfileVersion** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaMetadataStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Xml** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaMetadata

`func NewKalturaMetadata() *KalturaMetadata`

NewKalturaMetadata instantiates a new KalturaMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaMetadataWithDefaults

`func NewKalturaMetadataWithDefaults() *KalturaMetadata`

NewKalturaMetadataWithDefaults instantiates a new KalturaMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaMetadata) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaMetadata) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaMetadata) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *KalturaMetadata) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaMetadata) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaMetadata) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataObjectType

`func (o *KalturaMetadata) GetMetadataObjectType() string`

GetMetadataObjectType returns the MetadataObjectType field if non-nil, zero value otherwise.

### GetMetadataObjectTypeOk

`func (o *KalturaMetadata) GetMetadataObjectTypeOk() (*string, bool)`

GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObjectType

`func (o *KalturaMetadata) SetMetadataObjectType(v string)`

SetMetadataObjectType sets MetadataObjectType field to given value.

### HasMetadataObjectType

`func (o *KalturaMetadata) HasMetadataObjectType() bool`

HasMetadataObjectType returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *KalturaMetadata) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *KalturaMetadata) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *KalturaMetadata) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *KalturaMetadata) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetMetadataProfileVersion

`func (o *KalturaMetadata) GetMetadataProfileVersion() int32`

GetMetadataProfileVersion returns the MetadataProfileVersion field if non-nil, zero value otherwise.

### GetMetadataProfileVersionOk

`func (o *KalturaMetadata) GetMetadataProfileVersionOk() (*int32, bool)`

GetMetadataProfileVersionOk returns a tuple with the MetadataProfileVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileVersion

`func (o *KalturaMetadata) SetMetadataProfileVersion(v int32)`

SetMetadataProfileVersion sets MetadataProfileVersion field to given value.

### HasMetadataProfileVersion

`func (o *KalturaMetadata) HasMetadataProfileVersion() bool`

HasMetadataProfileVersion returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaMetadata) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaMetadata) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaMetadata) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaMetadata) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaMetadata) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaMetadata) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaMetadata) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaMetadata) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaMetadata) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaMetadata) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaMetadata) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaMetadata) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaMetadata) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaMetadata) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaMetadata) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaMetadata) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaMetadata) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaMetadata) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXml

`func (o *KalturaMetadata) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *KalturaMetadata) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *KalturaMetadata) SetXml(v string)`

SetXml sets Xml field to given value.

### HasXml

`func (o *KalturaMetadata) HasXml() bool`

HasXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



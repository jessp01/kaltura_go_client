# KalturaMetadataProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMode** | Pointer to **int32** | Enum Type: &#x60;KalturaMetadataProfileCreateMode&#x60; | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**DisableReIndexing** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MetadataObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaMetadataObjectType&#x60; | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaMetadataProfileStatus&#x60; | [optional] [readonly] 
**SystemName** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Views** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Xsd** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Xslt** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaMetadataProfile

`func NewKalturaMetadataProfile() *KalturaMetadataProfile`

NewKalturaMetadataProfile instantiates a new KalturaMetadataProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaMetadataProfileWithDefaults

`func NewKalturaMetadataProfileWithDefaults() *KalturaMetadataProfile`

NewKalturaMetadataProfileWithDefaults instantiates a new KalturaMetadataProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMode

`func (o *KalturaMetadataProfile) GetCreateMode() int32`

GetCreateMode returns the CreateMode field if non-nil, zero value otherwise.

### GetCreateModeOk

`func (o *KalturaMetadataProfile) GetCreateModeOk() (*int32, bool)`

GetCreateModeOk returns a tuple with the CreateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMode

`func (o *KalturaMetadataProfile) SetCreateMode(v int32)`

SetCreateMode sets CreateMode field to given value.

### HasCreateMode

`func (o *KalturaMetadataProfile) HasCreateMode() bool`

HasCreateMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaMetadataProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaMetadataProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaMetadataProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaMetadataProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaMetadataProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaMetadataProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaMetadataProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaMetadataProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableReIndexing

`func (o *KalturaMetadataProfile) GetDisableReIndexing() bool`

GetDisableReIndexing returns the DisableReIndexing field if non-nil, zero value otherwise.

### GetDisableReIndexingOk

`func (o *KalturaMetadataProfile) GetDisableReIndexingOk() (*bool, bool)`

GetDisableReIndexingOk returns a tuple with the DisableReIndexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableReIndexing

`func (o *KalturaMetadataProfile) SetDisableReIndexing(v bool)`

SetDisableReIndexing sets DisableReIndexing field to given value.

### HasDisableReIndexing

`func (o *KalturaMetadataProfile) HasDisableReIndexing() bool`

HasDisableReIndexing returns a boolean if a field has been set.

### GetId

`func (o *KalturaMetadataProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaMetadataProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaMetadataProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaMetadataProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadataObjectType

`func (o *KalturaMetadataProfile) GetMetadataObjectType() string`

GetMetadataObjectType returns the MetadataObjectType field if non-nil, zero value otherwise.

### GetMetadataObjectTypeOk

`func (o *KalturaMetadataProfile) GetMetadataObjectTypeOk() (*string, bool)`

GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataObjectType

`func (o *KalturaMetadataProfile) SetMetadataObjectType(v string)`

SetMetadataObjectType sets MetadataObjectType field to given value.

### HasMetadataObjectType

`func (o *KalturaMetadataProfile) HasMetadataObjectType() bool`

HasMetadataObjectType returns a boolean if a field has been set.

### GetName

`func (o *KalturaMetadataProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaMetadataProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaMetadataProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaMetadataProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaMetadataProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaMetadataProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaMetadataProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaMetadataProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaMetadataProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaMetadataProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaMetadataProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaMetadataProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaMetadataProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaMetadataProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaMetadataProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaMetadataProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaMetadataProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaMetadataProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaMetadataProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaMetadataProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaMetadataProfile) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaMetadataProfile) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaMetadataProfile) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaMetadataProfile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetViews

`func (o *KalturaMetadataProfile) GetViews() string`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *KalturaMetadataProfile) GetViewsOk() (*string, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *KalturaMetadataProfile) SetViews(v string)`

SetViews sets Views field to given value.

### HasViews

`func (o *KalturaMetadataProfile) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetXsd

`func (o *KalturaMetadataProfile) GetXsd() string`

GetXsd returns the Xsd field if non-nil, zero value otherwise.

### GetXsdOk

`func (o *KalturaMetadataProfile) GetXsdOk() (*string, bool)`

GetXsdOk returns a tuple with the Xsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXsd

`func (o *KalturaMetadataProfile) SetXsd(v string)`

SetXsd sets Xsd field to given value.

### HasXsd

`func (o *KalturaMetadataProfile) HasXsd() bool`

HasXsd returns a boolean if a field has been set.

### GetXslt

`func (o *KalturaMetadataProfile) GetXslt() string`

GetXslt returns the Xslt field if non-nil, zero value otherwise.

### GetXsltOk

`func (o *KalturaMetadataProfile) GetXsltOk() (*string, bool)`

GetXsltOk returns a tuple with the Xslt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXslt

`func (o *KalturaMetadataProfile) SetXslt(v string)`

SetXslt sets Xslt field to given value.

### HasXslt

`func (o *KalturaMetadataProfile) HasXslt() bool`

HasXslt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



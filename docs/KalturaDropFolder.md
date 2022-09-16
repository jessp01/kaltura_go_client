# KalturaDropFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoFileDeleteDays** | Pointer to **int32** |  | [optional] 
**CategoriesMetadataFieldName** | Pointer to **string** |  | [optional] 
**ConversionProfileId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Dc** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnforceEntitlement** | Pointer to **bool** |  | [optional] 
**ErrorCode** | Pointer to **string** | Enum Type: &#x60;KalturaDropFolderErrorCode&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**FileDeletePolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaDropFolderFileDeletePolicy&#x60; | [optional] 
**FileDeleteRegex** | Pointer to **string** |  | [optional] 
**FileHandlerConfig** | Pointer to [**KalturaDropFolderFileHandlerConfig**](KalturaDropFolderFileHandlerConfig.md) |  | [optional] 
**FileHandlerType** | Pointer to **string** | Enum Type: &#x60;KalturaDropFolderFileHandlerType&#x60; | [optional] 
**FileNamePatterns** | Pointer to **string** |  | [optional] 
**FileSizeCheckInterval** | Pointer to **int32** | The ammount of time, in seconds, that should pass so that a file with no change in size we&#39;ll be treated as \&quot;finished uploading to folder\&quot; | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IgnoreFileNamePatterns** | Pointer to **string** |  | [optional] 
**Incremental** | Pointer to **bool** |  | [optional] 
**LastAccessedAt** | Pointer to **int32** |  | [optional] 
**LastFileTimestamp** | Pointer to **int32** |  | [optional] 
**MetadataProfileId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ShouldValidateKS** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaDropFolderStatus&#x60; | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Enum Type: &#x60;KalturaDropFolderType&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaDropFolder

`func NewKalturaDropFolder() *KalturaDropFolder`

NewKalturaDropFolder instantiates a new KalturaDropFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDropFolderWithDefaults

`func NewKalturaDropFolderWithDefaults() *KalturaDropFolder`

NewKalturaDropFolderWithDefaults instantiates a new KalturaDropFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoFileDeleteDays

`func (o *KalturaDropFolder) GetAutoFileDeleteDays() int32`

GetAutoFileDeleteDays returns the AutoFileDeleteDays field if non-nil, zero value otherwise.

### GetAutoFileDeleteDaysOk

`func (o *KalturaDropFolder) GetAutoFileDeleteDaysOk() (*int32, bool)`

GetAutoFileDeleteDaysOk returns a tuple with the AutoFileDeleteDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFileDeleteDays

`func (o *KalturaDropFolder) SetAutoFileDeleteDays(v int32)`

SetAutoFileDeleteDays sets AutoFileDeleteDays field to given value.

### HasAutoFileDeleteDays

`func (o *KalturaDropFolder) HasAutoFileDeleteDays() bool`

HasAutoFileDeleteDays returns a boolean if a field has been set.

### GetCategoriesMetadataFieldName

`func (o *KalturaDropFolder) GetCategoriesMetadataFieldName() string`

GetCategoriesMetadataFieldName returns the CategoriesMetadataFieldName field if non-nil, zero value otherwise.

### GetCategoriesMetadataFieldNameOk

`func (o *KalturaDropFolder) GetCategoriesMetadataFieldNameOk() (*string, bool)`

GetCategoriesMetadataFieldNameOk returns a tuple with the CategoriesMetadataFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoriesMetadataFieldName

`func (o *KalturaDropFolder) SetCategoriesMetadataFieldName(v string)`

SetCategoriesMetadataFieldName sets CategoriesMetadataFieldName field to given value.

### HasCategoriesMetadataFieldName

`func (o *KalturaDropFolder) HasCategoriesMetadataFieldName() bool`

HasCategoriesMetadataFieldName returns a boolean if a field has been set.

### GetConversionProfileId

`func (o *KalturaDropFolder) GetConversionProfileId() int32`

GetConversionProfileId returns the ConversionProfileId field if non-nil, zero value otherwise.

### GetConversionProfileIdOk

`func (o *KalturaDropFolder) GetConversionProfileIdOk() (*int32, bool)`

GetConversionProfileIdOk returns a tuple with the ConversionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfileId

`func (o *KalturaDropFolder) SetConversionProfileId(v int32)`

SetConversionProfileId sets ConversionProfileId field to given value.

### HasConversionProfileId

`func (o *KalturaDropFolder) HasConversionProfileId() bool`

HasConversionProfileId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaDropFolder) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDropFolder) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDropFolder) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDropFolder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDc

`func (o *KalturaDropFolder) GetDc() int32`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *KalturaDropFolder) GetDcOk() (*int32, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *KalturaDropFolder) SetDc(v int32)`

SetDc sets Dc field to given value.

### HasDc

`func (o *KalturaDropFolder) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaDropFolder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaDropFolder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaDropFolder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaDropFolder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnforceEntitlement

`func (o *KalturaDropFolder) GetEnforceEntitlement() bool`

GetEnforceEntitlement returns the EnforceEntitlement field if non-nil, zero value otherwise.

### GetEnforceEntitlementOk

`func (o *KalturaDropFolder) GetEnforceEntitlementOk() (*bool, bool)`

GetEnforceEntitlementOk returns a tuple with the EnforceEntitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceEntitlement

`func (o *KalturaDropFolder) SetEnforceEntitlement(v bool)`

SetEnforceEntitlement sets EnforceEntitlement field to given value.

### HasEnforceEntitlement

`func (o *KalturaDropFolder) HasEnforceEntitlement() bool`

HasEnforceEntitlement returns a boolean if a field has been set.

### GetErrorCode

`func (o *KalturaDropFolder) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *KalturaDropFolder) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *KalturaDropFolder) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *KalturaDropFolder) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaDropFolder) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaDropFolder) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaDropFolder) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaDropFolder) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetFileDeletePolicy

`func (o *KalturaDropFolder) GetFileDeletePolicy() int32`

GetFileDeletePolicy returns the FileDeletePolicy field if non-nil, zero value otherwise.

### GetFileDeletePolicyOk

`func (o *KalturaDropFolder) GetFileDeletePolicyOk() (*int32, bool)`

GetFileDeletePolicyOk returns a tuple with the FileDeletePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDeletePolicy

`func (o *KalturaDropFolder) SetFileDeletePolicy(v int32)`

SetFileDeletePolicy sets FileDeletePolicy field to given value.

### HasFileDeletePolicy

`func (o *KalturaDropFolder) HasFileDeletePolicy() bool`

HasFileDeletePolicy returns a boolean if a field has been set.

### GetFileDeleteRegex

`func (o *KalturaDropFolder) GetFileDeleteRegex() string`

GetFileDeleteRegex returns the FileDeleteRegex field if non-nil, zero value otherwise.

### GetFileDeleteRegexOk

`func (o *KalturaDropFolder) GetFileDeleteRegexOk() (*string, bool)`

GetFileDeleteRegexOk returns a tuple with the FileDeleteRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDeleteRegex

`func (o *KalturaDropFolder) SetFileDeleteRegex(v string)`

SetFileDeleteRegex sets FileDeleteRegex field to given value.

### HasFileDeleteRegex

`func (o *KalturaDropFolder) HasFileDeleteRegex() bool`

HasFileDeleteRegex returns a boolean if a field has been set.

### GetFileHandlerConfig

`func (o *KalturaDropFolder) GetFileHandlerConfig() KalturaDropFolderFileHandlerConfig`

GetFileHandlerConfig returns the FileHandlerConfig field if non-nil, zero value otherwise.

### GetFileHandlerConfigOk

`func (o *KalturaDropFolder) GetFileHandlerConfigOk() (*KalturaDropFolderFileHandlerConfig, bool)`

GetFileHandlerConfigOk returns a tuple with the FileHandlerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHandlerConfig

`func (o *KalturaDropFolder) SetFileHandlerConfig(v KalturaDropFolderFileHandlerConfig)`

SetFileHandlerConfig sets FileHandlerConfig field to given value.

### HasFileHandlerConfig

`func (o *KalturaDropFolder) HasFileHandlerConfig() bool`

HasFileHandlerConfig returns a boolean if a field has been set.

### GetFileHandlerType

`func (o *KalturaDropFolder) GetFileHandlerType() string`

GetFileHandlerType returns the FileHandlerType field if non-nil, zero value otherwise.

### GetFileHandlerTypeOk

`func (o *KalturaDropFolder) GetFileHandlerTypeOk() (*string, bool)`

GetFileHandlerTypeOk returns a tuple with the FileHandlerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHandlerType

`func (o *KalturaDropFolder) SetFileHandlerType(v string)`

SetFileHandlerType sets FileHandlerType field to given value.

### HasFileHandlerType

`func (o *KalturaDropFolder) HasFileHandlerType() bool`

HasFileHandlerType returns a boolean if a field has been set.

### GetFileNamePatterns

`func (o *KalturaDropFolder) GetFileNamePatterns() string`

GetFileNamePatterns returns the FileNamePatterns field if non-nil, zero value otherwise.

### GetFileNamePatternsOk

`func (o *KalturaDropFolder) GetFileNamePatternsOk() (*string, bool)`

GetFileNamePatternsOk returns a tuple with the FileNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePatterns

`func (o *KalturaDropFolder) SetFileNamePatterns(v string)`

SetFileNamePatterns sets FileNamePatterns field to given value.

### HasFileNamePatterns

`func (o *KalturaDropFolder) HasFileNamePatterns() bool`

HasFileNamePatterns returns a boolean if a field has been set.

### GetFileSizeCheckInterval

`func (o *KalturaDropFolder) GetFileSizeCheckInterval() int32`

GetFileSizeCheckInterval returns the FileSizeCheckInterval field if non-nil, zero value otherwise.

### GetFileSizeCheckIntervalOk

`func (o *KalturaDropFolder) GetFileSizeCheckIntervalOk() (*int32, bool)`

GetFileSizeCheckIntervalOk returns a tuple with the FileSizeCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeCheckInterval

`func (o *KalturaDropFolder) SetFileSizeCheckInterval(v int32)`

SetFileSizeCheckInterval sets FileSizeCheckInterval field to given value.

### HasFileSizeCheckInterval

`func (o *KalturaDropFolder) HasFileSizeCheckInterval() bool`

HasFileSizeCheckInterval returns a boolean if a field has been set.

### GetId

`func (o *KalturaDropFolder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDropFolder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDropFolder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDropFolder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreFileNamePatterns

`func (o *KalturaDropFolder) GetIgnoreFileNamePatterns() string`

GetIgnoreFileNamePatterns returns the IgnoreFileNamePatterns field if non-nil, zero value otherwise.

### GetIgnoreFileNamePatternsOk

`func (o *KalturaDropFolder) GetIgnoreFileNamePatternsOk() (*string, bool)`

GetIgnoreFileNamePatternsOk returns a tuple with the IgnoreFileNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreFileNamePatterns

`func (o *KalturaDropFolder) SetIgnoreFileNamePatterns(v string)`

SetIgnoreFileNamePatterns sets IgnoreFileNamePatterns field to given value.

### HasIgnoreFileNamePatterns

`func (o *KalturaDropFolder) HasIgnoreFileNamePatterns() bool`

HasIgnoreFileNamePatterns returns a boolean if a field has been set.

### GetIncremental

`func (o *KalturaDropFolder) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *KalturaDropFolder) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *KalturaDropFolder) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *KalturaDropFolder) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.

### GetLastAccessedAt

`func (o *KalturaDropFolder) GetLastAccessedAt() int32`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *KalturaDropFolder) GetLastAccessedAtOk() (*int32, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *KalturaDropFolder) SetLastAccessedAt(v int32)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *KalturaDropFolder) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### GetLastFileTimestamp

`func (o *KalturaDropFolder) GetLastFileTimestamp() int32`

GetLastFileTimestamp returns the LastFileTimestamp field if non-nil, zero value otherwise.

### GetLastFileTimestampOk

`func (o *KalturaDropFolder) GetLastFileTimestampOk() (*int32, bool)`

GetLastFileTimestampOk returns a tuple with the LastFileTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFileTimestamp

`func (o *KalturaDropFolder) SetLastFileTimestamp(v int32)`

SetLastFileTimestamp sets LastFileTimestamp field to given value.

### HasLastFileTimestamp

`func (o *KalturaDropFolder) HasLastFileTimestamp() bool`

HasLastFileTimestamp returns a boolean if a field has been set.

### GetMetadataProfileId

`func (o *KalturaDropFolder) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *KalturaDropFolder) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *KalturaDropFolder) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.

### HasMetadataProfileId

`func (o *KalturaDropFolder) HasMetadataProfileId() bool`

HasMetadataProfileId returns a boolean if a field has been set.

### GetName

`func (o *KalturaDropFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDropFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDropFolder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDropFolder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDropFolder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDropFolder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDropFolder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDropFolder) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDropFolder) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDropFolder) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDropFolder) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDropFolder) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPath

`func (o *KalturaDropFolder) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *KalturaDropFolder) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *KalturaDropFolder) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *KalturaDropFolder) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetShouldValidateKS

`func (o *KalturaDropFolder) GetShouldValidateKS() bool`

GetShouldValidateKS returns the ShouldValidateKS field if non-nil, zero value otherwise.

### GetShouldValidateKSOk

`func (o *KalturaDropFolder) GetShouldValidateKSOk() (*bool, bool)`

GetShouldValidateKSOk returns a tuple with the ShouldValidateKS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldValidateKS

`func (o *KalturaDropFolder) SetShouldValidateKS(v bool)`

SetShouldValidateKS sets ShouldValidateKS field to given value.

### HasShouldValidateKS

`func (o *KalturaDropFolder) HasShouldValidateKS() bool`

HasShouldValidateKS returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDropFolder) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDropFolder) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDropFolder) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDropFolder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaDropFolder) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaDropFolder) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaDropFolder) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaDropFolder) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *KalturaDropFolder) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaDropFolder) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaDropFolder) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaDropFolder) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDropFolder) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDropFolder) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDropFolder) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDropFolder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



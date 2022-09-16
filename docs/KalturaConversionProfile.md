# KalturaConversionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CalculateComplexity** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Should calculate file conversion complexity | [optional] 
**ClipDuration** | Pointer to **int32** | Clipping duration (in miliseconds) | [optional] 
**ClipStart** | Pointer to **int32** | Clipping start position (in miliseconds) | [optional] 
**CollectionTags** | Pointer to **string** | Defines the tags that should be used to define &#39;collective&#39;/group/multi-flavor processing,  like &#39;mbr&#39; or &#39;ism&#39; | [optional] 
**ConditionalProfiles** | Pointer to **string** | JSON string with array of \&quot;condition,profile-id\&quot; pairs. | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CropDimensions** | Pointer to [**KalturaCropDimensions**](KalturaCropDimensions.md) |  | [optional] 
**DefaultAudioLang** | Pointer to **string** | Enum Type: &#x60;KalturaLanguage&#x60; | [optional] 
**DefaultEntryId** | Pointer to **string** | ID of the default entry to be used for template data | [optional] 
**DefaultReplacementOptions** | Pointer to [**KalturaEntryReplacementOptions**](KalturaEntryReplacementOptions.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the Conversion Profile | [optional] 
**DetectGOP** | Pointer to **int32** | When set, the ExtractMedia job should detect the source file GOP using this value as the max calculated period | [optional] 
**FlavorParamsIds** | Pointer to **string** | List of included flavor ids (comma separated) | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Conversion Profile | [optional] [readonly] 
**IsDefault** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Indicates that this conversion profile is system default | [optional] 
**IsPartnerDefault** | Pointer to **bool** | &#x60;readOnly&#x60;  Indicates that this conversion profile is partner default | [optional] [readonly] 
**MediaInfoXslTransformation** | Pointer to **string** | XSL to transform ingestion Media Info XML | [optional] 
**MediaParserType** | Pointer to **string** | Enum Type: &#x60;KalturaMediaParserType&#x60;  Media parser type to be used for extract media | [optional] 
**Name** | Pointer to **string** | The name of the Conversion Profile | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **string** | Enum Type: &#x60;KalturaConversionProfileStatus&#x60; | [optional] 
**StorageProfileId** | Pointer to **int32** | ID of default storage profile to be used for linked net-storage file syncs | [optional] 
**SystemName** | Pointer to **string** | System name of the Conversion Profile | [optional] 
**Tags** | Pointer to **string** | Comma separated tags | [optional] 
**Type** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaConversionProfileType&#x60; | [optional] 
**XslTransformation** | Pointer to **string** | XSL to transform ingestion MRSS XML | [optional] 

## Methods

### NewKalturaConversionProfile

`func NewKalturaConversionProfile() *KalturaConversionProfile`

NewKalturaConversionProfile instantiates a new KalturaConversionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaConversionProfileWithDefaults

`func NewKalturaConversionProfileWithDefaults() *KalturaConversionProfile`

NewKalturaConversionProfileWithDefaults instantiates a new KalturaConversionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalculateComplexity

`func (o *KalturaConversionProfile) GetCalculateComplexity() int32`

GetCalculateComplexity returns the CalculateComplexity field if non-nil, zero value otherwise.

### GetCalculateComplexityOk

`func (o *KalturaConversionProfile) GetCalculateComplexityOk() (*int32, bool)`

GetCalculateComplexityOk returns a tuple with the CalculateComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateComplexity

`func (o *KalturaConversionProfile) SetCalculateComplexity(v int32)`

SetCalculateComplexity sets CalculateComplexity field to given value.

### HasCalculateComplexity

`func (o *KalturaConversionProfile) HasCalculateComplexity() bool`

HasCalculateComplexity returns a boolean if a field has been set.

### GetClipDuration

`func (o *KalturaConversionProfile) GetClipDuration() int32`

GetClipDuration returns the ClipDuration field if non-nil, zero value otherwise.

### GetClipDurationOk

`func (o *KalturaConversionProfile) GetClipDurationOk() (*int32, bool)`

GetClipDurationOk returns a tuple with the ClipDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipDuration

`func (o *KalturaConversionProfile) SetClipDuration(v int32)`

SetClipDuration sets ClipDuration field to given value.

### HasClipDuration

`func (o *KalturaConversionProfile) HasClipDuration() bool`

HasClipDuration returns a boolean if a field has been set.

### GetClipStart

`func (o *KalturaConversionProfile) GetClipStart() int32`

GetClipStart returns the ClipStart field if non-nil, zero value otherwise.

### GetClipStartOk

`func (o *KalturaConversionProfile) GetClipStartOk() (*int32, bool)`

GetClipStartOk returns a tuple with the ClipStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipStart

`func (o *KalturaConversionProfile) SetClipStart(v int32)`

SetClipStart sets ClipStart field to given value.

### HasClipStart

`func (o *KalturaConversionProfile) HasClipStart() bool`

HasClipStart returns a boolean if a field has been set.

### GetCollectionTags

`func (o *KalturaConversionProfile) GetCollectionTags() string`

GetCollectionTags returns the CollectionTags field if non-nil, zero value otherwise.

### GetCollectionTagsOk

`func (o *KalturaConversionProfile) GetCollectionTagsOk() (*string, bool)`

GetCollectionTagsOk returns a tuple with the CollectionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTags

`func (o *KalturaConversionProfile) SetCollectionTags(v string)`

SetCollectionTags sets CollectionTags field to given value.

### HasCollectionTags

`func (o *KalturaConversionProfile) HasCollectionTags() bool`

HasCollectionTags returns a boolean if a field has been set.

### GetConditionalProfiles

`func (o *KalturaConversionProfile) GetConditionalProfiles() string`

GetConditionalProfiles returns the ConditionalProfiles field if non-nil, zero value otherwise.

### GetConditionalProfilesOk

`func (o *KalturaConversionProfile) GetConditionalProfilesOk() (*string, bool)`

GetConditionalProfilesOk returns a tuple with the ConditionalProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalProfiles

`func (o *KalturaConversionProfile) SetConditionalProfiles(v string)`

SetConditionalProfiles sets ConditionalProfiles field to given value.

### HasConditionalProfiles

`func (o *KalturaConversionProfile) HasConditionalProfiles() bool`

HasConditionalProfiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaConversionProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaConversionProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaConversionProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaConversionProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCropDimensions

`func (o *KalturaConversionProfile) GetCropDimensions() KalturaCropDimensions`

GetCropDimensions returns the CropDimensions field if non-nil, zero value otherwise.

### GetCropDimensionsOk

`func (o *KalturaConversionProfile) GetCropDimensionsOk() (*KalturaCropDimensions, bool)`

GetCropDimensionsOk returns a tuple with the CropDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCropDimensions

`func (o *KalturaConversionProfile) SetCropDimensions(v KalturaCropDimensions)`

SetCropDimensions sets CropDimensions field to given value.

### HasCropDimensions

`func (o *KalturaConversionProfile) HasCropDimensions() bool`

HasCropDimensions returns a boolean if a field has been set.

### GetDefaultAudioLang

`func (o *KalturaConversionProfile) GetDefaultAudioLang() string`

GetDefaultAudioLang returns the DefaultAudioLang field if non-nil, zero value otherwise.

### GetDefaultAudioLangOk

`func (o *KalturaConversionProfile) GetDefaultAudioLangOk() (*string, bool)`

GetDefaultAudioLangOk returns a tuple with the DefaultAudioLang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioLang

`func (o *KalturaConversionProfile) SetDefaultAudioLang(v string)`

SetDefaultAudioLang sets DefaultAudioLang field to given value.

### HasDefaultAudioLang

`func (o *KalturaConversionProfile) HasDefaultAudioLang() bool`

HasDefaultAudioLang returns a boolean if a field has been set.

### GetDefaultEntryId

`func (o *KalturaConversionProfile) GetDefaultEntryId() string`

GetDefaultEntryId returns the DefaultEntryId field if non-nil, zero value otherwise.

### GetDefaultEntryIdOk

`func (o *KalturaConversionProfile) GetDefaultEntryIdOk() (*string, bool)`

GetDefaultEntryIdOk returns a tuple with the DefaultEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEntryId

`func (o *KalturaConversionProfile) SetDefaultEntryId(v string)`

SetDefaultEntryId sets DefaultEntryId field to given value.

### HasDefaultEntryId

`func (o *KalturaConversionProfile) HasDefaultEntryId() bool`

HasDefaultEntryId returns a boolean if a field has been set.

### GetDefaultReplacementOptions

`func (o *KalturaConversionProfile) GetDefaultReplacementOptions() KalturaEntryReplacementOptions`

GetDefaultReplacementOptions returns the DefaultReplacementOptions field if non-nil, zero value otherwise.

### GetDefaultReplacementOptionsOk

`func (o *KalturaConversionProfile) GetDefaultReplacementOptionsOk() (*KalturaEntryReplacementOptions, bool)`

GetDefaultReplacementOptionsOk returns a tuple with the DefaultReplacementOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReplacementOptions

`func (o *KalturaConversionProfile) SetDefaultReplacementOptions(v KalturaEntryReplacementOptions)`

SetDefaultReplacementOptions sets DefaultReplacementOptions field to given value.

### HasDefaultReplacementOptions

`func (o *KalturaConversionProfile) HasDefaultReplacementOptions() bool`

HasDefaultReplacementOptions returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaConversionProfile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaConversionProfile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaConversionProfile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaConversionProfile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDetectGOP

`func (o *KalturaConversionProfile) GetDetectGOP() int32`

GetDetectGOP returns the DetectGOP field if non-nil, zero value otherwise.

### GetDetectGOPOk

`func (o *KalturaConversionProfile) GetDetectGOPOk() (*int32, bool)`

GetDetectGOPOk returns a tuple with the DetectGOP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectGOP

`func (o *KalturaConversionProfile) SetDetectGOP(v int32)`

SetDetectGOP sets DetectGOP field to given value.

### HasDetectGOP

`func (o *KalturaConversionProfile) HasDetectGOP() bool`

HasDetectGOP returns a boolean if a field has been set.

### GetFlavorParamsIds

`func (o *KalturaConversionProfile) GetFlavorParamsIds() string`

GetFlavorParamsIds returns the FlavorParamsIds field if non-nil, zero value otherwise.

### GetFlavorParamsIdsOk

`func (o *KalturaConversionProfile) GetFlavorParamsIdsOk() (*string, bool)`

GetFlavorParamsIdsOk returns a tuple with the FlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsIds

`func (o *KalturaConversionProfile) SetFlavorParamsIds(v string)`

SetFlavorParamsIds sets FlavorParamsIds field to given value.

### HasFlavorParamsIds

`func (o *KalturaConversionProfile) HasFlavorParamsIds() bool`

HasFlavorParamsIds returns a boolean if a field has been set.

### GetId

`func (o *KalturaConversionProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaConversionProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaConversionProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaConversionProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *KalturaConversionProfile) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *KalturaConversionProfile) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *KalturaConversionProfile) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *KalturaConversionProfile) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPartnerDefault

`func (o *KalturaConversionProfile) GetIsPartnerDefault() bool`

GetIsPartnerDefault returns the IsPartnerDefault field if non-nil, zero value otherwise.

### GetIsPartnerDefaultOk

`func (o *KalturaConversionProfile) GetIsPartnerDefaultOk() (*bool, bool)`

GetIsPartnerDefaultOk returns a tuple with the IsPartnerDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartnerDefault

`func (o *KalturaConversionProfile) SetIsPartnerDefault(v bool)`

SetIsPartnerDefault sets IsPartnerDefault field to given value.

### HasIsPartnerDefault

`func (o *KalturaConversionProfile) HasIsPartnerDefault() bool`

HasIsPartnerDefault returns a boolean if a field has been set.

### GetMediaInfoXslTransformation

`func (o *KalturaConversionProfile) GetMediaInfoXslTransformation() string`

GetMediaInfoXslTransformation returns the MediaInfoXslTransformation field if non-nil, zero value otherwise.

### GetMediaInfoXslTransformationOk

`func (o *KalturaConversionProfile) GetMediaInfoXslTransformationOk() (*string, bool)`

GetMediaInfoXslTransformationOk returns a tuple with the MediaInfoXslTransformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInfoXslTransformation

`func (o *KalturaConversionProfile) SetMediaInfoXslTransformation(v string)`

SetMediaInfoXslTransformation sets MediaInfoXslTransformation field to given value.

### HasMediaInfoXslTransformation

`func (o *KalturaConversionProfile) HasMediaInfoXslTransformation() bool`

HasMediaInfoXslTransformation returns a boolean if a field has been set.

### GetMediaParserType

`func (o *KalturaConversionProfile) GetMediaParserType() string`

GetMediaParserType returns the MediaParserType field if non-nil, zero value otherwise.

### GetMediaParserTypeOk

`func (o *KalturaConversionProfile) GetMediaParserTypeOk() (*string, bool)`

GetMediaParserTypeOk returns a tuple with the MediaParserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaParserType

`func (o *KalturaConversionProfile) SetMediaParserType(v string)`

SetMediaParserType sets MediaParserType field to given value.

### HasMediaParserType

`func (o *KalturaConversionProfile) HasMediaParserType() bool`

HasMediaParserType returns a boolean if a field has been set.

### GetName

`func (o *KalturaConversionProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaConversionProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaConversionProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaConversionProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaConversionProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaConversionProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaConversionProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaConversionProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaConversionProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaConversionProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaConversionProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaConversionProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageProfileId

`func (o *KalturaConversionProfile) GetStorageProfileId() int32`

GetStorageProfileId returns the StorageProfileId field if non-nil, zero value otherwise.

### GetStorageProfileIdOk

`func (o *KalturaConversionProfile) GetStorageProfileIdOk() (*int32, bool)`

GetStorageProfileIdOk returns a tuple with the StorageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileId

`func (o *KalturaConversionProfile) SetStorageProfileId(v int32)`

SetStorageProfileId sets StorageProfileId field to given value.

### HasStorageProfileId

`func (o *KalturaConversionProfile) HasStorageProfileId() bool`

HasStorageProfileId returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaConversionProfile) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaConversionProfile) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaConversionProfile) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaConversionProfile) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetTags

`func (o *KalturaConversionProfile) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaConversionProfile) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaConversionProfile) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaConversionProfile) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *KalturaConversionProfile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaConversionProfile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaConversionProfile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaConversionProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetXslTransformation

`func (o *KalturaConversionProfile) GetXslTransformation() string`

GetXslTransformation returns the XslTransformation field if non-nil, zero value otherwise.

### GetXslTransformationOk

`func (o *KalturaConversionProfile) GetXslTransformationOk() (*string, bool)`

GetXslTransformationOk returns a tuple with the XslTransformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXslTransformation

`func (o *KalturaConversionProfile) SetXslTransformation(v string)`

SetXslTransformation sets XslTransformation field to given value.

### HasXslTransformation

`func (o *KalturaConversionProfile) HasXslTransformation() bool`

HasXslTransformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



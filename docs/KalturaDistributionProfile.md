# KalturaDistributionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCreateFlavors** | Pointer to **string** | Comma separated flavor params ids that should be auto converted | [optional] 
**AutoCreateThumb** | Pointer to **string** | Comma separated thumbnail params ids that should be auto generated | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Profile creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**DeleteEnabled** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProfileActionStatus&#x60; | [optional] 
**DistributeTrigger** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributeTrigger&#x60;  The event that trigger the automatic distribute | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto generated unique id | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**OptionalAssetDistributionRules** | Pointer to [**[]KalturaAssetDistributionRule**](KalturaAssetDistributionRule.md) |  | [optional] 
**OptionalFlavorParamsIds** | Pointer to **string** | Comma separated flavor params ids that should be submitted if ready | [optional] 
**OptionalThumbDimensions** | Pointer to [**[]KalturaDistributionThumbDimensions**](KalturaDistributionThumbDimensions.md) |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ProviderType** | Pointer to **string** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaDistributionProviderType&#x60; | [optional] 
**RecommendedDcForDownload** | Pointer to **int32** | The best Kaltura data center to be used to download the asset files to | [optional] 
**RecommendedDcForExecute** | Pointer to **int32** | The best Kaltura data center to be used to execute the distribution job | [optional] 
**RecommendedStorageProfileForDownload** | Pointer to **int32** | The best external storage to be used to download the asset files from | [optional] 
**ReportEnabled** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProfileActionStatus&#x60; | [optional] 
**RequiredAssetDistributionRules** | Pointer to [**[]KalturaAssetDistributionRule**](KalturaAssetDistributionRule.md) |  | [optional] 
**RequiredFlavorParamsIds** | Pointer to **string** | Comma separated flavor params ids that required to be ready before submission | [optional] 
**RequiredThumbDimensions** | Pointer to [**[]KalturaDistributionThumbDimensions**](KalturaDistributionThumbDimensions.md) |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProfileStatus&#x60; | [optional] 
**SubmitEnabled** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProfileActionStatus&#x60; | [optional] 
**SunriseDefaultOffset** | Pointer to **int32** | If entry distribution sunrise not specified that will be the default since entry creation time, in seconds | [optional] 
**SunsetDefaultOffset** | Pointer to **int32** | If entry distribution sunset not specified that will be the default since entry creation time, in seconds | [optional] 
**SupportImageEntry** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UpdateEnabled** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProfileActionStatus&#x60; | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Profile last update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaDistributionProfile

`func NewKalturaDistributionProfile() *KalturaDistributionProfile`

NewKalturaDistributionProfile instantiates a new KalturaDistributionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDistributionProfileWithDefaults

`func NewKalturaDistributionProfileWithDefaults() *KalturaDistributionProfile`

NewKalturaDistributionProfileWithDefaults instantiates a new KalturaDistributionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCreateFlavors

`func (o *KalturaDistributionProfile) GetAutoCreateFlavors() string`

GetAutoCreateFlavors returns the AutoCreateFlavors field if non-nil, zero value otherwise.

### GetAutoCreateFlavorsOk

`func (o *KalturaDistributionProfile) GetAutoCreateFlavorsOk() (*string, bool)`

GetAutoCreateFlavorsOk returns a tuple with the AutoCreateFlavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateFlavors

`func (o *KalturaDistributionProfile) SetAutoCreateFlavors(v string)`

SetAutoCreateFlavors sets AutoCreateFlavors field to given value.

### HasAutoCreateFlavors

`func (o *KalturaDistributionProfile) HasAutoCreateFlavors() bool`

HasAutoCreateFlavors returns a boolean if a field has been set.

### GetAutoCreateThumb

`func (o *KalturaDistributionProfile) GetAutoCreateThumb() string`

GetAutoCreateThumb returns the AutoCreateThumb field if non-nil, zero value otherwise.

### GetAutoCreateThumbOk

`func (o *KalturaDistributionProfile) GetAutoCreateThumbOk() (*string, bool)`

GetAutoCreateThumbOk returns a tuple with the AutoCreateThumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateThumb

`func (o *KalturaDistributionProfile) SetAutoCreateThumb(v string)`

SetAutoCreateThumb sets AutoCreateThumb field to given value.

### HasAutoCreateThumb

`func (o *KalturaDistributionProfile) HasAutoCreateThumb() bool`

HasAutoCreateThumb returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaDistributionProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDistributionProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDistributionProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDistributionProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleteEnabled

`func (o *KalturaDistributionProfile) GetDeleteEnabled() int32`

GetDeleteEnabled returns the DeleteEnabled field if non-nil, zero value otherwise.

### GetDeleteEnabledOk

`func (o *KalturaDistributionProfile) GetDeleteEnabledOk() (*int32, bool)`

GetDeleteEnabledOk returns a tuple with the DeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEnabled

`func (o *KalturaDistributionProfile) SetDeleteEnabled(v int32)`

SetDeleteEnabled sets DeleteEnabled field to given value.

### HasDeleteEnabled

`func (o *KalturaDistributionProfile) HasDeleteEnabled() bool`

HasDeleteEnabled returns a boolean if a field has been set.

### GetDistributeTrigger

`func (o *KalturaDistributionProfile) GetDistributeTrigger() int32`

GetDistributeTrigger returns the DistributeTrigger field if non-nil, zero value otherwise.

### GetDistributeTriggerOk

`func (o *KalturaDistributionProfile) GetDistributeTriggerOk() (*int32, bool)`

GetDistributeTriggerOk returns a tuple with the DistributeTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeTrigger

`func (o *KalturaDistributionProfile) SetDistributeTrigger(v int32)`

SetDistributeTrigger sets DistributeTrigger field to given value.

### HasDistributeTrigger

`func (o *KalturaDistributionProfile) HasDistributeTrigger() bool`

HasDistributeTrigger returns a boolean if a field has been set.

### GetId

`func (o *KalturaDistributionProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDistributionProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDistributionProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDistributionProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaDistributionProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaDistributionProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaDistributionProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaDistributionProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDistributionProfile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDistributionProfile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDistributionProfile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDistributionProfile) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOptionalAssetDistributionRules

`func (o *KalturaDistributionProfile) GetOptionalAssetDistributionRules() []KalturaAssetDistributionRule`

GetOptionalAssetDistributionRules returns the OptionalAssetDistributionRules field if non-nil, zero value otherwise.

### GetOptionalAssetDistributionRulesOk

`func (o *KalturaDistributionProfile) GetOptionalAssetDistributionRulesOk() (*[]KalturaAssetDistributionRule, bool)`

GetOptionalAssetDistributionRulesOk returns a tuple with the OptionalAssetDistributionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAssetDistributionRules

`func (o *KalturaDistributionProfile) SetOptionalAssetDistributionRules(v []KalturaAssetDistributionRule)`

SetOptionalAssetDistributionRules sets OptionalAssetDistributionRules field to given value.

### HasOptionalAssetDistributionRules

`func (o *KalturaDistributionProfile) HasOptionalAssetDistributionRules() bool`

HasOptionalAssetDistributionRules returns a boolean if a field has been set.

### GetOptionalFlavorParamsIds

`func (o *KalturaDistributionProfile) GetOptionalFlavorParamsIds() string`

GetOptionalFlavorParamsIds returns the OptionalFlavorParamsIds field if non-nil, zero value otherwise.

### GetOptionalFlavorParamsIdsOk

`func (o *KalturaDistributionProfile) GetOptionalFlavorParamsIdsOk() (*string, bool)`

GetOptionalFlavorParamsIdsOk returns a tuple with the OptionalFlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalFlavorParamsIds

`func (o *KalturaDistributionProfile) SetOptionalFlavorParamsIds(v string)`

SetOptionalFlavorParamsIds sets OptionalFlavorParamsIds field to given value.

### HasOptionalFlavorParamsIds

`func (o *KalturaDistributionProfile) HasOptionalFlavorParamsIds() bool`

HasOptionalFlavorParamsIds returns a boolean if a field has been set.

### GetOptionalThumbDimensions

`func (o *KalturaDistributionProfile) GetOptionalThumbDimensions() []KalturaDistributionThumbDimensions`

GetOptionalThumbDimensions returns the OptionalThumbDimensions field if non-nil, zero value otherwise.

### GetOptionalThumbDimensionsOk

`func (o *KalturaDistributionProfile) GetOptionalThumbDimensionsOk() (*[]KalturaDistributionThumbDimensions, bool)`

GetOptionalThumbDimensionsOk returns a tuple with the OptionalThumbDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalThumbDimensions

`func (o *KalturaDistributionProfile) SetOptionalThumbDimensions(v []KalturaDistributionThumbDimensions)`

SetOptionalThumbDimensions sets OptionalThumbDimensions field to given value.

### HasOptionalThumbDimensions

`func (o *KalturaDistributionProfile) HasOptionalThumbDimensions() bool`

HasOptionalThumbDimensions returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDistributionProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDistributionProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDistributionProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDistributionProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProviderType

`func (o *KalturaDistributionProfile) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *KalturaDistributionProfile) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *KalturaDistributionProfile) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *KalturaDistributionProfile) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetRecommendedDcForDownload

`func (o *KalturaDistributionProfile) GetRecommendedDcForDownload() int32`

GetRecommendedDcForDownload returns the RecommendedDcForDownload field if non-nil, zero value otherwise.

### GetRecommendedDcForDownloadOk

`func (o *KalturaDistributionProfile) GetRecommendedDcForDownloadOk() (*int32, bool)`

GetRecommendedDcForDownloadOk returns a tuple with the RecommendedDcForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDcForDownload

`func (o *KalturaDistributionProfile) SetRecommendedDcForDownload(v int32)`

SetRecommendedDcForDownload sets RecommendedDcForDownload field to given value.

### HasRecommendedDcForDownload

`func (o *KalturaDistributionProfile) HasRecommendedDcForDownload() bool`

HasRecommendedDcForDownload returns a boolean if a field has been set.

### GetRecommendedDcForExecute

`func (o *KalturaDistributionProfile) GetRecommendedDcForExecute() int32`

GetRecommendedDcForExecute returns the RecommendedDcForExecute field if non-nil, zero value otherwise.

### GetRecommendedDcForExecuteOk

`func (o *KalturaDistributionProfile) GetRecommendedDcForExecuteOk() (*int32, bool)`

GetRecommendedDcForExecuteOk returns a tuple with the RecommendedDcForExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedDcForExecute

`func (o *KalturaDistributionProfile) SetRecommendedDcForExecute(v int32)`

SetRecommendedDcForExecute sets RecommendedDcForExecute field to given value.

### HasRecommendedDcForExecute

`func (o *KalturaDistributionProfile) HasRecommendedDcForExecute() bool`

HasRecommendedDcForExecute returns a boolean if a field has been set.

### GetRecommendedStorageProfileForDownload

`func (o *KalturaDistributionProfile) GetRecommendedStorageProfileForDownload() int32`

GetRecommendedStorageProfileForDownload returns the RecommendedStorageProfileForDownload field if non-nil, zero value otherwise.

### GetRecommendedStorageProfileForDownloadOk

`func (o *KalturaDistributionProfile) GetRecommendedStorageProfileForDownloadOk() (*int32, bool)`

GetRecommendedStorageProfileForDownloadOk returns a tuple with the RecommendedStorageProfileForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedStorageProfileForDownload

`func (o *KalturaDistributionProfile) SetRecommendedStorageProfileForDownload(v int32)`

SetRecommendedStorageProfileForDownload sets RecommendedStorageProfileForDownload field to given value.

### HasRecommendedStorageProfileForDownload

`func (o *KalturaDistributionProfile) HasRecommendedStorageProfileForDownload() bool`

HasRecommendedStorageProfileForDownload returns a boolean if a field has been set.

### GetReportEnabled

`func (o *KalturaDistributionProfile) GetReportEnabled() int32`

GetReportEnabled returns the ReportEnabled field if non-nil, zero value otherwise.

### GetReportEnabledOk

`func (o *KalturaDistributionProfile) GetReportEnabledOk() (*int32, bool)`

GetReportEnabledOk returns a tuple with the ReportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEnabled

`func (o *KalturaDistributionProfile) SetReportEnabled(v int32)`

SetReportEnabled sets ReportEnabled field to given value.

### HasReportEnabled

`func (o *KalturaDistributionProfile) HasReportEnabled() bool`

HasReportEnabled returns a boolean if a field has been set.

### GetRequiredAssetDistributionRules

`func (o *KalturaDistributionProfile) GetRequiredAssetDistributionRules() []KalturaAssetDistributionRule`

GetRequiredAssetDistributionRules returns the RequiredAssetDistributionRules field if non-nil, zero value otherwise.

### GetRequiredAssetDistributionRulesOk

`func (o *KalturaDistributionProfile) GetRequiredAssetDistributionRulesOk() (*[]KalturaAssetDistributionRule, bool)`

GetRequiredAssetDistributionRulesOk returns a tuple with the RequiredAssetDistributionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAssetDistributionRules

`func (o *KalturaDistributionProfile) SetRequiredAssetDistributionRules(v []KalturaAssetDistributionRule)`

SetRequiredAssetDistributionRules sets RequiredAssetDistributionRules field to given value.

### HasRequiredAssetDistributionRules

`func (o *KalturaDistributionProfile) HasRequiredAssetDistributionRules() bool`

HasRequiredAssetDistributionRules returns a boolean if a field has been set.

### GetRequiredFlavorParamsIds

`func (o *KalturaDistributionProfile) GetRequiredFlavorParamsIds() string`

GetRequiredFlavorParamsIds returns the RequiredFlavorParamsIds field if non-nil, zero value otherwise.

### GetRequiredFlavorParamsIdsOk

`func (o *KalturaDistributionProfile) GetRequiredFlavorParamsIdsOk() (*string, bool)`

GetRequiredFlavorParamsIdsOk returns a tuple with the RequiredFlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredFlavorParamsIds

`func (o *KalturaDistributionProfile) SetRequiredFlavorParamsIds(v string)`

SetRequiredFlavorParamsIds sets RequiredFlavorParamsIds field to given value.

### HasRequiredFlavorParamsIds

`func (o *KalturaDistributionProfile) HasRequiredFlavorParamsIds() bool`

HasRequiredFlavorParamsIds returns a boolean if a field has been set.

### GetRequiredThumbDimensions

`func (o *KalturaDistributionProfile) GetRequiredThumbDimensions() []KalturaDistributionThumbDimensions`

GetRequiredThumbDimensions returns the RequiredThumbDimensions field if non-nil, zero value otherwise.

### GetRequiredThumbDimensionsOk

`func (o *KalturaDistributionProfile) GetRequiredThumbDimensionsOk() (*[]KalturaDistributionThumbDimensions, bool)`

GetRequiredThumbDimensionsOk returns a tuple with the RequiredThumbDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredThumbDimensions

`func (o *KalturaDistributionProfile) SetRequiredThumbDimensions(v []KalturaDistributionThumbDimensions)`

SetRequiredThumbDimensions sets RequiredThumbDimensions field to given value.

### HasRequiredThumbDimensions

`func (o *KalturaDistributionProfile) HasRequiredThumbDimensions() bool`

HasRequiredThumbDimensions returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDistributionProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDistributionProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDistributionProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDistributionProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmitEnabled

`func (o *KalturaDistributionProfile) GetSubmitEnabled() int32`

GetSubmitEnabled returns the SubmitEnabled field if non-nil, zero value otherwise.

### GetSubmitEnabledOk

`func (o *KalturaDistributionProfile) GetSubmitEnabledOk() (*int32, bool)`

GetSubmitEnabledOk returns a tuple with the SubmitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitEnabled

`func (o *KalturaDistributionProfile) SetSubmitEnabled(v int32)`

SetSubmitEnabled sets SubmitEnabled field to given value.

### HasSubmitEnabled

`func (o *KalturaDistributionProfile) HasSubmitEnabled() bool`

HasSubmitEnabled returns a boolean if a field has been set.

### GetSunriseDefaultOffset

`func (o *KalturaDistributionProfile) GetSunriseDefaultOffset() int32`

GetSunriseDefaultOffset returns the SunriseDefaultOffset field if non-nil, zero value otherwise.

### GetSunriseDefaultOffsetOk

`func (o *KalturaDistributionProfile) GetSunriseDefaultOffsetOk() (*int32, bool)`

GetSunriseDefaultOffsetOk returns a tuple with the SunriseDefaultOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunriseDefaultOffset

`func (o *KalturaDistributionProfile) SetSunriseDefaultOffset(v int32)`

SetSunriseDefaultOffset sets SunriseDefaultOffset field to given value.

### HasSunriseDefaultOffset

`func (o *KalturaDistributionProfile) HasSunriseDefaultOffset() bool`

HasSunriseDefaultOffset returns a boolean if a field has been set.

### GetSunsetDefaultOffset

`func (o *KalturaDistributionProfile) GetSunsetDefaultOffset() int32`

GetSunsetDefaultOffset returns the SunsetDefaultOffset field if non-nil, zero value otherwise.

### GetSunsetDefaultOffsetOk

`func (o *KalturaDistributionProfile) GetSunsetDefaultOffsetOk() (*int32, bool)`

GetSunsetDefaultOffsetOk returns a tuple with the SunsetDefaultOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunsetDefaultOffset

`func (o *KalturaDistributionProfile) SetSunsetDefaultOffset(v int32)`

SetSunsetDefaultOffset sets SunsetDefaultOffset field to given value.

### HasSunsetDefaultOffset

`func (o *KalturaDistributionProfile) HasSunsetDefaultOffset() bool`

HasSunsetDefaultOffset returns a boolean if a field has been set.

### GetSupportImageEntry

`func (o *KalturaDistributionProfile) GetSupportImageEntry() bool`

GetSupportImageEntry returns the SupportImageEntry field if non-nil, zero value otherwise.

### GetSupportImageEntryOk

`func (o *KalturaDistributionProfile) GetSupportImageEntryOk() (*bool, bool)`

GetSupportImageEntryOk returns a tuple with the SupportImageEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportImageEntry

`func (o *KalturaDistributionProfile) SetSupportImageEntry(v bool)`

SetSupportImageEntry sets SupportImageEntry field to given value.

### HasSupportImageEntry

`func (o *KalturaDistributionProfile) HasSupportImageEntry() bool`

HasSupportImageEntry returns a boolean if a field has been set.

### GetUpdateEnabled

`func (o *KalturaDistributionProfile) GetUpdateEnabled() int32`

GetUpdateEnabled returns the UpdateEnabled field if non-nil, zero value otherwise.

### GetUpdateEnabledOk

`func (o *KalturaDistributionProfile) GetUpdateEnabledOk() (*int32, bool)`

GetUpdateEnabledOk returns a tuple with the UpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateEnabled

`func (o *KalturaDistributionProfile) SetUpdateEnabled(v int32)`

SetUpdateEnabled sets UpdateEnabled field to given value.

### HasUpdateEnabled

`func (o *KalturaDistributionProfile) HasUpdateEnabled() bool`

HasUpdateEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDistributionProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDistributionProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDistributionProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDistributionProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



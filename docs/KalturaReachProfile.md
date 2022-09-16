# KalturaReachProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDisplayHumanCaptionsOnPlayer** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**AutoDisplayMachineCaptionsOnPlayer** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**ContentDeletionPolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaReachProfileContentDeletionPolicy&#x60; | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Credit** | Pointer to [**KalturaBaseVendorCredit**](KalturaBaseVendorCredit.md) |  | [optional] 
**DefaultOutputFormat** | Pointer to **int32** | Enum Type: &#x60;KalturaVendorCatalogItemOutputFormat&#x60; | [optional] 
**Dictionaries** | Pointer to [**[]KalturaDictionary**](KalturaDictionary.md) |  | [optional] 
**EnableAudioTags** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**EnableHumanModeration** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**EnableMachineModeration** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**EnableMetadataExtraction** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**EnableProfanityRemoval** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**EnableSpeakerChangeIndication** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**FlavorParamsIds** | Pointer to **string** | Comma separated flavorParamsIds that the vendor should look for it matching asset when trying to download the asset | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LabelAdditionForHumanServiceType** | Pointer to **string** |  | [optional] 
**LabelAdditionForMachineServiceType** | Pointer to **string** |  | [optional] 
**MaxCharactersPerCaptionLine** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** | The name of the profile | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ProfileType** | Pointer to **int32** | Enum Type: &#x60;KalturaReachProfileType&#x60; | [optional] 
**Rules** | Pointer to [**[]KalturaRule**](KalturaRule.md) |  | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaReachProfileStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UsedCredit** | Pointer to **float32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**VendorTaskProcessingRegion** | Pointer to **int32** | Enum Type: &#x60;KalturaVendorTaskProcessingRegion&#x60;  Indicates in which region the task processing should task place | [optional] 

## Methods

### NewKalturaReachProfile

`func NewKalturaReachProfile() *KalturaReachProfile`

NewKalturaReachProfile instantiates a new KalturaReachProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReachProfileWithDefaults

`func NewKalturaReachProfileWithDefaults() *KalturaReachProfile`

NewKalturaReachProfileWithDefaults instantiates a new KalturaReachProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDisplayHumanCaptionsOnPlayer

`func (o *KalturaReachProfile) GetAutoDisplayHumanCaptionsOnPlayer() int32`

GetAutoDisplayHumanCaptionsOnPlayer returns the AutoDisplayHumanCaptionsOnPlayer field if non-nil, zero value otherwise.

### GetAutoDisplayHumanCaptionsOnPlayerOk

`func (o *KalturaReachProfile) GetAutoDisplayHumanCaptionsOnPlayerOk() (*int32, bool)`

GetAutoDisplayHumanCaptionsOnPlayerOk returns a tuple with the AutoDisplayHumanCaptionsOnPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDisplayHumanCaptionsOnPlayer

`func (o *KalturaReachProfile) SetAutoDisplayHumanCaptionsOnPlayer(v int32)`

SetAutoDisplayHumanCaptionsOnPlayer sets AutoDisplayHumanCaptionsOnPlayer field to given value.

### HasAutoDisplayHumanCaptionsOnPlayer

`func (o *KalturaReachProfile) HasAutoDisplayHumanCaptionsOnPlayer() bool`

HasAutoDisplayHumanCaptionsOnPlayer returns a boolean if a field has been set.

### GetAutoDisplayMachineCaptionsOnPlayer

`func (o *KalturaReachProfile) GetAutoDisplayMachineCaptionsOnPlayer() int32`

GetAutoDisplayMachineCaptionsOnPlayer returns the AutoDisplayMachineCaptionsOnPlayer field if non-nil, zero value otherwise.

### GetAutoDisplayMachineCaptionsOnPlayerOk

`func (o *KalturaReachProfile) GetAutoDisplayMachineCaptionsOnPlayerOk() (*int32, bool)`

GetAutoDisplayMachineCaptionsOnPlayerOk returns a tuple with the AutoDisplayMachineCaptionsOnPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDisplayMachineCaptionsOnPlayer

`func (o *KalturaReachProfile) SetAutoDisplayMachineCaptionsOnPlayer(v int32)`

SetAutoDisplayMachineCaptionsOnPlayer sets AutoDisplayMachineCaptionsOnPlayer field to given value.

### HasAutoDisplayMachineCaptionsOnPlayer

`func (o *KalturaReachProfile) HasAutoDisplayMachineCaptionsOnPlayer() bool`

HasAutoDisplayMachineCaptionsOnPlayer returns a boolean if a field has been set.

### GetContentDeletionPolicy

`func (o *KalturaReachProfile) GetContentDeletionPolicy() int32`

GetContentDeletionPolicy returns the ContentDeletionPolicy field if non-nil, zero value otherwise.

### GetContentDeletionPolicyOk

`func (o *KalturaReachProfile) GetContentDeletionPolicyOk() (*int32, bool)`

GetContentDeletionPolicyOk returns a tuple with the ContentDeletionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDeletionPolicy

`func (o *KalturaReachProfile) SetContentDeletionPolicy(v int32)`

SetContentDeletionPolicy sets ContentDeletionPolicy field to given value.

### HasContentDeletionPolicy

`func (o *KalturaReachProfile) HasContentDeletionPolicy() bool`

HasContentDeletionPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaReachProfile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaReachProfile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaReachProfile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaReachProfile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredit

`func (o *KalturaReachProfile) GetCredit() KalturaBaseVendorCredit`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *KalturaReachProfile) GetCreditOk() (*KalturaBaseVendorCredit, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *KalturaReachProfile) SetCredit(v KalturaBaseVendorCredit)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *KalturaReachProfile) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetDefaultOutputFormat

`func (o *KalturaReachProfile) GetDefaultOutputFormat() int32`

GetDefaultOutputFormat returns the DefaultOutputFormat field if non-nil, zero value otherwise.

### GetDefaultOutputFormatOk

`func (o *KalturaReachProfile) GetDefaultOutputFormatOk() (*int32, bool)`

GetDefaultOutputFormatOk returns a tuple with the DefaultOutputFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOutputFormat

`func (o *KalturaReachProfile) SetDefaultOutputFormat(v int32)`

SetDefaultOutputFormat sets DefaultOutputFormat field to given value.

### HasDefaultOutputFormat

`func (o *KalturaReachProfile) HasDefaultOutputFormat() bool`

HasDefaultOutputFormat returns a boolean if a field has been set.

### GetDictionaries

`func (o *KalturaReachProfile) GetDictionaries() []KalturaDictionary`

GetDictionaries returns the Dictionaries field if non-nil, zero value otherwise.

### GetDictionariesOk

`func (o *KalturaReachProfile) GetDictionariesOk() (*[]KalturaDictionary, bool)`

GetDictionariesOk returns a tuple with the Dictionaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDictionaries

`func (o *KalturaReachProfile) SetDictionaries(v []KalturaDictionary)`

SetDictionaries sets Dictionaries field to given value.

### HasDictionaries

`func (o *KalturaReachProfile) HasDictionaries() bool`

HasDictionaries returns a boolean if a field has been set.

### GetEnableAudioTags

`func (o *KalturaReachProfile) GetEnableAudioTags() int32`

GetEnableAudioTags returns the EnableAudioTags field if non-nil, zero value otherwise.

### GetEnableAudioTagsOk

`func (o *KalturaReachProfile) GetEnableAudioTagsOk() (*int32, bool)`

GetEnableAudioTagsOk returns a tuple with the EnableAudioTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioTags

`func (o *KalturaReachProfile) SetEnableAudioTags(v int32)`

SetEnableAudioTags sets EnableAudioTags field to given value.

### HasEnableAudioTags

`func (o *KalturaReachProfile) HasEnableAudioTags() bool`

HasEnableAudioTags returns a boolean if a field has been set.

### GetEnableHumanModeration

`func (o *KalturaReachProfile) GetEnableHumanModeration() int32`

GetEnableHumanModeration returns the EnableHumanModeration field if non-nil, zero value otherwise.

### GetEnableHumanModerationOk

`func (o *KalturaReachProfile) GetEnableHumanModerationOk() (*int32, bool)`

GetEnableHumanModerationOk returns a tuple with the EnableHumanModeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHumanModeration

`func (o *KalturaReachProfile) SetEnableHumanModeration(v int32)`

SetEnableHumanModeration sets EnableHumanModeration field to given value.

### HasEnableHumanModeration

`func (o *KalturaReachProfile) HasEnableHumanModeration() bool`

HasEnableHumanModeration returns a boolean if a field has been set.

### GetEnableMachineModeration

`func (o *KalturaReachProfile) GetEnableMachineModeration() int32`

GetEnableMachineModeration returns the EnableMachineModeration field if non-nil, zero value otherwise.

### GetEnableMachineModerationOk

`func (o *KalturaReachProfile) GetEnableMachineModerationOk() (*int32, bool)`

GetEnableMachineModerationOk returns a tuple with the EnableMachineModeration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMachineModeration

`func (o *KalturaReachProfile) SetEnableMachineModeration(v int32)`

SetEnableMachineModeration sets EnableMachineModeration field to given value.

### HasEnableMachineModeration

`func (o *KalturaReachProfile) HasEnableMachineModeration() bool`

HasEnableMachineModeration returns a boolean if a field has been set.

### GetEnableMetadataExtraction

`func (o *KalturaReachProfile) GetEnableMetadataExtraction() int32`

GetEnableMetadataExtraction returns the EnableMetadataExtraction field if non-nil, zero value otherwise.

### GetEnableMetadataExtractionOk

`func (o *KalturaReachProfile) GetEnableMetadataExtractionOk() (*int32, bool)`

GetEnableMetadataExtractionOk returns a tuple with the EnableMetadataExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetadataExtraction

`func (o *KalturaReachProfile) SetEnableMetadataExtraction(v int32)`

SetEnableMetadataExtraction sets EnableMetadataExtraction field to given value.

### HasEnableMetadataExtraction

`func (o *KalturaReachProfile) HasEnableMetadataExtraction() bool`

HasEnableMetadataExtraction returns a boolean if a field has been set.

### GetEnableProfanityRemoval

`func (o *KalturaReachProfile) GetEnableProfanityRemoval() int32`

GetEnableProfanityRemoval returns the EnableProfanityRemoval field if non-nil, zero value otherwise.

### GetEnableProfanityRemovalOk

`func (o *KalturaReachProfile) GetEnableProfanityRemovalOk() (*int32, bool)`

GetEnableProfanityRemovalOk returns a tuple with the EnableProfanityRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfanityRemoval

`func (o *KalturaReachProfile) SetEnableProfanityRemoval(v int32)`

SetEnableProfanityRemoval sets EnableProfanityRemoval field to given value.

### HasEnableProfanityRemoval

`func (o *KalturaReachProfile) HasEnableProfanityRemoval() bool`

HasEnableProfanityRemoval returns a boolean if a field has been set.

### GetEnableSpeakerChangeIndication

`func (o *KalturaReachProfile) GetEnableSpeakerChangeIndication() int32`

GetEnableSpeakerChangeIndication returns the EnableSpeakerChangeIndication field if non-nil, zero value otherwise.

### GetEnableSpeakerChangeIndicationOk

`func (o *KalturaReachProfile) GetEnableSpeakerChangeIndicationOk() (*int32, bool)`

GetEnableSpeakerChangeIndicationOk returns a tuple with the EnableSpeakerChangeIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSpeakerChangeIndication

`func (o *KalturaReachProfile) SetEnableSpeakerChangeIndication(v int32)`

SetEnableSpeakerChangeIndication sets EnableSpeakerChangeIndication field to given value.

### HasEnableSpeakerChangeIndication

`func (o *KalturaReachProfile) HasEnableSpeakerChangeIndication() bool`

HasEnableSpeakerChangeIndication returns a boolean if a field has been set.

### GetFlavorParamsIds

`func (o *KalturaReachProfile) GetFlavorParamsIds() string`

GetFlavorParamsIds returns the FlavorParamsIds field if non-nil, zero value otherwise.

### GetFlavorParamsIdsOk

`func (o *KalturaReachProfile) GetFlavorParamsIdsOk() (*string, bool)`

GetFlavorParamsIdsOk returns a tuple with the FlavorParamsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorParamsIds

`func (o *KalturaReachProfile) SetFlavorParamsIds(v string)`

SetFlavorParamsIds sets FlavorParamsIds field to given value.

### HasFlavorParamsIds

`func (o *KalturaReachProfile) HasFlavorParamsIds() bool`

HasFlavorParamsIds returns a boolean if a field has been set.

### GetId

`func (o *KalturaReachProfile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaReachProfile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaReachProfile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaReachProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabelAdditionForHumanServiceType

`func (o *KalturaReachProfile) GetLabelAdditionForHumanServiceType() string`

GetLabelAdditionForHumanServiceType returns the LabelAdditionForHumanServiceType field if non-nil, zero value otherwise.

### GetLabelAdditionForHumanServiceTypeOk

`func (o *KalturaReachProfile) GetLabelAdditionForHumanServiceTypeOk() (*string, bool)`

GetLabelAdditionForHumanServiceTypeOk returns a tuple with the LabelAdditionForHumanServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelAdditionForHumanServiceType

`func (o *KalturaReachProfile) SetLabelAdditionForHumanServiceType(v string)`

SetLabelAdditionForHumanServiceType sets LabelAdditionForHumanServiceType field to given value.

### HasLabelAdditionForHumanServiceType

`func (o *KalturaReachProfile) HasLabelAdditionForHumanServiceType() bool`

HasLabelAdditionForHumanServiceType returns a boolean if a field has been set.

### GetLabelAdditionForMachineServiceType

`func (o *KalturaReachProfile) GetLabelAdditionForMachineServiceType() string`

GetLabelAdditionForMachineServiceType returns the LabelAdditionForMachineServiceType field if non-nil, zero value otherwise.

### GetLabelAdditionForMachineServiceTypeOk

`func (o *KalturaReachProfile) GetLabelAdditionForMachineServiceTypeOk() (*string, bool)`

GetLabelAdditionForMachineServiceTypeOk returns a tuple with the LabelAdditionForMachineServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelAdditionForMachineServiceType

`func (o *KalturaReachProfile) SetLabelAdditionForMachineServiceType(v string)`

SetLabelAdditionForMachineServiceType sets LabelAdditionForMachineServiceType field to given value.

### HasLabelAdditionForMachineServiceType

`func (o *KalturaReachProfile) HasLabelAdditionForMachineServiceType() bool`

HasLabelAdditionForMachineServiceType returns a boolean if a field has been set.

### GetMaxCharactersPerCaptionLine

`func (o *KalturaReachProfile) GetMaxCharactersPerCaptionLine() int32`

GetMaxCharactersPerCaptionLine returns the MaxCharactersPerCaptionLine field if non-nil, zero value otherwise.

### GetMaxCharactersPerCaptionLineOk

`func (o *KalturaReachProfile) GetMaxCharactersPerCaptionLineOk() (*int32, bool)`

GetMaxCharactersPerCaptionLineOk returns a tuple with the MaxCharactersPerCaptionLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCharactersPerCaptionLine

`func (o *KalturaReachProfile) SetMaxCharactersPerCaptionLine(v int32)`

SetMaxCharactersPerCaptionLine sets MaxCharactersPerCaptionLine field to given value.

### HasMaxCharactersPerCaptionLine

`func (o *KalturaReachProfile) HasMaxCharactersPerCaptionLine() bool`

HasMaxCharactersPerCaptionLine returns a boolean if a field has been set.

### GetName

`func (o *KalturaReachProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaReachProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaReachProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaReachProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaReachProfile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaReachProfile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaReachProfile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaReachProfile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProfileType

`func (o *KalturaReachProfile) GetProfileType() int32`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *KalturaReachProfile) GetProfileTypeOk() (*int32, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *KalturaReachProfile) SetProfileType(v int32)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *KalturaReachProfile) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### GetRules

`func (o *KalturaReachProfile) GetRules() []KalturaRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *KalturaReachProfile) GetRulesOk() (*[]KalturaRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *KalturaReachProfile) SetRules(v []KalturaRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *KalturaReachProfile) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaReachProfile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaReachProfile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaReachProfile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaReachProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaReachProfile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaReachProfile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaReachProfile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaReachProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsedCredit

`func (o *KalturaReachProfile) GetUsedCredit() float32`

GetUsedCredit returns the UsedCredit field if non-nil, zero value otherwise.

### GetUsedCreditOk

`func (o *KalturaReachProfile) GetUsedCreditOk() (*float32, bool)`

GetUsedCreditOk returns a tuple with the UsedCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCredit

`func (o *KalturaReachProfile) SetUsedCredit(v float32)`

SetUsedCredit sets UsedCredit field to given value.

### HasUsedCredit

`func (o *KalturaReachProfile) HasUsedCredit() bool`

HasUsedCredit returns a boolean if a field has been set.

### GetVendorTaskProcessingRegion

`func (o *KalturaReachProfile) GetVendorTaskProcessingRegion() int32`

GetVendorTaskProcessingRegion returns the VendorTaskProcessingRegion field if non-nil, zero value otherwise.

### GetVendorTaskProcessingRegionOk

`func (o *KalturaReachProfile) GetVendorTaskProcessingRegionOk() (*int32, bool)`

GetVendorTaskProcessingRegionOk returns a tuple with the VendorTaskProcessingRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorTaskProcessingRegion

`func (o *KalturaReachProfile) SetVendorTaskProcessingRegion(v int32)`

SetVendorTaskProcessingRegion sets VendorTaskProcessingRegion field to given value.

### HasVendorTaskProcessingRegion

`func (o *KalturaReachProfile) HasVendorTaskProcessingRegion() bool`

HasVendorTaskProcessingRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



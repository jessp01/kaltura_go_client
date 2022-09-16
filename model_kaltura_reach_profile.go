/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KalturaReachProfile struct for KalturaReachProfile
type KalturaReachProfile struct {
	// Enum Type: `KalturaNullableBoolean`
	AutoDisplayHumanCaptionsOnPlayer *int32 `json:"autoDisplayHumanCaptionsOnPlayer,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	AutoDisplayMachineCaptionsOnPlayer *int32 `json:"autoDisplayMachineCaptionsOnPlayer,omitempty"`
	// Enum Type: `KalturaReachProfileContentDeletionPolicy`
	ContentDeletionPolicy *int32 `json:"contentDeletionPolicy,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	Credit *KalturaBaseVendorCredit `json:"credit,omitempty"`
	// Enum Type: `KalturaVendorCatalogItemOutputFormat`
	DefaultOutputFormat *int32 `json:"defaultOutputFormat,omitempty"`
	Dictionaries []KalturaDictionary `json:"dictionaries,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableAudioTags *int32 `json:"enableAudioTags,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableHumanModeration *int32 `json:"enableHumanModeration,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableMachineModeration *int32 `json:"enableMachineModeration,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableMetadataExtraction *int32 `json:"enableMetadataExtraction,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableProfanityRemoval *int32 `json:"enableProfanityRemoval,omitempty"`
	// Enum Type: `KalturaNullableBoolean`
	EnableSpeakerChangeIndication *int32 `json:"enableSpeakerChangeIndication,omitempty"`
	// Comma separated flavorParamsIds that the vendor should look for it matching asset when trying to download the asset
	FlavorParamsIds *string `json:"flavorParamsIds,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	LabelAdditionForHumanServiceType *string `json:"labelAdditionForHumanServiceType,omitempty"`
	LabelAdditionForMachineServiceType *string `json:"labelAdditionForMachineServiceType,omitempty"`
	MaxCharactersPerCaptionLine *int32 `json:"maxCharactersPerCaptionLine,omitempty"`
	// The name of the profile
	Name *string `json:"name,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// Enum Type: `KalturaReachProfileType`
	ProfileType *int32 `json:"profileType,omitempty"`
	Rules []KalturaRule `json:"rules,omitempty"`
	// `readOnly`  Enum Type: `KalturaReachProfileStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`
	UsedCredit *float32 `json:"usedCredit,omitempty"`
	// Enum Type: `KalturaVendorTaskProcessingRegion`  Indicates in which region the task processing should task place
	VendorTaskProcessingRegion *int32 `json:"vendorTaskProcessingRegion,omitempty"`
}

// NewKalturaReachProfile instantiates a new KalturaReachProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReachProfile() *KalturaReachProfile {
	this := KalturaReachProfile{}
	return &this
}

// NewKalturaReachProfileWithDefaults instantiates a new KalturaReachProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReachProfileWithDefaults() *KalturaReachProfile {
	this := KalturaReachProfile{}
	return &this
}

// GetAutoDisplayHumanCaptionsOnPlayer returns the AutoDisplayHumanCaptionsOnPlayer field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetAutoDisplayHumanCaptionsOnPlayer() int32 {
	if o == nil || o.AutoDisplayHumanCaptionsOnPlayer == nil {
		var ret int32
		return ret
	}
	return *o.AutoDisplayHumanCaptionsOnPlayer
}

// GetAutoDisplayHumanCaptionsOnPlayerOk returns a tuple with the AutoDisplayHumanCaptionsOnPlayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetAutoDisplayHumanCaptionsOnPlayerOk() (*int32, bool) {
	if o == nil || o.AutoDisplayHumanCaptionsOnPlayer == nil {
		return nil, false
	}
	return o.AutoDisplayHumanCaptionsOnPlayer, true
}

// HasAutoDisplayHumanCaptionsOnPlayer returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasAutoDisplayHumanCaptionsOnPlayer() bool {
	if o != nil && o.AutoDisplayHumanCaptionsOnPlayer != nil {
		return true
	}

	return false
}

// SetAutoDisplayHumanCaptionsOnPlayer gets a reference to the given int32 and assigns it to the AutoDisplayHumanCaptionsOnPlayer field.
func (o *KalturaReachProfile) SetAutoDisplayHumanCaptionsOnPlayer(v int32) {
	o.AutoDisplayHumanCaptionsOnPlayer = &v
}

// GetAutoDisplayMachineCaptionsOnPlayer returns the AutoDisplayMachineCaptionsOnPlayer field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetAutoDisplayMachineCaptionsOnPlayer() int32 {
	if o == nil || o.AutoDisplayMachineCaptionsOnPlayer == nil {
		var ret int32
		return ret
	}
	return *o.AutoDisplayMachineCaptionsOnPlayer
}

// GetAutoDisplayMachineCaptionsOnPlayerOk returns a tuple with the AutoDisplayMachineCaptionsOnPlayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetAutoDisplayMachineCaptionsOnPlayerOk() (*int32, bool) {
	if o == nil || o.AutoDisplayMachineCaptionsOnPlayer == nil {
		return nil, false
	}
	return o.AutoDisplayMachineCaptionsOnPlayer, true
}

// HasAutoDisplayMachineCaptionsOnPlayer returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasAutoDisplayMachineCaptionsOnPlayer() bool {
	if o != nil && o.AutoDisplayMachineCaptionsOnPlayer != nil {
		return true
	}

	return false
}

// SetAutoDisplayMachineCaptionsOnPlayer gets a reference to the given int32 and assigns it to the AutoDisplayMachineCaptionsOnPlayer field.
func (o *KalturaReachProfile) SetAutoDisplayMachineCaptionsOnPlayer(v int32) {
	o.AutoDisplayMachineCaptionsOnPlayer = &v
}

// GetContentDeletionPolicy returns the ContentDeletionPolicy field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetContentDeletionPolicy() int32 {
	if o == nil || o.ContentDeletionPolicy == nil {
		var ret int32
		return ret
	}
	return *o.ContentDeletionPolicy
}

// GetContentDeletionPolicyOk returns a tuple with the ContentDeletionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetContentDeletionPolicyOk() (*int32, bool) {
	if o == nil || o.ContentDeletionPolicy == nil {
		return nil, false
	}
	return o.ContentDeletionPolicy, true
}

// HasContentDeletionPolicy returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasContentDeletionPolicy() bool {
	if o != nil && o.ContentDeletionPolicy != nil {
		return true
	}

	return false
}

// SetContentDeletionPolicy gets a reference to the given int32 and assigns it to the ContentDeletionPolicy field.
func (o *KalturaReachProfile) SetContentDeletionPolicy(v int32) {
	o.ContentDeletionPolicy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaReachProfile) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetCredit() KalturaBaseVendorCredit {
	if o == nil || o.Credit == nil {
		var ret KalturaBaseVendorCredit
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetCreditOk() (*KalturaBaseVendorCredit, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given KalturaBaseVendorCredit and assigns it to the Credit field.
func (o *KalturaReachProfile) SetCredit(v KalturaBaseVendorCredit) {
	o.Credit = &v
}

// GetDefaultOutputFormat returns the DefaultOutputFormat field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetDefaultOutputFormat() int32 {
	if o == nil || o.DefaultOutputFormat == nil {
		var ret int32
		return ret
	}
	return *o.DefaultOutputFormat
}

// GetDefaultOutputFormatOk returns a tuple with the DefaultOutputFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetDefaultOutputFormatOk() (*int32, bool) {
	if o == nil || o.DefaultOutputFormat == nil {
		return nil, false
	}
	return o.DefaultOutputFormat, true
}

// HasDefaultOutputFormat returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasDefaultOutputFormat() bool {
	if o != nil && o.DefaultOutputFormat != nil {
		return true
	}

	return false
}

// SetDefaultOutputFormat gets a reference to the given int32 and assigns it to the DefaultOutputFormat field.
func (o *KalturaReachProfile) SetDefaultOutputFormat(v int32) {
	o.DefaultOutputFormat = &v
}

// GetDictionaries returns the Dictionaries field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetDictionaries() []KalturaDictionary {
	if o == nil || o.Dictionaries == nil {
		var ret []KalturaDictionary
		return ret
	}
	return o.Dictionaries
}

// GetDictionariesOk returns a tuple with the Dictionaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetDictionariesOk() ([]KalturaDictionary, bool) {
	if o == nil || o.Dictionaries == nil {
		return nil, false
	}
	return o.Dictionaries, true
}

// HasDictionaries returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasDictionaries() bool {
	if o != nil && o.Dictionaries != nil {
		return true
	}

	return false
}

// SetDictionaries gets a reference to the given []KalturaDictionary and assigns it to the Dictionaries field.
func (o *KalturaReachProfile) SetDictionaries(v []KalturaDictionary) {
	o.Dictionaries = v
}

// GetEnableAudioTags returns the EnableAudioTags field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableAudioTags() int32 {
	if o == nil || o.EnableAudioTags == nil {
		var ret int32
		return ret
	}
	return *o.EnableAudioTags
}

// GetEnableAudioTagsOk returns a tuple with the EnableAudioTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableAudioTagsOk() (*int32, bool) {
	if o == nil || o.EnableAudioTags == nil {
		return nil, false
	}
	return o.EnableAudioTags, true
}

// HasEnableAudioTags returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableAudioTags() bool {
	if o != nil && o.EnableAudioTags != nil {
		return true
	}

	return false
}

// SetEnableAudioTags gets a reference to the given int32 and assigns it to the EnableAudioTags field.
func (o *KalturaReachProfile) SetEnableAudioTags(v int32) {
	o.EnableAudioTags = &v
}

// GetEnableHumanModeration returns the EnableHumanModeration field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableHumanModeration() int32 {
	if o == nil || o.EnableHumanModeration == nil {
		var ret int32
		return ret
	}
	return *o.EnableHumanModeration
}

// GetEnableHumanModerationOk returns a tuple with the EnableHumanModeration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableHumanModerationOk() (*int32, bool) {
	if o == nil || o.EnableHumanModeration == nil {
		return nil, false
	}
	return o.EnableHumanModeration, true
}

// HasEnableHumanModeration returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableHumanModeration() bool {
	if o != nil && o.EnableHumanModeration != nil {
		return true
	}

	return false
}

// SetEnableHumanModeration gets a reference to the given int32 and assigns it to the EnableHumanModeration field.
func (o *KalturaReachProfile) SetEnableHumanModeration(v int32) {
	o.EnableHumanModeration = &v
}

// GetEnableMachineModeration returns the EnableMachineModeration field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableMachineModeration() int32 {
	if o == nil || o.EnableMachineModeration == nil {
		var ret int32
		return ret
	}
	return *o.EnableMachineModeration
}

// GetEnableMachineModerationOk returns a tuple with the EnableMachineModeration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableMachineModerationOk() (*int32, bool) {
	if o == nil || o.EnableMachineModeration == nil {
		return nil, false
	}
	return o.EnableMachineModeration, true
}

// HasEnableMachineModeration returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableMachineModeration() bool {
	if o != nil && o.EnableMachineModeration != nil {
		return true
	}

	return false
}

// SetEnableMachineModeration gets a reference to the given int32 and assigns it to the EnableMachineModeration field.
func (o *KalturaReachProfile) SetEnableMachineModeration(v int32) {
	o.EnableMachineModeration = &v
}

// GetEnableMetadataExtraction returns the EnableMetadataExtraction field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableMetadataExtraction() int32 {
	if o == nil || o.EnableMetadataExtraction == nil {
		var ret int32
		return ret
	}
	return *o.EnableMetadataExtraction
}

// GetEnableMetadataExtractionOk returns a tuple with the EnableMetadataExtraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableMetadataExtractionOk() (*int32, bool) {
	if o == nil || o.EnableMetadataExtraction == nil {
		return nil, false
	}
	return o.EnableMetadataExtraction, true
}

// HasEnableMetadataExtraction returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableMetadataExtraction() bool {
	if o != nil && o.EnableMetadataExtraction != nil {
		return true
	}

	return false
}

// SetEnableMetadataExtraction gets a reference to the given int32 and assigns it to the EnableMetadataExtraction field.
func (o *KalturaReachProfile) SetEnableMetadataExtraction(v int32) {
	o.EnableMetadataExtraction = &v
}

// GetEnableProfanityRemoval returns the EnableProfanityRemoval field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableProfanityRemoval() int32 {
	if o == nil || o.EnableProfanityRemoval == nil {
		var ret int32
		return ret
	}
	return *o.EnableProfanityRemoval
}

// GetEnableProfanityRemovalOk returns a tuple with the EnableProfanityRemoval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableProfanityRemovalOk() (*int32, bool) {
	if o == nil || o.EnableProfanityRemoval == nil {
		return nil, false
	}
	return o.EnableProfanityRemoval, true
}

// HasEnableProfanityRemoval returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableProfanityRemoval() bool {
	if o != nil && o.EnableProfanityRemoval != nil {
		return true
	}

	return false
}

// SetEnableProfanityRemoval gets a reference to the given int32 and assigns it to the EnableProfanityRemoval field.
func (o *KalturaReachProfile) SetEnableProfanityRemoval(v int32) {
	o.EnableProfanityRemoval = &v
}

// GetEnableSpeakerChangeIndication returns the EnableSpeakerChangeIndication field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetEnableSpeakerChangeIndication() int32 {
	if o == nil || o.EnableSpeakerChangeIndication == nil {
		var ret int32
		return ret
	}
	return *o.EnableSpeakerChangeIndication
}

// GetEnableSpeakerChangeIndicationOk returns a tuple with the EnableSpeakerChangeIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetEnableSpeakerChangeIndicationOk() (*int32, bool) {
	if o == nil || o.EnableSpeakerChangeIndication == nil {
		return nil, false
	}
	return o.EnableSpeakerChangeIndication, true
}

// HasEnableSpeakerChangeIndication returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasEnableSpeakerChangeIndication() bool {
	if o != nil && o.EnableSpeakerChangeIndication != nil {
		return true
	}

	return false
}

// SetEnableSpeakerChangeIndication gets a reference to the given int32 and assigns it to the EnableSpeakerChangeIndication field.
func (o *KalturaReachProfile) SetEnableSpeakerChangeIndication(v int32) {
	o.EnableSpeakerChangeIndication = &v
}

// GetFlavorParamsIds returns the FlavorParamsIds field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetFlavorParamsIds() string {
	if o == nil || o.FlavorParamsIds == nil {
		var ret string
		return ret
	}
	return *o.FlavorParamsIds
}

// GetFlavorParamsIdsOk returns a tuple with the FlavorParamsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetFlavorParamsIdsOk() (*string, bool) {
	if o == nil || o.FlavorParamsIds == nil {
		return nil, false
	}
	return o.FlavorParamsIds, true
}

// HasFlavorParamsIds returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasFlavorParamsIds() bool {
	if o != nil && o.FlavorParamsIds != nil {
		return true
	}

	return false
}

// SetFlavorParamsIds gets a reference to the given string and assigns it to the FlavorParamsIds field.
func (o *KalturaReachProfile) SetFlavorParamsIds(v string) {
	o.FlavorParamsIds = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaReachProfile) SetId(v int32) {
	o.Id = &v
}

// GetLabelAdditionForHumanServiceType returns the LabelAdditionForHumanServiceType field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetLabelAdditionForHumanServiceType() string {
	if o == nil || o.LabelAdditionForHumanServiceType == nil {
		var ret string
		return ret
	}
	return *o.LabelAdditionForHumanServiceType
}

// GetLabelAdditionForHumanServiceTypeOk returns a tuple with the LabelAdditionForHumanServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetLabelAdditionForHumanServiceTypeOk() (*string, bool) {
	if o == nil || o.LabelAdditionForHumanServiceType == nil {
		return nil, false
	}
	return o.LabelAdditionForHumanServiceType, true
}

// HasLabelAdditionForHumanServiceType returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasLabelAdditionForHumanServiceType() bool {
	if o != nil && o.LabelAdditionForHumanServiceType != nil {
		return true
	}

	return false
}

// SetLabelAdditionForHumanServiceType gets a reference to the given string and assigns it to the LabelAdditionForHumanServiceType field.
func (o *KalturaReachProfile) SetLabelAdditionForHumanServiceType(v string) {
	o.LabelAdditionForHumanServiceType = &v
}

// GetLabelAdditionForMachineServiceType returns the LabelAdditionForMachineServiceType field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetLabelAdditionForMachineServiceType() string {
	if o == nil || o.LabelAdditionForMachineServiceType == nil {
		var ret string
		return ret
	}
	return *o.LabelAdditionForMachineServiceType
}

// GetLabelAdditionForMachineServiceTypeOk returns a tuple with the LabelAdditionForMachineServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetLabelAdditionForMachineServiceTypeOk() (*string, bool) {
	if o == nil || o.LabelAdditionForMachineServiceType == nil {
		return nil, false
	}
	return o.LabelAdditionForMachineServiceType, true
}

// HasLabelAdditionForMachineServiceType returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasLabelAdditionForMachineServiceType() bool {
	if o != nil && o.LabelAdditionForMachineServiceType != nil {
		return true
	}

	return false
}

// SetLabelAdditionForMachineServiceType gets a reference to the given string and assigns it to the LabelAdditionForMachineServiceType field.
func (o *KalturaReachProfile) SetLabelAdditionForMachineServiceType(v string) {
	o.LabelAdditionForMachineServiceType = &v
}

// GetMaxCharactersPerCaptionLine returns the MaxCharactersPerCaptionLine field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetMaxCharactersPerCaptionLine() int32 {
	if o == nil || o.MaxCharactersPerCaptionLine == nil {
		var ret int32
		return ret
	}
	return *o.MaxCharactersPerCaptionLine
}

// GetMaxCharactersPerCaptionLineOk returns a tuple with the MaxCharactersPerCaptionLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetMaxCharactersPerCaptionLineOk() (*int32, bool) {
	if o == nil || o.MaxCharactersPerCaptionLine == nil {
		return nil, false
	}
	return o.MaxCharactersPerCaptionLine, true
}

// HasMaxCharactersPerCaptionLine returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasMaxCharactersPerCaptionLine() bool {
	if o != nil && o.MaxCharactersPerCaptionLine != nil {
		return true
	}

	return false
}

// SetMaxCharactersPerCaptionLine gets a reference to the given int32 and assigns it to the MaxCharactersPerCaptionLine field.
func (o *KalturaReachProfile) SetMaxCharactersPerCaptionLine(v int32) {
	o.MaxCharactersPerCaptionLine = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaReachProfile) SetName(v string) {
	o.Name = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaReachProfile) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetProfileType() int32 {
	if o == nil || o.ProfileType == nil {
		var ret int32
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetProfileTypeOk() (*int32, bool) {
	if o == nil || o.ProfileType == nil {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasProfileType() bool {
	if o != nil && o.ProfileType != nil {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given int32 and assigns it to the ProfileType field.
func (o *KalturaReachProfile) SetProfileType(v int32) {
	o.ProfileType = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetRules() []KalturaRule {
	if o == nil || o.Rules == nil {
		var ret []KalturaRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetRulesOk() ([]KalturaRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []KalturaRule and assigns it to the Rules field.
func (o *KalturaReachProfile) SetRules(v []KalturaRule) {
	o.Rules = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaReachProfile) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaReachProfile) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUsedCredit returns the UsedCredit field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetUsedCredit() float32 {
	if o == nil || o.UsedCredit == nil {
		var ret float32
		return ret
	}
	return *o.UsedCredit
}

// GetUsedCreditOk returns a tuple with the UsedCredit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetUsedCreditOk() (*float32, bool) {
	if o == nil || o.UsedCredit == nil {
		return nil, false
	}
	return o.UsedCredit, true
}

// HasUsedCredit returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasUsedCredit() bool {
	if o != nil && o.UsedCredit != nil {
		return true
	}

	return false
}

// SetUsedCredit gets a reference to the given float32 and assigns it to the UsedCredit field.
func (o *KalturaReachProfile) SetUsedCredit(v float32) {
	o.UsedCredit = &v
}

// GetVendorTaskProcessingRegion returns the VendorTaskProcessingRegion field value if set, zero value otherwise.
func (o *KalturaReachProfile) GetVendorTaskProcessingRegion() int32 {
	if o == nil || o.VendorTaskProcessingRegion == nil {
		var ret int32
		return ret
	}
	return *o.VendorTaskProcessingRegion
}

// GetVendorTaskProcessingRegionOk returns a tuple with the VendorTaskProcessingRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReachProfile) GetVendorTaskProcessingRegionOk() (*int32, bool) {
	if o == nil || o.VendorTaskProcessingRegion == nil {
		return nil, false
	}
	return o.VendorTaskProcessingRegion, true
}

// HasVendorTaskProcessingRegion returns a boolean if a field has been set.
func (o *KalturaReachProfile) HasVendorTaskProcessingRegion() bool {
	if o != nil && o.VendorTaskProcessingRegion != nil {
		return true
	}

	return false
}

// SetVendorTaskProcessingRegion gets a reference to the given int32 and assigns it to the VendorTaskProcessingRegion field.
func (o *KalturaReachProfile) SetVendorTaskProcessingRegion(v int32) {
	o.VendorTaskProcessingRegion = &v
}

func (o KalturaReachProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoDisplayHumanCaptionsOnPlayer != nil {
		toSerialize["autoDisplayHumanCaptionsOnPlayer"] = o.AutoDisplayHumanCaptionsOnPlayer
	}
	if o.AutoDisplayMachineCaptionsOnPlayer != nil {
		toSerialize["autoDisplayMachineCaptionsOnPlayer"] = o.AutoDisplayMachineCaptionsOnPlayer
	}
	if o.ContentDeletionPolicy != nil {
		toSerialize["contentDeletionPolicy"] = o.ContentDeletionPolicy
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.DefaultOutputFormat != nil {
		toSerialize["defaultOutputFormat"] = o.DefaultOutputFormat
	}
	if o.Dictionaries != nil {
		toSerialize["dictionaries"] = o.Dictionaries
	}
	if o.EnableAudioTags != nil {
		toSerialize["enableAudioTags"] = o.EnableAudioTags
	}
	if o.EnableHumanModeration != nil {
		toSerialize["enableHumanModeration"] = o.EnableHumanModeration
	}
	if o.EnableMachineModeration != nil {
		toSerialize["enableMachineModeration"] = o.EnableMachineModeration
	}
	if o.EnableMetadataExtraction != nil {
		toSerialize["enableMetadataExtraction"] = o.EnableMetadataExtraction
	}
	if o.EnableProfanityRemoval != nil {
		toSerialize["enableProfanityRemoval"] = o.EnableProfanityRemoval
	}
	if o.EnableSpeakerChangeIndication != nil {
		toSerialize["enableSpeakerChangeIndication"] = o.EnableSpeakerChangeIndication
	}
	if o.FlavorParamsIds != nil {
		toSerialize["flavorParamsIds"] = o.FlavorParamsIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LabelAdditionForHumanServiceType != nil {
		toSerialize["labelAdditionForHumanServiceType"] = o.LabelAdditionForHumanServiceType
	}
	if o.LabelAdditionForMachineServiceType != nil {
		toSerialize["labelAdditionForMachineServiceType"] = o.LabelAdditionForMachineServiceType
	}
	if o.MaxCharactersPerCaptionLine != nil {
		toSerialize["maxCharactersPerCaptionLine"] = o.MaxCharactersPerCaptionLine
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.ProfileType != nil {
		toSerialize["profileType"] = o.ProfileType
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UsedCredit != nil {
		toSerialize["usedCredit"] = o.UsedCredit
	}
	if o.VendorTaskProcessingRegion != nil {
		toSerialize["vendorTaskProcessingRegion"] = o.VendorTaskProcessingRegion
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaReachProfile struct {
	value *KalturaReachProfile
	isSet bool
}

func (v NullableKalturaReachProfile) Get() *KalturaReachProfile {
	return v.value
}

func (v *NullableKalturaReachProfile) Set(val *KalturaReachProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReachProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReachProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReachProfile(val *KalturaReachProfile) *NullableKalturaReachProfile {
	return &NullableKalturaReachProfile{value: val, isSet: true}
}

func (v NullableKalturaReachProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReachProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



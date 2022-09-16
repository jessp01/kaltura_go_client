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

// KalturaEntryDistribution struct for KalturaEntryDistribution
type KalturaEntryDistribution struct {
	// Comma separated asset ids
	AssetIds *string `json:"assetIds,omitempty"`
	// `readOnly`  Entry distribution creation date as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`  Enum Type: `KalturaEntryDistributionFlag`
	DirtyStatus *int32 `json:"dirtyStatus,omitempty"`
	// `insertOnly`
	DistributionProfileId *int32 `json:"distributionProfileId,omitempty"`
	// `insertOnly`
	EntryId *string `json:"entryId,omitempty"`
	// `readOnly`
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// `readOnly`
	ErrorNumber *int32 `json:"errorNumber,omitempty"`
	// `readOnly`  Enum Type: `KalturaBatchJobErrorTypes`
	ErrorType *int32 `json:"errorType,omitempty"`
	// Comma separated flavor asset ids
	FlavorAssetIds *string `json:"flavorAssetIds,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasDeleteResultsLog *int32 `json:"hasDeleteResultsLog,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasDeleteSentDataLog *int32 `json:"hasDeleteSentDataLog,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasSubmitResultsLog *int32 `json:"hasSubmitResultsLog,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasSubmitSentDataLog *int32 `json:"hasSubmitSentDataLog,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasUpdateResultsLog *int32 `json:"hasUpdateResultsLog,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`
	HasUpdateSentDataLog *int32 `json:"hasUpdateSentDataLog,omitempty"`
	// `readOnly`  Auto generated unique id
	Id *int32 `json:"id,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  The plays as retrieved from the remote destination reports
	Plays *int32 `json:"plays,omitempty"`
	// `readOnly`  The id as returned from the distributed destination
	RemoteId *string `json:"remoteId,omitempty"`
	// `readOnly`  Enum Type: `KalturaEntryDistributionStatus`
	Status *int32 `json:"status,omitempty"`
	// `readOnly`  Entry distribution submission date as Unix timestamp (In seconds)
	SubmittedAt *int32 `json:"submittedAt,omitempty"`
	// `readOnly`  Enum Type: `KalturaEntryDistributionSunStatus`
	SunStatus *int32 `json:"sunStatus,omitempty"`
	// Entry distribution publish time as Unix timestamp (In seconds)
	Sunrise *int32 `json:"sunrise,omitempty"`
	// Entry distribution un-publish time as Unix timestamp (In seconds)
	Sunset *int32 `json:"sunset,omitempty"`
	// Comma separated thumbnail asset ids
	ThumbAssetIds *string `json:"thumbAssetIds,omitempty"`
	// `readOnly`  Entry distribution last update date as Unix timestamp (In seconds)
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	ValidationErrors []KalturaDistributionValidationError `json:"validationErrors,omitempty"`
	// `readOnly`  The views as retrieved from the remote destination reports
	Views *int32 `json:"views,omitempty"`
}

// NewKalturaEntryDistribution instantiates a new KalturaEntryDistribution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryDistribution() *KalturaEntryDistribution {
	this := KalturaEntryDistribution{}
	return &this
}

// NewKalturaEntryDistributionWithDefaults instantiates a new KalturaEntryDistribution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryDistributionWithDefaults() *KalturaEntryDistribution {
	this := KalturaEntryDistribution{}
	return &this
}

// GetAssetIds returns the AssetIds field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetAssetIds() string {
	if o == nil || o.AssetIds == nil {
		var ret string
		return ret
	}
	return *o.AssetIds
}

// GetAssetIdsOk returns a tuple with the AssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetAssetIdsOk() (*string, bool) {
	if o == nil || o.AssetIds == nil {
		return nil, false
	}
	return o.AssetIds, true
}

// HasAssetIds returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasAssetIds() bool {
	if o != nil && o.AssetIds != nil {
		return true
	}

	return false
}

// SetAssetIds gets a reference to the given string and assigns it to the AssetIds field.
func (o *KalturaEntryDistribution) SetAssetIds(v string) {
	o.AssetIds = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaEntryDistribution) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDirtyStatus returns the DirtyStatus field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetDirtyStatus() int32 {
	if o == nil || o.DirtyStatus == nil {
		var ret int32
		return ret
	}
	return *o.DirtyStatus
}

// GetDirtyStatusOk returns a tuple with the DirtyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetDirtyStatusOk() (*int32, bool) {
	if o == nil || o.DirtyStatus == nil {
		return nil, false
	}
	return o.DirtyStatus, true
}

// HasDirtyStatus returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasDirtyStatus() bool {
	if o != nil && o.DirtyStatus != nil {
		return true
	}

	return false
}

// SetDirtyStatus gets a reference to the given int32 and assigns it to the DirtyStatus field.
func (o *KalturaEntryDistribution) SetDirtyStatus(v int32) {
	o.DirtyStatus = &v
}

// GetDistributionProfileId returns the DistributionProfileId field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetDistributionProfileId() int32 {
	if o == nil || o.DistributionProfileId == nil {
		var ret int32
		return ret
	}
	return *o.DistributionProfileId
}

// GetDistributionProfileIdOk returns a tuple with the DistributionProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetDistributionProfileIdOk() (*int32, bool) {
	if o == nil || o.DistributionProfileId == nil {
		return nil, false
	}
	return o.DistributionProfileId, true
}

// HasDistributionProfileId returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasDistributionProfileId() bool {
	if o != nil && o.DistributionProfileId != nil {
		return true
	}

	return false
}

// SetDistributionProfileId gets a reference to the given int32 and assigns it to the DistributionProfileId field.
func (o *KalturaEntryDistribution) SetDistributionProfileId(v int32) {
	o.DistributionProfileId = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaEntryDistribution) SetEntryId(v string) {
	o.EntryId = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *KalturaEntryDistribution) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorNumber returns the ErrorNumber field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetErrorNumber() int32 {
	if o == nil || o.ErrorNumber == nil {
		var ret int32
		return ret
	}
	return *o.ErrorNumber
}

// GetErrorNumberOk returns a tuple with the ErrorNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetErrorNumberOk() (*int32, bool) {
	if o == nil || o.ErrorNumber == nil {
		return nil, false
	}
	return o.ErrorNumber, true
}

// HasErrorNumber returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasErrorNumber() bool {
	if o != nil && o.ErrorNumber != nil {
		return true
	}

	return false
}

// SetErrorNumber gets a reference to the given int32 and assigns it to the ErrorNumber field.
func (o *KalturaEntryDistribution) SetErrorNumber(v int32) {
	o.ErrorNumber = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetErrorType() int32 {
	if o == nil || o.ErrorType == nil {
		var ret int32
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetErrorTypeOk() (*int32, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given int32 and assigns it to the ErrorType field.
func (o *KalturaEntryDistribution) SetErrorType(v int32) {
	o.ErrorType = &v
}

// GetFlavorAssetIds returns the FlavorAssetIds field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetFlavorAssetIds() string {
	if o == nil || o.FlavorAssetIds == nil {
		var ret string
		return ret
	}
	return *o.FlavorAssetIds
}

// GetFlavorAssetIdsOk returns a tuple with the FlavorAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetFlavorAssetIdsOk() (*string, bool) {
	if o == nil || o.FlavorAssetIds == nil {
		return nil, false
	}
	return o.FlavorAssetIds, true
}

// HasFlavorAssetIds returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasFlavorAssetIds() bool {
	if o != nil && o.FlavorAssetIds != nil {
		return true
	}

	return false
}

// SetFlavorAssetIds gets a reference to the given string and assigns it to the FlavorAssetIds field.
func (o *KalturaEntryDistribution) SetFlavorAssetIds(v string) {
	o.FlavorAssetIds = &v
}

// GetHasDeleteResultsLog returns the HasDeleteResultsLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasDeleteResultsLog() int32 {
	if o == nil || o.HasDeleteResultsLog == nil {
		var ret int32
		return ret
	}
	return *o.HasDeleteResultsLog
}

// GetHasDeleteResultsLogOk returns a tuple with the HasDeleteResultsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasDeleteResultsLogOk() (*int32, bool) {
	if o == nil || o.HasDeleteResultsLog == nil {
		return nil, false
	}
	return o.HasDeleteResultsLog, true
}

// HasHasDeleteResultsLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasDeleteResultsLog() bool {
	if o != nil && o.HasDeleteResultsLog != nil {
		return true
	}

	return false
}

// SetHasDeleteResultsLog gets a reference to the given int32 and assigns it to the HasDeleteResultsLog field.
func (o *KalturaEntryDistribution) SetHasDeleteResultsLog(v int32) {
	o.HasDeleteResultsLog = &v
}

// GetHasDeleteSentDataLog returns the HasDeleteSentDataLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasDeleteSentDataLog() int32 {
	if o == nil || o.HasDeleteSentDataLog == nil {
		var ret int32
		return ret
	}
	return *o.HasDeleteSentDataLog
}

// GetHasDeleteSentDataLogOk returns a tuple with the HasDeleteSentDataLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasDeleteSentDataLogOk() (*int32, bool) {
	if o == nil || o.HasDeleteSentDataLog == nil {
		return nil, false
	}
	return o.HasDeleteSentDataLog, true
}

// HasHasDeleteSentDataLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasDeleteSentDataLog() bool {
	if o != nil && o.HasDeleteSentDataLog != nil {
		return true
	}

	return false
}

// SetHasDeleteSentDataLog gets a reference to the given int32 and assigns it to the HasDeleteSentDataLog field.
func (o *KalturaEntryDistribution) SetHasDeleteSentDataLog(v int32) {
	o.HasDeleteSentDataLog = &v
}

// GetHasSubmitResultsLog returns the HasSubmitResultsLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasSubmitResultsLog() int32 {
	if o == nil || o.HasSubmitResultsLog == nil {
		var ret int32
		return ret
	}
	return *o.HasSubmitResultsLog
}

// GetHasSubmitResultsLogOk returns a tuple with the HasSubmitResultsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasSubmitResultsLogOk() (*int32, bool) {
	if o == nil || o.HasSubmitResultsLog == nil {
		return nil, false
	}
	return o.HasSubmitResultsLog, true
}

// HasHasSubmitResultsLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasSubmitResultsLog() bool {
	if o != nil && o.HasSubmitResultsLog != nil {
		return true
	}

	return false
}

// SetHasSubmitResultsLog gets a reference to the given int32 and assigns it to the HasSubmitResultsLog field.
func (o *KalturaEntryDistribution) SetHasSubmitResultsLog(v int32) {
	o.HasSubmitResultsLog = &v
}

// GetHasSubmitSentDataLog returns the HasSubmitSentDataLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasSubmitSentDataLog() int32 {
	if o == nil || o.HasSubmitSentDataLog == nil {
		var ret int32
		return ret
	}
	return *o.HasSubmitSentDataLog
}

// GetHasSubmitSentDataLogOk returns a tuple with the HasSubmitSentDataLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasSubmitSentDataLogOk() (*int32, bool) {
	if o == nil || o.HasSubmitSentDataLog == nil {
		return nil, false
	}
	return o.HasSubmitSentDataLog, true
}

// HasHasSubmitSentDataLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasSubmitSentDataLog() bool {
	if o != nil && o.HasSubmitSentDataLog != nil {
		return true
	}

	return false
}

// SetHasSubmitSentDataLog gets a reference to the given int32 and assigns it to the HasSubmitSentDataLog field.
func (o *KalturaEntryDistribution) SetHasSubmitSentDataLog(v int32) {
	o.HasSubmitSentDataLog = &v
}

// GetHasUpdateResultsLog returns the HasUpdateResultsLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasUpdateResultsLog() int32 {
	if o == nil || o.HasUpdateResultsLog == nil {
		var ret int32
		return ret
	}
	return *o.HasUpdateResultsLog
}

// GetHasUpdateResultsLogOk returns a tuple with the HasUpdateResultsLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasUpdateResultsLogOk() (*int32, bool) {
	if o == nil || o.HasUpdateResultsLog == nil {
		return nil, false
	}
	return o.HasUpdateResultsLog, true
}

// HasHasUpdateResultsLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasUpdateResultsLog() bool {
	if o != nil && o.HasUpdateResultsLog != nil {
		return true
	}

	return false
}

// SetHasUpdateResultsLog gets a reference to the given int32 and assigns it to the HasUpdateResultsLog field.
func (o *KalturaEntryDistribution) SetHasUpdateResultsLog(v int32) {
	o.HasUpdateResultsLog = &v
}

// GetHasUpdateSentDataLog returns the HasUpdateSentDataLog field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetHasUpdateSentDataLog() int32 {
	if o == nil || o.HasUpdateSentDataLog == nil {
		var ret int32
		return ret
	}
	return *o.HasUpdateSentDataLog
}

// GetHasUpdateSentDataLogOk returns a tuple with the HasUpdateSentDataLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetHasUpdateSentDataLogOk() (*int32, bool) {
	if o == nil || o.HasUpdateSentDataLog == nil {
		return nil, false
	}
	return o.HasUpdateSentDataLog, true
}

// HasHasUpdateSentDataLog returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasHasUpdateSentDataLog() bool {
	if o != nil && o.HasUpdateSentDataLog != nil {
		return true
	}

	return false
}

// SetHasUpdateSentDataLog gets a reference to the given int32 and assigns it to the HasUpdateSentDataLog field.
func (o *KalturaEntryDistribution) SetHasUpdateSentDataLog(v int32) {
	o.HasUpdateSentDataLog = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaEntryDistribution) SetId(v int32) {
	o.Id = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaEntryDistribution) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPlays returns the Plays field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetPlays() int32 {
	if o == nil || o.Plays == nil {
		var ret int32
		return ret
	}
	return *o.Plays
}

// GetPlaysOk returns a tuple with the Plays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetPlaysOk() (*int32, bool) {
	if o == nil || o.Plays == nil {
		return nil, false
	}
	return o.Plays, true
}

// HasPlays returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasPlays() bool {
	if o != nil && o.Plays != nil {
		return true
	}

	return false
}

// SetPlays gets a reference to the given int32 and assigns it to the Plays field.
func (o *KalturaEntryDistribution) SetPlays(v int32) {
	o.Plays = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *KalturaEntryDistribution) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaEntryDistribution) SetStatus(v int32) {
	o.Status = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetSubmittedAt() int32 {
	if o == nil || o.SubmittedAt == nil {
		var ret int32
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetSubmittedAtOk() (*int32, bool) {
	if o == nil || o.SubmittedAt == nil {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasSubmittedAt() bool {
	if o != nil && o.SubmittedAt != nil {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given int32 and assigns it to the SubmittedAt field.
func (o *KalturaEntryDistribution) SetSubmittedAt(v int32) {
	o.SubmittedAt = &v
}

// GetSunStatus returns the SunStatus field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetSunStatus() int32 {
	if o == nil || o.SunStatus == nil {
		var ret int32
		return ret
	}
	return *o.SunStatus
}

// GetSunStatusOk returns a tuple with the SunStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetSunStatusOk() (*int32, bool) {
	if o == nil || o.SunStatus == nil {
		return nil, false
	}
	return o.SunStatus, true
}

// HasSunStatus returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasSunStatus() bool {
	if o != nil && o.SunStatus != nil {
		return true
	}

	return false
}

// SetSunStatus gets a reference to the given int32 and assigns it to the SunStatus field.
func (o *KalturaEntryDistribution) SetSunStatus(v int32) {
	o.SunStatus = &v
}

// GetSunrise returns the Sunrise field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetSunrise() int32 {
	if o == nil || o.Sunrise == nil {
		var ret int32
		return ret
	}
	return *o.Sunrise
}

// GetSunriseOk returns a tuple with the Sunrise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetSunriseOk() (*int32, bool) {
	if o == nil || o.Sunrise == nil {
		return nil, false
	}
	return o.Sunrise, true
}

// HasSunrise returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasSunrise() bool {
	if o != nil && o.Sunrise != nil {
		return true
	}

	return false
}

// SetSunrise gets a reference to the given int32 and assigns it to the Sunrise field.
func (o *KalturaEntryDistribution) SetSunrise(v int32) {
	o.Sunrise = &v
}

// GetSunset returns the Sunset field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetSunset() int32 {
	if o == nil || o.Sunset == nil {
		var ret int32
		return ret
	}
	return *o.Sunset
}

// GetSunsetOk returns a tuple with the Sunset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetSunsetOk() (*int32, bool) {
	if o == nil || o.Sunset == nil {
		return nil, false
	}
	return o.Sunset, true
}

// HasSunset returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasSunset() bool {
	if o != nil && o.Sunset != nil {
		return true
	}

	return false
}

// SetSunset gets a reference to the given int32 and assigns it to the Sunset field.
func (o *KalturaEntryDistribution) SetSunset(v int32) {
	o.Sunset = &v
}

// GetThumbAssetIds returns the ThumbAssetIds field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetThumbAssetIds() string {
	if o == nil || o.ThumbAssetIds == nil {
		var ret string
		return ret
	}
	return *o.ThumbAssetIds
}

// GetThumbAssetIdsOk returns a tuple with the ThumbAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetThumbAssetIdsOk() (*string, bool) {
	if o == nil || o.ThumbAssetIds == nil {
		return nil, false
	}
	return o.ThumbAssetIds, true
}

// HasThumbAssetIds returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasThumbAssetIds() bool {
	if o != nil && o.ThumbAssetIds != nil {
		return true
	}

	return false
}

// SetThumbAssetIds gets a reference to the given string and assigns it to the ThumbAssetIds field.
func (o *KalturaEntryDistribution) SetThumbAssetIds(v string) {
	o.ThumbAssetIds = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaEntryDistribution) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetValidationErrors() []KalturaDistributionValidationError {
	if o == nil || o.ValidationErrors == nil {
		var ret []KalturaDistributionValidationError
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetValidationErrorsOk() ([]KalturaDistributionValidationError, bool) {
	if o == nil || o.ValidationErrors == nil {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasValidationErrors() bool {
	if o != nil && o.ValidationErrors != nil {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []KalturaDistributionValidationError and assigns it to the ValidationErrors field.
func (o *KalturaEntryDistribution) SetValidationErrors(v []KalturaDistributionValidationError) {
	o.ValidationErrors = v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *KalturaEntryDistribution) GetViews() int32 {
	if o == nil || o.Views == nil {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaEntryDistribution) GetViewsOk() (*int32, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *KalturaEntryDistribution) HasViews() bool {
	if o != nil && o.Views != nil {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *KalturaEntryDistribution) SetViews(v int32) {
	o.Views = &v
}

func (o KalturaEntryDistribution) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetIds != nil {
		toSerialize["assetIds"] = o.AssetIds
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DirtyStatus != nil {
		toSerialize["dirtyStatus"] = o.DirtyStatus
	}
	if o.DistributionProfileId != nil {
		toSerialize["distributionProfileId"] = o.DistributionProfileId
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorNumber != nil {
		toSerialize["errorNumber"] = o.ErrorNumber
	}
	if o.ErrorType != nil {
		toSerialize["errorType"] = o.ErrorType
	}
	if o.FlavorAssetIds != nil {
		toSerialize["flavorAssetIds"] = o.FlavorAssetIds
	}
	if o.HasDeleteResultsLog != nil {
		toSerialize["hasDeleteResultsLog"] = o.HasDeleteResultsLog
	}
	if o.HasDeleteSentDataLog != nil {
		toSerialize["hasDeleteSentDataLog"] = o.HasDeleteSentDataLog
	}
	if o.HasSubmitResultsLog != nil {
		toSerialize["hasSubmitResultsLog"] = o.HasSubmitResultsLog
	}
	if o.HasSubmitSentDataLog != nil {
		toSerialize["hasSubmitSentDataLog"] = o.HasSubmitSentDataLog
	}
	if o.HasUpdateResultsLog != nil {
		toSerialize["hasUpdateResultsLog"] = o.HasUpdateResultsLog
	}
	if o.HasUpdateSentDataLog != nil {
		toSerialize["hasUpdateSentDataLog"] = o.HasUpdateSentDataLog
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Plays != nil {
		toSerialize["plays"] = o.Plays
	}
	if o.RemoteId != nil {
		toSerialize["remoteId"] = o.RemoteId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SubmittedAt != nil {
		toSerialize["submittedAt"] = o.SubmittedAt
	}
	if o.SunStatus != nil {
		toSerialize["sunStatus"] = o.SunStatus
	}
	if o.Sunrise != nil {
		toSerialize["sunrise"] = o.Sunrise
	}
	if o.Sunset != nil {
		toSerialize["sunset"] = o.Sunset
	}
	if o.ThumbAssetIds != nil {
		toSerialize["thumbAssetIds"] = o.ThumbAssetIds
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.ValidationErrors != nil {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryDistribution struct {
	value *KalturaEntryDistribution
	isSet bool
}

func (v NullableKalturaEntryDistribution) Get() *KalturaEntryDistribution {
	return v.value
}

func (v *NullableKalturaEntryDistribution) Set(val *KalturaEntryDistribution) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryDistribution) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryDistribution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryDistribution(val *KalturaEntryDistribution) *NullableKalturaEntryDistribution {
	return &NullableKalturaEntryDistribution{value: val, isSet: true}
}

func (v NullableKalturaEntryDistribution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryDistribution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



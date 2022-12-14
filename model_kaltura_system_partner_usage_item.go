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

// KalturaSystemPartnerUsageItem struct for KalturaSystemPartnerUsageItem
type KalturaSystemPartnerUsageItem struct {
	// Number of new audio entries created during specific date range
	AudioEntriesCount *int32 `json:"audioEntriesCount,omitempty"`
	// The average amount of storage consumption during the given date range for the specific publisher
	AvgStorage *float32 `json:"avgStorage,omitempty"`
	// The total bandwidth usage during the given date range (in MB)
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	// The combined amount of bandwidth and storage consumed during the given date range for the specific publisher
	CombinedBandwidthStorage *float32 `json:"combinedBandwidthStorage,omitempty"`
	// Amount of deleted storage in MB
	DeletedStorage *float32 `json:"deletedStorage,omitempty"`
	// Number of new entries created during specific date range
	EntriesCount *int32 `json:"entriesCount,omitempty"`
	// Number of new image entries created during specific date range
	ImageEntriesCount *int32 `json:"imageEntriesCount,omitempty"`
	// Number of new mix entries created during specific date range
	MixEntriesCount *int32 `json:"mixEntriesCount,omitempty"`
	// Partner creation date (Unix timestamp)
	PartnerCreatedAt *int32 `json:"partnerCreatedAt,omitempty"`
	// Partner ID
	PartnerId *int32 `json:"partnerId,omitempty"`
	// Partner name
	PartnerName *string `json:"partnerName,omitempty"`
	// Partner package
	PartnerPackage *int32 `json:"partnerPackage,omitempty"`
	// Enum Type: `KalturaPartnerStatus`  Partner status
	PartnerStatus *int32 `json:"partnerStatus,omitempty"`
	// The peak amount of storage consumption during the given date range for the specific publisher
	PeakStorage *float32 `json:"peakStorage,omitempty"`
	// Number of plays in the specific date range
	Plays *int32 `json:"plays,omitempty"`
	// The change in storage consumption (new uploads) during the given date range (in MB)
	Storage *float32 `json:"storage,omitempty"`
	// Total number of entries
	TotalEntriesCount *int32 `json:"totalEntriesCount,omitempty"`
	// The total storage consumption (in MB)
	TotalStorage *float32 `json:"totalStorage,omitempty"`
	// Amount of transcoding usage in MB
	TranscodingUsage *float32 `json:"transcodingUsage,omitempty"`
	// Number of new video entries created during specific date range
	VideoEntriesCount *int32 `json:"videoEntriesCount,omitempty"`
	// Number of player loads in the specific date range
	Views *int32 `json:"views,omitempty"`
}

// NewKalturaSystemPartnerUsageItem instantiates a new KalturaSystemPartnerUsageItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSystemPartnerUsageItem() *KalturaSystemPartnerUsageItem {
	this := KalturaSystemPartnerUsageItem{}
	return &this
}

// NewKalturaSystemPartnerUsageItemWithDefaults instantiates a new KalturaSystemPartnerUsageItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSystemPartnerUsageItemWithDefaults() *KalturaSystemPartnerUsageItem {
	this := KalturaSystemPartnerUsageItem{}
	return &this
}

// GetAudioEntriesCount returns the AudioEntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetAudioEntriesCount() int32 {
	if o == nil || o.AudioEntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.AudioEntriesCount
}

// GetAudioEntriesCountOk returns a tuple with the AudioEntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetAudioEntriesCountOk() (*int32, bool) {
	if o == nil || o.AudioEntriesCount == nil {
		return nil, false
	}
	return o.AudioEntriesCount, true
}

// HasAudioEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasAudioEntriesCount() bool {
	if o != nil && o.AudioEntriesCount != nil {
		return true
	}

	return false
}

// SetAudioEntriesCount gets a reference to the given int32 and assigns it to the AudioEntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetAudioEntriesCount(v int32) {
	o.AudioEntriesCount = &v
}

// GetAvgStorage returns the AvgStorage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetAvgStorage() float32 {
	if o == nil || o.AvgStorage == nil {
		var ret float32
		return ret
	}
	return *o.AvgStorage
}

// GetAvgStorageOk returns a tuple with the AvgStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetAvgStorageOk() (*float32, bool) {
	if o == nil || o.AvgStorage == nil {
		return nil, false
	}
	return o.AvgStorage, true
}

// HasAvgStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasAvgStorage() bool {
	if o != nil && o.AvgStorage != nil {
		return true
	}

	return false
}

// SetAvgStorage gets a reference to the given float32 and assigns it to the AvgStorage field.
func (o *KalturaSystemPartnerUsageItem) SetAvgStorage(v float32) {
	o.AvgStorage = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetBandwidth() float32 {
	if o == nil || o.Bandwidth == nil {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetBandwidthOk() (*float32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *KalturaSystemPartnerUsageItem) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetCombinedBandwidthStorage returns the CombinedBandwidthStorage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetCombinedBandwidthStorage() float32 {
	if o == nil || o.CombinedBandwidthStorage == nil {
		var ret float32
		return ret
	}
	return *o.CombinedBandwidthStorage
}

// GetCombinedBandwidthStorageOk returns a tuple with the CombinedBandwidthStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetCombinedBandwidthStorageOk() (*float32, bool) {
	if o == nil || o.CombinedBandwidthStorage == nil {
		return nil, false
	}
	return o.CombinedBandwidthStorage, true
}

// HasCombinedBandwidthStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasCombinedBandwidthStorage() bool {
	if o != nil && o.CombinedBandwidthStorage != nil {
		return true
	}

	return false
}

// SetCombinedBandwidthStorage gets a reference to the given float32 and assigns it to the CombinedBandwidthStorage field.
func (o *KalturaSystemPartnerUsageItem) SetCombinedBandwidthStorage(v float32) {
	o.CombinedBandwidthStorage = &v
}

// GetDeletedStorage returns the DeletedStorage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetDeletedStorage() float32 {
	if o == nil || o.DeletedStorage == nil {
		var ret float32
		return ret
	}
	return *o.DeletedStorage
}

// GetDeletedStorageOk returns a tuple with the DeletedStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetDeletedStorageOk() (*float32, bool) {
	if o == nil || o.DeletedStorage == nil {
		return nil, false
	}
	return o.DeletedStorage, true
}

// HasDeletedStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasDeletedStorage() bool {
	if o != nil && o.DeletedStorage != nil {
		return true
	}

	return false
}

// SetDeletedStorage gets a reference to the given float32 and assigns it to the DeletedStorage field.
func (o *KalturaSystemPartnerUsageItem) SetDeletedStorage(v float32) {
	o.DeletedStorage = &v
}

// GetEntriesCount returns the EntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetEntriesCount() int32 {
	if o == nil || o.EntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.EntriesCount
}

// GetEntriesCountOk returns a tuple with the EntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetEntriesCountOk() (*int32, bool) {
	if o == nil || o.EntriesCount == nil {
		return nil, false
	}
	return o.EntriesCount, true
}

// HasEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasEntriesCount() bool {
	if o != nil && o.EntriesCount != nil {
		return true
	}

	return false
}

// SetEntriesCount gets a reference to the given int32 and assigns it to the EntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetEntriesCount(v int32) {
	o.EntriesCount = &v
}

// GetImageEntriesCount returns the ImageEntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetImageEntriesCount() int32 {
	if o == nil || o.ImageEntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.ImageEntriesCount
}

// GetImageEntriesCountOk returns a tuple with the ImageEntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetImageEntriesCountOk() (*int32, bool) {
	if o == nil || o.ImageEntriesCount == nil {
		return nil, false
	}
	return o.ImageEntriesCount, true
}

// HasImageEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasImageEntriesCount() bool {
	if o != nil && o.ImageEntriesCount != nil {
		return true
	}

	return false
}

// SetImageEntriesCount gets a reference to the given int32 and assigns it to the ImageEntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetImageEntriesCount(v int32) {
	o.ImageEntriesCount = &v
}

// GetMixEntriesCount returns the MixEntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetMixEntriesCount() int32 {
	if o == nil || o.MixEntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.MixEntriesCount
}

// GetMixEntriesCountOk returns a tuple with the MixEntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetMixEntriesCountOk() (*int32, bool) {
	if o == nil || o.MixEntriesCount == nil {
		return nil, false
	}
	return o.MixEntriesCount, true
}

// HasMixEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasMixEntriesCount() bool {
	if o != nil && o.MixEntriesCount != nil {
		return true
	}

	return false
}

// SetMixEntriesCount gets a reference to the given int32 and assigns it to the MixEntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetMixEntriesCount(v int32) {
	o.MixEntriesCount = &v
}

// GetPartnerCreatedAt returns the PartnerCreatedAt field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPartnerCreatedAt() int32 {
	if o == nil || o.PartnerCreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.PartnerCreatedAt
}

// GetPartnerCreatedAtOk returns a tuple with the PartnerCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPartnerCreatedAtOk() (*int32, bool) {
	if o == nil || o.PartnerCreatedAt == nil {
		return nil, false
	}
	return o.PartnerCreatedAt, true
}

// HasPartnerCreatedAt returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPartnerCreatedAt() bool {
	if o != nil && o.PartnerCreatedAt != nil {
		return true
	}

	return false
}

// SetPartnerCreatedAt gets a reference to the given int32 and assigns it to the PartnerCreatedAt field.
func (o *KalturaSystemPartnerUsageItem) SetPartnerCreatedAt(v int32) {
	o.PartnerCreatedAt = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaSystemPartnerUsageItem) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPartnerName returns the PartnerName field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPartnerName() string {
	if o == nil || o.PartnerName == nil {
		var ret string
		return ret
	}
	return *o.PartnerName
}

// GetPartnerNameOk returns a tuple with the PartnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPartnerNameOk() (*string, bool) {
	if o == nil || o.PartnerName == nil {
		return nil, false
	}
	return o.PartnerName, true
}

// HasPartnerName returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPartnerName() bool {
	if o != nil && o.PartnerName != nil {
		return true
	}

	return false
}

// SetPartnerName gets a reference to the given string and assigns it to the PartnerName field.
func (o *KalturaSystemPartnerUsageItem) SetPartnerName(v string) {
	o.PartnerName = &v
}

// GetPartnerPackage returns the PartnerPackage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPartnerPackage() int32 {
	if o == nil || o.PartnerPackage == nil {
		var ret int32
		return ret
	}
	return *o.PartnerPackage
}

// GetPartnerPackageOk returns a tuple with the PartnerPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPartnerPackageOk() (*int32, bool) {
	if o == nil || o.PartnerPackage == nil {
		return nil, false
	}
	return o.PartnerPackage, true
}

// HasPartnerPackage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPartnerPackage() bool {
	if o != nil && o.PartnerPackage != nil {
		return true
	}

	return false
}

// SetPartnerPackage gets a reference to the given int32 and assigns it to the PartnerPackage field.
func (o *KalturaSystemPartnerUsageItem) SetPartnerPackage(v int32) {
	o.PartnerPackage = &v
}

// GetPartnerStatus returns the PartnerStatus field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPartnerStatus() int32 {
	if o == nil || o.PartnerStatus == nil {
		var ret int32
		return ret
	}
	return *o.PartnerStatus
}

// GetPartnerStatusOk returns a tuple with the PartnerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPartnerStatusOk() (*int32, bool) {
	if o == nil || o.PartnerStatus == nil {
		return nil, false
	}
	return o.PartnerStatus, true
}

// HasPartnerStatus returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPartnerStatus() bool {
	if o != nil && o.PartnerStatus != nil {
		return true
	}

	return false
}

// SetPartnerStatus gets a reference to the given int32 and assigns it to the PartnerStatus field.
func (o *KalturaSystemPartnerUsageItem) SetPartnerStatus(v int32) {
	o.PartnerStatus = &v
}

// GetPeakStorage returns the PeakStorage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPeakStorage() float32 {
	if o == nil || o.PeakStorage == nil {
		var ret float32
		return ret
	}
	return *o.PeakStorage
}

// GetPeakStorageOk returns a tuple with the PeakStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPeakStorageOk() (*float32, bool) {
	if o == nil || o.PeakStorage == nil {
		return nil, false
	}
	return o.PeakStorage, true
}

// HasPeakStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPeakStorage() bool {
	if o != nil && o.PeakStorage != nil {
		return true
	}

	return false
}

// SetPeakStorage gets a reference to the given float32 and assigns it to the PeakStorage field.
func (o *KalturaSystemPartnerUsageItem) SetPeakStorage(v float32) {
	o.PeakStorage = &v
}

// GetPlays returns the Plays field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetPlays() int32 {
	if o == nil || o.Plays == nil {
		var ret int32
		return ret
	}
	return *o.Plays
}

// GetPlaysOk returns a tuple with the Plays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetPlaysOk() (*int32, bool) {
	if o == nil || o.Plays == nil {
		return nil, false
	}
	return o.Plays, true
}

// HasPlays returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasPlays() bool {
	if o != nil && o.Plays != nil {
		return true
	}

	return false
}

// SetPlays gets a reference to the given int32 and assigns it to the Plays field.
func (o *KalturaSystemPartnerUsageItem) SetPlays(v int32) {
	o.Plays = &v
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetStorage() float32 {
	if o == nil || o.Storage == nil {
		var ret float32
		return ret
	}
	return *o.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetStorageOk() (*float32, bool) {
	if o == nil || o.Storage == nil {
		return nil, false
	}
	return o.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasStorage() bool {
	if o != nil && o.Storage != nil {
		return true
	}

	return false
}

// SetStorage gets a reference to the given float32 and assigns it to the Storage field.
func (o *KalturaSystemPartnerUsageItem) SetStorage(v float32) {
	o.Storage = &v
}

// GetTotalEntriesCount returns the TotalEntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetTotalEntriesCount() int32 {
	if o == nil || o.TotalEntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalEntriesCount
}

// GetTotalEntriesCountOk returns a tuple with the TotalEntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetTotalEntriesCountOk() (*int32, bool) {
	if o == nil || o.TotalEntriesCount == nil {
		return nil, false
	}
	return o.TotalEntriesCount, true
}

// HasTotalEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasTotalEntriesCount() bool {
	if o != nil && o.TotalEntriesCount != nil {
		return true
	}

	return false
}

// SetTotalEntriesCount gets a reference to the given int32 and assigns it to the TotalEntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetTotalEntriesCount(v int32) {
	o.TotalEntriesCount = &v
}

// GetTotalStorage returns the TotalStorage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetTotalStorage() float32 {
	if o == nil || o.TotalStorage == nil {
		var ret float32
		return ret
	}
	return *o.TotalStorage
}

// GetTotalStorageOk returns a tuple with the TotalStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetTotalStorageOk() (*float32, bool) {
	if o == nil || o.TotalStorage == nil {
		return nil, false
	}
	return o.TotalStorage, true
}

// HasTotalStorage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasTotalStorage() bool {
	if o != nil && o.TotalStorage != nil {
		return true
	}

	return false
}

// SetTotalStorage gets a reference to the given float32 and assigns it to the TotalStorage field.
func (o *KalturaSystemPartnerUsageItem) SetTotalStorage(v float32) {
	o.TotalStorage = &v
}

// GetTranscodingUsage returns the TranscodingUsage field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetTranscodingUsage() float32 {
	if o == nil || o.TranscodingUsage == nil {
		var ret float32
		return ret
	}
	return *o.TranscodingUsage
}

// GetTranscodingUsageOk returns a tuple with the TranscodingUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetTranscodingUsageOk() (*float32, bool) {
	if o == nil || o.TranscodingUsage == nil {
		return nil, false
	}
	return o.TranscodingUsage, true
}

// HasTranscodingUsage returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasTranscodingUsage() bool {
	if o != nil && o.TranscodingUsage != nil {
		return true
	}

	return false
}

// SetTranscodingUsage gets a reference to the given float32 and assigns it to the TranscodingUsage field.
func (o *KalturaSystemPartnerUsageItem) SetTranscodingUsage(v float32) {
	o.TranscodingUsage = &v
}

// GetVideoEntriesCount returns the VideoEntriesCount field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetVideoEntriesCount() int32 {
	if o == nil || o.VideoEntriesCount == nil {
		var ret int32
		return ret
	}
	return *o.VideoEntriesCount
}

// GetVideoEntriesCountOk returns a tuple with the VideoEntriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetVideoEntriesCountOk() (*int32, bool) {
	if o == nil || o.VideoEntriesCount == nil {
		return nil, false
	}
	return o.VideoEntriesCount, true
}

// HasVideoEntriesCount returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasVideoEntriesCount() bool {
	if o != nil && o.VideoEntriesCount != nil {
		return true
	}

	return false
}

// SetVideoEntriesCount gets a reference to the given int32 and assigns it to the VideoEntriesCount field.
func (o *KalturaSystemPartnerUsageItem) SetVideoEntriesCount(v int32) {
	o.VideoEntriesCount = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *KalturaSystemPartnerUsageItem) GetViews() int32 {
	if o == nil || o.Views == nil {
		var ret int32
		return ret
	}
	return *o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaSystemPartnerUsageItem) GetViewsOk() (*int32, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *KalturaSystemPartnerUsageItem) HasViews() bool {
	if o != nil && o.Views != nil {
		return true
	}

	return false
}

// SetViews gets a reference to the given int32 and assigns it to the Views field.
func (o *KalturaSystemPartnerUsageItem) SetViews(v int32) {
	o.Views = &v
}

func (o KalturaSystemPartnerUsageItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioEntriesCount != nil {
		toSerialize["audioEntriesCount"] = o.AudioEntriesCount
	}
	if o.AvgStorage != nil {
		toSerialize["avgStorage"] = o.AvgStorage
	}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.CombinedBandwidthStorage != nil {
		toSerialize["combinedBandwidthStorage"] = o.CombinedBandwidthStorage
	}
	if o.DeletedStorage != nil {
		toSerialize["deletedStorage"] = o.DeletedStorage
	}
	if o.EntriesCount != nil {
		toSerialize["entriesCount"] = o.EntriesCount
	}
	if o.ImageEntriesCount != nil {
		toSerialize["imageEntriesCount"] = o.ImageEntriesCount
	}
	if o.MixEntriesCount != nil {
		toSerialize["mixEntriesCount"] = o.MixEntriesCount
	}
	if o.PartnerCreatedAt != nil {
		toSerialize["partnerCreatedAt"] = o.PartnerCreatedAt
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PartnerName != nil {
		toSerialize["partnerName"] = o.PartnerName
	}
	if o.PartnerPackage != nil {
		toSerialize["partnerPackage"] = o.PartnerPackage
	}
	if o.PartnerStatus != nil {
		toSerialize["partnerStatus"] = o.PartnerStatus
	}
	if o.PeakStorage != nil {
		toSerialize["peakStorage"] = o.PeakStorage
	}
	if o.Plays != nil {
		toSerialize["plays"] = o.Plays
	}
	if o.Storage != nil {
		toSerialize["storage"] = o.Storage
	}
	if o.TotalEntriesCount != nil {
		toSerialize["totalEntriesCount"] = o.TotalEntriesCount
	}
	if o.TotalStorage != nil {
		toSerialize["totalStorage"] = o.TotalStorage
	}
	if o.TranscodingUsage != nil {
		toSerialize["transcodingUsage"] = o.TranscodingUsage
	}
	if o.VideoEntriesCount != nil {
		toSerialize["videoEntriesCount"] = o.VideoEntriesCount
	}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSystemPartnerUsageItem struct {
	value *KalturaSystemPartnerUsageItem
	isSet bool
}

func (v NullableKalturaSystemPartnerUsageItem) Get() *KalturaSystemPartnerUsageItem {
	return v.value
}

func (v *NullableKalturaSystemPartnerUsageItem) Set(val *KalturaSystemPartnerUsageItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSystemPartnerUsageItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSystemPartnerUsageItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSystemPartnerUsageItem(val *KalturaSystemPartnerUsageItem) *NullableKalturaSystemPartnerUsageItem {
	return &NullableKalturaSystemPartnerUsageItem{value: val, isSet: true}
}

func (v NullableKalturaSystemPartnerUsageItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSystemPartnerUsageItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



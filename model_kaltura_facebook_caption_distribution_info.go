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

// KalturaFacebookCaptionDistributionInfo struct for KalturaFacebookCaptionDistributionInfo
type KalturaFacebookCaptionDistributionInfo struct {
	AssetId *string `json:"assetId,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	Label *string `json:"label,omitempty"`
	Language *string `json:"language,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewKalturaFacebookCaptionDistributionInfo instantiates a new KalturaFacebookCaptionDistributionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFacebookCaptionDistributionInfo() *KalturaFacebookCaptionDistributionInfo {
	this := KalturaFacebookCaptionDistributionInfo{}
	return &this
}

// NewKalturaFacebookCaptionDistributionInfoWithDefaults instantiates a new KalturaFacebookCaptionDistributionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFacebookCaptionDistributionInfoWithDefaults() *KalturaFacebookCaptionDistributionInfo {
	this := KalturaFacebookCaptionDistributionInfo{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *KalturaFacebookCaptionDistributionInfo) SetAssetId(v string) {
	o.AssetId = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *KalturaFacebookCaptionDistributionInfo) SetFilePath(v string) {
	o.FilePath = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *KalturaFacebookCaptionDistributionInfo) SetLabel(v string) {
	o.Label = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *KalturaFacebookCaptionDistributionInfo) SetLanguage(v string) {
	o.Language = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *KalturaFacebookCaptionDistributionInfo) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KalturaFacebookCaptionDistributionInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFacebookCaptionDistributionInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KalturaFacebookCaptionDistributionInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KalturaFacebookCaptionDistributionInfo) SetVersion(v string) {
	o.Version = &v
}

func (o KalturaFacebookCaptionDistributionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.RemoteId != nil {
		toSerialize["remoteId"] = o.RemoteId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFacebookCaptionDistributionInfo struct {
	value *KalturaFacebookCaptionDistributionInfo
	isSet bool
}

func (v NullableKalturaFacebookCaptionDistributionInfo) Get() *KalturaFacebookCaptionDistributionInfo {
	return v.value
}

func (v *NullableKalturaFacebookCaptionDistributionInfo) Set(val *KalturaFacebookCaptionDistributionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFacebookCaptionDistributionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFacebookCaptionDistributionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFacebookCaptionDistributionInfo(val *KalturaFacebookCaptionDistributionInfo) *NullableKalturaFacebookCaptionDistributionInfo {
	return &NullableKalturaFacebookCaptionDistributionInfo{value: val, isSet: true}
}

func (v NullableKalturaFacebookCaptionDistributionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFacebookCaptionDistributionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


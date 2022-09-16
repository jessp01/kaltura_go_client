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

// KalturaDailymotionDistributionCaptionInfo struct for KalturaDailymotionDistributionCaptionInfo
type KalturaDailymotionDistributionCaptionInfo struct {
	// Enum Type: `KalturaDailymotionDistributionCaptionAction`
	Action *int32 `json:"action,omitempty"`
	AssetId *string `json:"assetId,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	// Enum Type: `KalturaDailymotionDistributionCaptionFormat`
	Format *int32 `json:"format,omitempty"`
	Language *string `json:"language,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewKalturaDailymotionDistributionCaptionInfo instantiates a new KalturaDailymotionDistributionCaptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDailymotionDistributionCaptionInfo() *KalturaDailymotionDistributionCaptionInfo {
	this := KalturaDailymotionDistributionCaptionInfo{}
	return &this
}

// NewKalturaDailymotionDistributionCaptionInfoWithDefaults instantiates a new KalturaDailymotionDistributionCaptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDailymotionDistributionCaptionInfoWithDefaults() *KalturaDailymotionDistributionCaptionInfo {
	this := KalturaDailymotionDistributionCaptionInfo{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetAction() int32 {
	if o == nil || o.Action == nil {
		var ret int32
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetActionOk() (*int32, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given int32 and assigns it to the Action field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetAction(v int32) {
	o.Action = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetAssetId(v string) {
	o.AssetId = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetFormat() int32 {
	if o == nil || o.Format == nil {
		var ret int32
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetFormatOk() (*int32, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given int32 and assigns it to the Format field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetFormat(v int32) {
	o.Format = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetLanguage(v string) {
	o.Language = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetRemoteId() string {
	if o == nil || o.RemoteId == nil {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetRemoteIdOk() (*string, bool) {
	if o == nil || o.RemoteId == nil {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasRemoteId() bool {
	if o != nil && o.RemoteId != nil {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KalturaDailymotionDistributionCaptionInfo) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KalturaDailymotionDistributionCaptionInfo) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *KalturaDailymotionDistributionCaptionInfo) SetVersion(v string) {
	o.Version = &v
}

func (o KalturaDailymotionDistributionCaptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
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

type NullableKalturaDailymotionDistributionCaptionInfo struct {
	value *KalturaDailymotionDistributionCaptionInfo
	isSet bool
}

func (v NullableKalturaDailymotionDistributionCaptionInfo) Get() *KalturaDailymotionDistributionCaptionInfo {
	return v.value
}

func (v *NullableKalturaDailymotionDistributionCaptionInfo) Set(val *KalturaDailymotionDistributionCaptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDailymotionDistributionCaptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDailymotionDistributionCaptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDailymotionDistributionCaptionInfo(val *KalturaDailymotionDistributionCaptionInfo) *NullableKalturaDailymotionDistributionCaptionInfo {
	return &NullableKalturaDailymotionDistributionCaptionInfo{value: val, isSet: true}
}

func (v NullableKalturaDailymotionDistributionCaptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDailymotionDistributionCaptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



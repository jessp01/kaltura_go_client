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

// KalturaAttUverseDistributionFile struct for KalturaAttUverseDistributionFile
type KalturaAttUverseDistributionFile struct {
	AssetId *string `json:"assetId,omitempty"`
	// Enum Type: `KalturaAssetType`
	AssetType *string `json:"assetType,omitempty"`
	LocalFilePath *string `json:"localFilePath,omitempty"`
	RemoteFilename *string `json:"remoteFilename,omitempty"`
}

// NewKalturaAttUverseDistributionFile instantiates a new KalturaAttUverseDistributionFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttUverseDistributionFile() *KalturaAttUverseDistributionFile {
	this := KalturaAttUverseDistributionFile{}
	return &this
}

// NewKalturaAttUverseDistributionFileWithDefaults instantiates a new KalturaAttUverseDistributionFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttUverseDistributionFileWithDefaults() *KalturaAttUverseDistributionFile {
	this := KalturaAttUverseDistributionFile{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *KalturaAttUverseDistributionFile) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAttUverseDistributionFile) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *KalturaAttUverseDistributionFile) HasAssetId() bool {
	if o != nil && o.AssetId != nil {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *KalturaAttUverseDistributionFile) SetAssetId(v string) {
	o.AssetId = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *KalturaAttUverseDistributionFile) GetAssetType() string {
	if o == nil || o.AssetType == nil {
		var ret string
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAttUverseDistributionFile) GetAssetTypeOk() (*string, bool) {
	if o == nil || o.AssetType == nil {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *KalturaAttUverseDistributionFile) HasAssetType() bool {
	if o != nil && o.AssetType != nil {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given string and assigns it to the AssetType field.
func (o *KalturaAttUverseDistributionFile) SetAssetType(v string) {
	o.AssetType = &v
}

// GetLocalFilePath returns the LocalFilePath field value if set, zero value otherwise.
func (o *KalturaAttUverseDistributionFile) GetLocalFilePath() string {
	if o == nil || o.LocalFilePath == nil {
		var ret string
		return ret
	}
	return *o.LocalFilePath
}

// GetLocalFilePathOk returns a tuple with the LocalFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAttUverseDistributionFile) GetLocalFilePathOk() (*string, bool) {
	if o == nil || o.LocalFilePath == nil {
		return nil, false
	}
	return o.LocalFilePath, true
}

// HasLocalFilePath returns a boolean if a field has been set.
func (o *KalturaAttUverseDistributionFile) HasLocalFilePath() bool {
	if o != nil && o.LocalFilePath != nil {
		return true
	}

	return false
}

// SetLocalFilePath gets a reference to the given string and assigns it to the LocalFilePath field.
func (o *KalturaAttUverseDistributionFile) SetLocalFilePath(v string) {
	o.LocalFilePath = &v
}

// GetRemoteFilename returns the RemoteFilename field value if set, zero value otherwise.
func (o *KalturaAttUverseDistributionFile) GetRemoteFilename() string {
	if o == nil || o.RemoteFilename == nil {
		var ret string
		return ret
	}
	return *o.RemoteFilename
}

// GetRemoteFilenameOk returns a tuple with the RemoteFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAttUverseDistributionFile) GetRemoteFilenameOk() (*string, bool) {
	if o == nil || o.RemoteFilename == nil {
		return nil, false
	}
	return o.RemoteFilename, true
}

// HasRemoteFilename returns a boolean if a field has been set.
func (o *KalturaAttUverseDistributionFile) HasRemoteFilename() bool {
	if o != nil && o.RemoteFilename != nil {
		return true
	}

	return false
}

// SetRemoteFilename gets a reference to the given string and assigns it to the RemoteFilename field.
func (o *KalturaAttUverseDistributionFile) SetRemoteFilename(v string) {
	o.RemoteFilename = &v
}

func (o KalturaAttUverseDistributionFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssetId != nil {
		toSerialize["assetId"] = o.AssetId
	}
	if o.AssetType != nil {
		toSerialize["assetType"] = o.AssetType
	}
	if o.LocalFilePath != nil {
		toSerialize["localFilePath"] = o.LocalFilePath
	}
	if o.RemoteFilename != nil {
		toSerialize["remoteFilename"] = o.RemoteFilename
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAttUverseDistributionFile struct {
	value *KalturaAttUverseDistributionFile
	isSet bool
}

func (v NullableKalturaAttUverseDistributionFile) Get() *KalturaAttUverseDistributionFile {
	return v.value
}

func (v *NullableKalturaAttUverseDistributionFile) Set(val *KalturaAttUverseDistributionFile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttUverseDistributionFile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttUverseDistributionFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttUverseDistributionFile(val *KalturaAttUverseDistributionFile) *NullableKalturaAttUverseDistributionFile {
	return &NullableKalturaAttUverseDistributionFile{value: val, isSet: true}
}

func (v NullableKalturaAttUverseDistributionFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttUverseDistributionFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



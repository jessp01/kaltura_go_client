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

// KalturaFileContainer struct for KalturaFileContainer
type KalturaFileContainer struct {
	EncryptionKey *string `json:"encryptionKey,omitempty"`
	FilePath *string `json:"filePath,omitempty"`
	FileSize *int32 `json:"fileSize,omitempty"`
}

// NewKalturaFileContainer instantiates a new KalturaFileContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFileContainer() *KalturaFileContainer {
	this := KalturaFileContainer{}
	return &this
}

// NewKalturaFileContainerWithDefaults instantiates a new KalturaFileContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFileContainerWithDefaults() *KalturaFileContainer {
	this := KalturaFileContainer{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *KalturaFileContainer) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileContainer) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *KalturaFileContainer) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *KalturaFileContainer) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *KalturaFileContainer) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileContainer) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *KalturaFileContainer) HasFilePath() bool {
	if o != nil && o.FilePath != nil {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *KalturaFileContainer) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *KalturaFileContainer) GetFileSize() int32 {
	if o == nil || o.FileSize == nil {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileContainer) GetFileSizeOk() (*int32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *KalturaFileContainer) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *KalturaFileContainer) SetFileSize(v int32) {
	o.FileSize = &v
}

func (o KalturaFileContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionKey != nil {
		toSerialize["encryptionKey"] = o.EncryptionKey
	}
	if o.FilePath != nil {
		toSerialize["filePath"] = o.FilePath
	}
	if o.FileSize != nil {
		toSerialize["fileSize"] = o.FileSize
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFileContainer struct {
	value *KalturaFileContainer
	isSet bool
}

func (v NullableKalturaFileContainer) Get() *KalturaFileContainer {
	return v.value
}

func (v *NullableKalturaFileContainer) Set(val *KalturaFileContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFileContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFileContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFileContainer(val *KalturaFileContainer) *NullableKalturaFileContainer {
	return &NullableKalturaFileContainer{value: val, isSet: true}
}

func (v NullableKalturaFileContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFileContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



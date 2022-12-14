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

// KalturaFeedDropFolderFile struct for KalturaFeedDropFolderFile
type KalturaFeedDropFolderFile struct {
	KalturaDropFolderFile
}

// NewKalturaFeedDropFolderFile instantiates a new KalturaFeedDropFolderFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFeedDropFolderFile() *KalturaFeedDropFolderFile {
	this := KalturaFeedDropFolderFile{}
	return &this
}

// NewKalturaFeedDropFolderFileWithDefaults instantiates a new KalturaFeedDropFolderFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFeedDropFolderFileWithDefaults() *KalturaFeedDropFolderFile {
	this := KalturaFeedDropFolderFile{}
	return &this
}

func (o KalturaFeedDropFolderFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDropFolderFile, errKalturaDropFolderFile := json.Marshal(o.KalturaDropFolderFile)
	if errKalturaDropFolderFile != nil {
		return []byte{}, errKalturaDropFolderFile
	}
	errKalturaDropFolderFile = json.Unmarshal([]byte(serializedKalturaDropFolderFile), &toSerialize)
	if errKalturaDropFolderFile != nil {
		return []byte{}, errKalturaDropFolderFile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFeedDropFolderFile struct {
	value *KalturaFeedDropFolderFile
	isSet bool
}

func (v NullableKalturaFeedDropFolderFile) Get() *KalturaFeedDropFolderFile {
	return v.value
}

func (v *NullableKalturaFeedDropFolderFile) Set(val *KalturaFeedDropFolderFile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFeedDropFolderFile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFeedDropFolderFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFeedDropFolderFile(val *KalturaFeedDropFolderFile) *NullableKalturaFeedDropFolderFile {
	return &NullableKalturaFeedDropFolderFile{value: val, isSet: true}
}

func (v NullableKalturaFeedDropFolderFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFeedDropFolderFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaApFeedDropFolder struct for KalturaApFeedDropFolder
type KalturaApFeedDropFolder struct {
	KalturaFeedDropFolder
}

// NewKalturaApFeedDropFolder instantiates a new KalturaApFeedDropFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaApFeedDropFolder() *KalturaApFeedDropFolder {
	this := KalturaApFeedDropFolder{}
	return &this
}

// NewKalturaApFeedDropFolderWithDefaults instantiates a new KalturaApFeedDropFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaApFeedDropFolderWithDefaults() *KalturaApFeedDropFolder {
	this := KalturaApFeedDropFolder{}
	return &this
}

func (o KalturaApFeedDropFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFeedDropFolder, errKalturaFeedDropFolder := json.Marshal(o.KalturaFeedDropFolder)
	if errKalturaFeedDropFolder != nil {
		return []byte{}, errKalturaFeedDropFolder
	}
	errKalturaFeedDropFolder = json.Unmarshal([]byte(serializedKalturaFeedDropFolder), &toSerialize)
	if errKalturaFeedDropFolder != nil {
		return []byte{}, errKalturaFeedDropFolder
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaApFeedDropFolder struct {
	value *KalturaApFeedDropFolder
	isSet bool
}

func (v NullableKalturaApFeedDropFolder) Get() *KalturaApFeedDropFolder {
	return v.value
}

func (v *NullableKalturaApFeedDropFolder) Set(val *KalturaApFeedDropFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaApFeedDropFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaApFeedDropFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaApFeedDropFolder(val *KalturaApFeedDropFolder) *NullableKalturaApFeedDropFolder {
	return &NullableKalturaApFeedDropFolder{value: val, isSet: true}
}

func (v NullableKalturaApFeedDropFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaApFeedDropFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



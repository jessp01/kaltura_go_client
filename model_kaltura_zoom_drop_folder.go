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

// KalturaZoomDropFolder struct for KalturaZoomDropFolder
type KalturaZoomDropFolder struct {
	KalturaDropFolder
}

// NewKalturaZoomDropFolder instantiates a new KalturaZoomDropFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaZoomDropFolder() *KalturaZoomDropFolder {
	this := KalturaZoomDropFolder{}
	return &this
}

// NewKalturaZoomDropFolderWithDefaults instantiates a new KalturaZoomDropFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaZoomDropFolderWithDefaults() *KalturaZoomDropFolder {
	this := KalturaZoomDropFolder{}
	return &this
}

func (o KalturaZoomDropFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDropFolder, errKalturaDropFolder := json.Marshal(o.KalturaDropFolder)
	if errKalturaDropFolder != nil {
		return []byte{}, errKalturaDropFolder
	}
	errKalturaDropFolder = json.Unmarshal([]byte(serializedKalturaDropFolder), &toSerialize)
	if errKalturaDropFolder != nil {
		return []byte{}, errKalturaDropFolder
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaZoomDropFolder struct {
	value *KalturaZoomDropFolder
	isSet bool
}

func (v NullableKalturaZoomDropFolder) Get() *KalturaZoomDropFolder {
	return v.value
}

func (v *NullableKalturaZoomDropFolder) Set(val *KalturaZoomDropFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaZoomDropFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaZoomDropFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaZoomDropFolder(val *KalturaZoomDropFolder) *NullableKalturaZoomDropFolder {
	return &NullableKalturaZoomDropFolder{value: val, isSet: true}
}

func (v NullableKalturaZoomDropFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaZoomDropFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



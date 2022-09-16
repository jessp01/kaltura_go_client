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

// KalturaDropFolderImportJobData struct for KalturaDropFolderImportJobData
type KalturaDropFolderImportJobData struct {
	KalturaSshImportJobData
}

// NewKalturaDropFolderImportJobData instantiates a new KalturaDropFolderImportJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDropFolderImportJobData() *KalturaDropFolderImportJobData {
	this := KalturaDropFolderImportJobData{}
	return &this
}

// NewKalturaDropFolderImportJobDataWithDefaults instantiates a new KalturaDropFolderImportJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDropFolderImportJobDataWithDefaults() *KalturaDropFolderImportJobData {
	this := KalturaDropFolderImportJobData{}
	return &this
}

func (o KalturaDropFolderImportJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSshImportJobData, errKalturaSshImportJobData := json.Marshal(o.KalturaSshImportJobData)
	if errKalturaSshImportJobData != nil {
		return []byte{}, errKalturaSshImportJobData
	}
	errKalturaSshImportJobData = json.Unmarshal([]byte(serializedKalturaSshImportJobData), &toSerialize)
	if errKalturaSshImportJobData != nil {
		return []byte{}, errKalturaSshImportJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDropFolderImportJobData struct {
	value *KalturaDropFolderImportJobData
	isSet bool
}

func (v NullableKalturaDropFolderImportJobData) Get() *KalturaDropFolderImportJobData {
	return v.value
}

func (v *NullableKalturaDropFolderImportJobData) Set(val *KalturaDropFolderImportJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDropFolderImportJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDropFolderImportJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDropFolderImportJobData(val *KalturaDropFolderImportJobData) *NullableKalturaDropFolderImportJobData {
	return &NullableKalturaDropFolderImportJobData{value: val, isSet: true}
}

func (v NullableKalturaDropFolderImportJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDropFolderImportJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



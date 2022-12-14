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

// KalturaBulkUploadCategoryUserData This class represents object-specific data passed to the   bulk upload job.
type KalturaBulkUploadCategoryUserData struct {
	KalturaBulkUploadObjectData
}

// NewKalturaBulkUploadCategoryUserData instantiates a new KalturaBulkUploadCategoryUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadCategoryUserData() *KalturaBulkUploadCategoryUserData {
	this := KalturaBulkUploadCategoryUserData{}
	return &this
}

// NewKalturaBulkUploadCategoryUserDataWithDefaults instantiates a new KalturaBulkUploadCategoryUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadCategoryUserDataWithDefaults() *KalturaBulkUploadCategoryUserData {
	this := KalturaBulkUploadCategoryUserData{}
	return &this
}

func (o KalturaBulkUploadCategoryUserData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBulkUploadObjectData, errKalturaBulkUploadObjectData := json.Marshal(o.KalturaBulkUploadObjectData)
	if errKalturaBulkUploadObjectData != nil {
		return []byte{}, errKalturaBulkUploadObjectData
	}
	errKalturaBulkUploadObjectData = json.Unmarshal([]byte(serializedKalturaBulkUploadObjectData), &toSerialize)
	if errKalturaBulkUploadObjectData != nil {
		return []byte{}, errKalturaBulkUploadObjectData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBulkUploadCategoryUserData struct {
	value *KalturaBulkUploadCategoryUserData
	isSet bool
}

func (v NullableKalturaBulkUploadCategoryUserData) Get() *KalturaBulkUploadCategoryUserData {
	return v.value
}

func (v *NullableKalturaBulkUploadCategoryUserData) Set(val *KalturaBulkUploadCategoryUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadCategoryUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadCategoryUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadCategoryUserData(val *KalturaBulkUploadCategoryUserData) *NullableKalturaBulkUploadCategoryUserData {
	return &NullableKalturaBulkUploadCategoryUserData{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadCategoryUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadCategoryUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



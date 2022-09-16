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

// KalturaStorageDeleteJobData struct for KalturaStorageDeleteJobData
type KalturaStorageDeleteJobData struct {
	KalturaStorageJobData
}

// NewKalturaStorageDeleteJobData instantiates a new KalturaStorageDeleteJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaStorageDeleteJobData() *KalturaStorageDeleteJobData {
	this := KalturaStorageDeleteJobData{}
	return &this
}

// NewKalturaStorageDeleteJobDataWithDefaults instantiates a new KalturaStorageDeleteJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaStorageDeleteJobDataWithDefaults() *KalturaStorageDeleteJobData {
	this := KalturaStorageDeleteJobData{}
	return &this
}

func (o KalturaStorageDeleteJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaStorageJobData, errKalturaStorageJobData := json.Marshal(o.KalturaStorageJobData)
	if errKalturaStorageJobData != nil {
		return []byte{}, errKalturaStorageJobData
	}
	errKalturaStorageJobData = json.Unmarshal([]byte(serializedKalturaStorageJobData), &toSerialize)
	if errKalturaStorageJobData != nil {
		return []byte{}, errKalturaStorageJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaStorageDeleteJobData struct {
	value *KalturaStorageDeleteJobData
	isSet bool
}

func (v NullableKalturaStorageDeleteJobData) Get() *KalturaStorageDeleteJobData {
	return v.value
}

func (v *NullableKalturaStorageDeleteJobData) Set(val *KalturaStorageDeleteJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaStorageDeleteJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaStorageDeleteJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaStorageDeleteJobData(val *KalturaStorageDeleteJobData) *NullableKalturaStorageDeleteJobData {
	return &NullableKalturaStorageDeleteJobData{value: val, isSet: true}
}

func (v NullableKalturaStorageDeleteJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaStorageDeleteJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



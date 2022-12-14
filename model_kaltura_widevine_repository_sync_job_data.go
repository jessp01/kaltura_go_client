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

// KalturaWidevineRepositorySyncJobData struct for KalturaWidevineRepositorySyncJobData
type KalturaWidevineRepositorySyncJobData struct {
	KalturaJobData
}

// NewKalturaWidevineRepositorySyncJobData instantiates a new KalturaWidevineRepositorySyncJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWidevineRepositorySyncJobData() *KalturaWidevineRepositorySyncJobData {
	this := KalturaWidevineRepositorySyncJobData{}
	return &this
}

// NewKalturaWidevineRepositorySyncJobDataWithDefaults instantiates a new KalturaWidevineRepositorySyncJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWidevineRepositorySyncJobDataWithDefaults() *KalturaWidevineRepositorySyncJobData {
	this := KalturaWidevineRepositorySyncJobData{}
	return &this
}

func (o KalturaWidevineRepositorySyncJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaJobData, errKalturaJobData := json.Marshal(o.KalturaJobData)
	if errKalturaJobData != nil {
		return []byte{}, errKalturaJobData
	}
	errKalturaJobData = json.Unmarshal([]byte(serializedKalturaJobData), &toSerialize)
	if errKalturaJobData != nil {
		return []byte{}, errKalturaJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWidevineRepositorySyncJobData struct {
	value *KalturaWidevineRepositorySyncJobData
	isSet bool
}

func (v NullableKalturaWidevineRepositorySyncJobData) Get() *KalturaWidevineRepositorySyncJobData {
	return v.value
}

func (v *NullableKalturaWidevineRepositorySyncJobData) Set(val *KalturaWidevineRepositorySyncJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWidevineRepositorySyncJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWidevineRepositorySyncJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWidevineRepositorySyncJobData(val *KalturaWidevineRepositorySyncJobData) *NullableKalturaWidevineRepositorySyncJobData {
	return &NullableKalturaWidevineRepositorySyncJobData{value: val, isSet: true}
}

func (v NullableKalturaWidevineRepositorySyncJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWidevineRepositorySyncJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



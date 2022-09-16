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

// KalturaIndexJobData struct for KalturaIndexJobData
type KalturaIndexJobData struct {
	KalturaJobData
}

// NewKalturaIndexJobData instantiates a new KalturaIndexJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaIndexJobData() *KalturaIndexJobData {
	this := KalturaIndexJobData{}
	return &this
}

// NewKalturaIndexJobDataWithDefaults instantiates a new KalturaIndexJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaIndexJobDataWithDefaults() *KalturaIndexJobData {
	this := KalturaIndexJobData{}
	return &this
}

func (o KalturaIndexJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaIndexJobData struct {
	value *KalturaIndexJobData
	isSet bool
}

func (v NullableKalturaIndexJobData) Get() *KalturaIndexJobData {
	return v.value
}

func (v *NullableKalturaIndexJobData) Set(val *KalturaIndexJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaIndexJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaIndexJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaIndexJobData(val *KalturaIndexJobData) *NullableKalturaIndexJobData {
	return &NullableKalturaIndexJobData{value: val, isSet: true}
}

func (v NullableKalturaIndexJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaIndexJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


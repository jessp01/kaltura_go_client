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

// KalturaClipConcatJobData struct for KalturaClipConcatJobData
type KalturaClipConcatJobData struct {
	KalturaJobData
}

// NewKalturaClipConcatJobData instantiates a new KalturaClipConcatJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaClipConcatJobData() *KalturaClipConcatJobData {
	this := KalturaClipConcatJobData{}
	return &this
}

// NewKalturaClipConcatJobDataWithDefaults instantiates a new KalturaClipConcatJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaClipConcatJobDataWithDefaults() *KalturaClipConcatJobData {
	this := KalturaClipConcatJobData{}
	return &this
}

func (o KalturaClipConcatJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaClipConcatJobData struct {
	value *KalturaClipConcatJobData
	isSet bool
}

func (v NullableKalturaClipConcatJobData) Get() *KalturaClipConcatJobData {
	return v.value
}

func (v *NullableKalturaClipConcatJobData) Set(val *KalturaClipConcatJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaClipConcatJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaClipConcatJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaClipConcatJobData(val *KalturaClipConcatJobData) *NullableKalturaClipConcatJobData {
	return &NullableKalturaClipConcatJobData{value: val, isSet: true}
}

func (v NullableKalturaClipConcatJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaClipConcatJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



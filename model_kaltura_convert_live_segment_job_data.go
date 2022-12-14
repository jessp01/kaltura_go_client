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

// KalturaConvertLiveSegmentJobData struct for KalturaConvertLiveSegmentJobData
type KalturaConvertLiveSegmentJobData struct {
	KalturaJobData
}

// NewKalturaConvertLiveSegmentJobData instantiates a new KalturaConvertLiveSegmentJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConvertLiveSegmentJobData() *KalturaConvertLiveSegmentJobData {
	this := KalturaConvertLiveSegmentJobData{}
	return &this
}

// NewKalturaConvertLiveSegmentJobDataWithDefaults instantiates a new KalturaConvertLiveSegmentJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConvertLiveSegmentJobDataWithDefaults() *KalturaConvertLiveSegmentJobData {
	this := KalturaConvertLiveSegmentJobData{}
	return &this
}

func (o KalturaConvertLiveSegmentJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaConvertLiveSegmentJobData struct {
	value *KalturaConvertLiveSegmentJobData
	isSet bool
}

func (v NullableKalturaConvertLiveSegmentJobData) Get() *KalturaConvertLiveSegmentJobData {
	return v.value
}

func (v *NullableKalturaConvertLiveSegmentJobData) Set(val *KalturaConvertLiveSegmentJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConvertLiveSegmentJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConvertLiveSegmentJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConvertLiveSegmentJobData(val *KalturaConvertLiveSegmentJobData) *NullableKalturaConvertLiveSegmentJobData {
	return &NullableKalturaConvertLiveSegmentJobData{value: val, isSet: true}
}

func (v NullableKalturaConvertLiveSegmentJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConvertLiveSegmentJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



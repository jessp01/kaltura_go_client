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

// KalturaCopyJobData struct for KalturaCopyJobData
type KalturaCopyJobData struct {
	KalturaJobData
}

// NewKalturaCopyJobData instantiates a new KalturaCopyJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCopyJobData() *KalturaCopyJobData {
	this := KalturaCopyJobData{}
	return &this
}

// NewKalturaCopyJobDataWithDefaults instantiates a new KalturaCopyJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCopyJobDataWithDefaults() *KalturaCopyJobData {
	this := KalturaCopyJobData{}
	return &this
}

func (o KalturaCopyJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaCopyJobData struct {
	value *KalturaCopyJobData
	isSet bool
}

func (v NullableKalturaCopyJobData) Get() *KalturaCopyJobData {
	return v.value
}

func (v *NullableKalturaCopyJobData) Set(val *KalturaCopyJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCopyJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCopyJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCopyJobData(val *KalturaCopyJobData) *NullableKalturaCopyJobData {
	return &NullableKalturaCopyJobData{value: val, isSet: true}
}

func (v NullableKalturaCopyJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCopyJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaUsersCsvJobData struct for KalturaUsersCsvJobData
type KalturaUsersCsvJobData struct {
	KalturaMappedObjectsCsvJobData
}

// NewKalturaUsersCsvJobData instantiates a new KalturaUsersCsvJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUsersCsvJobData() *KalturaUsersCsvJobData {
	this := KalturaUsersCsvJobData{}
	return &this
}

// NewKalturaUsersCsvJobDataWithDefaults instantiates a new KalturaUsersCsvJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUsersCsvJobDataWithDefaults() *KalturaUsersCsvJobData {
	this := KalturaUsersCsvJobData{}
	return &this
}

func (o KalturaUsersCsvJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMappedObjectsCsvJobData, errKalturaMappedObjectsCsvJobData := json.Marshal(o.KalturaMappedObjectsCsvJobData)
	if errKalturaMappedObjectsCsvJobData != nil {
		return []byte{}, errKalturaMappedObjectsCsvJobData
	}
	errKalturaMappedObjectsCsvJobData = json.Unmarshal([]byte(serializedKalturaMappedObjectsCsvJobData), &toSerialize)
	if errKalturaMappedObjectsCsvJobData != nil {
		return []byte{}, errKalturaMappedObjectsCsvJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUsersCsvJobData struct {
	value *KalturaUsersCsvJobData
	isSet bool
}

func (v NullableKalturaUsersCsvJobData) Get() *KalturaUsersCsvJobData {
	return v.value
}

func (v *NullableKalturaUsersCsvJobData) Set(val *KalturaUsersCsvJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUsersCsvJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUsersCsvJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUsersCsvJobData(val *KalturaUsersCsvJobData) *NullableKalturaUsersCsvJobData {
	return &NullableKalturaUsersCsvJobData{value: val, isSet: true}
}

func (v NullableKalturaUsersCsvJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUsersCsvJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



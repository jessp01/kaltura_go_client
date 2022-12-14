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

// KalturaDistributionUpdateJobData struct for KalturaDistributionUpdateJobData
type KalturaDistributionUpdateJobData struct {
	KalturaDistributionJobData
}

// NewKalturaDistributionUpdateJobData instantiates a new KalturaDistributionUpdateJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDistributionUpdateJobData() *KalturaDistributionUpdateJobData {
	this := KalturaDistributionUpdateJobData{}
	return &this
}

// NewKalturaDistributionUpdateJobDataWithDefaults instantiates a new KalturaDistributionUpdateJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDistributionUpdateJobDataWithDefaults() *KalturaDistributionUpdateJobData {
	this := KalturaDistributionUpdateJobData{}
	return &this
}

func (o KalturaDistributionUpdateJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionJobData, errKalturaDistributionJobData := json.Marshal(o.KalturaDistributionJobData)
	if errKalturaDistributionJobData != nil {
		return []byte{}, errKalturaDistributionJobData
	}
	errKalturaDistributionJobData = json.Unmarshal([]byte(serializedKalturaDistributionJobData), &toSerialize)
	if errKalturaDistributionJobData != nil {
		return []byte{}, errKalturaDistributionJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDistributionUpdateJobData struct {
	value *KalturaDistributionUpdateJobData
	isSet bool
}

func (v NullableKalturaDistributionUpdateJobData) Get() *KalturaDistributionUpdateJobData {
	return v.value
}

func (v *NullableKalturaDistributionUpdateJobData) Set(val *KalturaDistributionUpdateJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDistributionUpdateJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDistributionUpdateJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDistributionUpdateJobData(val *KalturaDistributionUpdateJobData) *NullableKalturaDistributionUpdateJobData {
	return &NullableKalturaDistributionUpdateJobData{value: val, isSet: true}
}

func (v NullableKalturaDistributionUpdateJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDistributionUpdateJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



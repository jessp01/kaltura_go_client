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

// KalturaConvertCollectionJobData struct for KalturaConvertCollectionJobData
type KalturaConvertCollectionJobData struct {
	KalturaConvartableJobData
}

// NewKalturaConvertCollectionJobData instantiates a new KalturaConvertCollectionJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConvertCollectionJobData() *KalturaConvertCollectionJobData {
	this := KalturaConvertCollectionJobData{}
	return &this
}

// NewKalturaConvertCollectionJobDataWithDefaults instantiates a new KalturaConvertCollectionJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConvertCollectionJobDataWithDefaults() *KalturaConvertCollectionJobData {
	this := KalturaConvertCollectionJobData{}
	return &this
}

func (o KalturaConvertCollectionJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConvartableJobData, errKalturaConvartableJobData := json.Marshal(o.KalturaConvartableJobData)
	if errKalturaConvartableJobData != nil {
		return []byte{}, errKalturaConvartableJobData
	}
	errKalturaConvartableJobData = json.Unmarshal([]byte(serializedKalturaConvartableJobData), &toSerialize)
	if errKalturaConvartableJobData != nil {
		return []byte{}, errKalturaConvartableJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaConvertCollectionJobData struct {
	value *KalturaConvertCollectionJobData
	isSet bool
}

func (v NullableKalturaConvertCollectionJobData) Get() *KalturaConvertCollectionJobData {
	return v.value
}

func (v *NullableKalturaConvertCollectionJobData) Set(val *KalturaConvertCollectionJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConvertCollectionJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConvertCollectionJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConvertCollectionJobData(val *KalturaConvertCollectionJobData) *NullableKalturaConvertCollectionJobData {
	return &NullableKalturaConvertCollectionJobData{value: val, isSet: true}
}

func (v NullableKalturaConvertCollectionJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConvertCollectionJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



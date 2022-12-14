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

// KalturaDistributeObjectTask struct for KalturaDistributeObjectTask
type KalturaDistributeObjectTask struct {
	KalturaObjectTask
}

// NewKalturaDistributeObjectTask instantiates a new KalturaDistributeObjectTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDistributeObjectTask() *KalturaDistributeObjectTask {
	this := KalturaDistributeObjectTask{}
	return &this
}

// NewKalturaDistributeObjectTaskWithDefaults instantiates a new KalturaDistributeObjectTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDistributeObjectTaskWithDefaults() *KalturaDistributeObjectTask {
	this := KalturaDistributeObjectTask{}
	return &this
}

func (o KalturaDistributeObjectTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaObjectTask, errKalturaObjectTask := json.Marshal(o.KalturaObjectTask)
	if errKalturaObjectTask != nil {
		return []byte{}, errKalturaObjectTask
	}
	errKalturaObjectTask = json.Unmarshal([]byte(serializedKalturaObjectTask), &toSerialize)
	if errKalturaObjectTask != nil {
		return []byte{}, errKalturaObjectTask
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDistributeObjectTask struct {
	value *KalturaDistributeObjectTask
	isSet bool
}

func (v NullableKalturaDistributeObjectTask) Get() *KalturaDistributeObjectTask {
	return v.value
}

func (v *NullableKalturaDistributeObjectTask) Set(val *KalturaDistributeObjectTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDistributeObjectTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDistributeObjectTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDistributeObjectTask(val *KalturaDistributeObjectTask) *NullableKalturaDistributeObjectTask {
	return &NullableKalturaDistributeObjectTask{value: val, isSet: true}
}

func (v NullableKalturaDistributeObjectTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDistributeObjectTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



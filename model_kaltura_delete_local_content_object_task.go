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

// KalturaDeleteLocalContentObjectTask struct for KalturaDeleteLocalContentObjectTask
type KalturaDeleteLocalContentObjectTask struct {
	KalturaObjectTask
}

// NewKalturaDeleteLocalContentObjectTask instantiates a new KalturaDeleteLocalContentObjectTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeleteLocalContentObjectTask() *KalturaDeleteLocalContentObjectTask {
	this := KalturaDeleteLocalContentObjectTask{}
	return &this
}

// NewKalturaDeleteLocalContentObjectTaskWithDefaults instantiates a new KalturaDeleteLocalContentObjectTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeleteLocalContentObjectTaskWithDefaults() *KalturaDeleteLocalContentObjectTask {
	this := KalturaDeleteLocalContentObjectTask{}
	return &this
}

func (o KalturaDeleteLocalContentObjectTask) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDeleteLocalContentObjectTask struct {
	value *KalturaDeleteLocalContentObjectTask
	isSet bool
}

func (v NullableKalturaDeleteLocalContentObjectTask) Get() *KalturaDeleteLocalContentObjectTask {
	return v.value
}

func (v *NullableKalturaDeleteLocalContentObjectTask) Set(val *KalturaDeleteLocalContentObjectTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeleteLocalContentObjectTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeleteLocalContentObjectTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeleteLocalContentObjectTask(val *KalturaDeleteLocalContentObjectTask) *NullableKalturaDeleteLocalContentObjectTask {
	return &NullableKalturaDeleteLocalContentObjectTask{value: val, isSet: true}
}

func (v NullableKalturaDeleteLocalContentObjectTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeleteLocalContentObjectTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



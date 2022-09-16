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

// KalturaModifyEntryObjectTask struct for KalturaModifyEntryObjectTask
type KalturaModifyEntryObjectTask struct {
	KalturaObjectTask
}

// NewKalturaModifyEntryObjectTask instantiates a new KalturaModifyEntryObjectTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaModifyEntryObjectTask() *KalturaModifyEntryObjectTask {
	this := KalturaModifyEntryObjectTask{}
	return &this
}

// NewKalturaModifyEntryObjectTaskWithDefaults instantiates a new KalturaModifyEntryObjectTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaModifyEntryObjectTaskWithDefaults() *KalturaModifyEntryObjectTask {
	this := KalturaModifyEntryObjectTask{}
	return &this
}

func (o KalturaModifyEntryObjectTask) MarshalJSON() ([]byte, error) {
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

type NullableKalturaModifyEntryObjectTask struct {
	value *KalturaModifyEntryObjectTask
	isSet bool
}

func (v NullableKalturaModifyEntryObjectTask) Get() *KalturaModifyEntryObjectTask {
	return v.value
}

func (v *NullableKalturaModifyEntryObjectTask) Set(val *KalturaModifyEntryObjectTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaModifyEntryObjectTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaModifyEntryObjectTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaModifyEntryObjectTask(val *KalturaModifyEntryObjectTask) *NullableKalturaModifyEntryObjectTask {
	return &NullableKalturaModifyEntryObjectTask{value: val, isSet: true}
}

func (v NullableKalturaModifyEntryObjectTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaModifyEntryObjectTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



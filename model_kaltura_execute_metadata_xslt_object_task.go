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

// KalturaExecuteMetadataXsltObjectTask struct for KalturaExecuteMetadataXsltObjectTask
type KalturaExecuteMetadataXsltObjectTask struct {
	KalturaObjectTask
}

// NewKalturaExecuteMetadataXsltObjectTask instantiates a new KalturaExecuteMetadataXsltObjectTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaExecuteMetadataXsltObjectTask() *KalturaExecuteMetadataXsltObjectTask {
	this := KalturaExecuteMetadataXsltObjectTask{}
	return &this
}

// NewKalturaExecuteMetadataXsltObjectTaskWithDefaults instantiates a new KalturaExecuteMetadataXsltObjectTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaExecuteMetadataXsltObjectTaskWithDefaults() *KalturaExecuteMetadataXsltObjectTask {
	this := KalturaExecuteMetadataXsltObjectTask{}
	return &this
}

func (o KalturaExecuteMetadataXsltObjectTask) MarshalJSON() ([]byte, error) {
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

type NullableKalturaExecuteMetadataXsltObjectTask struct {
	value *KalturaExecuteMetadataXsltObjectTask
	isSet bool
}

func (v NullableKalturaExecuteMetadataXsltObjectTask) Get() *KalturaExecuteMetadataXsltObjectTask {
	return v.value
}

func (v *NullableKalturaExecuteMetadataXsltObjectTask) Set(val *KalturaExecuteMetadataXsltObjectTask) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaExecuteMetadataXsltObjectTask) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaExecuteMetadataXsltObjectTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaExecuteMetadataXsltObjectTask(val *KalturaExecuteMetadataXsltObjectTask) *NullableKalturaExecuteMetadataXsltObjectTask {
	return &NullableKalturaExecuteMetadataXsltObjectTask{value: val, isSet: true}
}

func (v NullableKalturaExecuteMetadataXsltObjectTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaExecuteMetadataXsltObjectTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



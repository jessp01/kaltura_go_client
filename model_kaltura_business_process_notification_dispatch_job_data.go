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

// KalturaBusinessProcessNotificationDispatchJobData struct for KalturaBusinessProcessNotificationDispatchJobData
type KalturaBusinessProcessNotificationDispatchJobData struct {
	KalturaEventNotificationDispatchJobData
}

// NewKalturaBusinessProcessNotificationDispatchJobData instantiates a new KalturaBusinessProcessNotificationDispatchJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessNotificationDispatchJobData() *KalturaBusinessProcessNotificationDispatchJobData {
	this := KalturaBusinessProcessNotificationDispatchJobData{}
	return &this
}

// NewKalturaBusinessProcessNotificationDispatchJobDataWithDefaults instantiates a new KalturaBusinessProcessNotificationDispatchJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessNotificationDispatchJobDataWithDefaults() *KalturaBusinessProcessNotificationDispatchJobData {
	this := KalturaBusinessProcessNotificationDispatchJobData{}
	return &this
}

func (o KalturaBusinessProcessNotificationDispatchJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEventNotificationDispatchJobData, errKalturaEventNotificationDispatchJobData := json.Marshal(o.KalturaEventNotificationDispatchJobData)
	if errKalturaEventNotificationDispatchJobData != nil {
		return []byte{}, errKalturaEventNotificationDispatchJobData
	}
	errKalturaEventNotificationDispatchJobData = json.Unmarshal([]byte(serializedKalturaEventNotificationDispatchJobData), &toSerialize)
	if errKalturaEventNotificationDispatchJobData != nil {
		return []byte{}, errKalturaEventNotificationDispatchJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBusinessProcessNotificationDispatchJobData struct {
	value *KalturaBusinessProcessNotificationDispatchJobData
	isSet bool
}

func (v NullableKalturaBusinessProcessNotificationDispatchJobData) Get() *KalturaBusinessProcessNotificationDispatchJobData {
	return v.value
}

func (v *NullableKalturaBusinessProcessNotificationDispatchJobData) Set(val *KalturaBusinessProcessNotificationDispatchJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessNotificationDispatchJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessNotificationDispatchJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessNotificationDispatchJobData(val *KalturaBusinessProcessNotificationDispatchJobData) *NullableKalturaBusinessProcessNotificationDispatchJobData {
	return &NullableKalturaBusinessProcessNotificationDispatchJobData{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessNotificationDispatchJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessNotificationDispatchJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



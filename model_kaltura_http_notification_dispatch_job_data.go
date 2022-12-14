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

// KalturaHttpNotificationDispatchJobData struct for KalturaHttpNotificationDispatchJobData
type KalturaHttpNotificationDispatchJobData struct {
	KalturaEventNotificationDispatchJobData
}

// NewKalturaHttpNotificationDispatchJobData instantiates a new KalturaHttpNotificationDispatchJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHttpNotificationDispatchJobData() *KalturaHttpNotificationDispatchJobData {
	this := KalturaHttpNotificationDispatchJobData{}
	return &this
}

// NewKalturaHttpNotificationDispatchJobDataWithDefaults instantiates a new KalturaHttpNotificationDispatchJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHttpNotificationDispatchJobDataWithDefaults() *KalturaHttpNotificationDispatchJobData {
	this := KalturaHttpNotificationDispatchJobData{}
	return &this
}

func (o KalturaHttpNotificationDispatchJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaHttpNotificationDispatchJobData struct {
	value *KalturaHttpNotificationDispatchJobData
	isSet bool
}

func (v NullableKalturaHttpNotificationDispatchJobData) Get() *KalturaHttpNotificationDispatchJobData {
	return v.value
}

func (v *NullableKalturaHttpNotificationDispatchJobData) Set(val *KalturaHttpNotificationDispatchJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHttpNotificationDispatchJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHttpNotificationDispatchJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHttpNotificationDispatchJobData(val *KalturaHttpNotificationDispatchJobData) *NullableKalturaHttpNotificationDispatchJobData {
	return &NullableKalturaHttpNotificationDispatchJobData{value: val, isSet: true}
}

func (v NullableKalturaHttpNotificationDispatchJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHttpNotificationDispatchJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



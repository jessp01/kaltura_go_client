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

// KalturaNotificationJobData struct for KalturaNotificationJobData
type KalturaNotificationJobData struct {
	KalturaJobData
}

// NewKalturaNotificationJobData instantiates a new KalturaNotificationJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaNotificationJobData() *KalturaNotificationJobData {
	this := KalturaNotificationJobData{}
	return &this
}

// NewKalturaNotificationJobDataWithDefaults instantiates a new KalturaNotificationJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaNotificationJobDataWithDefaults() *KalturaNotificationJobData {
	this := KalturaNotificationJobData{}
	return &this
}

func (o KalturaNotificationJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaNotificationJobData struct {
	value *KalturaNotificationJobData
	isSet bool
}

func (v NullableKalturaNotificationJobData) Get() *KalturaNotificationJobData {
	return v.value
}

func (v *NullableKalturaNotificationJobData) Set(val *KalturaNotificationJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaNotificationJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaNotificationJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaNotificationJobData(val *KalturaNotificationJobData) *NullableKalturaNotificationJobData {
	return &NullableKalturaNotificationJobData{value: val, isSet: true}
}

func (v NullableKalturaNotificationJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaNotificationJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



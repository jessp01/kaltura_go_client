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

// KalturaEmailNotificationStaticRecipientJobData JobData representing the static receipient array
type KalturaEmailNotificationStaticRecipientJobData struct {
	KalturaEmailNotificationRecipientJobData
}

// NewKalturaEmailNotificationStaticRecipientJobData instantiates a new KalturaEmailNotificationStaticRecipientJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEmailNotificationStaticRecipientJobData() *KalturaEmailNotificationStaticRecipientJobData {
	this := KalturaEmailNotificationStaticRecipientJobData{}
	return &this
}

// NewKalturaEmailNotificationStaticRecipientJobDataWithDefaults instantiates a new KalturaEmailNotificationStaticRecipientJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEmailNotificationStaticRecipientJobDataWithDefaults() *KalturaEmailNotificationStaticRecipientJobData {
	this := KalturaEmailNotificationStaticRecipientJobData{}
	return &this
}

func (o KalturaEmailNotificationStaticRecipientJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEmailNotificationRecipientJobData, errKalturaEmailNotificationRecipientJobData := json.Marshal(o.KalturaEmailNotificationRecipientJobData)
	if errKalturaEmailNotificationRecipientJobData != nil {
		return []byte{}, errKalturaEmailNotificationRecipientJobData
	}
	errKalturaEmailNotificationRecipientJobData = json.Unmarshal([]byte(serializedKalturaEmailNotificationRecipientJobData), &toSerialize)
	if errKalturaEmailNotificationRecipientJobData != nil {
		return []byte{}, errKalturaEmailNotificationRecipientJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEmailNotificationStaticRecipientJobData struct {
	value *KalturaEmailNotificationStaticRecipientJobData
	isSet bool
}

func (v NullableKalturaEmailNotificationStaticRecipientJobData) Get() *KalturaEmailNotificationStaticRecipientJobData {
	return v.value
}

func (v *NullableKalturaEmailNotificationStaticRecipientJobData) Set(val *KalturaEmailNotificationStaticRecipientJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEmailNotificationStaticRecipientJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEmailNotificationStaticRecipientJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEmailNotificationStaticRecipientJobData(val *KalturaEmailNotificationStaticRecipientJobData) *NullableKalturaEmailNotificationStaticRecipientJobData {
	return &NullableKalturaEmailNotificationStaticRecipientJobData{value: val, isSet: true}
}

func (v NullableKalturaEmailNotificationStaticRecipientJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEmailNotificationStaticRecipientJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


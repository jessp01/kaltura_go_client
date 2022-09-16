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

// KalturaIndexTagsByPrivacyContextJobData struct for KalturaIndexTagsByPrivacyContextJobData
type KalturaIndexTagsByPrivacyContextJobData struct {
	KalturaJobData
}

// NewKalturaIndexTagsByPrivacyContextJobData instantiates a new KalturaIndexTagsByPrivacyContextJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaIndexTagsByPrivacyContextJobData() *KalturaIndexTagsByPrivacyContextJobData {
	this := KalturaIndexTagsByPrivacyContextJobData{}
	return &this
}

// NewKalturaIndexTagsByPrivacyContextJobDataWithDefaults instantiates a new KalturaIndexTagsByPrivacyContextJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaIndexTagsByPrivacyContextJobDataWithDefaults() *KalturaIndexTagsByPrivacyContextJobData {
	this := KalturaIndexTagsByPrivacyContextJobData{}
	return &this
}

func (o KalturaIndexTagsByPrivacyContextJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaIndexTagsByPrivacyContextJobData struct {
	value *KalturaIndexTagsByPrivacyContextJobData
	isSet bool
}

func (v NullableKalturaIndexTagsByPrivacyContextJobData) Get() *KalturaIndexTagsByPrivacyContextJobData {
	return v.value
}

func (v *NullableKalturaIndexTagsByPrivacyContextJobData) Set(val *KalturaIndexTagsByPrivacyContextJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaIndexTagsByPrivacyContextJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaIndexTagsByPrivacyContextJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaIndexTagsByPrivacyContextJobData(val *KalturaIndexTagsByPrivacyContextJobData) *NullableKalturaIndexTagsByPrivacyContextJobData {
	return &NullableKalturaIndexTagsByPrivacyContextJobData{value: val, isSet: true}
}

func (v NullableKalturaIndexTagsByPrivacyContextJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaIndexTagsByPrivacyContextJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


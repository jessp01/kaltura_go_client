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

// KalturaBusinessProcessSignalNotificationTemplate struct for KalturaBusinessProcessSignalNotificationTemplate
type KalturaBusinessProcessSignalNotificationTemplate struct {
	KalturaBusinessProcessNotificationTemplate
}

// NewKalturaBusinessProcessSignalNotificationTemplate instantiates a new KalturaBusinessProcessSignalNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessSignalNotificationTemplate() *KalturaBusinessProcessSignalNotificationTemplate {
	this := KalturaBusinessProcessSignalNotificationTemplate{}
	return &this
}

// NewKalturaBusinessProcessSignalNotificationTemplateWithDefaults instantiates a new KalturaBusinessProcessSignalNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessSignalNotificationTemplateWithDefaults() *KalturaBusinessProcessSignalNotificationTemplate {
	this := KalturaBusinessProcessSignalNotificationTemplate{}
	return &this
}

func (o KalturaBusinessProcessSignalNotificationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBusinessProcessNotificationTemplate, errKalturaBusinessProcessNotificationTemplate := json.Marshal(o.KalturaBusinessProcessNotificationTemplate)
	if errKalturaBusinessProcessNotificationTemplate != nil {
		return []byte{}, errKalturaBusinessProcessNotificationTemplate
	}
	errKalturaBusinessProcessNotificationTemplate = json.Unmarshal([]byte(serializedKalturaBusinessProcessNotificationTemplate), &toSerialize)
	if errKalturaBusinessProcessNotificationTemplate != nil {
		return []byte{}, errKalturaBusinessProcessNotificationTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBusinessProcessSignalNotificationTemplate struct {
	value *KalturaBusinessProcessSignalNotificationTemplate
	isSet bool
}

func (v NullableKalturaBusinessProcessSignalNotificationTemplate) Get() *KalturaBusinessProcessSignalNotificationTemplate {
	return v.value
}

func (v *NullableKalturaBusinessProcessSignalNotificationTemplate) Set(val *KalturaBusinessProcessSignalNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessSignalNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessSignalNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessSignalNotificationTemplate(val *KalturaBusinessProcessSignalNotificationTemplate) *NullableKalturaBusinessProcessSignalNotificationTemplate {
	return &NullableKalturaBusinessProcessSignalNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessSignalNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessSignalNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



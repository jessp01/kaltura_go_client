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

// KalturaBusinessProcessStartNotificationTemplate struct for KalturaBusinessProcessStartNotificationTemplate
type KalturaBusinessProcessStartNotificationTemplate struct {
	KalturaBusinessProcessNotificationTemplate
}

// NewKalturaBusinessProcessStartNotificationTemplate instantiates a new KalturaBusinessProcessStartNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessStartNotificationTemplate() *KalturaBusinessProcessStartNotificationTemplate {
	this := KalturaBusinessProcessStartNotificationTemplate{}
	return &this
}

// NewKalturaBusinessProcessStartNotificationTemplateWithDefaults instantiates a new KalturaBusinessProcessStartNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessStartNotificationTemplateWithDefaults() *KalturaBusinessProcessStartNotificationTemplate {
	this := KalturaBusinessProcessStartNotificationTemplate{}
	return &this
}

func (o KalturaBusinessProcessStartNotificationTemplate) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBusinessProcessStartNotificationTemplate struct {
	value *KalturaBusinessProcessStartNotificationTemplate
	isSet bool
}

func (v NullableKalturaBusinessProcessStartNotificationTemplate) Get() *KalturaBusinessProcessStartNotificationTemplate {
	return v.value
}

func (v *NullableKalturaBusinessProcessStartNotificationTemplate) Set(val *KalturaBusinessProcessStartNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessStartNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessStartNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessStartNotificationTemplate(val *KalturaBusinessProcessStartNotificationTemplate) *NullableKalturaBusinessProcessStartNotificationTemplate {
	return &NullableKalturaBusinessProcessStartNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessStartNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessStartNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



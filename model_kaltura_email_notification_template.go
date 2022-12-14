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

// KalturaEmailNotificationTemplate struct for KalturaEmailNotificationTemplate
type KalturaEmailNotificationTemplate struct {
	KalturaEventNotificationTemplate
}

// NewKalturaEmailNotificationTemplate instantiates a new KalturaEmailNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEmailNotificationTemplate() *KalturaEmailNotificationTemplate {
	this := KalturaEmailNotificationTemplate{}
	return &this
}

// NewKalturaEmailNotificationTemplateWithDefaults instantiates a new KalturaEmailNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEmailNotificationTemplateWithDefaults() *KalturaEmailNotificationTemplate {
	this := KalturaEmailNotificationTemplate{}
	return &this
}

func (o KalturaEmailNotificationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEventNotificationTemplate, errKalturaEventNotificationTemplate := json.Marshal(o.KalturaEventNotificationTemplate)
	if errKalturaEventNotificationTemplate != nil {
		return []byte{}, errKalturaEventNotificationTemplate
	}
	errKalturaEventNotificationTemplate = json.Unmarshal([]byte(serializedKalturaEventNotificationTemplate), &toSerialize)
	if errKalturaEventNotificationTemplate != nil {
		return []byte{}, errKalturaEventNotificationTemplate
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEmailNotificationTemplate struct {
	value *KalturaEmailNotificationTemplate
	isSet bool
}

func (v NullableKalturaEmailNotificationTemplate) Get() *KalturaEmailNotificationTemplate {
	return v.value
}

func (v *NullableKalturaEmailNotificationTemplate) Set(val *KalturaEmailNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEmailNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEmailNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEmailNotificationTemplate(val *KalturaEmailNotificationTemplate) *NullableKalturaEmailNotificationTemplate {
	return &NullableKalturaEmailNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaEmailNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEmailNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



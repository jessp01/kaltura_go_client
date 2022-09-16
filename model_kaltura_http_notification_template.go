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

// KalturaHttpNotificationTemplate struct for KalturaHttpNotificationTemplate
type KalturaHttpNotificationTemplate struct {
	KalturaEventNotificationTemplate
}

// NewKalturaHttpNotificationTemplate instantiates a new KalturaHttpNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHttpNotificationTemplate() *KalturaHttpNotificationTemplate {
	this := KalturaHttpNotificationTemplate{}
	return &this
}

// NewKalturaHttpNotificationTemplateWithDefaults instantiates a new KalturaHttpNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHttpNotificationTemplateWithDefaults() *KalturaHttpNotificationTemplate {
	this := KalturaHttpNotificationTemplate{}
	return &this
}

func (o KalturaHttpNotificationTemplate) MarshalJSON() ([]byte, error) {
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

type NullableKalturaHttpNotificationTemplate struct {
	value *KalturaHttpNotificationTemplate
	isSet bool
}

func (v NullableKalturaHttpNotificationTemplate) Get() *KalturaHttpNotificationTemplate {
	return v.value
}

func (v *NullableKalturaHttpNotificationTemplate) Set(val *KalturaHttpNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHttpNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHttpNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHttpNotificationTemplate(val *KalturaHttpNotificationTemplate) *NullableKalturaHttpNotificationTemplate {
	return &NullableKalturaHttpNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaHttpNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHttpNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



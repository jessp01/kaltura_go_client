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

// KalturaKafkaNotificationTemplate struct for KalturaKafkaNotificationTemplate
type KalturaKafkaNotificationTemplate struct {
	KalturaEventNotificationTemplate
}

// NewKalturaKafkaNotificationTemplate instantiates a new KalturaKafkaNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaKafkaNotificationTemplate() *KalturaKafkaNotificationTemplate {
	this := KalturaKafkaNotificationTemplate{}
	return &this
}

// NewKalturaKafkaNotificationTemplateWithDefaults instantiates a new KalturaKafkaNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaKafkaNotificationTemplateWithDefaults() *KalturaKafkaNotificationTemplate {
	this := KalturaKafkaNotificationTemplate{}
	return &this
}

func (o KalturaKafkaNotificationTemplate) MarshalJSON() ([]byte, error) {
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

type NullableKalturaKafkaNotificationTemplate struct {
	value *KalturaKafkaNotificationTemplate
	isSet bool
}

func (v NullableKalturaKafkaNotificationTemplate) Get() *KalturaKafkaNotificationTemplate {
	return v.value
}

func (v *NullableKalturaKafkaNotificationTemplate) Set(val *KalturaKafkaNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaKafkaNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaKafkaNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaKafkaNotificationTemplate(val *KalturaKafkaNotificationTemplate) *NullableKalturaKafkaNotificationTemplate {
	return &NullableKalturaKafkaNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaKafkaNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaKafkaNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


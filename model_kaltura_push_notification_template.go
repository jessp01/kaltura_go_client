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

// KalturaPushNotificationTemplate struct for KalturaPushNotificationTemplate
type KalturaPushNotificationTemplate struct {
	KalturaEventNotificationTemplate
}

// NewKalturaPushNotificationTemplate instantiates a new KalturaPushNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPushNotificationTemplate() *KalturaPushNotificationTemplate {
	this := KalturaPushNotificationTemplate{}
	return &this
}

// NewKalturaPushNotificationTemplateWithDefaults instantiates a new KalturaPushNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPushNotificationTemplateWithDefaults() *KalturaPushNotificationTemplate {
	this := KalturaPushNotificationTemplate{}
	return &this
}

func (o KalturaPushNotificationTemplate) MarshalJSON() ([]byte, error) {
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

type NullableKalturaPushNotificationTemplate struct {
	value *KalturaPushNotificationTemplate
	isSet bool
}

func (v NullableKalturaPushNotificationTemplate) Get() *KalturaPushNotificationTemplate {
	return v.value
}

func (v *NullableKalturaPushNotificationTemplate) Set(val *KalturaPushNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPushNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPushNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPushNotificationTemplate(val *KalturaPushNotificationTemplate) *NullableKalturaPushNotificationTemplate {
	return &NullableKalturaPushNotificationTemplate{value: val, isSet: true}
}

func (v NullableKalturaPushNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPushNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



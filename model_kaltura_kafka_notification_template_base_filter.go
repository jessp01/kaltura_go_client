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

// KalturaKafkaNotificationTemplateBaseFilter `abstract`
type KalturaKafkaNotificationTemplateBaseFilter struct {
	KalturaEventNotificationTemplateFilter
}

// NewKalturaKafkaNotificationTemplateBaseFilter instantiates a new KalturaKafkaNotificationTemplateBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaKafkaNotificationTemplateBaseFilter() *KalturaKafkaNotificationTemplateBaseFilter {
	this := KalturaKafkaNotificationTemplateBaseFilter{}
	return &this
}

// NewKalturaKafkaNotificationTemplateBaseFilterWithDefaults instantiates a new KalturaKafkaNotificationTemplateBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaKafkaNotificationTemplateBaseFilterWithDefaults() *KalturaKafkaNotificationTemplateBaseFilter {
	this := KalturaKafkaNotificationTemplateBaseFilter{}
	return &this
}

func (o KalturaKafkaNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEventNotificationTemplateFilter, errKalturaEventNotificationTemplateFilter := json.Marshal(o.KalturaEventNotificationTemplateFilter)
	if errKalturaEventNotificationTemplateFilter != nil {
		return []byte{}, errKalturaEventNotificationTemplateFilter
	}
	errKalturaEventNotificationTemplateFilter = json.Unmarshal([]byte(serializedKalturaEventNotificationTemplateFilter), &toSerialize)
	if errKalturaEventNotificationTemplateFilter != nil {
		return []byte{}, errKalturaEventNotificationTemplateFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaKafkaNotificationTemplateBaseFilter struct {
	value *KalturaKafkaNotificationTemplateBaseFilter
	isSet bool
}

func (v NullableKalturaKafkaNotificationTemplateBaseFilter) Get() *KalturaKafkaNotificationTemplateBaseFilter {
	return v.value
}

func (v *NullableKalturaKafkaNotificationTemplateBaseFilter) Set(val *KalturaKafkaNotificationTemplateBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaKafkaNotificationTemplateBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaKafkaNotificationTemplateBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaKafkaNotificationTemplateBaseFilter(val *KalturaKafkaNotificationTemplateBaseFilter) *NullableKalturaKafkaNotificationTemplateBaseFilter {
	return &NullableKalturaKafkaNotificationTemplateBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaKafkaNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaKafkaNotificationTemplateBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


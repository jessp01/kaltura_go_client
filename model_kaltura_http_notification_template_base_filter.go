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

// KalturaHttpNotificationTemplateBaseFilter `abstract`
type KalturaHttpNotificationTemplateBaseFilter struct {
	KalturaEventNotificationTemplateFilter
}

// NewKalturaHttpNotificationTemplateBaseFilter instantiates a new KalturaHttpNotificationTemplateBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHttpNotificationTemplateBaseFilter() *KalturaHttpNotificationTemplateBaseFilter {
	this := KalturaHttpNotificationTemplateBaseFilter{}
	return &this
}

// NewKalturaHttpNotificationTemplateBaseFilterWithDefaults instantiates a new KalturaHttpNotificationTemplateBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHttpNotificationTemplateBaseFilterWithDefaults() *KalturaHttpNotificationTemplateBaseFilter {
	this := KalturaHttpNotificationTemplateBaseFilter{}
	return &this
}

func (o KalturaHttpNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaHttpNotificationTemplateBaseFilter struct {
	value *KalturaHttpNotificationTemplateBaseFilter
	isSet bool
}

func (v NullableKalturaHttpNotificationTemplateBaseFilter) Get() *KalturaHttpNotificationTemplateBaseFilter {
	return v.value
}

func (v *NullableKalturaHttpNotificationTemplateBaseFilter) Set(val *KalturaHttpNotificationTemplateBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHttpNotificationTemplateBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHttpNotificationTemplateBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHttpNotificationTemplateBaseFilter(val *KalturaHttpNotificationTemplateBaseFilter) *NullableKalturaHttpNotificationTemplateBaseFilter {
	return &NullableKalturaHttpNotificationTemplateBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaHttpNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHttpNotificationTemplateBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



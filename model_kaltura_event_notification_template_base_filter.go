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

// KalturaEventNotificationTemplateBaseFilter `abstract`
type KalturaEventNotificationTemplateBaseFilter struct {
	KalturaFilter
}

// NewKalturaEventNotificationTemplateBaseFilter instantiates a new KalturaEventNotificationTemplateBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEventNotificationTemplateBaseFilter() *KalturaEventNotificationTemplateBaseFilter {
	this := KalturaEventNotificationTemplateBaseFilter{}
	return &this
}

// NewKalturaEventNotificationTemplateBaseFilterWithDefaults instantiates a new KalturaEventNotificationTemplateBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEventNotificationTemplateBaseFilterWithDefaults() *KalturaEventNotificationTemplateBaseFilter {
	this := KalturaEventNotificationTemplateBaseFilter{}
	return &this
}

func (o KalturaEventNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFilter, errKalturaFilter := json.Marshal(o.KalturaFilter)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	errKalturaFilter = json.Unmarshal([]byte(serializedKalturaFilter), &toSerialize)
	if errKalturaFilter != nil {
		return []byte{}, errKalturaFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEventNotificationTemplateBaseFilter struct {
	value *KalturaEventNotificationTemplateBaseFilter
	isSet bool
}

func (v NullableKalturaEventNotificationTemplateBaseFilter) Get() *KalturaEventNotificationTemplateBaseFilter {
	return v.value
}

func (v *NullableKalturaEventNotificationTemplateBaseFilter) Set(val *KalturaEventNotificationTemplateBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEventNotificationTemplateBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEventNotificationTemplateBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEventNotificationTemplateBaseFilter(val *KalturaEventNotificationTemplateBaseFilter) *NullableKalturaEventNotificationTemplateBaseFilter {
	return &NullableKalturaEventNotificationTemplateBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaEventNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEventNotificationTemplateBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



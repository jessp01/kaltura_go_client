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

// KalturaBusinessProcessAbortNotificationTemplateBaseFilter `abstract`
type KalturaBusinessProcessAbortNotificationTemplateBaseFilter struct {
	KalturaBusinessProcessNotificationTemplateFilter
}

// NewKalturaBusinessProcessAbortNotificationTemplateBaseFilter instantiates a new KalturaBusinessProcessAbortNotificationTemplateBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBusinessProcessAbortNotificationTemplateBaseFilter() *KalturaBusinessProcessAbortNotificationTemplateBaseFilter {
	this := KalturaBusinessProcessAbortNotificationTemplateBaseFilter{}
	return &this
}

// NewKalturaBusinessProcessAbortNotificationTemplateBaseFilterWithDefaults instantiates a new KalturaBusinessProcessAbortNotificationTemplateBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBusinessProcessAbortNotificationTemplateBaseFilterWithDefaults() *KalturaBusinessProcessAbortNotificationTemplateBaseFilter {
	this := KalturaBusinessProcessAbortNotificationTemplateBaseFilter{}
	return &this
}

func (o KalturaBusinessProcessAbortNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBusinessProcessNotificationTemplateFilter, errKalturaBusinessProcessNotificationTemplateFilter := json.Marshal(o.KalturaBusinessProcessNotificationTemplateFilter)
	if errKalturaBusinessProcessNotificationTemplateFilter != nil {
		return []byte{}, errKalturaBusinessProcessNotificationTemplateFilter
	}
	errKalturaBusinessProcessNotificationTemplateFilter = json.Unmarshal([]byte(serializedKalturaBusinessProcessNotificationTemplateFilter), &toSerialize)
	if errKalturaBusinessProcessNotificationTemplateFilter != nil {
		return []byte{}, errKalturaBusinessProcessNotificationTemplateFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter struct {
	value *KalturaBusinessProcessAbortNotificationTemplateBaseFilter
	isSet bool
}

func (v NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) Get() *KalturaBusinessProcessAbortNotificationTemplateBaseFilter {
	return v.value
}

func (v *NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) Set(val *KalturaBusinessProcessAbortNotificationTemplateBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter(val *KalturaBusinessProcessAbortNotificationTemplateBaseFilter) *NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter {
	return &NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBusinessProcessAbortNotificationTemplateBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// KalturaInteractivityNodeFilter struct for KalturaInteractivityNodeFilter
type KalturaInteractivityNodeFilter struct {
	KalturaInteractivityDataFieldsFilter
}

// NewKalturaInteractivityNodeFilter instantiates a new KalturaInteractivityNodeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaInteractivityNodeFilter() *KalturaInteractivityNodeFilter {
	this := KalturaInteractivityNodeFilter{}
	return &this
}

// NewKalturaInteractivityNodeFilterWithDefaults instantiates a new KalturaInteractivityNodeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaInteractivityNodeFilterWithDefaults() *KalturaInteractivityNodeFilter {
	this := KalturaInteractivityNodeFilter{}
	return &this
}

func (o KalturaInteractivityNodeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaInteractivityDataFieldsFilter, errKalturaInteractivityDataFieldsFilter := json.Marshal(o.KalturaInteractivityDataFieldsFilter)
	if errKalturaInteractivityDataFieldsFilter != nil {
		return []byte{}, errKalturaInteractivityDataFieldsFilter
	}
	errKalturaInteractivityDataFieldsFilter = json.Unmarshal([]byte(serializedKalturaInteractivityDataFieldsFilter), &toSerialize)
	if errKalturaInteractivityDataFieldsFilter != nil {
		return []byte{}, errKalturaInteractivityDataFieldsFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaInteractivityNodeFilter struct {
	value *KalturaInteractivityNodeFilter
	isSet bool
}

func (v NullableKalturaInteractivityNodeFilter) Get() *KalturaInteractivityNodeFilter {
	return v.value
}

func (v *NullableKalturaInteractivityNodeFilter) Set(val *KalturaInteractivityNodeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaInteractivityNodeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaInteractivityNodeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaInteractivityNodeFilter(val *KalturaInteractivityNodeFilter) *NullableKalturaInteractivityNodeFilter {
	return &NullableKalturaInteractivityNodeFilter{value: val, isSet: true}
}

func (v NullableKalturaInteractivityNodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaInteractivityNodeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



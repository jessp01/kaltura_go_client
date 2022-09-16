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

// KalturaInteractivityRootFilter struct for KalturaInteractivityRootFilter
type KalturaInteractivityRootFilter struct {
	KalturaInteractivityDataFieldsFilter
}

// NewKalturaInteractivityRootFilter instantiates a new KalturaInteractivityRootFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaInteractivityRootFilter() *KalturaInteractivityRootFilter {
	this := KalturaInteractivityRootFilter{}
	return &this
}

// NewKalturaInteractivityRootFilterWithDefaults instantiates a new KalturaInteractivityRootFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaInteractivityRootFilterWithDefaults() *KalturaInteractivityRootFilter {
	this := KalturaInteractivityRootFilter{}
	return &this
}

func (o KalturaInteractivityRootFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaInteractivityRootFilter struct {
	value *KalturaInteractivityRootFilter
	isSet bool
}

func (v NullableKalturaInteractivityRootFilter) Get() *KalturaInteractivityRootFilter {
	return v.value
}

func (v *NullableKalturaInteractivityRootFilter) Set(val *KalturaInteractivityRootFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaInteractivityRootFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaInteractivityRootFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaInteractivityRootFilter(val *KalturaInteractivityRootFilter) *NullableKalturaInteractivityRootFilter {
	return &NullableKalturaInteractivityRootFilter{value: val, isSet: true}
}

func (v NullableKalturaInteractivityRootFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaInteractivityRootFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



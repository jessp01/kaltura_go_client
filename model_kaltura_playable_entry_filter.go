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

// KalturaPlayableEntryFilter struct for KalturaPlayableEntryFilter
type KalturaPlayableEntryFilter struct {
	KalturaPlayableEntryBaseFilter
}

// NewKalturaPlayableEntryFilter instantiates a new KalturaPlayableEntryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayableEntryFilter() *KalturaPlayableEntryFilter {
	this := KalturaPlayableEntryFilter{}
	return &this
}

// NewKalturaPlayableEntryFilterWithDefaults instantiates a new KalturaPlayableEntryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayableEntryFilterWithDefaults() *KalturaPlayableEntryFilter {
	this := KalturaPlayableEntryFilter{}
	return &this
}

func (o KalturaPlayableEntryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPlayableEntryBaseFilter, errKalturaPlayableEntryBaseFilter := json.Marshal(o.KalturaPlayableEntryBaseFilter)
	if errKalturaPlayableEntryBaseFilter != nil {
		return []byte{}, errKalturaPlayableEntryBaseFilter
	}
	errKalturaPlayableEntryBaseFilter = json.Unmarshal([]byte(serializedKalturaPlayableEntryBaseFilter), &toSerialize)
	if errKalturaPlayableEntryBaseFilter != nil {
		return []byte{}, errKalturaPlayableEntryBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayableEntryFilter struct {
	value *KalturaPlayableEntryFilter
	isSet bool
}

func (v NullableKalturaPlayableEntryFilter) Get() *KalturaPlayableEntryFilter {
	return v.value
}

func (v *NullableKalturaPlayableEntryFilter) Set(val *KalturaPlayableEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayableEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayableEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayableEntryFilter(val *KalturaPlayableEntryFilter) *NullableKalturaPlayableEntryFilter {
	return &NullableKalturaPlayableEntryFilter{value: val, isSet: true}
}

func (v NullableKalturaPlayableEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayableEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


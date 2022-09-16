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

// KalturaMixEntryBaseFilter `abstract`
type KalturaMixEntryBaseFilter struct {
	KalturaPlayableEntryFilter
}

// NewKalturaMixEntryBaseFilter instantiates a new KalturaMixEntryBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMixEntryBaseFilter() *KalturaMixEntryBaseFilter {
	this := KalturaMixEntryBaseFilter{}
	return &this
}

// NewKalturaMixEntryBaseFilterWithDefaults instantiates a new KalturaMixEntryBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMixEntryBaseFilterWithDefaults() *KalturaMixEntryBaseFilter {
	this := KalturaMixEntryBaseFilter{}
	return &this
}

func (o KalturaMixEntryBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPlayableEntryFilter, errKalturaPlayableEntryFilter := json.Marshal(o.KalturaPlayableEntryFilter)
	if errKalturaPlayableEntryFilter != nil {
		return []byte{}, errKalturaPlayableEntryFilter
	}
	errKalturaPlayableEntryFilter = json.Unmarshal([]byte(serializedKalturaPlayableEntryFilter), &toSerialize)
	if errKalturaPlayableEntryFilter != nil {
		return []byte{}, errKalturaPlayableEntryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMixEntryBaseFilter struct {
	value *KalturaMixEntryBaseFilter
	isSet bool
}

func (v NullableKalturaMixEntryBaseFilter) Get() *KalturaMixEntryBaseFilter {
	return v.value
}

func (v *NullableKalturaMixEntryBaseFilter) Set(val *KalturaMixEntryBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMixEntryBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMixEntryBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMixEntryBaseFilter(val *KalturaMixEntryBaseFilter) *NullableKalturaMixEntryBaseFilter {
	return &NullableKalturaMixEntryBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaMixEntryBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMixEntryBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


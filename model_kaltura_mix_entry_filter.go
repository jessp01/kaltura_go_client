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

// KalturaMixEntryFilter struct for KalturaMixEntryFilter
type KalturaMixEntryFilter struct {
	KalturaMixEntryBaseFilter
}

// NewKalturaMixEntryFilter instantiates a new KalturaMixEntryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMixEntryFilter() *KalturaMixEntryFilter {
	this := KalturaMixEntryFilter{}
	return &this
}

// NewKalturaMixEntryFilterWithDefaults instantiates a new KalturaMixEntryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMixEntryFilterWithDefaults() *KalturaMixEntryFilter {
	this := KalturaMixEntryFilter{}
	return &this
}

func (o KalturaMixEntryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMixEntryBaseFilter, errKalturaMixEntryBaseFilter := json.Marshal(o.KalturaMixEntryBaseFilter)
	if errKalturaMixEntryBaseFilter != nil {
		return []byte{}, errKalturaMixEntryBaseFilter
	}
	errKalturaMixEntryBaseFilter = json.Unmarshal([]byte(serializedKalturaMixEntryBaseFilter), &toSerialize)
	if errKalturaMixEntryBaseFilter != nil {
		return []byte{}, errKalturaMixEntryBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMixEntryFilter struct {
	value *KalturaMixEntryFilter
	isSet bool
}

func (v NullableKalturaMixEntryFilter) Get() *KalturaMixEntryFilter {
	return v.value
}

func (v *NullableKalturaMixEntryFilter) Set(val *KalturaMixEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMixEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMixEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMixEntryFilter(val *KalturaMixEntryFilter) *NullableKalturaMixEntryFilter {
	return &NullableKalturaMixEntryFilter{value: val, isSet: true}
}

func (v NullableKalturaMixEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMixEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



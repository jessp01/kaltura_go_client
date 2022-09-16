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

// KalturaLiveEntryBaseFilter `abstract`
type KalturaLiveEntryBaseFilter struct {
	KalturaMediaEntryFilter
}

// NewKalturaLiveEntryBaseFilter instantiates a new KalturaLiveEntryBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveEntryBaseFilter() *KalturaLiveEntryBaseFilter {
	this := KalturaLiveEntryBaseFilter{}
	return &this
}

// NewKalturaLiveEntryBaseFilterWithDefaults instantiates a new KalturaLiveEntryBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveEntryBaseFilterWithDefaults() *KalturaLiveEntryBaseFilter {
	this := KalturaLiveEntryBaseFilter{}
	return &this
}

func (o KalturaLiveEntryBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaEntryFilter, errKalturaMediaEntryFilter := json.Marshal(o.KalturaMediaEntryFilter)
	if errKalturaMediaEntryFilter != nil {
		return []byte{}, errKalturaMediaEntryFilter
	}
	errKalturaMediaEntryFilter = json.Unmarshal([]byte(serializedKalturaMediaEntryFilter), &toSerialize)
	if errKalturaMediaEntryFilter != nil {
		return []byte{}, errKalturaMediaEntryFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveEntryBaseFilter struct {
	value *KalturaLiveEntryBaseFilter
	isSet bool
}

func (v NullableKalturaLiveEntryBaseFilter) Get() *KalturaLiveEntryBaseFilter {
	return v.value
}

func (v *NullableKalturaLiveEntryBaseFilter) Set(val *KalturaLiveEntryBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveEntryBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveEntryBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveEntryBaseFilter(val *KalturaLiveEntryBaseFilter) *NullableKalturaLiveEntryBaseFilter {
	return &NullableKalturaLiveEntryBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaLiveEntryBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveEntryBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


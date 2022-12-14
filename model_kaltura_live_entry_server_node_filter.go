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

// KalturaLiveEntryServerNodeFilter struct for KalturaLiveEntryServerNodeFilter
type KalturaLiveEntryServerNodeFilter struct {
	KalturaLiveEntryServerNodeBaseFilter
}

// NewKalturaLiveEntryServerNodeFilter instantiates a new KalturaLiveEntryServerNodeFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveEntryServerNodeFilter() *KalturaLiveEntryServerNodeFilter {
	this := KalturaLiveEntryServerNodeFilter{}
	return &this
}

// NewKalturaLiveEntryServerNodeFilterWithDefaults instantiates a new KalturaLiveEntryServerNodeFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveEntryServerNodeFilterWithDefaults() *KalturaLiveEntryServerNodeFilter {
	this := KalturaLiveEntryServerNodeFilter{}
	return &this
}

func (o KalturaLiveEntryServerNodeFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaLiveEntryServerNodeBaseFilter, errKalturaLiveEntryServerNodeBaseFilter := json.Marshal(o.KalturaLiveEntryServerNodeBaseFilter)
	if errKalturaLiveEntryServerNodeBaseFilter != nil {
		return []byte{}, errKalturaLiveEntryServerNodeBaseFilter
	}
	errKalturaLiveEntryServerNodeBaseFilter = json.Unmarshal([]byte(serializedKalturaLiveEntryServerNodeBaseFilter), &toSerialize)
	if errKalturaLiveEntryServerNodeBaseFilter != nil {
		return []byte{}, errKalturaLiveEntryServerNodeBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveEntryServerNodeFilter struct {
	value *KalturaLiveEntryServerNodeFilter
	isSet bool
}

func (v NullableKalturaLiveEntryServerNodeFilter) Get() *KalturaLiveEntryServerNodeFilter {
	return v.value
}

func (v *NullableKalturaLiveEntryServerNodeFilter) Set(val *KalturaLiveEntryServerNodeFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveEntryServerNodeFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveEntryServerNodeFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveEntryServerNodeFilter(val *KalturaLiveEntryServerNodeFilter) *NullableKalturaLiveEntryServerNodeFilter {
	return &NullableKalturaLiveEntryServerNodeFilter{value: val, isSet: true}
}

func (v NullableKalturaLiveEntryServerNodeFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveEntryServerNodeFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



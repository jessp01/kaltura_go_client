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

// KalturaLiveChannelSegmentFilter struct for KalturaLiveChannelSegmentFilter
type KalturaLiveChannelSegmentFilter struct {
	KalturaLiveChannelSegmentBaseFilter
}

// NewKalturaLiveChannelSegmentFilter instantiates a new KalturaLiveChannelSegmentFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLiveChannelSegmentFilter() *KalturaLiveChannelSegmentFilter {
	this := KalturaLiveChannelSegmentFilter{}
	return &this
}

// NewKalturaLiveChannelSegmentFilterWithDefaults instantiates a new KalturaLiveChannelSegmentFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLiveChannelSegmentFilterWithDefaults() *KalturaLiveChannelSegmentFilter {
	this := KalturaLiveChannelSegmentFilter{}
	return &this
}

func (o KalturaLiveChannelSegmentFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaLiveChannelSegmentBaseFilter, errKalturaLiveChannelSegmentBaseFilter := json.Marshal(o.KalturaLiveChannelSegmentBaseFilter)
	if errKalturaLiveChannelSegmentBaseFilter != nil {
		return []byte{}, errKalturaLiveChannelSegmentBaseFilter
	}
	errKalturaLiveChannelSegmentBaseFilter = json.Unmarshal([]byte(serializedKalturaLiveChannelSegmentBaseFilter), &toSerialize)
	if errKalturaLiveChannelSegmentBaseFilter != nil {
		return []byte{}, errKalturaLiveChannelSegmentBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLiveChannelSegmentFilter struct {
	value *KalturaLiveChannelSegmentFilter
	isSet bool
}

func (v NullableKalturaLiveChannelSegmentFilter) Get() *KalturaLiveChannelSegmentFilter {
	return v.value
}

func (v *NullableKalturaLiveChannelSegmentFilter) Set(val *KalturaLiveChannelSegmentFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLiveChannelSegmentFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLiveChannelSegmentFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLiveChannelSegmentFilter(val *KalturaLiveChannelSegmentFilter) *NullableKalturaLiveChannelSegmentFilter {
	return &NullableKalturaLiveChannelSegmentFilter{value: val, isSet: true}
}

func (v NullableKalturaLiveChannelSegmentFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLiveChannelSegmentFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


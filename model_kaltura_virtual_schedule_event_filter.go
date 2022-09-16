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

// KalturaVirtualScheduleEventFilter struct for KalturaVirtualScheduleEventFilter
type KalturaVirtualScheduleEventFilter struct {
	KalturaVirtualScheduleEventBaseFilter
}

// NewKalturaVirtualScheduleEventFilter instantiates a new KalturaVirtualScheduleEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirtualScheduleEventFilter() *KalturaVirtualScheduleEventFilter {
	this := KalturaVirtualScheduleEventFilter{}
	return &this
}

// NewKalturaVirtualScheduleEventFilterWithDefaults instantiates a new KalturaVirtualScheduleEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirtualScheduleEventFilterWithDefaults() *KalturaVirtualScheduleEventFilter {
	this := KalturaVirtualScheduleEventFilter{}
	return &this
}

func (o KalturaVirtualScheduleEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaVirtualScheduleEventBaseFilter, errKalturaVirtualScheduleEventBaseFilter := json.Marshal(o.KalturaVirtualScheduleEventBaseFilter)
	if errKalturaVirtualScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaVirtualScheduleEventBaseFilter
	}
	errKalturaVirtualScheduleEventBaseFilter = json.Unmarshal([]byte(serializedKalturaVirtualScheduleEventBaseFilter), &toSerialize)
	if errKalturaVirtualScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaVirtualScheduleEventBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirtualScheduleEventFilter struct {
	value *KalturaVirtualScheduleEventFilter
	isSet bool
}

func (v NullableKalturaVirtualScheduleEventFilter) Get() *KalturaVirtualScheduleEventFilter {
	return v.value
}

func (v *NullableKalturaVirtualScheduleEventFilter) Set(val *KalturaVirtualScheduleEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirtualScheduleEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirtualScheduleEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirtualScheduleEventFilter(val *KalturaVirtualScheduleEventFilter) *NullableKalturaVirtualScheduleEventFilter {
	return &NullableKalturaVirtualScheduleEventFilter{value: val, isSet: true}
}

func (v NullableKalturaVirtualScheduleEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirtualScheduleEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


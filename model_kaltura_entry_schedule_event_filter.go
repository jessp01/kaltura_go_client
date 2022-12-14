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

// KalturaEntryScheduleEventFilter struct for KalturaEntryScheduleEventFilter
type KalturaEntryScheduleEventFilter struct {
	KalturaEntryScheduleEventBaseFilter
}

// NewKalturaEntryScheduleEventFilter instantiates a new KalturaEntryScheduleEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryScheduleEventFilter() *KalturaEntryScheduleEventFilter {
	this := KalturaEntryScheduleEventFilter{}
	return &this
}

// NewKalturaEntryScheduleEventFilterWithDefaults instantiates a new KalturaEntryScheduleEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryScheduleEventFilterWithDefaults() *KalturaEntryScheduleEventFilter {
	this := KalturaEntryScheduleEventFilter{}
	return &this
}

func (o KalturaEntryScheduleEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaEntryScheduleEventBaseFilter, errKalturaEntryScheduleEventBaseFilter := json.Marshal(o.KalturaEntryScheduleEventBaseFilter)
	if errKalturaEntryScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaEntryScheduleEventBaseFilter
	}
	errKalturaEntryScheduleEventBaseFilter = json.Unmarshal([]byte(serializedKalturaEntryScheduleEventBaseFilter), &toSerialize)
	if errKalturaEntryScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaEntryScheduleEventBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaEntryScheduleEventFilter struct {
	value *KalturaEntryScheduleEventFilter
	isSet bool
}

func (v NullableKalturaEntryScheduleEventFilter) Get() *KalturaEntryScheduleEventFilter {
	return v.value
}

func (v *NullableKalturaEntryScheduleEventFilter) Set(val *KalturaEntryScheduleEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryScheduleEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryScheduleEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryScheduleEventFilter(val *KalturaEntryScheduleEventFilter) *NullableKalturaEntryScheduleEventFilter {
	return &NullableKalturaEntryScheduleEventFilter{value: val, isSet: true}
}

func (v NullableKalturaEntryScheduleEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryScheduleEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



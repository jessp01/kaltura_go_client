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

// KalturaScheduleEventFilter struct for KalturaScheduleEventFilter
type KalturaScheduleEventFilter struct {
	KalturaScheduleEventBaseFilter
}

// NewKalturaScheduleEventFilter instantiates a new KalturaScheduleEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduleEventFilter() *KalturaScheduleEventFilter {
	this := KalturaScheduleEventFilter{}
	return &this
}

// NewKalturaScheduleEventFilterWithDefaults instantiates a new KalturaScheduleEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduleEventFilterWithDefaults() *KalturaScheduleEventFilter {
	this := KalturaScheduleEventFilter{}
	return &this
}

func (o KalturaScheduleEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleEventBaseFilter, errKalturaScheduleEventBaseFilter := json.Marshal(o.KalturaScheduleEventBaseFilter)
	if errKalturaScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaScheduleEventBaseFilter
	}
	errKalturaScheduleEventBaseFilter = json.Unmarshal([]byte(serializedKalturaScheduleEventBaseFilter), &toSerialize)
	if errKalturaScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaScheduleEventBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduleEventFilter struct {
	value *KalturaScheduleEventFilter
	isSet bool
}

func (v NullableKalturaScheduleEventFilter) Get() *KalturaScheduleEventFilter {
	return v.value
}

func (v *NullableKalturaScheduleEventFilter) Set(val *KalturaScheduleEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduleEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduleEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduleEventFilter(val *KalturaScheduleEventFilter) *NullableKalturaScheduleEventFilter {
	return &NullableKalturaScheduleEventFilter{value: val, isSet: true}
}

func (v NullableKalturaScheduleEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduleEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



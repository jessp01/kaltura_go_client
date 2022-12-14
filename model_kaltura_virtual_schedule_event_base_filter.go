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

// KalturaVirtualScheduleEventBaseFilter `abstract`
type KalturaVirtualScheduleEventBaseFilter struct {
	KalturaScheduleEventFilter
}

// NewKalturaVirtualScheduleEventBaseFilter instantiates a new KalturaVirtualScheduleEventBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirtualScheduleEventBaseFilter() *KalturaVirtualScheduleEventBaseFilter {
	this := KalturaVirtualScheduleEventBaseFilter{}
	return &this
}

// NewKalturaVirtualScheduleEventBaseFilterWithDefaults instantiates a new KalturaVirtualScheduleEventBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirtualScheduleEventBaseFilterWithDefaults() *KalturaVirtualScheduleEventBaseFilter {
	this := KalturaVirtualScheduleEventBaseFilter{}
	return &this
}

func (o KalturaVirtualScheduleEventBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleEventFilter, errKalturaScheduleEventFilter := json.Marshal(o.KalturaScheduleEventFilter)
	if errKalturaScheduleEventFilter != nil {
		return []byte{}, errKalturaScheduleEventFilter
	}
	errKalturaScheduleEventFilter = json.Unmarshal([]byte(serializedKalturaScheduleEventFilter), &toSerialize)
	if errKalturaScheduleEventFilter != nil {
		return []byte{}, errKalturaScheduleEventFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirtualScheduleEventBaseFilter struct {
	value *KalturaVirtualScheduleEventBaseFilter
	isSet bool
}

func (v NullableKalturaVirtualScheduleEventBaseFilter) Get() *KalturaVirtualScheduleEventBaseFilter {
	return v.value
}

func (v *NullableKalturaVirtualScheduleEventBaseFilter) Set(val *KalturaVirtualScheduleEventBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirtualScheduleEventBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirtualScheduleEventBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirtualScheduleEventBaseFilter(val *KalturaVirtualScheduleEventBaseFilter) *NullableKalturaVirtualScheduleEventBaseFilter {
	return &NullableKalturaVirtualScheduleEventBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaVirtualScheduleEventBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirtualScheduleEventBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



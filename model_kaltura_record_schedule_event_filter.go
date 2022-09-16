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

// KalturaRecordScheduleEventFilter struct for KalturaRecordScheduleEventFilter
type KalturaRecordScheduleEventFilter struct {
	KalturaRecordScheduleEventBaseFilter
}

// NewKalturaRecordScheduleEventFilter instantiates a new KalturaRecordScheduleEventFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaRecordScheduleEventFilter() *KalturaRecordScheduleEventFilter {
	this := KalturaRecordScheduleEventFilter{}
	return &this
}

// NewKalturaRecordScheduleEventFilterWithDefaults instantiates a new KalturaRecordScheduleEventFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaRecordScheduleEventFilterWithDefaults() *KalturaRecordScheduleEventFilter {
	this := KalturaRecordScheduleEventFilter{}
	return &this
}

func (o KalturaRecordScheduleEventFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaRecordScheduleEventBaseFilter, errKalturaRecordScheduleEventBaseFilter := json.Marshal(o.KalturaRecordScheduleEventBaseFilter)
	if errKalturaRecordScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaRecordScheduleEventBaseFilter
	}
	errKalturaRecordScheduleEventBaseFilter = json.Unmarshal([]byte(serializedKalturaRecordScheduleEventBaseFilter), &toSerialize)
	if errKalturaRecordScheduleEventBaseFilter != nil {
		return []byte{}, errKalturaRecordScheduleEventBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaRecordScheduleEventFilter struct {
	value *KalturaRecordScheduleEventFilter
	isSet bool
}

func (v NullableKalturaRecordScheduleEventFilter) Get() *KalturaRecordScheduleEventFilter {
	return v.value
}

func (v *NullableKalturaRecordScheduleEventFilter) Set(val *KalturaRecordScheduleEventFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaRecordScheduleEventFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaRecordScheduleEventFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaRecordScheduleEventFilter(val *KalturaRecordScheduleEventFilter) *NullableKalturaRecordScheduleEventFilter {
	return &NullableKalturaRecordScheduleEventFilter{value: val, isSet: true}
}

func (v NullableKalturaRecordScheduleEventFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaRecordScheduleEventFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


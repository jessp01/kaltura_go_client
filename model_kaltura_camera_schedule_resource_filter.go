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

// KalturaCameraScheduleResourceFilter struct for KalturaCameraScheduleResourceFilter
type KalturaCameraScheduleResourceFilter struct {
	KalturaCameraScheduleResourceBaseFilter
}

// NewKalturaCameraScheduleResourceFilter instantiates a new KalturaCameraScheduleResourceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCameraScheduleResourceFilter() *KalturaCameraScheduleResourceFilter {
	this := KalturaCameraScheduleResourceFilter{}
	return &this
}

// NewKalturaCameraScheduleResourceFilterWithDefaults instantiates a new KalturaCameraScheduleResourceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCameraScheduleResourceFilterWithDefaults() *KalturaCameraScheduleResourceFilter {
	this := KalturaCameraScheduleResourceFilter{}
	return &this
}

func (o KalturaCameraScheduleResourceFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCameraScheduleResourceBaseFilter, errKalturaCameraScheduleResourceBaseFilter := json.Marshal(o.KalturaCameraScheduleResourceBaseFilter)
	if errKalturaCameraScheduleResourceBaseFilter != nil {
		return []byte{}, errKalturaCameraScheduleResourceBaseFilter
	}
	errKalturaCameraScheduleResourceBaseFilter = json.Unmarshal([]byte(serializedKalturaCameraScheduleResourceBaseFilter), &toSerialize)
	if errKalturaCameraScheduleResourceBaseFilter != nil {
		return []byte{}, errKalturaCameraScheduleResourceBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCameraScheduleResourceFilter struct {
	value *KalturaCameraScheduleResourceFilter
	isSet bool
}

func (v NullableKalturaCameraScheduleResourceFilter) Get() *KalturaCameraScheduleResourceFilter {
	return v.value
}

func (v *NullableKalturaCameraScheduleResourceFilter) Set(val *KalturaCameraScheduleResourceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCameraScheduleResourceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCameraScheduleResourceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCameraScheduleResourceFilter(val *KalturaCameraScheduleResourceFilter) *NullableKalturaCameraScheduleResourceFilter {
	return &NullableKalturaCameraScheduleResourceFilter{value: val, isSet: true}
}

func (v NullableKalturaCameraScheduleResourceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCameraScheduleResourceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// KalturaCuePointFilter struct for KalturaCuePointFilter
type KalturaCuePointFilter struct {
	KalturaCuePointBaseFilter
}

// NewKalturaCuePointFilter instantiates a new KalturaCuePointFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCuePointFilter() *KalturaCuePointFilter {
	this := KalturaCuePointFilter{}
	return &this
}

// NewKalturaCuePointFilterWithDefaults instantiates a new KalturaCuePointFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCuePointFilterWithDefaults() *KalturaCuePointFilter {
	this := KalturaCuePointFilter{}
	return &this
}

func (o KalturaCuePointFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCuePointBaseFilter, errKalturaCuePointBaseFilter := json.Marshal(o.KalturaCuePointBaseFilter)
	if errKalturaCuePointBaseFilter != nil {
		return []byte{}, errKalturaCuePointBaseFilter
	}
	errKalturaCuePointBaseFilter = json.Unmarshal([]byte(serializedKalturaCuePointBaseFilter), &toSerialize)
	if errKalturaCuePointBaseFilter != nil {
		return []byte{}, errKalturaCuePointBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCuePointFilter struct {
	value *KalturaCuePointFilter
	isSet bool
}

func (v NullableKalturaCuePointFilter) Get() *KalturaCuePointFilter {
	return v.value
}

func (v *NullableKalturaCuePointFilter) Set(val *KalturaCuePointFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCuePointFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCuePointFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCuePointFilter(val *KalturaCuePointFilter) *NullableKalturaCuePointFilter {
	return &NullableKalturaCuePointFilter{value: val, isSet: true}
}

func (v NullableKalturaCuePointFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCuePointFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



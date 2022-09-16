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

// KalturaAdCuePointBaseFilter `abstract`
type KalturaAdCuePointBaseFilter struct {
	KalturaCuePointFilter
}

// NewKalturaAdCuePointBaseFilter instantiates a new KalturaAdCuePointBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAdCuePointBaseFilter() *KalturaAdCuePointBaseFilter {
	this := KalturaAdCuePointBaseFilter{}
	return &this
}

// NewKalturaAdCuePointBaseFilterWithDefaults instantiates a new KalturaAdCuePointBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAdCuePointBaseFilterWithDefaults() *KalturaAdCuePointBaseFilter {
	this := KalturaAdCuePointBaseFilter{}
	return &this
}

func (o KalturaAdCuePointBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCuePointFilter, errKalturaCuePointFilter := json.Marshal(o.KalturaCuePointFilter)
	if errKalturaCuePointFilter != nil {
		return []byte{}, errKalturaCuePointFilter
	}
	errKalturaCuePointFilter = json.Unmarshal([]byte(serializedKalturaCuePointFilter), &toSerialize)
	if errKalturaCuePointFilter != nil {
		return []byte{}, errKalturaCuePointFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAdCuePointBaseFilter struct {
	value *KalturaAdCuePointBaseFilter
	isSet bool
}

func (v NullableKalturaAdCuePointBaseFilter) Get() *KalturaAdCuePointBaseFilter {
	return v.value
}

func (v *NullableKalturaAdCuePointBaseFilter) Set(val *KalturaAdCuePointBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAdCuePointBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAdCuePointBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAdCuePointBaseFilter(val *KalturaAdCuePointBaseFilter) *NullableKalturaAdCuePointBaseFilter {
	return &NullableKalturaAdCuePointBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaAdCuePointBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAdCuePointBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


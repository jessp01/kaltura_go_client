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

// KalturaAdCuePointFilter struct for KalturaAdCuePointFilter
type KalturaAdCuePointFilter struct {
	KalturaAdCuePointBaseFilter
}

// NewKalturaAdCuePointFilter instantiates a new KalturaAdCuePointFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAdCuePointFilter() *KalturaAdCuePointFilter {
	this := KalturaAdCuePointFilter{}
	return &this
}

// NewKalturaAdCuePointFilterWithDefaults instantiates a new KalturaAdCuePointFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAdCuePointFilterWithDefaults() *KalturaAdCuePointFilter {
	this := KalturaAdCuePointFilter{}
	return &this
}

func (o KalturaAdCuePointFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAdCuePointBaseFilter, errKalturaAdCuePointBaseFilter := json.Marshal(o.KalturaAdCuePointBaseFilter)
	if errKalturaAdCuePointBaseFilter != nil {
		return []byte{}, errKalturaAdCuePointBaseFilter
	}
	errKalturaAdCuePointBaseFilter = json.Unmarshal([]byte(serializedKalturaAdCuePointBaseFilter), &toSerialize)
	if errKalturaAdCuePointBaseFilter != nil {
		return []byte{}, errKalturaAdCuePointBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAdCuePointFilter struct {
	value *KalturaAdCuePointFilter
	isSet bool
}

func (v NullableKalturaAdCuePointFilter) Get() *KalturaAdCuePointFilter {
	return v.value
}

func (v *NullableKalturaAdCuePointFilter) Set(val *KalturaAdCuePointFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAdCuePointFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAdCuePointFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAdCuePointFilter(val *KalturaAdCuePointFilter) *NullableKalturaAdCuePointFilter {
	return &NullableKalturaAdCuePointFilter{value: val, isSet: true}
}

func (v NullableKalturaAdCuePointFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAdCuePointFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



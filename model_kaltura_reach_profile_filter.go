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

// KalturaReachProfileFilter struct for KalturaReachProfileFilter
type KalturaReachProfileFilter struct {
	KalturaReachProfileBaseFilter
}

// NewKalturaReachProfileFilter instantiates a new KalturaReachProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReachProfileFilter() *KalturaReachProfileFilter {
	this := KalturaReachProfileFilter{}
	return &this
}

// NewKalturaReachProfileFilterWithDefaults instantiates a new KalturaReachProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReachProfileFilterWithDefaults() *KalturaReachProfileFilter {
	this := KalturaReachProfileFilter{}
	return &this
}

func (o KalturaReachProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaReachProfileBaseFilter, errKalturaReachProfileBaseFilter := json.Marshal(o.KalturaReachProfileBaseFilter)
	if errKalturaReachProfileBaseFilter != nil {
		return []byte{}, errKalturaReachProfileBaseFilter
	}
	errKalturaReachProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaReachProfileBaseFilter), &toSerialize)
	if errKalturaReachProfileBaseFilter != nil {
		return []byte{}, errKalturaReachProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaReachProfileFilter struct {
	value *KalturaReachProfileFilter
	isSet bool
}

func (v NullableKalturaReachProfileFilter) Get() *KalturaReachProfileFilter {
	return v.value
}

func (v *NullableKalturaReachProfileFilter) Set(val *KalturaReachProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReachProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReachProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReachProfileFilter(val *KalturaReachProfileFilter) *NullableKalturaReachProfileFilter {
	return &NullableKalturaReachProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaReachProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReachProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaFairplayDrmProfileFilter struct for KalturaFairplayDrmProfileFilter
type KalturaFairplayDrmProfileFilter struct {
	KalturaFairplayDrmProfileBaseFilter
}

// NewKalturaFairplayDrmProfileFilter instantiates a new KalturaFairplayDrmProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFairplayDrmProfileFilter() *KalturaFairplayDrmProfileFilter {
	this := KalturaFairplayDrmProfileFilter{}
	return &this
}

// NewKalturaFairplayDrmProfileFilterWithDefaults instantiates a new KalturaFairplayDrmProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFairplayDrmProfileFilterWithDefaults() *KalturaFairplayDrmProfileFilter {
	this := KalturaFairplayDrmProfileFilter{}
	return &this
}

func (o KalturaFairplayDrmProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFairplayDrmProfileBaseFilter, errKalturaFairplayDrmProfileBaseFilter := json.Marshal(o.KalturaFairplayDrmProfileBaseFilter)
	if errKalturaFairplayDrmProfileBaseFilter != nil {
		return []byte{}, errKalturaFairplayDrmProfileBaseFilter
	}
	errKalturaFairplayDrmProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaFairplayDrmProfileBaseFilter), &toSerialize)
	if errKalturaFairplayDrmProfileBaseFilter != nil {
		return []byte{}, errKalturaFairplayDrmProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFairplayDrmProfileFilter struct {
	value *KalturaFairplayDrmProfileFilter
	isSet bool
}

func (v NullableKalturaFairplayDrmProfileFilter) Get() *KalturaFairplayDrmProfileFilter {
	return v.value
}

func (v *NullableKalturaFairplayDrmProfileFilter) Set(val *KalturaFairplayDrmProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFairplayDrmProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFairplayDrmProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFairplayDrmProfileFilter(val *KalturaFairplayDrmProfileFilter) *NullableKalturaFairplayDrmProfileFilter {
	return &NullableKalturaFairplayDrmProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaFairplayDrmProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFairplayDrmProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



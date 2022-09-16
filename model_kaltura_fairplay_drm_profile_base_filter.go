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

// KalturaFairplayDrmProfileBaseFilter `abstract`
type KalturaFairplayDrmProfileBaseFilter struct {
	KalturaDrmProfileFilter
}

// NewKalturaFairplayDrmProfileBaseFilter instantiates a new KalturaFairplayDrmProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFairplayDrmProfileBaseFilter() *KalturaFairplayDrmProfileBaseFilter {
	this := KalturaFairplayDrmProfileBaseFilter{}
	return &this
}

// NewKalturaFairplayDrmProfileBaseFilterWithDefaults instantiates a new KalturaFairplayDrmProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFairplayDrmProfileBaseFilterWithDefaults() *KalturaFairplayDrmProfileBaseFilter {
	this := KalturaFairplayDrmProfileBaseFilter{}
	return &this
}

func (o KalturaFairplayDrmProfileBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDrmProfileFilter, errKalturaDrmProfileFilter := json.Marshal(o.KalturaDrmProfileFilter)
	if errKalturaDrmProfileFilter != nil {
		return []byte{}, errKalturaDrmProfileFilter
	}
	errKalturaDrmProfileFilter = json.Unmarshal([]byte(serializedKalturaDrmProfileFilter), &toSerialize)
	if errKalturaDrmProfileFilter != nil {
		return []byte{}, errKalturaDrmProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFairplayDrmProfileBaseFilter struct {
	value *KalturaFairplayDrmProfileBaseFilter
	isSet bool
}

func (v NullableKalturaFairplayDrmProfileBaseFilter) Get() *KalturaFairplayDrmProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaFairplayDrmProfileBaseFilter) Set(val *KalturaFairplayDrmProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFairplayDrmProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFairplayDrmProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFairplayDrmProfileBaseFilter(val *KalturaFairplayDrmProfileBaseFilter) *NullableKalturaFairplayDrmProfileBaseFilter {
	return &NullableKalturaFairplayDrmProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaFairplayDrmProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFairplayDrmProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



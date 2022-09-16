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

// KalturaQuestionCuePointFilter struct for KalturaQuestionCuePointFilter
type KalturaQuestionCuePointFilter struct {
	KalturaQuestionCuePointBaseFilter
}

// NewKalturaQuestionCuePointFilter instantiates a new KalturaQuestionCuePointFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuestionCuePointFilter() *KalturaQuestionCuePointFilter {
	this := KalturaQuestionCuePointFilter{}
	return &this
}

// NewKalturaQuestionCuePointFilterWithDefaults instantiates a new KalturaQuestionCuePointFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuestionCuePointFilterWithDefaults() *KalturaQuestionCuePointFilter {
	this := KalturaQuestionCuePointFilter{}
	return &this
}

func (o KalturaQuestionCuePointFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaQuestionCuePointBaseFilter, errKalturaQuestionCuePointBaseFilter := json.Marshal(o.KalturaQuestionCuePointBaseFilter)
	if errKalturaQuestionCuePointBaseFilter != nil {
		return []byte{}, errKalturaQuestionCuePointBaseFilter
	}
	errKalturaQuestionCuePointBaseFilter = json.Unmarshal([]byte(serializedKalturaQuestionCuePointBaseFilter), &toSerialize)
	if errKalturaQuestionCuePointBaseFilter != nil {
		return []byte{}, errKalturaQuestionCuePointBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuestionCuePointFilter struct {
	value *KalturaQuestionCuePointFilter
	isSet bool
}

func (v NullableKalturaQuestionCuePointFilter) Get() *KalturaQuestionCuePointFilter {
	return v.value
}

func (v *NullableKalturaQuestionCuePointFilter) Set(val *KalturaQuestionCuePointFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuestionCuePointFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuestionCuePointFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuestionCuePointFilter(val *KalturaQuestionCuePointFilter) *NullableKalturaQuestionCuePointFilter {
	return &NullableKalturaQuestionCuePointFilter{value: val, isSet: true}
}

func (v NullableKalturaQuestionCuePointFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuestionCuePointFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


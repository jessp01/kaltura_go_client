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

// KalturaQuizAdvancedFilter struct for KalturaQuizAdvancedFilter
type KalturaQuizAdvancedFilter struct {
	KalturaSearchItem
}

// NewKalturaQuizAdvancedFilter instantiates a new KalturaQuizAdvancedFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuizAdvancedFilter() *KalturaQuizAdvancedFilter {
	this := KalturaQuizAdvancedFilter{}
	return &this
}

// NewKalturaQuizAdvancedFilterWithDefaults instantiates a new KalturaQuizAdvancedFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuizAdvancedFilterWithDefaults() *KalturaQuizAdvancedFilter {
	this := KalturaQuizAdvancedFilter{}
	return &this
}

func (o KalturaQuizAdvancedFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchItem, errKalturaSearchItem := json.Marshal(o.KalturaSearchItem)
	if errKalturaSearchItem != nil {
		return []byte{}, errKalturaSearchItem
	}
	errKalturaSearchItem = json.Unmarshal([]byte(serializedKalturaSearchItem), &toSerialize)
	if errKalturaSearchItem != nil {
		return []byte{}, errKalturaSearchItem
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuizAdvancedFilter struct {
	value *KalturaQuizAdvancedFilter
	isSet bool
}

func (v NullableKalturaQuizAdvancedFilter) Get() *KalturaQuizAdvancedFilter {
	return v.value
}

func (v *NullableKalturaQuizAdvancedFilter) Set(val *KalturaQuizAdvancedFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuizAdvancedFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuizAdvancedFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuizAdvancedFilter(val *KalturaQuizAdvancedFilter) *NullableKalturaQuizAdvancedFilter {
	return &NullableKalturaQuizAdvancedFilter{value: val, isSet: true}
}

func (v NullableKalturaQuizAdvancedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuizAdvancedFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



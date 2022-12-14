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

// KalturaQuizFilter struct for KalturaQuizFilter
type KalturaQuizFilter struct {
	KalturaRelatedFilter
}

// NewKalturaQuizFilter instantiates a new KalturaQuizFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuizFilter() *KalturaQuizFilter {
	this := KalturaQuizFilter{}
	return &this
}

// NewKalturaQuizFilterWithDefaults instantiates a new KalturaQuizFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuizFilterWithDefaults() *KalturaQuizFilter {
	this := KalturaQuizFilter{}
	return &this
}

func (o KalturaQuizFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaRelatedFilter, errKalturaRelatedFilter := json.Marshal(o.KalturaRelatedFilter)
	if errKalturaRelatedFilter != nil {
		return []byte{}, errKalturaRelatedFilter
	}
	errKalturaRelatedFilter = json.Unmarshal([]byte(serializedKalturaRelatedFilter), &toSerialize)
	if errKalturaRelatedFilter != nil {
		return []byte{}, errKalturaRelatedFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuizFilter struct {
	value *KalturaQuizFilter
	isSet bool
}

func (v NullableKalturaQuizFilter) Get() *KalturaQuizFilter {
	return v.value
}

func (v *NullableKalturaQuizFilter) Set(val *KalturaQuizFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuizFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuizFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuizFilter(val *KalturaQuizFilter) *NullableKalturaQuizFilter {
	return &NullableKalturaQuizFilter{value: val, isSet: true}
}

func (v NullableKalturaQuizFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuizFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



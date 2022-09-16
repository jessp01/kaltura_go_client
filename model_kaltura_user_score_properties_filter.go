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

// KalturaUserScorePropertiesFilter struct for KalturaUserScorePropertiesFilter
type KalturaUserScorePropertiesFilter struct {
	KalturaUserScorePropertiesBaseFilter
}

// NewKalturaUserScorePropertiesFilter instantiates a new KalturaUserScorePropertiesFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserScorePropertiesFilter() *KalturaUserScorePropertiesFilter {
	this := KalturaUserScorePropertiesFilter{}
	return &this
}

// NewKalturaUserScorePropertiesFilterWithDefaults instantiates a new KalturaUserScorePropertiesFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserScorePropertiesFilterWithDefaults() *KalturaUserScorePropertiesFilter {
	this := KalturaUserScorePropertiesFilter{}
	return &this
}

func (o KalturaUserScorePropertiesFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserScorePropertiesBaseFilter, errKalturaUserScorePropertiesBaseFilter := json.Marshal(o.KalturaUserScorePropertiesBaseFilter)
	if errKalturaUserScorePropertiesBaseFilter != nil {
		return []byte{}, errKalturaUserScorePropertiesBaseFilter
	}
	errKalturaUserScorePropertiesBaseFilter = json.Unmarshal([]byte(serializedKalturaUserScorePropertiesBaseFilter), &toSerialize)
	if errKalturaUserScorePropertiesBaseFilter != nil {
		return []byte{}, errKalturaUserScorePropertiesBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUserScorePropertiesFilter struct {
	value *KalturaUserScorePropertiesFilter
	isSet bool
}

func (v NullableKalturaUserScorePropertiesFilter) Get() *KalturaUserScorePropertiesFilter {
	return v.value
}

func (v *NullableKalturaUserScorePropertiesFilter) Set(val *KalturaUserScorePropertiesFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserScorePropertiesFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserScorePropertiesFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserScorePropertiesFilter(val *KalturaUserScorePropertiesFilter) *NullableKalturaUserScorePropertiesFilter {
	return &NullableKalturaUserScorePropertiesFilter{value: val, isSet: true}
}

func (v NullableKalturaUserScorePropertiesFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserScorePropertiesFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



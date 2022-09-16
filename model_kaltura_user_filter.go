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

// KalturaUserFilter struct for KalturaUserFilter
type KalturaUserFilter struct {
	KalturaUserBaseFilter
}

// NewKalturaUserFilter instantiates a new KalturaUserFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserFilter() *KalturaUserFilter {
	this := KalturaUserFilter{}
	return &this
}

// NewKalturaUserFilterWithDefaults instantiates a new KalturaUserFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserFilterWithDefaults() *KalturaUserFilter {
	this := KalturaUserFilter{}
	return &this
}

func (o KalturaUserFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserBaseFilter, errKalturaUserBaseFilter := json.Marshal(o.KalturaUserBaseFilter)
	if errKalturaUserBaseFilter != nil {
		return []byte{}, errKalturaUserBaseFilter
	}
	errKalturaUserBaseFilter = json.Unmarshal([]byte(serializedKalturaUserBaseFilter), &toSerialize)
	if errKalturaUserBaseFilter != nil {
		return []byte{}, errKalturaUserBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUserFilter struct {
	value *KalturaUserFilter
	isSet bool
}

func (v NullableKalturaUserFilter) Get() *KalturaUserFilter {
	return v.value
}

func (v *NullableKalturaUserFilter) Set(val *KalturaUserFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserFilter(val *KalturaUserFilter) *NullableKalturaUserFilter {
	return &NullableKalturaUserFilter{value: val, isSet: true}
}

func (v NullableKalturaUserFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


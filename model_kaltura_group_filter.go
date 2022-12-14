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

// KalturaGroupFilter struct for KalturaGroupFilter
type KalturaGroupFilter struct {
	KalturaUserFilter
}

// NewKalturaGroupFilter instantiates a new KalturaGroupFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGroupFilter() *KalturaGroupFilter {
	this := KalturaGroupFilter{}
	return &this
}

// NewKalturaGroupFilterWithDefaults instantiates a new KalturaGroupFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGroupFilterWithDefaults() *KalturaGroupFilter {
	this := KalturaGroupFilter{}
	return &this
}

func (o KalturaGroupFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUserFilter, errKalturaUserFilter := json.Marshal(o.KalturaUserFilter)
	if errKalturaUserFilter != nil {
		return []byte{}, errKalturaUserFilter
	}
	errKalturaUserFilter = json.Unmarshal([]byte(serializedKalturaUserFilter), &toSerialize)
	if errKalturaUserFilter != nil {
		return []byte{}, errKalturaUserFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaGroupFilter struct {
	value *KalturaGroupFilter
	isSet bool
}

func (v NullableKalturaGroupFilter) Get() *KalturaGroupFilter {
	return v.value
}

func (v *NullableKalturaGroupFilter) Set(val *KalturaGroupFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGroupFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGroupFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGroupFilter(val *KalturaGroupFilter) *NullableKalturaGroupFilter {
	return &NullableKalturaGroupFilter{value: val, isSet: true}
}

func (v NullableKalturaGroupFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGroupFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



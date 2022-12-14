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

// KalturaIndexAdvancedFilter struct for KalturaIndexAdvancedFilter
type KalturaIndexAdvancedFilter struct {
	KalturaSearchItem
}

// NewKalturaIndexAdvancedFilter instantiates a new KalturaIndexAdvancedFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaIndexAdvancedFilter() *KalturaIndexAdvancedFilter {
	this := KalturaIndexAdvancedFilter{}
	return &this
}

// NewKalturaIndexAdvancedFilterWithDefaults instantiates a new KalturaIndexAdvancedFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaIndexAdvancedFilterWithDefaults() *KalturaIndexAdvancedFilter {
	this := KalturaIndexAdvancedFilter{}
	return &this
}

func (o KalturaIndexAdvancedFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaIndexAdvancedFilter struct {
	value *KalturaIndexAdvancedFilter
	isSet bool
}

func (v NullableKalturaIndexAdvancedFilter) Get() *KalturaIndexAdvancedFilter {
	return v.value
}

func (v *NullableKalturaIndexAdvancedFilter) Set(val *KalturaIndexAdvancedFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaIndexAdvancedFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaIndexAdvancedFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaIndexAdvancedFilter(val *KalturaIndexAdvancedFilter) *NullableKalturaIndexAdvancedFilter {
	return &NullableKalturaIndexAdvancedFilter{value: val, isSet: true}
}

func (v NullableKalturaIndexAdvancedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaIndexAdvancedFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



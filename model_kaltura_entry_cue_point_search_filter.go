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

// KalturaEntryCuePointSearchFilter struct for KalturaEntryCuePointSearchFilter
type KalturaEntryCuePointSearchFilter struct {
	KalturaSearchItem
}

// NewKalturaEntryCuePointSearchFilter instantiates a new KalturaEntryCuePointSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaEntryCuePointSearchFilter() *KalturaEntryCuePointSearchFilter {
	this := KalturaEntryCuePointSearchFilter{}
	return &this
}

// NewKalturaEntryCuePointSearchFilterWithDefaults instantiates a new KalturaEntryCuePointSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaEntryCuePointSearchFilterWithDefaults() *KalturaEntryCuePointSearchFilter {
	this := KalturaEntryCuePointSearchFilter{}
	return &this
}

func (o KalturaEntryCuePointSearchFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaEntryCuePointSearchFilter struct {
	value *KalturaEntryCuePointSearchFilter
	isSet bool
}

func (v NullableKalturaEntryCuePointSearchFilter) Get() *KalturaEntryCuePointSearchFilter {
	return v.value
}

func (v *NullableKalturaEntryCuePointSearchFilter) Set(val *KalturaEntryCuePointSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaEntryCuePointSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaEntryCuePointSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaEntryCuePointSearchFilter(val *KalturaEntryCuePointSearchFilter) *NullableKalturaEntryCuePointSearchFilter {
	return &NullableKalturaEntryCuePointSearchFilter{value: val, isSet: true}
}

func (v NullableKalturaEntryCuePointSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaEntryCuePointSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



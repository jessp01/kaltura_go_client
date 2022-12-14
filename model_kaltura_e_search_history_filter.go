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

// KalturaESearchHistoryFilter struct for KalturaESearchHistoryFilter
type KalturaESearchHistoryFilter struct {
	KalturaESearchBaseFilter
}

// NewKalturaESearchHistoryFilter instantiates a new KalturaESearchHistoryFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchHistoryFilter() *KalturaESearchHistoryFilter {
	this := KalturaESearchHistoryFilter{}
	return &this
}

// NewKalturaESearchHistoryFilterWithDefaults instantiates a new KalturaESearchHistoryFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchHistoryFilterWithDefaults() *KalturaESearchHistoryFilter {
	this := KalturaESearchHistoryFilter{}
	return &this
}

func (o KalturaESearchHistoryFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaESearchBaseFilter, errKalturaESearchBaseFilter := json.Marshal(o.KalturaESearchBaseFilter)
	if errKalturaESearchBaseFilter != nil {
		return []byte{}, errKalturaESearchBaseFilter
	}
	errKalturaESearchBaseFilter = json.Unmarshal([]byte(serializedKalturaESearchBaseFilter), &toSerialize)
	if errKalturaESearchBaseFilter != nil {
		return []byte{}, errKalturaESearchBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchHistoryFilter struct {
	value *KalturaESearchHistoryFilter
	isSet bool
}

func (v NullableKalturaESearchHistoryFilter) Get() *KalturaESearchHistoryFilter {
	return v.value
}

func (v *NullableKalturaESearchHistoryFilter) Set(val *KalturaESearchHistoryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchHistoryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchHistoryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchHistoryFilter(val *KalturaESearchHistoryFilter) *NullableKalturaESearchHistoryFilter {
	return &NullableKalturaESearchHistoryFilter{value: val, isSet: true}
}

func (v NullableKalturaESearchHistoryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchHistoryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



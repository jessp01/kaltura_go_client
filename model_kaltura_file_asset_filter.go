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

// KalturaFileAssetFilter struct for KalturaFileAssetFilter
type KalturaFileAssetFilter struct {
	KalturaFileAssetBaseFilter
}

// NewKalturaFileAssetFilter instantiates a new KalturaFileAssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFileAssetFilter() *KalturaFileAssetFilter {
	this := KalturaFileAssetFilter{}
	return &this
}

// NewKalturaFileAssetFilterWithDefaults instantiates a new KalturaFileAssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFileAssetFilterWithDefaults() *KalturaFileAssetFilter {
	this := KalturaFileAssetFilter{}
	return &this
}

func (o KalturaFileAssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFileAssetBaseFilter, errKalturaFileAssetBaseFilter := json.Marshal(o.KalturaFileAssetBaseFilter)
	if errKalturaFileAssetBaseFilter != nil {
		return []byte{}, errKalturaFileAssetBaseFilter
	}
	errKalturaFileAssetBaseFilter = json.Unmarshal([]byte(serializedKalturaFileAssetBaseFilter), &toSerialize)
	if errKalturaFileAssetBaseFilter != nil {
		return []byte{}, errKalturaFileAssetBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFileAssetFilter struct {
	value *KalturaFileAssetFilter
	isSet bool
}

func (v NullableKalturaFileAssetFilter) Get() *KalturaFileAssetFilter {
	return v.value
}

func (v *NullableKalturaFileAssetFilter) Set(val *KalturaFileAssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFileAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFileAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFileAssetFilter(val *KalturaFileAssetFilter) *NullableKalturaFileAssetFilter {
	return &NullableKalturaFileAssetFilter{value: val, isSet: true}
}

func (v NullableKalturaFileAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFileAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaAssetFilter struct for KalturaAssetFilter
type KalturaAssetFilter struct {
	KalturaAssetBaseFilter
}

// NewKalturaAssetFilter instantiates a new KalturaAssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetFilter() *KalturaAssetFilter {
	this := KalturaAssetFilter{}
	return &this
}

// NewKalturaAssetFilterWithDefaults instantiates a new KalturaAssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetFilterWithDefaults() *KalturaAssetFilter {
	this := KalturaAssetFilter{}
	return &this
}

func (o KalturaAssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAssetBaseFilter, errKalturaAssetBaseFilter := json.Marshal(o.KalturaAssetBaseFilter)
	if errKalturaAssetBaseFilter != nil {
		return []byte{}, errKalturaAssetBaseFilter
	}
	errKalturaAssetBaseFilter = json.Unmarshal([]byte(serializedKalturaAssetBaseFilter), &toSerialize)
	if errKalturaAssetBaseFilter != nil {
		return []byte{}, errKalturaAssetBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAssetFilter struct {
	value *KalturaAssetFilter
	isSet bool
}

func (v NullableKalturaAssetFilter) Get() *KalturaAssetFilter {
	return v.value
}

func (v *NullableKalturaAssetFilter) Set(val *KalturaAssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetFilter(val *KalturaAssetFilter) *NullableKalturaAssetFilter {
	return &NullableKalturaAssetFilter{value: val, isSet: true}
}

func (v NullableKalturaAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



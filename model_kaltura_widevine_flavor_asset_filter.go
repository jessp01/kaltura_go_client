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

// KalturaWidevineFlavorAssetFilter struct for KalturaWidevineFlavorAssetFilter
type KalturaWidevineFlavorAssetFilter struct {
	KalturaWidevineFlavorAssetBaseFilter
}

// NewKalturaWidevineFlavorAssetFilter instantiates a new KalturaWidevineFlavorAssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWidevineFlavorAssetFilter() *KalturaWidevineFlavorAssetFilter {
	this := KalturaWidevineFlavorAssetFilter{}
	return &this
}

// NewKalturaWidevineFlavorAssetFilterWithDefaults instantiates a new KalturaWidevineFlavorAssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWidevineFlavorAssetFilterWithDefaults() *KalturaWidevineFlavorAssetFilter {
	this := KalturaWidevineFlavorAssetFilter{}
	return &this
}

func (o KalturaWidevineFlavorAssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaWidevineFlavorAssetBaseFilter, errKalturaWidevineFlavorAssetBaseFilter := json.Marshal(o.KalturaWidevineFlavorAssetBaseFilter)
	if errKalturaWidevineFlavorAssetBaseFilter != nil {
		return []byte{}, errKalturaWidevineFlavorAssetBaseFilter
	}
	errKalturaWidevineFlavorAssetBaseFilter = json.Unmarshal([]byte(serializedKalturaWidevineFlavorAssetBaseFilter), &toSerialize)
	if errKalturaWidevineFlavorAssetBaseFilter != nil {
		return []byte{}, errKalturaWidevineFlavorAssetBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWidevineFlavorAssetFilter struct {
	value *KalturaWidevineFlavorAssetFilter
	isSet bool
}

func (v NullableKalturaWidevineFlavorAssetFilter) Get() *KalturaWidevineFlavorAssetFilter {
	return v.value
}

func (v *NullableKalturaWidevineFlavorAssetFilter) Set(val *KalturaWidevineFlavorAssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWidevineFlavorAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWidevineFlavorAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWidevineFlavorAssetFilter(val *KalturaWidevineFlavorAssetFilter) *NullableKalturaWidevineFlavorAssetFilter {
	return &NullableKalturaWidevineFlavorAssetFilter{value: val, isSet: true}
}

func (v NullableKalturaWidevineFlavorAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWidevineFlavorAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



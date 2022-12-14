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

// KalturaFlavorAssetFilter struct for KalturaFlavorAssetFilter
type KalturaFlavorAssetFilter struct {
	KalturaFlavorAssetBaseFilter
}

// NewKalturaFlavorAssetFilter instantiates a new KalturaFlavorAssetFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFlavorAssetFilter() *KalturaFlavorAssetFilter {
	this := KalturaFlavorAssetFilter{}
	return &this
}

// NewKalturaFlavorAssetFilterWithDefaults instantiates a new KalturaFlavorAssetFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFlavorAssetFilterWithDefaults() *KalturaFlavorAssetFilter {
	this := KalturaFlavorAssetFilter{}
	return &this
}

func (o KalturaFlavorAssetFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorAssetBaseFilter, errKalturaFlavorAssetBaseFilter := json.Marshal(o.KalturaFlavorAssetBaseFilter)
	if errKalturaFlavorAssetBaseFilter != nil {
		return []byte{}, errKalturaFlavorAssetBaseFilter
	}
	errKalturaFlavorAssetBaseFilter = json.Unmarshal([]byte(serializedKalturaFlavorAssetBaseFilter), &toSerialize)
	if errKalturaFlavorAssetBaseFilter != nil {
		return []byte{}, errKalturaFlavorAssetBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFlavorAssetFilter struct {
	value *KalturaFlavorAssetFilter
	isSet bool
}

func (v NullableKalturaFlavorAssetFilter) Get() *KalturaFlavorAssetFilter {
	return v.value
}

func (v *NullableKalturaFlavorAssetFilter) Set(val *KalturaFlavorAssetFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFlavorAssetFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFlavorAssetFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFlavorAssetFilter(val *KalturaFlavorAssetFilter) *NullableKalturaFlavorAssetFilter {
	return &NullableKalturaFlavorAssetFilter{value: val, isSet: true}
}

func (v NullableKalturaFlavorAssetFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFlavorAssetFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



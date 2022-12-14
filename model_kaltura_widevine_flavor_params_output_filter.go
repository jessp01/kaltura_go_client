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

// KalturaWidevineFlavorParamsOutputFilter struct for KalturaWidevineFlavorParamsOutputFilter
type KalturaWidevineFlavorParamsOutputFilter struct {
	KalturaWidevineFlavorParamsOutputBaseFilter
}

// NewKalturaWidevineFlavorParamsOutputFilter instantiates a new KalturaWidevineFlavorParamsOutputFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWidevineFlavorParamsOutputFilter() *KalturaWidevineFlavorParamsOutputFilter {
	this := KalturaWidevineFlavorParamsOutputFilter{}
	return &this
}

// NewKalturaWidevineFlavorParamsOutputFilterWithDefaults instantiates a new KalturaWidevineFlavorParamsOutputFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWidevineFlavorParamsOutputFilterWithDefaults() *KalturaWidevineFlavorParamsOutputFilter {
	this := KalturaWidevineFlavorParamsOutputFilter{}
	return &this
}

func (o KalturaWidevineFlavorParamsOutputFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaWidevineFlavorParamsOutputBaseFilter, errKalturaWidevineFlavorParamsOutputBaseFilter := json.Marshal(o.KalturaWidevineFlavorParamsOutputBaseFilter)
	if errKalturaWidevineFlavorParamsOutputBaseFilter != nil {
		return []byte{}, errKalturaWidevineFlavorParamsOutputBaseFilter
	}
	errKalturaWidevineFlavorParamsOutputBaseFilter = json.Unmarshal([]byte(serializedKalturaWidevineFlavorParamsOutputBaseFilter), &toSerialize)
	if errKalturaWidevineFlavorParamsOutputBaseFilter != nil {
		return []byte{}, errKalturaWidevineFlavorParamsOutputBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWidevineFlavorParamsOutputFilter struct {
	value *KalturaWidevineFlavorParamsOutputFilter
	isSet bool
}

func (v NullableKalturaWidevineFlavorParamsOutputFilter) Get() *KalturaWidevineFlavorParamsOutputFilter {
	return v.value
}

func (v *NullableKalturaWidevineFlavorParamsOutputFilter) Set(val *KalturaWidevineFlavorParamsOutputFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWidevineFlavorParamsOutputFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWidevineFlavorParamsOutputFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWidevineFlavorParamsOutputFilter(val *KalturaWidevineFlavorParamsOutputFilter) *NullableKalturaWidevineFlavorParamsOutputFilter {
	return &NullableKalturaWidevineFlavorParamsOutputFilter{value: val, isSet: true}
}

func (v NullableKalturaWidevineFlavorParamsOutputFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWidevineFlavorParamsOutputFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



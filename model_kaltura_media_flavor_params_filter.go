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

// KalturaMediaFlavorParamsFilter struct for KalturaMediaFlavorParamsFilter
type KalturaMediaFlavorParamsFilter struct {
	KalturaMediaFlavorParamsBaseFilter
}

// NewKalturaMediaFlavorParamsFilter instantiates a new KalturaMediaFlavorParamsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMediaFlavorParamsFilter() *KalturaMediaFlavorParamsFilter {
	this := KalturaMediaFlavorParamsFilter{}
	return &this
}

// NewKalturaMediaFlavorParamsFilterWithDefaults instantiates a new KalturaMediaFlavorParamsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMediaFlavorParamsFilterWithDefaults() *KalturaMediaFlavorParamsFilter {
	this := KalturaMediaFlavorParamsFilter{}
	return &this
}

func (o KalturaMediaFlavorParamsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMediaFlavorParamsBaseFilter, errKalturaMediaFlavorParamsBaseFilter := json.Marshal(o.KalturaMediaFlavorParamsBaseFilter)
	if errKalturaMediaFlavorParamsBaseFilter != nil {
		return []byte{}, errKalturaMediaFlavorParamsBaseFilter
	}
	errKalturaMediaFlavorParamsBaseFilter = json.Unmarshal([]byte(serializedKalturaMediaFlavorParamsBaseFilter), &toSerialize)
	if errKalturaMediaFlavorParamsBaseFilter != nil {
		return []byte{}, errKalturaMediaFlavorParamsBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMediaFlavorParamsFilter struct {
	value *KalturaMediaFlavorParamsFilter
	isSet bool
}

func (v NullableKalturaMediaFlavorParamsFilter) Get() *KalturaMediaFlavorParamsFilter {
	return v.value
}

func (v *NullableKalturaMediaFlavorParamsFilter) Set(val *KalturaMediaFlavorParamsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMediaFlavorParamsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMediaFlavorParamsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMediaFlavorParamsFilter(val *KalturaMediaFlavorParamsFilter) *NullableKalturaMediaFlavorParamsFilter {
	return &NullableKalturaMediaFlavorParamsFilter{value: val, isSet: true}
}

func (v NullableKalturaMediaFlavorParamsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMediaFlavorParamsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


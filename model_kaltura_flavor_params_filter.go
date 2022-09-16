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

// KalturaFlavorParamsFilter struct for KalturaFlavorParamsFilter
type KalturaFlavorParamsFilter struct {
	KalturaFlavorParamsBaseFilter
}

// NewKalturaFlavorParamsFilter instantiates a new KalturaFlavorParamsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFlavorParamsFilter() *KalturaFlavorParamsFilter {
	this := KalturaFlavorParamsFilter{}
	return &this
}

// NewKalturaFlavorParamsFilterWithDefaults instantiates a new KalturaFlavorParamsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFlavorParamsFilterWithDefaults() *KalturaFlavorParamsFilter {
	this := KalturaFlavorParamsFilter{}
	return &this
}

func (o KalturaFlavorParamsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParamsBaseFilter, errKalturaFlavorParamsBaseFilter := json.Marshal(o.KalturaFlavorParamsBaseFilter)
	if errKalturaFlavorParamsBaseFilter != nil {
		return []byte{}, errKalturaFlavorParamsBaseFilter
	}
	errKalturaFlavorParamsBaseFilter = json.Unmarshal([]byte(serializedKalturaFlavorParamsBaseFilter), &toSerialize)
	if errKalturaFlavorParamsBaseFilter != nil {
		return []byte{}, errKalturaFlavorParamsBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFlavorParamsFilter struct {
	value *KalturaFlavorParamsFilter
	isSet bool
}

func (v NullableKalturaFlavorParamsFilter) Get() *KalturaFlavorParamsFilter {
	return v.value
}

func (v *NullableKalturaFlavorParamsFilter) Set(val *KalturaFlavorParamsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFlavorParamsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFlavorParamsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFlavorParamsFilter(val *KalturaFlavorParamsFilter) *NullableKalturaFlavorParamsFilter {
	return &NullableKalturaFlavorParamsFilter{value: val, isSet: true}
}

func (v NullableKalturaFlavorParamsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFlavorParamsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



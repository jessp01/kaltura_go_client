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

// KalturaImageFlavorParamsOutputBaseFilter `abstract`
type KalturaImageFlavorParamsOutputBaseFilter struct {
	KalturaFlavorParamsOutputFilter
}

// NewKalturaImageFlavorParamsOutputBaseFilter instantiates a new KalturaImageFlavorParamsOutputBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaImageFlavorParamsOutputBaseFilter() *KalturaImageFlavorParamsOutputBaseFilter {
	this := KalturaImageFlavorParamsOutputBaseFilter{}
	return &this
}

// NewKalturaImageFlavorParamsOutputBaseFilterWithDefaults instantiates a new KalturaImageFlavorParamsOutputBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaImageFlavorParamsOutputBaseFilterWithDefaults() *KalturaImageFlavorParamsOutputBaseFilter {
	this := KalturaImageFlavorParamsOutputBaseFilter{}
	return &this
}

func (o KalturaImageFlavorParamsOutputBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParamsOutputFilter, errKalturaFlavorParamsOutputFilter := json.Marshal(o.KalturaFlavorParamsOutputFilter)
	if errKalturaFlavorParamsOutputFilter != nil {
		return []byte{}, errKalturaFlavorParamsOutputFilter
	}
	errKalturaFlavorParamsOutputFilter = json.Unmarshal([]byte(serializedKalturaFlavorParamsOutputFilter), &toSerialize)
	if errKalturaFlavorParamsOutputFilter != nil {
		return []byte{}, errKalturaFlavorParamsOutputFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaImageFlavorParamsOutputBaseFilter struct {
	value *KalturaImageFlavorParamsOutputBaseFilter
	isSet bool
}

func (v NullableKalturaImageFlavorParamsOutputBaseFilter) Get() *KalturaImageFlavorParamsOutputBaseFilter {
	return v.value
}

func (v *NullableKalturaImageFlavorParamsOutputBaseFilter) Set(val *KalturaImageFlavorParamsOutputBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaImageFlavorParamsOutputBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaImageFlavorParamsOutputBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaImageFlavorParamsOutputBaseFilter(val *KalturaImageFlavorParamsOutputBaseFilter) *NullableKalturaImageFlavorParamsOutputBaseFilter {
	return &NullableKalturaImageFlavorParamsOutputBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaImageFlavorParamsOutputBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaImageFlavorParamsOutputBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


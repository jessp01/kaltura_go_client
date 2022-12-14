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

// KalturaFlavorParamsOutputFilter struct for KalturaFlavorParamsOutputFilter
type KalturaFlavorParamsOutputFilter struct {
	KalturaFlavorParamsOutputBaseFilter
}

// NewKalturaFlavorParamsOutputFilter instantiates a new KalturaFlavorParamsOutputFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFlavorParamsOutputFilter() *KalturaFlavorParamsOutputFilter {
	this := KalturaFlavorParamsOutputFilter{}
	return &this
}

// NewKalturaFlavorParamsOutputFilterWithDefaults instantiates a new KalturaFlavorParamsOutputFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFlavorParamsOutputFilterWithDefaults() *KalturaFlavorParamsOutputFilter {
	this := KalturaFlavorParamsOutputFilter{}
	return &this
}

func (o KalturaFlavorParamsOutputFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParamsOutputBaseFilter, errKalturaFlavorParamsOutputBaseFilter := json.Marshal(o.KalturaFlavorParamsOutputBaseFilter)
	if errKalturaFlavorParamsOutputBaseFilter != nil {
		return []byte{}, errKalturaFlavorParamsOutputBaseFilter
	}
	errKalturaFlavorParamsOutputBaseFilter = json.Unmarshal([]byte(serializedKalturaFlavorParamsOutputBaseFilter), &toSerialize)
	if errKalturaFlavorParamsOutputBaseFilter != nil {
		return []byte{}, errKalturaFlavorParamsOutputBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFlavorParamsOutputFilter struct {
	value *KalturaFlavorParamsOutputFilter
	isSet bool
}

func (v NullableKalturaFlavorParamsOutputFilter) Get() *KalturaFlavorParamsOutputFilter {
	return v.value
}

func (v *NullableKalturaFlavorParamsOutputFilter) Set(val *KalturaFlavorParamsOutputFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFlavorParamsOutputFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFlavorParamsOutputFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFlavorParamsOutputFilter(val *KalturaFlavorParamsOutputFilter) *NullableKalturaFlavorParamsOutputFilter {
	return &NullableKalturaFlavorParamsOutputFilter{value: val, isSet: true}
}

func (v NullableKalturaFlavorParamsOutputFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFlavorParamsOutputFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaWidevineFlavorParamsBaseFilter `abstract`
type KalturaWidevineFlavorParamsBaseFilter struct {
	KalturaFlavorParamsFilter
}

// NewKalturaWidevineFlavorParamsBaseFilter instantiates a new KalturaWidevineFlavorParamsBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWidevineFlavorParamsBaseFilter() *KalturaWidevineFlavorParamsBaseFilter {
	this := KalturaWidevineFlavorParamsBaseFilter{}
	return &this
}

// NewKalturaWidevineFlavorParamsBaseFilterWithDefaults instantiates a new KalturaWidevineFlavorParamsBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWidevineFlavorParamsBaseFilterWithDefaults() *KalturaWidevineFlavorParamsBaseFilter {
	this := KalturaWidevineFlavorParamsBaseFilter{}
	return &this
}

func (o KalturaWidevineFlavorParamsBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFlavorParamsFilter, errKalturaFlavorParamsFilter := json.Marshal(o.KalturaFlavorParamsFilter)
	if errKalturaFlavorParamsFilter != nil {
		return []byte{}, errKalturaFlavorParamsFilter
	}
	errKalturaFlavorParamsFilter = json.Unmarshal([]byte(serializedKalturaFlavorParamsFilter), &toSerialize)
	if errKalturaFlavorParamsFilter != nil {
		return []byte{}, errKalturaFlavorParamsFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWidevineFlavorParamsBaseFilter struct {
	value *KalturaWidevineFlavorParamsBaseFilter
	isSet bool
}

func (v NullableKalturaWidevineFlavorParamsBaseFilter) Get() *KalturaWidevineFlavorParamsBaseFilter {
	return v.value
}

func (v *NullableKalturaWidevineFlavorParamsBaseFilter) Set(val *KalturaWidevineFlavorParamsBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWidevineFlavorParamsBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWidevineFlavorParamsBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWidevineFlavorParamsBaseFilter(val *KalturaWidevineFlavorParamsBaseFilter) *NullableKalturaWidevineFlavorParamsBaseFilter {
	return &NullableKalturaWidevineFlavorParamsBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaWidevineFlavorParamsBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWidevineFlavorParamsBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// KalturaBatchJobFilter struct for KalturaBatchJobFilter
type KalturaBatchJobFilter struct {
	KalturaBatchJobBaseFilter
}

// NewKalturaBatchJobFilter instantiates a new KalturaBatchJobFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBatchJobFilter() *KalturaBatchJobFilter {
	this := KalturaBatchJobFilter{}
	return &this
}

// NewKalturaBatchJobFilterWithDefaults instantiates a new KalturaBatchJobFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBatchJobFilterWithDefaults() *KalturaBatchJobFilter {
	this := KalturaBatchJobFilter{}
	return &this
}

func (o KalturaBatchJobFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBatchJobBaseFilter, errKalturaBatchJobBaseFilter := json.Marshal(o.KalturaBatchJobBaseFilter)
	if errKalturaBatchJobBaseFilter != nil {
		return []byte{}, errKalturaBatchJobBaseFilter
	}
	errKalturaBatchJobBaseFilter = json.Unmarshal([]byte(serializedKalturaBatchJobBaseFilter), &toSerialize)
	if errKalturaBatchJobBaseFilter != nil {
		return []byte{}, errKalturaBatchJobBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBatchJobFilter struct {
	value *KalturaBatchJobFilter
	isSet bool
}

func (v NullableKalturaBatchJobFilter) Get() *KalturaBatchJobFilter {
	return v.value
}

func (v *NullableKalturaBatchJobFilter) Set(val *KalturaBatchJobFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBatchJobFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBatchJobFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBatchJobFilter(val *KalturaBatchJobFilter) *NullableKalturaBatchJobFilter {
	return &NullableKalturaBatchJobFilter{value: val, isSet: true}
}

func (v NullableKalturaBatchJobFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBatchJobFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaBatchJobFilterExt struct for KalturaBatchJobFilterExt
type KalturaBatchJobFilterExt struct {
	KalturaBatchJobFilter
}

// NewKalturaBatchJobFilterExt instantiates a new KalturaBatchJobFilterExt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBatchJobFilterExt() *KalturaBatchJobFilterExt {
	this := KalturaBatchJobFilterExt{}
	return &this
}

// NewKalturaBatchJobFilterExtWithDefaults instantiates a new KalturaBatchJobFilterExt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBatchJobFilterExtWithDefaults() *KalturaBatchJobFilterExt {
	this := KalturaBatchJobFilterExt{}
	return &this
}

func (o KalturaBatchJobFilterExt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBatchJobFilter, errKalturaBatchJobFilter := json.Marshal(o.KalturaBatchJobFilter)
	if errKalturaBatchJobFilter != nil {
		return []byte{}, errKalturaBatchJobFilter
	}
	errKalturaBatchJobFilter = json.Unmarshal([]byte(serializedKalturaBatchJobFilter), &toSerialize)
	if errKalturaBatchJobFilter != nil {
		return []byte{}, errKalturaBatchJobFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBatchJobFilterExt struct {
	value *KalturaBatchJobFilterExt
	isSet bool
}

func (v NullableKalturaBatchJobFilterExt) Get() *KalturaBatchJobFilterExt {
	return v.value
}

func (v *NullableKalturaBatchJobFilterExt) Set(val *KalturaBatchJobFilterExt) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBatchJobFilterExt) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBatchJobFilterExt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBatchJobFilterExt(val *KalturaBatchJobFilterExt) *NullableKalturaBatchJobFilterExt {
	return &NullableKalturaBatchJobFilterExt{value: val, isSet: true}
}

func (v NullableKalturaBatchJobFilterExt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBatchJobFilterExt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



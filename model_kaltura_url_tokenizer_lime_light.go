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

// KalturaUrlTokenizerLimeLight struct for KalturaUrlTokenizerLimeLight
type KalturaUrlTokenizerLimeLight struct {
	KalturaUrlTokenizer
}

// NewKalturaUrlTokenizerLimeLight instantiates a new KalturaUrlTokenizerLimeLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlTokenizerLimeLight() *KalturaUrlTokenizerLimeLight {
	this := KalturaUrlTokenizerLimeLight{}
	return &this
}

// NewKalturaUrlTokenizerLimeLightWithDefaults instantiates a new KalturaUrlTokenizerLimeLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlTokenizerLimeLightWithDefaults() *KalturaUrlTokenizerLimeLight {
	this := KalturaUrlTokenizerLimeLight{}
	return &this
}

func (o KalturaUrlTokenizerLimeLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUrlTokenizer, errKalturaUrlTokenizer := json.Marshal(o.KalturaUrlTokenizer)
	if errKalturaUrlTokenizer != nil {
		return []byte{}, errKalturaUrlTokenizer
	}
	errKalturaUrlTokenizer = json.Unmarshal([]byte(serializedKalturaUrlTokenizer), &toSerialize)
	if errKalturaUrlTokenizer != nil {
		return []byte{}, errKalturaUrlTokenizer
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUrlTokenizerLimeLight struct {
	value *KalturaUrlTokenizerLimeLight
	isSet bool
}

func (v NullableKalturaUrlTokenizerLimeLight) Get() *KalturaUrlTokenizerLimeLight {
	return v.value
}

func (v *NullableKalturaUrlTokenizerLimeLight) Set(val *KalturaUrlTokenizerLimeLight) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlTokenizerLimeLight) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlTokenizerLimeLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlTokenizerLimeLight(val *KalturaUrlTokenizerLimeLight) *NullableKalturaUrlTokenizerLimeLight {
	return &NullableKalturaUrlTokenizerLimeLight{value: val, isSet: true}
}

func (v NullableKalturaUrlTokenizerLimeLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlTokenizerLimeLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


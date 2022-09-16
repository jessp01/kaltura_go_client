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

// KalturaUrlTokenizerL3 struct for KalturaUrlTokenizerL3
type KalturaUrlTokenizerL3 struct {
	KalturaUrlTokenizer
}

// NewKalturaUrlTokenizerL3 instantiates a new KalturaUrlTokenizerL3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlTokenizerL3() *KalturaUrlTokenizerL3 {
	this := KalturaUrlTokenizerL3{}
	return &this
}

// NewKalturaUrlTokenizerL3WithDefaults instantiates a new KalturaUrlTokenizerL3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlTokenizerL3WithDefaults() *KalturaUrlTokenizerL3 {
	this := KalturaUrlTokenizerL3{}
	return &this
}

func (o KalturaUrlTokenizerL3) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUrlTokenizerL3 struct {
	value *KalturaUrlTokenizerL3
	isSet bool
}

func (v NullableKalturaUrlTokenizerL3) Get() *KalturaUrlTokenizerL3 {
	return v.value
}

func (v *NullableKalturaUrlTokenizerL3) Set(val *KalturaUrlTokenizerL3) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlTokenizerL3) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlTokenizerL3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlTokenizerL3(val *KalturaUrlTokenizerL3) *NullableKalturaUrlTokenizerL3 {
	return &NullableKalturaUrlTokenizerL3{value: val, isSet: true}
}

func (v NullableKalturaUrlTokenizerL3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlTokenizerL3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


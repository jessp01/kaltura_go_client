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

// KalturaUrlTokenizerLevel3 struct for KalturaUrlTokenizerLevel3
type KalturaUrlTokenizerLevel3 struct {
	KalturaUrlTokenizer
}

// NewKalturaUrlTokenizerLevel3 instantiates a new KalturaUrlTokenizerLevel3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlTokenizerLevel3() *KalturaUrlTokenizerLevel3 {
	this := KalturaUrlTokenizerLevel3{}
	return &this
}

// NewKalturaUrlTokenizerLevel3WithDefaults instantiates a new KalturaUrlTokenizerLevel3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlTokenizerLevel3WithDefaults() *KalturaUrlTokenizerLevel3 {
	this := KalturaUrlTokenizerLevel3{}
	return &this
}

func (o KalturaUrlTokenizerLevel3) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUrlTokenizerLevel3 struct {
	value *KalturaUrlTokenizerLevel3
	isSet bool
}

func (v NullableKalturaUrlTokenizerLevel3) Get() *KalturaUrlTokenizerLevel3 {
	return v.value
}

func (v *NullableKalturaUrlTokenizerLevel3) Set(val *KalturaUrlTokenizerLevel3) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlTokenizerLevel3) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlTokenizerLevel3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlTokenizerLevel3(val *KalturaUrlTokenizerLevel3) *NullableKalturaUrlTokenizerLevel3 {
	return &NullableKalturaUrlTokenizerLevel3{value: val, isSet: true}
}

func (v NullableKalturaUrlTokenizerLevel3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlTokenizerLevel3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



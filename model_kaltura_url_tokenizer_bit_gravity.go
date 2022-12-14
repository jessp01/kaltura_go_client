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

// KalturaUrlTokenizerBitGravity struct for KalturaUrlTokenizerBitGravity
type KalturaUrlTokenizerBitGravity struct {
	KalturaUrlTokenizer
}

// NewKalturaUrlTokenizerBitGravity instantiates a new KalturaUrlTokenizerBitGravity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUrlTokenizerBitGravity() *KalturaUrlTokenizerBitGravity {
	this := KalturaUrlTokenizerBitGravity{}
	return &this
}

// NewKalturaUrlTokenizerBitGravityWithDefaults instantiates a new KalturaUrlTokenizerBitGravity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUrlTokenizerBitGravityWithDefaults() *KalturaUrlTokenizerBitGravity {
	this := KalturaUrlTokenizerBitGravity{}
	return &this
}

func (o KalturaUrlTokenizerBitGravity) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUrlTokenizerBitGravity struct {
	value *KalturaUrlTokenizerBitGravity
	isSet bool
}

func (v NullableKalturaUrlTokenizerBitGravity) Get() *KalturaUrlTokenizerBitGravity {
	return v.value
}

func (v *NullableKalturaUrlTokenizerBitGravity) Set(val *KalturaUrlTokenizerBitGravity) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUrlTokenizerBitGravity) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUrlTokenizerBitGravity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUrlTokenizerBitGravity(val *KalturaUrlTokenizerBitGravity) *NullableKalturaUrlTokenizerBitGravity {
	return &NullableKalturaUrlTokenizerBitGravity{value: val, isSet: true}
}

func (v NullableKalturaUrlTokenizerBitGravity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUrlTokenizerBitGravity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



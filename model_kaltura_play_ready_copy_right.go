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

// KalturaPlayReadyCopyRight struct for KalturaPlayReadyCopyRight
type KalturaPlayReadyCopyRight struct {
	KalturaPlayReadyRight
}

// NewKalturaPlayReadyCopyRight instantiates a new KalturaPlayReadyCopyRight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayReadyCopyRight() *KalturaPlayReadyCopyRight {
	this := KalturaPlayReadyCopyRight{}
	return &this
}

// NewKalturaPlayReadyCopyRightWithDefaults instantiates a new KalturaPlayReadyCopyRight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayReadyCopyRightWithDefaults() *KalturaPlayReadyCopyRight {
	this := KalturaPlayReadyCopyRight{}
	return &this
}

func (o KalturaPlayReadyCopyRight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaPlayReadyRight, errKalturaPlayReadyRight := json.Marshal(o.KalturaPlayReadyRight)
	if errKalturaPlayReadyRight != nil {
		return []byte{}, errKalturaPlayReadyRight
	}
	errKalturaPlayReadyRight = json.Unmarshal([]byte(serializedKalturaPlayReadyRight), &toSerialize)
	if errKalturaPlayReadyRight != nil {
		return []byte{}, errKalturaPlayReadyRight
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayReadyCopyRight struct {
	value *KalturaPlayReadyCopyRight
	isSet bool
}

func (v NullableKalturaPlayReadyCopyRight) Get() *KalturaPlayReadyCopyRight {
	return v.value
}

func (v *NullableKalturaPlayReadyCopyRight) Set(val *KalturaPlayReadyCopyRight) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayReadyCopyRight) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayReadyCopyRight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayReadyCopyRight(val *KalturaPlayReadyCopyRight) *NullableKalturaPlayReadyCopyRight {
	return &NullableKalturaPlayReadyCopyRight{value: val, isSet: true}
}

func (v NullableKalturaPlayReadyCopyRight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayReadyCopyRight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


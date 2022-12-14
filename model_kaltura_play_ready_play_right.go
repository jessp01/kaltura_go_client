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

// KalturaPlayReadyPlayRight struct for KalturaPlayReadyPlayRight
type KalturaPlayReadyPlayRight struct {
	KalturaPlayReadyRight
}

// NewKalturaPlayReadyPlayRight instantiates a new KalturaPlayReadyPlayRight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayReadyPlayRight() *KalturaPlayReadyPlayRight {
	this := KalturaPlayReadyPlayRight{}
	return &this
}

// NewKalturaPlayReadyPlayRightWithDefaults instantiates a new KalturaPlayReadyPlayRight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayReadyPlayRightWithDefaults() *KalturaPlayReadyPlayRight {
	this := KalturaPlayReadyPlayRight{}
	return &this
}

func (o KalturaPlayReadyPlayRight) MarshalJSON() ([]byte, error) {
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

type NullableKalturaPlayReadyPlayRight struct {
	value *KalturaPlayReadyPlayRight
	isSet bool
}

func (v NullableKalturaPlayReadyPlayRight) Get() *KalturaPlayReadyPlayRight {
	return v.value
}

func (v *NullableKalturaPlayReadyPlayRight) Set(val *KalturaPlayReadyPlayRight) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayReadyPlayRight) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayReadyPlayRight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayReadyPlayRight(val *KalturaPlayReadyPlayRight) *NullableKalturaPlayReadyPlayRight {
	return &NullableKalturaPlayReadyPlayRight{value: val, isSet: true}
}

func (v NullableKalturaPlayReadyPlayRight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayReadyPlayRight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



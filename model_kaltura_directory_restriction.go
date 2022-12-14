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

// KalturaDirectoryRestriction struct for KalturaDirectoryRestriction
type KalturaDirectoryRestriction struct {
	KalturaBaseRestriction
}

// NewKalturaDirectoryRestriction instantiates a new KalturaDirectoryRestriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDirectoryRestriction() *KalturaDirectoryRestriction {
	this := KalturaDirectoryRestriction{}
	return &this
}

// NewKalturaDirectoryRestrictionWithDefaults instantiates a new KalturaDirectoryRestriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDirectoryRestrictionWithDefaults() *KalturaDirectoryRestriction {
	this := KalturaDirectoryRestriction{}
	return &this
}

func (o KalturaDirectoryRestriction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBaseRestriction, errKalturaBaseRestriction := json.Marshal(o.KalturaBaseRestriction)
	if errKalturaBaseRestriction != nil {
		return []byte{}, errKalturaBaseRestriction
	}
	errKalturaBaseRestriction = json.Unmarshal([]byte(serializedKalturaBaseRestriction), &toSerialize)
	if errKalturaBaseRestriction != nil {
		return []byte{}, errKalturaBaseRestriction
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDirectoryRestriction struct {
	value *KalturaDirectoryRestriction
	isSet bool
}

func (v NullableKalturaDirectoryRestriction) Get() *KalturaDirectoryRestriction {
	return v.value
}

func (v *NullableKalturaDirectoryRestriction) Set(val *KalturaDirectoryRestriction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDirectoryRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDirectoryRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDirectoryRestriction(val *KalturaDirectoryRestriction) *NullableKalturaDirectoryRestriction {
	return &NullableKalturaDirectoryRestriction{value: val, isSet: true}
}

func (v NullableKalturaDirectoryRestriction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDirectoryRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



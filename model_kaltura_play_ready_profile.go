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

// KalturaPlayReadyProfile struct for KalturaPlayReadyProfile
type KalturaPlayReadyProfile struct {
	KalturaDrmProfile
}

// NewKalturaPlayReadyProfile instantiates a new KalturaPlayReadyProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayReadyProfile() *KalturaPlayReadyProfile {
	this := KalturaPlayReadyProfile{}
	return &this
}

// NewKalturaPlayReadyProfileWithDefaults instantiates a new KalturaPlayReadyProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayReadyProfileWithDefaults() *KalturaPlayReadyProfile {
	this := KalturaPlayReadyProfile{}
	return &this
}

func (o KalturaPlayReadyProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDrmProfile, errKalturaDrmProfile := json.Marshal(o.KalturaDrmProfile)
	if errKalturaDrmProfile != nil {
		return []byte{}, errKalturaDrmProfile
	}
	errKalturaDrmProfile = json.Unmarshal([]byte(serializedKalturaDrmProfile), &toSerialize)
	if errKalturaDrmProfile != nil {
		return []byte{}, errKalturaDrmProfile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayReadyProfile struct {
	value *KalturaPlayReadyProfile
	isSet bool
}

func (v NullableKalturaPlayReadyProfile) Get() *KalturaPlayReadyProfile {
	return v.value
}

func (v *NullableKalturaPlayReadyProfile) Set(val *KalturaPlayReadyProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayReadyProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayReadyProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayReadyProfile(val *KalturaPlayReadyProfile) *NullableKalturaPlayReadyProfile {
	return &NullableKalturaPlayReadyProfile{value: val, isSet: true}
}

func (v NullableKalturaPlayReadyProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayReadyProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



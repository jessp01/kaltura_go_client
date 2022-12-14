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

// KalturaFairplayDrmProfile struct for KalturaFairplayDrmProfile
type KalturaFairplayDrmProfile struct {
	KalturaDrmProfile
}

// NewKalturaFairplayDrmProfile instantiates a new KalturaFairplayDrmProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFairplayDrmProfile() *KalturaFairplayDrmProfile {
	this := KalturaFairplayDrmProfile{}
	return &this
}

// NewKalturaFairplayDrmProfileWithDefaults instantiates a new KalturaFairplayDrmProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFairplayDrmProfileWithDefaults() *KalturaFairplayDrmProfile {
	this := KalturaFairplayDrmProfile{}
	return &this
}

func (o KalturaFairplayDrmProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFairplayDrmProfile struct {
	value *KalturaFairplayDrmProfile
	isSet bool
}

func (v NullableKalturaFairplayDrmProfile) Get() *KalturaFairplayDrmProfile {
	return v.value
}

func (v *NullableKalturaFairplayDrmProfile) Set(val *KalturaFairplayDrmProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFairplayDrmProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFairplayDrmProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFairplayDrmProfile(val *KalturaFairplayDrmProfile) *NullableKalturaFairplayDrmProfile {
	return &NullableKalturaFairplayDrmProfile{value: val, isSet: true}
}

func (v NullableKalturaFairplayDrmProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFairplayDrmProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



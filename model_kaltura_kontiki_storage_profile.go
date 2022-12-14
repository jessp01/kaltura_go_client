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

// KalturaKontikiStorageProfile struct for KalturaKontikiStorageProfile
type KalturaKontikiStorageProfile struct {
	KalturaStorageProfile
}

// NewKalturaKontikiStorageProfile instantiates a new KalturaKontikiStorageProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaKontikiStorageProfile() *KalturaKontikiStorageProfile {
	this := KalturaKontikiStorageProfile{}
	return &this
}

// NewKalturaKontikiStorageProfileWithDefaults instantiates a new KalturaKontikiStorageProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaKontikiStorageProfileWithDefaults() *KalturaKontikiStorageProfile {
	this := KalturaKontikiStorageProfile{}
	return &this
}

func (o KalturaKontikiStorageProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaStorageProfile, errKalturaStorageProfile := json.Marshal(o.KalturaStorageProfile)
	if errKalturaStorageProfile != nil {
		return []byte{}, errKalturaStorageProfile
	}
	errKalturaStorageProfile = json.Unmarshal([]byte(serializedKalturaStorageProfile), &toSerialize)
	if errKalturaStorageProfile != nil {
		return []byte{}, errKalturaStorageProfile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaKontikiStorageProfile struct {
	value *KalturaKontikiStorageProfile
	isSet bool
}

func (v NullableKalturaKontikiStorageProfile) Get() *KalturaKontikiStorageProfile {
	return v.value
}

func (v *NullableKalturaKontikiStorageProfile) Set(val *KalturaKontikiStorageProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaKontikiStorageProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaKontikiStorageProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaKontikiStorageProfile(val *KalturaKontikiStorageProfile) *NullableKalturaKontikiStorageProfile {
	return &NullableKalturaKontikiStorageProfile{value: val, isSet: true}
}

func (v NullableKalturaKontikiStorageProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaKontikiStorageProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



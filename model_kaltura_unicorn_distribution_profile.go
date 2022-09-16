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

// KalturaUnicornDistributionProfile struct for KalturaUnicornDistributionProfile
type KalturaUnicornDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaUnicornDistributionProfile instantiates a new KalturaUnicornDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUnicornDistributionProfile() *KalturaUnicornDistributionProfile {
	this := KalturaUnicornDistributionProfile{}
	return &this
}

// NewKalturaUnicornDistributionProfileWithDefaults instantiates a new KalturaUnicornDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUnicornDistributionProfileWithDefaults() *KalturaUnicornDistributionProfile {
	this := KalturaUnicornDistributionProfile{}
	return &this
}

func (o KalturaUnicornDistributionProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConfigurableDistributionProfile, errKalturaConfigurableDistributionProfile := json.Marshal(o.KalturaConfigurableDistributionProfile)
	if errKalturaConfigurableDistributionProfile != nil {
		return []byte{}, errKalturaConfigurableDistributionProfile
	}
	errKalturaConfigurableDistributionProfile = json.Unmarshal([]byte(serializedKalturaConfigurableDistributionProfile), &toSerialize)
	if errKalturaConfigurableDistributionProfile != nil {
		return []byte{}, errKalturaConfigurableDistributionProfile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUnicornDistributionProfile struct {
	value *KalturaUnicornDistributionProfile
	isSet bool
}

func (v NullableKalturaUnicornDistributionProfile) Get() *KalturaUnicornDistributionProfile {
	return v.value
}

func (v *NullableKalturaUnicornDistributionProfile) Set(val *KalturaUnicornDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUnicornDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUnicornDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUnicornDistributionProfile(val *KalturaUnicornDistributionProfile) *NullableKalturaUnicornDistributionProfile {
	return &NullableKalturaUnicornDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaUnicornDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUnicornDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



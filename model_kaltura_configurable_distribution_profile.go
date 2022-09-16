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

// KalturaConfigurableDistributionProfile `abstract`
type KalturaConfigurableDistributionProfile struct {
	KalturaDistributionProfile
}

// NewKalturaConfigurableDistributionProfile instantiates a new KalturaConfigurableDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConfigurableDistributionProfile() *KalturaConfigurableDistributionProfile {
	this := KalturaConfigurableDistributionProfile{}
	return &this
}

// NewKalturaConfigurableDistributionProfileWithDefaults instantiates a new KalturaConfigurableDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConfigurableDistributionProfileWithDefaults() *KalturaConfigurableDistributionProfile {
	this := KalturaConfigurableDistributionProfile{}
	return &this
}

func (o KalturaConfigurableDistributionProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionProfile, errKalturaDistributionProfile := json.Marshal(o.KalturaDistributionProfile)
	if errKalturaDistributionProfile != nil {
		return []byte{}, errKalturaDistributionProfile
	}
	errKalturaDistributionProfile = json.Unmarshal([]byte(serializedKalturaDistributionProfile), &toSerialize)
	if errKalturaDistributionProfile != nil {
		return []byte{}, errKalturaDistributionProfile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaConfigurableDistributionProfile struct {
	value *KalturaConfigurableDistributionProfile
	isSet bool
}

func (v NullableKalturaConfigurableDistributionProfile) Get() *KalturaConfigurableDistributionProfile {
	return v.value
}

func (v *NullableKalturaConfigurableDistributionProfile) Set(val *KalturaConfigurableDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConfigurableDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConfigurableDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConfigurableDistributionProfile(val *KalturaConfigurableDistributionProfile) *NullableKalturaConfigurableDistributionProfile {
	return &NullableKalturaConfigurableDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaConfigurableDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConfigurableDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


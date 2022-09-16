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

// KalturaComcastMrssDistributionProfile struct for KalturaComcastMrssDistributionProfile
type KalturaComcastMrssDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaComcastMrssDistributionProfile instantiates a new KalturaComcastMrssDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaComcastMrssDistributionProfile() *KalturaComcastMrssDistributionProfile {
	this := KalturaComcastMrssDistributionProfile{}
	return &this
}

// NewKalturaComcastMrssDistributionProfileWithDefaults instantiates a new KalturaComcastMrssDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaComcastMrssDistributionProfileWithDefaults() *KalturaComcastMrssDistributionProfile {
	this := KalturaComcastMrssDistributionProfile{}
	return &this
}

func (o KalturaComcastMrssDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaComcastMrssDistributionProfile struct {
	value *KalturaComcastMrssDistributionProfile
	isSet bool
}

func (v NullableKalturaComcastMrssDistributionProfile) Get() *KalturaComcastMrssDistributionProfile {
	return v.value
}

func (v *NullableKalturaComcastMrssDistributionProfile) Set(val *KalturaComcastMrssDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaComcastMrssDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaComcastMrssDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaComcastMrssDistributionProfile(val *KalturaComcastMrssDistributionProfile) *NullableKalturaComcastMrssDistributionProfile {
	return &NullableKalturaComcastMrssDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaComcastMrssDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaComcastMrssDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaAttUverseDistributionProfile struct for KalturaAttUverseDistributionProfile
type KalturaAttUverseDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaAttUverseDistributionProfile instantiates a new KalturaAttUverseDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttUverseDistributionProfile() *KalturaAttUverseDistributionProfile {
	this := KalturaAttUverseDistributionProfile{}
	return &this
}

// NewKalturaAttUverseDistributionProfileWithDefaults instantiates a new KalturaAttUverseDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttUverseDistributionProfileWithDefaults() *KalturaAttUverseDistributionProfile {
	this := KalturaAttUverseDistributionProfile{}
	return &this
}

func (o KalturaAttUverseDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAttUverseDistributionProfile struct {
	value *KalturaAttUverseDistributionProfile
	isSet bool
}

func (v NullableKalturaAttUverseDistributionProfile) Get() *KalturaAttUverseDistributionProfile {
	return v.value
}

func (v *NullableKalturaAttUverseDistributionProfile) Set(val *KalturaAttUverseDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttUverseDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttUverseDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttUverseDistributionProfile(val *KalturaAttUverseDistributionProfile) *NullableKalturaAttUverseDistributionProfile {
	return &NullableKalturaAttUverseDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaAttUverseDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttUverseDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



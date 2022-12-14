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

// KalturaTvinciDistributionProfile struct for KalturaTvinciDistributionProfile
type KalturaTvinciDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaTvinciDistributionProfile instantiates a new KalturaTvinciDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTvinciDistributionProfile() *KalturaTvinciDistributionProfile {
	this := KalturaTvinciDistributionProfile{}
	return &this
}

// NewKalturaTvinciDistributionProfileWithDefaults instantiates a new KalturaTvinciDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTvinciDistributionProfileWithDefaults() *KalturaTvinciDistributionProfile {
	this := KalturaTvinciDistributionProfile{}
	return &this
}

func (o KalturaTvinciDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaTvinciDistributionProfile struct {
	value *KalturaTvinciDistributionProfile
	isSet bool
}

func (v NullableKalturaTvinciDistributionProfile) Get() *KalturaTvinciDistributionProfile {
	return v.value
}

func (v *NullableKalturaTvinciDistributionProfile) Set(val *KalturaTvinciDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTvinciDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTvinciDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTvinciDistributionProfile(val *KalturaTvinciDistributionProfile) *NullableKalturaTvinciDistributionProfile {
	return &NullableKalturaTvinciDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaTvinciDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTvinciDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// KalturaFtpDistributionProfile struct for KalturaFtpDistributionProfile
type KalturaFtpDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaFtpDistributionProfile instantiates a new KalturaFtpDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFtpDistributionProfile() *KalturaFtpDistributionProfile {
	this := KalturaFtpDistributionProfile{}
	return &this
}

// NewKalturaFtpDistributionProfileWithDefaults instantiates a new KalturaFtpDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFtpDistributionProfileWithDefaults() *KalturaFtpDistributionProfile {
	this := KalturaFtpDistributionProfile{}
	return &this
}

func (o KalturaFtpDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFtpDistributionProfile struct {
	value *KalturaFtpDistributionProfile
	isSet bool
}

func (v NullableKalturaFtpDistributionProfile) Get() *KalturaFtpDistributionProfile {
	return v.value
}

func (v *NullableKalturaFtpDistributionProfile) Set(val *KalturaFtpDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFtpDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFtpDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFtpDistributionProfile(val *KalturaFtpDistributionProfile) *NullableKalturaFtpDistributionProfile {
	return &NullableKalturaFtpDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaFtpDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFtpDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



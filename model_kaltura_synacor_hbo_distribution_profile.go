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

// KalturaSynacorHboDistributionProfile struct for KalturaSynacorHboDistributionProfile
type KalturaSynacorHboDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaSynacorHboDistributionProfile instantiates a new KalturaSynacorHboDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSynacorHboDistributionProfile() *KalturaSynacorHboDistributionProfile {
	this := KalturaSynacorHboDistributionProfile{}
	return &this
}

// NewKalturaSynacorHboDistributionProfileWithDefaults instantiates a new KalturaSynacorHboDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSynacorHboDistributionProfileWithDefaults() *KalturaSynacorHboDistributionProfile {
	this := KalturaSynacorHboDistributionProfile{}
	return &this
}

func (o KalturaSynacorHboDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaSynacorHboDistributionProfile struct {
	value *KalturaSynacorHboDistributionProfile
	isSet bool
}

func (v NullableKalturaSynacorHboDistributionProfile) Get() *KalturaSynacorHboDistributionProfile {
	return v.value
}

func (v *NullableKalturaSynacorHboDistributionProfile) Set(val *KalturaSynacorHboDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSynacorHboDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSynacorHboDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSynacorHboDistributionProfile(val *KalturaSynacorHboDistributionProfile) *NullableKalturaSynacorHboDistributionProfile {
	return &NullableKalturaSynacorHboDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaSynacorHboDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSynacorHboDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



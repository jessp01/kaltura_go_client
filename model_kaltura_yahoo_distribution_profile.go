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

// KalturaYahooDistributionProfile struct for KalturaYahooDistributionProfile
type KalturaYahooDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaYahooDistributionProfile instantiates a new KalturaYahooDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaYahooDistributionProfile() *KalturaYahooDistributionProfile {
	this := KalturaYahooDistributionProfile{}
	return &this
}

// NewKalturaYahooDistributionProfileWithDefaults instantiates a new KalturaYahooDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaYahooDistributionProfileWithDefaults() *KalturaYahooDistributionProfile {
	this := KalturaYahooDistributionProfile{}
	return &this
}

func (o KalturaYahooDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaYahooDistributionProfile struct {
	value *KalturaYahooDistributionProfile
	isSet bool
}

func (v NullableKalturaYahooDistributionProfile) Get() *KalturaYahooDistributionProfile {
	return v.value
}

func (v *NullableKalturaYahooDistributionProfile) Set(val *KalturaYahooDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaYahooDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaYahooDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaYahooDistributionProfile(val *KalturaYahooDistributionProfile) *NullableKalturaYahooDistributionProfile {
	return &NullableKalturaYahooDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaYahooDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaYahooDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


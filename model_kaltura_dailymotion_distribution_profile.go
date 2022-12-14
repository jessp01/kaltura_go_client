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

// KalturaDailymotionDistributionProfile struct for KalturaDailymotionDistributionProfile
type KalturaDailymotionDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaDailymotionDistributionProfile instantiates a new KalturaDailymotionDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDailymotionDistributionProfile() *KalturaDailymotionDistributionProfile {
	this := KalturaDailymotionDistributionProfile{}
	return &this
}

// NewKalturaDailymotionDistributionProfileWithDefaults instantiates a new KalturaDailymotionDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDailymotionDistributionProfileWithDefaults() *KalturaDailymotionDistributionProfile {
	this := KalturaDailymotionDistributionProfile{}
	return &this
}

func (o KalturaDailymotionDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDailymotionDistributionProfile struct {
	value *KalturaDailymotionDistributionProfile
	isSet bool
}

func (v NullableKalturaDailymotionDistributionProfile) Get() *KalturaDailymotionDistributionProfile {
	return v.value
}

func (v *NullableKalturaDailymotionDistributionProfile) Set(val *KalturaDailymotionDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDailymotionDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDailymotionDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDailymotionDistributionProfile(val *KalturaDailymotionDistributionProfile) *NullableKalturaDailymotionDistributionProfile {
	return &NullableKalturaDailymotionDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaDailymotionDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDailymotionDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



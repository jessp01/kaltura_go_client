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

// KalturaYoutubeApiDistributionProfile struct for KalturaYoutubeApiDistributionProfile
type KalturaYoutubeApiDistributionProfile struct {
	KalturaConfigurableDistributionProfile
}

// NewKalturaYoutubeApiDistributionProfile instantiates a new KalturaYoutubeApiDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaYoutubeApiDistributionProfile() *KalturaYoutubeApiDistributionProfile {
	this := KalturaYoutubeApiDistributionProfile{}
	return &this
}

// NewKalturaYoutubeApiDistributionProfileWithDefaults instantiates a new KalturaYoutubeApiDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaYoutubeApiDistributionProfileWithDefaults() *KalturaYoutubeApiDistributionProfile {
	this := KalturaYoutubeApiDistributionProfile{}
	return &this
}

func (o KalturaYoutubeApiDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaYoutubeApiDistributionProfile struct {
	value *KalturaYoutubeApiDistributionProfile
	isSet bool
}

func (v NullableKalturaYoutubeApiDistributionProfile) Get() *KalturaYoutubeApiDistributionProfile {
	return v.value
}

func (v *NullableKalturaYoutubeApiDistributionProfile) Set(val *KalturaYoutubeApiDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaYoutubeApiDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaYoutubeApiDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaYoutubeApiDistributionProfile(val *KalturaYoutubeApiDistributionProfile) *NullableKalturaYoutubeApiDistributionProfile {
	return &NullableKalturaYoutubeApiDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaYoutubeApiDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaYoutubeApiDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



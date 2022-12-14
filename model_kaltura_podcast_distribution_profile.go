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

// KalturaPodcastDistributionProfile struct for KalturaPodcastDistributionProfile
type KalturaPodcastDistributionProfile struct {
	KalturaDistributionProfile
}

// NewKalturaPodcastDistributionProfile instantiates a new KalturaPodcastDistributionProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPodcastDistributionProfile() *KalturaPodcastDistributionProfile {
	this := KalturaPodcastDistributionProfile{}
	return &this
}

// NewKalturaPodcastDistributionProfileWithDefaults instantiates a new KalturaPodcastDistributionProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPodcastDistributionProfileWithDefaults() *KalturaPodcastDistributionProfile {
	this := KalturaPodcastDistributionProfile{}
	return &this
}

func (o KalturaPodcastDistributionProfile) MarshalJSON() ([]byte, error) {
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

type NullableKalturaPodcastDistributionProfile struct {
	value *KalturaPodcastDistributionProfile
	isSet bool
}

func (v NullableKalturaPodcastDistributionProfile) Get() *KalturaPodcastDistributionProfile {
	return v.value
}

func (v *NullableKalturaPodcastDistributionProfile) Set(val *KalturaPodcastDistributionProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPodcastDistributionProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPodcastDistributionProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPodcastDistributionProfile(val *KalturaPodcastDistributionProfile) *NullableKalturaPodcastDistributionProfile {
	return &NullableKalturaPodcastDistributionProfile{value: val, isSet: true}
}

func (v NullableKalturaPodcastDistributionProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPodcastDistributionProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



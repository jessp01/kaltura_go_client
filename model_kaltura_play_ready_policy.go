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

// KalturaPlayReadyPolicy struct for KalturaPlayReadyPolicy
type KalturaPlayReadyPolicy struct {
	KalturaDrmPolicy
}

// NewKalturaPlayReadyPolicy instantiates a new KalturaPlayReadyPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayReadyPolicy() *KalturaPlayReadyPolicy {
	this := KalturaPlayReadyPolicy{}
	return &this
}

// NewKalturaPlayReadyPolicyWithDefaults instantiates a new KalturaPlayReadyPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayReadyPolicyWithDefaults() *KalturaPlayReadyPolicy {
	this := KalturaPlayReadyPolicy{}
	return &this
}

func (o KalturaPlayReadyPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDrmPolicy, errKalturaDrmPolicy := json.Marshal(o.KalturaDrmPolicy)
	if errKalturaDrmPolicy != nil {
		return []byte{}, errKalturaDrmPolicy
	}
	errKalturaDrmPolicy = json.Unmarshal([]byte(serializedKalturaDrmPolicy), &toSerialize)
	if errKalturaDrmPolicy != nil {
		return []byte{}, errKalturaDrmPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayReadyPolicy struct {
	value *KalturaPlayReadyPolicy
	isSet bool
}

func (v NullableKalturaPlayReadyPolicy) Get() *KalturaPlayReadyPolicy {
	return v.value
}

func (v *NullableKalturaPlayReadyPolicy) Set(val *KalturaPlayReadyPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayReadyPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayReadyPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayReadyPolicy(val *KalturaPlayReadyPolicy) *NullableKalturaPlayReadyPolicy {
	return &NullableKalturaPlayReadyPolicy{value: val, isSet: true}
}

func (v NullableKalturaPlayReadyPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayReadyPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


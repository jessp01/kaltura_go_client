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

// KalturaMsnDistributionProfileBaseFilter `abstract`
type KalturaMsnDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaMsnDistributionProfileBaseFilter instantiates a new KalturaMsnDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMsnDistributionProfileBaseFilter() *KalturaMsnDistributionProfileBaseFilter {
	this := KalturaMsnDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaMsnDistributionProfileBaseFilterWithDefaults instantiates a new KalturaMsnDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMsnDistributionProfileBaseFilterWithDefaults() *KalturaMsnDistributionProfileBaseFilter {
	this := KalturaMsnDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaMsnDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConfigurableDistributionProfileFilter, errKalturaConfigurableDistributionProfileFilter := json.Marshal(o.KalturaConfigurableDistributionProfileFilter)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	errKalturaConfigurableDistributionProfileFilter = json.Unmarshal([]byte(serializedKalturaConfigurableDistributionProfileFilter), &toSerialize)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaMsnDistributionProfileBaseFilter struct {
	value *KalturaMsnDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaMsnDistributionProfileBaseFilter) Get() *KalturaMsnDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaMsnDistributionProfileBaseFilter) Set(val *KalturaMsnDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMsnDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMsnDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMsnDistributionProfileBaseFilter(val *KalturaMsnDistributionProfileBaseFilter) *NullableKalturaMsnDistributionProfileBaseFilter {
	return &NullableKalturaMsnDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaMsnDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMsnDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


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

// KalturaHuluDistributionProfileFilter struct for KalturaHuluDistributionProfileFilter
type KalturaHuluDistributionProfileFilter struct {
	KalturaHuluDistributionProfileBaseFilter
}

// NewKalturaHuluDistributionProfileFilter instantiates a new KalturaHuluDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHuluDistributionProfileFilter() *KalturaHuluDistributionProfileFilter {
	this := KalturaHuluDistributionProfileFilter{}
	return &this
}

// NewKalturaHuluDistributionProfileFilterWithDefaults instantiates a new KalturaHuluDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHuluDistributionProfileFilterWithDefaults() *KalturaHuluDistributionProfileFilter {
	this := KalturaHuluDistributionProfileFilter{}
	return &this
}

func (o KalturaHuluDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaHuluDistributionProfileBaseFilter, errKalturaHuluDistributionProfileBaseFilter := json.Marshal(o.KalturaHuluDistributionProfileBaseFilter)
	if errKalturaHuluDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaHuluDistributionProfileBaseFilter
	}
	errKalturaHuluDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaHuluDistributionProfileBaseFilter), &toSerialize)
	if errKalturaHuluDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaHuluDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaHuluDistributionProfileFilter struct {
	value *KalturaHuluDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaHuluDistributionProfileFilter) Get() *KalturaHuluDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaHuluDistributionProfileFilter) Set(val *KalturaHuluDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHuluDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHuluDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHuluDistributionProfileFilter(val *KalturaHuluDistributionProfileFilter) *NullableKalturaHuluDistributionProfileFilter {
	return &NullableKalturaHuluDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaHuluDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHuluDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



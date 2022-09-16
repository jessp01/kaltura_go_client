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

// KalturaComcastMrssDistributionProfileFilter struct for KalturaComcastMrssDistributionProfileFilter
type KalturaComcastMrssDistributionProfileFilter struct {
	KalturaComcastMrssDistributionProfileBaseFilter
}

// NewKalturaComcastMrssDistributionProfileFilter instantiates a new KalturaComcastMrssDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaComcastMrssDistributionProfileFilter() *KalturaComcastMrssDistributionProfileFilter {
	this := KalturaComcastMrssDistributionProfileFilter{}
	return &this
}

// NewKalturaComcastMrssDistributionProfileFilterWithDefaults instantiates a new KalturaComcastMrssDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaComcastMrssDistributionProfileFilterWithDefaults() *KalturaComcastMrssDistributionProfileFilter {
	this := KalturaComcastMrssDistributionProfileFilter{}
	return &this
}

func (o KalturaComcastMrssDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaComcastMrssDistributionProfileBaseFilter, errKalturaComcastMrssDistributionProfileBaseFilter := json.Marshal(o.KalturaComcastMrssDistributionProfileBaseFilter)
	if errKalturaComcastMrssDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaComcastMrssDistributionProfileBaseFilter
	}
	errKalturaComcastMrssDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaComcastMrssDistributionProfileBaseFilter), &toSerialize)
	if errKalturaComcastMrssDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaComcastMrssDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaComcastMrssDistributionProfileFilter struct {
	value *KalturaComcastMrssDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaComcastMrssDistributionProfileFilter) Get() *KalturaComcastMrssDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaComcastMrssDistributionProfileFilter) Set(val *KalturaComcastMrssDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaComcastMrssDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaComcastMrssDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaComcastMrssDistributionProfileFilter(val *KalturaComcastMrssDistributionProfileFilter) *NullableKalturaComcastMrssDistributionProfileFilter {
	return &NullableKalturaComcastMrssDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaComcastMrssDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaComcastMrssDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



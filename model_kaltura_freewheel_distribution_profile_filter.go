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

// KalturaFreewheelDistributionProfileFilter struct for KalturaFreewheelDistributionProfileFilter
type KalturaFreewheelDistributionProfileFilter struct {
	KalturaFreewheelDistributionProfileBaseFilter
}

// NewKalturaFreewheelDistributionProfileFilter instantiates a new KalturaFreewheelDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFreewheelDistributionProfileFilter() *KalturaFreewheelDistributionProfileFilter {
	this := KalturaFreewheelDistributionProfileFilter{}
	return &this
}

// NewKalturaFreewheelDistributionProfileFilterWithDefaults instantiates a new KalturaFreewheelDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFreewheelDistributionProfileFilterWithDefaults() *KalturaFreewheelDistributionProfileFilter {
	this := KalturaFreewheelDistributionProfileFilter{}
	return &this
}

func (o KalturaFreewheelDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFreewheelDistributionProfileBaseFilter, errKalturaFreewheelDistributionProfileBaseFilter := json.Marshal(o.KalturaFreewheelDistributionProfileBaseFilter)
	if errKalturaFreewheelDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaFreewheelDistributionProfileBaseFilter
	}
	errKalturaFreewheelDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaFreewheelDistributionProfileBaseFilter), &toSerialize)
	if errKalturaFreewheelDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaFreewheelDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFreewheelDistributionProfileFilter struct {
	value *KalturaFreewheelDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaFreewheelDistributionProfileFilter) Get() *KalturaFreewheelDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaFreewheelDistributionProfileFilter) Set(val *KalturaFreewheelDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFreewheelDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFreewheelDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFreewheelDistributionProfileFilter(val *KalturaFreewheelDistributionProfileFilter) *NullableKalturaFreewheelDistributionProfileFilter {
	return &NullableKalturaFreewheelDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaFreewheelDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFreewheelDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



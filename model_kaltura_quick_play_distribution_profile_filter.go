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

// KalturaQuickPlayDistributionProfileFilter struct for KalturaQuickPlayDistributionProfileFilter
type KalturaQuickPlayDistributionProfileFilter struct {
	KalturaQuickPlayDistributionProfileBaseFilter
}

// NewKalturaQuickPlayDistributionProfileFilter instantiates a new KalturaQuickPlayDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuickPlayDistributionProfileFilter() *KalturaQuickPlayDistributionProfileFilter {
	this := KalturaQuickPlayDistributionProfileFilter{}
	return &this
}

// NewKalturaQuickPlayDistributionProfileFilterWithDefaults instantiates a new KalturaQuickPlayDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuickPlayDistributionProfileFilterWithDefaults() *KalturaQuickPlayDistributionProfileFilter {
	this := KalturaQuickPlayDistributionProfileFilter{}
	return &this
}

func (o KalturaQuickPlayDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaQuickPlayDistributionProfileBaseFilter, errKalturaQuickPlayDistributionProfileBaseFilter := json.Marshal(o.KalturaQuickPlayDistributionProfileBaseFilter)
	if errKalturaQuickPlayDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaQuickPlayDistributionProfileBaseFilter
	}
	errKalturaQuickPlayDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaQuickPlayDistributionProfileBaseFilter), &toSerialize)
	if errKalturaQuickPlayDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaQuickPlayDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuickPlayDistributionProfileFilter struct {
	value *KalturaQuickPlayDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaQuickPlayDistributionProfileFilter) Get() *KalturaQuickPlayDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaQuickPlayDistributionProfileFilter) Set(val *KalturaQuickPlayDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuickPlayDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuickPlayDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuickPlayDistributionProfileFilter(val *KalturaQuickPlayDistributionProfileFilter) *NullableKalturaQuickPlayDistributionProfileFilter {
	return &NullableKalturaQuickPlayDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaQuickPlayDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuickPlayDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



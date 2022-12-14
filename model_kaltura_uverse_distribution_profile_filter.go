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

// KalturaUverseDistributionProfileFilter struct for KalturaUverseDistributionProfileFilter
type KalturaUverseDistributionProfileFilter struct {
	KalturaUverseDistributionProfileBaseFilter
}

// NewKalturaUverseDistributionProfileFilter instantiates a new KalturaUverseDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUverseDistributionProfileFilter() *KalturaUverseDistributionProfileFilter {
	this := KalturaUverseDistributionProfileFilter{}
	return &this
}

// NewKalturaUverseDistributionProfileFilterWithDefaults instantiates a new KalturaUverseDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUverseDistributionProfileFilterWithDefaults() *KalturaUverseDistributionProfileFilter {
	this := KalturaUverseDistributionProfileFilter{}
	return &this
}

func (o KalturaUverseDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaUverseDistributionProfileBaseFilter, errKalturaUverseDistributionProfileBaseFilter := json.Marshal(o.KalturaUverseDistributionProfileBaseFilter)
	if errKalturaUverseDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaUverseDistributionProfileBaseFilter
	}
	errKalturaUverseDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaUverseDistributionProfileBaseFilter), &toSerialize)
	if errKalturaUverseDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaUverseDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUverseDistributionProfileFilter struct {
	value *KalturaUverseDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaUverseDistributionProfileFilter) Get() *KalturaUverseDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaUverseDistributionProfileFilter) Set(val *KalturaUverseDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUverseDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUverseDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUverseDistributionProfileFilter(val *KalturaUverseDistributionProfileFilter) *NullableKalturaUverseDistributionProfileFilter {
	return &NullableKalturaUverseDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaUverseDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUverseDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


